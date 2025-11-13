package main

import (
	"flag"
	"fmt"
	"os"
	"perfomate/src/convertors"
	"perfomate/src/generators"
	"perfomate/src/reviews"
	"perfomate/src/searchers"
	"strings"
)

func main() {
	inputFilePath := flag.String("i", "", "")
	ouputPath := flag.String("o", "./", "")
	flag.Parse()

	fullnames, _ := os.ReadFile("./users.txt")

	convertor := convertors.NewConvertor(*inputFilePath)
	generator := generators.NewGenerator(*ouputPath, "xlsx")
	searcher := searchers.NewFullnameSearcher(strings.Split(string(fullnames), "\r\n"))

	perfomanceReviewsSlice := convertor.Convert2PerfomanceReview()
	perfomanceReviewsMap := map[string][]reviews.PerfomanceReview{}

	for _, review := range perfomanceReviewsSlice {
		fullname, err := searcher.Search(review.WrittenFor)
		if err != nil {
			fmt.Printf("Для \"%v\" %v. Review от %v пропущено!\n", fullname, err, review.WhoWrited)
			continue
		}

		perfomanceReviewsMap[fullname] = append(perfomanceReviewsMap[fullname], review)
	}

	for fullname, perfomanceReviews := range perfomanceReviewsMap {
		generator.Generate(reviews.NewFinalPerfomanceReview(fullname, perfomanceReviews))
	}
}
