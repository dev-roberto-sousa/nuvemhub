package services

// HomeService define a estrutura de um serviço de home.
type HomeService struct{}

// NewHomeService cria uma nova instância de HomeService.
func NewHomeService() *HomeService {
	return &HomeService{}
}

// GetWelcomeMessage retorna uma mensagem de boas-vindas.
func (s *HomeService) GetWelcomeMessage() string {
	return "Bem-vindo ao NuvemHub!"
}
