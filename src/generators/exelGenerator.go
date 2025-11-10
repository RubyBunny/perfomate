package generators

import (
	"fmt"
	"path/filepath"
	"perfomate/src/reviews"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

type ExelGenerator struct {
	filePath string
}

func NewExelGenerator(filePath string) ExelGenerator {
	return ExelGenerator{filePath}
}

func (e ExelGenerator) Generate(finalReview reviews.FinalReview) {
	f := excelize.NewFile()
	defer f.Close()

	allCenterStyle, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Vertical:   "center",
			Horizontal: "center",
			WrapText:   true,
		},
	})

	textWrapStyle, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Vertical: "center",
			WrapText: true,
		},
	})

	f.SetCellStyle("Sheet1", "A2", "A2", allCenterStyle)
	f.SetCellStyle("Sheet1", "B14", "B14", textWrapStyle)
	f.SetCellStyle("Sheet1", "A18", "C29", textWrapStyle)
	f.SetCellStyle("Sheet1", "B13", "B29", allCenterStyle)
	f.SetCellStyle("Sheet1", "A17", "C17", allCenterStyle)

	f.SetColWidth("Sheet1", "A", "A", 54)
	f.SetColWidth("Sheet1", "B", "B", 20)
	f.SetColWidth("Sheet1", "C", "C", 85)

	f.MergeCell("Sheet1", "A2", "C2")
	f.MergeCell("Sheet1", "B5", "C5")
	f.MergeCell("Sheet1", "B6", "C6")
	f.MergeCell("Sheet1", "B7", "C7")
	f.MergeCell("Sheet1", "B8", "C8")
	f.MergeCell("Sheet1", "B9", "C9")
	f.MergeCell("Sheet1", "B10", "C10")

	writeRow(f, "A2", []any{"Результат Perfomance Review"})

	writeRow(f, "A5", []any{"ФИО", finalReview.Fullname})
	writeRow(f, "A6", []any{"Должность"})
	writeRow(f, "A7", []any{"Грейд"})
	writeRow(f, "A8", []any{"Проекты"})
	writeRow(f, "A9", []any{"Дата Review", time.Now().Format("02 January 2006")})
	writeRow(f, "A10", []any{"Респонденты", finalReview.Respondents})

	writeRow(f, "A13", []any{"Результат Perfomance Review", finalReview.Status})
	writeRow(f, "A14", []any{"Средняя оценка (по 5-бальной шкале)", fmt.Sprintf("%.1f", finalReview.AvgMark)})

	writeRow(f, "A17", []any{"Вопрос", "Оценка (по 5-бальной шкале)", "Комментарий"})

	for i, question := range finalReview.Questions.MarkedQuestions {
		cell := fmt.Sprintf("A%d", i+18)
		writeRow(
			f,
			cell,
			[]any{question.Question, question.Mark, question.Answer},
		)
	}

	for i, question := range finalReview.Questions.UnmarkedQuestions {
		cell := fmt.Sprintf("A%d", i+27)
		writeRow(
			f,
			cell,
			[]any{question.Question, "", question.Answer},
		)
	}

	err := f.SaveAs(
		filepath.Join(
			e.filePath,
			fmt.Sprintf("PR %v.xlsx", getFileName(finalReview.Fullname)),
		))
	if err != nil {
		fmt.Println(err)
	}
}

func writeRow(file *excelize.File, cell string, slice []any) {
	file.SetSheetRow("Sheet1", cell, &slice)
}

func getFileName(fullname string) string {
	now := time.Now()
	return fmt.Sprintf(
		"%v %v%v%v",
		strings.Split(fullname, " ")[0],
		now.Day(),
		int(now.Month()),
		now.Year(),
	)
}
