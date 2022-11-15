// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/gnasnik/titan-operator/generated/model"
)

func newLoginLog(db *gorm.DB, opts ...gen.DOOption) loginLog {
	_loginLog := loginLog{}

	_loginLog.loginLogDo.UseDB(db, opts...)
	_loginLog.loginLogDo.UseModel(&model.LoginLog{})

	tableName := _loginLog.loginLogDo.TableName()
	_loginLog.ALL = field.NewAsterisk(tableName)
	_loginLog.ID = field.NewInt64(tableName, "id")
	_loginLog.LoginUsername = field.NewString(tableName, "login_username")
	_loginLog.Ipaddr = field.NewString(tableName, "ipaddr")
	_loginLog.LoginLocation = field.NewString(tableName, "login_location")
	_loginLog.Browser = field.NewString(tableName, "browser")
	_loginLog.Os = field.NewString(tableName, "os")
	_loginLog.Status = field.NewInt32(tableName, "status")
	_loginLog.Msg = field.NewString(tableName, "msg")
	_loginLog.CreatedAt = field.NewTime(tableName, "created_at")

	_loginLog.fillFieldMap()

	return _loginLog
}

type loginLog struct {
	loginLogDo

	ALL           field.Asterisk
	ID            field.Int64
	LoginUsername field.String
	Ipaddr        field.String
	LoginLocation field.String
	Browser       field.String
	Os            field.String
	Status        field.Int32
	Msg           field.String
	CreatedAt     field.Time

	fieldMap map[string]field.Expr
}

func (l loginLog) Table(newTableName string) *loginLog {
	l.loginLogDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l loginLog) As(alias string) *loginLog {
	l.loginLogDo.DO = *(l.loginLogDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *loginLog) updateTableName(table string) *loginLog {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewInt64(table, "id")
	l.LoginUsername = field.NewString(table, "login_username")
	l.Ipaddr = field.NewString(table, "ipaddr")
	l.LoginLocation = field.NewString(table, "login_location")
	l.Browser = field.NewString(table, "browser")
	l.Os = field.NewString(table, "os")
	l.Status = field.NewInt32(table, "status")
	l.Msg = field.NewString(table, "msg")
	l.CreatedAt = field.NewTime(table, "created_at")

	l.fillFieldMap()

	return l
}

func (l *loginLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *loginLog) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 9)
	l.fieldMap["id"] = l.ID
	l.fieldMap["login_username"] = l.LoginUsername
	l.fieldMap["ipaddr"] = l.Ipaddr
	l.fieldMap["login_location"] = l.LoginLocation
	l.fieldMap["browser"] = l.Browser
	l.fieldMap["os"] = l.Os
	l.fieldMap["status"] = l.Status
	l.fieldMap["msg"] = l.Msg
	l.fieldMap["created_at"] = l.CreatedAt
}

func (l loginLog) clone(db *gorm.DB) loginLog {
	l.loginLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l loginLog) replaceDB(db *gorm.DB) loginLog {
	l.loginLogDo.ReplaceDB(db)
	return l
}

type loginLogDo struct{ gen.DO }

