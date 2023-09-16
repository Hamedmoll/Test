package userservice

import (
	"fmt"
	"time"
)

type CreateRequest struct {
	Region string
}

type CreateResponse struct {
}

func (s *Service) CreateRepresentation(r CreateRequest) (CreateResponse, error) {
	//Read Data
	newRep, rErr := s.Reader.Create()
	if rErr != nil {

		return CreateResponse{}, fmt.Errorf("cant create newreprestentaion %w", rErr)
	}

	newRep.ID = s.Memory.Last() + 1
	newRep.Region = r.Region
	newRep.RegisterDate = time.Now()

	//fmt.Println(newRep)

	if sErr := s.Memory.Save(newRep); sErr != nil {

		return CreateResponse{}, fmt.Errorf("cant save data to memoryStorage %w", sErr)
	}

	//Save Data In Repository
	if sErr := s.Repository.Save(newRep); sErr != nil {

		return CreateResponse{}, fmt.Errorf("cant save data to repository %w", sErr)
	}

	response := CreateResponse{}
	//fmt.Println(s.Memory)
	return response, nil
}
