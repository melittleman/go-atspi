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
	InterfaceSelection = "org.a11y.atspi.Selection"
)

// NewSelection creates and allocates org.a11y.atspi.Selection.
func NewSelection(object dbus.BusObject) *Selection {
	return &Selection{object}
}

// Selection implements org.a11y.atspi.Selection D-Bus interface.
type Selection struct {
	object dbus.BusObject
}

// GetSelectedChild calls org.a11y.atspi.Selection.GetSelectedChild method.
//
// Annotations:
//   @org.qtproject.QtDBus.QtTypeName.Out0 = QSpiObjectReference
func (o *Selection) GetSelectedChild(ctx context.Context, selectedChildIndex int32) (out0 struct {
	V0 string
	V1 dbus.ObjectPath
}, err error) {
	err = o.object.CallWithContext(ctx, InterfaceSelection+".GetSelectedChild", 0, selectedChildIndex).Store(&out0)
	return
}

// SelectChild calls org.a11y.atspi.Selection.SelectChild method.
func (o *Selection) SelectChild(ctx context.Context, childIndex int32) (out0 bool, err error) {
	err = o.object.CallWithContext(ctx, InterfaceSelection+".SelectChild", 0, childIndex).Store(&out0)
	return
}

// DeselectSelectedChild calls org.a11y.atspi.Selection.DeselectSelectedChild method.
func (o *Selection) DeselectSelectedChild(ctx context.Context, selectedChildIndex int32) (out0 bool, err error) {
	err = o.object.CallWithContext(ctx, InterfaceSelection+".DeselectSelectedChild", 0, selectedChildIndex).Store(&out0)
	return
}

// IsChildSelected calls org.a11y.atspi.Selection.IsChildSelected method.
func (o *Selection) IsChildSelected(ctx context.Context, childIndex int32) (out0 bool, err error) {
	err = o.object.CallWithContext(ctx, InterfaceSelection+".IsChildSelected", 0, childIndex).Store(&out0)
	return
}

// SelectAll calls org.a11y.atspi.Selection.SelectAll method.
func (o *Selection) SelectAll(ctx context.Context) (out0 bool, err error) {
	err = o.object.CallWithContext(ctx, InterfaceSelection+".SelectAll", 0).Store(&out0)
	return
}

// ClearSelection calls org.a11y.atspi.Selection.ClearSelection method.
func (o *Selection) ClearSelection(ctx context.Context) (out0 bool, err error) {
	err = o.object.CallWithContext(ctx, InterfaceSelection+".ClearSelection", 0).Store(&out0)
	return
}

// DeselectChild calls org.a11y.atspi.Selection.DeselectChild method.
func (o *Selection) DeselectChild(ctx context.Context, childIndex int32) (out0 bool, err error) {
	err = o.object.CallWithContext(ctx, InterfaceSelection+".DeselectChild", 0, childIndex).Store(&out0)
	return
}

// GetNSelectedChildren gets org.a11y.atspi.Selection.NSelectedChildren property.
func (o *Selection) GetNSelectedChildren(ctx context.Context) (nSelectedChildren int32, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceSelection, "NSelectedChildren").Store(&nSelectedChildren)
	return
}