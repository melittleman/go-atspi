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
	InterfaceAction = "org.a11y.atspi.Action"
)

// NewAction creates and allocates org.a11y.atspi.Action.
func NewAction(object dbus.BusObject) *Action {
	return &Action{object}
}

// Action implements org.a11y.atspi.Action D-Bus interface.
type Action struct {
	object dbus.BusObject
}

// GetDescription calls org.a11y.atspi.Action.GetDescription method.
func (o *Action) GetDescription(ctx context.Context, index int32) (out0 string, err error) {
	err = o.object.CallWithContext(ctx, InterfaceAction+".GetDescription", 0, index).Store(&out0)
	return
}

// GetName calls org.a11y.atspi.Action.GetName method.
func (o *Action) GetName(ctx context.Context, index int32) (out0 string, err error) {
	err = o.object.CallWithContext(ctx, InterfaceAction+".GetName", 0, index).Store(&out0)
	return
}

// GetLocalizedName calls org.a11y.atspi.Action.GetLocalizedName method.
func (o *Action) GetLocalizedName(ctx context.Context, index int32) (out0 string, err error) {
	err = o.object.CallWithContext(ctx, InterfaceAction+".GetLocalizedName", 0, index).Store(&out0)
	return
}

// GetKeyBinding calls org.a11y.atspi.Action.GetKeyBinding method.
func (o *Action) GetKeyBinding(ctx context.Context, index int32) (out0 string, err error) {
	err = o.object.CallWithContext(ctx, InterfaceAction+".GetKeyBinding", 0, index).Store(&out0)
	return
}

// GetActions calls org.a11y.atspi.Action.GetActions method.
//
// Annotations:
//   @org.qtproject.QtDBus.QtTypeName.Out0 = QSpiActionArray
func (o *Action) GetActions(ctx context.Context) (out0 []struct {
	V0 string
	V1 string
	V2 string
}, err error) {
	err = o.object.CallWithContext(ctx, InterfaceAction+".GetActions", 0).Store(&out0)
	return
}

// DoAction calls org.a11y.atspi.Action.DoAction method.
func (o *Action) DoAction(ctx context.Context, index int32) (out0 bool, err error) {
	err = o.object.CallWithContext(ctx, InterfaceAction+".DoAction", 0, index).Store(&out0)
	return
}

// GetNActions gets org.a11y.atspi.Action.NActions property.
func (o *Action) GetNActions(ctx context.Context) (nActions int32, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceAction, "NActions").Store(&nActions)
	return
}