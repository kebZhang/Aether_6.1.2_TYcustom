// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// ServiceLevelAAContainer 9.11.2.10
// ServiceLevelAAContainer Row, sBit, len = [0, INF], 8 , INF
type ServiceLevelAAContainer struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewServiceLevelAAContainer(iei uint8) (serviceLevelAAContainer *ServiceLevelAAContainer) {
	serviceLevelAAContainer = &ServiceLevelAAContainer{}
	serviceLevelAAContainer.SetIei(iei)
	return serviceLevelAAContainer
}

// ServiceLevelAAContainer 9.11.2.10
// Iei Row, sBit, len = [], 8, 8
func (a *ServiceLevelAAContainer) GetIei() (iei uint8) {
	return a.Iei
}

// ServiceLevelAAContainer 9.11.2.10
// Iei Row, sBit, len = [], 8, 8
func (a *ServiceLevelAAContainer) SetIei(iei uint8) {
	a.Iei = iei
}

// ServiceLevelAAContainer 9.11.2.10
// Len Row, sBit, len = [], 8, 16
func (a *ServiceLevelAAContainer) GetLen() (length uint16) {
	return a.Len
}

// ServiceLevelAAContainer 9.11.2.10
// Len Row, sBit, len = [], 8, 16
func (a *ServiceLevelAAContainer) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

// ServiceLevelAAContainer 9.11.2.10
// ServiceLevelAAContainer Row, sBit, len = [0, INF], 8 , INF
func (a *ServiceLevelAAContainer) GetServiceLevelAAContainer() (serviceLevelAAContainer []uint8) {
	serviceLevelAAContainer = make([]uint8, len(a.Buffer))
	copy(serviceLevelAAContainer, a.Buffer)
	return serviceLevelAAContainer
}

// ServiceLevelAAContainer 9.11.2.10
// ServiceLevelAAContainer Row, sBit, len = [0, INF], 8 , INF
func (a *ServiceLevelAAContainer) SetServiceLevelAAContainer(serviceLevelAAContainer []uint8) {
	copy(a.Buffer, serviceLevelAAContainer)
}
