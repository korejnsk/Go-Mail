package main

import (
	"gomail/internal/domain/campaign"
	"gomail/internal/endpoints"
	"gomail/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.NewDb()
	campaignService := &campaign.ServiceImp{
		Repository: &database.CampaignRepository{Db: db},
	}
	handler := endpoints.Handler{
		CampaignService: campaignService,
	}

	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/campaigns/{id}", endpoints.HandlerError(handler.CampaignGetById))

	http.ListenAndServe(":3000", r)
}
