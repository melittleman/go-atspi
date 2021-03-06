// Code generated by dbus-codegen-go DO NOT EDIT.
package server

import (
	"github.com/godbus/dbus/v5"
)

// Interface name constants.
const (
	InterfaceApplication = "org.a11y.atspi.Application"
)

// Applicationer is org.a11y.atspi.Application interface.
type Applicationer interface {
	// GetLocale is org.a11y.atspi.Application.GetLocale method.
	GetLocale(lctype uint32) (out0 string, err *dbus.Error)
	// RegisterEventListener is org.a11y.atspi.Application.RegisterEventListener method.
	RegisterEventListener(event string) (err *dbus.Error)
	// DeregisterEventListener is org.a11y.atspi.Application.DeregisterEventListener method.
	DeregisterEventListener(event string) (err *dbus.Error)
}

// ExportApplication exports the given object that implements org.a11y.atspi.Application on the bus.
func ExportApplication(conn *dbus.Conn, path dbus.ObjectPath, v Applicationer) error {
	return conn.ExportSubtreeMethodTable(map[string]interface{}{
		"GetLocale":               v.GetLocale,
		"RegisterEventListener":   v.RegisterEventListener,
		"DeregisterEventListener": v.DeregisterEventListener,
	}, path, InterfaceApplication)
}

// UnexportApplication unexports org.a11y.atspi.Application interface on the named path.
func UnexportApplication(conn *dbus.Conn, path dbus.ObjectPath) error {
	return conn.Export(nil, path, InterfaceApplication)
}

// UnimplementedApplication can be embedded to have forward compatible server implementations.
type UnimplementedApplication struct{}

func (*UnimplementedApplication) iface() string {
	return InterfaceApplication
}

func (*UnimplementedApplication) GetLocale(lctype uint32) (out0 string, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedApplication) RegisterEventListener(event string) (err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedApplication) DeregisterEventListener(event string) (err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}
