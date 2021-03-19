// +build mysql

package mysql

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func getDB() (*gorm.DB, error) {
	fmt.Println("Connecting to database")
	dsn := "root:Welcome1234@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening db connections: %v", err)
	}

	tx := db.Exec("CREATE DATABASE IF NOT EXISTS dev;")
	if tx.Error != nil {
		return nil, fmt.Errorf("error creating db: %v", err)
	}

	tx = db.Exec("USE dev;")
	if tx.Error != nil {
		return nil, fmt.Errorf("error selecting db: %v", err)
	}

	return db, err
}

func TestLivestreamSQL(t *testing.T) {
	db, err := getDB()
	assert.Nil(t, err)
	store := &LivestreamStore{db}
	store.Migrate()

	t.Run("Can create new Livestreams", func(t *testing.T) {
		db, err := getDB()
		assert.Nil(t, err)

		store := &LivestreamStore{db}

		err = store.CreateLivestream("unit_test_ls")
		assert.Nil(t, err)
		defer db.Where("name = ?", "unit_test_ls").Delete(&LiveStream{})

		var ls LiveStream
		tx := db.Where("name = ?", "unit_test_ls").First(&ls)
		assert.Nil(t, tx.Error)

		fmt.Printf("%v", ls)
		assert.NotNil(t, ls)
	})

	t.Run("Can get all livestreams", func(t *testing.T) {
		db, err := getDB()
		assert.Nil(t, err)

		store := &LivestreamStore{db}

		err = store.CreateLivestream("unit_test_1")
		assert.Nil(t, err)
		err = store.CreateLivestream("unit_test_2")
		assert.Nil(t, err)
		err = store.CreateLivestream("unit_test_3")
		assert.Nil(t, err)

		ls, err := store.GetLivestreams()
		assert.Nil(t, err)

		//want := []*models.LiveStream{
		//	{
		//		ID:        0,
		//		Name:      "unit_test_1",
		//		CreatedAt: time.Time{},
		//	},
		//	{
		//		ID:        1,
		//		Name:      "",
		//		CreatedAt: time.Time{},
		//	},
		//	{
		//		ID:        2,
		//		Name:      "",
		//		CreatedAt: time.Time{},
		//	},
		//}
		//assert.Equal(t, want, ls)
		assert.Len(t, ls, 3)
	})
}
