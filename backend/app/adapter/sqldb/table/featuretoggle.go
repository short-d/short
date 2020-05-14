package table

// FeatureToggle represents database table columns for 'feature_toggle' table.
var FeatureToggle = struct {
	TableName       string
	ColumnToggleID  string
	ColumnIsEnabled string
	ColumnType      string
}{
	TableName:       "feature_toggle",
	ColumnToggleID:  "toggle_id",
	ColumnIsEnabled: "is_enabled",
	ColumnType:      "type",
}
