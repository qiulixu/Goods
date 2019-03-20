// @APIVersion 1.0.0
// @Title 爱数商城
// @Description  开发一个简单的demo
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
// beego 路由信息
// 徐林峰
// 2019年03月18日19:02:56
// 更新记录 人员  格式：创建时间： yyyyMMdd
package routers

import (
	"Goods/controllers/goods"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/goods",
		// 商品模块
		beego.NSNamespace("/goods",
			//商品列表
			beego.NSInclude(
				//商品模块
				 &goods.GoodsController{Code:101},
				),
		),
	)
	beego.AddNamespace(ns)
	beego.SetStaticPath("/swagger", "swagger")
}
