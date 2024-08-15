package db

func (g *GormDB) CreateTask(task *TaskInput) error {
	newTask := Task{
		TaskInput: *task,
	}
	result := g.DB.Create(&newTask)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (g *GormDB) GetAllTasks() ([]Task, error) {
	tasks := []Task{}
	result := g.DB.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}
