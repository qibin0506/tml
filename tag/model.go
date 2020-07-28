package tag

import (
	"log"
	"bufio"
	"github.com/qibin0506/tml/utils"
)

func NewModel(ctx *TagContext) TagOp {
	return &Model{
		Tag: &Tag{
			Root: ctx.Prev.Root,
			Parent: ctx.Cur,
			Previous: ctx.Prev.Current,
			Current: ctx.Cur,
			Ext: ctx.Prev.Ext,
		},
	}
}

type Model struct {
	*Tag
}

func (m *Model) Name() string {
	return "model"
}

func (m *Model) Parse(writer *bufio.Writer) {
	writer.WriteString("# start build model.\n\n")
}

func (m *Model) Next() TagOp {
	firstChildNode := utils.FirstChildNode(m.Current)
	firstChildNodeName := utils.NodeName(firstChildNode)
	if firstChildNodeName == "" {
		log.Fatal("tag model must have children.")
	}

	nextTagFunc := tagMap[firstChildNodeName]
	if nextTagFunc == nil {
		log.Fatal("tag %s was not supported.", firstChildNodeName)
	}

	return nextTagFunc(CreateTagContext(m.Tag, firstChildNode))
}