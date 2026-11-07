// Copyright 2019 free5GC.org
// SPDX-FileCopyrightText: 2025 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package nasType

type Dnn []uint8

func (d *Dnn) MarshalBinary() (data []byte, err error) {
	data = append(data, uint8(len(*d)))
	data = append(data, (*d)...)

	return data, nil
}

func (d *Dnn) UnmarshalBinary(data []byte) error {
	(*d) = data[1:]
	return nil
}

// GetBitMask number, pos is shift bit
// >= lb
// < up
// TODO　exception check
func GetBitMask(ub uint8, lb uint8) (bitMask uint8) {
	bitMask = ((1<<(ub-lb) - 1) << lb)
	return bitMask
}
