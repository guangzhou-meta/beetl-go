package core

type PlaceHolderDelimeter struct {
	*Delimeter
}

func NewPlaceHolderDelimeter(start []rune, end []rune, start1 *[]rune, end1 *[]rune) *PlaceHolderDelimeter {
	inst := &PlaceHolderDelimeter{
		Delimeter: NewDelimeter(start, end, start1, end1),
	}
	return inst
}
