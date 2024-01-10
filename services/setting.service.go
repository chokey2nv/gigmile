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
)

const setting_collection = "settings"

type SettingService struct {
	*config.AppConfig
}

func (service *SettingService) CreateSetting(ctx context.Context, settingDto *dtos.SettingDto) (*models.Setting, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	col := service.DBClient.Database(service.Database).Collection(setting_collection)
	setting := &models.Setting{
		Id:          uuid.NewString(),
		Name:        settingDto.Name,
		Description: settingDto.Description,
		Value:       settingDto.Value,
		UpdateBy:    c.Value("userId").(string),
		UpdatedAt:   time.Now().UnixMilli(),
	}
	_, err := col.InsertOne(c, setting)
	if err != nil {
		return nil, err
	}
	return setting, nil
}
func (service *SettingService) GetSettings(ctx context.Context) ([]*models.Setting, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	collection := config.Config.DBClient.Database(config.Config.Database).Collection(setting_collection)
	cursor, err := collection.Find(
		c, bson.M{},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)
	// Iterate through the cursor and collect settings
	var settings []*models.Setting
	for cursor.Next(context.Background()) {
		var setting models.Setting
		if err := cursor.Decode(&setting); err != nil {
			log.Println(err)
			continue
		}
		settings = append(settings, &setting)
	}
	return settings, nil
}
func (service *SettingService) GetSetting(ctx context.Context, settingFilter *dtos.SettingDto) (*models.Setting, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	collection := config.Config.DBClient.Database(config.Config.Database).Collection(setting_collection)
	filter := bson.M{}
	config.StructToBSON(settingFilter, &filter)
	setting := models.Setting{}
	err := collection.FindOne(
		c,
		filter,
	).Decode(&setting)
	if err != nil {
		return nil, err
	}
	return &setting, nil
}
