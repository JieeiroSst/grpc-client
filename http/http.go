package http

import (
	"net/http"
	"strconv"

	"github.com/JIeeiroSst/go-app/entities"
	"github.com/JIeeiroSst/go-app/repositories"
	"github.com/labstack/echo/v4"
)

type Handler struct{
	repo repositories.ItemRepository
}

func (handler *Handler) SetRepo(repo repositories.ItemRepository) {
	handler.repo = repo
}

func (handler *Handler) GetAll(e echo.Context) error {
	items,err:=handler.repo.FindAll()
	if err!=nil {
		return e.JSON(http.StatusNoContent,err)	
	}
	return e.JSON(http.StatusOK,items)
}

func (handler *Handler) GetById(e echo.Context) error {
	item,err:=handler.repo.FindByID(e.Param("id"))
	if err!=nil {
		return e.JSON(http.StatusNoContent,err)
	}
	return e.JSON(http.StatusOK,item)
}

func (handler *Handler) CreateItem(e echo.Context) error {
	quantity,_:=strconv.Atoi( e.FormValue("quantity"))
	item:=entities.Item{
		ID: e.FormValue("id"),
		Quantity: quantity,
	}
	err:=handler.repo.CreateItem(item)
	if err!=nil {
		return e.JSON(http.StatusNoContent,"create failed")
	}
	return e.JSON(http.StatusOK,"create sucesss")
}

func (handler *Handler) DeleteItem(e echo.Context) error {
	id:=e.Param("id")
	err:=handler.repo.DeleteItem(id)
	if err!=nil {
		return e.JSON(http.StatusNoContent,"delete failed")
	}
	return e.JSON(http.StatusOK,"delete successs")
}

func (handler *Handler) UpdateItem(e echo.Context) error {
	quantity,_:=strconv.Atoi(e.FormValue("quantity"))
	id:= e.Param("id")
	item:=entities.Item{
		Quantity: quantity,
	}
	err:=handler.repo.UpdateItem(item,id)
	if err!=nil {
		return e.JSON(http.StatusNoContent,"update failed")
	}
	return e.JSON(http.StatusOK,"update sucesss")
}