package dtos

import "time"

type TrainerDTO struct {
	ID          int    		`json:"id"`
	FirstName   string 		`json:"firstName"`
	LastName    string 		`json:"lastName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Level   		string 		`json:"level"`
	YearOfExp		int 			`json:"yearOfExp"`
}
