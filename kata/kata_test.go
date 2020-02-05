package kata_test

import (
	"test/kata"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_CalculateAge_Input_Birth_Date_18042003_Should_be_16(t *testing.T) {
	expectedAge := 16
	birthDate, _ := time.Parse("02012006", "18042003")
	tranfer := kata.Tranfer{
		Now: func() time.Time {
			fixedTime, _ := time.Parse("02012006", "30012020")
			return fixedTime
		},
	}

	actual := tranfer.CalculateAge(birthDate)

	assert.Equal(t, expectedAge, actual)
}

func Test_CalulateFee_Input_User_BirtDate_18042003_And_Amount_Should_Be_Fee_0(t *testing.T) {
	expectedFee := 00.0
	amount := 5000.00
	birthDate, _ := time.Parse("02012006", "18042003")
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

func Test_CheckLimitPerDay_Input_Amount_500_And_History_Deposit_50000_Limit_Per_Day_100000_Should_Be_False(t *testing.T) {
	expected := false
	amount := 500.00
	historyDeposit := 50000.00
	limitPerDay := 100000.00

	actual := kata.CheckLimitPerDay(amount, historyDeposit, limitPerDay)

	assert.Equal(t, expected, actual)
}
