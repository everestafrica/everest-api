package handlers

import (
	"encoding/json"
	"github.com/everestafrica/everest-api/internal/commons/log"
	"github.com/everestafrica/everest-api/internal/commons/types"
	util "github.com/everestafrica/everest-api/internal/commons/util"
	"github.com/everestafrica/everest-api/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

const (
	requestIdentifier         = "REQ_ID"
	AuthUserContextKey string = "auth"
)

var utilToken = util.Token{}

// UserFromContext extracts the user_id from context
func UserFromContext(ctx *fiber.Ctx) (string, error) {
	id := ctx.GetRespHeader(AuthUserContextKey)
	if &id == nil {
		return "", ctx.Status(fiber.StatusUnauthorized).JSON("unable to fetch user info from token")
	}

	return id, nil
}

// StructToMap converts a struct of any type to a map[string]interface{}.
func StructToMap(inStruct interface{}) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	res, _ := json.Marshal(inStruct)

	if err := json.Unmarshal(res, &out); err != nil {
		return nil, err
	}

	return out, nil
}

// SecureAuth returns a middleware which secures all the private routes
func SecureAuth() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		logger := log.WithFields(log.FromContext(c.Context()).Fields)
		identifier := uuid.NewString()
		logger = logger.WithField(requestIdentifier, identifier)
		logger.Info("SecureAuth middleware")
		jwtKey := config.GetConf().JWTSecret

		accessToken, err := utilToken.ExtractBearerToken(c.Get("Authorization"))
		if err != nil {
			logger.Error("error while extracting token: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(types.GenericResponse{
				Success: false,
				Message: err.Error(),
			})
		}
		claims := new(types.Claims)
		token, err := jwt.ParseWithClaims(accessToken, claims,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtKey), nil
			})

		if err != nil {
			logger.Error("error while parsing claims: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(types.GenericResponse{
				Success: false,
				Message: err.Error(),
			})
		}

		if token.Valid {
			if claims.ExpiresAt < time.Now().Unix() {
				return c.Status(fiber.StatusUnauthorized).JSON(types.GenericResponse{
					Success: false,
					Message: "token expired",
				})
			}
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return c.SendStatus(fiber.StatusForbidden)
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return c.Status(fiber.StatusUnauthorized).JSON(types.GenericResponse{
					Success: false,
					Message: "token expired or not yet active",
				})
			} else {
				// cannot handle this token
				return c.Status(fiber.StatusForbidden).JSON(types.GenericResponse{
					Success: false,
					Message: "unable to handle this token or invalid token",
				})
			}
		}

		c.Set(AuthUserContextKey, claims.ID)
		return c.Next()
	}
}
