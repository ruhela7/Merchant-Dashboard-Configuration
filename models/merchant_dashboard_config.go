package models

import (
	"time"

	"github.com/google/uuid"
)

type MerchantDashboardConfigurations struct {
	Analytics struct {
		Enabled  bool `json:"enabled"`
		Sections struct {
			Order      bool `json:"order"`
			Growth     bool `json:"growth"`
			Others     bool `json:"others"`
			Payment    bool `json:"payment"`
			Business   bool `json:"business"`
			Conversion bool `json:"conversion"`
			Settlement bool `json:"settlement"`
		} `json:"sections"`
	} `json:"analytics"`

	OrderManager struct {
		Sections struct {
			EditOrders  bool `json:"edit_orders"`
			BulkRefunds bool `json:"bulk_refunds"`
		} `json:"sections"`
	} `json:"order_manager"`

	ConfigManager struct {
		Enabled  bool `json:"enabled"`
		Sections struct {
			CODConfiguration      bool `json:"cod_configuration"`
			ShippingConfiguration bool `json:"shipping_configuration"`
		} `json:"sections"`
	} `json:"config_manager"`

	CouponManager struct {
		Enabled bool `json:"enabled"`
	} `json:"coupon_manager"`
}

type MerchantDashboardConfig struct {
	ID         uuid.UUID                       `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	MerchantID uuid.UUID                       `json:"merchant_id" gorm:"column:merchant_id;type:uuid"`
	Config     MerchantDashboardConfigurations `json:"config" gorm:"column:config;type:jsonb"`
	CreatedAt  time.Time                       `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time                       `json:"updated_at" gorm:"column:updated_at"`
}

func (MerchantDashboardConfig) TableName() string {
	return "merchant_dashboard_configs"
}
