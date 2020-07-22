// Copyright 2017 Microsoft. All rights reserved.
// MIT License

package restserver

import (
	"github.com/Azure/azure-container-networking/cns"
	"github.com/Azure/azure-container-networking/cns/logger"
)

// This file contains the internal functions called by either HTTP APIs (api.go) or
// internal APIs (definde in internalapi.go).
// This will be used internally (say by RequestController in case of AKS)

// GetPartitionKey - Get dnc/service partition key
func (service *HTTPRestService) GetPartitionKey() (dncPartitionKey string) {
	service.RLock()
	dncPartitionKey = service.dncPartitionKey
	service.RUnlock()
	return
}

// This API will be called by CNS RequestController on CRD update.
func (service *HTTPRestService) CreateOrUpdateNetworkContainerInternal(req cns.CreateNetworkContainerRequest) int {
	if req.NetworkContainerid == "" {
		logger.Errorf("[Azure CNS] Error. NetworkContainerid is empty")
		return NetworkContainerNotSpecified
	}

	// For now only RequestController uses this API which will be initialized only for AKS scenario.
	// Validate ContainerType is set as Docker
	if service.state.OrchestratorType != cns.KubernetesCRD {
		logger.Errorf("[Azure CNS] Error. Unsupported OrchestratorType: %s", service.state.OrchestratorType)
		return UnsupportedOrchestratorType
	}

	// Validate PrimaryCA must never be empty
	err := validateIPConfig(req.IPConfiguration.IPSubnet)
	if err != nil {
		logger.Errorf("[Azure CNS] Error. PrimaryCA is invalid, NC Req: %v", req)
		return InvalidPrimaryIPConfig
	}

	// Validate SecondaryIPConfig
	for ipId, ipconfig := range req.SecondaryIPConfigs {
		// Validate Ipconfig
		err := validateIPConfig(ipconfig.IPSubnet)
		if err != nil {
			logger.Errorf("[Azure CNS] Error. SecondaryIpConfig, Id:%s is invalid, NC Req: %v", ipId, req)
			return InvalidSecndaryIPConfig
		}
	}

	// Validate if state exists already
	existing, ok := service.getNetworkContainerDetails(req.NetworkContainerid)

	if ok {
		existingReq := existing.CreateNetworkContainerRequest
		if existingReq.PrimaryInterfaceIdentifier != req.PrimaryInterfaceIdentifier {
			logger.Errorf("[Azure CNS] Error. PrimaryCA is not same, NCId %s, old CA %s, new CA %s", req.NetworkContainerid, existingReq.PrimaryInterfaceIdentifier, req.PrimaryInterfaceIdentifier)
			return PrimaryCANotSame
		}
	}

	// This will Create Or Update the NC state.
	returnCode, returnMessage := service.saveNetworkContainerGoalState(req)

	// If the NC was created successfully, log NC snapshot.
	if returnCode == 0 {
		logNCSnapshot(req)
	} else {
		logger.Errorf(returnMessage)
	}

	return returnCode
}