package campaign

import (
	"errors"
	"testing"

	"github.com/renatocardosoalves/email-wave/internal/contract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCampaignRepository struct {
	mock.Mock
}

func (m *MockCampaignRepository) Save(campaign *Campaign) error {
	args := m.Called(campaign)
	return args.Error(0)
}

func TestCampaignService_Save(t *testing.T) {
	repo := new(MockCampaignRepository)

	service := CampaignService{Repository: repo}

	newCampaign := contract.NewCampaign{
		Name:    "Campaign Name",
		Content: "Campaign Content",
		Emails:  []string{"username1@company.com", "username2@company.com"},
	}

	matcher := mock.MatchedBy(func(campaign *Campaign) bool {
		return campaign.Name == newCampaign.Name && campaign.Content == newCampaign.Content && len(campaign.Contacts) == len(newCampaign.Emails)
	})

	repo.On("Save", matcher).Return(nil)

	id, err := service.Save(newCampaign)

	repo.AssertExpectations(t)

	assert.NotNil(t, id)
	assert.Nil(t, err)
}

func TestCampaignService_Save_Error(t *testing.T) {
	repo := new(MockCampaignRepository)

	service := CampaignService{Repository: repo}

	newCampaign := contract.NewCampaign{
		Name:    "Campaign Name",
		Content: "Campaign Content",
		Emails:  []string{"username1@company.com", "username2@company.com"},
	}

	repo.On("Save", mock.Anything).Return(assert.AnError)

	id, err := service.Save(newCampaign)

	repo.AssertExpectations(t)

	assert.Empty(t, id)
	assert.Error(t, err)
}

func TestCampaignService_Save_ValidateDomainError(t *testing.T) {
	repo := new(MockCampaignRepository)

	service := CampaignService{Repository: repo}

	newCampaign := contract.NewCampaign{
		Name:    "",
		Content: "Campaign Content",
		Emails:  []string{"username1@company.com", "username2@company.com"},
	}

	expectedError := errors.New("name is required")

	repo.On("Save", mock.Anything).Return(expectedError)

	id, err := service.Save(newCampaign)

	assert.Empty(t, id)
	assert.EqualError(t, expectedError, err.Error())
}
