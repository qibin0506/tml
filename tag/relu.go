package tag

import (
	"bufio"
	"log"
)

func NewRelu(ctx *TagContext) TagOp {
	return &Relu{
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

type Relu struct {
	*HiddenLayer
}

func (r *Relu) Name() string {
	return "relu"
}

// <relu name="relu" layer="dropout"></relu>
func (r *Relu) Parse(writer *bufio.Writer) {
	name, exists := r.Current.Attr("name")
	if !exists {
		log.Fatal("tag relu must have a name attribute.")
	}

	prevName, exists := r.Previous.Attr("name")
	if !exists {
		log.Fatal("the tag above relu must have a name attribute.")
	}

	layer := r.Current.AttrOr("layer", prevName)

	// relu = tf.keras.layers.ReLU()(dropout)
	writer.WriteString(name)
	writer.WriteString(" = tf.keras.layers.ReLU()(")
	writer.WriteString(layer)
	writer.WriteString(")\n")

	r.HiddenLayer.CheckBuildModel(writer)
}

func (r *Relu) Next() TagOp {
	return r.HiddenLayer.Next()
}