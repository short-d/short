package table

// Progress represents database table columns for 'progress' table
var Progress = struct {
	TableName        string
	ColumnIncidentID string
	ColumnStatus     string
	ColumnInfo       string
	ColumnCreatedAt  string
}{
	TableName:        "progress",
	ColumnIncidentID: "incident_id",
	ColumnStatus:     "status",
	ColumnInfo:       "info",
	ColumnCreatedAt:  "created_at",
}
