package goalstateprocessor

import (
	"bytes"
	"context"
	"fmt"

	cp "github.com/Azure/azure-container-networking/npm/pkg/controlplane"
	"github.com/Azure/azure-container-networking/npm/pkg/dataplane"
	"github.com/Azure/azure-container-networking/npm/pkg/dataplane/ipsets"
	"github.com/Azure/azure-container-networking/npm/pkg/protos"

	"k8s.io/klog"
)

type GoalStateProcessor struct {
	ctx            context.Context
	cancel         context.CancelFunc
	nodeID         string
	podName        string
	controllerIP   string
	controllerPort int
	dp             dataplane.GenericDataplane
	inputChannel   chan *protos.Events
	backoffChannel chan *protos.Events
}

func NewGoalStateProcessor(
	ctx context.Context,
	nodeID string,
	podName string,
	inputChan chan *protos.Events,
	dp dataplane.GenericDataplane) *GoalStateProcessor {

	klog.Infof("Creating GoalStateProcessor for node %s", nodeID)
	return &GoalStateProcessor{
		ctx:            ctx,
		nodeID:         nodeID,
		podName:        podName,
		dp:             dp,
		inputChannel:   inputChan,
		backoffChannel: make(chan *protos.Events),
	}
}

func (gsp *GoalStateProcessor) Start() {
	klog.Infof("Starting GoalStateProcessor for node %s", gsp.nodeID)
	go gsp.run()
}

func (gsp *GoalStateProcessor) SetController(controllerIP string, controllerPort int) {
	klog.Infof("Setting controller for node %s", gsp.nodeID)
	gsp.controllerIP = controllerIP
	gsp.controllerPort = controllerPort
}

func (gsp *GoalStateProcessor) Stop() {
	klog.Infof("Stopping GoalStateProcessor for node %s", gsp.nodeID)
	gsp.cancel()
}

func (gsp *GoalStateProcessor) run() {
	if gsp.controllerIP != "" && gsp.controllerPort != 0 {
		klog.Warningf("Invaling controller for node %s, IP %s port %d", gsp.nodeID, gsp.controllerIP, gsp.controllerPort)
		return
	}
	klog.Infof("Starting dataplane for node %s", gsp.nodeID)
	for {
		select {
		case <-gsp.ctx.Done():
			klog.Infof("GoalStateProcessor for node %s stopped", gsp.nodeID)
			return
		case inputEvents := <-gsp.inputChannel:
			klog.Infof("Received event %s", inputEvents)
			gsp.process(inputEvents)
		case backoffEvents := <-gsp.backoffChannel:
			// For now keep it simple. Do not worry about backoff events
			// but if we need to handle them, we can do it here.
			klog.Infof("Received backoff event %s", backoffEvents)
			gsp.process(backoffEvents)
		}
	}
}

func (gsp *GoalStateProcessor) process(inputEvent *protos.Events) {
	klog.Infof("Processing event")

	// Process these individual buckkets in order
	// 1. Apply IPSET
	// 2. Apply POLICY
	// 3. Remove POLICY
	// 4. Remove IPSET
	payload := inputEvent.GetPayload()

	if !validatePayload(payload) {
		klog.Warningf("Empty payload in event %s", inputEvent)
		return
	}

	err := gsp.processIPSetsApplyEvent(payload[cp.IpsetApply])
	if err != nil {
		klog.Errorf("Error processing IPSET apply event %s", err)
	}

	err = gsp.processPolicyApplyEvent(payload[cp.PolicyApply])
	if err != nil {
		klog.Errorf("Error processing POLICY apply event %s", err)
	}

	err = gsp.processPolicyRemoveEvent(payload[cp.PolicyRemove])
	if err != nil {
		klog.Errorf("Error processing POLICY remove event %s", err)
	}

	err = gsp.processIPSetsRemoveEvent(payload[cp.IpsetRemove])
	if err != nil {
		klog.Errorf("Error processing IPSET remove event %s", err)
	}
}

func (gsp *GoalStateProcessor) processIPSetsApplyEvent(goalState *protos.GoalState) error {
	for _, gs := range goalState.GetData() {
		payload := bytes.NewBuffer(gs)
		ipset, err := cp.DecodeControllerIPSet(payload)
		if err != nil {
			return err
		}

		ipsetName := ipset.GetPrefixName()
		klog.Infof("Processing %s IPSET apply event", ipsetName)

		cachedIPSet := gsp.dp.GetIPSet(ipsetName)
		if cachedIPSet == nil {
			klog.Infof("IPSet %s not found in cache, adding to cache", ipsetName)
		}

		var applyErr error
		switch ipset.GetSetKind() {
		case ipsets.HashSet:
			applyErr = gsp.applySets(ipset, cachedIPSet)
		case ipsets.ListSet:
			applyErr = gsp.applyLists(ipset, cachedIPSet)
		case ipsets.UnknownKind:
			applyErr = fmt.Errorf("Unknown IPSet kind %s", cachedIPSet.Kind)
		}
		if applyErr != nil {
			return applyErr
		}
	}
	return nil
}

