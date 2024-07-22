package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wasupalonely/reservapp/internal/auth"
	"github.com/wasupalonely/reservapp/internal/booking"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/login", auth.LoginHandler)
		api.POST("/register", auth.RegisterHandler)

		authGroup := api.Group("/")
		authGroup.Use(auth.AuthMiddleware())
		{
			authGroup.GET("/categories", booking.GetCategoriesHandler)
			authGroup.POST("/categories", booking.CreateCategoryHandler)
			authGroup.GET("/locations", booking.GetLocationsHandler)
			authGroup.POST("/locations", booking.CreateLocationHandler)
			authGroup.GET("/spaces", booking.GetSpacesHandler)
			authGroup.POST("/spaces", booking.CreateSpaceHandler)
			authGroup.GET("/events", booking.GetEventsHandler)
			authGroup.POST("/events", booking.CreateEventHandler)
			authGroup.GET("/event_times/:event_id", booking.GetEventTimesHandler)
			authGroup.POST("/event_times", booking.CreateEventTimeHandler)
			authGroup.POST("/reservations", booking.CreateReservationHandler)
		}
	}

	return r
}
