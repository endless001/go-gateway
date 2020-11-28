package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// GrpcRule holds the schema definition for the GrpcRule entity.
type GrpcRule struct {
	ent.Schema
}

// Fields of the GrpcRule.
func (GrpcRule) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("service_id"),
		field.Int("Port"),
		field.String("header_transfor").
			Unique(),
	}
}

// Edges of the GrpcRule.
func (GrpcRule) Edges() []ent.Edge {
	return nil
}
