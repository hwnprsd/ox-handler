package database

import (
	"fmt"

	"github.com/hwnprsd/ox-handler/internal/models"
	"gorm.io/gorm/clause"
)

func (s *service) CreateOrderV1(order *models.OrderV1) (uint64, error) {
	if res := s.db.Clauses(clause.Returning{}).Create(order); res.Error != nil {
		return 0, fmt.Errorf("error creating order: %w", res.Error)
	}
	return order.ID, nil
}

func (s *service) GetOrderV1ById(id uint64) (models.OrderV1, error) {
	order := models.OrderV1{}
	if res := s.db.Where("id = ?", id).First(&order); res.Error != nil {
		return models.OrderV1{}, fmt.Errorf("error fetching order with id=%d; err=%w", id, res.Error)
	}
	return order, nil
}

func (s *service) GetOrderV1ForUser(address string) ([]models.OrderV1, error) {
	orders := make([]models.OrderV1, 0)
	if res := s.db.Where("user_address = ?", address).Find(&orders); res.Error != nil {
		return nil, fmt.Errorf("error fetching orders for user=%s; err=%w", address, res.Error)
	}
	return orders, nil
}

func (s *service) UpdateOrderV1Status(orderId uint64, status models.OrderStatus) error {
	if res := s.db.Model(&models.OrderV1{}).Where("id = ?", orderId).Update("status", uint32(status)); res.Error != nil {
		return fmt.Errorf("error updating order status; err=%w", res.Error)
	}
	return nil
}
