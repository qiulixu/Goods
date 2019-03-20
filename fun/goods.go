package fun

import (
	"Goods/models"
	"errors"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"reflect"
)

//用户中心 公共函数

// @Title 获取magic值
// @Description 获取用户magic值
// @Param  b beego.Controller   true    "beego。Controller资源"
// @Return magic string   true    "用户的magic值"
func GetUserToken(b *beego.Controller)(token string,err error){
	token = b.Ctx.Request.Header.Get("token")
	if token == "" || len(token) != 35 {
		return "",errors.New("格式不符合规范")
	}
	return
}

// @Title 获取用户详情信息
// @Description 获取用户id
// @Param  magic string   true    "用户的magic值"
// @Return userId int     true    "用户的UserId"
func GetUserId(magic string)(userId int,err error){
	//获取用户id缓存
	key := "uc:login:"+magic;
	//获取用户信息
	userId,err = models.DoInt("GET",key)
	//缓存获取为空，并且错误类型为key不存在
	if err  != nil && err == redis.ErrNil {
		//重新缓存数据
		userId,err = 1,nil
		err = models.DoExpire("SET",key,models.ExpireTime,userId)
		if err != nil{
			return
		}
	}
	return
}

// @Title 获取用户Magic获取到用户id
// @Description 通过用的magic获取用户的id
// @Param  magic string   true    "用户的magic值"
// @Return userId int     true    "用户的UserId"
func GetMagicUserId(b *beego.Controller)(userId int){
	controllerCode := 100
	methodCode := 13
	//获取magic值
	token,err := GetUserToken(b)
	if err != nil{
		RJ(b,controllerCode,methodCode,10,nil)
	}
	//获取用户id
	userId,err = GetUserId(token)
	if err != nil{
		RJ(b,controllerCode,methodCode,11,nil)
	}
	//返回用户id
	return userId
}

// @Title 结构体转成map
// @Description 2019/3/8:11:03 AM Fuzz
// @Param   obj          formData  interface            true       "传入结构体"
// @Return  ref          query     map[string]interface true       "返回map"
// @Failure 400 错误提示
func StructToMap(obj interface{}) (ref map[string]interface{}){
	relType := reflect.TypeOf(obj).Elem()
	v := reflect.ValueOf(obj).Elem()
	ref = make(map[string]interface{})
	for i := 0; i < relType.NumField(); i++ {
		ref[relType.Field(i).Name] = v.Field(i).Interface()
	}
	return ref
}

// @Title 获取商品缓存
// @Description  获取商品缓存
// @Auth  2019/3/19:5:28 PM   Fuzz
// @Param      string     ""
// @Return  err error     "错误信息"
func GetCommodity(id int) {

}

