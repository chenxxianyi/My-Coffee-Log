package repository

import (
	"encoding/json"
	"sort"
	"time"

	"my-coffee-log/internal/model"

	"gorm.io/gorm"
)

type StatsRepository struct {
	db *gorm.DB
}

func NewStatsRepository(db *gorm.DB) *StatsRepository {
	return &StatsRepository{db: db}
}

func (r *StatsRepository) GetMonthCount(userID uint) (int64, error) {
	var count int64
	now := time.Now()
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	if err := r.db.Model(&model.CoffeeLog{}).
		Where("user_id = ? AND drink_date >= ? AND is_test_data = ?", userID, monthStart, false).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *StatsRepository) GetTotalCount(userID uint) (int64, error) {
	var count int64
	if err := r.db.Model(&model.CoffeeLog{}).
		Where("user_id = ? AND is_test_data = ?", userID, false).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// GetValidSensoryCount returns the count of records where sensory scores were actively recorded
func (r *StatsRepository) GetValidSensoryCount(userID uint) (int64, error) {
	var count int64
	if err := r.db.Model(&model.CoffeeLog{}).
		Where("user_id = ? AND sensory_recorded = ? AND is_test_data = ?", userID, true, false).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// GetDateRange returns the earliest and latest drink dates for a user
func (r *StatsRepository) GetDateRange(userID uint) (string, string) {
	type dateResult struct {
		DateFrom string `gorm:"column:date_from"`
		DateTo   string `gorm:"column:date_to"`
	}
	var result dateResult
	r.db.Model(&model.CoffeeLog{}).
		Select("CAST(MIN(drink_date) AS CHAR) as date_from, CAST(MAX(drink_date) AS CHAR) as date_to").
		Where("user_id = ? AND is_test_data = ?", userID, false).
		Scan(&result)
	return result.DateFrom, result.DateTo
}

func (r *StatsRepository) GetFavoriteCoffeeType(userID uint) (string, error) {
	var result struct {
		CoffeeType string `json:"coffee_type"`
		Count      int64  `json:"count"`
	}
	if err := r.db.Model(&model.CoffeeLog{}).
		Select("coffee_type, COUNT(*) as count").
		Where("user_id = ?", userID).
		Group("coffee_type").
		Order("count DESC").
		First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		return "", err
	}
	return result.CoffeeType, nil
}

func (r *StatsRepository) GetFavoriteFlavorTag(userID uint) (string, error) {
	var result struct {
		FlavorTagID uint  `json:"flavor_tag_id"`
		Count       int64 `json:"count"`
	}
	if err := r.db.Table("coffee_log_flavor_tags").
		Select("flavor_tag_id, COUNT(*) as count").
		Joins("JOIN coffee_logs ON coffee_logs.id = coffee_log_flavor_tags.coffee_log_id").
		Where("coffee_logs.user_id = ?", userID).
		Group("flavor_tag_id").
		Order("count DESC").
		First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		return "", err
	}

	var tag model.FlavorTag
	if err := r.db.First(&tag, result.FlavorTagID).Error; err != nil {
		return "", err
	}
	return tag.Name, nil
}

type FlavorProfile struct {
	Acidity    float64 `json:"acidity"`
	Bitterness float64 `json:"bitterness"`
	Sweetness  float64 `json:"sweetness"`
	Body       float64 `json:"body"`
	Aroma      float64 `json:"aroma"`
	Aftertaste float64 `json:"aftertaste"`
}

