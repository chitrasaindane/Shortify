package utils

import (
	"errors"
	"time"
)

// # Get 'Current System Time' in 'IST' Timezone
func GetCurrentSystemTimeInIST() (*time.Time, error) {
	var err error
	var errMsg string

	// # Get the 'current system time'
	currentSystemTime := time.Now()

	// # Get the 'IST' timezone
	ISTTimezone := "Asia/Kolkata"

	// # Load the 'IST' timezone 'location'
	ISTTimezoneLocation, err := time.LoadLocation(ISTTimezone)
	if err != nil {
		LogError(err, "Utils.GetCurrentSystemTimeInIST")
		errMsg = "Failed to load the IST timezone location"
		err = errors.New(errMsg)
		return nil, err
	}

	// # Convert the 'current system time' to the 'IST' timezone
	ISTTime := currentSystemTime.In(ISTTimezoneLocation)

	// # Return the 'IST' time
	return &ISTTime, nil
}
