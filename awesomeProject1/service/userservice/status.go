package userservice

import "fmt"

type StatusRequest struct {
	Region string
}

type StatusResponse struct {
}

func (s *Service) Status(request StatusRequest) (StatusResponse, error) {
	repCount, StaffCount, sErr := s.Memory.Status(request.Region)
	if sErr != nil {

		return StatusResponse{}, fmt.Errorf("cant get status from memory %w", sErr)
	}
	
	wErr := s.Writer.GetStatus(repCount, StaffCount)
	if wErr != nil {
		return StatusResponse{}, fmt.Errorf("cant write status %w", wErr)
	}

	return StatusResponse{}, nil
}
