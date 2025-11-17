package cmd

import (
	"perfomate/src/convertors"
	"perfomate/src/generators"

	"github.com/spf13/cobra"
)

var selfCmd = &cobra.Command{
	Use:   "self",
	Short: "Generate self review",
	Run: func(cmd *cobra.Command, args []string) {
		convertor := convertors.NewConvertor(InputPath)
		generator := generators.NewGenerator(OutputPath, "xlsx")

		selfReviewsSlice := convertor.Convert2SelfReview()

		for _, review := range selfReviewsSlice {
			generator.GenerateSelfReview(review)
		}
	},
}
