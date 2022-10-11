package szrecord

import (
	"testing"
)

// ----------------------------------------------------------------------------
// Test the NewRecord function
func TestNewRecord_good(t *testing.T) {
	jsonLine := `{"DATA_SOURCE": "ICIJ", "RECORD_ID": "24000001", "ENTITY_TYPE": "ADDRESS", "RECORD_TYPE": "ADDRESS", "icij_source": "BAHAMAS", "icij_type": "ADDRESS", "COUNTRIES": [{"COUNTRY_OF_ASSOCIATION": "BHS"}], "ADDR_FULL": "ANNEX FREDERICK & SHIRLEY STS, P.O. BOX N-4805, NASSAU, BAHAMAS", "REL_ANCHOR_DOMAIN": "ICIJ_ID", "REL_ANCHOR_KEY": "24000001"}`
	record, err := NewRecord(jsonLine)
	if err != nil {
		t.Errorf("FAILED, received err: %s", err.Error())
	} else if record.Id != "24000001" {
		t.Errorf("FAILED, Id incorrect")
	} else if record.DataSource != "ICIJ" {
		t.Errorf("FAILED, DataSource incorrect")
	} else if record.Json != jsonLine {
		t.Errorf("FAILED, jsonLine incorrect")
	} else {
		t.Log("SUCCEDED. record created.")
	}
}

func TestNewRecord_invalidJson(t *testing.T) {
	jsonLine := `{"DATA_SOURCE": "ICIJ", "RECORD_ID": "24000005B" "ENTITY_TYPE": "ADDRESS", "RECORD_TYPE": "ADDRESS", "icij_source": "BAHAMAS", "icij_type": "ADDRESS", "COUNTRIES": [{"COUNTRY_OF_ASSOCIATION": "BHS"}], "ADDR_FULL": "LYFORD CAY HOUSE, 3RD FLOOR, LYFORD CAY, P.O. BOX N-3024, NASSAU, BAHAMAS", "REL_ANCHOR_DOMAIN": "ICIJ_ID", "REL_ANCHOR_KEY": "24000005"}`
	_, err := Validate(jsonLine)

	if err != nil {
		t.Logf("SUCCEDED, received err: %s", err.Error())
	} else {
		t.Error("FAILED, expected err.")
	}
}

// ----------------------------------------------------------------------------
// Test the Validate function
func TestValidate_good(t *testing.T) {
	jsonLine := `{"DATA_SOURCE": "ICIJ", "RECORD_ID": "24000001", "ENTITY_TYPE": "ADDRESS", "RECORD_TYPE": "ADDRESS", "icij_source": "BAHAMAS", "icij_type": "ADDRESS", "COUNTRIES": [{"COUNTRY_OF_ASSOCIATION": "BHS"}], "ADDR_FULL": "ANNEX FREDERICK & SHIRLEY STS, P.O. BOX N-4805, NASSAU, BAHAMAS", "REL_ANCHOR_DOMAIN": "ICIJ_ID", "REL_ANCHOR_KEY": "24000001"}`
	val, err := Validate(jsonLine)
	if err != nil {
		t.Errorf("FAILED, received err: %s", err.Error())
	}
	if !val {
		t.Error("FAILED, expected JSON to validate.")
	} else {
		t.Log("SUCCEDED. valid record.")
	}
}

func TestValidate_noRecordId(t *testing.T) {
	jsonLine := `{"DATA_SOURCE": "ICIJ", "ENTITY_TYPE": "ADDRESS", "RECORD_TYPE": "ADDRESS", "icij_source": "BAHAMAS", "icij_type": "ADDRESS", "COUNTRIES": [{"COUNTRY_OF_ASSOCIATION": "BHS"}], "ADDR_FULL": "ANNEX FREDERICK & SHIRLEY STS, P.O. BOX N-4805, NASSAU, BAHAMAS", "REL_ANCHOR_DOMAIN": "ICIJ_ID", "REL_ANCHOR_KEY": "24000001"}`
	val, err := Validate(jsonLine)

	if err != nil {
		t.Logf("SUCCEDED, received err: %s", err.Error())
	} else {
		t.Error("FAILED, expected err.")
	}
	if val {
		t.Error("FAILED, did NOT expected JSON to validate.")
	} else {
		t.Log("SUCCEDED, invalid record")
	}
}

func TestValidate_noDatasource(t *testing.T) {
	jsonLine := `{"RECORD_ID": "24000001", "ENTITY_TYPE": "ADDRESS", "RECORD_TYPE": "ADDRESS", "icij_source": "BAHAMAS", "icij_type": "ADDRESS", "COUNTRIES": [{"COUNTRY_OF_ASSOCIATION": "BHS"}], "ADDR_FULL": "ANNEX FREDERICK & SHIRLEY STS, P.O. BOX N-4805, NASSAU, BAHAMAS", "REL_ANCHOR_DOMAIN": "ICIJ_ID", "REL_ANCHOR_KEY": "24000001"}`
	val, err := Validate(jsonLine)

	if err != nil {
		t.Logf("SUCCEDED, received err: %s", err.Error())
	} else {
		t.Error("FAILED, expected err.")
	}
	if val {
		t.Error("FAILED, did NOT expected JSON to validate.")
	} else {
		t.Log("SUCCEDED, invalid record")
	}
}

func TestValidate_invalidJson(t *testing.T) {
	jsonLine := `{"DATA_SOURCE": "ICIJ", "RECORD_ID": "24000005B" "ENTITY_TYPE": "ADDRESS", "RECORD_TYPE": "ADDRESS", "icij_source": "BAHAMAS", "icij_type": "ADDRESS", "COUNTRIES": [{"COUNTRY_OF_ASSOCIATION": "BHS"}], "ADDR_FULL": "LYFORD CAY HOUSE, 3RD FLOOR, LYFORD CAY, P.O. BOX N-3024, NASSAU, BAHAMAS", "REL_ANCHOR_DOMAIN": "ICIJ_ID", "REL_ANCHOR_KEY": "24000005"}`
	val, err := Validate(jsonLine)

	if err != nil {
		t.Logf("SUCCEDED, received err: %s", err.Error())
	} else {
		t.Error("FAILED, expected err.")
	}
	if val {
		t.Error("FAILED, did NOT expected JSON to validate.")
	} else {
		t.Log("SUCCEDED, invalid record")
	}
}
