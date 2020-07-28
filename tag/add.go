package tag

import (
	"bufio"
	"log"
)

func NewAdd(ctx *TagContext) TagOp {
	return &Add{
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

type Add struct {
	*HiddenLayer
}

func (a *Add) Name() string {
	return "add"
}

// <add layers="conv1,conv2"></add>
func (a *Add) Parse(writer *bufio.Writer) {
	name, exists := a.Current.Attr("name")
	if !exists {
		log.Fatal("tag add must have a name attribute.")
	}

	layers, exists := a.Current.Attr("layers")
	if !exists {
		log.Fatal("tag add must have a layers attribute.")
	}

	// add = tf.keras.layers.Add()([x1, x2])
	writer.WriteString(name)
	writer.WriteString(" = tf.keras.layers.Add()([")
	writer.WriteString(layers)
	writer.WriteString("])\n")

	a.HiddenLayer.CheckBuildModel(writer)
}

func (a *Add) Next() TagOp {
	return a.HiddenLayer.Next()
}