func (r *StatsRepository) GetFlavorProfile(userID uint) (*FlavorProfile, error) {
	var profile FlavorProfile
	if err := r.db.Model(&model.CoffeeLog{}).
		Select("COALESCE(AVG(acidity), 0) as acidity, COALESCE(AVG(bitterness), 0) as bitterness, COALESCE(AVG(sweetness), 0) as sweetness, COALESCE(AVG(body), 0) as body, COALESCE(AVG(aroma), 0) as aroma, COALESCE(AVG(aftertaste), 0) as aftertaste").
		Where("user_id = ? AND sensory_recorded = ? AND is_test_data = ?", userID, true, false).
		Scan(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

type FlavorTagCount struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Count int64  `json:"count"`
}

func (r *StatsRepository) GetRecentFlavorTags(userID uint, limit int) ([]FlavorTagCount, error) {
	if limit <= 0 || limit > 20 {
		limit = 5
	}

	since := time.Now().AddDate(0, -1, 0)

	var results []FlavorTagCount
	if err := r.db.Table("coffee_log_flavor_tags").
		Select("flavor_tags.name, flavor_tags.label, COUNT(*) as count").
		Joins("JOIN coffee_logs ON coffee_logs.id = coffee_log_flavor_tags.coffee_log_id").
		Joins("JOIN flavor_tags ON flavor_tags.id = coffee_log_flavor_tags.flavor_tag_id").
		Where("coffee_logs.user_id = ? AND coffee_logs.drink_date >= ?", userID, since).
		Group("flavor_tags.id, flavor_tags.name, flavor_tags.label").
		Order("count DESC").
		Limit(limit).
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (r *StatsRepository) GetRecentMood(userID uint) (string, error) {
	var result struct {
		Mood string `json:"mood"`
	}
	if err := r.db.Model(&model.CoffeeLog{}).
		Select("mood").
		Where("user_id = ?", userID).
		Order("drink_date DESC, created_at DESC").
		First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		return "", err
	}
	return result.Mood, nil
}

type LifestyleTagCount struct {
	Tag   string `json:"tag"`
	Count int64  `json:"count"`
}

func (r *StatsRepository) GetLifestyleTagCounts(userID uint, column string) ([]LifestyleTagCount, error) {
	var results []LifestyleTagCount
	rows, err := r.db.Model(&model.CoffeeLog{}).
		Select(column).
		Where("user_id = ? AND "+column+" != ''", userID).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tagCounts := map[string]int64{}
	for rows.Next() {
		var jsonStr string
		if err := rows.Scan(&jsonStr); err != nil {
			continue
		}
		var tags []string
		if err := json.Unmarshal([]byte(jsonStr), &tags); err != nil {
			continue
		}
		for _, t := range tags {
			tagCounts[t]++
		}
	}

	for tag, count := range tagCounts {
		results = append(results, LifestyleTagCount{Tag: tag, Count: count})
	}
	sort.Slice(results, func(i, j int) bool { return results[i].Count > results[j].Count })
	if len(results) > 10 {
		results = results[:10]
	}
	return results, nil
}

type MonthlyCount struct {
	Month string `json:"month"`
	Count int64  `json:"count"`
}

func (r *StatsRepository) GetMonthlyCounts(userID uint) ([]MonthlyCount, error) {
	var monthly []MonthlyCount
	if err := r.db.Model(&model.CoffeeLog{}).
		Select("DATE_FORMAT(drink_date, '%Y-%m') as month, COUNT(*) as count").
		Where("user_id = ?", userID).
		Group("month").
		Order("month DESC").
		Find(&monthly).Error; err != nil {
		return nil, err
	}
	return monthly, nil
}

// ---- Monthly Review Aggregation ----

type CoffeeTypeCount struct {
	CoffeeType string `json:"coffee_type"`
	Count      int64  `json:"count"`
}

type CoffeeNameCount struct {
	CoffeeName string `json:"coffee_name"`
	Count      int64  `json:"count"`
}

type WeekdayCount struct {
	Weekday int   `json:"weekday"`
	Count   int64 `json:"count"`
}

func (r *StatsRepository) GetMonthCountByMonth(userID uint, month string) (int64, error) {
	var count int64
	if err := r.db.Model(&model.CoffeeLog{}).
		Where("user_id = ? AND DATE_FORMAT(drink_date, '%Y-%m') = ?", userID, month).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *StatsRepository) GetFavoriteCoffeeTypeByMonth(userID uint, month string) (string, error) {
	var result CoffeeTypeCount
	if err := r.db.Model(&model.CoffeeLog{}).
		Select("coffee_type, COUNT(*) as count").
		Where("user_id = ? AND DATE_FORMAT(drink_date, '%Y-%m') = ?", userID, month).
		Group("coffee_type").
		Order("count DESC").
		First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		return "", err
	}
	return result.CoffeeType, nil
}

