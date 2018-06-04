package main

import (
	"github.com/labstack/echo"
	h"github.com/free-way/riverwaveGateway/handlers"

	u"github.com/free-way/riverwaveGateway/utils"
	"google.golang.org/grpc"
	"gopkg.in/ini.v1"
	"flag"
	"fmt"
	"os"
	"github.com/free-way/riverwaveCommon/definitions"
	eMiddleware"github.com/labstack/echo/middleware"
	"github.com/free-way/riverwaveGateway/middleware"

)

var(
	err error
	cfg *ini.File
)

func main()  {
	cfgFlag := flag.String("config","","Env file")
	flag.Parse()
	cfg,err = ini.Load(*cfgFlag)
	if err != nil{
		fmt.Println("could not load configuration file due to: ",err.Error())
		os.Exit(-1)
	}
	u.ActorsConnection,u.Err = grpc.Dial(cfg.Section("Microservices").Key("actorsService").String(),grpc.WithInsecure())
	if u.Err != nil{
		fmt.Println("could not connect to actors service: ",u.Err.Error())
		os.Exit(-1)
	}
	u.ActorsClient = definitions.NewActorsServiceClient(u.ActorsConnection)
	e := echo.New()
	e.Use(eMiddleware.Logger())
	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(200,map[string]string{"Message":"River Wave v1.0"})
	})
	v1 :=e.Group("/api/v1",)
	v1.POST("/auth",h.Auth)
	protected := v1.Group("",middleware.AuthenticationMiddleware)
	//
	protected.GET("/users",h.GetAllUsers)
	protected.POST("/users",h.CreateUser)
	protected.PUT("/users/:id",h.EditUser)
	protected.DELETE("/users/:id",h.DeleteUser)


	e.Logger.Debug(e.Start(":8080"))
}