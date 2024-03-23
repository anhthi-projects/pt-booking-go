package models

import (
	"time"

	"github.com/uptrace/bun"
)

type TrainerModel struct {
	bun.BaseModel 					`bun:"table:trainers"`
	ID            int 			`bun:"id,pk,autoincrement"`
	FirstName     string 		`bun:"first_name,notnull"`
	LastName      string 		`bun:"last_name,notnull"`
	DateOfBirth   time.Time `bun:"date_of_birth,notnull"`
	Level   		  string 		`bun:"level"`
	YearOfExp		  int 			`bun:"year_of_exp"`
}