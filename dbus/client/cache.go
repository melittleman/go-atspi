// Code generated by dbus-codegen-go DO NOT EDIT.
package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/godbus/dbus/v5"
)

// Signal is a common interface for all signals.
type Signal interface {
	Name() string
	Interface() string
	Sender() string

	path() dbus.ObjectPath
	values() []interface{}
}

// ErrUnknownSignal is returned by LookupSignal when a signal cannot be resolved.
var ErrUnknownSignal = errors.New("unknown signal")

// LookupSignal converts the given raw D-Bus signal with variable body
// into one with typed structured body or returns ErrUnknownSignal error.
func LookupSignal(signal *dbus.Signal) (Signal, error) {
	switch signal.Name {
	case InterfaceCache + "." + "AddAccessible":
		v0, ok := signal.Body[0].(struct {
			V0 struct {
				V0 string
				V1 dbus.ObjectPath
			}
			V1 struct {
				V0 string
				V1 dbus.ObjectPath
			}
			V2 struct {
				V0 string
				V1 dbus.ObjectPath
			}
			V3 int32
			V4 int32
			V5 []string
			V6 string
			V7 uint32
			V8 string
			V9 []uint32
		})
		if !ok {
			return nil, fmt.Errorf("prop .NodeAdded is %T, not struct {V0 struct {V0 string;V1 dbus.ObjectPath};V1 struct {V0 string;V1 dbus.ObjectPath};V2 struct {V0 string;V1 dbus.ObjectPath};V3 int32;V4 int32;V5 []string;V6 string;V7 uint32;V8 string;V9 []uint32}", signal.Body[0])
		}
		return &CacheAddAccessibleSignal{
			sender: signal.Sender,
			Path:   signal.Path,
			Body: &CacheAddAccessibleSignalBody{
				NodeAdded: v0,
			},
		}, nil
	case InterfaceCache + "." + "RemoveAccessible":
		v0, ok := signal.Body[0].(struct {
			V0 string
			V1 dbus.ObjectPath
		})
		if !ok {
			return nil, fmt.Errorf("prop .NodeRemoved is %T, not struct {V0 string;V1 dbus.ObjectPath}", signal.Body[0])
		}
		return &CacheRemoveAccessibleSignal{
			sender: signal.Sender,
			Path:   signal.Path,
			Body: &CacheRemoveAccessibleSignalBody{
				NodeRemoved: v0,
			},
		}, nil
	default:
		return nil, ErrUnknownSignal
	}
}

// AddMatchSignal registers a match rule for the given signal,
// opts are appended to the automatically generated signal's rules.
func AddMatchSignal(conn *dbus.Conn, s Signal, opts ...dbus.MatchOption) error {
	return conn.AddMatchSignal(append([]dbus.MatchOption{
		dbus.WithMatchInterface(s.Interface()),
		dbus.WithMatchMember(s.Name()),
	}, opts...)...)
}

// RemoveMatchSignal unregisters the previously registered subscription.
func RemoveMatchSignal(conn *dbus.Conn, s Signal, opts ...dbus.MatchOption) error {
	return conn.RemoveMatchSignal(append([]dbus.MatchOption{
		dbus.WithMatchInterface(s.Interface()),
		dbus.WithMatchMember(s.Name()),
	}, opts...)...)
}

// Interface name constants.
const (
	InterfaceCache = "org.a11y.atspi.Cache"
)

// NewCache creates and allocates org.a11y.atspi.Cache.
func NewCache(object dbus.BusObject) *Cache {
	return &Cache{object}
}

// Cache implements org.a11y.atspi.Cache D-Bus interface.
type Cache struct {
	object dbus.BusObject
}

// GetItems calls org.a11y.atspi.Cache.GetItems method.
//
// Annotations:
//   @org.qtproject.QtDBus.QtTypeName.Out0 = QSpiAccessibleCacheArray
func (o *Cache) GetItems(ctx context.Context) (nodes []struct {
	V0 struct {
		V0 string
		V1 dbus.ObjectPath
	}
	V1 struct {
		V0 string
		V1 dbus.ObjectPath
	}
	V2 struct {
		V0 string
		V1 dbus.ObjectPath
	}
	V3 int32
	V4 int32
	V5 []string
	V6 string
	V7 uint32
	V8 string
	V9 []uint32
}, err error) {
	err = o.object.CallWithContext(ctx, InterfaceCache+".GetItems", 0).Store(&nodes)
	return
}

// CacheAddAccessibleSignal represents org.a11y.atspi.Cache.AddAccessible signal.
//
// Annotations:
//   @org.qtproject.QtDBus.QtTypeName.In0 = QSpiAccessibleCacheItem
type CacheAddAccessibleSignal struct {
	sender string
	Path   dbus.ObjectPath
	Body   *CacheAddAccessibleSignalBody
}

// Name returns the signal's name.
func (s *CacheAddAccessibleSignal) Name() string {
	return "AddAccessible"
}

// Interface returns the signal's interface.
func (s *CacheAddAccessibleSignal) Interface() string {
	return InterfaceCache
}

// Sender returns the signal's sender unique name.
func (s *CacheAddAccessibleSignal) Sender() string {
	return s.sender
}

func (s *CacheAddAccessibleSignal) path() dbus.ObjectPath {
	return s.Path
}

func (s *CacheAddAccessibleSignal) values() []interface{} {
	return []interface{}{s.Body.NodeAdded}
}

// CacheAddAccessibleSignalBody is body container.
type CacheAddAccessibleSignalBody struct {
	NodeAdded struct {
		V0 struct {
			V0 string
			V1 dbus.ObjectPath
		}
		V1 struct {
			V0 string
			V1 dbus.ObjectPath
		}
		V2 struct {
			V0 string
			V1 dbus.ObjectPath
		}
		V3 int32
		V4 int32
		V5 []string
		V6 string
		V7 uint32
		V8 string
		V9 []uint32
	}
}

// CacheRemoveAccessibleSignal represents org.a11y.atspi.Cache.RemoveAccessible signal.
//
// Annotations:
//   @org.qtproject.QtDBus.QtTypeName.In0 = QSpiObjectReference
type CacheRemoveAccessibleSignal struct {
	sender string
	Path   dbus.ObjectPath
	Body   *CacheRemoveAccessibleSignalBody
}

// Name returns the signal's name.
func (s *CacheRemoveAccessibleSignal) Name() string {
	return "RemoveAccessible"
}

// Interface returns the signal's interface.
func (s *CacheRemoveAccessibleSignal) Interface() string {
	return InterfaceCache
}

// Sender returns the signal's sender unique name.
func (s *CacheRemoveAccessibleSignal) Sender() string {
	return s.sender
}

func (s *CacheRemoveAccessibleSignal) path() dbus.ObjectPath {
	return s.Path
}

func (s *CacheRemoveAccessibleSignal) values() []interface{} {
	return []interface{}{s.Body.NodeRemoved}
}

// CacheRemoveAccessibleSignalBody is body container.
type CacheRemoveAccessibleSignalBody struct {
	NodeRemoved struct {
		V0 string
		V1 dbus.ObjectPath
	}
}