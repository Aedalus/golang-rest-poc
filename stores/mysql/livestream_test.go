package mysql

import (
	"fc-poc/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestLivestreamDTOs(t *testing.T) {
	t.Run("it can convert to mysql dto", func(t *testing.T) {
		createdTime := time.Now()
		domain := &models.LiveStream{
			ID:        1,
			Name:      "foo",
			CreatedAt: createdTime,
		}

		dto := toLivestreamDTO(domain)

		want := &LiveStream{
			Model: gorm.Model{
				ID: 1,
			},
			LiveStream: models.LiveStream{
				Name: "foo",
			},
		}
		assert.Equal(t, want, dto)
	})

	t.Run("it can convert from mysql dto", func(t *testing.T) {
		createdTime := time.Now()
		updatedTime := time.Now().Add(time.Minute * 60)
		dto := &LiveStream{
			Model: gorm.Model{
				ID:        10,
				CreatedAt: createdTime,
				UpdatedAt: updatedTime,
			},
			LiveStream: models.LiveStream{
				Name: "Foo",
			},
		}

		domain := fromLivestreamDTO(dto)

		want := &models.LiveStream{
			ID:        10,
			Name:      "Foo",
			CreatedAt: createdTime,
		}

		assert.Equal(t, want, domain)
	})
}
