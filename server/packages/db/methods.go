package db

func (g *GormDB) GetAllTasks() ([]Task, error) {
	tasks := []Task{}
	result := g.DB.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}
