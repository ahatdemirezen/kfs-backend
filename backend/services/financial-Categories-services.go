package services



// FinancialCategoryRequest servis modeli
type FinancialCategoryRequest struct {
	CategoryId uint   `json:"category_id"`
	Category   string `json:"category" validate:"required"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
