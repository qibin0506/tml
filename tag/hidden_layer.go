package tag

import (
	"bufio"
	"log"
	"github.com/qibin0506/tml/utils"
)

type HiddenLayer struct {
	*Tag
}

func (h *HiddenLayer) Name() string {
	return "must be implemented"
}

func (h *HiddenLayer) CheckBuildModel(writer *bufio.Writer) {
	nextNode := utils.NextNode(h.Current)
	if nextNode == nil {
		h.buildModel(writer)
	}
}

func (h *HiddenLayer) Next() TagOp {
	nextNode := utils.NextNode(h.Current)
	if nextNode == nil {
		train := utils.GetTagOrFatal(h.Root, "train")
		return tagMap["train"](CreateTagContext(h.Tag, train))
	}

	nextNodeName := utils.NodeName(nextNode)
	nextTagFunc := tagMap[nextNodeName]
	if nextTagFunc == nil {
		log.Fatalf("tag %s was not supported.", nextNodeName)
	}

	return nextTagFunc(CreateTagContext(h.Tag, nextNode))
}

func (h *HiddenLayer) buildModel(writer *bufio.Writer) {
	modelName := h.Parent.AttrOr("name", "model")
	inputs, exists := h.Parent.Attr("inputs")
	if !exists {
		log.Fatal("tag model must have an inputs attribute.")
	}

	outputs, exists := h.Parent.Attr("outputs")
	if !exists {
		log.Fatal("tag model must have an outputs attribute.")
	}

	writer.WriteRune('\n')
	writer.WriteString(modelName)
	writer.WriteString(" = tf.keras.models.Model(inputs=[")
	writer.WriteString(inputs)
	writer.WriteString("], ")
	writer.WriteString("outputs=[")
	writer.WriteString(outputs)
	writer.WriteString("])\n")
	writer.WriteString(modelName)
	writer.WriteString(".summary()\n")

	writer.WriteString("sys.stdout.flush()\n\n")
}