package server

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"fmt"
	"log"
	"time"

	"github.com/hwnprsd/ox-handler/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/coder/websocket"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*", "wss://*", "ws://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	e.GET("/", s.HelloWorldHandler)

	e.GET("/health", s.healthHandler)

	e.GET("/ws", s.websocketHandler)

	orderV1 := e.Group("/order/v1")
	orderV1.POST("/sig-data", s.getSigDataV1Handler)
	orderV1.POST("/signature", s.submitSignatureHandler)
	orderV1.GET("/id/:id", s.getOrderByIdHandler)
	orderV1.GET("/address/:address", s.getAllOrdersByAddressHandler)
	orderV1.GET("/tokens", s.getAllOrderTokensHandler)

	tradeV1 := e.Group("/trade/v1")
	tradeV1.GET("/address/:address", s.getAllTradesByAddressHandler)
	tradeV1.GET("/token/:address", s.getAllTradeTokensHandler)
	tradeV1.GET("/tokens", s.getAllTradeTokensHandler)

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) websocketHandler(c echo.Context) error {
	w := c.Response().Writer
	r := c.Request()
	options := &websocket.AcceptOptions{InsecureSkipVerify: true}
	socket, err := websocket.Accept(w, r, options)
	if err != nil {
		log.Printf("could not open websocket: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}
	defer socket.Close(websocket.StatusGoingAway, "server closing websocket")

	ctx := r.Context()

	// Track subscriptions
	subscriptions := make(map[string]string)

	// Create a channel to signal termination
	done := make(chan struct{})

	// Start a goroutine to read messages from the client
	go func() {
		defer close(done)

		for {
			messageType, message, err := socket.Read(ctx)
			if err != nil {
				log.Printf("error reading from websocket: %v", err)
				return
			}

			if messageType == websocket.MessageText {
				// Parse subscription message
				var subMsg struct {
					Type    string `json:"type"`
					Topic   string `json:"topic"`
					Address string `json:"address,omitempty"`
					Token   string `json:"token,omitempty"`
				}

				if err := json.Unmarshal(message, &subMsg); err != nil {
					log.Printf("error parsing subscription message: %v", err)
					errMsg := map[string]string{"error": "invalid subscription format"}
					errJSON, _ := json.Marshal(errMsg)
					socket.Write(ctx, websocket.MessageText, errJSON)
					continue
				}

				if subMsg.Type == "subscribe" {
					// Add subscription
					if subMsg.Topic == "address" && subMsg.Address != "" {
						subscriptions["address"] = subMsg.Address
						log.Printf("Client subscribed to address: %s", subMsg.Address)
					} else if subMsg.Topic == "token" && subMsg.Token != "" {
						subscriptions["token"] = subMsg.Token
						log.Printf("Client subscribed to token: %s", subMsg.Token)
					} else {
						errMsg := map[string]string{"error": "invalid subscription parameters"}
						errJSON, _ := json.Marshal(errMsg)
						socket.Write(ctx, websocket.MessageText, errJSON)
						continue
					}

					// Send confirmation
					confirm := map[string]string{
						"type":    "subscribed",
						"topic":   subMsg.Topic,
						"message": fmt.Sprintf("Successfully subscribed to %s", subMsg.Topic),
					}
					confirmJSON, _ := json.Marshal(confirm)
					socket.Write(ctx, websocket.MessageText, confirmJSON)
				} else if subMsg.Type == "unsubscribe" {
					// Remove subscription
					if subMsg.Topic == "address" {
						delete(subscriptions, "address")
						log.Printf("Client unsubscribed from address")
					} else if subMsg.Topic == "token" {
						delete(subscriptions, "token")
						log.Printf("Client unsubscribed from token")
					}

					// Send confirmation
					confirm := map[string]string{
						"type":    "unsubscribed",
						"topic":   subMsg.Topic,
						"message": fmt.Sprintf("Successfully unsubscribed from %s", subMsg.Topic),
					}
					confirmJSON, _ := json.Marshal(confirm)
					socket.Write(ctx, websocket.MessageText, confirmJSON)
				}
			}
		}
	}()

	// Send random trade data if there are active subscriptions
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return nil
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			// Only send data if there are active subscriptions
			if len(subscriptions) > 0 {
				trade := generateRandomTrade(subscriptions)
				message := map[string]interface{}{
					"type": "trade",
					"data": trade,
				}
				messageJSON, err := json.Marshal(message)
				if err != nil {
					log.Printf("Error marshaling trade data: %v", err)
					continue
				}

				if err := socket.Write(ctx, websocket.MessageText, messageJSON); err != nil {
					log.Printf("Error writing to websocket: %v", err)
					return nil
				}
			}
		}
	}
}

// Helper function to generate random trade data based on subscription topics
func generateRandomTrade(topics map[string]string) models.TradeV1 {
	// Generate a random ID
	id := fmt.Sprintf("trade_%d", time.Now().UnixNano())

	// Default values
	userAddress := fmt.Sprintf("0x%x", rand.Intn(1000000))
	assetA := fmt.Sprintf("0x%x", rand.Intn(1000000))
	assetB := fmt.Sprintf("0x%x", rand.Intn(1000000))

	// Use subscription topics if available
	if address, ok := topics["address"]; ok {
		userAddress = address
	}
	if token, ok := topics["token"]; ok {
		// Randomly assign the token to either assetA or assetB
		if rand.Intn(2) == 0 {
			assetA = token
		} else {
			assetB = token
		}
	}

	return models.TradeV1{
		Id:          id,
		UserAddress: userAddress,
		AssetA:      assetA,
		AssetB:      assetB,
		AmountA:     fmt.Sprintf("%d", rand.Intn(1000)+1),
		AmountB:     fmt.Sprintf("%d", rand.Intn(1000)+1),
		TxHash:      fmt.Sprintf("0x%x", rand.Intn(10000000)),
		Fee:         fmt.Sprintf("0.%d", rand.Intn(100)),
	}
}
