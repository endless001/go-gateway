package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// App holds the schema definition for the App entity.
type App struct {
	ent.Schema
}

// Fields of the App.
func (App) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("app_id"),
		field.String("name"),
		field.String("secret"),
		field.String("white_ips"),
		field.Int("qpd"),
		field.Int("Qps"),
		field.Time("create_at"),
		field.Time("update_at"),
		field.Int8("is_delete").
			Unique(),
	}
}

// Edges of the App.
func (App) Edges() []ent.Edge {
	return nil
}
