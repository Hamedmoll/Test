package consolewriter

import (
	"awesomeProject1/entity"
	"fmt"
)

type Writer struct {
}

func (w Writer) GetRepByID(rep entity.Representation) error {
	fmt.Printf("%+v\n", rep)

	return nil
}

func (w Writer) GetStatus(repCount, staffCount int) error {
	fmt.Printf("number of Representations: %d\nnumber of staff: %d\n", repCount, staffCount)

	return nil
}

func (w Writer) ListByRegion(list []entity.Representation) error {
	fmt.Printf("list of representations :\n")

	for _, rep := range list {
		fmt.Printf("%+v\n", rep)
	}

	return nil
}
