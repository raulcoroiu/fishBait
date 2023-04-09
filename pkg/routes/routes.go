package routes


import (
	"fmt"
	"github.com/raulcoroiu/fishBait/pkg/controllers"
	"github.com/gorilla/mux"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/users/singup", controllers.Signup())
	incomingRoutes.POST("/usres/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewAdmin())
	incomingRoutes.POST("/users/prodctview", controllers.SearchProduct())
	incomingRoutes.POST("/uses/search", controllers.SearchProductByQuery())
}

