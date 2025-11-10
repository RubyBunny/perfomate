package qapair

type QAPairRepository struct {
	MarkedQuestions   []MarkedQAPair
	UnmarkedQuestions []QAPair
}

type QAPair struct {
	Question string
	Answer   string
}

type MarkedQAPair struct {
	QAPair
	Mark float64
}

func NewMarkedQAPair(question, answer, markChoose string) MarkedQAPair {
	return MarkedQAPair{QAPair{question, answer}, markChoose2Mark(markChoose)}
}

func markChoose2Mark(markChoose string) float64 {
	switch markChoose {
	case "Полностью согласен(а)", "Полностью удовлетворен(а)":
		return 5
	case "Скорее согласен(а)", "Скорее удовлетворен(а)":
		return 4
	case "Скорее не согласен(а)", "Скорее неудовлетворен(а)":
		return 3
	case "Совершенно не согласен(а)", "Совершенно не удовлетворен(а)":
		return 2
	case "Не могу оценить":
		fallthrough
	default:
		return 0
	}
}
