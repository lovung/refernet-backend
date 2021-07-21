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
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now),
		field.String("title").MaxLen(128).NotEmpty(),
		field.String("location").MaxLen(128),
		field.Time("start_date"),
		field.Time("end_date").Optional(),
		field.String("description").MaxLen(1028),
	}
}

// Edges of the WorkExperience.
func (WorkExperience) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("experiences").Unique(),
		edge.From("in_company", Company.Type).Ref("staffs").Unique(),
		edge.From("skills", Skill.Type).Ref("experiences"),
	}
}
