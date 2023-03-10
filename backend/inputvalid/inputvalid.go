package inputvalid

import (
	"errors"
	"net"
	"strings"

	"github.com/AndreWongZH/iothome/models"
)

func CheckRoomInput(roomName string) error {
	if len(roomName) == 0 {
		return errors.New("cannot have empty room name")
	}

	if strings.Contains(roomName, " ") {
		return errors.New("room name must not contain spaces")
	}

	if len(roomName) > 16 {
		return errors.New("room name is too long")
	}

	return nil
}

func CheckDeviceInput(registeredDevice *models.RegisteredDevice) error {
	if len(registeredDevice.Name) == 0 || len(registeredDevice.Ipaddr) == 0 || len(registeredDevice.Type) == 0 {
		return errors.New("cannot have empty fields")
	}

	if len(registeredDevice.Name) > 15 {
		return errors.New("device name is too long")
	}

	if ip := net.ParseIP(registeredDevice.Ipaddr); ip == nil {
		return errors.New("ip is not valid")
	}

	// TODO
	// check for device type

	return nil
}
