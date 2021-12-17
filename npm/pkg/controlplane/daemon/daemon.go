package daemon

import (
	"context"

	"github.com/Azure/azure-container-networking/npm/pkg/dataplane"
	"k8s.io/klog"
)

type Daemon struct {
	ctx            context.Context
	cancel         context.CancelFunc
	nodeID         string
	podName        string
	controllerIP   string
	controllerPort int
	dp             dataplane.GenericDataplane
}

func NewDaemon(
	ctx context.Context,
	nodeID string,
	podName string,
	dp dataplane.GenericDataplane) *Daemon {

	klog.Infof("Creating daemon for node %s", nodeID)
	return &Daemon{
		ctx:     ctx,
		nodeID:  nodeID,
		podName: podName,
		dp:      dp,
	}
}

func (d *Daemon) Start() {
	klog.Infof("Starting daemon for node %s", d.nodeID)
	go d.run()
}

func (d *Daemon) SetController(controllerIP string, controllerPort int) {
	klog.Infof("Setting controller for node %s", d.nodeID)
	d.controllerIP = controllerIP
	d.controllerPort = controllerPort
}

func (d *Daemon) Stop() {
	klog.Infof("Stopping daemon for node %s", d.nodeID)
	d.cancel()
}

func (d *Daemon) run() {
	klog.Infof("Starting daemon for node %s", d.nodeID)
	for {
		select {
		case <-d.ctx.Done():
			klog.Infof("Daemon for node %s stopped", d.nodeID)
			return
		default:
			if d.controllerIP != "" && d.controllerPort != 0 {
				klog.Warningf("Invaling controller for node %s, IP %s port %d", d.nodeID, d.controllerIP, d.controllerPort)
				return
			}
			klog.Infof("Starting dataplane for node %s", d.nodeID)
		}
	}
}
