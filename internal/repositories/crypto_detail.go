package repositories

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type ICryptoDetailsRepository interface {
	Create(account *models.CryptoDetail) error
	Update(account *models.CryptoDetail) error
	FindByUserId(userId string) (*models.CryptoDetail, error)
}

type cryptoDetailsRepo struct {
	db *gorm.DB
}

// NewCryptoDetailsRepo will instantiate AccountDetails Repository
func NewCryptoDetailsRepo() ICryptoDetailsRepository {
	return &cryptoDetailsRepo{
		db: database.DB(),
	}
}

func (r *cryptoDetailsRepo) Create(account *models.CryptoDetail) error {
	return r.db.Create(account).Error
}

func (r *cryptoDetailsRepo) Update(account *models.CryptoDetail) error {
	return r.db.Save(account).Error
}

func (r *cryptoDetailsRepo) FindByUserId(userId string) (*models.CryptoDetail, error) {
	var crypto models.CryptoDetail
	if err := r.db.Where("user_id = ?", userId).First(&crypto).Error; err != nil {
		return nil, err
	}

	return &crypto, nil
}
