package commands

import (
	"github.com/nvlanvn/codingchallenges/task-tracker/repositories"
	"github.com/sanity-io/litter"
)

type ListCommand struct {
	Repo *repositories.TaskRepository
}

func (r *ListCommand) Execute(args []string) error {
	if len(args) == 0 {
		tasks, err := r.Repo.ListTasks()
		if err != nil {
			return err
		}
		litter.Dump(tasks)
	} else {
		status := args[0]
		tasks, err := r.Repo.ListTasksByStatus(status)
		if err != nil {
			return err
		}
		litter.Dump(tasks)
	}
	return nil
}
