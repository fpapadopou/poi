package database

import (
	"github.com/fpapadopou/poi/datastore/model"
	"github.com/go-pg/pg"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"
)

var service POIService

func TestMain(m *testing.M) {
	// Setup connection.
	service.Conn = pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "postgres",
		Addr:     "127.0.0.1:5432",
		Password: "",
	})

	os.Exit(m.Run())
}

func TestPOIService_Insert(t *testing.T) {
	truncateDB()

	start := time.Now()
	p := &model.POI{
		Name: "test point of interest",
		Longitude: 15.232334,
		Latitude: 42.324312,
	}
	actual, err := service.CreatePOI(p)

	assert.NoError(t, err)

	// Auto assign IDs - should always be set to 1.
	assert.Equal(t, 1, actual.ID)

	assert.Equal(t, 15.232334, actual.Longitude)
	assert.Equal(t, 42.324312, actual.Latitude)
	assert.InDelta(t, start.Add(1 * time.Second).Unix(), actual.CreatedAt.Unix(), 1)

	// Location should be created on the fly.
	assert.Equal(t, "0101000020E61000006B7F677BF4762E40A5D93C0E83294540", actual.Location)
}

func TestPOIService_GetPOIByID(t *testing.T) {
	truncateDB()

	p := &model.POI{
		Name: "another poi",
		Longitude: 15.232334,
		Latitude: 42.324312,
	}
	var err error
	_, err = service.CreatePOI(p)

	assert.NoError(t, err)

	actual, err := service.GetPOIByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "another poi", actual.Name)
}

func TestPOIService_GetAll(t *testing.T) {
	truncateDB()

	p1 := &model.POI{
		Name: "one",
		Longitude: 15.232334,
		Latitude: 42.324312,
	}
	p2 := &model.POI{
		Name: "two",
		Longitude: 15.232334,
		Latitude: 42.324312,
	}
	var err error
	_, err = service.CreatePOI(p1)
	assert.NoError(t, err)

	_, err = service.CreatePOI(p2)
	assert.NoError(t, err)

	actual, err := service.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, "one", actual[0].Name)
	assert.Equal(t, "two", actual[1].Name)
}

func TestPOIService_UpdatePOI(t *testing.T) {
	truncateDB()

	p := &model.POI{
		Name: "one",
		Longitude: 15.232334,
		Latitude: 42.324312,
	}

	actual, err := service.CreatePOI(p)
	assert.NoError(t, err)

	actual.Name = "fake"
	err = service.UpdatePOI(actual)
	assert.NoError(t, err)

	updated, err := service.GetPOIByID(actual.ID)
	assert.NoError(t, err)
	assert.Equal(t, "fake", updated.Name)
}

func TestPOIService_DeletePOI(t *testing.T) {
	truncateDB()

	p := &model.POI{
		Name: "one",
		Longitude: 15.232334,
		Latitude: 42.324312,
	}

	actual, err := service.CreatePOI(p)
	assert.NoError(t, err)

	err = service.DeletePOI(actual)
	assert.NoError(t, err)

	_, err = service.GetPOIByID(actual.ID)
	assert.EqualError(t, err, "pg: no rows in result set")
}

func truncateDB() {
	// Delete all rows before each test and reset ID sequence.
	if _, err := service.Conn.Exec(`TRUNCATE TABLE poi RESTART IDENTITY`); err != nil {
		log.Fatalf("DB truncate failed: %v", err)
	}
}
