package filerepository

import (
	"awesomeProject1/entity"
	"awesomeProject1/service/userservice"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var file *os.File

type FileStorage struct {
	filePath string
}

func openFile(path string) error {
	var fErr error
	file, fErr = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if fErr != nil {

		return fmt.Errorf("cant open file")
	}

	return nil
}

func (f FileStorage) Save(newRep entity.Representation) error {
	//open file
	fErr := openFile(f.filePath)
	if fErr != nil {
		return fErr
	}

	//convert to json
	jsonRep, jErr := json.Marshal(newRep)
	if jErr != nil {
		return fmt.Errorf("cant convert to json")
	}
	jsonRep = append(jsonRep, '\n')

	//write in file
	_, wErr := file.WriteString(string(jsonRep))
	if wErr != nil {
		return fmt.Errorf("cant write in file")
	}

	return nil
}

func (f FileStorage) Update(memory userservice.MemoryStorage) error {
	//open file
	if fErr := openFile(f.filePath); fErr != nil {
		return fErr
	}

	//replace data from file with ""
	if fErr := os.WriteFile(f.filePath, []byte(""), 0644); fErr != nil {
		return fmt.Errorf("cant clear data %w", fErr)
	}

	//write all data to file
	representations, lErr := memory.ListAll()
	if lErr != nil {

		return fmt.Errorf("cant call listAll")
	}

	//fmt.Println(representations)
	for _, rep := range representations {
		err := f.Save(rep)
		if err != nil {

			return err
		}
	}

	return nil
}

func (f FileStorage) Load() ([]entity.Representation, error) {
	oErr := openFile(f.filePath)
	if oErr != nil {
		return make([]entity.Representation, 0), fmt.Errorf("cant open file in load step")
	}

	representations := make([]entity.Representation, 0)

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		var tmpRepresentation entity.Representation
		jErr := json.Unmarshal([]byte(fileScanner.Text()), &tmpRepresentation)
		if jErr != nil {
			fmt.Println("cant decode", jErr)
		}

		representations = append(representations, tmpRepresentation)
	}

	return representations, nil
}

func New(path string) FileStorage {
	newFileStorage := FileStorage{
		filePath: path,
	}

	return newFileStorage
}
