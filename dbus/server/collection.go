// Code generated by dbus-codegen-go DO NOT EDIT.
package server

import (
	"github.com/godbus/dbus/v5"
)

// Interface name constants.
const (
	InterfaceCollection = "org.a11y.atspi.Collection"
)

// Collectioner is org.a11y.atspi.Collection interface.
type Collectioner interface {
	// GetMatches is org.a11y.atspi.Collection.GetMatches method.
	GetMatches(rule struct {
		V0 []int32
		V1 int32
		V2 map[string]string
		V3 int32
		V4 []int32
		V5 int32
		V6 []string
		V7 int32
		V8 bool
	}, sortby uint32, count int32, traverse bool) (out0 []struct {
		V0 string
		V1 dbus.ObjectPath
	}, err *dbus.Error)
	// GetMatchesTo is org.a11y.atspi.Collection.GetMatchesTo method.
	GetMatchesTo(currentObject dbus.ObjectPath, rule struct {
		V0 []int32
		V1 int32
		V2 map[string]string
		V3 int32
		V4 []int32
		V5 int32
		V6 []string
		V7 int32
		V8 bool
	}, sortby uint32, tree uint32, limitScope bool, count int32, traverse bool) (out0 []struct {
		V0 string
		V1 dbus.ObjectPath
	}, err *dbus.Error)
	// GetMatchesFrom is org.a11y.atspi.Collection.GetMatchesFrom method.
	GetMatchesFrom(currentObject dbus.ObjectPath, rule struct {
		V0 []int32
		V1 int32
		V2 map[string]string
		V3 int32
		V4 []int32
		V5 int32
		V6 []string
		V7 int32
		V8 bool
	}, sortby uint32, tree uint32, count int32, traverse bool) (out0 []struct {
		V0 string
		V1 dbus.ObjectPath
	}, err *dbus.Error)
	// GetActiveDescendant is org.a11y.atspi.Collection.GetActiveDescendant method.
	GetActiveDescendant() (out0 struct {
		V0 string
		V1 dbus.ObjectPath
	}, err *dbus.Error)
}

// ExportCollection exports the given object that implements org.a11y.atspi.Collection on the bus.
func ExportCollection(conn *dbus.Conn, path dbus.ObjectPath, v Collectioner) error {
	return conn.ExportSubtreeMethodTable(map[string]interface{}{
		"GetMatches":          v.GetMatches,
		"GetMatchesTo":        v.GetMatchesTo,
		"GetMatchesFrom":      v.GetMatchesFrom,
		"GetActiveDescendant": v.GetActiveDescendant,
	}, path, InterfaceCollection)
}

// UnexportCollection unexports org.a11y.atspi.Collection interface on the named path.
func UnexportCollection(conn *dbus.Conn, path dbus.ObjectPath) error {
	return conn.Export(nil, path, InterfaceCollection)
}

// UnimplementedCollection can be embedded to have forward compatible server implementations.
type UnimplementedCollection struct{}

func (*UnimplementedCollection) iface() string {
	return InterfaceCollection
}

func (*UnimplementedCollection) GetMatches(rule struct {
	V0 []int32
	V1 int32
	V2 map[string]string
	V3 int32
	V4 []int32
	V5 int32
	V6 []string
	V7 int32
	V8 bool
}, sortby uint32, count int32, traverse bool) (out0 []struct {
	V0 string
	V1 dbus.ObjectPath
}, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedCollection) GetMatchesTo(currentObject dbus.ObjectPath, rule struct {
	V0 []int32
	V1 int32
	V2 map[string]string
	V3 int32
	V4 []int32
	V5 int32
	V6 []string
	V7 int32
	V8 bool
}, sortby uint32, tree uint32, limitScope bool, count int32, traverse bool) (out0 []struct {
	V0 string
	V1 dbus.ObjectPath
}, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedCollection) GetMatchesFrom(currentObject dbus.ObjectPath, rule struct {
	V0 []int32
	V1 int32
	V2 map[string]string
	V3 int32
	V4 []int32
	V5 int32
	V6 []string
	V7 int32
	V8 bool
}, sortby uint32, tree uint32, count int32, traverse bool) (out0 []struct {
	V0 string
	V1 dbus.ObjectPath
}, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedCollection) GetActiveDescendant() (out0 struct {
	V0 string
	V1 dbus.ObjectPath
}, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}
