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
	// FieldMinSalary holds the string denoting the min_salary field in the database.
	FieldMinSalary = "min_salary"
	// FieldMaxSalary holds the string denoting the max_salary field in the database.
	FieldMaxSalary = "max_salary"
	// FieldSalaryUnit holds the string denoting the salary_unit field in the database.
	FieldSalaryUnit = "salary_unit"
	// FieldEmploymentType holds the string denoting the employment_type field in the database.
	FieldEmploymentType = "employment_type"
	// FieldRequirements holds the string denoting the requirements field in the database.
	FieldRequirements = "requirements"
	// FieldResponsibilities holds the string denoting the responsibilities field in the database.
	FieldResponsibilities = "responsibilities"
	// FieldBenefits holds the string denoting the benefits field in the database.
	FieldBenefits = "benefits"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeSkills holds the string denoting the skills edge name in mutations.
	EdgeSkills = "skills"
	// Table holds the table name of the job in the database.
	Table = "jobs"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "jobs"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_jobs"
	// SkillsTable is the table that holds the skills relation/edge. The primary key declared below.
	SkillsTable = "skill_jobs"
	// SkillsInverseTable is the table name for the Skill entity.
	// It exists in this package in order to avoid circular dependency with the "skill" package.
	SkillsInverseTable = "skills"
)

// Columns holds all SQL columns for job fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTitle,
	FieldMinSalary,
	FieldMaxSalary,
	FieldSalaryUnit,
	FieldEmploymentType,
	FieldRequirements,
	FieldResponsibilities,
	FieldBenefits,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "jobs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_jobs",
}

var (
	// SkillsPrimaryKey and SkillsColumn2 are the table columns denoting the
	// primary key for the skills relation (M2M).
	SkillsPrimaryKey = []string{"skill_id", "job_id"}
)

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
	// MinSalaryValidator is a validator for the "min_salary" field. It is called by the builders before save.
	MinSalaryValidator func(uint64) error
	// MaxSalaryValidator is a validator for the "max_salary" field. It is called by the builders before save.
	MaxSalaryValidator func(uint64) error
	// RequirementsValidator is a validator for the "requirements" field. It is called by the builders before save.
	RequirementsValidator func(string) error
	// ResponsibilitiesValidator is a validator for the "responsibilities" field. It is called by the builders before save.
	ResponsibilitiesValidator func(string) error
	// BenefitsValidator is a validator for the "benefits" field. It is called by the builders before save.
	BenefitsValidator func(string) error
)

// SalaryUnit defines the type for the "salary_unit" enum field.
type SalaryUnit string

// SalaryUnitVND is the default value of the SalaryUnit enum.
const DefaultSalaryUnit = SalaryUnitVND

// SalaryUnit values.
const (
	SalaryUnitVND SalaryUnit = "VND"
	SalaryUnitUSD SalaryUnit = "USD"
)

func (su SalaryUnit) String() string {
	return string(su)
}

// SalaryUnitValidator is a validator for the "salary_unit" field enum values. It is called by the builders before save.
func SalaryUnitValidator(su SalaryUnit) error {
	switch su {
	case SalaryUnitVND, SalaryUnitUSD:
		return nil
	default:
		return fmt.Errorf("job: invalid enum value for salary_unit field: %q", su)
	}
}

// EmploymentType defines the type for the "employment_type" enum field.
type EmploymentType string

// EmploymentTypeFulltime is the default value of the EmploymentType enum.
const DefaultEmploymentType = EmploymentTypeFulltime

// EmploymentType values.
const (
	EmploymentTypeFulltime  EmploymentType = "fulltime"
	EmploymentTypeParttime  EmploymentType = "parttime"
	EmploymentTypeContract  EmploymentType = "contract"
	EmploymentTypeIntern    EmploymentType = "intern"
	EmploymentTypeFreelance EmploymentType = "freelance"
	EmploymentTypeOther     EmploymentType = "other"
)

func (et EmploymentType) String() string {
	return string(et)
}

// EmploymentTypeValidator is a validator for the "employment_type" field enum values. It is called by the builders before save.
func EmploymentTypeValidator(et EmploymentType) error {
	switch et {
	case EmploymentTypeFulltime, EmploymentTypeParttime, EmploymentTypeContract, EmploymentTypeIntern, EmploymentTypeFreelance, EmploymentTypeOther:
		return nil
	default:
		return fmt.Errorf("job: invalid enum value for employment_type field: %q", et)
	}
}
