package tickets


import ("context")

type service struct {
	repo Repository
}

type Service interface {
	GetTotalTickets(ctx context.Context, destination string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

func NewService(repository Repository) Service {
	return &service{
		repo: repository,
	}
}


func(s *service) GetTotalTickets(ctx context.Context, destination string) (int, error){
	tickets, err := s.repo.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	
	return len(tickets), nil
}

func(s *service) AverageDestination(ctx context.Context, destination string) (float64, error){
	tickets, err := s.repo.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0.0, err
	}

	total, err := s.repo.GetAll(ctx)
	if err != nil {
		return 0.0, err
	}

	average := float64(len(tickets)) / float64(len(total))
	return average, nil
}