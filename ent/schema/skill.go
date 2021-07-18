package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Skill holds the schema definition for the Skill entity.
type Skill struct {
	ent.Schema
}

// Fields of the Skill.
func (Skill) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now),
		field.String("name").
			NotEmpty(),
		field.String("dark_logo_url").
			Default(""),
		field.String("light_logo_url").
			Default(""),
	}
}

// Edges of the Skill.
func (Skill) Edges() []ent.Edge {
	return nil
}
