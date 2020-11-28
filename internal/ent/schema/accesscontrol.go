package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// AccessControl holds the schema definition for the AccessControl entity.
type AccessControl struct {
	ent.Schema
}

// Fields of the AccessControl.
func (AccessControl) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("service_id"),
		field.Int("open_auth"),
		field.String("black_list"),
		field.String("white_list"),
		field.String("white_host_name"),
		field.Int("clientip_flow_limit"),
		field.Int("service_flow_limit").
			Unique(),
	}
}

// Edges of the AccessControl.
func (AccessControl) Edges() []ent.Edge {
	return nil
}
