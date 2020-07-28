package tag

import (
	"bufio"
	"log"
	"github.com/qibin0506/tml/utils"
)

func NewInput(ctx *TagContext) TagOp {
	return &Input{
		Tag: &Tag{
			Root: ctx.Prev.Root,
			Parent: ctx.Prev.Parent,
			Previous: ctx.Prev.Current,
			Current: ctx.Cur,
			Ext: ctx.Prev.Ext,
		},
	}
}

type Input struct {
	*Tag
}

func (i *Input) Name() string {
	return "input"
}

// <input shape="h,w,c" name="input" />
func (i *Input) Parse(writer *bufio.Writer) {
	shapeAttr, exists := i.Current.Attr("shape")
	if !exists {
		log.Fatal("tag input must have a shape attribute.")
	}

	nameAttr, exists := i.Current.Attr("name")
	if !exists {
		log.Fatal("tag input must have a name attribute.")
	}

	// inputs = tf.keras.layers.Input(shape=shape)
	writer.WriteString(nameAttr)
	writer.WriteString(" = tf.keras.layers.Input(shape=[")
	writer.WriteString(shapeAttr)
	writer.WriteString("])\n")
}

func (i *Input) Next() TagOp {
	nextNode := utils.NextNode(i.Current)
	nextNodeName := utils.NodeName(nextNode)
	if nextNodeName == "" {
		log.Fatal("below tag input must have at least one other tag, e.g. conv2d.")
	}

	nextTagFunc := tagMap[nextNodeName]
	if nextTagFunc == nil {
		log.Fatalf("tag %s was not supported.", nextNodeName)
	}

	return nextTagFunc(CreateTagContext(i.Tag, nextNode))
}