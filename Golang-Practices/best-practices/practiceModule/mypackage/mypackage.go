package person

import (
	"strconv"
)

var Name string = "Anik Adnan"

func IntArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, i := range intArr {
		strArr = append(strArr, strconv.Itoa(i))
	}
	return strArr

}
