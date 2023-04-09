package main

import (
	"github.com/raulcoroiu/fishBait/pkg/controllers"
	"github.com/raulcoroiu/fishBait/pkg/database"
	"github.com/raulcoroiu/fishBait/pkg/middleware"
	"github.com/raulcoroiu/fishBait/pkg/routes"
	"github.com/gin-gonic/gin"
)


func main(){
	//de citi os package
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}

	app:= controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	
	//de citit gin package
	router := gin.New()
	router.User(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authenticator())

	router.GET("/addtocart", app.AddToCard())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))


}