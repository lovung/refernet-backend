package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Company holds the schema definition for the Company entity.
type Company struct {
	ent.Schema
}

// Fields of the Company.
func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now),
		field.String("name").NotEmpty(),
		field.Text("overview").MaxLen(1024),
		field.String("website").MaxLen(128),
		field.Strings("industries"),
		field.Strings("locations"),
		field.String("logo_url").MaxLen(128),
		field.Enum("size").
			Values("startup", "small", "medium", "big", "huge").
			Default("medium"),
		field.Int("founded_at").Positive(),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("staffs", WorkExperience.Type),
	}
}
