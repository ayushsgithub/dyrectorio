package route

import (
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.com/dyrector_io/dyrector.io/go/pkg/crane/controller"
	"gitlab.com/dyrector_io/dyrector.io/go/pkg/crane/docs"

	"github.com/gin-gonic/gin"
)

// SetupRouterV1 contains all types of routes
// it is mixed with original / new API calls
// todo: V2 http api
func SetupRouterV1(r *gin.Engine) *gin.Engine {
	support := r.Group("api")
	support.GET("/containers", controller.GetDeployments)
	support.GET("/Status/GetStatus", controller.GetDeploymentStatus)
	support.POST("/Distribution/DeployImage", controller.DeployImage)

	// DeployRequest -> facade -> configmaps + services + deployments + ingressobjects
	// deployments -> facade ->

	api := r.Group("v1")
	api.GET("/deployments", controller.GetDeployments)
	api.GET("/namespaces", controller.GetNamespaces)

	// based-off of dagents interface, cries for a refactor
	// two different domains to be synced into one API
	api.GET("/containers/:preName/:name/status", controller.GetDeploymentStatus)
	api.GET("/containers/:preName/:name/logs", controller.GetDeploymentLogs)
	api.GET("/containers/:preName/:name/inspect", controller.DeleteDeployment)
	api.GET("/containers", controller.GetDeployments)
	api.DELETE("/containers/:preName/:name", controller.DeleteDeployment)

	api.POST("/deploy", controller.DeployImage)
	api.POST("/deploy/batch", controller.BatchDeployImage)
	api.POST("/deploy/version", controller.DeployVersion)

	api.GET("/swagger", swaggerRedirect)
	r.GET("swagger", swaggerRedirect)

	docs.SwaggerInfo.BasePath = "/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func swaggerRedirect(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
}

// SetupUpdate is for application image update weebhooks
func SetupUpdate(r *gin.Engine) *gin.Engine {
	r.POST("update", controller.UpdateRunningCrane)
	return r
}
