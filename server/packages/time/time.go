package time

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ParseLargeDuration acts like time.ParseDuration but can also handle D (day), W (week), M (month = 30D), Y (year = 365D)
// TODO: test this function
func ParseLargeDuration(s string) (time.Duration, error) {
	// TODO: verify the string with a regex first
	// Extract the values we are interested in using regex
	regex := regexp.MustCompile(`(\d+[WYDM])`)
	extractedValues := regex.FindAllString(s, -1)

	// remove the extractedValues from s
	for _, v := range extractedValues {
		s = strings.Replace(s, v, "", 1)
	}

	// Pass the others to the time.ParseDuration function
	var d time.Duration
	var err error
	if s != "" {
		d, err = time.ParseDuration(s)
		if err != nil {
			return d, err
		}
	}

	// add our values to the duration
	for _, v := range extractedValues {
		cutP := len(v) - 1
		valueS, timeTypeC := v[:cutP], v[cutP]

		var timeType time.Duration
		switch timeTypeC {
		case 'D':
			timeType = 24 * time.Hour
		case 'W':
			timeType = 7 * 24 * time.Hour
		case 'M':
			timeType = 30 * 24 * time.Hour
		case 'Y':
			timeType = 365 * 24 * time.Hour
		}

		valueI, err := strconv.ParseInt(valueS, 10, 64)
		if err != nil {
			return time.Nanosecond, fmt.Errorf("cannot parse value %s", valueS)
		}
		d += time.Duration(valueI) * timeType
	}

	return d, nil
}
