// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"refernet/internal/ent/company"
	"refernet/internal/ent/predicate"
	"refernet/internal/ent/skill"
	"refernet/internal/ent/user"
	"refernet/internal/ent/workexperience"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorkExperienceUpdate is the builder for updating WorkExperience entities.
type WorkExperienceUpdate struct {
	config
	hooks    []Hook
	mutation *WorkExperienceMutation
}

// Where adds a new predicate for the WorkExperienceUpdate builder.
func (weu *WorkExperienceUpdate) Where(ps ...predicate.WorkExperience) *WorkExperienceUpdate {
	weu.mutation.predicates = append(weu.mutation.predicates, ps...)
	return weu
}

// SetUpdatedAt sets the "updated_at" field.
func (weu *WorkExperienceUpdate) SetUpdatedAt(t time.Time) *WorkExperienceUpdate {
	weu.mutation.SetUpdatedAt(t)
	return weu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (weu *WorkExperienceUpdate) SetNillableUpdatedAt(t *time.Time) *WorkExperienceUpdate {
	if t != nil {
		weu.SetUpdatedAt(*t)
	}
	return weu
}

// SetTitle sets the "title" field.
func (weu *WorkExperienceUpdate) SetTitle(s string) *WorkExperienceUpdate {
	weu.mutation.SetTitle(s)
	return weu
}

// SetLocation sets the "location" field.
func (weu *WorkExperienceUpdate) SetLocation(s string) *WorkExperienceUpdate {
	weu.mutation.SetLocation(s)
	return weu
}

// SetStartDate sets the "start_date" field.
func (weu *WorkExperienceUpdate) SetStartDate(t time.Time) *WorkExperienceUpdate {
	weu.mutation.SetStartDate(t)
	return weu
}

// SetEndDate sets the "end_date" field.
func (weu *WorkExperienceUpdate) SetEndDate(t time.Time) *WorkExperienceUpdate {
	weu.mutation.SetEndDate(t)
	return weu
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (weu *WorkExperienceUpdate) SetNillableEndDate(t *time.Time) *WorkExperienceUpdate {
	if t != nil {
		weu.SetEndDate(*t)
	}
	return weu
}

// ClearEndDate clears the value of the "end_date" field.
func (weu *WorkExperienceUpdate) ClearEndDate() *WorkExperienceUpdate {
	weu.mutation.ClearEndDate()
	return weu
}

// SetDescription sets the "description" field.
func (weu *WorkExperienceUpdate) SetDescription(s string) *WorkExperienceUpdate {
	weu.mutation.SetDescription(s)
	return weu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (weu *WorkExperienceUpdate) SetOwnerID(id int) *WorkExperienceUpdate {
	weu.mutation.SetOwnerID(id)
	return weu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (weu *WorkExperienceUpdate) SetNillableOwnerID(id *int) *WorkExperienceUpdate {
	if id != nil {
		weu = weu.SetOwnerID(*id)
	}
	return weu
}

// SetOwner sets the "owner" edge to the User entity.
func (weu *WorkExperienceUpdate) SetOwner(u *User) *WorkExperienceUpdate {
	return weu.SetOwnerID(u.ID)
}

// SetInCompanyID sets the "in_company" edge to the Company entity by ID.
func (weu *WorkExperienceUpdate) SetInCompanyID(id int) *WorkExperienceUpdate {
	weu.mutation.SetInCompanyID(id)
	return weu
}

// SetNillableInCompanyID sets the "in_company" edge to the Company entity by ID if the given value is not nil.
func (weu *WorkExperienceUpdate) SetNillableInCompanyID(id *int) *WorkExperienceUpdate {
	if id != nil {
		weu = weu.SetInCompanyID(*id)
	}
	return weu
}

// SetInCompany sets the "in_company" edge to the Company entity.
func (weu *WorkExperienceUpdate) SetInCompany(c *Company) *WorkExperienceUpdate {
	return weu.SetInCompanyID(c.ID)
}

// AddSkillIDs adds the "skills" edge to the Skill entity by IDs.
func (weu *WorkExperienceUpdate) AddSkillIDs(ids ...int) *WorkExperienceUpdate {
	weu.mutation.AddSkillIDs(ids...)
	return weu
}

// AddSkills adds the "skills" edges to the Skill entity.
func (weu *WorkExperienceUpdate) AddSkills(s ...*Skill) *WorkExperienceUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return weu.AddSkillIDs(ids...)
}

// Mutation returns the WorkExperienceMutation object of the builder.
func (weu *WorkExperienceUpdate) Mutation() *WorkExperienceMutation {
	return weu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (weu *WorkExperienceUpdate) ClearOwner() *WorkExperienceUpdate {
	weu.mutation.ClearOwner()
	return weu
}

// ClearInCompany clears the "in_company" edge to the Company entity.
func (weu *WorkExperienceUpdate) ClearInCompany() *WorkExperienceUpdate {
	weu.mutation.ClearInCompany()
	return weu
}

// ClearSkills clears all "skills" edges to the Skill entity.
func (weu *WorkExperienceUpdate) ClearSkills() *WorkExperienceUpdate {
	weu.mutation.ClearSkills()
	return weu
}

// RemoveSkillIDs removes the "skills" edge to Skill entities by IDs.
func (weu *WorkExperienceUpdate) RemoveSkillIDs(ids ...int) *WorkExperienceUpdate {
	weu.mutation.RemoveSkillIDs(ids...)
	return weu
}

// RemoveSkills removes "skills" edges to Skill entities.
func (weu *WorkExperienceUpdate) RemoveSkills(s ...*Skill) *WorkExperienceUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return weu.RemoveSkillIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (weu *WorkExperienceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(weu.hooks) == 0 {
		if err = weu.check(); err != nil {
			return 0, err
		}
		affected, err = weu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkExperienceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = weu.check(); err != nil {
				return 0, err
			}
			weu.mutation = mutation
			affected, err = weu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(weu.hooks) - 1; i >= 0; i-- {
			mut = weu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, weu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (weu *WorkExperienceUpdate) SaveX(ctx context.Context) int {
	affected, err := weu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (weu *WorkExperienceUpdate) Exec(ctx context.Context) error {
	_, err := weu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (weu *WorkExperienceUpdate) ExecX(ctx context.Context) {
	if err := weu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (weu *WorkExperienceUpdate) check() error {
	if v, ok := weu.mutation.Title(); ok {
		if err := workexperience.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := weu.mutation.Location(); ok {
		if err := workexperience.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf("ent: validator failed for field \"location\": %w", err)}
		}
	}
	if v, ok := weu.mutation.Description(); ok {
		if err := workexperience.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf("ent: validator failed for field \"description\": %w", err)}
		}
	}
	return nil
}

func (weu *WorkExperienceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workexperience.Table,
			Columns: workexperience.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workexperience.FieldID,
			},
		},
	}
	if ps := weu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := weu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workexperience.FieldUpdatedAt,
		})
	}
	if value, ok := weu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workexperience.FieldTitle,
		})
	}
	if value, ok := weu.mutation.Location(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workexperience.FieldLocation,
		})
	}
	if value, ok := weu.mutation.StartDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workexperience.FieldStartDate,
		})
	}
	if value, ok := weu.mutation.EndDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workexperience.FieldEndDate,
		})
	}
	if weu.mutation.EndDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: workexperience.FieldEndDate,
		})
	}
	if value, ok := weu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workexperience.FieldDescription,
		})
	}
	if weu.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := weu.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if weu.mutation.InCompanyCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := weu.mutation.InCompanyIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if weu.mutation.SkillsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := weu.mutation.RemovedSkillsIDs(); len(nodes) > 0 && !weu.mutation.SkillsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := weu.mutation.SkillsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, weu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workexperience.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// WorkExperienceUpdateOne is the builder for updating a single WorkExperience entity.
type WorkExperienceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkExperienceMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (weuo *WorkExperienceUpdateOne) SetUpdatedAt(t time.Time) *WorkExperienceUpdateOne {
	weuo.mutation.SetUpdatedAt(t)
	return weuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (weuo *WorkExperienceUpdateOne) SetNillableUpdatedAt(t *time.Time) *WorkExperienceUpdateOne {
	if t != nil {
		weuo.SetUpdatedAt(*t)
	}
	return weuo
}

// SetTitle sets the "title" field.
func (weuo *WorkExperienceUpdateOne) SetTitle(s string) *WorkExperienceUpdateOne {
	weuo.mutation.SetTitle(s)
	return weuo
}

// SetLocation sets the "location" field.
func (weuo *WorkExperienceUpdateOne) SetLocation(s string) *WorkExperienceUpdateOne {
	weuo.mutation.SetLocation(s)
	return weuo
}

// SetStartDate sets the "start_date" field.
func (weuo *WorkExperienceUpdateOne) SetStartDate(t time.Time) *WorkExperienceUpdateOne {
	weuo.mutation.SetStartDate(t)
	return weuo
}

// SetEndDate sets the "end_date" field.
func (weuo *WorkExperienceUpdateOne) SetEndDate(t time.Time) *WorkExperienceUpdateOne {
	weuo.mutation.SetEndDate(t)
	return weuo
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (weuo *WorkExperienceUpdateOne) SetNillableEndDate(t *time.Time) *WorkExperienceUpdateOne {
	if t != nil {
		weuo.SetEndDate(*t)
	}
	return weuo
}

// ClearEndDate clears the value of the "end_date" field.
func (weuo *WorkExperienceUpdateOne) ClearEndDate() *WorkExperienceUpdateOne {
	weuo.mutation.ClearEndDate()
	return weuo
}

// SetDescription sets the "description" field.
func (weuo *WorkExperienceUpdateOne) SetDescription(s string) *WorkExperienceUpdateOne {
	weuo.mutation.SetDescription(s)
	return weuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (weuo *WorkExperienceUpdateOne) SetOwnerID(id int) *WorkExperienceUpdateOne {
	weuo.mutation.SetOwnerID(id)
	return weuo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (weuo *WorkExperienceUpdateOne) SetNillableOwnerID(id *int) *WorkExperienceUpdateOne {
	if id != nil {
		weuo = weuo.SetOwnerID(*id)
	}
	return weuo
}

// SetOwner sets the "owner" edge to the User entity.
func (weuo *WorkExperienceUpdateOne) SetOwner(u *User) *WorkExperienceUpdateOne {
	return weuo.SetOwnerID(u.ID)
}

// SetInCompanyID sets the "in_company" edge to the Company entity by ID.
func (weuo *WorkExperienceUpdateOne) SetInCompanyID(id int) *WorkExperienceUpdateOne {
	weuo.mutation.SetInCompanyID(id)
	return weuo
}

// SetNillableInCompanyID sets the "in_company" edge to the Company entity by ID if the given value is not nil.
func (weuo *WorkExperienceUpdateOne) SetNillableInCompanyID(id *int) *WorkExperienceUpdateOne {
	if id != nil {
		weuo = weuo.SetInCompanyID(*id)
	}
	return weuo
}

// SetInCompany sets the "in_company" edge to the Company entity.
func (weuo *WorkExperienceUpdateOne) SetInCompany(c *Company) *WorkExperienceUpdateOne {
	return weuo.SetInCompanyID(c.ID)
}

// AddSkillIDs adds the "skills" edge to the Skill entity by IDs.
func (weuo *WorkExperienceUpdateOne) AddSkillIDs(ids ...int) *WorkExperienceUpdateOne {
	weuo.mutation.AddSkillIDs(ids...)
	return weuo
}

// AddSkills adds the "skills" edges to the Skill entity.
func (weuo *WorkExperienceUpdateOne) AddSkills(s ...*Skill) *WorkExperienceUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return weuo.AddSkillIDs(ids...)
}

// Mutation returns the WorkExperienceMutation object of the builder.
func (weuo *WorkExperienceUpdateOne) Mutation() *WorkExperienceMutation {
	return weuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (weuo *WorkExperienceUpdateOne) ClearOwner() *WorkExperienceUpdateOne {
	weuo.mutation.ClearOwner()
	return weuo
}

// ClearInCompany clears the "in_company" edge to the Company entity.
func (weuo *WorkExperienceUpdateOne) ClearInCompany() *WorkExperienceUpdateOne {
	weuo.mutation.ClearInCompany()
	return weuo
}

// ClearSkills clears all "skills" edges to the Skill entity.
func (weuo *WorkExperienceUpdateOne) ClearSkills() *WorkExperienceUpdateOne {
	weuo.mutation.ClearSkills()
	return weuo
}

// RemoveSkillIDs removes the "skills" edge to Skill entities by IDs.
func (weuo *WorkExperienceUpdateOne) RemoveSkillIDs(ids ...int) *WorkExperienceUpdateOne {
	weuo.mutation.RemoveSkillIDs(ids...)
	return weuo
}

// RemoveSkills removes "skills" edges to Skill entities.
func (weuo *WorkExperienceUpdateOne) RemoveSkills(s ...*Skill) *WorkExperienceUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return weuo.RemoveSkillIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (weuo *WorkExperienceUpdateOne) Select(field string, fields ...string) *WorkExperienceUpdateOne {
	weuo.fields = append([]string{field}, fields...)
	return weuo
}

// Save executes the query and returns the updated WorkExperience entity.
func (weuo *WorkExperienceUpdateOne) Save(ctx context.Context) (*WorkExperience, error) {
	var (
		err  error
		node *WorkExperience
	)
	if len(weuo.hooks) == 0 {
		if err = weuo.check(); err != nil {
			return nil, err
		}
		node, err = weuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkExperienceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = weuo.check(); err != nil {
				return nil, err
			}
			weuo.mutation = mutation
			node, err = weuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(weuo.hooks) - 1; i >= 0; i-- {
			mut = weuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, weuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (weuo *WorkExperienceUpdateOne) SaveX(ctx context.Context) *WorkExperience {
	node, err := weuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (weuo *WorkExperienceUpdateOne) Exec(ctx context.Context) error {
	_, err := weuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (weuo *WorkExperienceUpdateOne) ExecX(ctx context.Context) {
	if err := weuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (weuo *WorkExperienceUpdateOne) check() error {
	if v, ok := weuo.mutation.Title(); ok {
		if err := workexperience.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := weuo.mutation.Location(); ok {
		if err := workexperience.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf("ent: validator failed for field \"location\": %w", err)}
		}
	}
	if v, ok := weuo.mutation.Description(); ok {
		if err := workexperience.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf("ent: validator failed for field \"description\": %w", err)}
		}
	}
	return nil
}

func (weuo *WorkExperienceUpdateOne) sqlSave(ctx context.Context) (_node *WorkExperience, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workexperience.Table,
			Columns: workexperience.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workexperience.FieldID,
			},
		},
	}
	id, ok := weuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing WorkExperience.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := weuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workexperience.FieldID)
		for _, f := range fields {
			if !workexperience.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workexperience.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := weuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := weuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workexperience.FieldUpdatedAt,
		})
	}
	if value, ok := weuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workexperience.FieldTitle,
		})
	}
	if value, ok := weuo.mutation.Location(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workexperience.FieldLocation,
		})
	}
	if value, ok := weuo.mutation.StartDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workexperience.FieldStartDate,
		})
	}
	if value, ok := weuo.mutation.EndDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workexperience.FieldEndDate,
		})
	}
	if weuo.mutation.EndDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: workexperience.FieldEndDate,
		})
	}
	if value, ok := weuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workexperience.FieldDescription,
		})
	}
	if weuo.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := weuo.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if weuo.mutation.InCompanyCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := weuo.mutation.InCompanyIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if weuo.mutation.SkillsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := weuo.mutation.RemovedSkillsIDs(); len(nodes) > 0 && !weuo.mutation.SkillsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := weuo.mutation.SkillsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &WorkExperience{config: weuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, weuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workexperience.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