type ILoginLogDo interface {
	gen.SubQuery
	Debug() ILoginLogDo
	WithContext(ctx context.Context) ILoginLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILoginLogDo
	WriteDB() ILoginLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILoginLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILoginLogDo
	Not(conds ...gen.Condition) ILoginLogDo
	Or(conds ...gen.Condition) ILoginLogDo
	Select(conds ...field.Expr) ILoginLogDo
	Where(conds ...gen.Condition) ILoginLogDo
	Order(conds ...field.Expr) ILoginLogDo
	Distinct(cols ...field.Expr) ILoginLogDo
	Omit(cols ...field.Expr) ILoginLogDo
	Join(table schema.Tabler, on ...field.Expr) ILoginLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILoginLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILoginLogDo
	Group(cols ...field.Expr) ILoginLogDo
	Having(conds ...gen.Condition) ILoginLogDo
	Limit(limit int) ILoginLogDo
	Offset(offset int) ILoginLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILoginLogDo
	Unscoped() ILoginLogDo
	Create(values ...*model.LoginLog) error
	CreateInBatches(values []*model.LoginLog, batchSize int) error
	Save(values ...*model.LoginLog) error
	First() (*model.LoginLog, error)
	Take() (*model.LoginLog, error)
	Last() (*model.LoginLog, error)
	Find() ([]*model.LoginLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LoginLog, err error)
	FindInBatches(result *[]*model.LoginLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.LoginLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILoginLogDo
	Assign(attrs ...field.AssignExpr) ILoginLogDo
	Joins(fields ...field.RelationField) ILoginLogDo
	Preload(fields ...field.RelationField) ILoginLogDo
	FirstOrInit() (*model.LoginLog, error)
	FirstOrCreate() (*model.LoginLog, error)
	FindByPage(offset int, limit int) (result []*model.LoginLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILoginLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l loginLogDo) Debug() ILoginLogDo {
	return l.withDO(l.DO.Debug())
}

func (l loginLogDo) WithContext(ctx context.Context) ILoginLogDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l loginLogDo) ReadDB() ILoginLogDo {
	return l.Clauses(dbresolver.Read)
}

func (l loginLogDo) WriteDB() ILoginLogDo {
	return l.Clauses(dbresolver.Write)
}

func (l loginLogDo) Session(config *gorm.Session) ILoginLogDo {
	return l.withDO(l.DO.Session(config))
}

func (l loginLogDo) Clauses(conds ...clause.Expression) ILoginLogDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l loginLogDo) Returning(value interface{}, columns ...string) ILoginLogDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l loginLogDo) Not(conds ...gen.Condition) ILoginLogDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l loginLogDo) Or(conds ...gen.Condition) ILoginLogDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l loginLogDo) Select(conds ...field.Expr) ILoginLogDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l loginLogDo) Where(conds ...gen.Condition) ILoginLogDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l loginLogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ILoginLogDo {
	return l.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (l loginLogDo) Order(conds ...field.Expr) ILoginLogDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l loginLogDo) Distinct(cols ...field.Expr) ILoginLogDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l loginLogDo) Omit(cols ...field.Expr) ILoginLogDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l loginLogDo) Join(table schema.Tabler, on ...field.Expr) ILoginLogDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l loginLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILoginLogDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l loginLogDo) RightJoin(table schema.Tabler, on ...field.Expr) ILoginLogDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l loginLogDo) Group(cols ...field.Expr) ILoginLogDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l loginLogDo) Having(conds ...gen.Condition) ILoginLogDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l loginLogDo) Limit(limit int) ILoginLogDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l loginLogDo) Offset(offset int) ILoginLogDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l loginLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILoginLogDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l loginLogDo) Unscoped() ILoginLogDo {
	return l.withDO(l.DO.Unscoped())
}

func (l loginLogDo) Create(values ...*model.LoginLog) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l loginLogDo) CreateInBatches(values []*model.LoginLog, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l loginLogDo) Save(values ...*model.LoginLog) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l loginLogDo) First() (*model.LoginLog, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoginLog), nil
	}
}

func (l loginLogDo) Take() (*model.LoginLog, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoginLog), nil
	}
}

func (l loginLogDo) Last() (*model.LoginLog, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoginLog), nil
	}
}

func (l loginLogDo) Find() ([]*model.LoginLog, error) {
	result, err := l.DO.Find()
	return result.([]*model.LoginLog), err
}

func (l loginLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LoginLog, err error) {
	buf := make([]*model.LoginLog, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l loginLogDo) FindInBatches(result *[]*model.LoginLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l loginLogDo) Attrs(attrs ...field.AssignExpr) ILoginLogDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l loginLogDo) Assign(attrs ...field.AssignExpr) ILoginLogDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l loginLogDo) Joins(fields ...field.RelationField) ILoginLogDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l loginLogDo) Preload(fields ...field.RelationField) ILoginLogDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l loginLogDo) FirstOrInit() (*model.LoginLog, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoginLog), nil
	}
}

func (l loginLogDo) FirstOrCreate() (*model.LoginLog, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoginLog), nil
	}
}

func (l loginLogDo) FindByPage(offset int, limit int) (result []*model.LoginLog, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l loginLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l loginLogDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l loginLogDo) Delete(models ...*model.LoginLog) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *loginLogDo) withDO(do gen.Dao) *loginLogDo {
	l.DO = *do.(*gen.DO)
	return l
}
