package db

import (
	"net/url"
	"time"

	otime "github.com/humanbojack/again/server/packages/time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskInput
}

type TaskInput struct {
	Title           string        `json:"title,omitempty"`
	Description     string        `json:"description,omitempty"`
	Frequency       time.Duration `json:"interval,omitempty"`
	LastCompletedAt *time.Time    `json:"last_completed_at,omitempty"`
}

// TaskInputFromForm creates a new TaskInput from a HTML form
// TODO: move to a service instead
func TaskInputFromForm(form url.Values) (TaskInput, error) {
	t := TaskInput{}

	// TODO: validate fields

	// Map the fields to the model
	for k, v := range form {
		if k == "Frequency" {
			d, err := otime.ParseLargeDuration(v[0])
			if err != nil {
				return TaskInput{}, err
			}
			t.Frequency = d
			continue
		}

		if k == "LastCompleted" {
			var lastCompleted time.Time
			if v[0] == "" {
				lastCompleted = time.Now()
			} else {
				var err error
				lastCompleted, err = time.Parse("2006-01-02T15:04", v[0])
				if err != nil {
					return TaskInput{}, err
				}
			}
			t.LastCompletedAt = &lastCompleted
			continue
		}

		switch k {
		case "Title":
			t.Title = v[0]
		case "Description":
			t.Description = v[0]
		}
	}
	return t, nil
}
