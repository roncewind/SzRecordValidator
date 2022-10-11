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
func NewRecord(line string) (*Record, error) {
	var record Record
	err := json.Unmarshal([]byte(line), &record)
	if err == nil {
		record.Json = line
		return &record, nil
	}
	//TODO:  should we return err ???
	return &record, errors.New("JSON-line not well formed.")
}

// ----------------------------------------------------------------------------
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
func ValidateRecord(record Record) (bool, error) {
	// FIXME: errors should be specific to the input method
	//  ala rabbitmq message ID?
	if record.DataSource == "" {
		return false, errors.New("A DATA_SOURCE field is required.")
	}
	if record.Id == "" {
		return false, errors.New("A RECORD_ID field is required.")
	}
	return true, nil
}
