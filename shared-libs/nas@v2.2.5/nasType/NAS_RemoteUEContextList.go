// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// RemoteUEContextList 9.11.4.29
// RemoteUEContextList Row, sBit, len = [0, INF], 8 , INF
type RemoteUEContextList struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewRemoteUEContextList(iei uint8) (remoteUEContextList *RemoteUEContextList) {
	remoteUEContextList = &RemoteUEContextList{}
	remoteUEContextList.SetIei(iei)
	return remoteUEContextList
}

// RemoteUEContextList 9.11.4.29
// Iei Row, sBit, len = [], 8, 8
func (a *RemoteUEContextList) GetIei() (iei uint8) {
	return a.Iei
}

// RemoteUEContextList 9.11.4.29
// Iei Row, sBit, len = [], 8, 8
func (a *RemoteUEContextList) SetIei(iei uint8) {
	a.Iei = iei
}

// RemoteUEContextList 9.11.4.29
// Len Row, sBit, len = [], 8, 16
func (a *RemoteUEContextList) GetLen() (length uint16) {
	return a.Len
}

// RemoteUEContextList 9.11.4.29
// Len Row, sBit, len = [], 8, 16
func (a *RemoteUEContextList) SetLen(length uint16) {
	a.Len = length
	a.Buffer = make([]uint8, a.Len)
}

// RemoteUEContextList 9.11.4.29
// RemoteUEContextList Row, sBit, len = [0, INF], 8 , INF
func (a *RemoteUEContextList) GetRemoteUEContextList() (remoteUEContextList []uint8) {
	remoteUEContextList = make([]uint8, len(a.Buffer))
	copy(remoteUEContextList, a.Buffer)
	return remoteUEContextList
}

// RemoteUEContextList 9.11.4.29
// RemoteUEContextList Row, sBit, len = [0, INF], 8 , INF
func (a *RemoteUEContextList) SetRemoteUEContextList(remoteUEContextList []uint8) {
	copy(a.Buffer, remoteUEContextList)
}
