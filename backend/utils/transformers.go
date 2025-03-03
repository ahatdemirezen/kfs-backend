package utils

import (
	"kfs-backend/models"
	"kfs-backend/services"
	"time"
)

// ConvertRequestToModel, genel bir request nesnesini ilgili model nesnesine dönüştürür.
func ConvertRequestToModel[R any, T any](req R) T {
    var model T
    switch v := any(req).(type) {
    case services.VentureLocationRequest:
        model = any(models.VentureLocation{
            CampaignId: v.CampaignId,
            Location:   v.Location,
            CreatedAt:  time.Now(),
            UpdatedAt:  time.Now(),
        }).(T)

    case services.VentureCategoryRequest:
        model = any(models.VentureCategory{
            CampaignId: v.CampaignId,
            Category:   v.Category,
            CreatedAt:  time.Now(),
            UpdatedAt:  time.Now(),
        }).(T)

    case services.VentureBusinessModalRequest:
        model = any(models.VentureBusinessModal{
            CampaignId:      v.CampaignId,
            BusinessModal:   v.BusinessModal,
            CreatedAt:       time.Now(),
            UpdatedAt:       time.Now(),
        }).(T)

    case services.VentureSectorRequest:
        model = any(models.VentureSector{
            CampaignId: v.CampaignId,
            Sector:     v.Sector,
            CreatedAt:  time.Now(),
            UpdatedAt:  time.Now(),
        }).(T)

    case services.ParticipantEmailRequest:
        model = any(models.ParticipantEmail{
            CampaignId: v.CampaignId,
            Email:      v.Email,
            CreatedAt:  time.Now(),
            UpdatedAt:  time.Now(),
        }).(T)

    case services.PastCampaignInfoRequest:
        statusPointer := new(bool) // Yeni bir bool pointer oluştur
        *statusPointer = v.Status  // Pointer'a JSON'dan gelen değeri ata
        model = any(models.PastCampaignInfo{
            CampaignId:  v.CampaignId,
            Status:      statusPointer, // Modeldeki Status alanına pointer'ı ata
            Description: v.Description,
            CreatedAt:   time.Now(),
            UpdatedAt:   time.Now(),
        }).(T)

    case services.PatentRequest:
        model = any(models.Patent{
            CampaignId:     v.CampaignId,
            DocumentKey:    v.DocumentKey,
            DocumentNumber: v.DocumentNumber,
            Description:    v.Description,
            CreatedAt:      time.Now(),
            UpdatedAt:      time.Now(),
        }).(T)

    case services.AchievementRequest:
        model = any(models.Achievement{
            CampaignId:   v.CampaignId,
            Date:         v.Date,
            Foundation:   v.Foundation,
            Description:  v.Description,
            DocumentKey:  v.DocumentKey,
            CreatedAt:    time.Now(),
            UpdatedAt:    time.Now(),
        }).(T)

    case services.PermissionRequest:
        model = any(models.Permission{
            CampaignId:   v.CampaignId,
            DocumentKey:  v.DocumentKey,
            Subject:      v.Subject,
            Description:  v.Description,
            CreatedAt:    time.Now(),
            UpdatedAt:    time.Now(),
        }).(T)
    case services.TeamMemberRequest:
        model = any(models.TeamMember{
            CampaignId:   v.CampaignId,
            Name:         v.Name,
            Surname:      v.Surname,
            Position:     v.Position,
            ResumeKey:    v.ResumeKey,
            PhotoKey:     v.PhotoKey,
            Biography:    v.Biography,
            Responsibility: v.Responsibility,
            Profession:     v.Profession,
            Relation:       v.Relation,
            Email:          v.Email,
            Instagram:      v.Instagram,
            Twitter:        v.Twitter,
            Linkedin:       v.Linkedin,
            CreatedAt:      time.Now(),
            UpdatedAt:      time.Now(),
        }).(T)

    case services.ProductModelInfoRequest:
        model = any(models.ProductModelInfo{
            CampaignId:           v.CampaignId,
            ProductSummary:       v.ProductSummary,
            AboutProduct:         v.AboutProduct,
            Problem:              v.Problem,
            Solve:                v.Solve,
            ValueProposition:     v.ValueProposition,
            ProcessSummary:       v.ProcessSummary,
            AboutProcess:         v.AboutProcess,
            AboutSideProduct:     v.AboutSideProduct,
            TechnicalAnalyses:    v.TechnicalAnalyses,
            ArgeSummary:          v.ArgeSummary,
            PreviousSales:        v.PreviousSales,
            AboutProductKey:      v.AboutProductKey,
            ProcessSummaryKey:    v.ProcessSummaryKey,
            AboutProcessKey:      v.AboutProcessKey,
            AboutSideKey:         v.AboutSideKey,
            TechnicalAnalysesKey: v.TechnicalAnalysesKey,
            ArgeSummaryKey:       v.ArgeSummaryKey,
            PreviousSalesKey:     v.PreviousSalesKey,
            CreatedAt:            time.Now(),
            UpdatedAt:            time.Now(),
        }).(T)

    case services.OtherProductTopicRequest:
        model = any(models.OtherProductTopic{
            ProductModelInfoId:  v.ProductModelInfoId,
            DocumentKey:         v.DocumentKey,
            Subject:             v.Subject,
            Description:         v.Description,
            CreatedAt:           time.Now(),
            UpdatedAt:           time.Now(),
        }).(T)

    case services.MarketInfoRequest:
        model = any(models.MarketInfo{
            CampaignId:                 v.CampaignId,
            AboutMarket:                v.AboutMarket,
            AboutCompetition:           v.AboutCompetition,
            TargetSummary:              v.TargetSummary,
            CommercializationSummary:   v.CommercializationSummary,
            AboutMarketKey:             v.AboutMarketKey,
            AboutCompetitionKey:        v.AboutCompetitionKey,
            TargetSummaryKey:           v.TargetSummaryKey,
            CommercializationSummaryKey: v.CommercializationSummaryKey,
            CreatedAt:                  time.Now(),
            UpdatedAt:                  time.Now(),
        }).(T)

    case services.OtherMarketTopicRequest:
        model = any(models.OtherMarketTopic{
            MarketInfoId:  v.MarketInfoId,
            DocumentKey:   v.DocumentKey,
            Subject:       v.Subject,
            Description:   v.Description,
            CreatedAt:     time.Now(),
            UpdatedAt:     time.Now(),
        }).(T)

    case services.AnalysisInfoRequest:
        model = any(models.AnalysisInfo{
            CampaignId:     v.CampaignId,
            SwotKey:        v.SwotKey,
            BusinessKey:    v.BusinessKey,
            InvestorKey:    v.InvestorKey,
            CreatedAt:      time.Now(),
            UpdatedAt:      time.Now(),
        }).(T)

    case services.ProsRequest:
        model = any(models.Pros{
            AnalysisInfoId: v.AnalysisInfoId,
            Pro:            v.Pro,
            CreatedAt:      time.Now(),
            UpdatedAt:      time.Now(),
        }).(T)

    case services.ConsRequest:
        model = any(models.Cons{
            AnalysisInfoId: v.AnalysisInfoId,
            Con:            v.Con,
            CreatedAt:      time.Now(),
            UpdatedAt:      time.Now(),
        }).(T)

    case services.OpportunityRequest:
        model = any(models.Opportunity{
            AnalysisInfoId: v.AnalysisInfoId,
            Opportunity:    v.Opportunity,
            CreatedAt:      time.Now(),
            UpdatedAt:      time.Now(),
        }).(T)

    case services.ThreatRequest:
        model = any(models.Threat{
            AnalysisInfoId: v.AnalysisInfoId,
            Threat:         v.Threat,
            CreatedAt:      time.Now(),
            UpdatedAt:      time.Now(),
        }).(T)
    case services.ConsPlanRequest:
        model = any(models.ConsPlan{
            AnalysisInfoId: v.AnalysisInfoId,
            ConPlan:        v.ConPlan,
            CreatedAt:      time.Now(),
            UpdatedAt:      time.Now(),
        }).(T)
    case services.ThreatPlanRequest:
        model = any(models.ThreatPlan{
            AnalysisInfoId: v.AnalysisInfoId,
            ThreatPlan:     v.ThreatPlan,
            CreatedAt:      time.Now(),
            UpdatedAt:      time.Now(),
        }).(T)
    case services.RisksInfoRequest:
        model = any(models.RisksInfo{
            CampaignId:  v.CampaignId,
            ProjectRisk: v.ProjectRisk,
            SectorRisk:  v.SectorRisk,
            ShareRisk:   v.ShareRisk,
            OtherRisk:   v.OtherRisk,
            CreatedAt:   time.Now(),
            UpdatedAt:   time.Now(),
        }).(T)
    case services.FundingInfoRequest:
        extraFundingPointer := new(bool) // Yeni bir bool pointer oluştur
        *extraFundingPointer = v.ExtraFunding // Pointer'a JSON'dan gelen değeri ata
        model = any(models.FundingInfo{
            CampaignId:           v.CampaignId,
            VentureValue:         v.VentureValue,
            RequiredVentureFund:  v.RequiredVentureFund,
            FundingMonths:        v.FundingMonths,
            EvaluationReportKey:  v.EvaluationReportKey,
            SharePercentage:      v.SharePercentage,
            ExtraFunding:         extraFundingPointer, // Modeldeki ExtraFunding alanına pointer'ı ata
            ComparingPartnership: v.ComparingPartnership,
            GeneralReason:        v.GeneralReason,
            CreatedAt:            time.Now(),
            UpdatedAt:            time.Now(),
        }).(T)    
    case services.UsageRequest:
        model = any(models.Usage{
            FundingInfoId: v.FundingInfoId,
            Description:   v.Description,
            StartingDate:  v.StartingDate,
            EndingDate:    v.EndingDate,
            Amount:        v.Amount,
            CreatedAt:     time.Now(),
            UpdatedAt:     time.Now(),
        }).(T)
    case services.ExtraFinancingResourceRequest:
        model = any(models.ExtraFinancingResource{
            FundingInfoId:     v.FundingInfoId,
            Description:       v.Description,
            SupplyDate:        v.SupplyDate,
            Amount:            v.Amount,
            CreatedAt:         time.Now(),
            UpdatedAt:         time.Now(),
        }).(T)
    case services.EnterpriseInfoRequest:
        model = any(models.EnterpriseInfo{
            CampaignId:        v.CampaignId,
            EnterpriseName:    v.EnterpriseName,
            EnterpriseCapital: v.EnterpriseCapital,
            EnterpriseCity:    v.EnterpriseCity,
            EnterpriseTown:    v.EnterpriseTown,
            EnterpriseAddress: v.EnterpriseAddress,
            CreatedAt:         time.Now(),
            UpdatedAt:         time.Now(),
        }).(T)

    case services.AfterFundingFounderPartnerRequest:
        model = any(models.AfterFundingFounderPartner{
            EnterpriseInfoId:       v.EnterpriseInfoId,
            PartnerName:            v.PartnerName,
            PartnerSurname:         v.PartnerSurname,
            PartnerTitle:           v.PartnerTitle,
            PartnerSchool:          v.PartnerSchool,
            PartnerGpa:             v.PartnerGpa,
            ResumeKey:              v.ResumeKey,
            Citizenship:            v.Citizenship,
            CapitalShareAmount:     v.CapitalShareAmount,
            CapitalSharePercentage: v.CapitalSharePercentage,
            VotePercentage:         v.VotePercentage,
            Privilege:              v.Privilege,
            CampaignRelation:       v.CampaignRelation,
            Experience:             v.Experience,
            Profession:             v.Profession,
            CreatedAt:              time.Now(),
            UpdatedAt:              time.Now(),
        }).(T)
    case services.VisualInfoRequest:
        model = any(models.VisualInfo{
            CampaignId:       v.CampaignId,
            ShowcasePhotoKey: v.ShowcasePhotoKey,
            CreatedAt:        time.Now(),
            UpdatedAt:        time.Now(),
        }).(T)
    case services.VideosRequest:
        model = any(models.Videos{
            VisualInfoId: v.VisualInfoId,
            VideoUrl:     v.VideoUrl,
            CreatedAt:    time.Now(),
            UpdatedAt:    time.Now(),
        }).(T)
    case services.OtherPhotosRequest:
        model = any(models.OtherPhotos{
            VisualInfoId: v.VisualInfoId,
            PhotoKey:     v.PhotoKey,
            CreatedAt:    time.Now(),
            UpdatedAt:    time.Now(),
        }).(T)
        case services.OtherDocumentsInfoRequest:
        model = any(models.OtherDocumentsInfo{
            CampaignId:  v.CampaignId,
            DocumentKey: v.DocumentKey,
            CreatedAt:   time.Now(),
            UpdatedAt:   time.Now(),
        }).(T)
        case services.ProfitForecastRequest:
        model = any(models.ProfitForecast{
            CampaignId:     v.CampaignId,
            ProfitForecast: v.ProfitForecast,
            CreatedAt:      time.Now(),
            UpdatedAt:      time.Now(),
        }).(T)
        case services.InvestmentBudgetsRequest:
        model = any(models.InvestmentBudgets{
            CampaignId:      v.CampaignId,
            YearOneBudget:   v.YearOneBudget,
            YearTwoBudget:   v.YearTwoBudget,
            YearThreeBudget: v.YearThreeBudget,
            YearFourBudget:  v.YearFourBudget,
            YearFiveBudget:  v.YearFiveBudget,
            CreatedAt:       time.Now(),
            UpdatedAt:       time.Now(),
        }).(T)
    case services.IncomeItemsRequest:
        model = any(models.IncomeItems{
            CampaignId:   v.CampaignId,
            Title:        v.Title,
            SalePrice:    v.SalePrice,
            Cost:         v.Cost,
            CreatedAt:    time.Now(),
            UpdatedAt:    time.Now(),
        }).(T)
    case services.FinancialDocumentsRequest:
        model = any(models.FinancialDocuments{
            CampaignId:          v.CampaignId,
            Subject:             v.Subject,
            DocumentKey:         v.DocumentKey,
            CreatedAt:           time.Now(),
            UpdatedAt:           time.Now(),
        }).(T)
        case services.ExplanationsRequest:
        model = any(models.Explanations{
            CampaignId:    v.CampaignId,
            Explanation:   v.Explanation,
            CreatedAt:     time.Now(),
            UpdatedAt:     time.Now(),
        }).(T)
        case services.SaleGoalsRequest:
        model = any(models.SaleGoals{
            CampaignId:   v.CampaignId,
            IncomeItemId: v.IncomeItemId,
            YearOne:      v.YearOne,
            YearTwo:      v.YearTwo,
            YearThree:    v.YearThree,
            YearFour:     v.YearFour,
            YearFive:     v.YearFive,
            CreatedAt:    time.Now(),
            UpdatedAt:    time.Now(),
        }).(T)
        case services.FinancialExpenseRequest:
		model = any(models.FinancialExpense{
			CampaignId:    v.CampaignId,
			Year:          v.Year,
			SubCategoryId: v.SubCategoryId,
			Value:         v.Value,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}).(T)
        case services.FinancialCategoryRequest:
		model = any(models.FinancialCategory{
			Category:   v.Category,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}).(T)
        case services.FinancialSubCategoryRequest:
		model = any(models.FinancialSubCategory{
			CategoryId:  v.CategoryId,
			SubCategory: v.SubCategory,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}).(T)

    }
    return model
}
