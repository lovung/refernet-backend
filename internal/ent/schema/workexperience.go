package schema

import (
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entproto.Field(2)),
		field.Time("updated_at").Default(time.Now).Annotations(entproto.Field(3)),
		field.String("title").MaxLen(128).NotEmpty().Annotations(entproto.Field(4)),
		field.String("location").MaxLen(128).Annotations(entproto.Field(5)),
		field.Time("start_date").Annotations(entproto.Field(6)),
		field.Time("end_date").Optional().Annotations(entproto.Field(7)),
		field.String("description").MaxLen(1028).Annotations(entproto.Field(8)),
	}
}

// Edges of the WorkExperience.
func (WorkExperience) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("experiences").Unique().Annotations(entproto.Field(9)),
		edge.From("in_company", Company.Type).Ref("staffs").Unique().Annotations(entproto.Field(10)),
		edge.From("skills", Skill.Type).Ref("experiences").Annotations(entproto.Field(11)),
	}
}

// Annotations of the WorkExperience
func (WorkExperience) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
