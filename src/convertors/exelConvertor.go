package convertors

import (
	"perfomate/src/qapair"
	"perfomate/src/reviews"
	"strings"

	"github.com/xuri/excelize/v2"
)

type ExelConvertor struct {
	absolutePath string
}

func NewExelConvertor(absolutePath string) ExelConvertor {
	return ExelConvertor{absolutePath}
}

func (e ExelConvertor) Convert() []reviews.Review {
	f, _ := excelize.OpenFile(e.absolutePath)
	rows, _ := f.GetRows(f.GetSheetName(0))
	defer f.Close()

	var reviews []reviews.Review
	for _, answerRow := range rows[1:] {
		reviews = append(reviews, row2review(rows[0], answerRow))
	}

	return reviews
}

func row2review(questionRow, answerRow []string) reviews.Review {
	markedQuestions := row2markedQuestions(questionRow[2:20], answerRow[2:20])
	unmarkedQuestions := row2unmarkedQuestions(questionRow[20:], answerRow[20:])

	return reviews.Review{
		WhoWrited:  strings.TrimSpace(answerRow[0]),
		WrittenFor: strings.TrimSpace(answerRow[1]),
		Questions: qapair.QAPairRepository{
			MarkedQuestions:   markedQuestions,
			UnmarkedQuestions: unmarkedQuestions,
		},
	}
}

func row2markedQuestions(questionRow, answerRow []string) []qapair.MarkedQAPair {
	var markedQuestions []qapair.MarkedQAPair

	for i := 0; i < len(questionRow); i += 2 {
		markedQuestions = append(markedQuestions, qapair.NewMarkedQAPair(
			questionRow[i], strings.TrimSpace(answerRow[i+1]), answerRow[i],
		))
	}

	return markedQuestions
}

func row2unmarkedQuestions(questionRow, answerRow []string) []qapair.QAPair {
	unmarkedQuestions := make([]qapair.QAPair, len(questionRow))

	for i, question := range questionRow {
		unmarkedQuestions[i].Question = question
	}

	for i, answer := range answerRow {
		unmarkedQuestions[i].Answer = strings.TrimSpace(answer)
	}

	return unmarkedQuestions
}
