package table

var FeatureToggle = struct {
	TableName       string
	ColumnToggleID  string
	ColumnIsEnabled string
}{
	TableName:       "feature_toggle",
	ColumnToggleID:  "toggle_id",
	ColumnIsEnabled: "is_enabled",
}
