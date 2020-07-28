package tag

import (
	"bufio"
	"log"
)

func NewSoftmax(ctx *TagContext) TagOp {
	return &Softmax{
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

type Softmax struct {
	*HiddenLayer
}

func (s *Softmax) Name() string {
	return "softmax"
}

// softmax name="softmax" layer="conv-transpose"></softmax>
func (s *Softmax) Parse(writer *bufio.Writer) {
	name, exists := s.Current.Attr("name")
	if !exists {
		log.Fatal("tag softmax must have a name attribute.")
	}

	prevName, exists := s.Previous.Attr("name")
	if !exists {
		log.Fatal("the tag above softmax must have a name attribute.")
	}

	layer := s.Current.AttrOr("layer", prevName)

	// softmax = tf.keras.layers.Softmax()(upsample)
	writer.WriteString(name)
	writer.WriteString(" = tf.keras.layers.Softmax()(")
	writer.WriteString(layer)
	writer.WriteString(")\n")

	s.HiddenLayer.CheckBuildModel(writer)
}

func (s *Softmax) Next() TagOp {
	return s.HiddenLayer.Next()
}