package http

import (
	"net/http"
	"encoding/json"
	"sync"
	"context"
	"time"
	"github.com/JIeeiroSst/go-app/inventory"
	"github.com/JIeeiroSst/go-app/entities"
	"github.com/labstack/echo/v4"
)

type GrpcHandler struct{
	client inventory.CheckInventoryServiceClient
}

func (handler *GrpcHandler) MakeOrder(c echo.Context) error {
	body := c.Request().Body
	var cart entities.Item
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&cart)
	if err != nil {
		return c.JSON(http.StatusOK, "Fail")
	}
	var items []entities.Item
	itemList := make([]*inventory.Item, 0)
	for _, cartItem := range items {
		item := inventory.Item{
			ProductId:    cartItem.ID,
			Quantity: int32(cartItem.Quantity),
		}
		itemList = append(itemList, &item)
	}
	request := inventory.CheckInventory{
		Items: itemList,
	}
	wait := &sync.WaitGroup{}
	wait.Add(5)
	for i := 0; i < 5; i++ {
		handler.sendRequest(&request, wait)
	}
	
	wait.Wait()
	return c.String(http.StatusBadRequest,"===========")
}

func (handler *GrpcHandler) sendRequest(request *inventory.CheckInventory, wait *sync.WaitGroup) error {
	defer wait.Done()
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	_,err:=handler.client.Check(ctx,request)
	return err
}


func NewGrpcHandler(client inventory.CheckInventoryServiceClient)GrpcHandler{
	return GrpcHandler{
		client:client,
	}
}