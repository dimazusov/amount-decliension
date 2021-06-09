package amount_decliension

type Declension struct {
	DeclensionOne  string
	DeclensionSome string
	DeclensionMany string
	Gender         bool
}

WordZero     = "ноль"
WordHundreds = map[int]string{
	1: "сто",
	2: "двести",
	3: "триста",
	4: "четыреста",
	5: "пятьсот",
	6: "шестьсот",
	7: "семьсот",
	8: "восемьсот",
	9: "девятьсот",
}
WordDozens = map[int]string{
	1: "десять",
	2: "двадцать",
	3: "тридцать",
	4: "сорок",
	5: "пятьдесят",
	6: "шестьдесят",
	7: "семьдесят",
	8: "восемьдесят",
	9: "девяносто",
}
WordUnitFemale = map[int]string{
	1: "одна",
	2: "две",
	3: "три",
	4: "четыре",
	5: "пять",
	6: "шесть",
	7: "семь",
	8: "восемь",
	9: "девять",
}
WordUnitMale = map[int]string{
	1: "один",
	2: "два",
	3: "три",
	4: "четыре",
	5: "пять",
	6: "шесть",
	7: "семь",
	8: "восемь",
	9: "девять",
}
WordTeen = map[int]string{
	1: "одиннадцать",
	2: "двенадцать",
	3: "тринадцать",
	4: "четырнадцать",
	5: "пятнадцать",
	6: "шестнадцать",
	7: "семнадцать",
	8: "восемнадцать",
	9: "девятнадцать",
}
NumberParts = map[int]NumberPart{
	0: NumberPart{
		IsDeclRequired: true,
		DeclensionOne:  "рубль",
		DeclensionSome: "рубля",
		DeclensionMany: "рублей",
		Gender:         GENDER_MALE,
	},
	1: NumberPart{
		DeclensionOne:  "тысяча",
		DeclensionSome: "тысячи",
		DeclensionMany: "тысяч",
		Gender:         GENDER_FEMALE,
	},
	2: NumberPart{
		DeclensionOne:  "миллион",
		DeclensionSome: "миллиона",
		DeclensionMany: "миллионов",
		Gender:         GENDER_MALE,
	},
	3: NumberPart{
		DeclensionOne:  "миллиард",
		DeclensionSome: "миллиарда",
		DeclensionMany: "миллиардов",
		Gender:         GENDER_MALE,
	},
}
FractionalPart = NumberPart{
	DeclensionOne:  "копейка",
	DeclensionSome: "копейки",
	DeclensionMany: "копеек",
	Gender:         GENDER_FEMALE,
}
