// Code generated by entc, DO NOT EDIT.

package job

import (
	"refernet/internal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// MinSalary applies equality check predicate on the "min_salary" field. It's identical to MinSalaryEQ.
func MinSalary(v uint64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMinSalary), v))
	})
}

// MaxSalary applies equality check predicate on the "max_salary" field. It's identical to MaxSalaryEQ.
func MaxSalary(v uint64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxSalary), v))
	})
}

// Requirements applies equality check predicate on the "requirements" field. It's identical to RequirementsEQ.
func Requirements(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRequirements), v))
	})
}

// Responsibilities applies equality check predicate on the "responsibilities" field. It's identical to ResponsibilitiesEQ.
func Responsibilities(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResponsibilities), v))
	})
}

// Benefits applies equality check predicate on the "benefits" field. It's identical to BenefitsEQ.
func Benefits(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBenefits), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// MinSalaryEQ applies the EQ predicate on the "min_salary" field.
func MinSalaryEQ(v uint64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMinSalary), v))
	})
}

// MinSalaryNEQ applies the NEQ predicate on the "min_salary" field.
func MinSalaryNEQ(v uint64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMinSalary), v))
	})
}

// MinSalaryIn applies the In predicate on the "min_salary" field.
func MinSalaryIn(vs ...uint64) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMinSalary), v...))
	})
}

// MinSalaryNotIn applies the NotIn predicate on the "min_salary" field.
func MinSalaryNotIn(vs ...uint64) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMinSalary), v...))
	})
}

// MinSalaryGT applies the GT predicate on the "min_salary" field.
func MinSalaryGT(v uint64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMinSalary), v))
	})
}

// MinSalaryGTE applies the GTE predicate on the "min_salary" field.
func MinSalaryGTE(v uint64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMinSalary), v))
	})
}

// MinSalaryLT applies the LT predicate on the "min_salary" field.
func MinSalaryLT(v uint64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMinSalary), v))
	})
}

// MinSalaryLTE applies the LTE predicate on the "min_salary" field.
func MinSalaryLTE(v uint64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMinSalary), v))
	})
}

// MaxSalaryEQ applies the EQ predicate on the "max_salary" field.
func MaxSalaryEQ(v uint64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxSalary), v))
	})
}

// MaxSalaryNEQ applies the NEQ predicate on the "max_salary" field.
func MaxSalaryNEQ(v uint64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMaxSalary), v))
	})
}

// MaxSalaryIn applies the In predicate on the "max_salary" field.
func MaxSalaryIn(vs ...uint64) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMaxSalary), v...))
	})
}

// MaxSalaryNotIn applies the NotIn predicate on the "max_salary" field.
func MaxSalaryNotIn(vs ...uint64) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMaxSalary), v...))
	})
}

// MaxSalaryGT applies the GT predicate on the "max_salary" field.
func MaxSalaryGT(v uint64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMaxSalary), v))
	})
}

// MaxSalaryGTE applies the GTE predicate on the "max_salary" field.
func MaxSalaryGTE(v uint64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMaxSalary), v))
	})
}

// MaxSalaryLT applies the LT predicate on the "max_salary" field.
func MaxSalaryLT(v uint64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMaxSalary), v))
	})
}

// MaxSalaryLTE applies the LTE predicate on the "max_salary" field.
func MaxSalaryLTE(v uint64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMaxSalary), v))
	})
}

// SalaryUnitEQ applies the EQ predicate on the "salary_unit" field.
func SalaryUnitEQ(v SalaryUnit) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSalaryUnit), v))
	})
}

// SalaryUnitNEQ applies the NEQ predicate on the "salary_unit" field.
func SalaryUnitNEQ(v SalaryUnit) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSalaryUnit), v))
	})
}

// SalaryUnitIn applies the In predicate on the "salary_unit" field.
func SalaryUnitIn(vs ...SalaryUnit) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSalaryUnit), v...))
	})
}

// SalaryUnitNotIn applies the NotIn predicate on the "salary_unit" field.
func SalaryUnitNotIn(vs ...SalaryUnit) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSalaryUnit), v...))
	})
}

// EmploymentTypeEQ applies the EQ predicate on the "employment_type" field.
func EmploymentTypeEQ(v EmploymentType) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmploymentType), v))
	})
}

// EmploymentTypeNEQ applies the NEQ predicate on the "employment_type" field.
func EmploymentTypeNEQ(v EmploymentType) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmploymentType), v))
	})
}

// EmploymentTypeIn applies the In predicate on the "employment_type" field.
func EmploymentTypeIn(vs ...EmploymentType) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmploymentType), v...))
	})
}

// EmploymentTypeNotIn applies the NotIn predicate on the "employment_type" field.
func EmploymentTypeNotIn(vs ...EmploymentType) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmploymentType), v...))
	})
}

// RequirementsEQ applies the EQ predicate on the "requirements" field.
func RequirementsEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRequirements), v))
	})
}

// RequirementsNEQ applies the NEQ predicate on the "requirements" field.
func RequirementsNEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRequirements), v))
	})
}

// RequirementsIn applies the In predicate on the "requirements" field.
func RequirementsIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRequirements), v...))
	})
}

// RequirementsNotIn applies the NotIn predicate on the "requirements" field.
func RequirementsNotIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRequirements), v...))
	})
}

// RequirementsGT applies the GT predicate on the "requirements" field.
func RequirementsGT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRequirements), v))
	})
}

