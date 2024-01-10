package config

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitMongo initializes the MongoDB client
func InitMongo(uri string) *mongo.Client {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(c, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client
}
func isZero(value interface{}) bool {
	return reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
}
func allObjects(arr primitive.A) bool {
	for _, v := range arr {
		if _, ok := v.(primitive.M); !ok {
			return false
		}
	}
	return true
}
func CleanUpBSON(bsonResult *primitive.M, stepKey string) {
	for key, value := range *bsonResult {
		if value == nil || isZero(value) {
			delete(*bsonResult, key)
			continue
		}
		switch v := value.(type) {
		case primitive.A:
			if len(v) == 0 {
				delete(*bsonResult, key)
			} else if !allObjects(v) {
				(*bsonResult)[key] = v
			} else {
				for key1, value1 := range v {
					if m, ok := value1.(primitive.M); ok {
						nextStepKey := fmt.Sprintf("%s.%d.", key, key1)
						CleanUpBSON(&m, nextStepKey)
					}
				}
			}
		case primitive.M:
			delete(*bsonResult, key)
			for key1, value1 := range v {
				nextStepKey := key + "." + key1
				if stepKey != "" {
					nextStepKey = stepKey + "." + nextStepKey
				}
				if m, ok := value1.(primitive.M); ok {
					CleanUpBSON(&m, nextStepKey)
				}
				(*bsonResult)[nextStepKey] = value1
			}
		}
	}
}
func StructToBSON(structObject interface{}, bsonResult *bson.M) error {
	byteValue, err := bson.Marshal(structObject)
	if err != nil {
		return err
	}
	err = bson.Unmarshal(byteValue, bsonResult)
	if err != nil {
		return err
	}
	CleanUpBSON(bsonResult, "")
	return nil
}
