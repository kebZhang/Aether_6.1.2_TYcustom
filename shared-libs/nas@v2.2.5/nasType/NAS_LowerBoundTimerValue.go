// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// LowerBoundTimerValue 9.11.2.5
// UnitTimerValue Row, sBit, len = [0, 0], 8 , 3
// TimerValue Row, sBit, len = [0, 0], 5 , 5
type LowerBoundTimerValue struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

func NewLowerBoundTimerValue(iei uint8) (lowerBoundTimerValue *LowerBoundTimerValue) {
	lowerBoundTimerValue = &LowerBoundTimerValue{}
	lowerBoundTimerValue.SetIei(iei)
	return lowerBoundTimerValue
}

// LowerBoundTimerValue 9.11.2.5
// Iei Row, sBit, len = [], 8, 8
func (a *LowerBoundTimerValue) GetIei() (iei uint8) {
	return a.Iei
}

// LowerBoundTimerValue 9.11.2.5
// Iei Row, sBit, len = [], 8, 8
func (a *LowerBoundTimerValue) SetIei(iei uint8) {
	a.Iei = iei
}

// LowerBoundTimerValue 9.11.2.5
// Len Row, sBit, len = [], 8, 8
func (a *LowerBoundTimerValue) GetLen() (length uint8) {
	return a.Len
}

// LowerBoundTimerValue 9.11.2.5
// Len Row, sBit, len = [], 8, 8
func (a *LowerBoundTimerValue) SetLen(length uint8) {
	a.Len = length
}

// LowerBoundTimerValue 9.11.2.5
// UnitTimerValue Row, sBit, len = [0, 0], 8 , 3
func (a *LowerBoundTimerValue) GetUnitTimerValue() (unitTimerValue uint8) {
	return a.Octet & GetBitMask(8, 5) >> 5
}

// LowerBoundTimerValue 9.11.2.5
// UnitTimerValue Row, sBit, len = [0, 0], 8 , 3
func (a *LowerBoundTimerValue) SetUnitTimerValue(unitTimerValue uint8) {
	a.Octet = (a.Octet & 31) + ((unitTimerValue & 7) << 5)
}

// LowerBoundTimerValue 9.11.2.5
// TimerValue Row, sBit, len = [0, 0], 5 , 5
func (a *LowerBoundTimerValue) GetTimerValue() (timerValue uint8) {
	return a.Octet & GetBitMask(5, 0)
}

// LowerBoundTimerValue 9.11.2.5
// TimerValue Row, sBit, len = [0, 0], 5 , 5
func (a *LowerBoundTimerValue) SetTimerValue(timerValue uint8) {
	a.Octet = (a.Octet & 224) + (timerValue & 31)
}
