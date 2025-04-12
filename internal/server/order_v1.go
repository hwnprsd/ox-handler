package server

import (
	"strconv"
	"time"

	"github.com/hwnprsd/ox-handler/internal/models"
	"github.com/labstack/echo/v4"
)

type sigData struct {
	UserAddress string `json:"user_address"`
	Token       string `json:"token"`
	Amount      string `json:"amount"`
	OutputToken string `json:"output_token"`
}

func (s *Server) getSigDataV1Handler(ctx echo.Context) error {
	data := sigData{}
	if err := ctx.Bind(&data); err != nil {
		return ctx.String(400, "error binding body to request")
	}
	return ctx.JSON(200, data)
}

type sigDataWithSignature struct {
	UserAddress string `json:"user_address"`
	InputToken  string `json:"input_token"`
	Amount      string `json:"amount"`
	OutputToken string `json:"output_token"`
	Signature   string `json:"signature"`
}

func (s *Server) submitSignatureHandler(ctx echo.Context) error {
	data := sigDataWithSignature{}
	if err := ctx.Bind(&data); err != nil {
		return ctx.String(400, "error binding body to request")
	}
	// TODO: Add this to a offchain db queue. Notify fillers
	// for testing, create a new order and then update it after 5s
	orderV1 := models.NewOrderV1(data.UserAddress, data.InputToken, data.Amount, data.OutputToken)
	id, err := s.db.CreateOrderV1(orderV1)
	if err != nil {
		return ctx.String(400, err.Error())
	}
	orderV1.ID = id

	go func(id uint64) {
		time.Sleep(time.Second * 5)
		s.db.UpdateOrderV1Status(id, models.OrderStatusFilled)
	}(id)

	return ctx.JSON(200, orderV1)
}

func (s *Server) getOrderByIdHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.String(400, "invalid id")
	}
	order, err := s.db.GetOrderV1ById(uint64(id))
	if err != nil {
		return ctx.String(400, err.Error())
	}

	return ctx.JSON(200, &order)
}

func (s *Server) getAllOrdersByAddressHandler(ctx echo.Context) error {
	address := ctx.Param("address")
	if address == "" {
		return ctx.String(400, "invalid address")
	}
	orders, err := s.db.GetOrderV1ForUser(address)
	if err != nil {
		return ctx.String(400, err.Error())
	}
	return ctx.JSON(200, orders)
}
