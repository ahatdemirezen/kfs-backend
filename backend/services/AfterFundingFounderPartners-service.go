package services

import (
	"time"
)

// AfterFundingFounderPartnerRequest service model
type AfterFundingFounderPartnerRequest struct {
	EnterpriseInfoId       uint      `json:"enterprise_info_id" validate:"required"` // Foreign key
	PartnerName            string    `json:"partner_name" validate:"required"`       // Partner's name
	PartnerSurname         string    `json:"partner_surname" validate:"required"`    // Partner's surname
	PartnerTitle           string    `json:"partner_title"`                          // Partner's title
	PartnerSchool          string    `json:"partner_school"`                         // Partner's graduated school
	PartnerGpa             int       `json:"partner_gpa"`                            // Graduation GPA
	ResumeKey              string    `json:"resume_key"`                             // Resume file key
	Citizenship            string    `json:"citizenship"`                            // Citizenship information
	CapitalShareAmount     int       `json:"capital_share_amount" validate:"required"` // Capital share amount
	CapitalSharePercentage int       `json:"capital_share_percentage" validate:"required"` // Capital share percentage
	VotePercentage         int       `json:"vote_percentage" validate:"required"`    // Voting right percentage
	Privilege              string    `json:"privilege"`                              // Privileges
	CampaignRelation       string    `json:"campaign_relation"`                      // Relation to the campaign
	Experience             string    `json:"experience"`                             // Experience
	Profession             string    `json:"profession"`                             // Profession
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}
