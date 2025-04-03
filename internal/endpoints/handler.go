package endpoints

import "gomail/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
