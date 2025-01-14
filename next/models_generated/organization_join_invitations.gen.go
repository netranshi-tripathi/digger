// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models_generated

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/diggerhq/digger/next/model"
)

func newOrganizationJoinInvitation(db *gorm.DB, opts ...gen.DOOption) organizationJoinInvitation {
	_organizationJoinInvitation := organizationJoinInvitation{}

	_organizationJoinInvitation.organizationJoinInvitationDo.UseDB(db, opts...)
	_organizationJoinInvitation.organizationJoinInvitationDo.UseModel(&model.OrganizationJoinInvitation{})

	tableName := _organizationJoinInvitation.organizationJoinInvitationDo.TableName()
	_organizationJoinInvitation.ALL = field.NewAsterisk(tableName)
	_organizationJoinInvitation.CreatedAt = field.NewTime(tableName, "created_at")
	_organizationJoinInvitation.InviterUserID = field.NewString(tableName, "inviter_user_id")
	_organizationJoinInvitation.Status = field.NewString(tableName, "status")
	_organizationJoinInvitation.ID = field.NewString(tableName, "id")
	_organizationJoinInvitation.InviteeUserEmail = field.NewString(tableName, "invitee_user_email")
	_organizationJoinInvitation.OrganizationID = field.NewString(tableName, "organization_id")
	_organizationJoinInvitation.InviteeOrganizationRole = field.NewString(tableName, "invitee_organization_role")
	_organizationJoinInvitation.InviteeUserID = field.NewString(tableName, "invitee_user_id")

	_organizationJoinInvitation.fillFieldMap()

	return _organizationJoinInvitation
}

type organizationJoinInvitation struct {
	organizationJoinInvitationDo

	ALL                     field.Asterisk
	CreatedAt               field.Time
	InviterUserID           field.String
	Status                  field.String
	ID                      field.String
	InviteeUserEmail        field.String
	OrganizationID          field.String
	InviteeOrganizationRole field.String
	InviteeUserID           field.String

	fieldMap map[string]field.Expr
}

func (o organizationJoinInvitation) Table(newTableName string) *organizationJoinInvitation {
	o.organizationJoinInvitationDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o organizationJoinInvitation) As(alias string) *organizationJoinInvitation {
	o.organizationJoinInvitationDo.DO = *(o.organizationJoinInvitationDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *organizationJoinInvitation) updateTableName(table string) *organizationJoinInvitation {
	o.ALL = field.NewAsterisk(table)
	o.CreatedAt = field.NewTime(table, "created_at")
	o.InviterUserID = field.NewString(table, "inviter_user_id")
	o.Status = field.NewString(table, "status")
	o.ID = field.NewString(table, "id")
	o.InviteeUserEmail = field.NewString(table, "invitee_user_email")
	o.OrganizationID = field.NewString(table, "organization_id")
	o.InviteeOrganizationRole = field.NewString(table, "invitee_organization_role")
	o.InviteeUserID = field.NewString(table, "invitee_user_id")

	o.fillFieldMap()

	return o
}

func (o *organizationJoinInvitation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *organizationJoinInvitation) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 8)
	o.fieldMap["created_at"] = o.CreatedAt
	o.fieldMap["inviter_user_id"] = o.InviterUserID
	o.fieldMap["status"] = o.Status
	o.fieldMap["id"] = o.ID
	o.fieldMap["invitee_user_email"] = o.InviteeUserEmail
	o.fieldMap["organization_id"] = o.OrganizationID
	o.fieldMap["invitee_organization_role"] = o.InviteeOrganizationRole
	o.fieldMap["invitee_user_id"] = o.InviteeUserID
}

func (o organizationJoinInvitation) clone(db *gorm.DB) organizationJoinInvitation {
	o.organizationJoinInvitationDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o organizationJoinInvitation) replaceDB(db *gorm.DB) organizationJoinInvitation {
	o.organizationJoinInvitationDo.ReplaceDB(db)
	return o
}

type organizationJoinInvitationDo struct{ gen.DO }

type IOrganizationJoinInvitationDo interface {
	gen.SubQuery
	Debug() IOrganizationJoinInvitationDo
	WithContext(ctx context.Context) IOrganizationJoinInvitationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOrganizationJoinInvitationDo
	WriteDB() IOrganizationJoinInvitationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOrganizationJoinInvitationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOrganizationJoinInvitationDo
	Not(conds ...gen.Condition) IOrganizationJoinInvitationDo
	Or(conds ...gen.Condition) IOrganizationJoinInvitationDo
	Select(conds ...field.Expr) IOrganizationJoinInvitationDo
	Where(conds ...gen.Condition) IOrganizationJoinInvitationDo
	Order(conds ...field.Expr) IOrganizationJoinInvitationDo
	Distinct(cols ...field.Expr) IOrganizationJoinInvitationDo
	Omit(cols ...field.Expr) IOrganizationJoinInvitationDo
	Join(table schema.Tabler, on ...field.Expr) IOrganizationJoinInvitationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOrganizationJoinInvitationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOrganizationJoinInvitationDo
	Group(cols ...field.Expr) IOrganizationJoinInvitationDo
	Having(conds ...gen.Condition) IOrganizationJoinInvitationDo
	Limit(limit int) IOrganizationJoinInvitationDo
	Offset(offset int) IOrganizationJoinInvitationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOrganizationJoinInvitationDo
	Unscoped() IOrganizationJoinInvitationDo
	Create(values ...*model.OrganizationJoinInvitation) error
	CreateInBatches(values []*model.OrganizationJoinInvitation, batchSize int) error
	Save(values ...*model.OrganizationJoinInvitation) error
	First() (*model.OrganizationJoinInvitation, error)
	Take() (*model.OrganizationJoinInvitation, error)
	Last() (*model.OrganizationJoinInvitation, error)
	Find() ([]*model.OrganizationJoinInvitation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OrganizationJoinInvitation, err error)
	FindInBatches(result *[]*model.OrganizationJoinInvitation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OrganizationJoinInvitation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOrganizationJoinInvitationDo
	Assign(attrs ...field.AssignExpr) IOrganizationJoinInvitationDo
	Joins(fields ...field.RelationField) IOrganizationJoinInvitationDo
	Preload(fields ...field.RelationField) IOrganizationJoinInvitationDo
	FirstOrInit() (*model.OrganizationJoinInvitation, error)
	FirstOrCreate() (*model.OrganizationJoinInvitation, error)
	FindByPage(offset int, limit int) (result []*model.OrganizationJoinInvitation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOrganizationJoinInvitationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o organizationJoinInvitationDo) Debug() IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.Debug())
}

