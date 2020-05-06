package email

import (
	"fmt"
	"github.com/Gssssssssy/ns-stored/internal/task"
	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGenerateContent(t *testing.T) {
	result := task.Result{
		Name1:      "test1",
		Name2:      "test2",
		Price1:     "3999.99",
		Price2:     "4999.99",
		Available1: "Yes",
		Available2: "No",
	}
	htmlContext, err := generateTextContent(&result)
	assert.NoError(t, err)
	fmt.Printf("%# v\n", pretty.Formatter(htmlContext))
}

func TestEmail(t *testing.T) {
	result := task.Result{
		Name1:       "test1",
		Name2:       "test2",
		Price1:      "3999.99",
		Price2:      "4999.99",
		Available1:  "Yes",
		Available2:  "No",
		UpdatedTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	sdr := NewSender()
	err := sdr.Do(nil, &result)
	assert.NoError(t, err)
}
