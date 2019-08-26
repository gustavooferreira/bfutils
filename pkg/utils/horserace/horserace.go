package horserace

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func DistanceToFurlongs(distance string) (uint, error) {

	// # Use regex to validate input.
	match, _ := regexp.MatchString("([1-7]m[1-7]f)|([1-7][mf])", distance)
	if !match {
		return 0, errors.New("distance not properly formatted")
	}

	m_f := false
	f_f := false

	if strings.Index(distance, "m") != -1 {
		m_f = true
	}

	if strings.Index(distance, "f") != -1 {
		f_f = true
	}

	var m int
	var f int

	if m_f && f_f {
		m, _ = strconv.Atoi(string(distance[0]))
		f, _ = strconv.Atoi(string(distance[2]))
	} else if m_f && !f_f {
		m, _ = strconv.Atoi(string(distance[0]))
	} else if !m_f && f_f {
		f, _ = strconv.Atoi(string(distance[0]))

	}

	return uint(8*m + f), nil
}

func DistanceFromFurlongs(distance uint) string {

	m := distance / 8
	f := distance % 8

	var result string

	if m == 0 {
		result = fmt.Sprintf("%df", f)
	} else if f == 0 {
		result = fmt.Sprintf("%dm", m)
	} else {
		result = fmt.Sprintf("%dm%df", m, f)
	}

	return result
}

// getNameAndDistance returns the name of the race, the distance
func GetNameAndDistance(marketName string) (name string, distance string) {
	words := strings.Fields(marketName)

	// check length as well
	distance = words[0]

	name = strings.Join(words[1:], " ")

	return name, distance
}
