package Handlers

import (
	"github.com/labstack/echo"
	"github.com/free-way/riverwaveGateway/utils"
	"context"
	"github.com/free-way/riverwaveCommon/definitions"
	"strconv"
	"errors"
)
var(
	err error
)
func GetAllUsers(ctx echo.Context) error {
	payload := definitions.Empty{}
	res,err := utils.ActorsClient.GetAllUsers(context.Background(),&payload)
	if err != nil{
		goto ErrorOccurred
	}
	return ctx.JSON(200,map[string]interface{}{
		"Message":"Success",
		"Payload": res.Users,
	})
	ErrorOccurred:
		return ctx.JSON(500,map[string]string{"Error": err.Error()})
}

func CreateUser(ctx echo.Context) error {
	payload := definitions.CreateUserRequest{}
	ctx.Bind(&payload)
	res,err := utils.ActorsClient.CreateUser(context.Background(),&payload)
	if err != nil{
		goto ErrorOccurred
	}
	return ctx.JSON(200,map[string]string{"Message":res.Message})
ErrorOccurred:
	return ctx.JSON(500,map[string]string{"Error": err.Error()})
}

func EditUser(ctx echo.Context) error {
	var res *definitions.GeneralResponse
	payload := definitions.EditUserRequest{}
	ctx.Bind(&payload)
	id,err := strconv.Atoi(ctx.Param("id"))
	if err != nil{
		err = errors.New("provide valid integer as id param")
		goto ErrorOccurred
	}
	payload.ID = int32(id)
	res,err = utils.ActorsClient.EditUser(context.Background(),&payload)
	if err != nil{
		goto ErrorOccurred
	}
	return ctx.JSON(200,map[string]string{"Message":res.Message})
ErrorOccurred:
	return ctx.JSON(500,map[string]string{"Error": err.Error()})

}

func DeleteUser(ctx echo.Context) error  {
	var res *definitions.GeneralResponse
	payload := definitions.DeleteUserRequest{}
	id,err := strconv.Atoi(ctx.Param("id"))
	if err != nil{
		err = errors.New("provide valid integer as id param")
		goto ErrorOccurred
	}
	payload.UserID = int32(id)
	res,err = utils.ActorsClient.DeleteUser(context.Background(),&payload)
	if err != nil{
		goto ErrorOccurred
	}
	return ctx.JSON(200,map[string]string{"Message":res.Message})
ErrorOccurred:
	return ctx.JSON(500,map[string]string{"Error": err.Error()})
}

func Auth(ctx echo.Context) error {
	params := definitions.AuthenticationParams{}
	ctx.Bind(&params)
	res,err := utils.ActorsClient.Authenticate(context.Background(),&params)
	if err != nil{
		return ctx.JSON(500,map[string]string{"Error": err.Error()})
	}
	return ctx.JSON(200,map[string]string{"Message":"Success","Payload":res.Message})

}
