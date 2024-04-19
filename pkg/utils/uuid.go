package utils

import (
	"database/sql/driver"
	"errors"

	"github.com/GoWebProd/uuid7"
)

type UUID uuid7.UUID

var uuid7Gen = uuid7.New()

func GenUUID() UUID {
	return UUID(uuid7Gen.Next())
}
func (s UUID) Scan(value interface{}) error {

	if bv, err := driver.String.ConvertValue(value); err == nil {
		// if this is a bool type
		if v, ok := bv.(string); ok {
			// set the value of the pointer yne to YesNoEnum(v)
			s = UUID([]byte(v))
			return nil
		}
	}
	// otherwise, return an error
	return errors.New("failed to scan YesNoEnum")

}

func (s *UUID) Value() (driver.Value, error) {

	return string(s[:]), nil
}
