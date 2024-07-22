package booking

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCategoriesHandler(c *gin.Context) {
	categories, err := GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func CreateCategoryHandler(c *gin.Context) {
	var category Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := CreateCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category created successfully"})
}

func GetLocationsHandler(c *gin.Context) {
	locations, err := GetAllLocations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch locations"})
		return
	}
	c.JSON(http.StatusOK, locations)
}

func CreateLocationHandler(c *gin.Context) {
	var location Location
	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := CreateLocation(&location); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create location"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Location created successfully"})
}

func GetSpacesHandler(c *gin.Context) {
	spaces, err := GetAllSpaces()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch spaces"})
		return
	}
	c.JSON(http.StatusOK, spaces)
}

func CreateSpaceHandler(c *gin.Context) {
	var space Space
	if err := c.ShouldBindJSON(&space); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := CreateSpace(&space); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create space"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Space created successfully"})
}

func GetEventsHandler(c *gin.Context) {
	events, err := GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch events"})
		return
	}
	c.JSON(http.StatusOK, events)
}

func CreateEventHandler(c *gin.Context) {
	var event Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := CreateEvent(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event created successfully"})
}

func GetEventTimesHandler(c *gin.Context) {
	eventID, err := strconv.ParseUint(c.Param("event_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	eventTimes, err := GetEventTimesByEventID(uint(eventID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event times"})
		return
	}
	c.JSON(http.StatusOK, eventTimes)
}

func CreateEventTimeHandler(c *gin.Context) {
	var eventTime EventTime
	if err := c.ShouldBindJSON(&eventTime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := CreateEventTime(&eventTime); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create event time"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event time created successfully"})
}

func CreateReservationHandler(c *gin.Context) {
	var reservation Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := CreateReservation(&reservation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create reservation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation created successfully"})
}
