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
	InterfaceDocument = "org.a11y.atspi.Document"
)

// NewDocument creates and allocates org.a11y.atspi.Document.
func NewDocument(object dbus.BusObject) *Document {
	return &Document{object}
}

// Document implements org.a11y.atspi.Document D-Bus interface.
type Document struct {
	object dbus.BusObject
}

// GetLocale calls org.a11y.atspi.Document.GetLocale method.
func (o *Document) GetLocale(ctx context.Context) (out0 string, err error) {
	err = o.object.CallWithContext(ctx, InterfaceDocument+".GetLocale", 0).Store(&out0)
	return
}

// GetAttributeValue calls org.a11y.atspi.Document.GetAttributeValue method.
func (o *Document) GetAttributeValue(ctx context.Context, attributename string) (out0 string, err error) {
	err = o.object.CallWithContext(ctx, InterfaceDocument+".GetAttributeValue", 0, attributename).Store(&out0)
	return
}

// GetAttributes calls org.a11y.atspi.Document.GetAttributes method.
//
// Annotations:
//   @org.qtproject.QtDBus.QtTypeName.Out0 = QSpiAttributeSet
func (o *Document) GetAttributes(ctx context.Context) (out0 map[string]string, err error) {
	err = o.object.CallWithContext(ctx, InterfaceDocument+".GetAttributes", 0).Store(&out0)
	return
}

// GetCurrentPageNumber gets org.a11y.atspi.Document.CurrentPageNumber property.
func (o *Document) GetCurrentPageNumber(ctx context.Context) (currentPageNumber int32, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceDocument, "CurrentPageNumber").Store(&currentPageNumber)
	return
}

// GetPageCount gets org.a11y.atspi.Document.PageCount property.
func (o *Document) GetPageCount(ctx context.Context) (pageCount int32, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceDocument, "PageCount").Store(&pageCount)
	return
}
