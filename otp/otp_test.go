package otp

import (
	"reflect"
	"testing"
)

func TestGoogleAuthenticatorCode(t *testing.T) {
	test_cases := []struct {
		Name          string
		Input         Author
		ExpectedCalls string
	}{
		{
			"Golden test",
			Author{"some data with \x00 and \ufeff", 1},
			"ONXW2ZJAMRQXIYJAO5UXI2BAAAQGC3TEEDX3XPY=",
		},
	}

	for _, test := range test_cases {
		t.Run(test.Name, func(t *testing.T) {
			got := GoogleAuthenticatorCode(test.Input)

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("Not deep equal, got %v expected %v", got, test.ExpectedCalls)
			}
		})
	}

}
