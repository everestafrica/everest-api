package services

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/repositories"
)

type ReportService interface {
}

type reportService struct {
	cashRepo               repositories.ICashAccountRepository
	accountDetailsService  ICashAccountService
	cashTransactionService ICashTransactionService
	accountTxnRepo         repositories.ICashTransactionRepository
	cryptoRepo             repositories.ICoinRepository
	cryptoService          ICoinService
	cryptoTxnRepo          repositories.ICoinTransactionRepository
}

func NewReportService() ReportService {
	return &reportService{
		cashRepo:               repositories.NewCashRepo(),
		accountDetailsService:  NewCashAccountService(),
		cashTransactionService: NewAccountTransactionService(),
		accountTxnRepo:         repositories.NewCashTransactionRepo(),
		cryptoRepo:             repositories.NewCoinRepo(),
		cryptoService:          NewCoinService(),
		cryptoTxnRepo:          repositories.NewCryptoTransactionRepo(),
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
	acctIncome, err := rs.cashTransactionService.GetInflow(dateRange, userId)
	acctExpense, err := rs.cashTransactionService.GetOutflow(dateRange, userId)
	cryptoIncome, err := rs.cryptoService.GetInflow(dateRange, userId)
	cryptoExpense, err := rs.cryptoService.GetOutflow(dateRange, userId)
	if err != nil {
		return nil, err
	}
	details, err := rs.accountDetailsService.GetAllCashAccountsDetails(userId)
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
