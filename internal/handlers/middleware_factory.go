package handlers

//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/gofiber/fiber/v2"
//	repository "github.com/wirepay/wirepay-api/internal/repositories"
//)
//
//type middleware struct {
//	Jwt	*fiber.Ctx
//}
//
//var registeredMiddleWare middleware
//
//func InitializeMiddleWares()  {
//
//	var (
//		userRepo 	= repository.NewUserRepo()
//	)
//
//	registeredMiddleWare = middleware{
//		Jwt: AuthJWT(userRepo),
//	}
//}
//
//func MiddleWare() *middleware {
//	return &registeredMiddleWare
//}
