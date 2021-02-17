package descriptor

import (
	"database/sql"
	"fmt"

	"github.com/PumpkinSeed/sqlfuzz/pkg/flags"
	"github.com/volatiletech/null"
)

// FieldDescriptor represents a field described by the table in the SQL database
type FieldDescriptor struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default null.String
	Extra   string
}

// Describe try to get the fields of the table in the SQL database
func Describe(db *sql.DB, f flags.Flags) ([]FieldDescriptor, error) {
	results, err := db.Query(fmt.Sprintf("DESCRIBE %s;", f.Table))
	if err != nil {
		return nil, err
	}

	var fields []FieldDescriptor
	for results.Next() {
		var d FieldDescriptor

		err = results.Scan(&d.Field, &d.Type, &d.Null, &d.Key, &d.Default, &d.Extra)
		if err != nil {
			return nil, err
		}

		fields = append(fields, d)
	}

	return fields, nil
}