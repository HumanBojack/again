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

func (g *GormDB) GetTask(id string) (*Task, error) {
	task := Task{}
	result := g.DB.First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (g *GormDB) UpdateTask(id string, task *TaskInput) error {
	newTask := Task{
		TaskInput: *task,
	}
	result := g.DB.Model(&Task{}).Where("id = ?", id).Updates(newTask)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (g *GormDB) DeleteTask(id string) error {
	result := g.DB.Delete(&Task{}, id)
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
