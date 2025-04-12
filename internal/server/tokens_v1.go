package server

import "github.com/labstack/echo/v4"

type token struct {
	Name     string
	FullName string
	Address  string
	Decimals uint32
}

type tokenPair struct {
	TokenA token
	TokenB token
	Name   string
}

func (s *Server) getAllTradeTokensHandler(ctx echo.Context) error {
	wS := "0x5A91D3042b71A92f6757Fa937763D03Cc65ED8BC"
	RiftX := "0x2b258Ba5F33A72DcEf2E70A11B75Dd0eF0730942"
	pair1 := tokenPair{
		TokenA: token{FullName: "Rift (ODX)", Name: "Rift.X", Address: RiftX, Decimals: 18},
		TokenB: token{FullName: "Wrapped Sonic", Name: "wS", Address: wS, Decimals: 18},
		Name:   "Rift.X/wS",
	}
	pairs := []tokenPair{pair1}
	return ctx.JSON(200, pairs)
}

func (s *Server) getAllOrderTokensHandler(ctx echo.Context) error {
	xUSDT := "0x2d4b1eDa9514675a9F8CB13b3f3a7475ebb81024"
	RiftX := "0x2b258Ba5F33A72DcEf2E70A11B75Dd0eF0730942"
	pair1 := tokenPair{
		TokenA: token{FullName: "Rift (ODX)", Name: "Rift.X", Address: RiftX, Decimals: 18},
		TokenB: token{FullName: "USDT (ODX)", Name: "xUSDT", Address: xUSDT, Decimals: 6},
		Name:   "Rift.X/wS",
	}
	pairs := []tokenPair{pair1}
	return ctx.JSON(200, pairs)
}
