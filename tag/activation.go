package tag

import (
	"bufio"
	"log"
)

func NewActivation(ctx *TagContext) TagOp {
	return &Activation{
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

type Activation struct {
	*HiddenLayer
}

func (a *Activation) Name() string {
	return "activation"
}

// <activation name="sigmoid" type="sigmoid"></activation>
func (a *Activation) Parse(writer *bufio.Writer) {
	name, exists := a.Current.Attr("name")
	if !exists {
		log.Fatal("tag activation must have a name attribute.")
	}

	acType, exists := a.Current.Attr("type")
	if !exists {
		log.Fatal("tag activation must have a type attribute.")
	}

	prevName, exists := a.Previous.Attr("name")
	if !exists {
		log.Fatal("the tag above activation must have a name attribute.")
	}

	layer := a.Current.AttrOr("layer", prevName)

	// activation = tf.keras.layers.Activation("sigmoid")(bn)
	writer.WriteString(name)
	writer.WriteString(" = tf.keras.layers.Activation(\"")
	writer.WriteString(acType)
	writer.WriteString("\")(")
	writer.WriteString(layer)
	writer.WriteString(")\n")

	a.HiddenLayer.CheckBuildModel(writer)
}

func (a *Activation) Next() TagOp {
	return a.HiddenLayer.Next()
}