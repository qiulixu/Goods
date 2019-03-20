package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"strings"
	"time"
  "gopkg.in/mgo.v2/bson"
)
//@Title Mongo连接数据库
//@Description Mongo连接数据库
//@Author Mick
//@Time 19-02-18
//@参考文档  https://www.jb51.net/article/143208.htm
var (
	mongodbHost  = strings.Split(beego.AppConfig.String("mongodbHostPort"),",")
    mongodbUser  = beego.AppConfig.String("mongodbUser")
	mongodbPassword  = beego.AppConfig.String("mongodbPassword")
	mongodbAuthDb = beego.AppConfig.String("mongodbAuthDb")
	mongodbDb = beego.AppConfig.String("mongodbDb")
	mongodbReplSet = beego.AppConfig.String("mongodbReplSet")
)
var globalS *mgo.Session
func init() {
	dialInfo := &mgo.DialInfo{
		Addrs: mongodbHost,
		Timeout: 10*time.Second,
		Source: mongodbAuthDb,
		Username: mongodbUser,
		Password: mongodbPassword,
		PoolLimit: 512,
		ReplicaSetName:mongodbReplSet,
		Database:mongodbDb,
	}
	var err error
	globalS, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	globalS.SetMode(mgo.Monotonic, true)
	// 第二种连接方式
    //url := "mongodb://audioMark:vgurvLXQEdcEdCcw7mbp53xL2VJjwSdX@47.92.107.58:27717,47.92.255.215:27717,39.98.78.28:27717/audioMark?replicaSet=MagicGo"
	//globalS,err = mgo.Dial(url)
	//if err != nil{
	//	fmt.Println("Mongo连接异常",err)
	//}
}
func Connect(collection string) (*mgo.Session, *mgo.Collection) {
	ms := globalS.Copy()
	c := ms.DB(mongodbDb).C(collection)
	return ms, c
}
// @Title 插入数据
// @Description  数据插入操作
// @Param	collection   string        true    "集合的名称"
// @Param   doc          interface{}   true    "文档数据"
// @Return  error        error         true    "错误状态"
func MgoInsert(collection string, doc interface{}) error {
	ms, c := Connect(collection)
	defer ms.Close()
	return c.Insert(doc)
}
// @Title 获取单条数据
// @Description  根据指定条件获取符合标准的一条数据
// @Param	collection   string        true    "集合的名称"
// @Param   query        interface{}   true    "查询条件"
// @Param   selector     interface{}   true    "查询字段过滤器，不过略直接写nil"
// @Param   result       interface{}   true    "结果保存变量，传引用类型"
// @Return  error        error         true    "错误状态"
func MgoFindOne(collection string, query bson.M, selector, result interface{}) error {
	ms, c := Connect(collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}
// @Title 获取所有数据
// @Description  根据指定条件获取符合标准的所有数据
// @Param	collection   string        true    "集合的名称"
// @Param   query        interface{}   true    "查询条件"
// @Param   selector     interface{}   true    "查询字段过滤器，不过略直接写nil"
// @Param   result       interface{}   true    "结果保存变量，传切片类型"
// @Return  error        error         true    "错误状态"
func MgoFindAll( collection string, query, selector, result interface{}) error {
	ms, c := Connect( collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}
// @Title 更新一条数据(不存在则不更新)
// @Description  根据指定条件更新符合标准的一条数据
// @Param	collection   string        true    "集合的名称"
// @Param   selector     interface{}   true    "更新条件"
// @Param   update       interface{}   true    "更新数据"
// @Return  error        error         true    "错误状态"
func MgoUpdate(collection string, query, update interface{}) error {
	ms, c := Connect(collection)
	defer ms.Close()
	return c.Update(query, update)
}
// @Title 更新一条数据(不存在则插入)
// @Description  根据指定条件更新符合标准的一条数据(不存在则插入)
// @Param	collection   string        true    "集合的名称"
// @Param   selector     interface{}   true    "更新条件"
// @Param   update       interface{}   true    "更新数据"
// @Return  error        error         true    "错误状态"
func MgoUpsert(collection string, selector, update interface{}) error {
	ms, c := Connect( collection)
	defer ms.Close()
	_, err := c.Upsert(selector, update)
	return err
}
// @Title 更新所有数据
// @Description  根据指定条件更新符合标准的所有数据
// @Param	collection   string        true    "集合的名称"
// @Param   selector     interface{}   true    "更新条件"
// @Param   update       interface{}   true    "更新数据"
// @Return  error        error         true    "错误状态"
func MgoUpdateAll(collection string, selector, update interface{}) error {
	ms, c := Connect( collection)
	defer ms.Close()
	_, err := c.UpdateAll(selector, update)
	return err
}
