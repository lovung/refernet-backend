// Code generated by entc, DO NOT EDIT.

package skill

import (
	"time"
)

const (
	// Label holds the string label denoting the skill type in the database.
	Label = "skill"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLogoURL holds the string denoting the logo_url field in the database.
	FieldLogoURL = "logo_url"
	// EdgeExperiences holds the string denoting the experiences edge name in mutations.
	EdgeExperiences = "experiences"
	// EdgeJobs holds the string denoting the jobs edge name in mutations.
	EdgeJobs = "jobs"
	// Table holds the table name of the skill in the database.
	Table = "skills"
	// ExperiencesTable is the table the holds the experiences relation/edge. The primary key declared below.
	ExperiencesTable = "skill_experiences"
	// ExperiencesInverseTable is the table name for the WorkExperience entity.
	// It exists in this package in order to avoid circular dependency with the "workexperience" package.
	ExperiencesInverseTable = "work_experiences"
	// JobsTable is the table the holds the jobs relation/edge. The primary key declared below.
	JobsTable = "skill_jobs"
	// JobsInverseTable is the table name for the Job entity.
	// It exists in this package in order to avoid circular dependency with the "job" package.
	JobsInverseTable = "jobs"
)

// Columns holds all SQL columns for skill fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldLogoURL,
}

var (
	// ExperiencesPrimaryKey and ExperiencesColumn2 are the table columns denoting the
	// primary key for the experiences relation (M2M).
	ExperiencesPrimaryKey = []string{"skill_id", "work_experience_id"}
	// JobsPrimaryKey and JobsColumn2 are the table columns denoting the
	// primary key for the jobs relation (M2M).
	JobsPrimaryKey = []string{"skill_id", "job_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultLogoURL holds the default value on creation for the "logo_url" field.
	DefaultLogoURL string
)
