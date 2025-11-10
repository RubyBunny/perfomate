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

	reviewsSlice := convertor.Convert()
	reviewsMap := map[string][]reviews.Review{}

	for _, review := range reviewsSlice {
		fullname, err := searcher.Search(review.WrittenFor)
		if err != nil {
			fmt.Printf("Для \"%v\" %v. Review от %v пропущено!\n", fullname, err, review.WhoWrited)
			continue
		}

		reviewsMap[fullname] = append(reviewsMap[fullname], review)
	}

	for fullname, reviewArr := range reviewsMap {
		generator.Generate(reviews.NewFinalReview(fullname, reviewArr))
	}
}
