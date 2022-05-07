package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/siwodevilheart/bookings/internal/config"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing
	default:
		t.Error(fmt.Sprintf("type is not Chi mux, but %T", v))
	}
}
