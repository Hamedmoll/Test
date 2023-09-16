package validate

import (
	"awesomeProject1/entity"
	"strconv"
)

func Name(name string) bool {
	if len(name) < 3 {

		return false
	}

	return true
}

func Address(address string) bool {
	if len(address) == 0 {

		return false
	}

	return true
}

func PhoneNumber(phoneNumber string) bool {
	//TODO - VERIFY +98 FORMAT
	if len(phoneNumber) != 11 {

		return false
	}
	if phoneNumber[:2] != "09" {

		return false
	}

	return true
}

func StaffNumber(staffNumber string) bool {
	number, err := strconv.Atoi(staffNumber)
	if err != nil || number < 0 {

		return false
	}

	return true
}

func ID(id int, last int) bool {
	if id < 0 || id > last {

		return false
	}

	return true
}

func MatchIDRegion(rep entity.Representation, region string) bool {
	if rep.Region != region {

		return false
	}

	return true
}