func (o organizationJoinInvitationDo) WithContext(ctx context.Context) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o organizationJoinInvitationDo) ReadDB() IOrganizationJoinInvitationDo {
	return o.Clauses(dbresolver.Read)
}

func (o organizationJoinInvitationDo) WriteDB() IOrganizationJoinInvitationDo {
	return o.Clauses(dbresolver.Write)
}

func (o organizationJoinInvitationDo) Session(config *gorm.Session) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.Session(config))
}

func (o organizationJoinInvitationDo) Clauses(conds ...clause.Expression) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o organizationJoinInvitationDo) Returning(value interface{}, columns ...string) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o organizationJoinInvitationDo) Not(conds ...gen.Condition) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o organizationJoinInvitationDo) Or(conds ...gen.Condition) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o organizationJoinInvitationDo) Select(conds ...field.Expr) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o organizationJoinInvitationDo) Where(conds ...gen.Condition) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o organizationJoinInvitationDo) Order(conds ...field.Expr) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o organizationJoinInvitationDo) Distinct(cols ...field.Expr) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o organizationJoinInvitationDo) Omit(cols ...field.Expr) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o organizationJoinInvitationDo) Join(table schema.Tabler, on ...field.Expr) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o organizationJoinInvitationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o organizationJoinInvitationDo) RightJoin(table schema.Tabler, on ...field.Expr) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o organizationJoinInvitationDo) Group(cols ...field.Expr) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o organizationJoinInvitationDo) Having(conds ...gen.Condition) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o organizationJoinInvitationDo) Limit(limit int) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o organizationJoinInvitationDo) Offset(offset int) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o organizationJoinInvitationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o organizationJoinInvitationDo) Unscoped() IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.Unscoped())
}

func (o organizationJoinInvitationDo) Create(values ...*model.OrganizationJoinInvitation) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o organizationJoinInvitationDo) CreateInBatches(values []*model.OrganizationJoinInvitation, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o organizationJoinInvitationDo) Save(values ...*model.OrganizationJoinInvitation) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o organizationJoinInvitationDo) First() (*model.OrganizationJoinInvitation, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrganizationJoinInvitation), nil
	}
}

func (o organizationJoinInvitationDo) Take() (*model.OrganizationJoinInvitation, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrganizationJoinInvitation), nil
	}
}

func (o organizationJoinInvitationDo) Last() (*model.OrganizationJoinInvitation, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrganizationJoinInvitation), nil
	}
}

func (o organizationJoinInvitationDo) Find() ([]*model.OrganizationJoinInvitation, error) {
	result, err := o.DO.Find()
	return result.([]*model.OrganizationJoinInvitation), err
}

func (o organizationJoinInvitationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OrganizationJoinInvitation, err error) {
	buf := make([]*model.OrganizationJoinInvitation, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o organizationJoinInvitationDo) FindInBatches(result *[]*model.OrganizationJoinInvitation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o organizationJoinInvitationDo) Attrs(attrs ...field.AssignExpr) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o organizationJoinInvitationDo) Assign(attrs ...field.AssignExpr) IOrganizationJoinInvitationDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o organizationJoinInvitationDo) Joins(fields ...field.RelationField) IOrganizationJoinInvitationDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o organizationJoinInvitationDo) Preload(fields ...field.RelationField) IOrganizationJoinInvitationDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o organizationJoinInvitationDo) FirstOrInit() (*model.OrganizationJoinInvitation, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrganizationJoinInvitation), nil
	}
}

func (o organizationJoinInvitationDo) FirstOrCreate() (*model.OrganizationJoinInvitation, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrganizationJoinInvitation), nil
	}
}

func (o organizationJoinInvitationDo) FindByPage(offset int, limit int) (result []*model.OrganizationJoinInvitation, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o organizationJoinInvitationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o organizationJoinInvitationDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o organizationJoinInvitationDo) Delete(models ...*model.OrganizationJoinInvitation) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *organizationJoinInvitationDo) withDO(do gen.Dao) *organizationJoinInvitationDo {
	o.DO = *do.(*gen.DO)
	return o
}
