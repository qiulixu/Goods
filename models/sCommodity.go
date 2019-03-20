// @Title 
// @Description 
// @Author  Fuzz (2019/3/20:7:00 PM)
// @Update  Fuzz (2019/3/20:7:00 PM)
package models

//s_commodity 表
type SCommodity struct {
	Mid			    string	  `orm:"pk" description:"商品id"`
	Name		    string	  `description:"商品名称"`
	PersonId 	  	int	      `description:"商品添加人id"`
	Price		    int 	  `description:"商品价格"`
	Stock		    int 	  `description:"总计商品库存"`
	UsedStock	  	int 	  `description:"剩余商品库存"`
	Time		    int64	  `description:"商品添加时间"`
	Stat        	int       `json:"stat" description:"商品状态 1：未删除 2：删除"`
}

// 添加商品
func SCommodityAdd(mid string,name string,user_id int,price int,stock int,time int64)(int64,error){
	SC := SCommodity{
		Mid:mid,
		Name:name,
		PersonId:user_id,
		Price:price,
		Stock:stock,
		UsedStock:stock,
		Time:time,
		Stat:1,
	}
	return Insert(&SC)
}