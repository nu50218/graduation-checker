package checker

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/nu50218/nucredit"
)

// Checkers Checkerの集合
type Checkers []Checker

// Run 指定したSubjectsに対してCheckerを実行する
func (cs Checkers) Run(s nucredit.Subjects) error {
	for _, c := range cs {
		ok, err := c.Check(s)
		if err != nil {
			return err
		}

		if ok {
			color.New(color.FgHiGreen).Print("[✔] ")
			fmt.Println(c.Title())
		} else {
			color.New(color.FgRed).Print("[✗] ")
			fmt.Print(c.Title())
			color.New(color.FgYellow).Print(" [条件] ")
			fmt.Println(c.Description())
		}
	}

	return nil
}

// RunFile 指定したファイルに対してチェックする
func (cs Checkers) RunFile(name string) error {
	s, err := nucredit.FromFile(name)
	if err != nil {
		return err
	}
	return cs.Run(s)
}