// RequirementsGTE applies the GTE predicate on the "requirements" field.
func RequirementsGTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRequirements), v))
	})
}

// RequirementsLT applies the LT predicate on the "requirements" field.
func RequirementsLT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRequirements), v))
	})
}

// RequirementsLTE applies the LTE predicate on the "requirements" field.
func RequirementsLTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRequirements), v))
	})
}

// RequirementsContains applies the Contains predicate on the "requirements" field.
func RequirementsContains(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRequirements), v))
	})
}

// RequirementsHasPrefix applies the HasPrefix predicate on the "requirements" field.
func RequirementsHasPrefix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRequirements), v))
	})
}

// RequirementsHasSuffix applies the HasSuffix predicate on the "requirements" field.
func RequirementsHasSuffix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRequirements), v))
	})
}

// RequirementsEqualFold applies the EqualFold predicate on the "requirements" field.
func RequirementsEqualFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRequirements), v))
	})
}

// RequirementsContainsFold applies the ContainsFold predicate on the "requirements" field.
func RequirementsContainsFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRequirements), v))
	})
}

// ResponsibilitiesEQ applies the EQ predicate on the "responsibilities" field.
func ResponsibilitiesEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResponsibilities), v))
	})
}

// ResponsibilitiesNEQ applies the NEQ predicate on the "responsibilities" field.
func ResponsibilitiesNEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldResponsibilities), v))
	})
}

// ResponsibilitiesIn applies the In predicate on the "responsibilities" field.
func ResponsibilitiesIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldResponsibilities), v...))
	})
}

// ResponsibilitiesNotIn applies the NotIn predicate on the "responsibilities" field.
func ResponsibilitiesNotIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldResponsibilities), v...))
	})
}

// ResponsibilitiesGT applies the GT predicate on the "responsibilities" field.
func ResponsibilitiesGT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldResponsibilities), v))
	})
}

// ResponsibilitiesGTE applies the GTE predicate on the "responsibilities" field.
func ResponsibilitiesGTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldResponsibilities), v))
	})
}

// ResponsibilitiesLT applies the LT predicate on the "responsibilities" field.
func ResponsibilitiesLT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldResponsibilities), v))
	})
}

// ResponsibilitiesLTE applies the LTE predicate on the "responsibilities" field.
func ResponsibilitiesLTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldResponsibilities), v))
	})
}

// ResponsibilitiesContains applies the Contains predicate on the "responsibilities" field.
func ResponsibilitiesContains(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldResponsibilities), v))
	})
}

// ResponsibilitiesHasPrefix applies the HasPrefix predicate on the "responsibilities" field.
func ResponsibilitiesHasPrefix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldResponsibilities), v))
	})
}

// ResponsibilitiesHasSuffix applies the HasSuffix predicate on the "responsibilities" field.
func ResponsibilitiesHasSuffix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldResponsibilities), v))
	})
}

// ResponsibilitiesEqualFold applies the EqualFold predicate on the "responsibilities" field.
func ResponsibilitiesEqualFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldResponsibilities), v))
	})
}

// ResponsibilitiesContainsFold applies the ContainsFold predicate on the "responsibilities" field.
func ResponsibilitiesContainsFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldResponsibilities), v))
	})
}

// BenefitsEQ applies the EQ predicate on the "benefits" field.
func BenefitsEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBenefits), v))
	})
}

// BenefitsNEQ applies the NEQ predicate on the "benefits" field.
func BenefitsNEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBenefits), v))
	})
}

// BenefitsIn applies the In predicate on the "benefits" field.
func BenefitsIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBenefits), v...))
	})
}

// BenefitsNotIn applies the NotIn predicate on the "benefits" field.
func BenefitsNotIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBenefits), v...))
	})
}

// BenefitsGT applies the GT predicate on the "benefits" field.
func BenefitsGT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBenefits), v))
	})
}

// BenefitsGTE applies the GTE predicate on the "benefits" field.
func BenefitsGTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBenefits), v))
	})
}

// BenefitsLT applies the LT predicate on the "benefits" field.
func BenefitsLT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBenefits), v))
	})
}

// BenefitsLTE applies the LTE predicate on the "benefits" field.
func BenefitsLTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBenefits), v))
	})
}

// BenefitsContains applies the Contains predicate on the "benefits" field.
func BenefitsContains(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBenefits), v))
	})
}

// BenefitsHasPrefix applies the HasPrefix predicate on the "benefits" field.
func BenefitsHasPrefix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBenefits), v))
	})
}

// BenefitsHasSuffix applies the HasSuffix predicate on the "benefits" field.
func BenefitsHasSuffix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBenefits), v))
	})
}

// BenefitsEqualFold applies the EqualFold predicate on the "benefits" field.
func BenefitsEqualFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBenefits), v))
	})
}

// BenefitsContainsFold applies the ContainsFold predicate on the "benefits" field.
func BenefitsContainsFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBenefits), v))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSkills applies the HasEdge predicate on the "skills" edge.
func HasSkills() predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SkillsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, SkillsTable, SkillsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSkillsWith applies the HasEdge predicate on the "skills" edge with a given conditions (other predicates).
func HasSkillsWith(preds ...predicate.Skill) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SkillsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, SkillsTable, SkillsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Job) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Job) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Job) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		p(s.Not())
	})
}
