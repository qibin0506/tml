package tag

import (
	"bufio"
	"log"
)

func NewPadding(ctx *TagContext) TagOp {
	return &Padding{
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

type Padding struct {
	*HiddenLayer
}

func (p *Padding) Name() string {
	return "padding"
}

// <relu name="relu" layer="dropout"></relu>
func (p *Padding) Parse(writer *bufio.Writer) {
	name, exists := p.Current.Attr("name")
	if !exists {
		log.Fatal("tag relu must have a name attribute.")
	}

	prevName, exists := p.Previous.Attr("name")
	if !exists {
		log.Fatal("the tag above relu must have a name attribute.")
	}

	layer := p.Current.AttrOr("layer", prevName)
	size := p.Current.AttrOr("size", "1")

	// tf.keras.layers.ZeroPadding2D([2, 2])(inputs)
	writer.WriteString(name)
	writer.WriteString(" = tf.keras.layers.ZeroPadding2D([")
	writer.WriteString(size)
	writer.WriteString("])(")
	writer.WriteString(layer)
	writer.WriteString(")\n")

	p.HiddenLayer.CheckBuildModel(writer)
}

func (p *Padding) Next() TagOp {
	return p.HiddenLayer.Next()
}