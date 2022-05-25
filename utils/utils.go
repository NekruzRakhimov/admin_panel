package utils

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func ConvertTime(date string) (string, error) {
	timeSplit := strings.Split(date, ".")
	if len(timeSplit) != 3 {
		return "", errors.New("len of time must be 3")
	}
	fmt.Println(timeSplit)
	convertYear, err := strconv.Atoi(timeSplit[2])
	if err != nil {
		log.Println(err)
		return "", err
	}
	convertYear -= 1
	updateTime := fmt.Sprintf("%s.%s.%d", timeSplit[0], timeSplit[1], convertYear)
	//fmt.Println(sprintf)

	return updateTime, nil
}

func ReverseString(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func FloatToMoneyFormat(numIn float64) (numOut string) {
	counter := 0
	temp := fmt.Sprintf("%.2f", numIn)
	fmt.Println(temp)
	temp = ReverseString(temp)

	for i := 0; i < len(temp); i++ {
		if i < 3 {
			numOut = string(temp[i]) + numOut
			continue
		}

		if counter == 3 {
			numOut = " " + numOut
			counter = 0
		}

		numOut = string(temp[i]) + numOut
		counter++

		//fmt.Println("i: ", i)
		//fmt.Println("	temp[i]: ", string(temp[i]))
		//fmt.Println("	counter: ", counter)
		//fmt.Println("	numOut: ", numOut)
	}

	return numOut
}
