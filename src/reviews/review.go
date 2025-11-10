package reviews

import "perfomate/src/qapair"

type Review struct {
	WrittenFor string
	WhoWrited  string
	Questions  qapair.QAPairRepository
}
