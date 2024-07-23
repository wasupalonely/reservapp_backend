package validation

type CreateCategoryData struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type CreateLocationData struct {
	Name       string `json:"name" binding:"required"`
	CategoryID uint   `json:"category_id" binding:"required"`
	Address    string `json:"address" binding:"required"`
}

type CreateSpaceData struct {
	LocationID  uint   `json:"location_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Capacity    int    `json:"capacity" binding:"required"`
}

type CreateEventData struct {
	SpaceID     uint   `json:"space_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	StartDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date" binding:"required"`
}

type CreateReservationData struct {
	UserID      uint   `json:"user_id" binding:"required"`
	EventTimeID uint   `json:"event_time_id" binding:"required"`
	Status      string `json:"status" binding:"required"`
}
