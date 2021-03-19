package main

import (
	"fc-poc/models"
)

// -----------------------------------------------------------------------------
// General Store Interfaces
// -----------------------------------------------------------------------------
type LivestreamStore interface {
	CreateLivestream(name string) error
	GetLivestreams() ([]models.LiveStream, error)
}
