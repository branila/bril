package adder

import (
	"log"
	"regexp"
	"time"
)

func getExpiration(due string) time.Time {
	if due == "" {
		return time.Time{}
	}

	/*
	 *	Supported time formats:
	 *	  1. 13/11/2006-10:55
	 *	  2. 10:55-13/11/2006
	 *	  3. 13/11/2006
	 *	  4. 10:55
	 */
	regexp1 := regexp.MustCompile(`^\d{2}/\d{2}/\d{4}-\d{2}:\d{2}$`)
	regexp2 := regexp.MustCompile(`^\d{2}:\d{2}-\d{2}/\d{2}/\d{4}$`)
	regexp3 := regexp.MustCompile(`^\d{2}/\d{2}/\d{4}$`)
	regexp4 := regexp.MustCompile(`^\d{2}:\d{2}$`)

	var expiration time.Time
	var err error

	switch {
	case regexp1.MatchString(due):
		expiration, err = time.Parse("02/01/2006-15:04", due)

	case regexp2.MatchString(due):
		expiration, err = time.Parse("15:04-02/01/2006", due)

	case regexp3.MatchString(due):
		expiration, err = time.Parse("02/01/2006", due)

	case regexp4.MatchString(due):
		expiration, err = time.Parse("15:04", due)

	default:
		log.Fatal("Unsupported time format")
	}

	if err != nil {
		log.Fatal("Error ", err)
	}

	return expiration
}
