package searchers

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

type FullnameSearcher struct {
	fullnamesSlice []string
}

func NewFullnameSearcher(fullnamesSlice []string) FullnameSearcher {
	return FullnameSearcher{fullnamesSlice}
}

func (s FullnameSearcher) Search(fullname string) (string, error) {
	suitableFullnames := recursiveSearch(
		s.fullnamesSlice,
		strings.Split(fullname, " "),
		0,
	)

	if len(suitableFullnames) > 1 {
		fmt.Printf("Обнаружена коллизия! Искомая строка \"%v\" соответствует нескольким значениям!\n", fullname)
		fmt.Println("Выберете нужный вариант: ")
		for i, fullname := range suitableFullnames {
			fmt.Printf("[%v] %v\n", i, fullname)
		}

		var i int
		fmt.Print("Введите число: ")
		fmt.Scanln(&i)

		suitableFullnames = []string{suitableFullnames[i]}
	}

	if len(suitableFullnames) == 0 {
		return fullname, errors.New("не найдено ни одного подходящего элемента")
	}

	return suitableFullnames[0], nil
}

func recursiveSearch(stringsArr, substrings []string, depthLevel uint) []string {
	var filtredArr []string

	for _, _string := range stringsArr {
		if slices.Contains(strings.Split(_string, " "), substrings[depthLevel]) {
			filtredArr = append(filtredArr, _string)
		}
	}

	if len(stringsArr) == 1 || int(depthLevel) == len(substrings)-1 {
		return filtredArr
	}

	return recursiveSearch(filtredArr, substrings, depthLevel+1)
}
