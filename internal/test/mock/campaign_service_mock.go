package mock

import (
	"gomail/internal/domain/campaign/contract"

	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}

func (r *CampaignServiceMock) Create(newCampaign contract.NewCampaign) (string, error) {
	args := r.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func (r *CampaignServiceMock) GetBy(id string) (*contract.CampaignResponse, error) {
	//args := r.Called(id)
	return nil, nil
}

func (r *CampaignServiceMock) Cancel(id string) error {
	return nil
}
