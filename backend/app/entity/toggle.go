package entity

type ToggleType string

const PermissionToggle ToggleType = "permission"

// Toggle represents a controllable switch that can be turned on or off.
type Toggle struct {
	ID        string
	IsEnabled bool
	Type      ToggleType
}
