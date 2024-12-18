package helper

import "time"

func CoalesceInt(newValue, oldValue int) int {
	if newValue != 0 {
		return newValue
	}
	return oldValue
}

func CoalesceString(newValue, oldValue string) string {
	if newValue != "" {
		return newValue
	}
	return oldValue
}

func CoalesceBoolPtr(newValue, oldValue *bool) *bool {
	if newValue != nil {
		return newValue
	}
	return oldValue
}

func CoalesceTime(newValue, oldValue time.Time) time.Time {
	if !newValue.IsZero() {
		return newValue
	}
	return oldValue
}
