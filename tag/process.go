package tag

import (
	"bufio"
	"github.com/qibin0506/tml/utils"
)

func NewProcess(ctx *TagContext) TagOp {
	return &Process{
		Tag: &Tag{
			Root: ctx.Prev.Root,
			Parent: ctx.Prev.Parent,
			Previous: ctx.Prev.Current,
			Current: ctx.Cur,
			Ext: ctx.Prev.Ext,
		},
	}
}

type Process struct {
	*Tag
}

func (p *Process) Name() string {
	return "process"
}

func (p *Process) Parse(writer *bufio.Writer) {
	isBuildIn := p.Parent.AttrOr("type", "buildin") == "buildin"

	writer.WriteString("train_dataset = tf.data.Dataset.from_tensor_slices((x_train, y_train))\n")
	if p.Ext.HasTestData {
		writer.WriteString("test_dataset = tf.data.Dataset.from_tensor_slices((x_test, y_test))\n")
	}
	writer.WriteRune('\n')

	p.writeProcessFunc(writer, isBuildIn, true)
	p.writeProcessFunc(writer, isBuildIn, false)

	writer.WriteString("train_dataset = train_dataset.map(train_process)\n")
	if p.Ext.HasTestData {
		writer.WriteString("test_dataset = test_dataset.map(test_process)\n\n")
	}
}

func (p *Process) Next() TagOp {
	batch := utils.GetTagOrFatal(p.Root, "batch")
	return tagMap["batch"](CreateTagContext(p.Tag, batch))
}

// <process channels="3" resize="3, 3" x_mean="/ 255.0" y_mean="/ 28.0"></process>
func (p *Process) writeProcessFunc(writer *bufio.Writer, isBuildIn bool, isTrain bool) {
	if !isTrain && !p.Ext.HasTestData {
		return
	}

	xMean, xMeanExists := p.Current.Attr("x_mean")
	yMean, yMeanExists := p.Current.Attr("y_mean")
	resize, resizeExists := p.Current.Attr("resize")
	cropResize, cropResizeExists := p.Current.Attr("crop_resize")

	if (isBuildIn && !xMeanExists && !yMeanExists && !resizeExists && !cropResizeExists) {
		return
	}

	if isTrain {
		writer.WriteString("def train_process(x, y):\n")
	} else {
		writer.WriteString("def test_process(x, y):\n")
	}
	
	if (!isBuildIn) {
		writer.WriteString("    ")
		writer.WriteString("x = tf.io.read_file(x)\n")
		
		channels, exists := p.Parent.Attr("channels")
		writer.WriteString("    ")
		writer.WriteString("x = tf.image.decode_image(x")
		if exists {
			writer.WriteString(",")
			writer.WriteString(channels)
		}
		writer.WriteString(")\n")
	}

	if resizeExists {
		writer.WriteString("    ")
		writer.WriteString("x = tf.image.resize_with_pad(x, ")
		writer.WriteString(resize)
		writer.WriteString(")\n")
	}

	if cropResizeExists {
		writer.WriteString("    ")
		writer.WriteString("x = tf.image.resize_with_crop_or_pad(x, ")
		writer.WriteString(cropResize)
		writer.WriteString(")\n")
	}

	if (xMeanExists) {
		writer.WriteString("    ")
		writer.WriteString("x = tf.keras.backend.cast_to_floatx(x)\n")

		writer.WriteString("    ")
		writer.WriteString("x = x ")
		writer.WriteString(xMean)
		writer.WriteString("\n")
	}

	if (yMeanExists) {
		writer.WriteString("    ")
		writer.WriteString("y = tf.keras.backend.cast_to_floatx(y)\n")

		writer.WriteString("    ")
		writer.WriteString("y = y ")
		writer.WriteString(yMean)
		writer.WriteString("\n")
	}

	writer.WriteString("\n    ")
	writer.WriteString("return x, y\n\n")
}