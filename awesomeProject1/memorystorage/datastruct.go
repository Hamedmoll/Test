package memorystorage

import "awesomeProject1/entity"

type DataStruct struct {
	representations       map[int]entity.Representation
	representationCounter map[string]int
	staffCounter          map[string]int
}

func (d *DataStruct) ListAll() ([]entity.Representation, error) {
	list := make([]entity.Representation, 0)
	for _, rep := range d.representations {
		list = append(list, rep)
	}

	return list, nil
}

func (d *DataStruct) Last() int {

	return len(d.representations)
}

func (d *DataStruct) GetRepByID(id int) (entity.Representation, error) {
	return d.representations[id], nil
}

func (d *DataStruct) Save(newRep entity.Representation) error {
	d.representations[newRep.ID] = newRep
	d.representationCounter[newRep.Region]++
	d.staffCounter[newRep.Region] += newRep.StaffNumber

	return nil
}

func (d *DataStruct) Update(editedRep entity.Representation) error {

	oldRep := d.representations[editedRep.ID]
	d.representations[oldRep.ID] = editedRep

	d.staffCounter[oldRep.Region] -= oldRep.StaffNumber
	d.staffCounter[editedRep.Region] += editedRep.StaffNumber

	d.representationCounter[oldRep.Region]--
	d.representationCounter[editedRep.Region]++

	return nil
}

func (d *DataStruct) Status(region string) (int, int, error) {

	return d.representationCounter[region], d.staffCounter[region], nil
}

func New(representations map[int]entity.Representation, representationCounter map[string]int,
	staffCounter map[string]int) DataStruct {
	newDataStruct := DataStruct{
		representations:       representations,
		representationCounter: representationCounter,
		staffCounter:          staffCounter,
	}

	return newDataStruct
}
