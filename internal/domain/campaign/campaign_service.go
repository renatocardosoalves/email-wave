package campaign

import "github.com/renatocardosoalves/email-wave/internal/contract"

type CampaignService struct {
	Repository CampaignRepository
}

func (s *CampaignService) Save(newCampaign contract.NewCampaign) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails...)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)

	if err != nil {
		return "", err
	}

	return campaign.ID, nil
}
