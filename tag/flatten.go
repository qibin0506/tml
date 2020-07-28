package tag

import (
	"bufio"
	"log"
)

func NewFlatten(ctx *TagContext) TagOp {
	return &Flatten{
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

type Flatten struct {
	*HiddenLayer
}

func (f *Flatten) Name() string {
	return "flatten"
}

// <flatten name="flatten" layer="concat"></flatten>
func (f *Flatten) Parse(writer *bufio.Writer) {
	name, exists := f.Current.Attr("name")
	if !exists {
		log.Fatal("tag flatten must have a name attribute.")
	}

	prevName, exists := f.Previous.Attr("name")
	if !exists {
		log.Fatal("the tag above flatten must have a name attribute.")
	}

	layer := f.Current.AttrOr("layer", prevName)

	// flatten = tf.keras.layers.Flatten()(add)
	writer.WriteString(name)
	writer.WriteString(" = tf.keras.layers.Flatten()(")
	writer.WriteString(layer)
	writer.WriteString(")\n")

	f.HiddenLayer.CheckBuildModel(writer)
}

func (f *Flatten) Next() TagOp {
	return f.HiddenLayer.Next()
}