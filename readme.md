package main

import (
	"fmt"
	cProto "goproj/cart/proto"
	Cmongo "goproj/cart/repository/mongo"
	Csql "goproj/cart/repository/mysql"

	Chttp "goproj/cart/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)


func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":1889", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Ket noi that bai")
	}
	defer conn.Close()
	client := cProto.NewCheckAndOrderServiceClient(conn)
	grpcHandler := Chttp.CreateGrpcHandler(client)

	e := echo.New()
	// dao := &Cmongo.MongoDb{}
	// dao.InitDb(&config.MongoConfig)
	Csql.(&config.MysqlConfig)
	dao := &Csql.CartDao{}
	dao.InitDao(&config.MysqlConfig)
	handler := &Chttp.Handler{}
	handler.SetRepo(dao)
	e.GET("/getCart", handler.FindCartByUser)
	e.DELETE("/delete/cart", handler.RemoveCart)
	e.PUT("/update/cart", handler.UpdateCart)
	e.DELETE("/delete/cart-item", handler.RemoveCartItem)
	e.PUT("/update/cart-item", handler.UpdateCartItem)
	e.POST("/add/cart-item", handler.AddCartItem)
	e.POST("/add/cart", handler.AddCart)

	//

	e.POST("/make-order", grpcHandler.MakeOrder)
	//
	e.Logger.Fatal(e.Start(":1323"))
}
