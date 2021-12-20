package models

type completion struct {
	body struct {
		code           string
		fileType       string
		line           int
		column         int
		wordToComplete string
		offset         int
	}
}
