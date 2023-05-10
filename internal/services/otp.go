package services

import (
	"errors"
	"fmt"
	"github.com/everestafrica/everest-api/internal/config"
	"github.com/everestafrica/everest-api/internal/database/redis"
	"math/rand"
	"time"
)

const defaultOTPCode = "111111"

var otpCodeTime = 5

type IOTPService interface {
	Generate(key string) (*string, error)
	Validate(key string, code string) error
}

type otpService struct {
	cache *redis.Client
	cfg   *config.Config
}

// NewOTPService will instantiate AuthService
func NewOTPService() IOTPService {
	return &otpService{
		cache: redis.RedisClient(),
		cfg:   config.GetConf(),
	}
}

func (otp *otpService) Generate(key string) (*string, error) {

	var otpCode = ""

	if otp.cfg.Env != "production" {
		otpCode = defaultOTPCode
	} else {
		otpCode = randomCode()
	}

	err := otp.cache.Set(key, otpCode, time.Duration(otpCodeTime)*time.Minute)
	if err != nil {
		return nil, err
	}

	return &otpCode, nil

}

func randomCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var code [6]byte
	for i := 0; i < 6; i++ {
		code[i] = uint8(48 + r.Intn(10))
	}

	otpCode := string(code[:])
	return otpCode
}

func (otp *otpService) Validate(key string, code string) error {

	val, err := otp.cache.Get(key)

	if err != nil {
		return errors.New("invalid code")
	}
	if code != fmt.Sprintf("%v", val) {
		return errors.New("invalid code")
	}

	return nil

}
