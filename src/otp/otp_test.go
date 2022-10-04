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
			Author{"dummySECRETdummy", 1523822557}, //Secret=dummySECRETdummy
			"563916",
		},
	}

	for _, test := range test_cases {
		t.Run(test.Name, func(t *testing.T) {
			got, _ := GoogleAuthenticatorCode(test.Input)

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("Not deep equal, got %v expected %v", got, test.ExpectedCalls)
			}
		})
	}

}
