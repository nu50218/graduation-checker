package sizen

import (
	"strings"

	checker "github.com/nu50218/graduation-checker"
	"github.com/nu50218/graduation-checker/checkfuncs"
	"github.com/nu50218/nucredit"
)

var Checkers = checker.Checkers{
	checker.NewSimpleChecker(
		"合計128単位",
		"合計で128単位取得すること。",
		checkfuncs.CreditBorder(128),
	),

	checker.NewSimpleChecker(
		"基礎セミナー",
		"基礎セミナーA又は基礎セミナーBから2単位を含むこと。",
		checkfuncs.WithFilter(
			checkfuncs.CreditBorder(2),
			func(s *nucredit.Subject) bool {
				return strings.HasPrefix(s.Name, "基礎セミナー")
			},
		),
	),

	checker.NewSimpleChecker(
		"英語",
		"「言語文化」として英語6単位を含むこと。",
		checkfuncs.WithFilter(
			checkfuncs.CreditBorder(6),
			func(s *nucredit.Subject) bool {
				return strings.HasPrefix(s.Name, "英語")
			},
		),
	),

	checker.NewSimpleChecker(
		"第二外国語",
		"英語以外の外国語(ドイツ語，フランス語，ロシア語，中国語，スペイン語，朝鮮・韓国語及び日本語(外国人留学生対象))のうちから1外国語6単位を含むこと。",
		checkfuncs.WithFilter(
			checkfuncs.CreditBorder(6),
			func(s *nucredit.Subject) bool {
				return s.Category == "言語文化" && !strings.HasPrefix(s.Name, "英語")
			},
		),
	),

	checker.NewSimpleChecker(
		"健康・スポーツ科学",
		"「健康・スポーツ科学」として講義又は実習から2単位を含むこと。",
		checkfuncs.WithFilter(
			checkfuncs.CreditBorder(2),
			func(s *nucredit.Subject) bool {
				return strings.HasPrefix(s.Category, "健康・スポーツ科学")
			},
		),
	),

	checker.NewSimpleChecker(
		"文系基礎・教養科目",
		"「文系基礎科目」及び「文系教養科目」6単位を含むこと。",
		checkfuncs.WithFilter(
			checkfuncs.CreditBorder(6),
			func(s *nucredit.Subject) bool {
				return strings.HasPrefix(s.Category, "文系基礎科目") || strings.HasPrefix(s.Category, "文系教養科目")
			},
		),
	),

	checker.NewSimpleChecker(
		"理系教養科目",
		"「理系教養科目」2単位を含むこと。",
		checkfuncs.WithFilter(
			checkfuncs.CreditBorder(2),
			func(s *nucredit.Subject) bool {
				return strings.HasPrefix(s.Category, "理系教養科目")
			},
		),
	),

	checker.NewSimpleChecker(
		"全学教養科目",
		"「全学教養科目」2単位を含むこと。",
		checkfuncs.WithFilter(
			checkfuncs.CreditBorder(2),
			func(s *nucredit.Subject) bool {
				return strings.HasPrefix(s.Category, "全学教養科目")
			},
		),
	),

	checker.NewSimpleChecker(
		"専門基礎科目",
		"30単位以上取得すること。",
		checkfuncs.WithFilter(
			checkfuncs.CreditBorder(30),
			func(s *nucredit.Subject) bool {
				return strings.Contains(s.Category, "専門基礎科目")
			},
		),
	),

	checker.NewSimpleChecker(
		"インフォマティックス1~4",
		"「インフォマティックス1〜4」を含むこと",
		checkfuncs.WithFilter(
			checkfuncs.CreditBorder(4),
			func(s *nucredit.Subject) bool {
				return strings.Contains(s.Name, "インフォマティックス")
			},
		),
	),

	checker.NewSimpleChecker(
		"情報の挑戦者・開拓者たち",
		"「情報の挑戦者・開拓者たち」を含むこと",
		checkfuncs.WithFilter(
			checkfuncs.CreditBorder(2),
			func(s *nucredit.Subject) bool {
				return strings.Contains(s.Name, "情報の挑戦者・開拓者たち")
			},
		),
	),

	checker.NewSimpleChecker(
		"情報セキュリティとリテラシー1~2",
		"「情報セキュリティとリテラシー1~2」を含むこと",
		checkfuncs.WithFilter(
			checkfuncs.CreditBorder(2),
			func(s *nucredit.Subject) bool {
				return strings.Contains(s.Name, "情報セキュリティとリテラシー")
			},
		),
	),

	checker.NewSimpleChecker(
		"プログラミング1~2",
		"「プログラミング1~2」を含むこと",
		checkfuncs.WithFilter(
			checkfuncs.CreditBorder(2),
			func(s *nucredit.Subject) bool {
				return strings.Contains(s.Name, "プログラミング")
			},
		),
	),

	checker.NewSimpleChecker(
		"論理学1",
		"「論理学1」を含むこと",
		checkfuncs.WithFilter(
			checkfuncs.CreditBorder(1),
			func(s *nucredit.Subject) bool {
				return strings.Contains(s.Name, "論理学１")
			},
		),
	),

	checker.NewSimpleChecker(
		"データマイニング入門",
		"データマイニング入門",
		checkfuncs.WithFilter(
			checkfuncs.CreditBorder(1),
			func(s *nucredit.Subject) bool {
				return strings.Contains(s.Name, "データマイニング入門")
			},
		),
	),

	checker.NewSimpleChecker(
		"専門科目",
		"40単位以上取得すること。",
		checkfuncs.WithFilter(
			checkfuncs.CreditBorder(40),
			func(s *nucredit.Subject) bool {
				return strings.Contains(s.Category, "専門科目")
			},
		),
	),

	checker.NewSimpleChecker(
		"関連専門科目",
		"2単位以上取得すること。",
		checkfuncs.WithFilter(
			checkfuncs.CreditBorder(2),
			func(s *nucredit.Subject) bool {
				return strings.Contains(s.Category, "関連専門科目")
			},
		),
	),

	checker.NewSimpleChecker(
		"情報倫理と法",
		"「情報倫理と法」を含むこと。",
		checkfuncs.WithFilter(
			checkfuncs.CreditBorder(1),
			func(s *nucredit.Subject) bool {
				return strings.Contains(s.Name, "情報倫理と法")
			},
		),
	),
}
