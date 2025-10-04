package commands

import (
	"fmt"

	"github.com/nvlanvn/codingchallenges/task-tracker/factories"
	"github.com/nvlanvn/codingchallenges/task-tracker/repositories"
)

type AddCommand struct {
	Repo *repositories.TaskRepository
}

func (c *AddCommand) Execute(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing the description")
	}
	task := factories.NewTask(args[0], c.Repo)
	err := c.Repo.Add(task)
	if err != nil {
		return err
	}
	fmt.Printf("Task added successfully: %d", task.Id)
	return nil
}
