// @Title 商品模块
// @Description 商品列表、增加商品、编辑商品、删除商品
// @Author  Fuzz (2019/3/20:6:25 PM)
// @Update  Fuzz (2019/3/20:6:25 PM)
package goods

import (
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"Goods/fun"
	"Goods/models"
	"time"
)

//商品模块
type GoodsController struct{
	Code int
	beego.Controller
}

func(g *GoodsController) RegistMethodCode()(map[string]int){
	return map[string]int{
		"Get":10,
		"Post":11,
	}
}

// @Title 商品列表
// @Description 2019/3/20:6:33 PM Fuzz 获取商品列表
// @Success 200
// @Failure 400 错误提示
// @router /goods [get]
func(g *GoodsController) Get()(){
	fmt.Println("Get")
}

// @Title 添加商品
// @Description 2019/3/20:6:33 PM Fuzz 添加商品
// @Param   token   	header     string	true   "用户标识符"
// @Param   goodsName   formData   string	true   "商品名称"
// @Param   price       formData   int		true   "商品价格"
// @Param   stock       formData   int		true   "商品库存"
// @Success 200
// @Failure 10 数据添加失败，请稍后重试
// @router /goods [post]
func(g *GoodsController) Post()(){
	//mongodbID
	mid := bson.NewObjectId().Hex()
	personId := fun.GetMagicUserId(&g.Controller)
	// 获取商品名称
	goodsName := fun.LenStr(&g.Controller,"goodsName",[2]int{2,18})
	// 获取价格
	price := fun.LenInt(&g.Controller,"price",[2]int{0,9999})
	// 库存
	stock := fun.LenInt(&g.Controller,"stock",[2]int{0,9999})
	addTime := time.Now().Unix()
	//fmt.Println(mid,personId,goodsName,price,stock)
	inserId,err := models.SCommodityAdd(mid,goodsName,personId,price,stock,addTime)
	if err != nil {
		fun.Rj(g,10,err)
	}
	fun.Rj(g,200,inserId)

}

// @Title 编辑商品
// @Description 2019/3/20:6:33 PM Fuzz 编辑商品
// @Success 200
// @Failure 400 错误提示
// @router /goods [put]
func(g *GoodsController) Put()(){
	fmt.Println("Put")
}


// @Title 删除商品
// @Description 2019/3/20:6:33 PM Fuzz	删除商品
// @Success 200
// @Failure 400 错误提示
// @router /goods [delete]
func(g *GoodsController) Delete()(){
	fmt.Println("Delete")
}
