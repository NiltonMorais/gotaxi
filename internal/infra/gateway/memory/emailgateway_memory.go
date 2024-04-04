package gateway

import (
	"fmt"
)

// EmailMemoryGateway é uma implementação de EmailGateway que simula o envio de e-mails, armazenando os dados em memória.
type EmailMemoryGateway struct {
	// Aqui você pode adicionar qualquer campo adicional necessário para a implementação
}

// NewEmailMemoryGateway cria uma nova instância de EmailMemoryGateway.
func NewEmailMemoryGateway() *EmailMemoryGateway {
	return &EmailMemoryGateway{}
}

// SendEmail simula o envio de um e-mail, armazenando as informações em memória.
func (eg *EmailMemoryGateway) SendEmail(to, subject, body string) error {
	// Aqui você pode implementar a lógica para armazenar os detalhes do e-mail em memória.
	fmt.Printf("E-mail enviado para: %s\nAssunto: %s\nCorpo: %s\n", to, subject, body)
	return nil
}
