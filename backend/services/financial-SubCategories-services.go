package services



// FinancialSubCategoryRequest servis modeli
type FinancialSubCategoryRequest struct {
	SubCategoryId uint   `json:"sub_category_id"`
	CategoryId    uint   `json:"category_id" validate:"required"`
	SubCategory   string `json:"sub_category" validate:"required"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}