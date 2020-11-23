package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// LoadBalance holds the schema definition for the LoadBalance entity.
type LoadBalance struct {
	ent.Schema
}

// Fields of the LoadBalance.
func (LoadBalance) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("service_id"),
		field.Int("check_method"),
		field.Int("check_timeout"),
		field.Int("check_interval"),
		field.Int("round_type"),
		field.String("ip_list"),
		field.String("weight_list"),
		field.String("forbid_list"),
		field.Int("upstream_connect_timeout"),
		field.Int("upstream_header_timeout"),
		field.Int("upstream_idle_timeout"),
		field.Int("upstream_max_idle").
			Unique(),
	}
}

// Edges of the LoadBalance.
func (LoadBalance) Edges() []ent.Edge {
	return nil
}
