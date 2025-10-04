package commands

import (
	"fmt"
	"strconv"

	"github.com/nvlanvn/codingchallenges/task-tracker/repositories"
)

type UpdateCommand struct {
	Repo *repositories.TaskRepository
}

func (r *UpdateCommand) Execute(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("missing id or description")
	}
	id, _ := strconv.Atoi(args[0])
	description := args[1]

	task, err := r.Repo.Update(id, description)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(task)

	return nil
}