func (gsp *GoalStateProcessor) applySets(ipSet *cp.ControllerIPSets, cachedIPSet *ipsets.IPSet) error {
	if len(ipSet.IPPodMetadata) == 0 {
		gsp.dp.CreateIPSets([]*ipsets.IPSetMetadata{ipSet.GetMetadata()})
		return nil
	}

	setMetadata := ipSet.GetMetadata()
	for _, podMetadata := range ipSet.IPPodMetadata {
		err := gsp.dp.AddToSets([]*ipsets.IPSetMetadata{setMetadata}, podMetadata)
		if err != nil {
			return err
		}
	}

	if cachedIPSet != nil {
		toDeleteMembers := make(map[string]string)
		for podIP, podKey := range cachedIPSet.IPPodKey {
			if _, ok := ipSet.IPPodMetadata[podIP]; !ok {
				toDeleteMembers[podIP] = podKey
			}
		}

		if len(toDeleteMembers) > 0 {
			for podIP, podKey := range toDeleteMembers {
				err := gsp.dp.RemoveFromSets([]*ipsets.IPSetMetadata{setMetadata}, dataplane.NewPodMetadata(podIP, podKey, ""))
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (gsp *GoalStateProcessor) applyLists(ipSet *cp.ControllerIPSets, cachedIPSet *ipsets.IPSet) error {
	if len(ipSet.MemberIPSets) == 0 {
		gsp.dp.CreateIPSets([]*ipsets.IPSetMetadata{ipSet.GetMetadata()})
		return nil
	}

	setMetadata := ipSet.GetMetadata()
	membersToAdd := make([]*ipsets.IPSetMetadata, len(ipSet.MemberIPSets))
	idx := 0
	for _, memberIPSet := range ipSet.MemberIPSets {
		membersToAdd[idx] = memberIPSet
		idx++
	}
	err := gsp.dp.AddToLists([]*ipsets.IPSetMetadata{setMetadata}, membersToAdd)
	if err != nil {
		return err
	}

	if cachedIPSet != nil {
		toDeleteMembers := make([]*ipsets.IPSetMetadata, 0)
		for _, memberSet := range cachedIPSet.MemberIPSets {
			if _, ok := ipSet.MemberIPSets[memberSet.Name]; !ok {
				toDeleteMembers = append(toDeleteMembers, memberSet.GetSetMetadata())
			}
		}

		if len(toDeleteMembers) > 0 {
			err := gsp.dp.RemoveFromList(setMetadata, toDeleteMembers)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (gsp *GoalStateProcessor) processIPSetsRemoveEvent(goalState *protos.GoalState) error {
	for _, gs := range goalState.GetData() {
		payload := bytes.NewBuffer(gs)
		ipsetName, err := cp.DecodeString(payload)
		if err != nil {
			return err
		}
		klog.Infof("Processing %s IPSET remove event", ipsetName)

		cachedIPSet := gsp.dp.GetIPSet(ipsetName)
		if cachedIPSet == nil {
			klog.Infof("IPSet %s not found in cache, adding to cache", ipsetName)
			return nil
		}

		gsp.dp.DeleteIPSet(ipsets.NewIPSetMetadata(cachedIPSet.Name, cachedIPSet.Type))
	}
	return nil
}

func (gsp *GoalStateProcessor) processPolicyApplyEvent(goalState *protos.GoalState) error {
	for _, gs := range goalState.GetData() {
		payload := bytes.NewBuffer(gs)
		netpol, err := cp.DecodeNPMNetworkPolicy(payload)
		if err != nil {
			return err
		}
		klog.Infof("Processing %s Policy ADD event", netpol.Name)

		err = gsp.dp.UpdatePolicy(netpol)
		if err != nil {
			klog.Errorf("Error removing policy %s from dataplane %s", netpol.Name, err)
			return err
		}
	}
	return nil
}

func (gsp *GoalStateProcessor) processPolicyRemoveEvent(goalState *protos.GoalState) error {
	for _, gs := range goalState.GetData() {
		payload := bytes.NewBuffer(gs)
		netpolName, err := cp.DecodeString(payload)
		if err != nil {
			return err
		}
		klog.Infof("Processing %s Policy remove event", netpolName)

		err = gsp.dp.RemovePolicy(netpolName)
		if err != nil {
			klog.Errorf("Error removing policy %s from dataplane %s", netpolName, err)
			return err
		}
	}
	return nil
}

func validatePayload(payload map[string]*protos.GoalState) bool {
	for _, v := range payload {
		if len(v.GetData()) != 0 {
			return true
		}
	}
	return false
}
