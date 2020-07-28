package tag

import (
	"bufio"
	"log"
)

func NewReshape(ctx *TagContext) TagOp {
	return &Reshape{
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

type Reshape struct {
	*HiddenLayer
}

func (r *Reshape) Name() string {
	return "reshape"
}

// <reshape name="reshape" target_shape="128,128,3" layer="flatten" />
func (r *Reshape) Parse(writer *bufio.Writer) {
	name, exists := r.Current.Attr("name")
	if !exists {
		log.Fatal("tag reshape must have a name attribute.")
	}

	targetShape, exists := r.Current.Attr("target_shape")
	if !exists {
		log.Fatal("tag reshape must have a target_shape attribute.")
	}

	prevName, exists := r.Previous.Attr("name")
	if !exists {
		log.Fatal("the tag above reshape must have a name attribute.")
	}

	layer := r.Current.AttrOr("layer", prevName)

	// reshape = tf.keras.layers.Reshape(target_shape=[128, 128, 3])(flatten)
	writer.WriteString(name)
	writer.WriteString(" = tf.keras.layers.Reshape(target_shape=[")
	writer.WriteString(targetShape)
	writer.WriteString("])(")
	writer.WriteString(layer)
	writer.WriteString(")\n")

	r.HiddenLayer.CheckBuildModel(writer)
}

func (r *Reshape) Next() TagOp {
	return r.HiddenLayer.Next()
}