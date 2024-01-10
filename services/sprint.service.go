package services

import (
	"context"
	"log"
	"time"

	"github.com/chokey2nv/gigmile/config"
	"github.com/chokey2nv/gigmile/dtos"
	"github.com/chokey2nv/gigmile/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const sprint_collection = "sprints"

type SprintService struct {
	*config.AppConfig
}

func (service *SprintService) CreateSprint(ctx context.Context, sprintDto *dtos.SprintDto) (*models.Sprint, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	col := service.DBClient.Database(service.Database).Collection(sprint_collection)
	format := "2006-02-01"
	startTimestamp, err := time.Parse(format, sprintDto.StartDate)
	if err != nil {
		return nil, Errorf(err)
	}
	endTimestamp, err := time.Parse(format, sprintDto.EndDate)
	if err != nil {
		return nil, Errorf(err)
	}
	sprint := &models.Sprint{
		Name:           sprintDto.Name,
		Description:    sprintDto.Description,
		StartTimestamp: startTimestamp.UnixMilli(),
		EndTimestamp:   endTimestamp.UnixMilli(),
		Id:             uuid.NewString(),
	}
	_, err = col.InsertOne(c, sprint)
	if err != nil {
		return nil, err
	}
	return sprint, nil
}
func (service *SprintService) GetSprints(ctx context.Context, pageOption *dtos.PageOption) ([]*models.Sprint, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	collection := config.Config.DBClient.Database(config.Config.Database).Collection(sprint_collection)
	cursor, err := collection.Find(
		c, bson.M{},
		&options.FindOptions{Limit: &pageOption.Limit, Skip: &pageOption.Skip},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)
	// Iterate through the cursor and collect sprints
	sprints := []*models.Sprint{}
	for cursor.Next(context.Background()) {
		var sprint models.Sprint
		if err := cursor.Decode(&sprint); err != nil {
			log.Println(err)
			continue
		}
		sprints = append(sprints, &sprint)
	}
	return sprints, nil
}
