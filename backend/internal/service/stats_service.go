package service

import (
	"my-coffee-log/internal/repository"
)

type StatsService struct {
	statsRepo *repository.StatsRepository
}

func NewStatsService(statsRepo *repository.StatsRepository) *StatsService {
	return &StatsService{statsRepo: statsRepo}
}

type FlavorTagItem struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Count int64  `json:"count"`
}

type OverviewResponse struct {
	MonthCount         int64           `json:"month_count"`
	TotalCount         int64           `json:"total_count"`
	FavoriteCoffeeType string          `json:"favorite_coffee_type"`
	FavoriteFlavorTag  string          `json:"favorite_flavor_tag"`
	RecentFlavorTags   []FlavorTagItem `json:"recent_flavor_tags"`
}

func (s *StatsService) GetOverview(userID uint) (*OverviewResponse, error) {
	monthCount, err := s.statsRepo.GetMonthCount(userID)
	if err != nil {
		return nil, err
	}

	totalCount, err := s.statsRepo.GetTotalCount(userID)
	if err != nil {
		return nil, err
	}

	favType, err := s.statsRepo.GetFavoriteCoffeeType(userID)
	if err != nil {
		return nil, err
	}

	favTag, err := s.statsRepo.GetFavoriteFlavorTag(userID)
	if err != nil {
		return nil, err
	}

	recentTags, err := s.statsRepo.GetRecentFlavorTags(userID, 5)
	if err != nil {
		return nil, err
	}

	tagItems := make([]FlavorTagItem, 0, len(recentTags))
	for _, t := range recentTags {
		tagItems = append(tagItems, FlavorTagItem{
			Name:  t.Name,
			Label: t.Label,
			Count: t.Count,
		})
	}

	return &OverviewResponse{
		MonthCount:         monthCount,
		TotalCount:         totalCount,
		FavoriteCoffeeType: favType,
		FavoriteFlavorTag:  favTag,
		RecentFlavorTags:   tagItems,
	}, nil
}

func (s *StatsService) GetFlavorProfile(userID uint) (*repository.FlavorProfile, error) {
	return s.statsRepo.GetFlavorProfile(userID)
}

func (s *StatsService) GetMonthly(userID uint) ([]repository.MonthlyCount, error) {
	return s.statsRepo.GetMonthlyCounts(userID)
}

func (s *StatsService) GetRecentMood(userID uint) (string, error) {
	return s.statsRepo.GetRecentMood(userID)
}
