package consolereader

import (
	"awesomeProject1/entity"
	"awesomeProject1/pkg/validate"
	"awesomeProject1/service/userservice"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

var scanner = bufio.NewScanner(os.Stdin)

type Reader struct {
}

func (r Reader) Create() (entity.Representation, error) {
	fmt.Println("Please enter your name")
	scanner.Scan()
	name := scanner.Text()

	if !validate.Name(name) {

		return entity.Representation{}, fmt.Errorf("your name is invalid")
	}

	fmt.Println("Please enter your address")
	scanner.Scan()
	address := scanner.Text()

	if !validate.Address(address) {

		return entity.Representation{}, fmt.Errorf("your address is invalid")

	}

	fmt.Println("Please enter your phoneNumber")
	scanner.Scan()
	phoneNumber := scanner.Text()

	if !validate.PhoneNumber(phoneNumber) {

		return entity.Representation{}, fmt.Errorf("your phonenumber is invalid")
	}

	fmt.Println("Please enter your staffNumber")
	scanner.Scan()
	staffNumberString := scanner.Text()

	var staffNumber int
	if !validate.StaffNumber(staffNumberString) {
		return entity.Representation{}, fmt.Errorf("your staffnumber is invalid")
	} else {
		staffNumber, _ = strconv.Atoi(staffNumberString)
	}

	return entity.Representation{
		Name:         name,
		Address:      address,
		PhoneNumber:  phoneNumber,
		StaffNumber:  staffNumber,
		RegisterDate: time.Now(),
	}, nil

}

func (r Reader) Edit(memory userservice.MemoryStorage, region string) (entity.Representation, error) {
	//Read and Validate ID
	fmt.Println("Please enter your ID")
	scanner.Scan()
	idString := scanner.Text()
	id, strErr := strconv.Atoi(idString)
	if strErr != nil {
		return entity.Representation{}, fmt.Errorf("cant convert string to int")
	}

	if !validate.ID(id, memory.Last()) {

		return entity.Representation{}, fmt.Errorf("your id is invalid")
	}

	//Read and Validate Name
	fmt.Println("Please enter your new name")
	scanner.Scan()
	newName := scanner.Text()

	if !validate.Name(newName) {

		return entity.Representation{}, fmt.Errorf("your name is invalid")
	}

	//Read and Validate Address
	fmt.Println("Please enter your new address")
	scanner.Scan()
	newAddress := scanner.Text()

	if !validate.Address(newAddress) {

		return entity.Representation{}, fmt.Errorf("your address is invalid")

	}

	//Read and Validate Phone Number
	fmt.Println("Please enter your new phoneNumber")
	scanner.Scan()
	newPhoneNumber := scanner.Text()

	if !validate.PhoneNumber(newPhoneNumber) {

		return entity.Representation{}, fmt.Errorf("your phonenumber is invalid")
	}

	//Read and Validate Staff Number
	fmt.Println("Please enter your new staffNumber")
	scanner.Scan()
	newStaffNumberString := scanner.Text()

	var newStaffNumber int
	if !validate.StaffNumber(newStaffNumberString) {

		return entity.Representation{}, fmt.Errorf("your staffnumber is invalid")
	} else {
		newStaffNumber, _ = strconv.Atoi(newStaffNumberString)
	}

	//Make Response
	old, gErr := memory.GetRepByID(id)
	if gErr != nil {
		return entity.Representation{}, fmt.Errorf("cant get representation by id")
	}

	editedRep := entity.Representation{
		ID:           id,
		StaffNumber:  newStaffNumber,
		Name:         newName,
		Region:       region,
		PhoneNumber:  newPhoneNumber,
		Address:      newAddress,
		RegisterDate: old.RegisterDate,
	}

	return editedRep, nil
}

func (r Reader) GetID(memory userservice.MemoryStorage) (int, error) {
	fmt.Println("Please enter your id")
	scanner.Scan()
	idString := scanner.Text()
	id, iErr := strconv.Atoi(idString)
	if iErr != nil {
		return 0, fmt.Errorf("cant call strconv")
	}

	if !validate.ID(id, memory.Last()) {
		return 0, fmt.Errorf("id is ivalid")
	}

	return id, nil
}

func (r Reader) CommandRegion() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter your command")
	scanner.Scan()
	command := scanner.Text()

	fmt.Println("Pleases enter your region")
	scanner.Scan()
	region := scanner.Text()

	return command, region
}
