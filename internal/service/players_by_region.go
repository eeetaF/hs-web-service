package service

import (
	"context"
	"encoding/json"
	"fmt"
	"hs-web-service/internal/model"
	"io"
)

func (i *Implementation) GetPlayerByAccountID(_ context.Context, accountID string) (*model.Player, error) {
	resp, err := i.httpClient.Get("https://hearthstone.blizzard.com/api/community/leaderboardsData?region=US&leaderboardId=battlegrounds&page=1")
	if err != nil {
		return nil, fmt.Errorf("http Get: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("http Get: %w", err)
	}
	var retrievedData model.RetrievedLeaderboardData

	err = json.Unmarshal(body, &retrievedData)
	if err != nil {
		return nil, fmt.Errorf("json Unmarshal: %w", err)
	}
	fmt.Println(retrievedData)
	return nil
}
