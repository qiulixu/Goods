package fun

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)
var(
	platformCode = beego.AppConfig.String("platformCode")
	defaultControllerCode = 100
)

// @Title 结构体返回json数据
// @Description 继承了beego.Controller的结构体实例返回json数据
// @Param  obj           interface{}   true    "组合了beego.Controller的结构体实例"
// @Param  statusCode    int           true    "状态码"
// @Param  msg           string        true    "状态码对应的名词解释"
// @Param  data          interface{}   true    "要返回的数据"
func Rj(obj interface{},statusCode int,data interface{}){
	// 反射类性值，指针类型需要执行elem方法
	refV := reflect.ValueOf(obj).Elem()
	controllerRef := refV.FieldByName("Controller")
	// 控制器编号
	controllerCode := refV.FieldByName("Code").Int()
	// 方法名称
	actionName := controllerRef.FieldByName("actionName").String()
	// 获取调用的方法
	registMethodCode := refV.Addr().MethodByName("RegistMethodCode")
	if !registMethodCode.IsValid(){
		panic("RegistMethodCode is UnValid")
	}
	// 执行获取方法编号的方法
	methodCodesName :=  registMethodCode.Call(nil)
	methodCodes := methodCodesName[0].Interface().(map[string]int)
	// 获取方法对应的状态码
	methodCode,ok := methodCodes[actionName]
	if !ok {
		panic(fmt.Sprintf("%s 方法对应的编码不存在",actionName))
	}
	// 类型断言，转换成
	controller := controllerRef.Interface().(beego.Controller)
	RJ(&controller,int(controllerCode),methodCode,statusCode,data)
}

// @Title 基类返回json数据
// @Description  基类不需要反射直接返回json数据
// @Param  controller    interface{}   true    "beego.Controller的结构体实例"
// @Param  statusCode    int           true    "状态码"
// @Param  msg           string        true    "状态码对应的名词解释"
// @Param  data          interface{}   true    "要返回的数据"
func RJ(controller *beego.Controller,controllerCode int,methodCode int,statusCode int,data interface{}){
	var code  int = 200
	// 检测装填码，判断是否格式化
	if statusCode != 200 {
		code,_ = strconv.Atoi(fmt.Sprintf("%s%d%d%d",platformCode,controllerCode,methodCode,statusCode))
	}
	iniConf,err := config.NewConfig("ini", "lang/zh.ini")
	//获取语言结果
	if err != nil{
		panic(err)
	}
	lang := strings.Split(iniConf.String(strconv.Itoa(code)),"|#|")
	msg:= lang[0]
	if "en" == controller.Ctx.Request.Header.Get("lang"){
		msg = lang[1]
	}
	// 格式化返回的数据
	controller.Data["json"] = struct{
		Code int
		Msg string
		Data interface{}
	}{
		Code:code,
		Msg:msg,
		Data:data,
	}
	// 记录日志
	logs.Debug("%v",controller.Data["json"])
	// 返回数据并退出
	controller.ServeJSON()
	controller.StopRun()
}

// @Title 表单数据之字符串长度验证
// @Description 验证字符串的长度是否符合格式(utf8格式验证)
// @Param  obj           beego.controller   true    "beego.Controller的结构体实例"
// @Param  filedName     string             true    "要获取的表单字段名称"
// @Param  between       [2]int             true    "字符串长度区间"
// @Param  str           string             true    "符合验证的字符串"
func LenStr(obj *beego.Controller,filedName string,between [2]int)(str string){
	methodCode := 11
	str = obj.GetString(filedName)
	length := utf8.RuneCountInString(str)
	// 判断是否限制最大长度
	if  between[1] != 0{
		// 判断最大值是否大于等于最小值
		if between[0] > between[1] {
			return ""
		}
	}
	// 检测区间真正的数值
	if (between[0] !=0 && length < between[0]) || (between[1] != 0 && length > between[1]) {
		RJ(obj,defaultControllerCode,methodCode,10,nil)
	}
	return
}

// @Title 表单数据之数字长度验证
// @Description 验证数字是否再次范围之内
// @Param  obj           beego.controller   true    "beego.Controller的结构体实例"
// @Param  filedName     string    true    "要获取的表单字段名称"
// @Param  between       [2]int    true    "数字长度区间"
// @Return num           int       true    "符合验证的数值"
func LenInt(obj *beego.Controller,filedName string,between [2]int)(num int){
	methodCode := 12
	num,err := obj.GetInt(filedName)
	if err != nil{
		RJ(obj,defaultControllerCode,methodCode,10,nil)
	}
	// 判断是否限制最大区间
	if  between[1] != 0{
		// 判断最大值是否大于等于最小值
		if between[0] > between[1] {
			return 0
		}
	}
	if (between[0] != 0 && num < between[0]) || (between[1] != 0 && num > between[1]){
		RJ(obj,defaultControllerCode,methodCode,11,nil)
	}
	return
}

// @Title 生成随机数
// @Description 生成指定长度的数字
// @Param  length    int   true    "要生成数字的长度"
// @Return num       int   true    "生成指定长度后的数字"
func RandNum(length int)(num int){
	if length <= 0{
		return 0
	}
	var code string
	ran := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<length;i++{
		tmp  :=  ran.Intn(9)
		if i==0 && tmp == 0{
			tmp = 1
		}
		code += strconv.Itoa(tmp)
	}
	num,_ = strconv.Atoi(code)
	return
}
// @Title 随机字符
// @Description 生成指定长度的字符串数字大小写字母
// @Param  length	string   true    "要生成的字符长度"
// @Return str      string   true    "生成指定长度后的字符串"
func RandomString(length int)(str string) {
	baseStr := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(baseStr)
	result := []byte{}
	ran := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[ran.Intn(len(bytes))])
	}
	return string(result)
}
// @Title 加密字符串
// @Description 使用md5对原始字符+盐进行2次字符串加密
// @Param  str    string   true    "要加密的字符串"
// @Param  salt   string   true    "盐"
// @Return encStr string   true    "加密后的字符串"
func EncryptPassword(str string,salt string)(encStr string){
	enc1 := md5.Sum([]byte(str + salt))
	enc2 := enc1[:]
	has := md5.Sum(enc2)
	return fmt.Sprintf("%x", has)
}
