package search

import (
	"bytes"
	"github.com/yanyiwu/gojieba"
)

func cut(s string) []string {
	j := gojieba.NewJieba()
	words := j.CutForSearch(s, true)
	return words
}

func strConn(a ...string) string {
	var buffer bytes.Buffer
	for _, str := range a {
		buffer.WriteString(str)
	}
	return buffer.String()
}

func runeConn(a ...rune) string {
	var buffer bytes.Buffer
	for _, char := range a {
		buffer.WriteRune(char)
	}
	return buffer.String()
}

func allSubStrings(s string) []string {
	var (
		chars []rune
		strs  []string
	)
	
	for _, char := range s {
		chars = append(chars, char)
	}
	
	for i := 0; i < len(chars); i++ {
		for ii := 0; ii < len(chars)-i+1; ii++ {
			if i == 0 {
				continue
			}
			strs = append(strs, runeConn(chars[ii:ii+i]...))
		}
	}
	strs = append(strs, s)
	
	return strs
}

func exhaustTags(a ...string) []string {
	var bigTags, smallTags []string
	for _, v := range a {
		bigTags = append(bigTags, cut(v)...)
	}
	
	for _, v := range bigTags {
		smallTags = append(smallTags, allSubStrings(v)...)
	}
	smallTags = func() []string {
		var strs []string
		for _, v := range smallTags {
			if v != "" {
				strs = append(strs, v)
			}
		}
		return strs
	}()
	return smallTags
}

