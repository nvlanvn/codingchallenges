package repositories

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nvlanvn/codingchallenges/task-tracker/models"
)

type TaskRepository struct {
	file string
}

func NewTaskRepository(file string) *TaskRepository {
	return &TaskRepository{file: file}
}

func (r *TaskRepository) load() ([]models.Task, error) {
	if _, err := os.Stat(r.file); os.IsNotExist(err) {
		return []models.Task{}, nil
	}
	data, err := os.ReadFile(r.file)
	if err != nil {
		return nil, err
	}
	var tasks []models.Task
	json.Unmarshal(data, &tasks)
	return tasks, nil
}

func (r *TaskRepository) save(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, " ", "")
	if err != nil {
		return nil
	}
	return os.WriteFile(r.file, data, 0644)
}

func (r *TaskRepository) Add(task models.Task) error {
	tasks, err := r.load()
	if err != nil {
		return err
	}
	tasks = append(tasks, task)
	return r.save(tasks)
}

func (r *TaskRepository) Update(id int, description string) (models.Task, error) {
	tasks, err := r.load()
	if err != nil {
		return models.Task{}, err
	}
	for index, task := range tasks {
		if task.Id == id {
			tasks[index].Description = description
			if err := r.save(tasks); err != nil {
				return models.Task{}, fmt.Errorf("the task cannot update")
			}
			return tasks[index], nil
		}
	}
	return models.Task{}, fmt.Errorf("the task is not found")
}

func (r *TaskRepository) UpdateStatus(id int, status models.Status) (models.Task, error) {
	tasks, err := r.load()
	if err != nil {
		return models.Task{}, err
	}
	for index, task := range tasks {
		if task.Id == id {
			tasks[index].Status = status
			if err := r.save(tasks); err != nil {
				return models.Task{}, fmt.Errorf("the task cannot update status %s", status)
			}
			return tasks[index], nil
		}
	}
	return models.Task{}, fmt.Errorf("the task is not found")
}

func (r *TaskRepository) Delete(id int) (models.Task, error) {
	tasks, err := r.load()
	if err != nil {
		return models.Task{}, err
	}
	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			if err := r.save(tasks); err != nil {
				return models.Task{}, err
			}
			return task, nil
		}
	}
	return models.Task{}, fmt.Errorf("the task with id %d is not found", id)
}

func (r *TaskRepository) FindLatestTask() models.Task {
	tasks, _ := r.load()
	if len(tasks) > 0 {
		return tasks[len(tasks)-1]
	}
	return models.Task{}
}

func (r *TaskRepository) ListTasksByStatus(status string) ([]models.Task, error) {
	tasks, err := r.load()
	if err != nil {
		return []models.Task{}, err
	}
	var tasksResult []models.Task
	for _, task := range tasks {
		if string(task.Status) == status {
			tasksResult = append(tasksResult, task)
		}
	}
	return tasksResult, nil
}

func (r *TaskRepository) ListTasks() ([]models.Task, error) {
	tasks, err := r.load()
	if err != nil {
		return []models.Task{}, err
	}
	return tasks, nil
}
