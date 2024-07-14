package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	_, ok := fnb.(FancyNumber)
	if !ok {
		return 0
	}

	i, _ := strconv.Atoi(fnb.Value())
	return i
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	_, ok := fnb.(FancyNumber)
	if ok {
		fancyFloat, _ := strconv.ParseFloat(fnb.Value(), 64)
		return fmt.Sprintf("This is a fancy box containing the number %.1f", fancyFloat)
	}
	return "This is a fancy box containing the number 0.0"
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {

	fancyNum, okFancyNum := i.(FancyNumber)
	if okFancyNum {
		return DescribeFancyNumberBox(fancyNum)
	}

	return "Return to sender"
}
