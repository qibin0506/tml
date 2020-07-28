package tag

import (
	"bufio"
	"github.com/qibin0506/tml/utils"
)

func NewTrain(ctx *TagContext) TagOp {
	return &Train{
		Tag: &Tag{
			Root: ctx.Prev.Root,
			Parent: ctx.Cur,
			Previous: ctx.Prev.Current,
			Current: ctx.Cur,
			Ext: ctx.Prev.Ext,
		},
	}
}

type Train struct {
	*Tag
}

func (t *Train) Name() string {
	return "train"
}

func (t *Train) Parse(writer *bufio.Writer) {
	writer.WriteString("# start train data.\n\n")
}

func (t *Train) Next() TagOp {
	loss := utils.GetTagOrFatal(t.Parent, "loss")
	return tagMap["loss"](CreateTagContext(t.Tag, loss))
}