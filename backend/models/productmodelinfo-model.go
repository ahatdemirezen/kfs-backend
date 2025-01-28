package models

import (
	"time"
)

// ProductModelInfo modeli
type ProductModelInfo struct {
	ProductModelInfoId uint      `gorm:"primaryKey;autoIncrement;column:product_model_info_id"`
	CampaignId         uint      `gorm:"not null;column:campaign_id"` // Foreign key
	ProductSummary     string    `gorm:"type:text;column:product_summary"`
	AboutProduct       string    `gorm:"type:text;column:about_product"`
	Problem            string    `gorm:"type:text;column:problem"`
	Solve              string    `gorm:"type:text;column:solve"`
	ValueProposition   string    `gorm:"type:text;column:value_proposition"`
	ProcessSummary     string    `gorm:"type:text;column:process_summary"`
	AboutProcess       string    `gorm:"type:text;column:about_process"`
	AboutSideProduct   string    `gorm:"type:text;column:about_side_product"`
	TechnicalAnalyses  string    `gorm:"type:text;column:technical_analyses"`
	ArgeSummary       string    `gorm:"type:text;column:angel_summary"`
	PreviousSales      string    `gorm:"type:text;column:previous_sales"`
	AboutProductKey    string    `gorm:"type:text;column:about_product_key"`
	ProcessSummaryKey  string    `gorm:"type:text;column:process_summary_key"`
	AboutProcessKey    string    `gorm:"type:text;column:about_process_key"`
	AboutSideKey       string    `gorm:"type:text;column:about_side_key"`
	TechnicalAnalysesKey string  `gorm:"type:text;column:technical_analyses_key"`
	ArgeSummaryKey    string    `gorm:"type:text;column:angel_summary_key"`
	PreviousSalesKey   string    `gorm:"type:text;column:previous_sales_key"`
	CreatedAt          time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt          time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId"`
}

// TableName tablosunun adını belirtir
func (ProductModelInfo) TableName() string {
	return "product_model_info"
}
