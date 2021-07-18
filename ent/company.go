// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"refernet/ent/company"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Company is the model entity for the Company schema.
type Company struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Overview holds the value of the "overview" field.
	Overview string `json:"overview,omitempty"`
	// Website holds the value of the "website" field.
	Website string `json:"website,omitempty"`
	// Industry holds the value of the "industry" field.
	Industry []string `json:"industry,omitempty"`
	// LogoURL holds the value of the "logo_url" field.
	LogoURL string `json:"logo_url,omitempty"`
	// Size holds the value of the "size" field.
	Size company.Size `json:"size,omitempty"`
	// FoundedAt holds the value of the "founded_at" field.
	FoundedAt int `json:"founded_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Company) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case company.FieldIndustry:
			values[i] = new([]byte)
		case company.FieldID, company.FieldFoundedAt:
			values[i] = new(sql.NullInt64)
		case company.FieldName, company.FieldOverview, company.FieldWebsite, company.FieldLogoURL, company.FieldSize:
			values[i] = new(sql.NullString)
		case company.FieldCreatedAt, company.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Company", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Company fields.
func (c *Company) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case company.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case company.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case company.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case company.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case company.FieldOverview:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field overview", values[i])
			} else if value.Valid {
				c.Overview = value.String
			}
		case company.FieldWebsite:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field website", values[i])
			} else if value.Valid {
				c.Website = value.String
			}
		case company.FieldIndustry:

			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field industry", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Industry); err != nil {
					return fmt.Errorf("unmarshal field industry: %w", err)
				}
			}
		case company.FieldLogoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logo_url", values[i])
			} else if value.Valid {
				c.LogoURL = value.String
			}
		case company.FieldSize:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				c.Size = company.Size(value.String)
			}
		case company.FieldFoundedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field founded_at", values[i])
			} else if value.Valid {
				c.FoundedAt = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Company.
// Note that you need to call Company.Unwrap() before calling this method if this Company
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Company) Update() *CompanyUpdateOne {
	return (&CompanyClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Company entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Company) Unwrap() *Company {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Company is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Company) String() string {
	var builder strings.Builder
	builder.WriteString("Company(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteString(", overview=")
	builder.WriteString(c.Overview)
	builder.WriteString(", website=")
	builder.WriteString(c.Website)
	builder.WriteString(", industry=")
	builder.WriteString(fmt.Sprintf("%v", c.Industry))
	builder.WriteString(", logo_url=")
	builder.WriteString(c.LogoURL)
	builder.WriteString(", size=")
	builder.WriteString(fmt.Sprintf("%v", c.Size))
	builder.WriteString(", founded_at=")
	builder.WriteString(fmt.Sprintf("%v", c.FoundedAt))
	builder.WriteByte(')')
	return builder.String()
}

// Companies is a parsable slice of Company.
type Companies []*Company

func (c Companies) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
