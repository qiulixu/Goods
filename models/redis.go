package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"time"
)
var (
	redisURL string
	redisPassword  string
	redisMaxIdle int
	redisIdleTimeoutSecond time.Duration
	redisPrefix string
	ExpireTime int64
)
func init(){
	// 初始化连接方式
	redisURL  =  "redis://" + beego.AppConfig.String("redisHost")+":"+beego.AppConfig.String("redisPort")
	// 初始化密码
	redisPassword   = beego.AppConfig.String("redisPassword")
	// 初始最大空闲时间
	redisMaxIdle,_ = beego.AppConfig.Int("redisMaxidle")
	// 初始空闲时间
	redisIdleTimeoutSecond,_ = time.ParseDuration(beego.AppConfig.String("redisIdleTimeoutSecond"))
	// 初始化字段前缀
	redisPrefix  = beego.AppConfig.String("redisPrefix")
	ExpireTime = time.Now().Unix()+86400
}
// @Title 连接池
// @Description  生成连接池，并发访问会使用连接池中的资源句柄
// @Return *redis.Pool    struct   true    "返回连接池"
func NewRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     redisMaxIdle,
		IdleTimeout: redisIdleTimeoutSecond * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(redisURL)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			//验证redis密码
			if _, authErr := c.Do("AUTH", redisPassword); authErr != nil {
				return nil, fmt.Errorf("redis auth password error: %s", authErr)
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
}
// @Title 操作入口
// @Description  像命令行一样的执行命令
// @Return  interface    interface{}   true    "执行命令返回的状态"
// @Return  error        error         true    "错误状态"
func Do(comm string,key string,args ...interface{})(interface{},error){
	c := NewRedisPool().Get()
	defer c.Close()
	key = redisPrefix + ":"+key
	param := []interface{}{key}
	param = append(param,args...)
	return  c.Do(comm,param...)
}

// @Title 获取单个数据
// @Description  获取的数据都是字节，需要转换成字符串
// @Return  bool    bool   true    "执行命令返回的状态 true 成功 false 失败"
// @Return  error   error  true    "错误状态"
func DoBool(comm string,key string,args ...interface{})(bool,error){
	return  redis.Bool(Do(comm ,key ,args ...))
}

// @Title 获取单个数据
// @Description  获取的数据都是字节，需要转换成字符串
// @Return  interface    interface{}   true    "执行命令返回的状态"
// @Return  error        error         true    "错误状态"
func DoString(comm string,key string,args ...interface{})(interface{},error){
	return  redis.String(Do(comm ,key ,args ...))
}

// @Title 获取批量数据
// @Description  获取的数据都是字节，需要转换成字符串
// @Return  interface    interface{}   true    "执行命令返回的状态"
// @Return  error        error         true    "错误状态"
func DoStrings(comm string,key string,args ...interface{})(interface{},error){
	return  redis.Strings(Do(comm ,key ,args ...))
}

func DoInt(comm string,key string,args ...interface{})(int,error){
	return  redis.Int(Do(comm ,key ,args ...))
}

func DoValues(comm string,key string,args ...interface{})(desc []interface{},err error){
	return redis.Values(Do(comm ,key ,args ...))
}

func DoStruct(comm string,key string,data interface{},args ...interface{})(err error){
	tmp,err := redis.Values(Do(comm ,key ,args ...))
	if err != nil{
		return
	}
	err = redis.ScanStruct(tmp,data)
	if err != nil{
		return
	}
	return
}

func DoExpire(comm string,key string,ExpireTime int64,args ...interface{})(err error){
	_,err = Do(comm ,key ,args ...)
	if err != nil{
		return
	}
	_,err = Do("EXPIREAT" ,key ,ExpireTime)
	return
}

//传入结构体，把结构体通过json转换成 map方便redis入库操作
func StructToSlice(obj interface{})(data []interface{},err error){
	j,err := json.Marshal(obj)
	if err != nil{
		return
	}
	tmp := make(map[string]interface{})
	d := json.NewDecoder(bytes.NewReader(j))
	d.UseNumber()
	err = d.Decode(&tmp)
	if err != nil{
		return
	}
	for k,v := range tmp{
		data = append(data,k,v)
	}
	return
}