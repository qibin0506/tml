package tag

import (
	"bufio"
	"github.com/qibin0506/tml/utils"
)

func NewBatch(ctx *TagContext) TagOp {
	return &Batch{
		Tag: &Tag{
			Root: ctx.Prev.Root,
			Parent: ctx.Prev.Parent,
			Previous: ctx.Prev.Current,
			Current: ctx.Cur,
			Ext: ctx.Prev.Ext,
		},
	}
}

type Batch struct {
	*Tag
}

func (b *Batch) Name() string {
	return "batch"
}

// <batch batch_size="32" shuffle_size="10000" />
func (b *Batch) Parse(writer *bufio.Writer) {
	hasTestData := b.Ext.HasTestData

	batchSizeAttr := b.Current.AttrOr("batch_size", "1")
	shuffleSizeAttr := b.Current.AttrOr("shuffle_size", "1000")

	writer.WriteString("train_dataset = train_dataset.shuffle(")
	writer.WriteString(shuffleSizeAttr)
	writer.WriteString(")")
	writer.WriteString(".batch(")
	writer.WriteString(batchSizeAttr)
	writer.WriteString(")\n")

	if hasTestData {
		writer.WriteString("test_dataset = test_dataset.shuffle(")
		writer.WriteString(shuffleSizeAttr)
		writer.WriteString(")")
		writer.WriteString(".batch(")
		writer.WriteString(batchSizeAttr)
		writer.WriteString(")\n")
	}
	writer.WriteRune('\n')
}

func (b *Batch) Next() TagOp {
	model := utils.GetTagOrFatal(b.Root, "model")
	return tagMap["model"](CreateTagContext(b.Tag, model))
}