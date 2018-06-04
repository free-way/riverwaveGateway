package middleware

import (
	"github.com/labstack/echo"
	"github.com/free-way/riverwaveGateway/utils"
	"github.com/free-way/riverwaveCommon/definitions"
	ctx"context"
)

func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		token := context.Request().Header.Get("Authorization")
		res,err := utils.ActorsClient.ValidateToken(ctx.Background(),&definitions.LoginParams{
			Token: token,
		})

		if err != nil || res.IsValid == false{
			return context.JSON(401,map[string]interface{}{
				"Message": "Authentication Failed",
				"Origin Error": err.Error(),

			})
		}
		return next(context)
	}
}
