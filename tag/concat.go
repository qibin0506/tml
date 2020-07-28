package tag

import (
	"bufio"
	"log"
)

func NewConcat(ctx *TagContext) TagOp {
	return &Concat{
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

type Concat struct {
	*HiddenLayer
}

func (c *Concat) Name() string {
	return "concat"
}

// <concat name="concat" layers="conv1,conv2"></concat>
func (c *Concat) Parse(writer *bufio.Writer) {
	name, exists := c.Current.Attr("name")
	if !exists {
		log.Fatal("tag concat must have a name attribute.")
	}

	layers, exists := c.Current.Attr("layers")
	if !exists {
		log.Fatal("tag concat must have a layers attribute.")
	}

	// concat = tf.keras.layers.Concatenate()([x1, x2])
	writer.WriteString(name)
	writer.WriteString(" = tf.keras.layers.Concatenate()([")
	writer.WriteString(layers)
	writer.WriteString("])\n")

	c.HiddenLayer.CheckBuildModel(writer)
}

func (c *Concat) Next() TagOp {
	return c.HiddenLayer.Next()
}