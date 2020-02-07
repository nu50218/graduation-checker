package checkfuncs

import (
	checker "github.com/nu50218/graduation-checker"
	"github.com/nu50218/nucredit"
)

const eps = 1e-9

// CreditBorder 単位数がボーダー以上かを返すCheckFuncを返す
func CreditBorder(border float64) checker.CheckFunc {
	return func(ss nucredit.Subjects) (bool, error) {
		sum := 0.0
		for _, s := range ss {
			sum += s.Credit
		}
		return sum+eps > border, nil
	}
}

func equal(name1, name2 string) bool {
	// TODO: 表記ゆれ吸収
	return name1 == name2
}

// Acquired 指定した科目名の単位を取得したかを返すCheckFuncを返す
func Acquired(name string) checker.CheckFunc {
	return func(ss nucredit.Subjects) (bool, error) {
		for _, s := range ss {
			if equal(s.Name, name) {
				if s.Credit > eps {
					return true, nil
				}
			}
		}
		return false, nil
	}
}

// AcquiredAll 指定した科目名全ての単位を取得したかを返すCheckFuncを返す
func AcquiredAll(names ...string) checker.CheckFunc {
	checkFuncs := make([]checker.CheckFunc, len(names))
	for i := range checkFuncs {
		checkFuncs[i] = Acquired(names[i])
	}
	return AND(checkFuncs...)
}

// AND checkFuncsの全てがtrueならtrueを返すCheckFuncを返す
func AND(checkFuncs ...checker.CheckFunc) checker.CheckFunc {
	return func(ss nucredit.Subjects) (bool, error) {
		for _, checkFn := range checkFuncs {
			ok, err := checkFn(ss)
			if err != nil {
				return false, err
			}
			if !ok {
				return false, nil
			}
		}
		return true, nil
	}
}

// OR checkFuncsの一つでもtrueならtrueを返すCheckFuncを返す
func OR(checkFuncs ...checker.CheckFunc) checker.CheckFunc {
	return func(ss nucredit.Subjects) (bool, error) {
		for _, checkFn := range checkFuncs {
			ok, err := checkFn(ss)
			if err != nil {
				return false, err
			}
			if ok {
				return true, nil
			}
		}
		return false, nil
	}
}

// WithFilter filterをかけた上でチェックする
func WithFilter(checkFn checker.CheckFunc, filterFn nucredit.FilterFunc) checker.CheckFunc {
	return func(ss nucredit.Subjects) (bool, error) {
		return checkFn(ss.Filter(filterFn))
	}
}
