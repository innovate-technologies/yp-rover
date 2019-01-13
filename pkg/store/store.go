package store

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo/options"

	"github.com/mongodb/mongo-go-driver/x/bsonx"

	"github.com/innovate-technologies/yp-rover/internal/config"
	"github.com/innovate-technologies/yp-rover/pkg/shoutcastcom"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/pkg/errors"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// Store allows interaction with the discover.fm database
type Store struct {
	db *mongo.Database
}

// New sets up a connection to the discover.fm database
func New(conf config.Config) (*Store, error) {
	client, err := mongo.NewClient(conf.MongoDBURL)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	db := client.Database(conf.MongoDBDatabase)

	return &Store{
		db: db,
	}, nil
}

// Close closes the connection to the server
func (s *Store) Close() {
	s.db.Client().Disconnect(context.Background())
}

// Migrate will make sure indexes are set
func (s *Store) Migrate() {
	opts := options.Index()
	opts.SetBackground(true)
	opts.SetName("sc_genre")
	err := s.EnsureIndex(context.Background(), s.db.Collection("sc_genres"), []string{"name"}, opts)
	if err != nil {
		log.Printf("Ensureindex error: %s", err)
	}

	opts = options.Index()
	opts.SetBackground(true)
	opts.SetName("tunein_genre")
	err = s.EnsureIndex(context.Background(), s.db.Collection("tunein_genres"), []string{"name"}, opts)
	if err != nil {
		log.Printf("Ensureindex error: %s", err)
	}

	opts = options.Index()
	opts.SetBackground(true)
	opts.SetUnique(true)
	opts.SetName("sc_id")
	err = s.EnsureIndex(context.Background(), s.db.Collection("sc_stations"), []string{"shoutcastID"}, opts)
	if err != nil {
		log.Printf("Ensureindex error: %s", err)
	}

	opts = options.Index()
	opts.SetBackground(true)
	opts.SetUnique(true)
	opts.SetName("tunein_name")
	err = s.EnsureIndex(context.Background(), s.db.Collection("tunein_stations"), []string{"name"}, opts)
	if err != nil {
		log.Printf("Ensureindex error: %s", err)
	}
}

// GetSHOUTcastGenres gets the roved SHOUTcast.com genres
func (s *Store) GetSHOUTcastGenres() ([]string, error) {
	ctx := context.Background()
	cur, err := s.db.Collection("sc_genres").Find(ctx, nil)
	if err != nil {
		return nil, err
	}

	genres := []string{}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result shoutcastGenre
		err := cur.Decode(&result)
		if err != nil {
			log.Println(err)
		}

		genres = append(genres, result.Name)
	}

	return genres, nil
}

// AddSHOUTcastGenre adds a roved SHOUTcast.com genre
func (s *Store) AddSHOUTcastGenre(genre string) error {
	var res shoutcastGenre
	err := s.db.Collection("sc_genres").FindOne(context.Background(), bson.M{"name": genre}).Decode(&res)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}

	if err == mongo.ErrNoDocuments {
		_, err = s.db.Collection("sc_genres").InsertOne(context.Background(), shoutcastGenre{Name: genre})
		return err
	}
	return nil
}

// AddSHOUTcastStation adds a roved SHOUTcast.com station
func (s *Store) AddSHOUTcastStation(station shoutcastcom.Station) error {
	var res shoutcastGenre
	err := s.db.Collection("sc_stations").FindOne(context.Background(), bson.M{"shoutcastID": station.ID}).Decode(&res)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}
	if err == mongo.ErrNoDocuments {
		_, err = s.db.Collection("sc_stations").InsertOne(context.Background(), station)
		return err
	}
	return nil
}

// EnsureIndex will ensure the index model provided is on the given collection.
func (s *Store) EnsureIndex(ctx context.Context, c *mongo.Collection, keys []string, opts *options.IndexOptions) error {
	idxs := c.Indexes()

	ks := bson.M{}
	for _, k := range keys {
		ks[k] = -1
	}
	idm := mongo.IndexModel{
		Keys:    ks,
		Options: opts,
	}

	v := idm.Options.Name
	if v == nil {
		return errors.New("must provide a key name for index")
	}
	expectedName := v

	cur, err := idxs.List(ctx)
	if err != nil {
		return errors.Wrap(err, "unable to list indexes")
	}

	found := false
	for cur.Next(ctx) {
		d := &bsonx.Doc{}

		if err := cur.Decode(d); err != nil {
			return errors.Wrap(err, "unable to decode bson index document")
		}

		v, ok := d.Lookup("name").StringValueOK()
		if ok && v == *expectedName {
			found = true
			break
		}
	}

	if found {
		return nil
	}

	_, err = idxs.CreateOne(ctx, idm)
	return err
}
