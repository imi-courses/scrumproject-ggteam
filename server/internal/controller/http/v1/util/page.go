package util

import (
	"errors"
	"strconv"
)

func GetPage(page string, count string) (int, int, error) {
	var currentPage, currentCount int
	var err error
	if page == "" {
		currentPage = 1
	} else {
		currentPage, err = strconv.Atoi(page)
		if err != nil {
			return 1, 10, err
		}
		if currentPage < 1 {
			return 1, 10, errors.New("page must be greater than 0")
		}
	}
	if count == "" {
		currentCount = 10
	} else {
		currentCount, err = strconv.Atoi(count)
		if err != nil {
			return currentPage, 10, err
		}
		if currentCount < 1 {
			return currentPage, 10, errors.New("count must be greater than 0")
		}
	}

	return currentPage, currentCount, nil
}
