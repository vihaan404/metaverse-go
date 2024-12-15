// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type UserRole string

const (
	UserRoleAdmin UserRole = "Admin"
	UserRoleUser  UserRole = "User"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRole: %T", src)
	}
	return nil
}

type NullUserRole struct {
	UserRole UserRole
	Valid    bool // Valid is true if UserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.UserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserRole), nil
}

type Avatar struct {
	ID       string
	ImageUrl sql.NullString
	Name     sql.NullString
}

type Element struct {
	ID       string
	Width    int32
	Height   int32
	ImageUrl string
}

type Map struct {
	ID     string
	Width  int32
	Height int32
	Name   string
}

type MapElement struct {
	ID        string
	MapID     string
	ElementID string
	X         sql.NullInt32
	Y         sql.NullInt32
}

type Space struct {
	ID        string
	Name      string
	Width     int32
	Height    sql.NullInt32
	Thumbnail sql.NullString
	CreatorID string
}

type SpaceElement struct {
	ID        string
	ElementID string
	SpaceID   string
	X         int32
	Y         int32
}

type User struct {
	ID       string
	Username string
	Password string
	AvatarID sql.NullString
	Role     UserRole
}