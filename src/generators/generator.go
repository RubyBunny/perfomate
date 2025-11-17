package generators

import "perfomate/src/reviews"

type GeneratorCore interface {
	GeneratePerfomanceReview(finalReview reviews.FinalPerfomanceReview)
	GenerateSelfReview(selfReview *reviews.SelfReview)
}

type Generator struct {
	core GeneratorCore
}

func NewGenerator(filePath, fileType string) Generator {
	core := NewExelGenerator(filePath)
	return Generator{core}
}

func (g Generator) GeneratePerfomanceReview(finalReview reviews.FinalPerfomanceReview) {
	g.core.GeneratePerfomanceReview(finalReview)
}

func (g Generator) GenerateSelfReview(selfReview *reviews.SelfReview) {
	g.core.GenerateSelfReview(selfReview)
}
