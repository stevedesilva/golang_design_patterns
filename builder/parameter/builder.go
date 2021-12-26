package parameter

import (
	"fmt"
	"strings"
)

/*
  Ensure user of the API user builder is used instead of underlined object directly
*/

// package scope
type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("from should contain @")
	}
	b.email.from = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.email.to = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

func sendEmailImpl(email *email) string {
	// sends email
	return fmt.Sprintf("%s", email)
}

type build func(builder *EmailBuilder)

func SendEmail(action build) string {
	eb := EmailBuilder{}
	action(&eb)
	return sendEmailImpl(&eb.email)
}
