package repository

import (
	"my-coffee-log/internal/model"
	"my-coffee-log/internal/utils"

	"gorm.io/gorm"
)

type CoffeeShopRepository struct {
	db *gorm.DB
}

func NewCoffeeShopRepository(db *gorm.DB) *CoffeeShopRepository {
	return &CoffeeShopRepository{db: db}
}

func (r *CoffeeShopRepository) Create(shop *model.CoffeeShop) error {
	return r.db.Create(shop).Error
}

func (r *CoffeeShopRepository) FindByID(id, userID uint) (*model.CoffeeShop, error) {
	var shop model.CoffeeShop
	if err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&shop).Error; err != nil {
		return nil, err
	}
	return &shop, nil
}

func (r *CoffeeShopRepository) FindList(userID uint, pagination *utils.Pagination, search string) ([]model.CoffeeShop, int64, error) {
	var shops []model.CoffeeShop
	var total int64

	query := r.db.Model(&model.CoffeeShop{}).Where("user_id = ?", userID)

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("visit_count DESC, last_visit_at DESC, created_at DESC").
		Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit()).
		Find(&shops).Error; err != nil {
		return nil, 0, err
	}

	pagination.Total = total
	return shops, total, nil
}

func (r *CoffeeShopRepository) FindByName(userID uint, name string) (*model.CoffeeShop, error) {
	var shop model.CoffeeShop
	if err := r.db.Where("user_id = ? AND name = ?", userID, name).First(&shop).Error; err != nil {
		return nil, err
	}
	return &shop, nil
}

func (r *CoffeeShopRepository) Update(shop *model.CoffeeShop) error {
	return r.db.Model(shop).Where("id = ? AND user_id = ?", shop.ID, shop.UserID).Updates(shop).Error
}

func (r *CoffeeShopRepository) Delete(id, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&model.CoffeeShop{}).Error
}

func (r *CoffeeShopRepository) FindShopNames(userID uint) ([]string, error) {
	var names []string
	if err := r.db.Model(&model.CoffeeShop{}).
		Where("user_id = ?", userID).
		Order("visit_count DESC, name ASC").
		Pluck("name", &names).Error; err != nil {
		return nil, err
	}
	return names, nil
}

func (r *CoffeeShopRepository) IncrementVisitCount(id, userID uint) error {
	return r.db.Model(&model.CoffeeShop{}).
		Where("id = ? AND user_id = ?", id, userID).
		Updates(map[string]interface{}{
			"visit_count":   gorm.Expr("visit_count + 1"),
			"last_visit_at": gorm.Expr("CURRENT_DATE"),
		}).Error
}
