package commands

import (
	"fmt"
	"strconv"

	"github.com/nvlanvn/codingchallenges/task-tracker/repositories"
)

type DeleteCommand struct {
	Repo *repositories.TaskRepository
}

func (r *DeleteCommand) Execute(args []string) error {
	if len(args) < 1 {
		fmt.Println("the task id is required")
	}
	taskId, _ := strconv.Atoi(args[0])
	_, err := r.Repo.Delete(taskId)
	if err != nil {
		return fmt.Errorf("cannot remove task id %d", taskId)
	}
	return nil
}
