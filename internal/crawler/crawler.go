package crawler

import (
	"context"
	"hs-web-service/internal/model"
	"net/http"
)

type Crawler interface {
	LoadLeaderboardPage(ctx context.Context, region model.Region, Leaderboard model.LeaderboardID, page int) ([]model.Player, error)
}

type Implementation struct {
	httpClient *http.Client
}

func New() *Implementation {
	return &Implementation{httpClient: http.DefaultClient}
}
