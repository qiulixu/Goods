package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["Goods/controllers/shopCart:ShopCartController"] = append(beego.GlobalControllerRouter["Goods/controllers/shopCart:ShopCartController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/ShopCart`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Goods/controllers/shopCart:ShopCartController"] = append(beego.GlobalControllerRouter["Goods/controllers/shopCart:ShopCartController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/ShopCart`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Goods/controllers/shopCart:ShopCartController"] = append(beego.GlobalControllerRouter["Goods/controllers/shopCart:ShopCartController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/ShopCart`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Goods/controllers/shopCart:ShopCartController"] = append(beego.GlobalControllerRouter["Goods/controllers/shopCart:ShopCartController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/ShopCart`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
