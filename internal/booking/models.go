package booking

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique"`
	Description string `json:"description"`
}

type Location struct {
	gorm.Model
	Name       string `json:"name"`
	CategoryID uint   `json:"category_id"`
	Address    string `json:"address"`
}

type Space struct {
	gorm.Model
	LocationID uint   `json:"location_id"`
	Name       string `json:"name"`
	Description string `json:"description"`
	Capacity   int    `json:"capacity"`
}

type Event struct {
	gorm.Model
	SpaceID     uint   `json:"space_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EventTime struct {
	gorm.Model
	EventID   uint      `json:"event_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type Reservation struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	EventTimeID uint   `json:"event_time_id"`
	Status      string `json:"status"`
}
