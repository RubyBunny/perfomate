package cmd

import (
	"fmt"
	"os"
	"perfomate/src/convertors"
	"perfomate/src/generators"
	"perfomate/src/searchers"
	"strings"

	"github.com/spf13/cobra"
)

var selfCmd = &cobra.Command{
	Use:   "self",
	Short: "Generate self review",
	Run: func(cmd *cobra.Command, args []string) {

		var convertorCore convertors.ConvertorCore

		if IsJSON {
			convertorCore = convertors.NewJSONConvertor(InputPath)
		} else {
			convertorCore = convertors.NewExelConvertor(InputPath)
		}

		convertor := convertors.NewConvertor(convertorCore)
		generator := generators.NewGenerator(OutputPath, "xlsx")

		fullnames, _ := os.ReadFile("./users.txt")
		searcher := searchers.NewFullnameSearcher(strings.Split(string(fullnames), "\r\n"))

		selfReviewsSlice := convertor.Convert2SelfReview()

		for _, review := range selfReviewsSlice {
			if fullname, err := searcher.Search(review.WhoWrited); err == nil {
				review.WhoWrited = fullname
			} else {
				fmt.Printf(
					"ФИО \"%v\" не найдено! Использовано ФИО из входного файла\n",
					fullname,
				)
			}

			generator.GenerateSelfReview(review)
		}
	},
}
