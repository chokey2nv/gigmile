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

const task_collection = "tasks"

type TaskService struct {
	*config.AppConfig
}

func (service *TaskService) CreateTask(ctx context.Context, taskDto *dtos.TaskDto) (*models.Task, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	col := service.DBClient.Database(service.Database).Collection(task_collection)
	task := &models.Task{
		Name:        taskDto.Name,
		Description: taskDto.Description,
		Id:          uuid.NewString(),
		CreateBy:    ctx.Value("userId").(string),
		CreateAt:    time.Now().UnixMilli(),
	}
	_, err := col.InsertOne(c, task)
	if err != nil {
		return nil, err
	}
	return task, nil
}
func (service *TaskService) GetTasks(ctx context.Context, pageOption *dtos.PageOption) ([]*models.Task, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	collection := config.Config.DBClient.Database(config.Config.Database).Collection(task_collection)
	cursor, err := collection.Find(
		c, bson.M{},
		&options.FindOptions{Limit: &pageOption.Limit, Skip: &pageOption.Skip},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)
	// Iterate through the cursor and collect tasks
	var tasks []*models.Task
	for cursor.Next(context.Background()) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			log.Println(err)
			continue
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}
