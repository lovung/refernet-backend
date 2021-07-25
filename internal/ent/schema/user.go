package schema

import (
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entproto.Field(2)),
		field.Time("updated_at").Default(time.Now).Annotations(entproto.Field(3)),
		field.String("username").MaxLen(128).NotEmpty().Annotations(entproto.Field(4)),
		field.String("fullname").MaxLen(128).NotEmpty().Annotations(entproto.Field(5)),
		field.String("password").MaxLen(256).NotEmpty().Annotations(entproto.Field(6)),
		field.String("email").MaxLen(128).NotEmpty().Annotations(entproto.Field(7)),
		field.String("phone").MaxLen(20).Annotations(entproto.Field(8)),
		field.String("bio").MaxLen(64).Annotations(entproto.Field(9)),
		field.Text("intro").MaxLen(1024).Annotations(entproto.Field(10)),
		field.String("github_profile").MaxLen(128).Annotations(entproto.Field(11)),
		field.String("profile_picture_url").MaxLen(128).Annotations(entproto.Field(12)),
		field.Enum("status").
			Values("new", "active", "inactive").
			Default("new").
			Annotations(
				entproto.Field(13),
				entproto.Enum(
					map[string]int32{
						"new":      0,
						"active":   1,
						"inactive": 2,
					},
				),
			),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("jobs", Job.Type).Annotations(entproto.Field(14)),
		edge.To("experiences", WorkExperience.Type).Annotations(entproto.Field(15)),
	}
}

// Annotations of the User
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
