package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// WorkExperience holds the schema definition for the WorkExperience entity.
type WorkExperience struct {
	ent.Schema
}

// Fields of the WorkExperience.
func (WorkExperience) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now),
		field.String("title").
			NotEmpty(),
		field.String("location"),
		field.Time("start_date"),
		field.Time("end_date").
			Optional(),
		field.String("description"),
	}
}

// Edges of the WorkExperience.
func (WorkExperience) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("experiences"),
		edge.From("company", Company.Type).
			Ref("staffs"),
	}
}
