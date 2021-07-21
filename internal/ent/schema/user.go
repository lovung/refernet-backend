package schema

import (
	"time"

	"entgo.io/ent"
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
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now),
		field.String("username").MaxLen(128).NotEmpty(),
		field.String("fullname").MaxLen(128).NotEmpty(),
		field.String("password").MaxLen(256).NotEmpty(),
		field.String("email").MaxLen(128).NotEmpty(),
		field.String("phone").MaxLen(20),
		field.String("bio").MaxLen(64),
		field.Text("intro").MaxLen(1024),
		field.String("github_profile").MaxLen(128),
		field.String("profile_picture_url").MaxLen(128),
		field.Enum("status").Values("new", "verified", "inactive").Default("new"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("jobs", Job.Type),
		edge.To("experiences", WorkExperience.Type),
	}
}
