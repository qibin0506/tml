package tag

import (
	"bufio"
	"log"
)

func NewConv2D(ctx *TagContext) TagOp {
	return &Conv2D{
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

type Conv2D struct {
	*HiddenLayer
}

func (c *Conv2D) Name() string {
	return "conv2d"
}

// <conv2d filters="32" kernel="3" name="conv2" activation="relu" input="conv1" />
func (c *Conv2D) Parse(writer *bufio.Writer) {
	name, exists := c.Current.Attr("name")
	if !exists {
		log.Fatal("tag conv2d must have a name attribute.")
	}

	filters, exists := c.Current.Attr("filters")
	if !exists {
		log.Fatal("tag conv2d must have a filters attribute.")
	}

	kernel, exists := c.Current.Attr("kernel")
	if !exists {
		log.Fatal("tag conv2d must have a kernel attribute.")
	}

	activation := c.Current.AttrOr("activation", "")

	prevName, exists := c.Previous.Attr("name")
	if !exists {
		log.Fatal("the tag above conv2d must have a name attribute.")
	}

	layer := c.Current.AttrOr("layer", prevName)
	strides := c.Current.AttrOr("strides", "")
	padding := c.Current.AttrOr("padding", "")

	// x1 = tf.keras.layers.Conv2D(filters=16, kernel_size=3, activation='relu')(inputs)
	writer.WriteString(name)
	writer.WriteString(" = tf.keras.layers.Conv2D(filters=")
	writer.WriteString(filters)
	writer.WriteString(", kernel_size=")
	writer.WriteString(kernel)
	if strides != "" {
		writer.WriteString(", strides=")
		writer.WriteString(strides)
	}
	if activation != "" {
		writer.WriteString(", activation=\"")
		writer.WriteString(activation)
		writer.WriteString("\"")
	}
	if padding != "" {
		writer.WriteString(", padding=\"")
		writer.WriteString(padding)
		writer.WriteString("\"")
	}
	writer.WriteString(")(")
	writer.WriteString(layer)
	writer.WriteString(")\n")

	c.HiddenLayer.CheckBuildModel(writer)
}

func (c *Conv2D) Next() TagOp {
	return c.HiddenLayer.Next()
}