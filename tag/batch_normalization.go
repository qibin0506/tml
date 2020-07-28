package tag

import (
	"bufio"
	"log"
)

func NewBatchNormalization(ctx *TagContext) TagOp {
	return &BatchNormalization{
		&HiddenLayer {
			&Tag{
				Root: ctx.Prev.Root,
				Parent: ctx.Prev.Parent,
				Previous: ctx.Prev.Current,
				Current: ctx.Cur,
				Ext: ctx.Prev.Ext,
			},
		},
	}
}

type BatchNormalization struct {
	*HiddenLayer
}

func (b *BatchNormalization) Name() string {
	return "batch-normalization"
}

// <batch-normalization name="bn"></batch-normalization>
func (b *BatchNormalization) Parse(writer *bufio.Writer) {
	name, exists := b.Current.Attr("name")
	if !exists {
		log.Fatal("tag batch-normalization must have a name attribute.")
	}

	prevName, exists := b.Previous.Attr("name")
	if !exists {
		log.Fatal("the tag above batch-normalization must have a name attribute.")
	}

	layer := b.Current.AttrOr("layer", prevName)

	// bn = tf.keras.layers.BatchNormalization()(dense)
	writer.WriteString(name)
	writer.WriteString(" = tf.keras.layers.BatchNormalization()(")
	writer.WriteString(layer)
	writer.WriteString(")\n")

	b.HiddenLayer.CheckBuildModel(writer)
}

func (b *BatchNormalization) Next() TagOp {
	return b.HiddenLayer.Next()
}