package bestbuy_com

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoInquiry(t *testing.T) {
	result, err := doInquiry(nil, nil)
	assert.NoError(t, err)
}
