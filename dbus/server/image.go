// Code generated by dbus-codegen-go DO NOT EDIT.
package server

import (
	"github.com/godbus/dbus/v5"
)

// Interface name constants.
const (
	InterfaceImage = "org.a11y.atspi.Image"
)

// Imageer is org.a11y.atspi.Image interface.
type Imageer interface {
	// GetImageExtents is org.a11y.atspi.Image.GetImageExtents method.
	GetImageExtents(coordType uint32) (out0 struct {
		V0 int32
		V1 int32
		V2 int32
		V3 int32
	}, err *dbus.Error)
	// GetImagePosition is org.a11y.atspi.Image.GetImagePosition method.
	GetImagePosition(coordType uint32) (x int32, y int32, err *dbus.Error)
	// GetImageSize is org.a11y.atspi.Image.GetImageSize method.
	GetImageSize() (width int32, height int32, err *dbus.Error)
}

// ExportImage exports the given object that implements org.a11y.atspi.Image on the bus.
func ExportImage(conn *dbus.Conn, path dbus.ObjectPath, v Imageer) error {
	return conn.ExportSubtreeMethodTable(map[string]interface{}{
		"GetImageExtents":  v.GetImageExtents,
		"GetImagePosition": v.GetImagePosition,
		"GetImageSize":     v.GetImageSize,
	}, path, InterfaceImage)
}

// UnexportImage unexports org.a11y.atspi.Image interface on the named path.
func UnexportImage(conn *dbus.Conn, path dbus.ObjectPath) error {
	return conn.Export(nil, path, InterfaceImage)
}

// UnimplementedImage can be embedded to have forward compatible server implementations.
type UnimplementedImage struct{}

func (*UnimplementedImage) iface() string {
	return InterfaceImage
}

func (*UnimplementedImage) GetImageExtents(coordType uint32) (out0 struct {
	V0 int32
	V1 int32
	V2 int32
	V3 int32
}, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedImage) GetImagePosition(coordType uint32) (x int32, y int32, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedImage) GetImageSize() (width int32, height int32, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}