func (r *StatsRepository) GetTopCoffeeTypesByMonth(userID uint, month string, limit int) ([]CoffeeTypeCount, error) {
	if limit <= 0 || limit > 10 {
		limit = 5
	}
	var results []CoffeeTypeCount
	if err := r.db.Model(&model.CoffeeLog{}).
		Select("coffee_type, COUNT(*) as count").
		Where("user_id = ? AND DATE_FORMAT(drink_date, '%Y-%m') = ?", userID, month).
		Group("coffee_type").
		Order("count DESC").
		Limit(limit).
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (r *StatsRepository) GetTopFlavorTagsByMonth(userID uint, month string, limit int) ([]FlavorTagCount, error) {
	if limit <= 0 || limit > 20 {
		limit = 5
	}
	var results []FlavorTagCount
	if err := r.db.Table("coffee_log_flavor_tags").
		Select("flavor_tags.name, flavor_tags.label, COUNT(*) as count").
		Joins("JOIN coffee_logs ON coffee_logs.id = coffee_log_flavor_tags.coffee_log_id").
		Joins("JOIN flavor_tags ON flavor_tags.id = coffee_log_flavor_tags.flavor_tag_id").
		Where("coffee_logs.user_id = ? AND DATE_FORMAT(coffee_logs.drink_date, '%Y-%m') = ?", userID, month).
		Group("flavor_tags.id, flavor_tags.name, flavor_tags.label").
		Order("count DESC").
		Limit(limit).
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (r *StatsRepository) GetTopCoffeeNamesByMonth(userID uint, month string, limit int) ([]CoffeeNameCount, error) {
	if limit <= 0 || limit > 10 {
		limit = 5
	}
	var results []CoffeeNameCount
	if err := r.db.Model(&model.CoffeeLog{}).
		Select("coffee_name, COUNT(*) as count").
		Where("user_id = ? AND DATE_FORMAT(drink_date, '%Y-%m') = ? AND coffee_name != ''", userID, month).
		Group("coffee_name").
		Order("count DESC").
		Limit(limit).
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (r *StatsRepository) GetTopDrinkWeekdayByMonth(userID uint, month string) (*WeekdayCount, error) {
	var result WeekdayCount
	if err := r.db.Model(&model.CoffeeLog{}).
		Select("DAYOFWEEK(drink_date) as weekday, COUNT(*) as count").
		Where("user_id = ? AND DATE_FORMAT(drink_date, '%Y-%m') = ?", userID, month).
		Group("weekday").
		Order("count DESC").
		First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

func (r *StatsRepository) GetLifestyleTagCountsByMonth(userID uint, month string, column string) ([]LifestyleTagCount, error) {
	var results []LifestyleTagCount
	rows, err := r.db.Model(&model.CoffeeLog{}).
		Select(column).
		Where("user_id = ? AND DATE_FORMAT(drink_date, '%Y-%m') = ? AND "+column+" != ''", userID, month).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tagCounts := map[string]int64{}
	for rows.Next() {
		var jsonStr string
		if err := rows.Scan(&jsonStr); err != nil {
			continue
		}
		var tags []string
		if err := json.Unmarshal([]byte(jsonStr), &tags); err != nil {
			continue
		}
		for _, t := range tags {
			tagCounts[t]++
		}
	}

	for tag, count := range tagCounts {
		results = append(results, LifestyleTagCount{Tag: tag, Count: count})
	}
	sort.Slice(results, func(i, j int) bool { return results[i].Count > results[j].Count })
	if len(results) > 10 {
		results = results[:10]
	}
	return results, nil
}

func (r *StatsRepository) GetFlavorProfileByMonth(userID uint, month string) (*FlavorProfile, error) {
	var profile FlavorProfile
	if err := r.db.Model(&model.CoffeeLog{}).
		Select("COALESCE(AVG(acidity), 0) as acidity, COALESCE(AVG(bitterness), 0) as bitterness, COALESCE(AVG(sweetness), 0) as sweetness, COALESCE(AVG(body), 0) as body, COALESCE(AVG(aroma), 0) as aroma, COALESCE(AVG(aftertaste), 0) as aftertaste").
		Where("user_id = ? AND DATE_FORMAT(drink_date, '%Y-%m') = ?", userID, month).
		Scan(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

// ---- Date Range Methods for Weekly Review ----

func (r *StatsRepository) GetCountByDateRange(userID uint, startDate, endDate string) (int64, error) {
	var count int64
	if err := r.db.Model(&model.CoffeeLog{}).
		Where("user_id = ? AND drink_date >= ? AND drink_date <= ?", userID, startDate, endDate).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *StatsRepository) GetFavoriteCoffeeTypeByDateRange(userID uint, startDate, endDate string) (string, error) {
	var result CoffeeTypeCount
	if err := r.db.Model(&model.CoffeeLog{}).
		Select("coffee_type, COUNT(*) as count").
		Where("user_id = ? AND drink_date >= ? AND drink_date <= ?", userID, startDate, endDate).
		Group("coffee_type").
		Order("count DESC").
		First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		return "", err
	}
	return result.CoffeeType, nil
}

func (r *StatsRepository) GetTopCoffeeTypesByDateRange(userID uint, startDate, endDate string, limit int) ([]CoffeeTypeCount, error) {
	if limit <= 0 || limit > 10 {
		limit = 5
	}
	var results []CoffeeTypeCount
	if err := r.db.Model(&model.CoffeeLog{}).
		Select("coffee_type, COUNT(*) as count").
		Where("user_id = ? AND drink_date >= ? AND drink_date <= ?", userID, startDate, endDate).
		Group("coffee_type").
		Order("count DESC").
		Limit(limit).
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (r *StatsRepository) GetTopFlavorTagsByDateRange(userID uint, startDate, endDate string, limit int) ([]FlavorTagCount, error) {
	if limit <= 0 || limit > 20 {
		limit = 5
	}
	var results []FlavorTagCount
	if err := r.db.Table("coffee_log_flavor_tags").
		Select("flavor_tags.name, flavor_tags.label, COUNT(*) as count").
		Joins("JOIN coffee_logs ON coffee_logs.id = coffee_log_flavor_tags.coffee_log_id").
		Joins("JOIN flavor_tags ON flavor_tags.id = coffee_log_flavor_tags.flavor_tag_id").
		Where("coffee_logs.user_id = ? AND coffee_logs.drink_date >= ? AND coffee_logs.drink_date <= ?", userID, startDate, endDate).
		Group("flavor_tags.id, flavor_tags.name, flavor_tags.label").
		Order("count DESC").
		Limit(limit).
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (r *StatsRepository) GetLifestyleTagCountsByDateRange(userID uint, startDate, endDate string, column string) ([]LifestyleTagCount, error) {
	var results []LifestyleTagCount
	rows, err := r.db.Model(&model.CoffeeLog{}).
		Select(column).
		Where("user_id = ? AND drink_date >= ? AND drink_date <= ? AND "+column+" != ''", userID, startDate, endDate).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tagCounts := map[string]int64{}
	for rows.Next() {
		var jsonStr string
		if err := rows.Scan(&jsonStr); err != nil {
			continue
		}
		var tags []string
		if err := json.Unmarshal([]byte(jsonStr), &tags); err != nil {
			continue
		}
		for _, t := range tags {
			tagCounts[t]++
		}
	}

	for tag, count := range tagCounts {
		results = append(results, LifestyleTagCount{Tag: tag, Count: count})
	}
	sort.Slice(results, func(i, j int) bool { return results[i].Count > results[j].Count })
	if len(results) > 10 {
		results = results[:10]
	}
	return results, nil
}

func (r *StatsRepository) GetFlavorProfileByDateRange(userID uint, startDate, endDate string) (*FlavorProfile, error) {
	var profile FlavorProfile
	if err := r.db.Model(&model.CoffeeLog{}).
		Select("COALESCE(AVG(acidity), 0) as acidity, COALESCE(AVG(bitterness), 0) as bitterness, COALESCE(AVG(sweetness), 0) as sweetness, COALESCE(AVG(body), 0) as body, COALESCE(AVG(aroma), 0) as aroma, COALESCE(AVG(aftertaste), 0) as aftertaste").
		Where("user_id = ? AND drink_date >= ? AND drink_date <= ?", userID, startDate, endDate).
		Scan(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}
