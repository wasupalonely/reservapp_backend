package booking

import "github.com/wasupalonely/reservapp/pkg/db"

func GetAllCategories() ([]Category, error) {
	var categories []Category
	if err := db.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func CreateCategory(category *Category) error {
	return db.DB.Create(category).Error
}

func GetAllLocations() ([]Location, error) {
	var locations []Location
	if err := db.DB.Find(&locations).Error; err != nil {
		return nil, err
	}
	return locations, nil
}

func CreateLocation(location *Location) error {
	return db.DB.Create(location).Error
}

func GetAllSpaces() ([]Space, error) {
	var spaces []Space
	if err := db.DB.Find(&spaces).Error; err != nil {
		return nil, err
	}
	return spaces, nil
}

func CreateSpace(space *Space) error {
	return db.DB.Create(space).Error
}

func GetAllEvents() ([]Event, error) {
	var events []Event
	if err := db.DB.Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func CreateEvent(event *Event) error {
	return db.DB.Create(event).Error
}

func GetEventTimesByEventID(eventID uint) ([]EventTime, error) {
	var eventTimes []EventTime
	if err := db.DB.Where("event_id = ?", eventID).Find(&eventTimes).Error; err != nil {
		return nil, err
	}
	return eventTimes, nil
}

func CreateEventTime(eventTime *EventTime) error {
	return db.DB.Create(eventTime).Error
}

func CreateReservation(reservation *Reservation) error {
	return db.DB.Create(reservation).Error
}
