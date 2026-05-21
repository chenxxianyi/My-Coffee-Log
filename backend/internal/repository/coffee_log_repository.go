package repository

import (
	"my-coffee-log/internal/model"
	"my-coffee-log/internal/utils"

	"gorm.io/gorm"
)

type CoffeeLogRepository struct {
	db *gorm.DB
}

func NewCoffeeLogRepository(db *gorm.DB) *CoffeeLogRepository {
	return &CoffeeLogRepository{db: db}
}

func (r *CoffeeLogRepository) Create(log *model.CoffeeLog, flavorTagIDs []uint) error {
	tx := r.db.Begin()
	if err := tx.Create(log).Error; err != nil {
		tx.Rollback()
		return err
	}
	if len(flavorTagIDs) > 0 {
		var tags []model.FlavorTag
		if err := tx.Where("id IN ?", flavorTagIDs).Find(&tags).Error; err != nil {
			tx.Rollback()
			return err
		}
		if err := tx.Model(log).Association("FlavorTags").Replace(tags); err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (r *CoffeeLogRepository) FindByID(id uint, userID uint) (*model.CoffeeLog, error) {
	var log model.CoffeeLog
	if err := r.db.Preload("FlavorTags").Where("id = ? AND user_id = ?", id, userID).First(&log).Error; err != nil {
		return nil, err
	}
	return &log, nil
}

func (r *CoffeeLogRepository) FindList(userID uint, pagination *utils.Pagination, month, coffeeType string, tagID uint) ([]model.CoffeeLog, int64, error) {
	var logs []model.CoffeeLog
	var total int64

	query := r.db.Model(&model.CoffeeLog{}).Where("user_id = ?", userID)

	if month != "" {
		query = query.Where("DATE_FORMAT(drink_date, '%Y-%m') = ?", month)
	}
	if coffeeType != "" {
		query = query.Where("coffee_type = ?", coffeeType)
	}
	if tagID > 0 {
		query = query.Joins("JOIN coffee_log_flavor_tags ON coffee_log_flavor_tags.coffee_log_id = coffee_logs.id").
			Where("coffee_log_flavor_tags.flavor_tag_id = ?", tagID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Preload("FlavorTags").
		Order("drink_date DESC, created_at DESC").
		Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit()).
		Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	pagination.Total = total
	return logs, total, nil
}

func (r *CoffeeLogRepository) Update(log *model.CoffeeLog, flavorTagIDs []uint) error {
	tx := r.db.Begin()
	if err := tx.Model(log).Where("id = ? AND user_id = ?", log.ID, log.UserID).Updates(log).Error; err != nil {
		tx.Rollback()
		return err
	}
	if flavorTagIDs != nil {
		var tags []model.FlavorTag
		if len(flavorTagIDs) > 0 {
			if err := tx.Where("id IN ?", flavorTagIDs).Find(&tags).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
		if err := tx.Model(log).Association("FlavorTags").Replace(tags); err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (r *CoffeeLogRepository) Delete(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&model.CoffeeLog{}).Error
}

func (r *CoffeeLogRepository) FindRecentByUserID(userID uint, limit int) ([]model.CoffeeLog, error) {
	var logs []model.CoffeeLog
	if err := r.db.Preload("FlavorTags").Where("user_id = ?", userID).
		Order("drink_date DESC, created_at DESC").
		Limit(limit).
		Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}
