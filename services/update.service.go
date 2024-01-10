package services

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/chokey2nv/gigmile/config"
	"github.com/chokey2nv/gigmile/dtos"
	"github.com/chokey2nv/gigmile/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	update_collection = "updates"
	ErrCreateUpdate   = errors.New("unable to create update")
	ErrGetUpdate      = errors.New("unable to get update")
)

type UpdateService struct {
	*config.AppConfig
}

func (service *UpdateService) CreateUpdate(ctx context.Context, updateDto *dtos.UpdateDto) (*models.Update, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	collection := service.DBClient.Database(service.Database).Collection(update_collection)

	update := &models.Update{
		Id:                       uuid.NewString(),
		EmployeeID:               ctx.Value("userId").(string), //TODO: get this from access token
		Timestamp:                time.Now().UnixMilli(),
		Week:                     updateDto.Week,
		SprintID:                 updateDto.SprintID,
		TaskIDs:                  updateDto.TaskIDs,
		PreviousCompletedTaskIDs: updateDto.PreviousCompletedTaskIDs,
		CurrentTaskIDs:           updateDto.CurrentTaskIDs,
		BlockedByEmployeeIDs:     updateDto.BlockedByEmployeeIDs,
		Breakaway:                updateDto.Breakaway,
		CreateAt:                 primitive.NewDateTimeFromTime(time.Now()),
	}
	// Insert the update into MongoDB
	_, err := collection.InsertOne(c, update)
	if err != nil {
		return nil, Errorf(ErrCreateUpdate)
	}
	return update, nil
}

func (service *UpdateService) GetUpdates(
	ctx context.Context,
	filter *dtos.GetUPdateFilterDto,
	pageOption *dtos.PageOption,
) ([]*dtos.UpdateResponseDto, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	collection := service.DBClient.Database(service.Database).Collection(update_collection)
	lookupEmployee := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: User_collection},
			{Key: "localField", Value: "employeeId"},
			{Key: "foreignField", Value: "id"},
			{Key: "as", Value: "employee"},
		}},
	}
	unwindEmployee := bson.D{
		{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$employee"},
			{Key: "preserveNullAndEmptyArrays", Value: true},
		}},
	}
	lookUpSprintBadge := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: sprint_collection},
			{Key: "localField", Value: "sprintId"},
			{Key: "foreignField", Value: "id"},
			{Key: "as", Value: "sprint"},
		}},
	}
	unwindSprintStage := bson.D{
		{Key: "$unwind", Value: "$sprint"},
	}
	addFieldsStage := bson.D{
		{Key: "$addFields", Value: bson.D{
			{Key: "dayOfWeek", Value: bson.D{
				{Key: "$dayOfWeek", Value: "$createdAt"},
			}},
		}},
	}

	pipeline := mongo.Pipeline{lookupEmployee, unwindEmployee, lookUpSprintBadge, unwindSprintStage, addFieldsStage}
	if filter != nil {
		if filter.EmployeeID != "" {
			fetchById := bson.D{{
				Key: "$match", Value: bson.D{
					{Key: "employeeId", Value: filter.EmployeeID},
				}}}
			pipeline = append(pipeline, fetchById)
		}
		if filter.EmployeeName != "" {
			fetchById := bson.D{{
				Key: "$match", Value: bson.D{
					{Key: "$or", Value: bson.A{
						bson.D{{Key: "employee.lastName", Value: bson.D{{Key: "$regex", Value: filter.EmployeeName}, {Key: "$options", Value: "i"}}}},
						bson.D{{Key: "employee.firstName", Value: bson.D{{Key: "$regex", Value: filter.EmployeeName}, {Key: "$options", Value: "i"}}}},
					}},
				}}}
			pipeline = append(pipeline, fetchById)
		}
		if filter.Week != 0 {
			filterByWeek := bson.D{{
				Key: "$match", Value: bson.D{
					{Key: "week", Value: filter.Week},
				}}}
			pipeline = append(pipeline, filterByWeek)
		}
		if filter.DayOfWeek != 0 {
			filterByWeek := bson.D{{
				Key: "$match", Value: bson.D{
					{Key: "dayOfWeek", Value: filter.DayOfWeek},
				}}}
			pipeline = append(pipeline, filterByWeek)
		}
		if filter.SprintName != "" {
			filterByWeek := bson.D{{
				Key:   "$match",
				Value: bson.D{{Key: "sprint.name", Value: bson.D{{Key: "$regex", Value: filter.SprintName}, {Key: "$options", Value: "i"}}}},
			}}
			pipeline = append(pipeline, filterByWeek)
		}
	}
	cursor, err := collection.Aggregate(c, pipeline)
	if err != nil {
		return nil, Errorf(err)
	}
	defer cursor.Close(c)

	settingService := SettingService{AppConfig: service.AppConfig}
	standUpStartTime, err := settingService.GetSetting(ctx, &dtos.SettingDto{
		Name: models.SettingName_StandUpStartTime,
	})
	if err != nil {
		return nil, Errorf(err)
	}
	standUpEndTime, err := settingService.GetSetting(ctx, &dtos.SettingDto{
		Name: models.SettingName_StandUpEndTime,
	})
	if err != nil {
		return nil, Errorf(err)
	}
	endTimeInt, err := strconv.ParseInt(strings.Join(strings.Split(standUpEndTime.Value, ":"), ""), 10, 64)
	if err != nil {
		return nil, Errorf(err)
	}
	startTimeInt, err := strconv.ParseInt(strings.Join(strings.Split(standUpStartTime.Value, ":"), ""), 10, 64)
	if err != nil {
		return nil, Errorf(err)
	}
	// Iterate through the cursor and collect updates
	var updates []*dtos.UpdateResponseDto
	for cursor.Next(c) {
		var update models.Update
		var a interface{}
		if err := cursor.Decode(&update); err != nil {
			Errorf(err)
			continue
		}
		Errorf(a)
		date := time.UnixMilli(update.Timestamp)
		checkedInTime := date.Format("15:04")
		checkedInTimeInt, err := strconv.ParseInt(strings.Join(strings.Split(checkedInTime, ":"), ""), 10, 64)
		if err != nil {
			return nil, Errorf(err)
		}
		var status string
		if checkedInTimeInt > endTimeInt {
			status = "After stand-up"
		} else if checkedInTimeInt < startTimeInt {
			status = "Before stand-up"
		} else {
			status = "Within stand-up"
		}
		updates = append(updates, &dtos.UpdateResponseDto{
			Id:                       update.Id,
			EmployeeID:               update.EmployeeID,
			Timestamp:                update.Timestamp,
			SprintID:                 update.SprintID,
			TaskIDs:                  update.TaskIDs,
			PreviousCompletedTaskIDs: update.PreviousCompletedTaskIDs,
			CurrentTaskIDs:           update.CurrentTaskIDs,
			BlockedByEmployeeIDs:     update.BlockedByEmployeeIDs,
			Breakaway:                update.Breakaway,
			CheckedInTime:            checkedInTime,
			Employee: &dtos.UpdateUserResponseDto{
				LastName:  update.Employee.LastName,
				FirstName: update.Employee.FirstName,
				Email:     update.Employee.Email,
				UserRole:  update.Employee.UserRole,
			},
			CreateAt:  update.CreateAt,
			Status:    status,
			Week:      update.Week,
			DayOfWeek: update.DayOfWeek,
		})
	}

	return updates, nil
}
