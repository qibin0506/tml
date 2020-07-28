package tag

import (
	"bufio"
	"log"
)

func NewLRelu(ctx *TagContext) TagOp {
	return &LRelu{
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

type LRelu struct {
	*HiddenLayer
}

func (l *LRelu) Name() string {
	return "lrelu"
}

// <lrelu name="relu" alpha="0.3" layer="dropout"></relu>
func (l *LRelu) Parse(writer *bufio.Writer) {
	name, exists := l.Current.Attr("name")
	if !exists {
		log.Fatal("tag lrelu must have a name attribute.")
	}

	prevName, exists := l.Previous.Attr("name")
	if !exists {
		log.Fatal("the tag above lrelu must have a name attribute.")
	}

	layer := l.Current.AttrOr("layer", prevName)
	alpha := l.Current.AttrOr("alpha", "")

	// lrelu = tf.keras.layers.LeakyReLU(0.3)(dropout)
	writer.WriteString(name)
	writer.WriteString(" = tf.keras.layers.LeakyReLU(")
	if alpha != "" {
		writer.WriteString(alpha)
	}
	writer.WriteString(")(")
	writer.WriteString(layer)
	writer.WriteString(")\n")

	l.HiddenLayer.CheckBuildModel(writer)
}

func (l *LRelu) Next() TagOp {
	return l.HiddenLayer.Next()
}