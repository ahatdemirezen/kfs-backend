package services



// FinancialExpenseRequest servis modeli
type FinancialExpenseRequest struct {
	ExpenseId     uint   `json:"expense_id"`
	CampaignId    uint   `json:"campaign_id" validate:"required"`
	Year          int    `json:"year" validate:"required"`
	SubCategoryId uint   `json:"sub_category_id" validate:"required"`
	Value         int    `json:"value" validate:"required"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}