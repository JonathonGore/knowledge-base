package answer

import (
	"fmt"
	"time"
)

const (
	minContentLength = 10
	maxContentLength = 1000
)

type Answer struct {
	ID          int       `json:"id"`
	SubmittedOn time.Time `json:"submitted-on"`
	Author      int       `json:"author"`
	Username    string    `json:"username"`
	Content     string    `json:"content"`
	Accepted    bool      `json:"accepted"`
	Question    int       `json:"question"`
}

// TODO
func Validate(ans Answer) error {
	err := validateContent(ans.Content)
	if err != nil {
		return err
	}

	return nil
}

func validateContent(content string) error {
	if len(content) < minContentLength {
		return fmt.Errorf("content length must be at least %v characters", minContentLength)
	} else if len(content) > maxContentLength {
		return fmt.Errorf("content length must be at most %v characters", maxContentLength)
	}

	return nil
}
