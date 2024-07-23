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
			categoriesGroup := authGroup.Group("/categories")
			{
				categoriesGroup.GET("/", booking.GetCategoriesHandler)
				categoriesGroup.POST("/", booking.CreateCategoryHandler)
			}

			locationsGroup := authGroup.Group("/locations")
			{
				locationsGroup.GET("/", booking.GetLocationsHandler)
				locationsGroup.POST("/", booking.CreateLocationHandler)
			}

			spacesGroup := authGroup.Group("/spaces")
			{
				spacesGroup.GET("/", booking.GetSpacesHandler)
				spacesGroup.POST("/", booking.CreateSpaceHandler)
			}

			eventsGroup := authGroup.Group("/events")
			{
				eventsGroup.GET("/", booking.GetEventsHandler)
				eventsGroup.POST("/", booking.CreateEventHandler)
			}

			reservationsGroup := authGroup.Group("/reservations")
			{
				reservationsGroup.POST("/", booking.CreateReservationHandler)

			}
			eventTimesGroup := authGroup.Group("/event_times")
			{
				eventTimesGroup.GET("/:event_id", booking.GetEventTimesHandler)
				eventTimesGroup.POST("/", booking.CreateEventTimeHandler)
			}

		}
	}

	return r
}
