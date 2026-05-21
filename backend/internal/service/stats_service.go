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

type OverviewResponse struct {
	MonthCount        int64  `json:"month_count"`
	TotalCount        int64  `json:"total_count"`
	FavoriteCoffeeType string `json:"favorite_coffee_type"`
	FavoriteFlavorTag  string `json:"favorite_flavor_tag"`
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

	return &OverviewResponse{
		MonthCount:        monthCount,
		TotalCount:        totalCount,
		FavoriteCoffeeType: favType,
		FavoriteFlavorTag:  favTag,
	}, nil
}

func (s *StatsService) GetFlavorProfile(userID uint) (*repository.FlavorProfile, error) {
	return s.statsRepo.GetFlavorProfile(userID)
}

func (s *StatsService) GetMonthly(userID uint) ([]repository.MonthlyCount, error) {
	return s.statsRepo.GetMonthlyCounts(userID)
}
