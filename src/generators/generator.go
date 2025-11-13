package generators

import "perfomate/src/reviews"

type GeneratorCore interface {
	GeneratePerfomanceReview(finalReview reviews.FinalPerfomanceReview)
}

type Generator struct {
	core GeneratorCore
}

func NewGenerator(filePath, fileType string) Generator {
	core := NewExelGenerator(filePath)
	return Generator{core}
}

func (g Generator) Generate(finalReview reviews.FinalPerfomanceReview) {
	g.core.GeneratePerfomanceReview(finalReview)
}
