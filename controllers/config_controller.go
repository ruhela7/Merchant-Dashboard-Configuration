package controllers

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
)

type ConfigController struct {
	configService *services.ConfigService
}

func NewConfigController(db *gorm.DB) *ConfigController {
	return &ConfigController{
		configService: services.NewConfigService(db),
	}
}

func (c *ConfigController) GetConfigHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	merchantID := "e9257895-df22-4dd7-9ef5-5eedf647123a"
	config, err := c.configService.GetConfigByMerchantID(merchantID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to retrieve config: %v", err)
		return
	}

	jsonData, err := json.Marshal(config)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to encode config: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (c *ConfigController) UpdateConfigHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var config models.MerchantDashboardConfig
	err := json.NewDecoder(r.Body).Decode(&config)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body: %v", err)
		return
	}

	err = c.configService.CreateConfig(&config)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to create config: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Config updated successfully")
}
