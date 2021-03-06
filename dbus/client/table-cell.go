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
	InterfaceTableCell = "org.a11y.atspi.TableCell"
)

// NewTableCell creates and allocates org.a11y.atspi.TableCell.
func NewTableCell(object dbus.BusObject) *TableCell {
	return &TableCell{object}
}

// TableCell implements org.a11y.atspi.TableCell D-Bus interface.
type TableCell struct {
	object dbus.BusObject
}

// GetRowColumnSpan calls org.a11y.atspi.TableCell.GetRowColumnSpan method.
func (o *TableCell) GetRowColumnSpan(ctx context.Context) (out0 bool, row int32, col int32, rowExtents int32, colExtents int32, err error) {
	err = o.object.CallWithContext(ctx, InterfaceTableCell+".GetRowColumnSpan", 0).Store(&out0, &row, &col, &rowExtents, &colExtents)
	return
}

// GetColumnSpan gets org.a11y.atspi.TableCell.ColumnSpan property.
func (o *TableCell) GetColumnSpan(ctx context.Context) (columnSpan int32, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceTableCell, "ColumnSpan").Store(&columnSpan)
	return
}

// GetPosition gets org.a11y.atspi.TableCell.Position property.
//
// Annotations:
//   @org.qtproject.QtDBus.QtTypeName = QPoint
func (o *TableCell) GetPosition(ctx context.Context) (position struct {
	V0 int32
	V1 int32
}, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceTableCell, "Position").Store(&position)
	return
}

// GetRowSpan gets org.a11y.atspi.TableCell.RowSpan property.
func (o *TableCell) GetRowSpan(ctx context.Context) (rowSpan int32, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceTableCell, "RowSpan").Store(&rowSpan)
	return
}

// GetTable gets org.a11y.atspi.TableCell.Table property.
//
// Annotations:
//   @org.qtproject.QtDBus.QtTypeName = QSpiObjectReference
func (o *TableCell) GetTable(ctx context.Context) (table struct {
	V0 string
	V1 dbus.ObjectPath
}, err error) {
	err = o.object.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, InterfaceTableCell, "Table").Store(&table)
	return
}
