package services

import (
	"awesomeProject/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type ConfigService struct {
	db *gorm.DB
}

func NewConfigService(db *gorm.DB) *ConfigService {
	return &ConfigService{
		db: db,
	}
}

func (s *ConfigService) GetConfigByMerchantID(merchantID string) (*models.MerchantDashboardConfig, error) {
	var config models.MerchantDashboardConfig
	err := s.db.Where("merchant_id = ?", uuid.MustParse(merchantID)).Order("updated_at DESC").First(&config).Error
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (s *ConfigService) CreateConfig(config *models.MerchantDashboardConfig) error {
	config.ID = uuid.New()
	config.CreatedAt = time.Now()
	config.UpdatedAt = time.Now()

	err := s.db.Create(config).Error
	return err
}
