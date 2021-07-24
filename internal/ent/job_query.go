// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"refernet/internal/ent/job"
	"refernet/internal/ent/predicate"
	"refernet/internal/ent/skill"
	"refernet/internal/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JobQuery is the builder for querying Job entities.
type JobQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Job
	// eager-loading edges.
	withOwner  *UserQuery
	withSkills *SkillQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the JobQuery builder.
func (jq *JobQuery) Where(ps ...predicate.Job) *JobQuery {
	jq.predicates = append(jq.predicates, ps...)
	return jq
}

// Limit adds a limit step to the query.
func (jq *JobQuery) Limit(limit int) *JobQuery {
	jq.limit = &limit
	return jq
}

// Offset adds an offset step to the query.
func (jq *JobQuery) Offset(offset int) *JobQuery {
	jq.offset = &offset
	return jq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (jq *JobQuery) Unique(unique bool) *JobQuery {
	jq.unique = &unique
	return jq
}

// Order adds an order step to the query.
func (jq *JobQuery) Order(o ...OrderFunc) *JobQuery {
	jq.order = append(jq.order, o...)
	return jq
}

// QueryOwner chains the current query on the "owner" edge.
func (jq *JobQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: jq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := jq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := jq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(job.Table, job.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, job.OwnerTable, job.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(jq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySkills chains the current query on the "skills" edge.
func (jq *JobQuery) QuerySkills() *SkillQuery {
	query := &SkillQuery{config: jq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := jq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := jq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(job.Table, job.FieldID, selector),
			sqlgraph.To(skill.Table, skill.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, job.SkillsTable, job.SkillsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(jq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Job entity from the query.
// Returns a *NotFoundError when no Job was found.
func (jq *JobQuery) First(ctx context.Context) (*Job, error) {
	nodes, err := jq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{job.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (jq *JobQuery) FirstX(ctx context.Context) *Job {
	node, err := jq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Job ID from the query.
// Returns a *NotFoundError when no Job ID was found.
func (jq *JobQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{job.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (jq *JobQuery) FirstIDX(ctx context.Context) int {
	id, err := jq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Job entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Job entity is not found.
// Returns a *NotFoundError when no Job entities are found.
func (jq *JobQuery) Only(ctx context.Context) (*Job, error) {
	nodes, err := jq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{job.Label}
	default:
		return nil, &NotSingularError{job.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (jq *JobQuery) OnlyX(ctx context.Context) *Job {
	node, err := jq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Job ID in the query.
// Returns a *NotSingularError when exactly one Job ID is not found.
// Returns a *NotFoundError when no entities are found.
func (jq *JobQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{job.Label}
	default:
		err = &NotSingularError{job.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (jq *JobQuery) OnlyIDX(ctx context.Context) int {
	id, err := jq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Jobs.
func (jq *JobQuery) All(ctx context.Context) ([]*Job, error) {
	if err := jq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return jq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (jq *JobQuery) AllX(ctx context.Context) []*Job {
	nodes, err := jq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Job IDs.
func (jq *JobQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := jq.Select(job.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (jq *JobQuery) IDsX(ctx context.Context) []int {
	ids, err := jq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (jq *JobQuery) Count(ctx context.Context) (int, error) {
	if err := jq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return jq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (jq *JobQuery) CountX(ctx context.Context) int {
	count, err := jq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (jq *JobQuery) Exist(ctx context.Context) (bool, error) {
	if err := jq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return jq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (jq *JobQuery) ExistX(ctx context.Context) bool {
	exist, err := jq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the JobQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (jq *JobQuery) Clone() *JobQuery {
	if jq == nil {
		return nil
	}
	return &JobQuery{
		config:     jq.config,
		limit:      jq.limit,
		offset:     jq.offset,
		order:      append([]OrderFunc{}, jq.order...),
		predicates: append([]predicate.Job{}, jq.predicates...),
		withOwner:  jq.withOwner.Clone(),
		withSkills: jq.withSkills.Clone(),
		// clone intermediate query.
		sql:  jq.sql.Clone(),
		path: jq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (jq *JobQuery) WithOwner(opts ...func(*UserQuery)) *JobQuery {
	query := &UserQuery{config: jq.config}
	for _, opt := range opts {
		opt(query)
	}
	jq.withOwner = query
	return jq
}

// WithSkills tells the query-builder to eager-load the nodes that are connected to
// the "skills" edge. The optional arguments are used to configure the query builder of the edge.
func (jq *JobQuery) WithSkills(opts ...func(*SkillQuery)) *JobQuery {
	query := &SkillQuery{config: jq.config}
	for _, opt := range opts {
		opt(query)
	}
	jq.withSkills = query
	return jq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Job.Query().
//		GroupBy(job.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (jq *JobQuery) GroupBy(field string, fields ...string) *JobGroupBy {
	group := &JobGroupBy{config: jq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := jq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return jq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Job.Query().
//		Select(job.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (jq *JobQuery) Select(fields ...string) *JobSelect {
	jq.fields = append(jq.fields, fields...)
	return &JobSelect{JobQuery: jq}
}

func (jq *JobQuery) prepareQuery(ctx context.Context) error {
	for _, f := range jq.fields {
		if !job.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if jq.path != nil {
		prev, err := jq.path(ctx)
		if err != nil {
			return err
		}
		jq.sql = prev
	}
	return nil
}

func (jq *JobQuery) sqlAll(ctx context.Context) ([]*Job, error) {
	var (
		nodes       = []*Job{}
		withFKs     = jq.withFKs
		_spec       = jq.querySpec()
		loadedTypes = [2]bool{
			jq.withOwner != nil,
			jq.withSkills != nil,
		}
	)
	if jq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, job.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Job{config: jq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, jq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := jq.withOwner; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Job)
		for i := range nodes {
			if nodes[i].user_jobs == nil {
				continue
			}
			fk := *nodes[i].user_jobs
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_jobs" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Owner = n
			}
		}
	}

	if query := jq.withSkills; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Job, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Skills = []*Skill{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*Job)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   job.SkillsTable,
				Columns: job.SkillsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(job.SkillsPrimaryKey[1], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, jq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "skills": %w`, err)
		}
		query.Where(skill.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "skills" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Skills = append(nodes[i].Edges.Skills, n)
			}
		}
	}

	return nodes, nil
}

func (jq *JobQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := jq.querySpec()
	return sqlgraph.CountNodes(ctx, jq.driver, _spec)
}

func (jq *JobQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := jq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (jq *JobQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   job.Table,
			Columns: job.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: job.FieldID,
			},
		},
		From:   jq.sql,
		Unique: true,
	}
	if unique := jq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := jq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, job.FieldID)
		for i := range fields {
			if fields[i] != job.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := jq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := jq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := jq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := jq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (jq *JobQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(jq.driver.Dialect())
	t1 := builder.Table(job.Table)
	columns := jq.fields
	if len(columns) == 0 {
		columns = job.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if jq.sql != nil {
		selector = jq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range jq.predicates {
		p(selector)
	}
	for _, p := range jq.order {
		p(selector)
	}
	if offset := jq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := jq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// JobGroupBy is the group-by builder for Job entities.
type JobGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (jgb *JobGroupBy) Aggregate(fns ...AggregateFunc) *JobGroupBy {
	jgb.fns = append(jgb.fns, fns...)
	return jgb
}

// Scan applies the group-by query and scans the result into the given value.
func (jgb *JobGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := jgb.path(ctx)
	if err != nil {
		return err
	}
	jgb.sql = query
	return jgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (jgb *JobGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := jgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (jgb *JobGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(jgb.fields) > 1 {
		return nil, errors.New("ent: JobGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := jgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (jgb *JobGroupBy) StringsX(ctx context.Context) []string {
	v, err := jgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (jgb *JobGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = jgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{job.Label}
	default:
		err = fmt.Errorf("ent: JobGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (jgb *JobGroupBy) StringX(ctx context.Context) string {
	v, err := jgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (jgb *JobGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(jgb.fields) > 1 {
		return nil, errors.New("ent: JobGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := jgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (jgb *JobGroupBy) IntsX(ctx context.Context) []int {
	v, err := jgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (jgb *JobGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = jgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{job.Label}
	default:
		err = fmt.Errorf("ent: JobGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (jgb *JobGroupBy) IntX(ctx context.Context) int {
	v, err := jgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (jgb *JobGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(jgb.fields) > 1 {
		return nil, errors.New("ent: JobGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := jgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (jgb *JobGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := jgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (jgb *JobGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = jgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{job.Label}
	default:
		err = fmt.Errorf("ent: JobGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (jgb *JobGroupBy) Float64X(ctx context.Context) float64 {
	v, err := jgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (jgb *JobGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(jgb.fields) > 1 {
		return nil, errors.New("ent: JobGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := jgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (jgb *JobGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := jgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (jgb *JobGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = jgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{job.Label}
	default:
		err = fmt.Errorf("ent: JobGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (jgb *JobGroupBy) BoolX(ctx context.Context) bool {
	v, err := jgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (jgb *JobGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range jgb.fields {
		if !job.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := jgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := jgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (jgb *JobGroupBy) sqlQuery() *sql.Selector {
	selector := jgb.sql.Select()
	aggregation := make([]string, 0, len(jgb.fns))
	for _, fn := range jgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(jgb.fields)+len(jgb.fns))
		for _, f := range jgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(jgb.fields...)...)
}

// JobSelect is the builder for selecting fields of Job entities.
type JobSelect struct {
	*JobQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (js *JobSelect) Scan(ctx context.Context, v interface{}) error {
	if err := js.prepareQuery(ctx); err != nil {
		return err
	}
	js.sql = js.JobQuery.sqlQuery(ctx)
	return js.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (js *JobSelect) ScanX(ctx context.Context, v interface{}) {
	if err := js.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (js *JobSelect) Strings(ctx context.Context) ([]string, error) {
	if len(js.fields) > 1 {
		return nil, errors.New("ent: JobSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := js.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (js *JobSelect) StringsX(ctx context.Context) []string {
	v, err := js.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (js *JobSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = js.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{job.Label}
	default:
		err = fmt.Errorf("ent: JobSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (js *JobSelect) StringX(ctx context.Context) string {
	v, err := js.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (js *JobSelect) Ints(ctx context.Context) ([]int, error) {
	if len(js.fields) > 1 {
		return nil, errors.New("ent: JobSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := js.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (js *JobSelect) IntsX(ctx context.Context) []int {
	v, err := js.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (js *JobSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = js.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{job.Label}
	default:
		err = fmt.Errorf("ent: JobSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (js *JobSelect) IntX(ctx context.Context) int {
	v, err := js.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (js *JobSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(js.fields) > 1 {
		return nil, errors.New("ent: JobSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := js.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (js *JobSelect) Float64sX(ctx context.Context) []float64 {
	v, err := js.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (js *JobSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = js.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{job.Label}
	default:
		err = fmt.Errorf("ent: JobSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (js *JobSelect) Float64X(ctx context.Context) float64 {
	v, err := js.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (js *JobSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(js.fields) > 1 {
		return nil, errors.New("ent: JobSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := js.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (js *JobSelect) BoolsX(ctx context.Context) []bool {
	v, err := js.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (js *JobSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = js.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{job.Label}
	default:
		err = fmt.Errorf("ent: JobSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (js *JobSelect) BoolX(ctx context.Context) bool {
	v, err := js.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (js *JobSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := js.sql.Query()
	if err := js.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
