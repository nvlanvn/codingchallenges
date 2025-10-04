package factories

import (
	"time"

	"github.com/nvlanvn/codingchallenges/task-tracker/models"
	"github.com/nvlanvn/codingchallenges/task-tracker/repositories"
)

func NewTask(description string, r *repositories.TaskRepository) models.Task {
	idCounter := 1
	latestTask := r.FindLatestTask()
	if latestTask != (models.Task{}) {
		idCounter = latestTask.Id + 1
	}

	newTask := models.Task{
		Id:          idCounter,
		Description: description,
		Status:      models.Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return newTask
}
