package szrecord

import (
	"encoding/json"
	"errors"
)

// ----------------------------------------------------------------------------
type Record struct {
	DataSource string `json:"DATA_SOURCE"`
	Id string `json:"RECORD_ID"`
	Json string
}

// ----------------------------------------------------------------------------
// returns a valid Record or an error if validation fails
func NewRecord(line string) (*Record, error) {
	var record Record
	err := json.Unmarshal([]byte(line), &record)
	if err == nil {
		record.Json = line
		_, validationErr := ValidateRecord(record)
		if validationErr == nil {
			return &record, nil
		} else {
			return &record, validationErr
		}
	}
	//TODO:  should we return err ???
	return &record, errors.New("JSON-line not well formed.")
}

// ----------------------------------------------------------------------------
// a string is only a valid Record, if it is a well formed JSONLine
// and it has a DataSource field
// and it has an Id field
func Validate(line string) (bool, error) {
	var record Record
	valid := json.Unmarshal([]byte(line), &record) == nil
	if valid {
		return ValidateRecord(record)
	}
	//TODO: should we return the actual parse error???
	return valid, errors.New("JSON-line not well formed.")
}

// ----------------------------------------------------------------------------
// a Record is only valid if it has a DataSource field
// and it has an Id field
func ValidateRecord(record Record) (bool, error) {

	if record.DataSource == "" {
		return false, errors.New("A DATA_SOURCE field is required.")
	}
	if record.Id == "" {
		return false, errors.New("A RECORD_ID field is required.")
	}
	return true, nil
}
