package schema

import (
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entproto.Field(2)),
		field.Time("updated_at").Default(time.Now).Annotations(entproto.Field(3)),
		field.String("title").MaxLen(128).NotEmpty().Annotations(entproto.Field(4)),
		field.Uint64("min_salary").Positive().Annotations(entproto.Field(5)),
		field.Uint64("max_salary").Positive().Annotations(entproto.Field(6)),
		field.Enum("salary_unit").Values("VND", "USD").
			Default("VND").
			Annotations(
				entproto.Field(7),
				entproto.Enum(
					map[string]int32{
						"VND": 0,
						"USD": 1,
					},
				),
			),
		field.Enum("employment_type").
			Values("fulltime", "parttime", "contract", "intern", "freelance", "other").
			Default("fulltime").
			Annotations(
				entproto.Field(8),
				entproto.Enum(
					map[string]int32{
						"fulltime":  0,
						"parttime":  1,
						"contract":  2,
						"intern":    3,
						"freelance": 4,
						"other":     5,
					},
				),
			),
		field.Text("requirements").MaxLen(1024).Annotations(entproto.Field(9)),
		field.Text("responsibilities").MaxLen(1024).Annotations(entproto.Field(10)),
		field.Text("benefits").MaxLen(1024).Annotations(entproto.Field(11)),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("jobs").Unique().Annotations(entproto.Field(12)),
		edge.From("skills", Skill.Type).Ref("jobs").Annotations(entproto.Field(13)),
	}
}

// Annotations of the Job
func (Job) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
