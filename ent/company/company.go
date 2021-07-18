// Code generated by entc, DO NOT EDIT.

package company

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the company type in the database.
	Label = "company"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldOverview holds the string denoting the overview field in the database.
	FieldOverview = "overview"
	// FieldWebsite holds the string denoting the website field in the database.
	FieldWebsite = "website"
	// FieldIndustry holds the string denoting the industry field in the database.
	FieldIndustry = "industry"
	// FieldLogoURL holds the string denoting the logo_url field in the database.
	FieldLogoURL = "logo_url"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// FieldFoundedAt holds the string denoting the founded_at field in the database.
	FieldFoundedAt = "founded_at"
	// EdgeStaffs holds the string denoting the staffs edge name in mutations.
	EdgeStaffs = "staffs"
	// Table holds the table name of the company in the database.
	Table = "companies"
	// StaffsTable is the table the holds the staffs relation/edge. The primary key declared below.
	StaffsTable = "company_staffs"
	// StaffsInverseTable is the table name for the WorkExperience entity.
	// It exists in this package in order to avoid circular dependency with the "workexperience" package.
	StaffsInverseTable = "work_experiences"
)

// Columns holds all SQL columns for company fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldOverview,
	FieldWebsite,
	FieldIndustry,
	FieldLogoURL,
	FieldSize,
	FieldFoundedAt,
}

var (
	// StaffsPrimaryKey and StaffsColumn2 are the table columns denoting the
	// primary key for the staffs relation (M2M).
	StaffsPrimaryKey = []string{"company_id", "work_experience_id"}
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
	// FoundedAtValidator is a validator for the "founded_at" field. It is called by the builders before save.
	FoundedAtValidator func(int) error
)

// Size defines the type for the "size" enum field.
type Size string

// SizeMedium is the default value of the Size enum.
const DefaultSize = SizeMedium

// Size values.
const (
	SizeStartup Size = "startup"
	SizeSmall   Size = "small"
	SizeMedium  Size = "medium"
	SizeBig     Size = "big"
	SizeHuge    Size = "huge"
)

func (s Size) String() string {
	return string(s)
}

// SizeValidator is a validator for the "size" field enum values. It is called by the builders before save.
func SizeValidator(s Size) error {
	switch s {
	case SizeStartup, SizeSmall, SizeMedium, SizeBig, SizeHuge:
		return nil
	default:
		return fmt.Errorf("company: invalid enum value for size field: %q", s)
	}
}
