// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"refernet/internal/ent/company"
	"refernet/internal/ent/skill"
	"refernet/internal/ent/user"
	"refernet/internal/ent/workexperience"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorkExperienceCreate is the builder for creating a WorkExperience entity.
type WorkExperienceCreate struct {
	config
	mutation *WorkExperienceMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (wec *WorkExperienceCreate) SetCreatedAt(t time.Time) *WorkExperienceCreate {
	wec.mutation.SetCreatedAt(t)
	return wec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wec *WorkExperienceCreate) SetNillableCreatedAt(t *time.Time) *WorkExperienceCreate {
	if t != nil {
		wec.SetCreatedAt(*t)
	}
	return wec
}

// SetUpdatedAt sets the "updated_at" field.
func (wec *WorkExperienceCreate) SetUpdatedAt(t time.Time) *WorkExperienceCreate {
	wec.mutation.SetUpdatedAt(t)
	return wec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wec *WorkExperienceCreate) SetNillableUpdatedAt(t *time.Time) *WorkExperienceCreate {
	if t != nil {
		wec.SetUpdatedAt(*t)
	}
	return wec
}

// SetTitle sets the "title" field.
func (wec *WorkExperienceCreate) SetTitle(s string) *WorkExperienceCreate {
	wec.mutation.SetTitle(s)
	return wec
}

// SetLocation sets the "location" field.
func (wec *WorkExperienceCreate) SetLocation(s string) *WorkExperienceCreate {
	wec.mutation.SetLocation(s)
	return wec
}

// SetStartDate sets the "start_date" field.
func (wec *WorkExperienceCreate) SetStartDate(t time.Time) *WorkExperienceCreate {
	wec.mutation.SetStartDate(t)
	return wec
}

// SetEndDate sets the "end_date" field.
func (wec *WorkExperienceCreate) SetEndDate(t time.Time) *WorkExperienceCreate {
	wec.mutation.SetEndDate(t)
	return wec
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (wec *WorkExperienceCreate) SetNillableEndDate(t *time.Time) *WorkExperienceCreate {
	if t != nil {
		wec.SetEndDate(*t)
	}
	return wec
}

// SetDescription sets the "description" field.
func (wec *WorkExperienceCreate) SetDescription(s string) *WorkExperienceCreate {
	wec.mutation.SetDescription(s)
	return wec
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (wec *WorkExperienceCreate) SetOwnerID(id int) *WorkExperienceCreate {
	wec.mutation.SetOwnerID(id)
	return wec
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (wec *WorkExperienceCreate) SetNillableOwnerID(id *int) *WorkExperienceCreate {
	if id != nil {
		wec = wec.SetOwnerID(*id)
	}
	return wec
}

// SetOwner sets the "owner" edge to the User entity.
func (wec *WorkExperienceCreate) SetOwner(u *User) *WorkExperienceCreate {
	return wec.SetOwnerID(u.ID)
}

// SetInCompanyID sets the "in_company" edge to the Company entity by ID.
func (wec *WorkExperienceCreate) SetInCompanyID(id int) *WorkExperienceCreate {
	wec.mutation.SetInCompanyID(id)
	return wec
}

// SetNillableInCompanyID sets the "in_company" edge to the Company entity by ID if the given value is not nil.
func (wec *WorkExperienceCreate) SetNillableInCompanyID(id *int) *WorkExperienceCreate {
	if id != nil {
		wec = wec.SetInCompanyID(*id)
	}
	return wec
}

// SetInCompany sets the "in_company" edge to the Company entity.
func (wec *WorkExperienceCreate) SetInCompany(c *Company) *WorkExperienceCreate {
	return wec.SetInCompanyID(c.ID)
}

// AddSkillIDs adds the "skills" edge to the Skill entity by IDs.
func (wec *WorkExperienceCreate) AddSkillIDs(ids ...int) *WorkExperienceCreate {
	wec.mutation.AddSkillIDs(ids...)
	return wec
}

// AddSkills adds the "skills" edges to the Skill entity.
func (wec *WorkExperienceCreate) AddSkills(s ...*Skill) *WorkExperienceCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return wec.AddSkillIDs(ids...)
}

// Mutation returns the WorkExperienceMutation object of the builder.
func (wec *WorkExperienceCreate) Mutation() *WorkExperienceMutation {
	return wec.mutation
}

// Save creates the WorkExperience in the database.
func (wec *WorkExperienceCreate) Save(ctx context.Context) (*WorkExperience, error) {
	var (
		err  error
		node *WorkExperience
	)
	wec.defaults()
	if len(wec.hooks) == 0 {
		if err = wec.check(); err != nil {
			return nil, err
		}
		node, err = wec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkExperienceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wec.check(); err != nil {
				return nil, err
			}
			wec.mutation = mutation
			if node, err = wec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(wec.hooks) - 1; i >= 0; i-- {
			if wec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wec *WorkExperienceCreate) SaveX(ctx context.Context) *WorkExperience {
	v, err := wec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wec *WorkExperienceCreate) Exec(ctx context.Context) error {
	_, err := wec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wec *WorkExperienceCreate) ExecX(ctx context.Context) {
	if err := wec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wec *WorkExperienceCreate) defaults() {
	if _, ok := wec.mutation.CreatedAt(); !ok {
		v := workexperience.DefaultCreatedAt()
		wec.mutation.SetCreatedAt(v)
	}
	if _, ok := wec.mutation.UpdatedAt(); !ok {
		v := workexperience.DefaultUpdatedAt()
		wec.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wec *WorkExperienceCreate) check() error {
	if _, ok := wec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := wec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := wec.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	if v, ok := wec.mutation.Title(); ok {
		if err := workexperience.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "title": %w`, err)}
		}
	}
	if _, ok := wec.mutation.Location(); !ok {
		return &ValidationError{Name: "location", err: errors.New(`ent: missing required field "location"`)}
	}
	if v, ok := wec.mutation.Location(); ok {
		if err := workexperience.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf(`ent: validator failed for field "location": %w`, err)}
		}
	}
	if _, ok := wec.mutation.StartDate(); !ok {
		return &ValidationError{Name: "start_date", err: errors.New(`ent: missing required field "start_date"`)}
	}
	if _, ok := wec.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "description"`)}
	}
	if v, ok := wec.mutation.Description(); ok {
		if err := workexperience.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "description": %w`, err)}
		}
	}
	return nil
}

func (wec *WorkExperienceCreate) sqlSave(ctx context.Context) (*WorkExperience, error) {
	_node, _spec := wec.createSpec()
	if err := sqlgraph.CreateNode(ctx, wec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (wec *WorkExperienceCreate) createSpec() (*WorkExperience, *sqlgraph.CreateSpec) {
	var (
		_node = &WorkExperience{config: wec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: workexperience.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workexperience.FieldID,
			},
		}
	)
	if value, ok := wec.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workexperience.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := wec.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workexperience.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := wec.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workexperience.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := wec.mutation.Location(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workexperience.FieldLocation,
		})
		_node.Location = value
	}
	if value, ok := wec.mutation.StartDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workexperience.FieldStartDate,
		})
		_node.StartDate = value
	}
	if value, ok := wec.mutation.EndDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workexperience.FieldEndDate,
		})
		_node.EndDate = value
	}
	if value, ok := wec.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workexperience.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := wec.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workexperience.OwnerTable,
			Columns: []string{workexperience.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_experiences = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wec.mutation.InCompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workexperience.InCompanyTable,
			Columns: []string{workexperience.InCompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: company.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.company_staffs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wec.mutation.SkillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   workexperience.SkillsTable,
			Columns: workexperience.SkillsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: skill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkExperienceCreateBulk is the builder for creating many WorkExperience entities in bulk.
type WorkExperienceCreateBulk struct {
	config
	builders []*WorkExperienceCreate
}

// Save creates the WorkExperience entities in the database.
func (wecb *WorkExperienceCreateBulk) Save(ctx context.Context) ([]*WorkExperience, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wecb.builders))
	nodes := make([]*WorkExperience, len(wecb.builders))
	mutators := make([]Mutator, len(wecb.builders))
	for i := range wecb.builders {
		func(i int, root context.Context) {
			builder := wecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkExperienceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wecb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wecb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wecb *WorkExperienceCreateBulk) SaveX(ctx context.Context) []*WorkExperience {
	v, err := wecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
