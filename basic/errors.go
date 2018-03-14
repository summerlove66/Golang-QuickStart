package basic

func divsion(x, y float64) (error, float64) {
	if y != 0 {
		return nil, y / x
	} else {
		return divsionError{}, 0
	}
}

type divsionError struct {
}

func (d divsionError) Error() string {
	return "divsion error because of 0"
}
