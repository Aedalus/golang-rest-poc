package mysql

import (
	"fc-poc/models"
	"fmt"
	"gorm.io/gorm"
)

// -----------------------------------------------------------------------------
// DTOs
// -----------------------------------------------------------------------------
func toLivestreamDTO(l *models.LiveStream) *LiveStream {
	return &LiveStream{
		Model: gorm.Model{
			ID: l.ID,
		},
		LiveStream: models.LiveStream{
			Name: l.Name,
		},
	}
}

func fromLivestreamDTO(l *LiveStream) *models.LiveStream {
	return &models.LiveStream{
		ID:        l.Model.ID,
		Name:      l.Name,
		CreatedAt: l.Model.CreatedAt,
	}
}

// -----------------------------------------------------------------------------
// MySQL Store
// -----------------------------------------------------------------------------

type LivestreamStore struct {
	db *gorm.DB
}

func (s LivestreamStore) Migrate() {
	// Perform Migrations
	s.db.AutoMigrate(&LiveStream{})
}

func (LiveStream) TableName() string {
	return "live_streams"
}

type LiveStream struct {
	gorm.Model
	models.LiveStream
}

func (s *LivestreamStore) CreateLivestream(name string) error {
	result := s.db.Create(&LiveStream{
		LiveStream: models.LiveStream{
			Name: name,
		},
	})

	if result.Error != nil {
		return fmt.Errorf("error creating livestream: %v", result.Error)
	}

	return nil
}

func (s *LivestreamStore) GetLivestreams() ([]*models.LiveStream, error) {
	var resultModels []*LiveStream

	result := s.db.Find(&resultModels)

	if result.Error != nil {
		return nil, fmt.Errorf("error getting livestreams: %v", result.Error)
	}

	var livestreams []*models.LiveStream
	for _, m := range resultModels {
		livestreams = append(livestreams, fromLivestreamDTO(m))
	}

	return livestreams, nil
}
