package util

import (
	"cashbag-me-mini/models"
)

// HelperReversedArrayMilestone ...
func HelperReverseArrayMilestone(array []models.LoyaltyProgramMilestone) []models.LoyaltyProgramMilestone {
	reversed := []models.LoyaltyProgramMilestone{}
	for i:= range array {
		n := array[len(array)-1-i]
		reversed = append(reversed, n)
	}
	return reversed
}
