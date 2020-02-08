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
		"理系基礎科目",
		"18単位以上取得すること。",
		checkfuncs.WithFilter(
			checkfuncs.CreditBorder(18),
			func(s *nucredit.Subject) bool {
				return strings.Contains(s.Category, "理系基礎科目")
			},
		),
	),

	checker.NewSimpleChecker(
		"微積・線形ⅠⅡ",
		"微分積分学I，微分積分学II，線形代数学I及び線形代数学IIの4科目8単位を含むこと。",
		checkfuncs.AcquiredAll(
			"微分積分学1",
			"微分積分学2",
			"線形代数学1",
			"線形代数学2",
		),
	),

	checker.NewSimpleChecker(
		"理科基礎・実験",
		"「物理学基礎I，物理学基礎II及び物理学実験」，「化学基礎I，化学基礎II及び化学実験」，「生物学基礎I，生物学基礎II及び生物学実験」又は「地球科学基礎I，地球科学基礎II及び地球科学実験」のうちから1組，3科目，計5.5単位を含むこと。前号以外の「物理学基礎I及び物理学基礎II」，「化学基礎Ⅰ及び化学基礎Ⅱ」，「生物学基礎Ⅰ及び生物学基礎Ⅱ」又は「地球科学基礎Ⅰ及び地球科学基礎Ⅱ」から１組, 2科目，計4単位以上を含むこと。",
		checkfuncs.OR(
			checkfuncs.AND(
				checkfuncs.AcquiredAll("物理学基礎1", "物理学基礎2"),
				checkfuncs.AcquiredAll("化学基礎1", "化学基礎2"),
				checkfuncs.OR(
					checkfuncs.Acquired("物理学実験"),
					checkfuncs.Acquired("化学実験"),
				),
			),
			checkfuncs.AND(
				checkfuncs.AcquiredAll("物理学基礎1", "物理学基礎2"),
				checkfuncs.AcquiredAll("生物学基礎1", "生物学基礎2"),
				checkfuncs.OR(
					checkfuncs.Acquired("物理学実験"),
					checkfuncs.Acquired("生物学実験"),
				),
			),
			checkfuncs.AND(
				checkfuncs.AcquiredAll("物理学基礎1", "物理学基礎2"),
				checkfuncs.AcquiredAll("地球科学基礎1", "地球科学基礎2"),
				checkfuncs.OR(
					checkfuncs.Acquired("物理学実験"),
					checkfuncs.Acquired("地球科学実験"),
				),
			),
			checkfuncs.AND(
				checkfuncs.AcquiredAll("化学基礎1", "化学基礎2"),
				checkfuncs.AcquiredAll("生物学基礎1", "生物学基礎2"),
				checkfuncs.OR(
					checkfuncs.Acquired("化学実験"),
					checkfuncs.Acquired("生物学実験"),
				),
			),
			checkfuncs.AND(
				checkfuncs.AcquiredAll("化学基礎1", "化学基礎2"),
				checkfuncs.AcquiredAll("地球科学基礎1", "地球科学基礎2"),
				checkfuncs.OR(
					checkfuncs.Acquired("化学実験"),
					checkfuncs.Acquired("地球科学実験"),
				),
			),
			checkfuncs.AND(
				checkfuncs.AcquiredAll("生物学基礎1", "生物学基礎2"),
				checkfuncs.AcquiredAll("地球科学基礎1", "地球科学基礎2"),
				checkfuncs.OR(
					checkfuncs.Acquired("生物学実験"),
					checkfuncs.Acquired("地球科学実験"),
				),
			),
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
