package typeconversions

import (
	"strconv"
)

func ConverToID(id string) (uint, error) {
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return 0, err
	}

	return uint(idInt), nil
}
