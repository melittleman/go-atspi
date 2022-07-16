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
	InterfaceHyperlink = "org.a11y.atspi.Hyperlink"
)

// NewHyperlink creates and allocates org.a11y.atspi.Hyperlink.
func NewHyperlink(object dbus.BusObject) *Hyperlink {
	return &Hyperlink{object}
}

// Hyperlink implements org.a11y.atspi.Hyperlink D-Bus interface.
type Hyperlink struct {
	object dbus.BusObject
}

// GetObject calls org.a11y.atspi.Hyperlink.GetObject method.
//
// Annotations:
//   @org.qtproject.QtDBus.QtTypeName.Out0 = QSpiObjectReference
func (o *Hyperlink) GetObject(ctx context.Context, i int32) (out0 struct {
	V0 string
	V1 dbus.ObjectPath
}, err error) {
	err = o.object.CallWithContext(ctx, InterfaceHyperlink+".GetObject", 0, i).Store(&out0)
	return
}

// GetURI calls org.a11y.atspi.Hyperlink.GetURI method.
func (o *Hyperlink) GetURI(ctx context.Context, i int32) (out0 string, err error) {
	err = o.object.CallWithContext(ctx, InterfaceHyperlink+".GetURI", 0, i).Store(&out0)
	return
}

// IsValid calls org.a11y.atspi.Hyperlink.IsValid method.
func (o *Hyperlink) IsValid(ctx context.Context) (out0 bool, err error) {
	err = o.object.CallWithContext(ctx, InterfaceHyperlink+".IsValid", 0).Store(&out0)
	return
}

// GetNAnchors gets org.a11y.atspi.Hyperlink.NAnchors property.
func (o *Hyperlink) GetNAnchors(ctx context.Context) (nAnchors int16, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceHyperlink, "NAnchors").Store(&nAnchors)
	return
}

// GetStartIndex gets org.a11y.atspi.Hyperlink.StartIndex property.
func (o *Hyperlink) GetStartIndex(ctx context.Context) (startIndex int32, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceHyperlink, "StartIndex").Store(&startIndex)
	return
}

// GetEndIndex gets org.a11y.atspi.Hyperlink.EndIndex property.
func (o *Hyperlink) GetEndIndex(ctx context.Context) (endIndex int32, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceHyperlink, "EndIndex").Store(&endIndex)
	return
}