package reviews

import "perfomate/src/qapair"

type Review struct {
	WhoWrited string
	Questions qapair.QAPairRepository
}

type PerfomanceReview struct {
	Review
	WrittenFor string
}

func NewPerfomanceReview(whoWrited, writtenFor string, questions qapair.QAPairRepository) PerfomanceReview {
	return PerfomanceReview{Review{WhoWrited: whoWrited, Questions: questions}, writtenFor}
}

type SelfReview struct {
	Review
}

func NewSelfReview(whoWrited string, questions qapair.QAPairRepository) SelfReview {
	return SelfReview{Review{WhoWrited: whoWrited, Questions: questions}}
}
