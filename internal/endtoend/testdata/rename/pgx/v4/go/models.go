// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package querytest

import (
	"database/sql/driver"
	"fmt"
)

type IPProtocol string

const (
	IPProtocolTCP  IPProtocol = "tcp"
	IpProtocolIp   IPProtocol = "ip"
	IpProtocolIcmp IPProtocol = "icmp"
)

func (e *IPProtocol) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = IPProtocol(s)
	case string:
		*e = IPProtocol(s)
	default:
		return fmt.Errorf("unsupported scan type for IPProtocol: %T", src)
	}
	return nil
}

type NullIPProtocol struct {
	IPProtocol IPProtocol
	Valid      bool // Valid is true if IPProtocol is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullIPProtocol) Scan(value interface{}) error {
	if value == nil {
		ns.IPProtocol, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.IPProtocol.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullIPProtocol) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.IPProtocol, nil
}

type BarNew struct {
	IDNew int32
	IpOld IPProtocol
}