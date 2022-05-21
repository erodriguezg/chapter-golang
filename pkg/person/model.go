package person

import "time"

type Person struct {
	Id        *int64
	Rut       int
	FirstName string
	LastName  string
	BirthDay  *time.Time
	Active    bool
}
