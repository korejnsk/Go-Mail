package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gomail/internal/domain/campaign/contract"
	internalmock "gomail/internal/test/mock"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Campaign_should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaign{
		Name:    "teste",
		Content: "Hi everyone",
		Emails:  []string{"teste@teste.com"},
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		if request.Name == body.Name && request.Content == body.Content {
			return true
		} else {
			return false
		}
	})).Return("123", nil)
	handler := Handler{CampaignService: service}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req := httptest.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignPost(rr, req)

	assert.Equal(201, status)
	assert.Nil(err)
}

func Test_Campaign_should_inform_error(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaign{
		Name:    "teste",
		Content: "Hi everyone",
		Emails:  []string{"teste@teste.com"},
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{CampaignService: service}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req := httptest.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()

	_, _, err := handler.CampaignPost(rr, req)

	assert.NotNil(err)
}
