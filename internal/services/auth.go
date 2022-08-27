package services

import (
	"errors"
	"github.com/everestafrica/everest-api/internal/commons/types"
	util "github.com/everestafrica/everest-api/internal/commons/utils"
	"github.com/everestafrica/everest-api/internal/config"
	"github.com/everestafrica/everest-api/internal/models"
	"github.com/everestafrica/everest-api/internal/repositories"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type IAuthService interface {
	Register(body types.RegisterRequest) (*types.RegisterResponse, error)
	//SendOTPCode(request *types.SendCodeRequest) error
	Login(body types.LoginRequest) (*types.LoginResponse, error)
	IssueToken(u *models.User) (*types.TokenResponse, error)
	ParseToken(token string) (*types.Claims, error)
	RefreshToken(token string) (*types.TokenResponse, error)
}

type authService struct {
	jwtSecret  string
	userRepo   repositories.IUserRepository
	otpService IOTPService
}

// NewAuthService will instantiate AuthService
func NewAuthService() IAuthService {
	return &authService{
		jwtSecret:  config.GetConf().JWTSecret,
		userRepo:   repositories.NewUserRepo(),
		otpService: NewOTPService(),
	}
}

//func (as *authService) SendOTPCode(request *types.SendCodeRequest) error {
//
//	code, err := as.otpService.Generate(request.Receiver)
//
//	if err != nil {
//		return errors.New("oops an error occurred please try again")
//	}
//
//	message := fmt.Sprintf("Your wirepay code is %s", *code)
//
//	if !request.IsEmail {
//		go rehook.SendSMS(message, request.Receiver)
//	} else {
//		go sendgrid.SendEmail(sendgrid.Email{
//			ToName:  "",
//			ToEmail: request.Receiver,
//			Subject: "Your Wirepay code",
//			HTML:    fmt.Sprintf("<html><body>%s</body></html>", message),
//			Text:    message,
//		})
//	}
//
//	return nil
//
//}

func (as *authService) Register(body types.RegisterRequest) (*types.RegisterResponse, error) {
	stringUtil := util.StringUtil{}

	body.FirstName = stringUtil.CapitalizeFirstCharacter(body.FirstName)
	body.LastName = stringUtil.CapitalizeFirstCharacter(body.LastName)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("could not set password")
	}
	password := string(hashedPassword)
	user := models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  password,
	}
	err = as.userRepo.Create(&user)
	if err != nil {
		return nil, err
	}

	return &types.RegisterResponse{
		FirstName: body.FirstName,
		LastName:  body.LastName,
	}, nil
}

func (as *authService) Login(body types.LoginRequest) (*types.LoginResponse, error) {
	user, err := as.userRepo.FindByEmail(body.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	issueResponse, err := as.IssueToken(user)

	return &types.LoginResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Token:     issueResponse,
	}, nil
}

func (as *authService) IssueToken(u *models.User) (*types.TokenResponse, error) {
	nowTime := time.Now()
	expConf, err := strconv.Atoi(config.GetConf().ExpiryTime)

	if err != nil {
		return nil, err
	}
	expireTime := nowTime.Add(time.Duration(int64(expConf)) * time.Minute)

	claims := types.Claims{
		Email: u.Email,
		ID:    u.UserId,
		StandardClaims: jwt.StandardClaims{
			Subject:   u.UserId,
			IssuedAt:  nowTime.Unix(),
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Everest",
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := jwtToken.SignedString([]byte(as.jwtSecret))

	if err != nil {
		return nil, err
	}

	return &types.TokenResponse{
		AccessToken: accessToken,
		ExpiresAt:   claims.ExpiresAt,
		Issuer:      claims.Issuer,
	}, nil
}

func (as *authService) ParseToken(token string) (*types.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&types.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(as.jwtSecret), nil
		},
	)

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*types.Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func (as *authService) RefreshToken(token string) (*types.TokenResponse, error) {

	claims, err := as.ParseToken(token)

	if err != nil {
		return nil, err
	}

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 60 seconds of expiry. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 60*time.Second {
		return &types.TokenResponse{
			AccessToken: token,
			ExpiresAt:   claims.ExpiresAt,
			Issuer:      claims.Issuer,
		}, nil
	}

	// get configure exp time
	expConf, err := strconv.Atoi(config.GetConf().ExpiryTime)

	if err != nil {
		return nil, err
	}
	expiry := time.Duration(expConf) * time.Minute

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(expiry)
	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().Unix()

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = newToken.SignedString([]byte(as.jwtSecret))

	if err != nil {
		return nil, err
	}

	return &types.TokenResponse{
		AccessToken: token,
		ExpiresAt:   claims.ExpiresAt,
		Issuer:      claims.Issuer,
	}, nil
}

//func (as *authService) SendOTPCode(request *types.SendCodeRequest) error {
//
//	code, err := as.otpService.Generate(request.Receiver)
//
//	if err != nil {
//		return errors.New("oops an error occurred please try again")
//	}
//
//	message := fmt.Sprintf("Your Everest code is %s", *code)
//
//	if !request.IsEmail {
//		//go channels.SendSMS(message, request.Receiver)
//	} else {
//		go channels.SendMail(&channels.Email{
//			Sender:    "Everest",
//			Subject:   "OTP",
//			Body:      message,
//			Recipient: request.Receiver,
//		})
//	}
//
//	return nil
//
//}
