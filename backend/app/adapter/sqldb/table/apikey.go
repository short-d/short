package table

var ApiKey = struct {
	TableName       string
	ColumnAppID     string
	ColumnKey       string
	ColumnDisabled  string
	ColumnCreatedAt string
}{
	TableName:       "api_key",
	ColumnAppID:     "app_id",
	ColumnKey:       "key",
	ColumnDisabled:  "disabled",
	ColumnCreatedAt: "created_at",
}
