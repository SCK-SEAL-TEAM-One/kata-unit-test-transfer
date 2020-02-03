package kata_test

import (
	"test/kata"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_CalculateAge_Input_Birth_Date_18042011_Should_be_8(t *testing.T) {
	//preconditions (Arrange)
	expectedAge := 8                                   //expected result
	birthDate, _ := time.Parse("02012006", "18042011") //setup
	tranfer := kata.Tranfer{
		Now: func() time.Time {
			fixedTime, _ := time.Parse("02012006", "30012020")
			return fixedTime
		},
	}

	//action (Act)
	actual := tranfer.CalculateAge(birthDate)

	//validate result vs expect (Assert)
	assert.Equal(t, expectedAge, actual)
}

func Test_CalulateFee_Input_User_BirtDate_18042011_And_Amount_Should_Be_Fee_20(t *testing.T) {
	expectedFee := 20.0
	amount := 5000.00
	birthDate, _ := time.Parse("02012006", "18042011")
	userInfo := kata.UserInfo{
		UserID:    1,
		BirthDate: birthDate,
	}
	tranfer := kata.Tranfer{
		Now: func() time.Time {
			fixedTime, _ := time.Parse("02012006", "30012020")
			return fixedTime
		},
	}

	actual := tranfer.CalulateFee(userInfo, amount)

	assert.Equal(t, expectedFee, actual)
}

func Test_CalulateFee_Input_User_BirtDate_18041996_And_Amount_Should_Be_Fee_15(t *testing.T) {
	expectedFee := 15.0
	amount := 5000.00
	birthDate, _ := time.Parse("02012006", "18041996")
	userInfo := kata.UserInfo{
		UserID:    1,
		BirthDate: birthDate,
	}
	tranfer := kata.Tranfer{
		Now: func() time.Time {
			fixedTime, _ := time.Parse("02012006", "30012020")
			return fixedTime
		},
	}

	actual := tranfer.CalulateFee(userInfo, amount)

	assert.Equal(t, expectedFee, actual)
}
