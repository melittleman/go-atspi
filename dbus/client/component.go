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
	InterfaceComponent = "org.a11y.atspi.Component"
)

// NewComponent creates and allocates org.a11y.atspi.Component.
func NewComponent(object dbus.BusObject) *Component {
	return &Component{object}
}

// Component implements org.a11y.atspi.Component D-Bus interface.
type Component struct {
	object dbus.BusObject
}

// Contains calls org.a11y.atspi.Component.Contains method.
func (o *Component) Contains(ctx context.Context, x int32, y int32, coordType uint32) (out0 bool, err error) {
	err = o.object.CallWithContext(ctx, InterfaceComponent+".Contains", 0, x, y, coordType).Store(&out0)
	return
}

// GetAccessibleAtPoint calls org.a11y.atspi.Component.GetAccessibleAtPoint method.
//
// Annotations:
//   @org.qtproject.QtDBus.QtTypeName.Out0 = QSpiObjectReference
func (o *Component) GetAccessibleAtPoint(ctx context.Context, x int32, y int32, coordType uint32) (out0 struct {
	V0 string
	V1 dbus.ObjectPath
}, err error) {
	err = o.object.CallWithContext(ctx, InterfaceComponent+".GetAccessibleAtPoint", 0, x, y, coordType).Store(&out0)
	return
}

// GetExtents calls org.a11y.atspi.Component.GetExtents method.
//
// Annotations:
//   @org.qtproject.QtDBus.QtTypeName.Out0 = QSpiRect
func (o *Component) GetExtents(ctx context.Context, coordType uint32) (out0 struct {
	V0 int32
	V1 int32
	V2 int32
	V3 int32
}, err error) {
	err = o.object.CallWithContext(ctx, InterfaceComponent+".GetExtents", 0, coordType).Store(&out0)
	return
}

// GetPosition calls org.a11y.atspi.Component.GetPosition method.
func (o *Component) GetPosition(ctx context.Context, coordType uint32) (x int32, y int32, err error) {
	err = o.object.CallWithContext(ctx, InterfaceComponent+".GetPosition", 0, coordType).Store(&x, &y)
	return
}

// GetSize calls org.a11y.atspi.Component.GetSize method.
func (o *Component) GetSize(ctx context.Context) (width int32, height int32, err error) {
	err = o.object.CallWithContext(ctx, InterfaceComponent+".GetSize", 0).Store(&width, &height)
	return
}

// GetLayer calls org.a11y.atspi.Component.GetLayer method.
func (o *Component) GetLayer(ctx context.Context) (out0 uint32, err error) {
	err = o.object.CallWithContext(ctx, InterfaceComponent+".GetLayer", 0).Store(&out0)
	return
}

// GetMDIZOrder calls org.a11y.atspi.Component.GetMDIZOrder method.
func (o *Component) GetMDIZOrder(ctx context.Context) (out0 int16, err error) {
	err = o.object.CallWithContext(ctx, InterfaceComponent+".GetMDIZOrder", 0).Store(&out0)
	return
}

// GrabFocus calls org.a11y.atspi.Component.GrabFocus method.
func (o *Component) GrabFocus(ctx context.Context) (out0 bool, err error) {
	err = o.object.CallWithContext(ctx, InterfaceComponent+".GrabFocus", 0).Store(&out0)
	return
}

// GetAlpha calls org.a11y.atspi.Component.GetAlpha method.
func (o *Component) GetAlpha(ctx context.Context) (out0 float64, err error) {
	err = o.object.CallWithContext(ctx, InterfaceComponent+".GetAlpha", 0).Store(&out0)
	return
}

// SetExtents calls org.a11y.atspi.Component.SetExtents method.
func (o *Component) SetExtents(ctx context.Context, x int32, y int32, width int32, height int32, coordType uint32) (out0 bool, err error) {
	err = o.object.CallWithContext(ctx, InterfaceComponent+".SetExtents", 0, x, y, width, height, coordType).Store(&out0)
	return
}

// SetPosition calls org.a11y.atspi.Component.SetPosition method.
func (o *Component) SetPosition(ctx context.Context, x int32, y int32, coordType uint32) (out0 bool, err error) {
	err = o.object.CallWithContext(ctx, InterfaceComponent+".SetPosition", 0, x, y, coordType).Store(&out0)
	return
}

// SetSize calls org.a11y.atspi.Component.SetSize method.
func (o *Component) SetSize(ctx context.Context, width int32, height int32) (out0 bool, err error) {
	err = o.object.CallWithContext(ctx, InterfaceComponent+".SetSize", 0, width, height).Store(&out0)
	return
}

// ScrollTo calls org.a11y.atspi.Component.ScrollTo method.
func (o *Component) ScrollTo(ctx context.Context, inType uint32) (out0 bool, err error) {
	err = o.object.CallWithContext(ctx, InterfaceComponent+".ScrollTo", 0, inType).Store(&out0)
	return
}

// ScrollToPoint calls org.a11y.atspi.Component.ScrollToPoint method.
func (o *Component) ScrollToPoint(ctx context.Context, inType uint32, x int32, y int32) (out0 bool, err error) {
	err = o.object.CallWithContext(ctx, InterfaceComponent+".ScrollToPoint", 0, inType, x, y).Store(&out0)
	return
}
