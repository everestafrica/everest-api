package repositories

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type ICryptoDetailsRepository interface {
	Create(crypto *models.CryptoDetail) error
	Update(crypto *models.CryptoDetail) error
	FindByUserId(userId string) (*models.CryptoDetail, error)
}

type cryptoDetailsRepo struct {
	db *gorm.DB
}

// NewCryptoDetailsRepo will instantiate CryptoDetails Repository
func NewCryptoDetailsRepo() ICryptoDetailsRepository {
	return &cryptoDetailsRepo{
		db: database.DB(),
	}
}

func (r *cryptoDetailsRepo) Create(crypto *models.CryptoDetail) error {
	return r.db.Create(crypto).Error
}

func (r *cryptoDetailsRepo) Update(crypto *models.CryptoDetail) error {
	return r.db.Save(crypto).Error
}

func (r *cryptoDetailsRepo) FindByUserId(userId string) (*models.CryptoDetail, error) {
	var crypto models.CryptoDetail
	if err := r.db.Where("user_id = ?", userId).First(&crypto).Error; err != nil {
		return nil, err
	}

	return &crypto, nil
}
