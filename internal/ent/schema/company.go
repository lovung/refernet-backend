package schema

import (
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entproto.Field(2)),
		field.Time("updated_at").Default(time.Now).Annotations(entproto.Field(3)),
		field.String("name").NotEmpty().Annotations(entproto.Field(4)),
		field.Text("overview").MaxLen(1024).Annotations(entproto.Field(5)),
		field.String("website").MaxLen(128).Annotations(entproto.Field(6)),
		field.String("logo_url").MaxLen(128).Annotations(entproto.Field(7)),
		field.Enum("size").
			Values("startup", "small", "medium", "big", "huge").
			Default("startup").
			Annotations(
				entproto.Field(8),
				entproto.Enum(
					map[string]int32{
						"startup": 0,
						"small":   1,
						"medium":  2,
						"big":     3,
						"huge":    4,
					},
				),
			),
		field.Int("founded_at").Positive().Annotations(entproto.Field(9)),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("staffs", WorkExperience.Type).
			Annotations(entproto.Field(10)),
	}
}

// Annotations of the Company
func (Company) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
