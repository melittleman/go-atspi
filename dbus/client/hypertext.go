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
	InterfaceHypertext = "org.a11y.atspi.Hypertext"
)

// NewHypertext creates and allocates org.a11y.atspi.Hypertext.
func NewHypertext(object dbus.BusObject) *Hypertext {
	return &Hypertext{object}
}

// Hypertext implements org.a11y.atspi.Hypertext D-Bus interface.
type Hypertext struct {
	object dbus.BusObject
}

// GetNLinks calls org.a11y.atspi.Hypertext.GetNLinks method.
func (o *Hypertext) GetNLinks(ctx context.Context) (out0 int32, err error) {
	err = o.object.CallWithContext(ctx, InterfaceHypertext+".GetNLinks", 0).Store(&out0)
	return
}

// GetLink calls org.a11y.atspi.Hypertext.GetLink method.
//
// Annotations:
//   @org.qtproject.QtDBus.QtTypeName.Out0 = QSpiObjectReference
func (o *Hypertext) GetLink(ctx context.Context, linkIndex int32) (out0 struct {
	V0 string
	V1 dbus.ObjectPath
}, err error) {
	err = o.object.CallWithContext(ctx, InterfaceHypertext+".GetLink", 0, linkIndex).Store(&out0)
	return
}

// GetLinkIndex calls org.a11y.atspi.Hypertext.GetLinkIndex method.
func (o *Hypertext) GetLinkIndex(ctx context.Context, characterIndex int32) (out0 int32, err error) {
	err = o.object.CallWithContext(ctx, InterfaceHypertext+".GetLinkIndex", 0, characterIndex).Store(&out0)
	return
}