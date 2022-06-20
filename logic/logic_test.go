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
			mobileNum:   "(212) 814969671",
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
			mobileNum:   "(237)269309891",
			isNil:       false,
			message:     "Valid country number",
			countryName: "Cameroon",
		},
		{
			mobileNum:   "(237)3984797",
			isNil:       false,
			message:     "Invalid country number",
			countryName: "Cameroon",
		},
		{
			mobileNum:   "(258)239141337",
			isNil:       false,
			message:     "Valid country number",
			countryName: "Mozambique",
		},
		{
			mobileNum:   "(259)239141337",
			isNil:       true,
			message:     "Invalid country number",
			countryName: "Mozambique",
		},
		{
			mobileNum:   "(256)458116714",
			isNil:       false,
			message:     "Valid country number",
			countryName: "Uganda",
		},
		{
			mobileNum:   "(256)8260233",
			isNil:       false,
			message:     "Invalid country number",
			countryName: "Uganda",
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
