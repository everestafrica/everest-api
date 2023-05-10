package service

type IRatesService interface {
	SetRates() error
	DeleteRates() error
}
