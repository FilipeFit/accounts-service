package repositories

import (
	"github.com/filipeFit/account-service/domain"
	"github.com/filipeFit/account-service/domain/api"
	"gorm.io/gorm"
)

type accountRepository struct {
	DB *gorm.DB
}

type accountRepositoryInterface interface {
	Create(request *api.CreateAccountRequest) (*domain.Account, error)
	FindAll() ([]domain.Account, error)
	FindByID(id uint64) (*domain.Account, error)
	FindByCustomerID(customerId uint64) (*domain.Account, error)
	UpdateBalance(accountID uint64, balance float64) (float64, error)
}

var (
	_ accountRepositoryInterface
)

func init() {
	_ = &accountRepository{}
}

func NewAccountRepository(db *gorm.DB) *accountRepository {
	return &accountRepository{DB: db}
}

func (s *accountRepository) Create(request *api.CreateAccountRequest) (*domain.Account, error) {
	account := domain.Account{Active: true,
		AllowOverdraft: request.AllowOverdraft,
		CustomerId:     request.CustomerId,
		Balance:        0}
	result := s.DB.Create(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

func (s *accountRepository) FindAll() ([]domain.Account, error) {
	var accounts []domain.Account
	result := s.DB.Find(&accounts)
	if result.Error != nil {
		return nil, result.Error
	}
	return accounts, nil
}

func (s *accountRepository) FindByID(id uint64) (*domain.Account, error) {
	var account domain.Account
	result := s.DB.First(&account, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

func (s *accountRepository) FindByCustomerID(customerId uint64) (*domain.Account, error) {
	var account domain.Account
	result := s.DB.First(&account, "customer_id = ?", customerId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

func (s *accountRepository) UpdateBalance(accountID uint64, balance float64) (float64, error) {
	account := &domain.Account{}
	result := s.DB.Model(account).Where("id = ?", accountID).Update("balance", balance)
	if result.Error != nil {
		return 0, result.Error
	}
	return account.Balance, nil
}
