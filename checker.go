package checker

import (
	"github.com/nu50218/nucredit"
)

// Checker 卒業要件それぞれについてCheckerを用意する
type Checker interface {
	// Title タイトル
	Title() string
	// Description 詳細
	Description() string
	// Checkする
	Check(nucredit.Subjects) (bool, error)
}

// CheckFunc チェックする関数
type CheckFunc func(nucredit.Subjects) (bool, error)

// SimpleChecker 最低限の機能だけ備えたChecker
type SimpleChecker struct {
	title       string
	description string
	checkFn     CheckFunc
}

// Title タイトル
func (c *SimpleChecker) Title() string {
	return c.title
}

func (c *SimpleChecker) Description() string {
	return c.description
}

// Check チェックする
func (c *SimpleChecker) Check(s nucredit.Subjects) (bool, error) {
	return c.checkFn(s)
}

// NewSimpleChecker SimpleなCheckerをつくる
func NewSimpleChecker(title, description string, checkFn CheckFunc) *SimpleChecker {
	return &SimpleChecker{
		title:       title,
		description: description,
		checkFn:     checkFn,
	}
}
