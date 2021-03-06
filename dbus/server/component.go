// Code generated by dbus-codegen-go DO NOT EDIT.
package server

import (
	"github.com/godbus/dbus/v5"
)

// Interface name constants.
const (
	InterfaceComponent = "org.a11y.atspi.Component"
)

// Componenter is org.a11y.atspi.Component interface.
type Componenter interface {
	// Contains is org.a11y.atspi.Component.Contains method.
	Contains(x int32, y int32, coordType uint32) (out0 bool, err *dbus.Error)
	// GetAccessibleAtPoint is org.a11y.atspi.Component.GetAccessibleAtPoint method.
	GetAccessibleAtPoint(x int32, y int32, coordType uint32) (out0 struct {
		V0 string
		V1 dbus.ObjectPath
	}, err *dbus.Error)
	// GetExtents is org.a11y.atspi.Component.GetExtents method.
	GetExtents(coordType uint32) (out0 struct {
		V0 int32
		V1 int32
		V2 int32
		V3 int32
	}, err *dbus.Error)
	// GetPosition is org.a11y.atspi.Component.GetPosition method.
	GetPosition(coordType uint32) (x int32, y int32, err *dbus.Error)
	// GetSize is org.a11y.atspi.Component.GetSize method.
	GetSize() (width int32, height int32, err *dbus.Error)
	// GetLayer is org.a11y.atspi.Component.GetLayer method.
	GetLayer() (out0 uint32, err *dbus.Error)
	// GetMDIZOrder is org.a11y.atspi.Component.GetMDIZOrder method.
	GetMDIZOrder() (out0 int16, err *dbus.Error)
	// GrabFocus is org.a11y.atspi.Component.GrabFocus method.
	GrabFocus() (out0 bool, err *dbus.Error)
	// GetAlpha is org.a11y.atspi.Component.GetAlpha method.
	GetAlpha() (out0 float64, err *dbus.Error)
	// SetExtents is org.a11y.atspi.Component.SetExtents method.
	SetExtents(x int32, y int32, width int32, height int32, coordType uint32) (out0 bool, err *dbus.Error)
	// SetPosition is org.a11y.atspi.Component.SetPosition method.
	SetPosition(x int32, y int32, coordType uint32) (out0 bool, err *dbus.Error)
	// SetSize is org.a11y.atspi.Component.SetSize method.
	SetSize(width int32, height int32) (out0 bool, err *dbus.Error)
	// ScrollTo is org.a11y.atspi.Component.ScrollTo method.
	ScrollTo(inType uint32) (out0 bool, err *dbus.Error)
	// ScrollToPoint is org.a11y.atspi.Component.ScrollToPoint method.
	ScrollToPoint(inType uint32, x int32, y int32) (out0 bool, err *dbus.Error)
}

// ExportComponent exports the given object that implements org.a11y.atspi.Component on the bus.
func ExportComponent(conn *dbus.Conn, path dbus.ObjectPath, v Componenter) error {
	return conn.ExportSubtreeMethodTable(map[string]interface{}{
		"Contains":             v.Contains,
		"GetAccessibleAtPoint": v.GetAccessibleAtPoint,
		"GetExtents":           v.GetExtents,
		"GetPosition":          v.GetPosition,
		"GetSize":              v.GetSize,
		"GetLayer":             v.GetLayer,
		"GetMDIZOrder":         v.GetMDIZOrder,
		"GrabFocus":            v.GrabFocus,
		"GetAlpha":             v.GetAlpha,
		"SetExtents":           v.SetExtents,
		"SetPosition":          v.SetPosition,
		"SetSize":              v.SetSize,
		"ScrollTo":             v.ScrollTo,
		"ScrollToPoint":        v.ScrollToPoint,
	}, path, InterfaceComponent)
}

// UnexportComponent unexports org.a11y.atspi.Component interface on the named path.
func UnexportComponent(conn *dbus.Conn, path dbus.ObjectPath) error {
	return conn.Export(nil, path, InterfaceComponent)
}

// UnimplementedComponent can be embedded to have forward compatible server implementations.
type UnimplementedComponent struct{}

func (*UnimplementedComponent) iface() string {
	return InterfaceComponent
}

func (*UnimplementedComponent) Contains(x int32, y int32, coordType uint32) (out0 bool, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedComponent) GetAccessibleAtPoint(x int32, y int32, coordType uint32) (out0 struct {
	V0 string
	V1 dbus.ObjectPath
}, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedComponent) GetExtents(coordType uint32) (out0 struct {
	V0 int32
	V1 int32
	V2 int32
	V3 int32
}, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedComponent) GetPosition(coordType uint32) (x int32, y int32, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedComponent) GetSize() (width int32, height int32, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedComponent) GetLayer() (out0 uint32, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedComponent) GetMDIZOrder() (out0 int16, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedComponent) GrabFocus() (out0 bool, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedComponent) GetAlpha() (out0 float64, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedComponent) SetExtents(x int32, y int32, width int32, height int32, coordType uint32) (out0 bool, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedComponent) SetPosition(x int32, y int32, coordType uint32) (out0 bool, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedComponent) SetSize(width int32, height int32) (out0 bool, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedComponent) ScrollTo(inType uint32) (out0 bool, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedComponent) ScrollToPoint(inType uint32, x int32, y int32) (out0 bool, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}
