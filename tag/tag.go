package tag

import (
	"bufio"
	"github.com/PuerkitoBio/goquery"
)

type Tag struct {
	Root *goquery.Selection
	Parent *goquery.Selection
	Previous *goquery.Selection
	Current *goquery.Selection

	Ext *TagExt
}

type TagExt struct {
	HasTestData bool
}

type TagOp interface {
	Name() string
	Parse(writer *bufio.Writer)
	Next() TagOp
}

type TagContext struct {
	Prev *Tag
	Cur *goquery.Selection
}

func CreateTagContext(prev *Tag, cur *goquery.Selection) *TagContext {
	return &TagContext{
		Prev: prev,
		Cur: cur,
	}
}

func NewTagExt() *TagExt {
	return &TagExt{}
}