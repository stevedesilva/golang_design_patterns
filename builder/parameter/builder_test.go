package parameter_test

import (
	"golangDesignPatterns/builder/parameter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendEmail(t *testing.T) {
	result := parameter.SendEmail(func(builder *parameter.EmailBuilder) {
		builder.From("steve@test.com").To("test@test.com").Subject("Meeting").Body("Let's setup a meeting.")
	})
	expected := "&{steve@test.com test@test.com Meeting Let's setup a meeting.}"
	assert.Equal(t, result, expected)
}
