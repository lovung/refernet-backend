// Code generated by entc, DO NOT EDIT.

package job

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the job type in the database.
	Label = "job"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldMinSalary holds the string denoting the min_salary field in the database.
	FieldMinSalary = "min_salary"
	// FieldMaxSalary holds the string denoting the max_salary field in the database.
	FieldMaxSalary = "max_salary"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldRequirements holds the string denoting the requirements field in the database.
	FieldRequirements = "requirements"
	// FieldResponsibilities holds the string denoting the responsibilities field in the database.
	FieldResponsibilities = "responsibilities"
	// FieldBenefits holds the string denoting the benefits field in the database.
	FieldBenefits = "benefits"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the job in the database.
	Table = "jobs"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "jobs"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_jobs"
)

// Columns holds all SQL columns for job fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTitle,
	FieldLocation,
	FieldMinSalary,
	FieldMaxSalary,
	FieldType,
	FieldRequirements,
	FieldResponsibilities,
	FieldBenefits,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "jobs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_jobs",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// RequirementsValidator is a validator for the "requirements" field. It is called by the builders before save.
	RequirementsValidator func(string) error
	// ResponsibilitiesValidator is a validator for the "responsibilities" field. It is called by the builders before save.
	ResponsibilitiesValidator func(string) error
	// BenefitsValidator is a validator for the "benefits" field. It is called by the builders before save.
	BenefitsValidator func(string) error
)

// Type defines the type for the "type" enum field.
type Type string

// TypeFulltime is the default value of the Type enum.
const DefaultType = TypeFulltime

// Type values.
const (
	TypeFulltime  Type = "fulltime"
	TypeParttime  Type = "parttime"
	TypeContract  Type = "contract"
	TypeIntern    Type = "intern"
	TypeFreelance Type = "freelance"
	TypeOther     Type = "other"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeFulltime, TypeParttime, TypeContract, TypeIntern, TypeFreelance, TypeOther:
		return nil
	default:
		return fmt.Errorf("job: invalid enum value for type field: %q", _type)
	}
}
