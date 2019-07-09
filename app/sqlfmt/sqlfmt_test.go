package sqlfmt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMustParseDatetime(t *testing.T) {
	dataTime := MustParseDatetime("2019-05-01 08:02:16")

	expectedDateTime, err := time.Parse(time.RFC3339, "2019-05-01T08:02:16Z00:00")

	assert.Nil(t, err)
	assert.Equal(t, expectedDateTime, dataTime)
}
