package userservice

import (
	"awesomeProject1/entity"
	"fmt"
)

type ListRequest struct {
	Region string
}

type ListResponse struct {
}

func (s *Service) ListByRegion(request ListRequest) (ListResponse, error) {
	list := make([]entity.Representation, 0)
	allRep, lErr := s.Memory.ListAll()
	if lErr != nil {

		return ListResponse{}, fmt.Errorf("cant get list of all representation %w", lErr)
	}

	for _, rep := range allRep {
		if rep.Region == request.Region {
			list = append(list, rep)
		}
	}

	wErr := s.Writer.ListByRegion(list)
	if wErr != nil {

		return ListResponse{}, fmt.Errorf("cant write list %w", lErr)
	}

	return ListResponse{}, nil
}
