package services

type IRatesService interface {
	SetRates() error
	DeleteRates() error
}
