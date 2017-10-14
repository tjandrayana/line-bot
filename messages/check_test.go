package messages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestcheckSimilarity(t *testing.T) {

	msg1 := "i have graduate from duta wacana christian university"
	msg2 := "i have graduate from ambassador christian university "

	assert.Equal(t, 0.90, checkSimilarity(msg1, msg2))

}
