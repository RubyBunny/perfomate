package cmd

import (
	"fmt"
	"os"
	"perfomate/src/convertors"
	"perfomate/src/generators"
	"perfomate/src/reviews"
	"perfomate/src/searchers"
	"strings"

	"github.com/spf13/cobra"
)

var perfomanceCmd = &cobra.Command{
	Use:   "perfomance",
	Short: "Generate perfomance review",
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

		perfomanceReviewsSlice := convertor.Convert2PerfomanceReview()
		perfomanceReviewsMap := map[string][]*reviews.PerfomanceReview{}

		for _, review := range perfomanceReviewsSlice {
			fullname, err := searcher.Search(review.WrittenFor)
			if err != nil {
				fmt.Printf("Для \"%v\" %v. Review от %v пропущено!\n", fullname, err, review.WhoWrited)
				continue
			}

			perfomanceReviewsMap[fullname] = append(perfomanceReviewsMap[fullname], review)
		}

		for fullname, perfomanceReviews := range perfomanceReviewsMap {
			generator.GeneratePerfomanceReview(reviews.NewFinalPerfomanceReview(fullname, perfomanceReviews))
		}
	},
}
