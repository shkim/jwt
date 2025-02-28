package jwt_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/shkim/jwt/v4"
)

func TestNumericDate(t *testing.T) {
	var s struct {
		Iat jwt.NumericDate `json:"iat"`
		Exp jwt.NumericDate `json:"exp"`
	}

	oldPrecision := jwt.TimePrecision

	jwt.TimePrecision = time.Microsecond

	raw := `{"iat":1516239022.000000,"exp":1516239022.123450}`

	if err := json.Unmarshal([]byte(raw), &s); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	b, _ := json.Marshal(s)

	if raw != string(b) {
		t.Errorf("Serialized format of numeric date mismatch. Expecting: %s  Got: %s", string(raw), string(b))
	}

	jwt.TimePrecision = oldPrecision
}

func TestSingleArrayMarshal(t *testing.T) {
	jwt.MarshalSingleStringAsArray = false

	s := jwt.ClaimStrings{"test"}
	expected := `"test"`

	b, err := json.Marshal(s)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if expected != string(b) {
		t.Errorf("Serialized format of string array mismatch. Expecting: %s  Got: %s", string(expected), string(b))
	}

	jwt.MarshalSingleStringAsArray = true

	expected = `["test"]`

	b, err = json.Marshal(s)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if expected != string(b) {
		t.Errorf("Serialized format of string array mismatch. Expecting: %s  Got: %s", string(expected), string(b))
	}
}

func TestNumericDate_MarshalJSON(t *testing.T) {
	// Do not run this test in parallel because it's changing
	// global state.
	oldPrecision := jwt.TimePrecision
	t.Cleanup(func() {
		jwt.TimePrecision = oldPrecision
	})

	tt := []struct {
		in        time.Time
		want      string
		precision time.Duration
	}{
		{time.Unix(5243700879, 0), "5243700879", time.Second},
		{time.Unix(5243700879, 0), "5243700879.000", time.Millisecond},
		{time.Unix(5243700879, 0), "5243700879.000001", time.Microsecond},
		{time.Unix(5243700879, 0), "5243700879.000000954", time.Nanosecond},
		//
		{time.Unix(4239425898, 0), "4239425898", time.Second},
		{time.Unix(4239425898, 0), "4239425898.000", time.Millisecond},
		{time.Unix(4239425898, 0), "4239425898.000000", time.Microsecond},
		{time.Unix(4239425898, 0), "4239425898.000000000", time.Nanosecond},
		//
		{time.Unix(0, 1644285000210402000), "1644285000", time.Second},
		{time.Unix(0, 1644285000210402000), "1644285000.210", time.Millisecond},
		{time.Unix(0, 1644285000210402000), "1644285000.210402", time.Microsecond},
		{time.Unix(0, 1644285000210402000), "1644285000.210402012", time.Nanosecond},
		//
		{time.Unix(0, 1644285315063096000), "1644285315", time.Second},
		{time.Unix(0, 1644285315063096000), "1644285315.063", time.Millisecond},
		{time.Unix(0, 1644285315063096000), "1644285315.063096", time.Microsecond},
		{time.Unix(0, 1644285315063096000), "1644285315.063096046", time.Nanosecond},
	}

	for i, tc := range tt {
		jwt.TimePrecision = tc.precision
		by, err := jwt.NewNumericDate(tc.in).MarshalJSON()
		if err != nil {
			t.Fatal(err)
		}
		if got := string(by); got != tc.want {
			t.Errorf("[%d]: failed encoding: got %q want %q", i, got, tc.want)
		}
	}
}
