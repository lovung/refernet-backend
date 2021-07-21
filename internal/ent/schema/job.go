package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Job holds the schema definition for the Job entity.
type Job struct {
	ent.Schema
}

// Fields of the Job.
func (Job) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now),
		field.String("title").
			NotEmpty(),
		field.String("location"),
		field.Uint64("min_salary"),
		field.Uint64("max_salary"),
		field.Enum("type").
			Values("fulltime", "parttime", "contract", "intern", "freelance", "other").
			Default("fulltime"),
		field.Text("requirements").
			MaxLen(1024),
		field.Text("responsibilities").
			MaxLen(1024),
		field.Text("benefits").
			MaxLen(1024),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("jobs").
			Unique(),
	}
}
