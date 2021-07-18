// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"refernet/ent/job"
	"refernet/ent/user"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Job is the model entity for the Job schema.
type Job struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// MinSalary holds the value of the "min_salary" field.
	MinSalary uint64 `json:"min_salary,omitempty"`
	// MaxSalary holds the value of the "max_salary" field.
	MaxSalary uint64 `json:"max_salary,omitempty"`
	// Type holds the value of the "type" field.
	Type job.Type `json:"type,omitempty"`
	// Requirements holds the value of the "requirements" field.
	Requirements string `json:"requirements,omitempty"`
	// Responsibilities holds the value of the "responsibilities" field.
	Responsibilities string `json:"responsibilities,omitempty"`
	// Benefits holds the value of the "benefits" field.
	Benefits string `json:"benefits,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the JobQuery when eager-loading is set.
	Edges     JobEdges `json:"edges"`
	user_jobs *int
}

// JobEdges holds the relations/edges for other nodes in the graph.
type JobEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Job) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case job.FieldID, job.FieldMinSalary, job.FieldMaxSalary:
			values[i] = new(sql.NullInt64)
		case job.FieldTitle, job.FieldLocation, job.FieldType, job.FieldRequirements, job.FieldResponsibilities, job.FieldBenefits:
			values[i] = new(sql.NullString)
		case job.FieldCreatedAt, job.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case job.ForeignKeys[0]: // user_jobs
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Job", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Job fields.
func (j *Job) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case job.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			j.ID = int(value.Int64)
		case job.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				j.CreatedAt = value.Time
			}
		case job.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				j.UpdatedAt = value.Time
			}
		case job.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				j.Title = value.String
			}
		case job.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				j.Location = value.String
			}
		case job.FieldMinSalary:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field min_salary", values[i])
			} else if value.Valid {
				j.MinSalary = uint64(value.Int64)
			}
		case job.FieldMaxSalary:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field max_salary", values[i])
			} else if value.Valid {
				j.MaxSalary = uint64(value.Int64)
			}
		case job.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				j.Type = job.Type(value.String)
			}
		case job.FieldRequirements:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field requirements", values[i])
			} else if value.Valid {
				j.Requirements = value.String
			}
		case job.FieldResponsibilities:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field responsibilities", values[i])
			} else if value.Valid {
				j.Responsibilities = value.String
			}
		case job.FieldBenefits:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field benefits", values[i])
			} else if value.Valid {
				j.Benefits = value.String
			}
		case job.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_jobs", value)
			} else if value.Valid {
				j.user_jobs = new(int)
				*j.user_jobs = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Job entity.
func (j *Job) QueryOwner() *UserQuery {
	return (&JobClient{config: j.config}).QueryOwner(j)
}

// Update returns a builder for updating this Job.
// Note that you need to call Job.Unwrap() before calling this method if this Job
// was returned from a transaction, and the transaction was committed or rolled back.
func (j *Job) Update() *JobUpdateOne {
	return (&JobClient{config: j.config}).UpdateOne(j)
}

// Unwrap unwraps the Job entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (j *Job) Unwrap() *Job {
	tx, ok := j.config.driver.(*txDriver)
	if !ok {
		panic("ent: Job is not a transactional entity")
	}
	j.config.driver = tx.drv
	return j
}

// String implements the fmt.Stringer.
func (j *Job) String() string {
	var builder strings.Builder
	builder.WriteString("Job(")
	builder.WriteString(fmt.Sprintf("id=%v", j.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(j.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(j.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", title=")
	builder.WriteString(j.Title)
	builder.WriteString(", location=")
	builder.WriteString(j.Location)
	builder.WriteString(", min_salary=")
	builder.WriteString(fmt.Sprintf("%v", j.MinSalary))
	builder.WriteString(", max_salary=")
	builder.WriteString(fmt.Sprintf("%v", j.MaxSalary))
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", j.Type))
	builder.WriteString(", requirements=")
	builder.WriteString(j.Requirements)
	builder.WriteString(", responsibilities=")
	builder.WriteString(j.Responsibilities)
	builder.WriteString(", benefits=")
	builder.WriteString(j.Benefits)
	builder.WriteByte(')')
	return builder.String()
}

// Jobs is a parsable slice of Job.
type Jobs []*Job

func (j Jobs) config(cfg config) {
	for _i := range j {
		j[_i].config = cfg
	}
}
