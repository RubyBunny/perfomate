package reviews

import (
	"fmt"
	"math"
	"perfomate/src/qapair"
	"strings"
)

const (
	AnswerFormat = "— %v"
)

type FinalReview struct {
	Fullname    string
	Respondents string
	Questions   qapair.QAPairRepository
	AvgMark     float64
	Status      string
}

func NewFinalReview(fullname string, reviews []Review) FinalReview {
	respondents := concatRespondents(reviews)
	questions := joinQuestions(reviews)

	var marks []float64
	for _, mq := range questions.MarkedQuestions {
		marks = append(marks, mq.Mark)
	}

	avg := calcAverageMark(marks)
	status := avgMark2Result(avg)

	return FinalReview{
		Fullname:    fullname,
		Respondents: respondents,
		Questions:   questions,
		AvgMark:     avg,
		Status:      status,
	}
}

func concatRespondents(reviews []Review) string {
	var respondents []string
	for _, review := range reviews {
		respondents = append(respondents, review.WhoWrited)
	}

	return strings.Join(respondents, ", ")
}

func joinQuestions(reviews []Review) qapair.QAPairRepository {
	var markedQuestions []qapair.MarkedQAPair
	var unmarkedQuestions []qapair.QAPair

	for qi := range reviews[0].Questions.MarkedQuestions {
		var answers []string
		marks := make([]float64, len(reviews))

		for ri := range reviews {
			if reviews[ri].Questions.MarkedQuestions[qi].Answer != "" {
				answers = append(
					answers,
					fmt.Sprintf(AnswerFormat, reviews[ri].Questions.MarkedQuestions[qi].Answer),
				)
			}

			marks[ri] = reviews[ri].Questions.MarkedQuestions[qi].Mark
		}

		markedQuestions = append(markedQuestions, qapair.MarkedQAPair{
			QAPair: qapair.QAPair{
				Question: reviews[0].Questions.MarkedQuestions[qi].Question,
				Answer:   strings.Join(answers, "\n"),
			},
			Mark: calcAverageMark(marks),
		})
	}

	for qi := range reviews[0].Questions.UnmarkedQuestions {
		var answers []string

		for ri := range reviews {
			if reviews[ri].Questions.UnmarkedQuestions[qi].Answer != "" {
				answers = append(
					answers,
					fmt.Sprintf(AnswerFormat, reviews[ri].Questions.UnmarkedQuestions[qi].Answer),
				)
			}
		}

		unmarkedQuestions = append(unmarkedQuestions, qapair.QAPair{
			Question: reviews[0].Questions.UnmarkedQuestions[qi].Question,
			Answer:   strings.Join(answers, "\n"),
		})
	}

	return qapair.QAPairRepository{
		MarkedQuestions:   markedQuestions,
		UnmarkedQuestions: unmarkedQuestions,
	}
}

func calcAverageMark(marks []float64) float64 {
	var sum float64
	var zeroes int

	for _, mark := range marks {
		if mark == 0 {
			zeroes++
		}

		sum += mark
	}

	if len(marks) == zeroes {
		return 0
	}

	return sum / float64(len(marks)-zeroes)
}

func avgMark2Result(avg float64) string {
	switch math.Round(avg) {
	case 5:
		return "Успешно пройден"
	case 4:
		return "Пройден"
	case 3:
		return "Пройден с замечаниями"
	case 2:
		return "Не пройден"
	default:
		return "N/A"
	}
}
