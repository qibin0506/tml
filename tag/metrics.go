package tag

import (
	"bufio"
	"log"
	"github.com/qibin0506/tml/utils"
)

func NewMetrics(ctx *TagContext) TagOp {
	return &Metrics{
		Tag: &Tag{
			Root: ctx.Prev.Root,
			Parent: ctx.Prev.Parent,
			Previous: ctx.Prev.Current,
			Current: ctx.Cur,
			Ext: ctx.Prev.Ext,
		},
	}
}

type Metrics struct {
	*Tag
}

func (m *Metrics) Name() string {
	return "metrics"
}

// <metrics type="sparse_categorical"></metrics>
func (m *Metrics) Parse(writer *bufio.Writer) {
	metricsKey, exists := m.Current.Attr("type")
	if !exists {
		log.Fatal("tag metrics must have a type attribute.")
	}

	metricsValue := metrics[metricsKey]
	if metricsValue == "" {
		log.Fatalf("metrics type %s is not supported.", metricsKey)
	}

	// train_accuracy = tf.keras.metrics.SparseCategoricalAccuracy(name="train_accuracy")
	writer.WriteString("train_accuracy = tf.keras.metrics.")
	writer.WriteString(metricsValue)
	writer.WriteString("(name=\"train_accuracy\")\n")

	if m.Ext.HasTestData {
		writer.WriteString("test_accuracy = tf.keras.metrics.")
		writer.WriteString(metricsValue)
		writer.WriteString("(name=\"test_accuracy\")\n")
	}
	writer.WriteRune('\n')
}

func (m *Metrics) Next() TagOp {
	optimizer := utils.GetTagOrFatal(m.Parent, "optimizer")
	return tagMap["optimizer"](CreateTagContext(m.Tag, optimizer))
}