package db

func (g *GormDB) CreateTask(task *Task) error {
	result := g.DB.Create(&task)
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

func (g *GormDB) UpdateTask(id string, task *Task) error {
	result := g.DB.Model(&Task{}).Where("id = ?", id).Updates(task)
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
	result := g.DB.Order("ID").Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}
