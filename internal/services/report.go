package services

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/repositories"
)

type ReportService interface {
}

type reportService struct {
	accountRepo               repositories.IAccountDetailsRepository
	accountDetailsService     IAccountDetailsService
	accountTransactionService IAccountTransactionService
	accountTxnRepo            repositories.IAccountTransactionRepository
	cryptoRepo                repositories.ICryptoDetailsRepository
	cryptoService             ICryptoService
	cryptoTxnRepo             repositories.ICryptoTransactionRepository
}

func NewReportService() ReportService {
	return &reportService{
		accountRepo:               repositories.NewAccountDetailsRepo(),
		accountDetailsService:     NewAccountDetailsService(),
		accountTransactionService: NewAccountTransactionService(),
		accountTxnRepo:            repositories.NewAccountTransactionRepo(),
		cryptoRepo:                repositories.NewCryptoDetailsRepo(),
		cryptoService:             NewCryptoService(),
		cryptoTxnRepo:             repositories.NewCryptoTransactionRepo(),
	}
}

type ReportResponse struct {
	DateRange      types.DateRange
	Networth       float64
	TotalIncome    float64
	TotalExpense   float64
	AccountIncome  float64
	CryptoIncome   float64
	AccountExpense float64
	CryptoExpense  float64
}

func (rs reportService) GetAccountDetailsReport(userId string, dateRange types.DateRange) (*ReportResponse, error) {
	acctIncome, err := rs.accountTransactionService.GetInflow(dateRange, userId)
	acctExpense, err := rs.accountTransactionService.GetOutflow(dateRange, userId)
	cryptoIncome, err := rs.cryptoService.GetInflow(dateRange, userId)
	cryptoExpense, err := rs.cryptoService.GetOutflow(dateRange, userId)
	if err != nil {
		return nil, err
	}
	details, err := rs.accountDetailsService.GetAllAccountsDetails(userId)
	if err != nil {
		return nil, err
	}
	var networth float64
	for _, detail := range *details {
		networth += float64(detail.Balance)
	}
	return &ReportResponse{
		DateRange:      dateRange,
		Networth:       networth,
		TotalIncome:    acctIncome.Total + cryptoIncome.Total,
		TotalExpense:   acctExpense.Total + cryptoExpense.Total,
		AccountIncome:  acctIncome.Total,
		CryptoIncome:   cryptoIncome.Total,
		AccountExpense: acctExpense.Total,
		CryptoExpense:  cryptoExpense.Total,
	}, nil
}

// balance, fees, net income and expenses

// Week summary, Month summary, Year summary
