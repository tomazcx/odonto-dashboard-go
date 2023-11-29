package tableutils

import "math"

func CalculatePagination(totalEntities int, quantityPerPage int) int {
	numOfPages := math.Ceil(float64((totalEntities + quantityPerPage - 1) / quantityPerPage))
	return int(numOfPages)
}
