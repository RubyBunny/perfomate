package convertors

import (
	"perfomate/src/reviews"
)

type ConvertorCore interface {
	Convert() []reviews.Review
}

type Convertor struct {
	core ConvertorCore
}

func NewConvertor(absolutePath string) Convertor {
	var convertorEngine ConvertorCore = NewExelConvertor(absolutePath)

	return Convertor{convertorEngine}
}

func (c Convertor) Convert() []reviews.Review {
	return c.core.Convert()
}
