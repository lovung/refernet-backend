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
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now),
		field.String("username").
			NotEmpty(),
		field.String("fullname").
			NotEmpty(),
		field.String("email").NotEmpty(),
		field.String("phone"),
		field.String("bio"),
		field.Text("intro"),
		field.String("github_profile"),
		field.String("profile_picture_url"),
		field.Enum("status").
			Values("new", "verified", "inactive").
			Default("new"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("jobs", Job.Type),
		edge.To("experiences", WorkExperience.Type),
	}
}
