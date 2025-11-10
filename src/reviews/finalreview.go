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
	avgMark := calcAverageMark(questions)
	status := avgMark2Result(avgMark)

	return FinalReview{
		Fullname:    fullname,
		Respondents: respondents,
		Questions:   questions,
		AvgMark:     avgMark,
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
		var avgMark float64

		for ri := range reviews {
			if reviews[ri].Questions.MarkedQuestions[qi].Answer != "" {
				answers = append(
					answers,
					fmt.Sprintf(AnswerFormat, reviews[ri].Questions.MarkedQuestions[qi].Answer),
				)
			}
			avgMark += reviews[ri].Questions.MarkedQuestions[qi].Mark
		}

		markedQuestions = append(markedQuestions, qapair.MarkedQAPair{
			QAPair: qapair.QAPair{
				Question: reviews[0].Questions.MarkedQuestions[qi].Question,
				Answer:   strings.Join(answers, "\n"),
			},
			Mark: avgMark / float64(len(reviews)),
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

func calcAverageMark(questions qapair.QAPairRepository) float64 {
	var avgMark float64

	for _, q := range questions.MarkedQuestions {
		avgMark += q.Mark
	}

	return avgMark / float64(len(questions.MarkedQuestions))
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
		fallthrough
	default:
		return "Не пройден"
	}
}
