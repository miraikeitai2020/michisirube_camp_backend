package main

import(
	// import gin library
	"github.com/gin-gonic/gin"

	// import sample API packages
	"github.com/miraikeitai2020/michisirube_camp_backend/evaluation/pkg/server/controller"
)

func initializeController() controller.Controller {
	return controller.Controller{}
}

func setupRooter(ctrl controller.Controller) *gin.Engine {
	router := gin.Default()
	// API Handlers List
	// router.POST("/", )
	return router
}

func main() {
	ctrl := initializeController()
	router := setupRooter(ctrl)
	err := router.Run(":9030")
	if err != nil {
		panic("server Painc")
	}
}