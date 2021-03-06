// Copyright 2014, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package accessory

import (
	"github.com/kylelemons/gousb/usb"
)

func getProtocol(dev *usb.Device) (uint16, error) {
	if dev == nil {
		return 0, ErrorNoAccessoryDevice
	}

	var data = make([]byte, 2)
	n, err := dev.Control(RTYPE_IN, GET_PROTOCOL, 0, 0, data)
	if err != nil {
		return 0, err
	}
	if n != 2 {
		return 0, ErrorFailedToGetProtocol
	}

	return (uint16(data[1])<<8 | uint16(data[0])), nil
}

func sendString(dev *usb.Device, idx uint16, str string) error {
	if dev == nil {
		return ErrorNoAccessoryDevice
	}

	data := []byte(str + "\x00")
	_, err := dev.Control(RTYPE_OUT, SEND_STRING, 0, idx, data)
	return err
}

func start(dev *usb.Device) error {
	if dev == nil {
		return ErrorNoAccessoryDevice
	}

	_, err := dev.Control(RTYPE_OUT, START, 0, 0, nil)
	return err
}
