package models

import (
	models "dinosaur-cage/models/dinosaur"
)

// Cage represents a dinosaur cage with specific requirements.
type Cage struct {
	ID               int               `json:"id"`
	Name             string            `json:"name"`
	PowerStatus      PowerStatus       `json:"power_status"`
	MaxCapacity      int               `json:"max_capacity"`
	CurrentCapacity  int               `json:"current_capacity"`
	CurrentDinosaurs []models.Dinosaur `json:"current_dinosaurs,omitempty"`
}

// PowerStatus represents the power status of a cage.
type PowerStatus string

// DefaultPowerStatus represents the default power status for a new cage.
const DefaultPowerStatus PowerStatus = "DOWN"

// List of available power status values.
const (
	PowerStatusActive PowerStatus = "ACTIVE"
	PowerStatusDown   PowerStatus = "DOWN"
)
