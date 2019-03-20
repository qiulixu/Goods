package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["Goods/controllers/goods:GoodsController"] = append(beego.GlobalControllerRouter["Goods/controllers/goods:GoodsController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/goods`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Goods/controllers/goods:GoodsController"] = append(beego.GlobalControllerRouter["Goods/controllers/goods:GoodsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/goods`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Goods/controllers/goods:GoodsController"] = append(beego.GlobalControllerRouter["Goods/controllers/goods:GoodsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/goods`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Goods/controllers/goods:GoodsController"] = append(beego.GlobalControllerRouter["Goods/controllers/goods:GoodsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/goods`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
