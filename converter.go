package amount_decliension

import (
	"fmt"
	"strconv"
	"strings"
)

type declensioner struct {}

func NewDeclensioner() *declensioner{
	return &declensioner{}
}

func (m *declensioner) AmountToString(totalSum float64) (stringAmount string) {
	partOfNumbers := strings.Split(strconv.FormatFloat(totalSum, 'f', 2, 64), ".")

	intPart, _ := strconv.Atoi(partOfNumbers[0])
	stringAmount = m.GetNumberInWords(float64(intPart))

	if len(partOfNumbers) == 2 {
		stringAmount += " " + partOfNumbers[1]

		fractalPart, _ := strconv.Atoi(partOfNumbers[1])
		if fractalPart >= 11 && fractalPart <= 19 {
			stringAmount += " " + FractionalPart.DeclensionMany

			return
		}

		lastDigit := string(partOfNumbers[1][1])

		if m.inArray(lastDigit, []string{"2", "3", "4"}) {
			stringAmount += " " + FractionalPart.DeclensionSome
		} else if lastDigit == "1" {
			stringAmount += " " + FractionalPart.DeclensionOne
		} else {
			stringAmount += " " + FractionalPart.DeclensionMany
		}
	}

	return stringAmount
}

func (m *declensioner) GetNumberInWords(num float64) (numInWords string) {
	partOfNumbers := strings.Split(strconv.FormatFloat(num, 'f', 2, 64), ".")
	intPart, _ := strconv.Atoi(partOfNumbers[0])

	if intPart != 0 {
		parts := m.splitStringByLength(fmt.Sprint(intPart), 3)
		strNumParts := []string{}

		for i, part := range m.reverse(parts) {
			numPart := NumberParts[i]
			numPart.Init(part)

			if isEmpty, strNumPart := numPart.ToString(); !isEmpty {
				strNumParts = append(strNumParts, strNumPart)
			}
		}

		numInWords = strings.Join(m.reverse(strNumParts), " ")
	} else {
		numInWords = WordZero + " " + NumberParts[0].DeclensionMany
	}

	if len(partOfNumbers) == 2 {
		parts := m.splitStringByLength(partOfNumbers[1], 2)
		fractionalPart := FractionalPart
		fractionalPart.Init(parts[0])

		isEmpty, strNumPart := fractionalPart.ToString()
		if !isEmpty {
			if len(numInWords) != 0 {
				numInWords += " "
			}
			numInWords += strNumPart
		}
	}

	return
}

func (m *declensioner) splitStringByLength(str string, count int) []string {
	strLen := len(str)
	parts := []string{}

	for i := strLen; i >= 0; i -= count {
		if i == 0 {
			continue
		}

		if i-count < 0 {
			parts = append(parts, str[0:i])
			continue
		}

		parts = append(parts, str[i-count:i])
	}

	return m.reverse(parts)
}


func (m *declensioner) inArray(val string, array []string) bool {
	for _, v := range array {
		if val == v {
			return true
		}
	}

	return false
}

func (m *declensioner) reverse(numbers []string) []string {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}