package convertors

import (
	"perfomate/src/reviews"
)

type ConvertorCore interface {
	Convert2PerfomanceReview() []*reviews.PerfomanceReview
	Convert2SelfReview() []*reviews.SelfReview
}

type Convertor struct {
	core ConvertorCore
}

func NewConvertor(absolutePath string) Convertor {
	var convertorEngine ConvertorCore = NewExelConvertor(absolutePath)

	return Convertor{convertorEngine}
}

func (c Convertor) Convert2PerfomanceReview() []*reviews.PerfomanceReview {
	return c.core.Convert2PerfomanceReview()
}

func (c Convertor) Convert2SelfReview() []*reviews.SelfReview {
	return c.core.Convert2SelfReview()
}
