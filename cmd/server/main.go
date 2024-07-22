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
	// Inicializar la configuración
	config.InitConfig()

	// Inicializar la base de datos
	if err := db.Init(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Migrar los modelos
	db.GetDB().AutoMigrate(&user.User{}, &booking.Category{}, &booking.Location{}, &booking.Space{}, &booking.Event{}, &booking.EventTime{}, &booking.Reservation{})

	// Configurar el router
	r := router.SetupRouter()

	// Correr el servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
