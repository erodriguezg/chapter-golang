package demosql

import "time"

type Person struct {
	Id        *int64
	FirstName string
	LastName  string
	BirthDay  *time.Time
	Active    bool
}
