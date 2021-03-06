// Code generated by dbus-codegen-go DO NOT EDIT.
package server

import (
	"github.com/godbus/dbus/v5"
)

// Interface name constants.
const (
	InterfaceTable = "org.a11y.atspi.Table"
)

// Tableer is org.a11y.atspi.Table interface.
type Tableer interface {
	// GetAccessibleAt is org.a11y.atspi.Table.GetAccessibleAt method.
	GetAccessibleAt(row int32, column int32) (out0 struct {
		V0 string
		V1 dbus.ObjectPath
	}, err *dbus.Error)
	// GetIndexAt is org.a11y.atspi.Table.GetIndexAt method.
	GetIndexAt(row int32, column int32) (out0 int32, err *dbus.Error)
	// GetRowAtIndex is org.a11y.atspi.Table.GetRowAtIndex method.
	GetRowAtIndex(index int32) (out0 int32, err *dbus.Error)
	// GetColumnAtIndex is org.a11y.atspi.Table.GetColumnAtIndex method.
	GetColumnAtIndex(index int32) (out0 int32, err *dbus.Error)
	// GetRowDescription is org.a11y.atspi.Table.GetRowDescription method.
	GetRowDescription(row int32) (out0 string, err *dbus.Error)
	// GetColumnDescription is org.a11y.atspi.Table.GetColumnDescription method.
	GetColumnDescription(column int32) (out0 string, err *dbus.Error)
	// GetRowExtentAt is org.a11y.atspi.Table.GetRowExtentAt method.
	GetRowExtentAt(row int32, column int32) (out0 int32, err *dbus.Error)
	// GetColumnExtentAt is org.a11y.atspi.Table.GetColumnExtentAt method.
	GetColumnExtentAt(row int32, column int32) (out0 int32, err *dbus.Error)
	// GetRowHeader is org.a11y.atspi.Table.GetRowHeader method.
	GetRowHeader(row int32) (out0 struct {
		V0 string
		V1 dbus.ObjectPath
	}, err *dbus.Error)
	// GetColumnHeader is org.a11y.atspi.Table.GetColumnHeader method.
	GetColumnHeader(column int32) (out0 struct {
		V0 string
		V1 dbus.ObjectPath
	}, err *dbus.Error)
	// GetSelectedRows is org.a11y.atspi.Table.GetSelectedRows method.
	GetSelectedRows() (out0 []int32, err *dbus.Error)
	// GetSelectedColumns is org.a11y.atspi.Table.GetSelectedColumns method.
	GetSelectedColumns() (out0 []int32, err *dbus.Error)
	// IsRowSelected is org.a11y.atspi.Table.IsRowSelected method.
	IsRowSelected(row int32) (out0 bool, err *dbus.Error)
	// IsColumnSelected is org.a11y.atspi.Table.IsColumnSelected method.
	IsColumnSelected(column int32) (out0 bool, err *dbus.Error)
	// IsSelected is org.a11y.atspi.Table.IsSelected method.
	IsSelected(row int32, column int32) (out0 bool, err *dbus.Error)
	// AddRowSelection is org.a11y.atspi.Table.AddRowSelection method.
	AddRowSelection(row int32) (out0 bool, err *dbus.Error)
	// AddColumnSelection is org.a11y.atspi.Table.AddColumnSelection method.
	AddColumnSelection(column int32) (out0 bool, err *dbus.Error)
	// RemoveRowSelection is org.a11y.atspi.Table.RemoveRowSelection method.
	RemoveRowSelection(row int32) (out0 bool, err *dbus.Error)
	// RemoveColumnSelection is org.a11y.atspi.Table.RemoveColumnSelection method.
	RemoveColumnSelection(column int32) (out0 bool, err *dbus.Error)
	// GetRowColumnExtentsAtIndex is org.a11y.atspi.Table.GetRowColumnExtentsAtIndex method.
	GetRowColumnExtentsAtIndex(index int32) (out0 bool, row int32, col int32, rowExtents int32, colExtents int32, isSelected bool, err *dbus.Error)
}

// ExportTable exports the given object that implements org.a11y.atspi.Table on the bus.
func ExportTable(conn *dbus.Conn, path dbus.ObjectPath, v Tableer) error {
	return conn.ExportSubtreeMethodTable(map[string]interface{}{
		"GetAccessibleAt":            v.GetAccessibleAt,
		"GetIndexAt":                 v.GetIndexAt,
		"GetRowAtIndex":              v.GetRowAtIndex,
		"GetColumnAtIndex":           v.GetColumnAtIndex,
		"GetRowDescription":          v.GetRowDescription,
		"GetColumnDescription":       v.GetColumnDescription,
		"GetRowExtentAt":             v.GetRowExtentAt,
		"GetColumnExtentAt":          v.GetColumnExtentAt,
		"GetRowHeader":               v.GetRowHeader,
		"GetColumnHeader":            v.GetColumnHeader,
		"GetSelectedRows":            v.GetSelectedRows,
		"GetSelectedColumns":         v.GetSelectedColumns,
		"IsRowSelected":              v.IsRowSelected,
		"IsColumnSelected":           v.IsColumnSelected,
		"IsSelected":                 v.IsSelected,
		"AddRowSelection":            v.AddRowSelection,
		"AddColumnSelection":         v.AddColumnSelection,
		"RemoveRowSelection":         v.RemoveRowSelection,
		"RemoveColumnSelection":      v.RemoveColumnSelection,
		"GetRowColumnExtentsAtIndex": v.GetRowColumnExtentsAtIndex,
	}, path, InterfaceTable)
}

// UnexportTable unexports org.a11y.atspi.Table interface on the named path.
func UnexportTable(conn *dbus.Conn, path dbus.ObjectPath) error {
	return conn.Export(nil, path, InterfaceTable)
}

// UnimplementedTable can be embedded to have forward compatible server implementations.
type UnimplementedTable struct{}

func (*UnimplementedTable) iface() string {
	return InterfaceTable
}

func (*UnimplementedTable) GetAccessibleAt(row int32, column int32) (out0 struct {
	V0 string
	V1 dbus.ObjectPath
}, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedTable) GetIndexAt(row int32, column int32) (out0 int32, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedTable) GetRowAtIndex(index int32) (out0 int32, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedTable) GetColumnAtIndex(index int32) (out0 int32, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedTable) GetRowDescription(row int32) (out0 string, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedTable) GetColumnDescription(column int32) (out0 string, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedTable) GetRowExtentAt(row int32, column int32) (out0 int32, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedTable) GetColumnExtentAt(row int32, column int32) (out0 int32, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedTable) GetRowHeader(row int32) (out0 struct {
	V0 string
	V1 dbus.ObjectPath
}, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedTable) GetColumnHeader(column int32) (out0 struct {
	V0 string
	V1 dbus.ObjectPath
}, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedTable) GetSelectedRows() (out0 []int32, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedTable) GetSelectedColumns() (out0 []int32, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedTable) IsRowSelected(row int32) (out0 bool, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedTable) IsColumnSelected(column int32) (out0 bool, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedTable) IsSelected(row int32, column int32) (out0 bool, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedTable) AddRowSelection(row int32) (out0 bool, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedTable) AddColumnSelection(column int32) (out0 bool, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedTable) RemoveRowSelection(row int32) (out0 bool, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedTable) RemoveColumnSelection(column int32) (out0 bool, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}

func (*UnimplementedTable) GetRowColumnExtentsAtIndex(index int32) (out0 bool, row int32, col int32, rowExtents int32, colExtents int32, isSelected bool, err *dbus.Error) {
	err = &dbus.ErrMsgUnknownMethod
	return
}
