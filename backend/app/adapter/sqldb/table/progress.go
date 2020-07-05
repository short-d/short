package table

// Progress represents database table columns for 'progress' table
var Progress = struct {
  TableName           string
  ColumnIncidentId    string
  ColumnStatus        string
  ColumnInfo          string
  ColumnCreatedAt     string
}{
  TableName:          "progress",
  ColumnIncidentId:   "incident_id",
  ColumnStatus:       "status",
  ColumnInfo:         "info",
  ColumnCreatedAt:    "created_at",
}
