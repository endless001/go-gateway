package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// ServiceInfo holds the schema definition for the ServiceInfo entity.
type ServiceInfo struct {
	ent.Schema
}

// Fields of the ServiceInfo.
func (ServiceInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int("load_type"),
		field.String("service_name"),
		field.String("service_desc"),
		field.Time("create_at"),
		field.Time("update_at"),
		field.Int8("is_delete").
			Unique(),
	}
}

// Edges of the ServiceInfo.
func (ServiceInfo) Edges() []ent.Edge {
	return nil
}
