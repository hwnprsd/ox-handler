package database

import "github.com/hwnprsd/ox-handler/internal/models"

func (s *service) migrate() {
	s.db.AutoMigrate(&models.OrderV1{})
}
