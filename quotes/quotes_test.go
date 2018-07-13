package quotes

import (
	"testing"
)

func TestGetJSON(t *testing.T) {

	json := GetJSON(1, 4)

	if json == "" {
		t.Fatalf("No json returned")
	}
}

func TestGetRawJSON(t *testing.T) {
	raw := GetJSONRaw(1, 4)

	if raw.Created == "" || len(raw.Data) <= 0 {
		t.Fatalf("No json returned")
	}
}
