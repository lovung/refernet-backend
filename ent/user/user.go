// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldFullname holds the string denoting the fullname field in the database.
	FieldFullname = "fullname"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldBio holds the string denoting the bio field in the database.
	FieldBio = "bio"
	// FieldIntro holds the string denoting the intro field in the database.
	FieldIntro = "intro"
	// FieldGithubProfile holds the string denoting the github_profile field in the database.
	FieldGithubProfile = "github_profile"
	// FieldProfilePictureURL holds the string denoting the profile_picture_url field in the database.
	FieldProfilePictureURL = "profile_picture_url"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeJobs holds the string denoting the jobs edge name in mutations.
	EdgeJobs = "jobs"
	// EdgeExperiences holds the string denoting the experiences edge name in mutations.
	EdgeExperiences = "experiences"
	// Table holds the table name of the user in the database.
	Table = "users"
	// JobsTable is the table the holds the jobs relation/edge.
	JobsTable = "jobs"
	// JobsInverseTable is the table name for the Job entity.
	// It exists in this package in order to avoid circular dependency with the "job" package.
	JobsInverseTable = "jobs"
	// JobsColumn is the table column denoting the jobs relation/edge.
	JobsColumn = "user_jobs"
	// ExperiencesTable is the table the holds the experiences relation/edge. The primary key declared below.
	ExperiencesTable = "user_experiences"
	// ExperiencesInverseTable is the table name for the WorkExperience entity.
	// It exists in this package in order to avoid circular dependency with the "workexperience" package.
	ExperiencesInverseTable = "work_experiences"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUsername,
	FieldFullname,
	FieldEmail,
	FieldPhone,
	FieldBio,
	FieldIntro,
	FieldGithubProfile,
	FieldProfilePictureURL,
	FieldStatus,
}

var (
	// ExperiencesPrimaryKey and ExperiencesColumn2 are the table columns denoting the
	// primary key for the experiences relation (M2M).
	ExperiencesPrimaryKey = []string{"user_id", "work_experience_id"}
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
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// FullnameValidator is a validator for the "fullname" field. It is called by the builders before save.
	FullnameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
)

// Status defines the type for the "status" enum field.
type Status string

// StatusNew is the default value of the Status enum.
const DefaultStatus = StatusNew

// Status values.
const (
	StatusNew      Status = "new"
	StatusVerified Status = "verified"
	StatusInactive Status = "inactive"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusNew, StatusVerified, StatusInactive:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for status field: %q", s)
	}
}
