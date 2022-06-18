package logic

import "testing"

func TestFindCountry(t *testing.T) {

	tests := []struct {
		mobileNum   string
		isNil       bool
		message     string
		countryName string
	}{
		{
			mobileNum:   "(212) 654642448",
			isNil:       false,
			message:     "Valid country number",
			countryName: "Morocco",
		},
		{
			mobileNum:   "21(2 6007989253",
			isNil:       true,
			message:     "Invalid country number",
			countryName: "Morocco",
		},
		{
			mobileNum:   "(2120) 698054317",
			isNil:       true,
			message:     "Invalid country number",
			countryName: "Morocco",
		},
	}
	for _, tt := range tests {
		t.Run(tt.message, func(t *testing.T) {
			got, _ := FindCountry(tt.mobileNum)
			if (got == nil) != tt.isNil {
				t.Errorf("Check if FindCountry() = %v, isNil %v", got, tt.isNil)
			}
		})
	}
}
