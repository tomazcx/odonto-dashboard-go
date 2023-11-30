package tableutils

import (
	"errors"
	"math"
	"strings"
)

func CalculatePagination(totalEntities int, quantityPerPage int) int {
	numOfPages := math.Ceil(float64((totalEntities + quantityPerPage - 1) / quantityPerPage))
	return int(numOfPages)
}

func GetFieldAndAscending(str string, field *string, ascendingSort *bool) error {

	if len(str) == 0 {
		*field = ""
		*ascendingSort = true
		return nil
	}

	parts := strings.Split(str, "/")

	if len(parts) != 2 {
		return errors.New("Opção de ordenação inválida")
	}

	*field = parts[0]

	if parts[1] == "descending" {
		*ascendingSort = false
	} else {
		*ascendingSort = true
	}

	return nil
}
