package handlers

import (
	"github.com/labstack/echo"
	"github.com/free-way/riverwaveGateway/utils"
	"context"
	"github.com/free-way/riverwaveCommon/definitions"
	"strconv"
)

func GetAllResources(ctx echo.Context) error {
	resources,err := utils.ResourcesClient.GetResources(context.Background(),&definitions.Empty{})
	if err != nil{
		return ctx.JSON(500,map[string]interface{}{
			"Message":"Internal Server Error",
			"Original Message": err.Error(),
		})
	}
	return ctx.JSON(200,map[string]interface{}{
		"Message":"Success",
		"Payload":resources,
	})
}

func CreateResource(ctx echo.Context)error{
	resource := definitions.CreateResourceParams{}
	ctx.Bind(&resource)
	res,err := utils.ResourcesClient.AddResource(context.Background(),&resource)
	if err != nil{
		return ctx.JSON(500,map[string]interface{}{
			"Message":"Internal Server Error",
			"Original Message": err.Error(),
		})
	}
	return ctx.JSON(200,map[string]interface{}{
		"Message":res.Message,
	})
}

func EditResource(ctx echo.Context) error{
	resource := definitions.EditResourceParams{}
	id,err := strconv.Atoi(ctx.Param("id"))
	if err != nil{
		return ctx.JSON(500,map[string]interface{}{
			"Message":"Internal Server Error",
			"Original Message": err.Error(),
		})
	}

	resource.ResourceId = int32(id)
	if err = ctx.Bind(&resource); err != nil{
		return ctx.JSON(500,map[string]interface{}{
			"Message":"Internal Server Error",
			"Original Message": err.Error(),
		})
	}
	res,err := utils.ResourcesClient.EditResource(context.Background(),&resource)
	if err != nil{
		return ctx.JSON(500,map[string]interface{}{
			"Message":"Internal Server Error",
			"Original Message": err.Error(),
		})
	}

	return ctx.JSON(200,map[string]interface{}{
		"Message":res.Message,
	})

}

func DeleteResource(ctx echo.Context) error  {
	resource := definitions.DeleteResourceParams{}
	id,err := strconv.Atoi(ctx.Param("id"))
	if err != nil{
		return ctx.JSON(500,map[string]interface{}{
			"Message":"Internal Server Error",
			"Original Message": err.Error(),
		})
	}

	resource.ResourceId = int32(id)
	res,err := utils.ResourcesClient.DeleteResource(context.Background(),&resource)
	if err != nil{
		return ctx.JSON(500,map[string]interface{}{
			"Message":"Internal Server Error",
			"Original Message": err.Error(),
		})
	}
	return ctx.JSON(200,map[string]interface{}{
		"Message":res.Message,
	})
}
