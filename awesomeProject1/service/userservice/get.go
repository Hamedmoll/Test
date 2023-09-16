package userservice

import (
	"awesomeProject1/pkg/validate"
	"fmt"
)

type GetRequest struct {
	Region string
}

type GetResponse struct {
}

func (s *Service) GetByID(request GetRequest) (GetResponse, error) {
	id, rErr := s.Reader.GetID(s.Memory)
	if rErr != nil {

		return GetResponse{}, fmt.Errorf("cant read %w", rErr)
	}

	rep, gErr := s.Memory.GetRepByID(id)
	if gErr != nil {
		return GetResponse{}, fmt.Errorf("cant get user %w", gErr)
	}

	//fmt.Println(rep, "\n\n\n\n", request.Region)

	if !validate.MatchIDRegion(rep, request.Region) {

		return GetResponse{}, fmt.Errorf("your id or region is incorrect")
	}

	wErr := s.Writer.GetRepByID(rep)
	if wErr != nil {

		return GetResponse{}, fmt.Errorf("cant write response %w", wErr)
	}

	response := GetResponse{}

	return response, nil
}
