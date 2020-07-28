package tag

import (
	"bufio"
	"log"
)

func NewDropout(ctx *TagContext) TagOp {
	return &Dropout{
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

type Dropout struct {
	*HiddenLayer
}

func (d *Dropout) Name() string {
	return "dropout"
}

// <dropout rate="0.2"></dropout>
func (d *Dropout) Parse(writer *bufio.Writer) {
	name, exists := d.Current.Attr("name")
	if !exists {
		log.Fatal("tag dropout must have a name attribute.")
	}

	prevName, exists := d.Previous.Attr("name")
	if !exists {
		log.Fatal("the tag above dropout must have a name attribute.")
	}

	layer := d.Current.AttrOr("layer", prevName)
	rate := d.Current.AttrOr("rate", "0")

	// dropout = tf.keras.layers.Dropout(0.2)(softmax)
	writer.WriteString(name)
	writer.WriteString(" = tf.keras.layers.Dropout(")
	writer.WriteString(rate)
	writer.WriteString(")(")
	writer.WriteString(layer)
	writer.WriteString(")\n")

	d.HiddenLayer.CheckBuildModel(writer)
}

func (d *Dropout) Next() TagOp {
	return d.HiddenLayer.Next()
}