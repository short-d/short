package entity

import "database/sql"

type ToggleType = sql.NullString

var PermissionToggle = ToggleType{String: "permission", Valid: true}

// Toggle represents a controllable switch that can be turned on or off.
type Toggle struct {
	ID        string
	IsEnabled bool
	Type      ToggleType
}
