// Code generated by dbus-codegen-go DO NOT EDIT.
package server

import (
	"github.com/godbus/dbus/v5"
)

// Interface name constants.
const (
	InterfaceDocument = "org.a11y.atspi.Document"
)

// Documenter is org.a11y.atspi.Document interface.
type Documenter interface {
	// GetLocale is org.a11y.atspi.Document.GetLocale method.
	GetLocale() (out0 string, err *dbus.Error)
	// GetAttributeValue is org.a11y.atspi.Document.GetAttributeValue method.
	GetAttributeValue(attributename string) (out0 string, err *dbus.Error)
	// GetAttributes is org.a11y.atspi.Document.GetAttributes method.
	GetAttributes() (out0 map[string]string, err *dbus.Error)
}

// ExportDocument exports the given object that implements org.a11y.atspi.Document on the bus.
func ExportDocument(conn *dbus.Conn, path dbus.ObjectPath, v Documenter) error {
	return conn.ExportSubtreeMethodTable(map[string]interface{}{
		"GetLocale":         v.GetLocale,
		"GetAttributeValue": v.GetAttributeValue,
		"GetAttributes":     v.GetAttributes,
	}, path, InterfaceDocument)
}

// UnexportDocument unexports org.a11y.atspi.Document interface on the named path.
func UnexportDocument(conn *dbus.Conn, path dbus.ObjectPath) error {
	return conn.Export(nil, path, InterfaceDocument)
}

// UnimplementedDocument can be embedded to have forward compatible server implementations.
type UnimplementedDocument struct{}

func (*UnimplementedDocument) iface() string {
	return InterfaceDocument
}

func (*UnimplementedDocument) GetLocale() (out0 string, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedDocument) GetAttributeValue(attributename string) (out0 string, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedDocument) GetAttributes() (out0 map[string]string, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}
