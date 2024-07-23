package main

import (
	"log"
	"os/user"

	"github.com/wasupalonely/reservapp/config"
	"github.com/wasupalonely/reservapp/internal/booking"
	"github.com/wasupalonely/reservapp/internal/router"
	"github.com/wasupalonely/reservapp/pkg/db"
)

func main() {
	config.InitConfig()

	if err := db.Init(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	db.GetDB().AutoMigrate(&user.User{}, &booking.Category{}, &booking.Location{}, &booking.Space{}, &booking.Event{}, &booking.EventTime{}, &booking.Reservation{})

	r := router.SetupRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
