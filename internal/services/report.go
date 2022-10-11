package services

import "github.com/everestafrica/everest-api/internal/repositories"

type ReportService interface {
}

type reportService struct {
	accountRepo    repositories.IAccountDetailsRepository
	accountTxnRepo repositories.IAccountTransactionRepository
	cryptoRepo     repositories.ICryptoDetailsRepository
	cryptoTxnRepo  repositories.ICryptoTransactionRepository
}

func NewReportService() ReportService {
	return &reportService{
		accountRepo:    repositories.NewAccountDetailsRepo(),
		accountTxnRepo: repositories.NewAccountTransactionRepo(),
		cryptoRepo:     repositories.NewCryptoDetailsRepo(),
		cryptoTxnRepo:  repositories.NewCryptoTransactionRepo(),
	}
}

// balance, fees, net income and expenses
