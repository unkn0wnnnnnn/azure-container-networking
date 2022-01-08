package controlplane

import (
	"bytes"
	"encoding/gob"

	"github.com/Azure/azure-container-networking/npm/pkg/dataplane/policies"
)

func EncodeString(name string) (*bytes.Buffer, error) {
	var payloadBuffer *bytes.Buffer
	err := gob.NewEncoder(payloadBuffer).Encode(&name)
	if err != nil {
		return nil, err
	}
	return payloadBuffer, nil
}

func DecodeString(payload *bytes.Buffer) (string, error) {
	var name string
	err := gob.NewDecoder(payload).Decode(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func EncodeControllerIPSet(ipset *ControllerIPSets) (*bytes.Buffer, error) {
	var payloadBuffer *bytes.Buffer
	err := gob.NewEncoder(payloadBuffer).Encode(&ipset)
	if err != nil {
		return nil, err
	}
	return payloadBuffer, nil
}

func DecodeControllerIPSet(payload *bytes.Buffer) (*ControllerIPSets, error) {
	var ipset ControllerIPSets
	err := gob.NewDecoder(payload).Decode(&ipset)
	if err != nil {
		return nil, err
	}
	return &ipset, nil
}

func EncodeNPMNetworkPolicy(netpol *policies.NPMNetworkPolicy) (*bytes.Buffer, error) {
	var payloadBuffer *bytes.Buffer
	err := gob.NewEncoder(payloadBuffer).Encode(&netpol)
	if err != nil {
		return nil, err
	}
	return payloadBuffer, nil
}

func DecodeNPMNetworkPolicy(payload *bytes.Buffer) (*policies.NPMNetworkPolicy, error) {
	var netpol policies.NPMNetworkPolicy
	err := gob.NewDecoder(payload).Decode(&netpol)
	if err != nil {
		return nil, err
	}
	return &netpol, nil
}
