package schema

import "github.com/facebook/ent"

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	ent.Schema
}

// Fields of the Tenant.
func (Tenant) Fields() []ent.Field {
	return nil
}

// Edges of the Tenant.
func (Tenant) Edges() []ent.Edge {
	return nil
}
