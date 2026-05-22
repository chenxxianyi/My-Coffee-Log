package repository

import (
	"my-coffee-log/internal/model"
	"my-coffee-log/internal/utils"

	"gorm.io/gorm"
)

type CoffeeBeanRepository struct {
	db *gorm.DB
}

func NewCoffeeBeanRepository(db *gorm.DB) *CoffeeBeanRepository {
	return &CoffeeBeanRepository{db: db}
}

func (r *CoffeeBeanRepository) Create(bean *model.CoffeeBean) error {
	return r.db.Create(bean).Error
}

func (r *CoffeeBeanRepository) FindByID(id, userID uint) (*model.CoffeeBean, error) {
	var bean model.CoffeeBean
	if err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&bean).Error; err != nil {
		return nil, err
	}
	return &bean, nil
}

func (r *CoffeeBeanRepository) FindList(userID uint, pagination *utils.Pagination, search string) ([]model.CoffeeBean, int64, error) {
	var beans []model.CoffeeBean
	var total int64

	query := r.db.Model(&model.CoffeeBean{}).Where("user_id = ?", userID)

	if search != "" {
		query = query.Where("name LIKE ? OR origin LIKE ? OR roaster LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("usage_count DESC, created_at DESC").
		Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit()).
		Find(&beans).Error; err != nil {
		return nil, 0, err
	}

	pagination.Total = total
	return beans, total, nil
}

func (r *CoffeeBeanRepository) FindByName(userID uint, name string) (*model.CoffeeBean, error) {
	var bean model.CoffeeBean
	if err := r.db.Where("user_id = ? AND name = ?", userID, name).First(&bean).Error; err != nil {
		return nil, err
	}
	return &bean, nil
}

func (r *CoffeeBeanRepository) Update(bean *model.CoffeeBean) error {
	return r.db.Model(bean).Where("id = ? AND user_id = ?", bean.ID, bean.UserID).Updates(bean).Error
}

func (r *CoffeeBeanRepository) Delete(id, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&model.CoffeeBean{}).Error
}

func (r *CoffeeBeanRepository) FindBeanNames(userID uint) ([]model.CoffeeBean, error) {
	var beans []model.CoffeeBean
	if err := r.db.Model(&model.CoffeeBean{}).
		Where("user_id = ?", userID).
		Order("usage_count DESC, name ASC").
		Find(&beans).Error; err != nil {
		return nil, err
	}
	return beans, nil
}

func (r *CoffeeBeanRepository) IncrementUsageCount(id, userID uint) error {
	return r.db.Model(&model.CoffeeBean{}).
		Where("id = ? AND user_id = ?", id, userID).
		Update("usage_count", gorm.Expr("usage_count + 1")).Error
}
