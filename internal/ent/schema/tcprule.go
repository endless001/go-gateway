package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// TcpRule holds the schema definition for the TcpRule entity.
type TcpRule struct {
	ent.Schema
}

// Fields of the TcpRule.
func (TcpRule) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("service_id"),
		field.Int("Port").
			Unique(),
	}
}

// Edges of the TcpRule.
func (TcpRule) Edges() []ent.Edge {
	return nil
}
