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

func (e ExelConvertor) Convert2PerfomanceReview() []*reviews.PerfomanceReview {
	f, _ := excelize.OpenFile(e.absolutePath)
	rows, _ := f.GetRows(f.GetSheetName(0))
	defer f.Close()

	var reviews []*reviews.PerfomanceReview
	for _, answerRow := range rows[1:] {
		reviews = append(reviews, row2perfomanceReview(rows[0], answerRow))
	}

	return reviews
}

func row2perfomanceReview(questionRow, answerRow []string) *reviews.PerfomanceReview {
	markedQuestions := row2markedQuestions(questionRow[2:20], answerRow[2:20])
	unmarkedQuestions := row2unmarkedQuestions(questionRow[20:], answerRow[20:])

	return reviews.NewPerfomanceReview(
		strings.TrimSpace(answerRow[0]),
		strings.TrimSpace(answerRow[1]),
		qapair.QAPairRepository{
			UnmarkedQuestions: unmarkedQuestions,
			MarkedQuestions:   markedQuestions,
		},
	)
}

func (e ExelConvertor) Convert2SelfReview() []*reviews.SelfReview {
	f, _ := excelize.OpenFile(e.absolutePath)
	rows, _ := f.GetRows(f.GetSheetName(0))
	defer f.Close()

	var reviews []*reviews.SelfReview
	for _, answerRow := range rows[1:] {
		reviews = append(reviews, row2selfReview(rows[0], answerRow))
	}

	return reviews
}

func row2selfReview(questionRow, answerRow []string) *reviews.SelfReview {
	unmarkedQuestions := row2unmarkedQuestions(questionRow[1:], answerRow[1:])

	return reviews.NewSelfReview(
		strings.TrimSpace(answerRow[0]),
		qapair.QAPairRepository{
			UnmarkedQuestions: unmarkedQuestions,
		},
	)
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
