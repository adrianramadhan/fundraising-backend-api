package campaign

import "fundraising-backend-api/user"

type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Name             string    `json:"name"`
	ShortDescription string    `json:"short_description"`
	Description      string    `json:"description"`
	GoalAmount       int       `json:"goal_amount"`
	Perks            string    `json:"perks"`
	User             user.User `json:"user"`
}

type CreateCampaignImageInput struct {
	CampaignID int  `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary" binding:"required"`
}