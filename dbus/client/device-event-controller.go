// Code generated by dbus-codegen-go DO NOT EDIT.
package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/godbus/dbus/v5"
)

// Interface name constants.
const (
	InterfaceDeviceEventController = "org.a11y.atspi.DeviceEventController"
)

// NewDeviceEventController creates and allocates org.a11y.atspi.DeviceEventController.
func NewDeviceEventController(object dbus.BusObject) *DeviceEventController {
	return &DeviceEventController{object}
}

// DeviceEventController implements org.a11y.atspi.DeviceEventController D-Bus interface.
type DeviceEventController struct {
	object dbus.BusObject
}

// RegisterKeystrokeListener calls org.a11y.atspi.DeviceEventController.RegisterKeystrokeListener method.
//
// Annotations:
//   @org.qtproject.QtDBus.QtTypeName.In1 = QSpiKeyTypeArray
//   @org.qtproject.QtDBus.QtTypeName.In3 = QSpiEventTypeArray
//   @org.qtproject.QtDBus.QtTypeName.In4 = QSpiEventMode
func (o *DeviceEventController) RegisterKeystrokeListener(ctx context.Context, listener dbus.ObjectPath, keys []struct {
	V0 int32
	V1 int32
	V2 string
	V3 int32
}, mask uint32, inType []uint32, mode struct {
	V0 bool
	V1 bool
	V2 bool
}) (out0 bool, err error) {
	err = o.object.CallWithContext(ctx, InterfaceDeviceEventController+".RegisterKeystrokeListener", 0, listener, keys, mask, inType, mode).Store(&out0)
	return
}

// DeregisterKeystrokeListener calls org.a11y.atspi.DeviceEventController.DeregisterKeystrokeListener method.
//
// Annotations:
//   @org.qtproject.QtDBus.QtTypeName.In1 = QSpiKeyTypeArray
func (o *DeviceEventController) DeregisterKeystrokeListener(ctx context.Context, listener dbus.ObjectPath, keys []struct {
	V0 int32
	V1 int32
	V2 string
	V3 int32
}, mask uint32, inType uint32) (err error) {
	err = o.object.CallWithContext(ctx, InterfaceDeviceEventController+".DeregisterKeystrokeListener", 0, listener, keys, mask, inType).Store()
	return
}

// RegisterDeviceEventListener calls org.a11y.atspi.DeviceEventController.RegisterDeviceEventListener method.
func (o *DeviceEventController) RegisterDeviceEventListener(ctx context.Context, listener dbus.ObjectPath, types uint32) (out0 bool, err error) {
	err = o.object.CallWithContext(ctx, InterfaceDeviceEventController+".RegisterDeviceEventListener", 0, listener, types).Store(&out0)
	return
}

// DeregisterDeviceEventListener calls org.a11y.atspi.DeviceEventController.DeregisterDeviceEventListener method.
func (o *DeviceEventController) DeregisterDeviceEventListener(ctx context.Context, listener dbus.ObjectPath, types uint32) (err error) {
	err = o.object.CallWithContext(ctx, InterfaceDeviceEventController+".DeregisterDeviceEventListener", 0, listener, types).Store()
	return
}

// GenerateKeyboardEvent calls org.a11y.atspi.DeviceEventController.GenerateKeyboardEvent method.
func (o *DeviceEventController) GenerateKeyboardEvent(ctx context.Context, keycode int32, keystring string, inType uint32) (err error) {
	err = o.object.CallWithContext(ctx, InterfaceDeviceEventController+".GenerateKeyboardEvent", 0, keycode, keystring, inType).Store()
	return
}

// GenerateMouseEvent calls org.a11y.atspi.DeviceEventController.GenerateMouseEvent method.
func (o *DeviceEventController) GenerateMouseEvent(ctx context.Context, x int32, y int32, eventName string) (err error) {
	err = o.object.CallWithContext(ctx, InterfaceDeviceEventController+".GenerateMouseEvent", 0, x, y, eventName).Store()
	return
}

// NotifyListenersSync calls org.a11y.atspi.DeviceEventController.NotifyListenersSync method.
//
// Annotations:
//   @org.qtproject.QtDBus.QtTypeName.In0 = QSpiDeviceEvent
func (o *DeviceEventController) NotifyListenersSync(ctx context.Context, event struct {
	V0 uint32
	V1 int32
	V2 uint32
	V3 uint32
	V4 int32
	V5 string
	V6 bool
}) (out0 bool, err error) {
	err = o.object.CallWithContext(ctx, InterfaceDeviceEventController+".NotifyListenersSync", 0, event).Store(&out0)
	return
}

// NotifyListenersAsync calls org.a11y.atspi.DeviceEventController.NotifyListenersAsync method.
//
// Annotations:
//   @org.qtproject.QtDBus.QtTypeName.In0 = QSpiDeviceEvent
func (o *DeviceEventController) NotifyListenersAsync(ctx context.Context, event struct {
	V0 uint32
	V1 int32
	V2 uint32
	V3 uint32
	V4 int32
	V5 string
	V6 bool
}) (err error) {
	err = o.object.CallWithContext(ctx, InterfaceDeviceEventController+".NotifyListenersAsync", 0, event).Store()
	return
}
