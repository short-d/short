package table

// Incident represents database table columns for 'incident' table
var Incident = struct {
  TableName       string
  ColumnID        string
  ColumnTitle     string
  ColumnCreatedAt string
}{
  TableName:        "incident",
  ColumnID:         "id",
  ColumnTitle:      "title",
  ColumnCreatedAt:  "created_at",
}
