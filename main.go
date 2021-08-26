package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/muhammadisa/nobita/endpoint"
	authrepositoryv1 "github.com/muhammadisa/nobita/repository/v1/auth"
	featurerepositoryv1 "github.com/muhammadisa/nobita/repository/v1/feature"
	redisrepositoryv1 "github.com/muhammadisa/nobita/repository/v1/redis"
	routeversion "github.com/muhammadisa/nobita/route"
	authroutev1 "github.com/muhammadisa/nobita/route/v1/auth"
	fooroutev1 "github.com/muhammadisa/nobita/route/v1/foo"

	authusecasev1 "github.com/muhammadisa/nobita/usecase/v1/auth"
	"github.com/muhammadisa/nobita/util/dbc"
	"github.com/muhammadisa/nobita/util/otp"
	"github.com/muhammadisa/nobita/util/vlt"
)

func main() {
	router := gin.Default()

	vault, err := vlt.NewVLT("myroot", "http://0.0.0.0:8300", "/secret")
	if err != nil {
		panic(err)
	}

	var senderConfig otp.Config
	{
		senderConfig.Email = vault.Get("/email_otp:email")
		senderConfig.Password = vault.Get("/email_otp:password")
		senderConfig.TemplatePath = vault.Get("/email_otp:html_template")
	}

	var redisConfig dbc.Config
	{
		redisConfig.Host = vault.Get("/redis_database:host")
		redisConfig.Port = vault.Get("/redis_database:port")
		redisConfig.Password = vault.Get("/redis_database:password")
	}

	var mysqlConfig dbc.Config
	{
		mysqlConfig.Username = vault.Get("/sql_database:username")
		mysqlConfig.Password = vault.Get("/sql_database:password")
		mysqlConfig.Host = vault.Get("/sql_database:host")
		mysqlConfig.Port = vault.Get("/sql_database:port")
		mysqlConfig.Name = vault.Get("/sql_database:db")
	}

	fmt.Println("redis config", redisConfig)
	fmt.Println("mysql config", mysqlConfig)
	fmt.Println("sender config", senderConfig)

	_, err = redisrepositoryv1.NewRedisRepo(redisConfig)
	if err != nil {
		panic(err)
	}

	authRepo, err := authrepositoryv1.NewAuthRepo(mysqlConfig)
	if err != nil {
		panic(err)
	}

	featureRepo, err := featurerepositoryv1.NewFeatureRepo(mysqlConfig)
	if err != nil {
		panic(err)
	}

	endpoint.NewEndpoint(router, featureRepo, routeversion.Versions{
		RouteAuthV1: authroutev1.NewAuthRouteV1(authusecasev1.NewAuthUseCaseV1(authRepo, senderConfig)),
		RouteFooV1:  fooroutev1.NewFooRouteV1(),
	})

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
