package generators

import "perfomate/src/reviews"

type GeneratorCore interface {
	Generate(finalReview reviews.FinalReview)
}

type Generator struct {
	core GeneratorCore
}

func NewGenerator(filePath, fileType string) Generator {
	core := NewExelGenerator(filePath)
	return Generator{core}
}

func (g Generator) Generate(finalReview reviews.FinalReview) {
	g.core.Generate(finalReview)
}
