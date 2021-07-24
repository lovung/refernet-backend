package schema

import (
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Skill holds the schema definition for the Skill entity.
type Skill struct {
	ent.Schema
}

// Fields of the Skill.
func (Skill) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entproto.Field(2)),
		field.Time("updated_at").Default(time.Now).Annotations(entproto.Field(3)),
		field.String("name").MaxLen(32).NotEmpty().Annotations(entproto.Field(4)),
		field.String("logo_url").Default("").Annotations(entproto.Field(5)),
	}
}

// Edges of the Skill.
func (Skill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("experiences", WorkExperience.Type).Annotations(entproto.Field(6)),
		edge.To("jobs", Job.Type).Annotations(entproto.Field(7)),
	}
}

// Annotations of the Skill
func (Skill) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
