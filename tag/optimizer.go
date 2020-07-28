package tag

import (
	"bufio"
	"log"
	"github.com/qibin0506/tml/utils"
)

func NewOptimizer(ctx *TagContext) TagOp {
	return &Optimizer{
		Tag: &Tag{
			Root: ctx.Prev.Root,
			Parent: ctx.Prev.Parent,
			Previous: ctx.Prev.Current,
			Current: ctx.Cur,
			Ext: ctx.Prev.Ext,
		},
	}
}

type Optimizer struct {
	*Tag
}

func (o *Optimizer) Name() string {
	return "optimizer"
}

// <optimizer type="adam" lr="1e-4"></optimizer>
func (o *Optimizer) Parse(writer *bufio.Writer) {
	optimizerKey, exists := o.Current.Attr("type")
	if !exists {
		log.Fatal("tag optimizer must have a type attribute.")
	}

	optimizerValue := optimizers[optimizerKey]
	if optimizerValue == "" {
		log.Fatalf("optimizer type %s is not supported.", optimizerKey)
	}

	lr, exists := o.Current.Attr("lr")
	if !exists {
		log.Fatal("tag optimizer must have a lr attribute.")
	}

	// optimizer = tf.keras.optimizers.Adam(1e-4)
	writer.WriteString("optimizer = tf.keras.optimizers.")
	writer.WriteString(optimizerValue)
	writer.WriteString("(")
	writer.WriteString("learning_rate=")
	writer.WriteString(lr)
	writer.WriteString(")\n\n")

}

func (o *Optimizer) Next() TagOp {
	result := utils.GetTagOrFatal(o.Parent, "result")
	return tagMap["result"](CreateTagContext(o.Tag, result))
}