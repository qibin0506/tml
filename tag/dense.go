package tag

import (
	"bufio"
	"log"
)

func NewDense(ctx *TagContext) TagOp {
	return &Dense{
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

type Dense struct {
	*HiddenLayer
}

func (d *Dense) Name() string {
	return "dense"
}

// <dense units="10" name="dense1"></dense>
func (d *Dense) Parse(writer *bufio.Writer) {
	name, exists := d.Current.Attr("name")
	if !exists {
		log.Fatal("tag dense must have a name attribute.")
	}

	prevName, exists := d.Previous.Attr("name")
	if !exists {
		log.Fatal("the tag above dense must have a name attribute.")
	}

	layer := d.Current.AttrOr("layer", prevName)
	units, exists := d.Current.Attr("units")
	if !exists {
		log.Fatal("tag dense must have a unit attribute.")
	}

	activation := d.Current.AttrOr("activation", "")

	// dense = tf.keras.layers.Dense(128)(pool)
	writer.WriteString(name)
	writer.WriteString(" = tf.keras.layers.Dense(units=")
	writer.WriteString(units)
	if activation != "" {
		writer.WriteString(", activation=\"")
		writer.WriteString(activation)
		writer.WriteString("\"")
	}
	writer.WriteString(")(")
	writer.WriteString(layer)
	writer.WriteString(")\n")

	d.HiddenLayer.CheckBuildModel(writer)
}

func (d *Dense) Next() TagOp {
	return d.HiddenLayer.Next()
}