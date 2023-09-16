package userservice

import (
	"fmt"
)

type EditRequest struct {
	Region string
}

type EditResponse struct {
}

func (s *Service) EditRepresentation(e EditRequest) (EditResponse, error) {
	//Read Data
	editedRep, eErr := s.Reader.Edit(s.Memory, e.Region)
	if eErr != nil {
		return EditResponse{}, fmt.Errorf("cant edit representation %w", eErr)
	}

	//Modify MemoryStorage
	if uErr := s.Memory.Update(editedRep); uErr != nil {

		return EditResponse{}, fmt.Errorf("cant update memoryStorage %w", uErr)
	}

	//Modify Repository
	if uErr := s.Repository.Update(s.Memory); uErr != nil {

		return EditResponse{}, fmt.Errorf("cant update repositoryStorage %w", uErr)
	}

	response := EditResponse{}

	return response, nil
}
