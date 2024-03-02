// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"app/app/model/schemas"
)

func newTest(db *gorm.DB, opts ...gen.DOOption) test {
	_test := test{}

	_test.testDo.UseDB(db, opts...)
	_test.testDo.UseModel(&schemas.Test{})

	tableName := _test.testDo.TableName()
	_test.ALL = field.NewAsterisk(tableName)
	_test.Name = field.NewString(tableName, "name")

	_test.fillFieldMap()

	return _test
}

type test struct {
	testDo

	ALL  field.Asterisk
	Name field.String

	fieldMap map[string]field.Expr
}

func (t test) Table(newTableName string) *test {
	t.testDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t test) As(alias string) *test {
	t.testDo.DO = *(t.testDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *test) updateTableName(table string) *test {
	t.ALL = field.NewAsterisk(table)
	t.Name = field.NewString(table, "name")

	t.fillFieldMap()

	return t
}

func (t *test) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *test) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 1)
	t.fieldMap["name"] = t.Name
}

func (t test) clone(db *gorm.DB) test {
	t.testDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t test) replaceDB(db *gorm.DB) test {
	t.testDo.ReplaceDB(db)
	return t
}

type testDo struct{ gen.DO }

type ITestDo interface {
	gen.SubQuery
	Debug() ITestDo
	WithContext(ctx context.Context) ITestDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITestDo
	WriteDB() ITestDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITestDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITestDo
	Not(conds ...gen.Condition) ITestDo
	Or(conds ...gen.Condition) ITestDo
	Select(conds ...field.Expr) ITestDo
	Where(conds ...gen.Condition) ITestDo
	Order(conds ...field.Expr) ITestDo
	Distinct(cols ...field.Expr) ITestDo
	Omit(cols ...field.Expr) ITestDo
	Join(table schema.Tabler, on ...field.Expr) ITestDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITestDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITestDo
	Group(cols ...field.Expr) ITestDo
	Having(conds ...gen.Condition) ITestDo
	Limit(limit int) ITestDo
	Offset(offset int) ITestDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITestDo
	Unscoped() ITestDo
	Create(values ...*schemas.Test) error
	CreateInBatches(values []*schemas.Test, batchSize int) error
	Save(values ...*schemas.Test) error
	First() (*schemas.Test, error)
	Take() (*schemas.Test, error)
	Last() (*schemas.Test, error)
	Find() ([]*schemas.Test, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*schemas.Test, err error)
	FindInBatches(result *[]*schemas.Test, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*schemas.Test) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITestDo
	Assign(attrs ...field.AssignExpr) ITestDo
	Joins(fields ...field.RelationField) ITestDo
	Preload(fields ...field.RelationField) ITestDo
	FirstOrInit() (*schemas.Test, error)
	FirstOrCreate() (*schemas.Test, error)
	FindByPage(offset int, limit int) (result []*schemas.Test, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITestDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByName(name string) (result []schemas.Test, err error)
}

// select * from @@table where name = @name
func (t testDo) GetByName(name string) (result []schemas.Test, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, name)
	generateSQL.WriteString("select * from tests where name = ? ")

	var executeSQL *gorm.DB
	executeSQL = t.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (t testDo) Debug() ITestDo {
	return t.withDO(t.DO.Debug())
}

func (t testDo) WithContext(ctx context.Context) ITestDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t testDo) ReadDB() ITestDo {
	return t.Clauses(dbresolver.Read)
}

func (t testDo) WriteDB() ITestDo {
	return t.Clauses(dbresolver.Write)
}

func (t testDo) Session(config *gorm.Session) ITestDo {
	return t.withDO(t.DO.Session(config))
}

func (t testDo) Clauses(conds ...clause.Expression) ITestDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t testDo) Returning(value interface{}, columns ...string) ITestDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t testDo) Not(conds ...gen.Condition) ITestDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t testDo) Or(conds ...gen.Condition) ITestDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t testDo) Select(conds ...field.Expr) ITestDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t testDo) Where(conds ...gen.Condition) ITestDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t testDo) Order(conds ...field.Expr) ITestDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t testDo) Distinct(cols ...field.Expr) ITestDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t testDo) Omit(cols ...field.Expr) ITestDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t testDo) Join(table schema.Tabler, on ...field.Expr) ITestDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t testDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITestDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t testDo) RightJoin(table schema.Tabler, on ...field.Expr) ITestDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t testDo) Group(cols ...field.Expr) ITestDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t testDo) Having(conds ...gen.Condition) ITestDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t testDo) Limit(limit int) ITestDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t testDo) Offset(offset int) ITestDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t testDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITestDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t testDo) Unscoped() ITestDo {
	return t.withDO(t.DO.Unscoped())
}

func (t testDo) Create(values ...*schemas.Test) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t testDo) CreateInBatches(values []*schemas.Test, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t testDo) Save(values ...*schemas.Test) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t testDo) First() (*schemas.Test, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*schemas.Test), nil
	}
}

func (t testDo) Take() (*schemas.Test, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*schemas.Test), nil
	}
}

func (t testDo) Last() (*schemas.Test, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*schemas.Test), nil
	}
}

func (t testDo) Find() ([]*schemas.Test, error) {
	result, err := t.DO.Find()
	return result.([]*schemas.Test), err
}

func (t testDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*schemas.Test, err error) {
	buf := make([]*schemas.Test, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t testDo) FindInBatches(result *[]*schemas.Test, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t testDo) Attrs(attrs ...field.AssignExpr) ITestDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t testDo) Assign(attrs ...field.AssignExpr) ITestDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t testDo) Joins(fields ...field.RelationField) ITestDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t testDo) Preload(fields ...field.RelationField) ITestDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t testDo) FirstOrInit() (*schemas.Test, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*schemas.Test), nil
	}
}

func (t testDo) FirstOrCreate() (*schemas.Test, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*schemas.Test), nil
	}
}

func (t testDo) FindByPage(offset int, limit int) (result []*schemas.Test, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t testDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t testDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t testDo) Delete(models ...*schemas.Test) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *testDo) withDO(do gen.Dao) *testDo {
	t.DO = *do.(*gen.DO)
	return t
}
