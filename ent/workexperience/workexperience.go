// Code generated by entc, DO NOT EDIT.

package workexperience

import (
	"time"
)

const (
	// Label holds the string label denoting the workexperience type in the database.
	Label = "work_experience"
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
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the end_date field in the database.
	FieldEndDate = "end_date"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeCompany holds the string denoting the company edge name in mutations.
	EdgeCompany = "company"
	// Table holds the table name of the workexperience in the database.
	Table = "work_experiences"
	// UserTable is the table the holds the user relation/edge. The primary key declared below.
	UserTable = "user_experiences"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// CompanyTable is the table the holds the company relation/edge. The primary key declared below.
	CompanyTable = "company_staffs"
	// CompanyInverseTable is the table name for the Company entity.
	// It exists in this package in order to avoid circular dependency with the "company" package.
	CompanyInverseTable = "companies"
)

// Columns holds all SQL columns for workexperience fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTitle,
	FieldLocation,
	FieldStartDate,
	FieldEndDate,
	FieldDescription,
}

var (
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"user_id", "work_experience_id"}
	// CompanyPrimaryKey and CompanyColumn2 are the table columns denoting the
	// primary key for the company relation (M2M).
	CompanyPrimaryKey = []string{"company_id", "work_experience_id"}
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
)
