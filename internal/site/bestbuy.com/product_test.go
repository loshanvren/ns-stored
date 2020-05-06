package bestbuy_com

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoInquiry(t *testing.T) {
	_, err := doInquiry(nil)
	assert.NoError(t, err)
}
