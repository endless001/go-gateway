package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// HttpRule holds the schema definition for the HttpRule entity.
type HttpRule struct {
	ent.Schema
}

// Fields of the HttpRule.
func (HttpRule) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("service_id"),
		field.Int("rule_type"),
		field.String("rule"),
		field.Int("need_https"),
		field.Int("need_websocket"),
		field.Int("need_strip_uri"),
		field.String("url_rewrite"),
		field.String("header_transfor").
			Unique(),
	}
}

// Edges of the HttpRule.
func (HttpRule) Edges() []ent.Edge {
	return nil
}
