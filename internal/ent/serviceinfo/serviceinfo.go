// Code generated by entc, DO NOT EDIT.

package serviceinfo

const (
	// Label holds the string label denoting the serviceinfo type in the database.
	Label = "service_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLoadType holds the string denoting the load_type field in the database.
	FieldLoadType = "load_type"
	// FieldServiceName holds the string denoting the service_name field in the database.
	FieldServiceName = "service_name"
	// FieldServiceDesc holds the string denoting the service_desc field in the database.
	FieldServiceDesc = "service_desc"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldIsDelete holds the string denoting the is_delete field in the database.
	FieldIsDelete = "is_delete"

	// Table holds the table name of the serviceinfo in the database.
	Table = "service_infos"
)

// Columns holds all SQL columns for serviceinfo fields.
var Columns = []string{
	FieldID,
	FieldLoadType,
	FieldServiceName,
	FieldServiceDesc,
	FieldCreateAt,
	FieldUpdateAt,
	FieldIsDelete,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
