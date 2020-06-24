package controllers

import (
	"fmt"
	"strconv"
)

func convertStringToInt(routeParam string) uint64 {
	var err error
	var id uint64
	id, err = strconv.ParseUint(routeParam, 10, 32)

	fmt.Println(err)
	return id
}
