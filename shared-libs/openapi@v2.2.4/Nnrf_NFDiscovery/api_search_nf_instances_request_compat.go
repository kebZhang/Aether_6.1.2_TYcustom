// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package Nnrf_NFDiscovery

import "github.com/omec-project/openapi/v2/models"

func (r ApiSearchNFInstancesRequest) GetServiceNames() *[]models.ServiceName {
	return r.serviceNames
}

func (r ApiSearchNFInstancesRequest) GetTargetPlmnList() *[]models.PlmnId {
	return r.targetPlmnList
}

func (r ApiSearchNFInstancesRequest) GetTargetNfInstanceId() *string {
	return r.targetNfInstanceId
}

func (r ApiSearchNFInstancesRequest) GetSnssais() *[]models.Snssai {
	return r.snssais
}

func (r ApiSearchNFInstancesRequest) GetDnn() *string {
	return r.dnn
}

func (r ApiSearchNFInstancesRequest) GetGuami() *models.Guami {
	return r.guami
}

func (r ApiSearchNFInstancesRequest) GetAmfRegionId() *string {
	return r.amfRegionId
}

func (r ApiSearchNFInstancesRequest) GetAmfSetId() *string {
	return r.amfSetId
}

func (r ApiSearchNFInstancesRequest) GetSupi() *string {
	return r.supi
}

func (r ApiSearchNFInstancesRequest) GetTai() *models.Tai {
	return r.tai
}
