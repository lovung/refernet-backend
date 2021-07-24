// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"refernet/internal/ent/company"
	"refernet/internal/ent/workexperience"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CompanyCreate is the builder for creating a Company entity.
type CompanyCreate struct {
	config
	mutation *CompanyMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CompanyCreate) SetCreatedAt(t time.Time) *CompanyCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CompanyCreate) SetNillableCreatedAt(t *time.Time) *CompanyCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CompanyCreate) SetUpdatedAt(t time.Time) *CompanyCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CompanyCreate) SetNillableUpdatedAt(t *time.Time) *CompanyCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *CompanyCreate) SetName(s string) *CompanyCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetOverview sets the "overview" field.
func (cc *CompanyCreate) SetOverview(s string) *CompanyCreate {
	cc.mutation.SetOverview(s)
	return cc
}

// SetWebsite sets the "website" field.
func (cc *CompanyCreate) SetWebsite(s string) *CompanyCreate {
	cc.mutation.SetWebsite(s)
	return cc
}

// SetLogoURL sets the "logo_url" field.
func (cc *CompanyCreate) SetLogoURL(s string) *CompanyCreate {
	cc.mutation.SetLogoURL(s)
	return cc
}

// SetSize sets the "size" field.
func (cc *CompanyCreate) SetSize(c company.Size) *CompanyCreate {
	cc.mutation.SetSize(c)
	return cc
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (cc *CompanyCreate) SetNillableSize(c *company.Size) *CompanyCreate {
	if c != nil {
		cc.SetSize(*c)
	}
	return cc
}

// SetFoundedAt sets the "founded_at" field.
func (cc *CompanyCreate) SetFoundedAt(i int) *CompanyCreate {
	cc.mutation.SetFoundedAt(i)
	return cc
}

// AddStaffIDs adds the "staffs" edge to the WorkExperience entity by IDs.
func (cc *CompanyCreate) AddStaffIDs(ids ...int) *CompanyCreate {
	cc.mutation.AddStaffIDs(ids...)
	return cc
}

// AddStaffs adds the "staffs" edges to the WorkExperience entity.
func (cc *CompanyCreate) AddStaffs(w ...*WorkExperience) *CompanyCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return cc.AddStaffIDs(ids...)
}

// Mutation returns the CompanyMutation object of the builder.
func (cc *CompanyCreate) Mutation() *CompanyMutation {
	return cc.mutation
}

// Save creates the Company in the database.
func (cc *CompanyCreate) Save(ctx context.Context) (*Company, error) {
	var (
		err  error
		node *Company
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CompanyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CompanyCreate) SaveX(ctx context.Context) *Company {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CompanyCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CompanyCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CompanyCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := company.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := company.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.Size(); !ok {
		v := company.DefaultSize
		cc.mutation.SetSize(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CompanyCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := company.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Overview(); !ok {
		return &ValidationError{Name: "overview", err: errors.New(`ent: missing required field "overview"`)}
	}
	if v, ok := cc.mutation.Overview(); ok {
		if err := company.OverviewValidator(v); err != nil {
			return &ValidationError{Name: "overview", err: fmt.Errorf(`ent: validator failed for field "overview": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Website(); !ok {
		return &ValidationError{Name: "website", err: errors.New(`ent: missing required field "website"`)}
	}
	if v, ok := cc.mutation.Website(); ok {
		if err := company.WebsiteValidator(v); err != nil {
			return &ValidationError{Name: "website", err: fmt.Errorf(`ent: validator failed for field "website": %w`, err)}
		}
	}
	if _, ok := cc.mutation.LogoURL(); !ok {
		return &ValidationError{Name: "logo_url", err: errors.New(`ent: missing required field "logo_url"`)}
	}
	if v, ok := cc.mutation.LogoURL(); ok {
		if err := company.LogoURLValidator(v); err != nil {
			return &ValidationError{Name: "logo_url", err: fmt.Errorf(`ent: validator failed for field "logo_url": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "size"`)}
	}
	if v, ok := cc.mutation.Size(); ok {
		if err := company.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf(`ent: validator failed for field "size": %w`, err)}
		}
	}
	if _, ok := cc.mutation.FoundedAt(); !ok {
		return &ValidationError{Name: "founded_at", err: errors.New(`ent: missing required field "founded_at"`)}
	}
	if v, ok := cc.mutation.FoundedAt(); ok {
		if err := company.FoundedAtValidator(v); err != nil {
			return &ValidationError{Name: "founded_at", err: fmt.Errorf(`ent: validator failed for field "founded_at": %w`, err)}
		}
	}
	return nil
}

func (cc *CompanyCreate) sqlSave(ctx context.Context) (*Company, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CompanyCreate) createSpec() (*Company, *sqlgraph.CreateSpec) {
	var (
		_node = &Company{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: company.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: company.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: company.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: company.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: company.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.Overview(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: company.FieldOverview,
		})
		_node.Overview = value
	}
	if value, ok := cc.mutation.Website(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: company.FieldWebsite,
		})
		_node.Website = value
	}
	if value, ok := cc.mutation.LogoURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: company.FieldLogoURL,
		})
		_node.LogoURL = value
	}
	if value, ok := cc.mutation.Size(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: company.FieldSize,
		})
		_node.Size = value
	}
	if value, ok := cc.mutation.FoundedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: company.FieldFoundedAt,
		})
		_node.FoundedAt = value
	}
	if nodes := cc.mutation.StaffsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.StaffsTable,
			Columns: []string{company.StaffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workexperience.FieldID,
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

// CompanyCreateBulk is the builder for creating many Company entities in bulk.
type CompanyCreateBulk struct {
	config
	builders []*CompanyCreate
}

// Save creates the Company entities in the database.
func (ccb *CompanyCreateBulk) Save(ctx context.Context) ([]*Company, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Company, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CompanyMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CompanyCreateBulk) SaveX(ctx context.Context) []*Company {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
