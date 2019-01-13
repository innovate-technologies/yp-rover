package store

import (
	"context"

	"github.com/innovate-technologies/yp-rover/pkg/tunein"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// AddTuneInGenre adds a roved TuneIn genre
func (s *Store) AddTuneInGenre(genre tunein.Genre) error {
	var res tunein.Genre
	err := s.db.Collection("tunein_genres").FindOne(context.Background(), bson.M{"name": genre.Name}).Decode(&res)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}

	_, err = s.db.Collection("tunein_genres").InsertOne(context.Background(), genre)
	return err
}

// AddTuneInStation adds a roved TuneIn station
func (s *Store) AddTuneInStation(station tunein.Station) error {
	var res tunein.Station
	err := s.db.Collection("tunein_stations").FindOne(context.Background(), bson.M{"tuneInURL": station.TuneInURL}).Decode(&res)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}

	_, err = s.db.Collection("tunein_stations").InsertOne(context.Background(), station)
	return err
}
