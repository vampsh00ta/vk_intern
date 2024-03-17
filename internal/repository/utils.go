package repository

func sortStatement(sortBy, orderBy string) string {
	var sortVal string
	sortVal = sortBy
	if sortVal == "" {
		sortVal = "rating"
	}
	if orderBy == "asc" {
		sortVal += " asc"
	} else {
		sortVal += " desc"

	}
	return sortVal
}
