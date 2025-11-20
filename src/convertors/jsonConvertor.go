package convertors

import (
	"encoding/json"
	"fmt"
	"os"
	"perfomate/src/qapair"
	"perfomate/src/reviews"
	"strings"
)

type JSONConvertor struct {
	absolutePath string
}

func NewJSONConvertor(absolutePath string) JSONConvertor {
	return JSONConvertor{absolutePath}
}

func (c JSONConvertor) Convert2PerfomanceReview() []*reviews.PerfomanceReview {
	var perfomanceReviews []*reviews.PerfomanceReview

	file, _ := os.Open(c.absolutePath)
	defer file.Close()

	var rawData [][][]string
	err := json.NewDecoder(file).Decode(&rawData)
	if err != nil {
		fmt.Println(err)
	}

	for _, row := range rawData {
		perfomanceReviews = append(perfomanceReviews, c.row2perfomanceReview(row))
	}

	return perfomanceReviews
}

func (c JSONConvertor) Convert2SelfReview() []*reviews.SelfReview {
	var selfReviews []*reviews.SelfReview

	file, _ := os.Open(c.absolutePath)
	defer file.Close()

	var rawData [][][]string
	err := json.NewDecoder(file).Decode(&rawData)
	if err != nil {
		fmt.Println(err)
	}

	for _, row := range rawData {
		selfReviews = append(selfReviews, c.row2selfReview(row))
	}

	return selfReviews
}

func (c JSONConvertor) row2perfomanceReview(row [][]string) *reviews.PerfomanceReview {
	markedQuestions := c.row2markedQuestions(row[2:20])
	unmarkedQuestions := c.row2unmarkedQuestions(row[20:])

	return reviews.NewPerfomanceReview(
		strings.TrimSpace(row[0][1]),
		strings.TrimSpace(row[1][1]),
		qapair.QAPairRepository{
			MarkedQuestions:   markedQuestions,
			UnmarkedQuestions: unmarkedQuestions,
		},
	)
}

func (c JSONConvertor) row2selfReview(row [][]string) *reviews.SelfReview {
	unmarkedQuestions := c.row2unmarkedQuestions(row[1:])

	return reviews.NewSelfReview(
		strings.TrimSpace(row[0][1]),
		qapair.QAPairRepository{
			UnmarkedQuestions: unmarkedQuestions,
		},
	)
}

func (c JSONConvertor) row2markedQuestions(row [][]string) []qapair.MarkedQAPair {
	var markedQuestions []qapair.MarkedQAPair

	for i := 0; i < len(row); i += 2 {
		markedQuestions = append(markedQuestions, qapair.NewMarkedQAPair(
			row[i][0], strings.TrimSpace(row[i+1][1]), row[i][1],
		))
	}

	return markedQuestions
}

func (c JSONConvertor) row2unmarkedQuestions(row [][]string) []qapair.QAPair {
	unmarkedQuestions := make([]qapair.QAPair, len(row))

	for i, qa := range row {
		unmarkedQuestions[i] = qapair.QAPair{Question: qa[0], Answer: strings.TrimSpace(qa[1])}
	}

	return unmarkedQuestions
}
