package route

import (
	"github.com/gin-gonic/gin"
	"ilham/config"
	"ilham/controller"
)

func User()  {
	db := config.DBInit()
	inDB := &controller.InDB{DB: db}
	route := gin.Default();

	route.GET("/", controller.Index)
	route.GET("/user/:id", inDB.GetUser)
	route.GET("/users", inDB.GetUsers)
	route.POST("/user", inDB.CreateUser)
	route.PUT("/user", inDB.UpdateUser)
	route.DELETE("/user/:id", inDB.DeleteUser)
	route.Run();
}
