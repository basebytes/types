package types

import (
	"database/sql/driver"
	"errors"
)

type Json []byte

func (j *Json) Value() (driver.Value, error) {
	if j == nil || len(*j) == 0 || string(*j) == "null" {
		return nil, nil
	}
	return string(*j), nil
}

func (j *Json) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	if s, OK := value.([]byte); OK {
		*j = append((*j)[0:0], s...)
		return nil
	} else {
		return errors.New("Invalid Scan Source ")
	}
}

func (j *Json) MarshalJSON() ([]byte, error) {
	if *j == nil {
		return []byte{}, nil
	}
	return *j, nil
}

func (j *Json) UnmarshalJSON(data []byte) (err error) {
	if *j == nil {
		*j = []byte{}
	}
	*j = append((*j)[0:0], data...)
	return
}
