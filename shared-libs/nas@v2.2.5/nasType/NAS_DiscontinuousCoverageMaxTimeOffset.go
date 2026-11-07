// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// DiscontinuousCoverageMaxTimeOffset 9.11.2.5
// Iei Row, sBit, len = [], 8, 8
// Len Row, sBit, len = [], 8, 8
// UnitTimerValue Row, sBit, len = [0, 0], 8 , 3
// TimerValue Row, sBit, len = [0, 0], 5 , 5
type DiscontinuousCoverageMaxTimeOffset struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

func NewDiscontinuousCoverageMaxTimeOffset(iei uint8) (x *DiscontinuousCoverageMaxTimeOffset) {
	x = &DiscontinuousCoverageMaxTimeOffset{}
	x.SetIei(iei)
	x.Len = 1
	return x
}

func (a *DiscontinuousCoverageMaxTimeOffset) GetIei() (iei uint8)    { return a.Iei }
func (a *DiscontinuousCoverageMaxTimeOffset) SetIei(iei uint8)       { a.Iei = iei }
func (a *DiscontinuousCoverageMaxTimeOffset) GetLen() (length uint8) { return a.Len }
func (a *DiscontinuousCoverageMaxTimeOffset) SetLen(length uint8)    { a.Len = length }

func (a *DiscontinuousCoverageMaxTimeOffset) GetUnitTimerValue() uint8 {
	return a.Octet & GetBitMask(8, 5) >> 5
}

func (a *DiscontinuousCoverageMaxTimeOffset) SetUnitTimerValue(unitTimerValue uint8) {
	a.Octet = (a.Octet & 0x1F) | ((unitTimerValue & 0x07) << 5)
}

func (a *DiscontinuousCoverageMaxTimeOffset) GetTimerValue() uint8 {
	return a.Octet & GetBitMask(5, 0)
}

func (a *DiscontinuousCoverageMaxTimeOffset) SetTimerValue(timerValue uint8) {
	a.Octet = (a.Octet & 0xE0) | (timerValue & 0x1F)
}
