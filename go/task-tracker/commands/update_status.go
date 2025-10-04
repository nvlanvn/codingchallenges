package commands

import (
	"fmt"
	"strconv"

	"github.com/nvlanvn/codingchallenges/task-tracker/models"
	"github.com/nvlanvn/codingchallenges/task-tracker/repositories"
)

type UpdateStatusCommand struct {
	Repo   *repositories.TaskRepository
	Status models.Status
}

func (r *UpdateStatusCommand) Execute(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing task id")
	}
	taskId, _ := strconv.Atoi(args[0])
	_, err := r.Repo.UpdateStatus(taskId, r.Status)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
