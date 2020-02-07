package suuri

import (
	"strings"

	checker "github.com/nu50218/graduation-checker"
	"github.com/nu50218/graduation-checker/checkers/sizen"
	"github.com/nu50218/graduation-checker/checkfuncs"
	"github.com/nu50218/nucredit"
)

var Checkers = append(sizen.Checkers,
	checker.NewSimpleChecker(
		"数理情報学序論1~2",
		"「数理情報学序論1~2」を含むこと。",
		checkfuncs.WithFilter(
			checkfuncs.CreditBorder(1),
			func(s *nucredit.Subject) bool {
				return strings.Contains(s.Name, "数理情報学序論")
			},
		),
	),

	checker.NewSimpleChecker(
		"数理情報系開講の専門科目16単位",
		"数理情報系が開講する専門科目16単位以上及び卒業研究を含むこと。（判定できないので自分で判定してください）",
		func(nucredit.Subjects) (bool, error) {
			return false, nil
		},
	),
)
