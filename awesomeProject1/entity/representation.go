package entity

import "time"

type Representation struct {
	ID           int
	StaffNumber  int
	Name         string
	Region       string
	PhoneNumber  string
	Address      string
	RegisterDate time.Time
}
