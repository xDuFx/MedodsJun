package repository

import (
	"context"
	"log"

	"testjun/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *Session) Create(lg, refresh string) (interface{}, error) {
	err := s.con.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	u := models.UserData{
		GUID: lg,
		RefreshToken: refresh,
	}
	insertResult, err := s.cll.InsertOne(context.TODO(), u)
	return insertResult.InsertedID, err
}

func (s *Session) Find(guid string) (*models.UserData, bool, error) {
	err := s.con.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.D{{Key: "guid", Value: guid}}
	var result models.UserData
	err = s.cll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return &models.UserData{}, false, err
	}
	return &result, true, nil
}

func (s *Session) Update(guid, token string)  error {
	err := s.con.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.D{{Key: "guid", Value: guid}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "refreshtoken", Value: token}}}}
	_, err = s.cll.UpdateOne(context.TODO(), filter, update)
	return err
}
