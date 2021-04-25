package database

// Count type for queries, contains a single int64 field
type Count struct {
	C int64 `db:"c"`
}

// Helper class for Index definitions
type Index struct {
	Name       string `db:"n"`
	Definition string `db:"d"`
}
