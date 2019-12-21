package api

import (
	"github.com/gin-gonic/gin"
	"github.com/justericgg/go-ddd-coffee-shop/infra/event"
	"github.com/justericgg/go-ddd-coffee-shop/order/infra/repository"
	"github.com/justericgg/go-ddd-coffee-shop/order/usecase"
	"log"
	"net/http"
)

type CreateReq struct {
	TableNo string `json:"table_no" binding:"required"`
	Items   []struct {
		ProductID string `json:"product_id" binding:"required"`
		Qty       int    `json:"qty" binding:"required"`
		Price     int64  `json:"price" binding:"required"`
	} `json:"items" binding:"required"`
}

func create(c *gin.Context) {
	var req CreateReq
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("bad request %+v", err)
		c.JSON(http.StatusBadRequest, Error{
			Code:    http.StatusBadRequest,
			Message: "bad request",
		})
		return
	}

	cmd := usecase.CreateOrderCmd{
		TableNo: req.TableNo,
		Items:   nil,
	}
	for _, item := range req.Items {
		cmd.Items = append(cmd.Items, usecase.Item{
			ProductID: item.ProductID,
			Qty:       item.Qty,
			Price:     item.Price,
		})
	}

	svc := usecase.NewCreateOrderSvc(repository.OrderRepository{}, event.Cwe{})
	err = svc.CreateOrder(cmd)
	if err != nil {
		log.Printf("internal server error %+v", err)
		c.JSON(http.StatusInternalServerError, Error{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, req)
}
