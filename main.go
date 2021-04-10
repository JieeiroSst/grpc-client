package main

import (
	"fmt"

	"github.com/JIeeiroSst/go-app/config"
	"github.com/JIeeiroSst/go-app/http"
	"github.com/JIeeiroSst/go-app/inventory"
	"github.com/JIeeiroSst/go-app/repositories/mongo"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func main(){
	fmt.Println(config.Config)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(config.Config.URL, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Ket noi that bai")
	}else {
		fmt.Println("ket noi thanh cong voi server")
	}
	defer conn.Close()
	client:=inventory.NewCheckInventoryServiceClient(conn)
	grpcHandler:=http.NewGrpcHandler(client)

	e := echo.New()

	dao:=mongo.NewMongoSqlRepo(config.Config.MongoConfig)

	handler:=http.Handler{}
	handler.SetRepo(dao)

	r:=e.Group("/api")
	r.GET("/item",handler.GetAll)
	r.GET("/item/:id",handler.GetById)
	r.POST("/item/update/:id",handler.UpdateItem)
	r.POST("/item/create",handler.CreateItem)
	r.POST("/item/delete/:id",handler.DeleteItem)

	e.POST("/order",grpcHandler.MakeOrder)

	e.Logger.Fatal(e.Start(":8080"))
}