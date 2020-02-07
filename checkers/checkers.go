package checkers

import (
	checker "github.com/nu50218/graduation-checker"
	"github.com/nu50218/graduation-checker/checkers/fukuzatsu"
	"github.com/nu50218/graduation-checker/checkers/sizen"
	"github.com/nu50218/graduation-checker/checkers/suuri"
)

var checkersList = map[string]checker.Checkers{
	"sizen":     sizen.Checkers,
	"suuri":     suuri.Checkers,
	"fukuzatsu": fukuzatsu.Checkers,
}

func GetCheckers(target string) (checker.Checkers, bool) {
	cs, exists := checkersList[target]
	return cs, exists
}
