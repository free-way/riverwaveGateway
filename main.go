package main

import (
	"github.com/labstack/echo"
	h"github.com/free-way/riverwaveGateway/Handlers"

	u"github.com/free-way/riverwaveGateway/utils"
	"google.golang.org/grpc"
	"gopkg.in/ini.v1"
	"flag"
	"fmt"
	"os"
	"github.com/free-way/riverwaveCommon/definitions"
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
	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(200,map[string]string{"Message":"River Wave v1.0"})
	})
	v1 :=e.Group("/api/v1",)
	v1.GET("/users",h.GetAllUsers)
	v1.POST("/users",h.CreateUser)
	v1.PUT("/users/:id",h.EditUser)
	v1.DELETE("/users/:id",h.DeleteUser)

	v1.POST("/auth",h.Auth)
	e.Logger.Fatal(e.Start(":8080"))
}