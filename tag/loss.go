package tag

import (
	"bufio"
	"log"
	"github.com/qibin0506/tml/utils"
)

func NewLoss(ctx *TagContext) TagOp {
	return &Loss{
		Tag: &Tag{
			Root: ctx.Prev.Root,
			Parent: ctx.Prev.Parent,
			Previous: ctx.Prev.Current,
			Current: ctx.Cur,
			Ext: ctx.Prev.Ext,
		},
	}
}

type Loss struct {
	*Tag
}

func (l *Loss) Name() string {
	return "loss"
}

func (l *Loss) Parse(writer *bufio.Writer) {
	lossKey, exists := l.Current.Attr("type")
	if !exists {
		log.Fatal("tag loss must have a type attribute.")
	}

	lossValue := losses[lossKey]
	if lossValue == "" {
		log.Fatalf("loss type %s is not supported.", lossKey)
	}

	// loss_object = tf.keras.losses.SparseCategoricalCrossentropy()
	writer.WriteString("loss_object = tf.keras.losses.")
	writer.WriteString(lossValue)
	writer.WriteString("()\n")

	// train_loss = tf.keras.metrics.Mean(name="train_loss")
	writer.WriteString("train_loss = tf.keras.metrics.Mean(name=\"train_loss\")\n")
	if l.Ext.HasTestData {
		writer.WriteString("test_loss = tf.keras.metrics.Mean(name=\"test_loss\")\n")
	}
	writer.WriteRune('\n')
}

func (l *Loss) Next() TagOp {
	metrics := utils.GetTagOrFatal(l.Parent, "metrics")
	return tagMap["metrics"](CreateTagContext(l.Tag, metrics))
}