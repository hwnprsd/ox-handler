package server

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/hwnprsd/ox-handler/internal/models"
	"github.com/labstack/echo/v4"
)

func (s *Server) getAllTradesByAddressHandler(c echo.Context) error {
	address := c.Param("address")
	if address == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "address parameter is required"})
	}

	// Get 10 mock trades for the specified address
	trades := getMockTrades(10, address, "")

	return c.JSON(http.StatusOK, trades)
}

func (s *Server) getAllTradesByTokenHandler(c echo.Context) error {
	token := c.Param("token")
	if token == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "token parameter is required"})
	}

	// Get 10 mock trades involving the specified token
	trades := getMockTrades(10, "", token)

	return c.JSON(http.StatusOK, trades)
}

// Helper function to get mock trades
func getMockTrades(count int, address string, token string) []models.TradeV1 {
	trades := make([]models.TradeV1, count)

	for i := range count {
		userAddress := address
		if userAddress == "" {
			userAddress = fmt.Sprintf("0x%x", rand.Intn(1000000))
		}

		assetA := fmt.Sprintf("0x%x", rand.Intn(1000000))
		assetB := fmt.Sprintf("0x%x", rand.Intn(1000000))

		// If token is specified, randomly assign it to either assetA or assetB
		if token != "" {
			if rand.Intn(2) == 0 {
				assetA = token
			} else {
				assetB = token
			}
		}

		trades[i] = models.TradeV1{
			Id:          fmt.Sprintf("trade_%d", i+1),
			UserAddress: userAddress,
			AssetA:      assetA,
			AssetB:      assetB,
			AmountA:     fmt.Sprintf("%d", rand.Intn(1000)+1),
			AmountB:     fmt.Sprintf("%d", rand.Intn(1000)+1),
			TxHash:      fmt.Sprintf("0x%x", rand.Intn(10000000)),
			Fee:         fmt.Sprintf("0.%d", rand.Intn(100)),
		}
	}

	return trades
}
