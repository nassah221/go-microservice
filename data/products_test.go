package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{Name: "Haanss", Price: 1.0, SKU: "abc-abc-abc"}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
