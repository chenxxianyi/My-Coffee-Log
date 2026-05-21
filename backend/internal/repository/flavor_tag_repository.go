package repository

import (
	"my-coffee-log/internal/model"

	"gorm.io/gorm"
)

type FlavorTagRepository struct {
	db *gorm.DB
}

func NewFlavorTagRepository(db *gorm.DB) *FlavorTagRepository {
	return &FlavorTagRepository{db: db}
}

func (r *FlavorTagRepository) FindAll() ([]model.FlavorTag, error) {
	var tags []model.FlavorTag
	if err := r.db.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *FlavorTagRepository) FindByIDs(ids []uint) ([]model.FlavorTag, error) {
	var tags []model.FlavorTag
	if err := r.db.Where("id IN ?", ids).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *FlavorTagRepository) FindByNames(names []string) ([]model.FlavorTag, error) {
	var tags []model.FlavorTag
	if err := r.db.Where("name IN ?", names).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *FlavorTagRepository) Seed(tags []model.FlavorTag) error {
	for _, tag := range tags {
		var existing model.FlavorTag
		if err := r.db.Where("name = ?", tag.Name).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := r.db.Create(&tag).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}
	return nil
}
