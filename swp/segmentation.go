package swp

import "github.com/yanyiwu/gojieba"

func cut(s string) []string {
	j := gojieba.NewJieba()
	words := j.CutForSearch(s, true)
	return words
}
