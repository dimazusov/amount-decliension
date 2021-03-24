package amount_decliension

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type amountToStringTestCase struct {
	Amount        float64
	StringAmount string
}

func TestDeclensioner_GetAmountInWords(t *testing.T) {
	declensioner := NewDeclensioner()

	testCases := getTestCases()
	for _, testCase := range testCases {
		givenAmount := declensioner.AmountToString(testCase.Amount)
		require.Equal(t, testCase.StringAmount, result)
	}
}

func getTestCases() []amountToStringTestCase {
	return []amountToStringTestCase{
		{0.01, "ноль рублей одна копейка"},
		{0.02, "ноль рублей две копейки"},
		{0.07, "ноль рублей семь копеек"},
		{0.1, "ноль рублей десять копеек"},
		{0.2, "ноль рублей двадцать копеек"},
		{0.17, "ноль рублей семнадцать копеек"},
		{0.23, "ноль рублей двадцать три копейки"},
		{1.23, "один рубль двадцать три копейки"},
		{1, "один рубль"},
		{2, "два рубля"},
		{3, "три рубля"},
		{4, "четыре рубля"},
		{5, "пять рублей"},
		{6, "шесть рублей"},
		{7, "семь рублей"},
		{8, "восемь рублей"},
		{9, "девять рублей"},
		{10, "десять рублей"},
		{11, "одиннадцать рублей"},
		{12, "двенадцать рублей"},
		{15, "пятнадцать рублей"},
		{19, "девятнадцать рублей"},
		{51, "пятьдесят один рубль"},
		{54, "пятьдесят четыре рубля"},
		{55, "пятьдесят пять рублей"},
		{71, "семьдесят один рубль"},
		{74, "семьдесят четыре рубля"},
		{75, "семьдесят пять рублей"},
		{79, "семьдесят девять рублей"},
		{91, "девяносто один рубль"},
		{94, "девяносто четыре рубля"},
		{100, "сто рублей"},
		{111, "сто одиннадцать рублей"},
		{134, "сто тридцать четыре рубля"},
		{221, "двести двадцать один рубль"},
		{675, "шестьсот семьдесят пять рублей"},
		{868, "восемьсот шестьдесят восемь рублей"},
		{1223, "одна тысяча двести двадцать три рубля"},
		{1868, "одна тысяча восемьсот шестьдесят восемь рублей"},
		{10868, "десять тысяч восемьсот шестьдесят восемь рублей"},
		{11000, "одиннадцать тысяч рублей"},
		{11521, "одиннадцать тысяч пятьсот двадцать один рубль"},
		{11543, "одиннадцать тысяч пятьсот сорок три рубля"},
		{11527, "одиннадцать тысяч пятьсот двадцать семь рублей"},
		{100000, "сто тысяч рублей"},
		{100001, "сто тысяч один рубль"},
		{100200, "сто тысяч двести рублей"},
		{101000, "сто одна тысяча рублей"},
		{1000000, "один миллион рублей"},
		{1101001, "один миллион сто одна тысяча один рубль"},
		{11000003, "одиннадцать миллионов три рубля"},
		{71000003, "семьдесят один миллион три рубля"},
		{111000003, "сто одиннадцать миллионов три рубля"},
		{222222222, "двести двадцать два миллиона двести двадцать две тысячи двести двадцать два рубля"},
		{4111000003, "четыре миллиарда сто одиннадцать миллионов три рубля"},
		{4111001003, "четыре миллиарда сто одиннадцать миллионов одна тысяча три рубля"},
		{4124375151.35, "четыре миллиарда сто двадцать четыре миллиона триста семьдесят пять тысяч сто пятьдесят один рубль тридцать пять копеек"},
	}
}