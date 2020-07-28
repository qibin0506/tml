package tag

import (
	"bufio"
	"log"
)

func NewMaxPool(ctx *TagContext) TagOp {
	return &MaxPool{
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

type MaxPool struct {
	*HiddenLayer
}

func (m *MaxPool) Name() string {
	return "maxpool"
}

// <maxpool name="pool" size="3" strides="2"></max-pool>
func (m *MaxPool) Parse(writer *bufio.Writer) {
	name, exists := m.Current.Attr("name")
	if !exists {
		log.Fatal("tag maxpool must have a name attribute.")
	}

	prevName, exists := m.Previous.Attr("name")
	if !exists {
		log.Fatal("the tag above maxpool must have a name attribute.")
	}

	layer := m.Current.AttrOr("layer", prevName)
	size := m.Current.AttrOr("size", "3")
	strides := m.Current.AttrOr("strides", "2")

	// pool = tf.keras.layers.MaxPool2D(pool_size=3, strides=2)(lrelu)
	writer.WriteString(name)
	writer.WriteString(" = tf.keras.layers.MaxPool2D(pool_size=")
	writer.WriteString(size)
	writer.WriteString(", strides=")
	writer.WriteString(strides)
	writer.WriteString(")(")
	writer.WriteString(layer)
	writer.WriteString(")\n")

	m.HiddenLayer.CheckBuildModel(writer)
}

func (m *MaxPool) Next() TagOp {
	return m.HiddenLayer.Next()
}