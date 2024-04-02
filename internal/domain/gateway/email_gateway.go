package gateway

type EmailGateway interface {
	SendEmail(to, subject, body string) error
}
