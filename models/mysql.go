package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //导出sql库
)
var(
	mysqlHost = beego.AppConfig.String("mysqlHost")
	mysqlPort = beego.AppConfig.String("mysqlPort")
	mysqlUser = beego.AppConfig.String("mysqlUser")
	mysqlPassword = beego.AppConfig.String("mysqlPassword")
	mysqlDb =beego.AppConfig.String("mysqlDb")
)
func init() {
	connectInfo := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDb + "?charset=utf8"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql",connectInfo )
	orm.SetMaxIdleConns("default", 0)
	orm.SetMaxOpenConns("default", 512)
	orm.RegisterModel(new(SCommodity))
}

// 查询单挑记录
func Select(sql string,data interface{},where ...interface{})(error){
  o := orm.NewOrm()
  err := o.Raw(sql,where).QueryRow(data)
  return err
}

// 查询数据库返回多条几率
func SelectAll(sql string,data interface{},where ...interface{})(int64,error){
  o := orm.NewOrm()
  num,err := o.Raw(sql,where).QueryRows(data)
  return num,err
}

// 添加
func Insert(data interface{})(int64,error){
  o := orm.NewOrm()
  id, err := o.Insert(data)
  return id,err
}

// 修改
func Update(data interface{},cols ...string)(int64,error){
  o := orm.NewOrm()
  id, err := o.Update(data,cols...)
  return id,err
}

// 更新操作
func ExecSql(sql string) (int64,error){
  o := orm.NewOrm()
  res,err := o.Raw(sql).Exec()
  row,err := res.RowsAffected()
  return row,err
}
