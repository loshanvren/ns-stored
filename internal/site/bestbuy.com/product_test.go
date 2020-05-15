package bestbuy_com

import (
	"fmt"
	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoInquiry(t *testing.T) {
	result, err := doInquiry(nil)
	assert.NoError(t, err)
	fmt.Printf("%# v\n", pretty.Formatter(result))
}
