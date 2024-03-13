package models

import (
	"time"

	"github.com/uptrace/bun"
)

type TrainerModel struct {
	bun.BaseModel 					`bun:"table:trainers"`
	ID            int 			`bun:"id,pk,autoincrement"`
	FirstName     string 		`bun:"firstName,notnull"`
	LastName      string 		`bun:"lastName,notnull"`
	DateOfBirth   time.Time `bun:"dateOfBirth,notnull"`
	Level   		  string 		`bun:"level"`
	YearOfExp		  int 			`bun:"yearOfExp"`
}