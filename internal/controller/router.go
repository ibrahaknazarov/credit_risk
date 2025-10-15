package controller

import (
	"fmt"
	// _ "github.com/credit_risk/api/docs"
	"github.com/ibrahaknzarov/credit_risk/internal/configs"
	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func (ctrl *Controller) InitRoutes() error {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", ctrl.ping)

	authG := r.Group("/auth")
	{
		authG.POST("/sign-up", ctrl.SignUp)
		authG.POST("/sign-in", ctrl.SignIn)
		authG.GET("/refresh", ctrl.RefreshTokenPair)
	}

	apiV1G := r.Group("/api/v1", ctrl.checkUserAuthentication)
	{
		apiV1G.GET("/profile", ctrl.GetProfile)
		apiV1G.PUT("/profile", ctrl.UpdateProfile)
		apiV1G.POST("/profile/avatar", ctrl.UploadAvatar)
		apiV1G.GET("/profile/avatar", ctrl.GetAvatar)
		apiV1G.DELETE("/profile/avatar", ctrl.DeleteAvatar)

		apiV1G.POST("/credit-applications", ctrl.CreateCreditApplication)	
		apiV1G.GET("/credit-applications", ctrl.ListCreditApplications)
		apiV1G.GET("/credit-applications/:id", ctrl.GetCreditApplicationByID)
		apiV1G.PUT("/credit-applications/:id", ctrl.UpdateCreditApplicationByID)
		apiV1G.DELETE("/credit-applications/:id", ctrl.DeleteCreditApplicationByID)
	}

	err := r.Run(fmt.Sprintf(":%s", configs.AppSettings.AppParams.PortRun))
	if err != nil {
		return err
	}

	return nil
}

// Ping
// @Summary     Healthcheck
// @Description Роут проверки сервиса
// @Tags        Ping
// @Produce     json
// @Success     200 {object} CommonResponse
// @Failure     500 {object} CommonError
// @Router      /ping [get]
func (ctrl *Controller) ping(c *gin.Context) {
	c.JSON(http.StatusOK, CommonResponse{
		Message: "Server is running",
	})
}
