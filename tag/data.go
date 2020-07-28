package tag

import (
	"bufio"
	"log"
	"github.com/qibin0506/tml/utils"
)

func NewData(ctx *TagContext) TagOp {
	return &Data{
		Tag: &Tag{
			Root: ctx.Prev.Root,
			Parent: ctx.Cur,
			Previous: ctx.Prev.Current,
			Current: ctx.Cur,
			Ext: ctx.Prev.Ext,
		},
	}
}

type Data struct {
	*Tag
}

func (d *Data) Name() string {
	return "data"
}

func (d *Data) Parse(writer *bufio.Writer) {
	dtype, exist := d.Current.Attr("dtype")
	if exist {
		writer.WriteString("tf.keras.backend.set_floatx(\"")
		writer.WriteString(dtype)
		writer.WriteString("\")\n\n")
	}
	
	writer.WriteString("# start prepare data.\n\n")

	typeAttr, exist := d.Current.Attr("type")
	if !exist {
		log.Fatal("tag data must have a type attribute.")
	}

	switch typeAttr {
		case "buildin":
			d.parseBuildinDataLoader(writer)
		case "local":
			d.parseLocalDataLoader(writer)
		default:
			log.Fatalf("tag data type attribute value %s was not supported.", typeAttr)
	}
}

func (d *Data) Next() TagOp {
	typeAttr := d.Current.AttrOr("type", "buildin")
	if typeAttr == "local" {
		csv := utils.GetTagOrFatal(d.Parent, "csv")
		return tagMap["csv"](CreateTagContext(d.Tag, csv))
	}

	process := utils.GetTagOrFatal(d.Parent, "process")
	return tagMap["process"](CreateTagContext(d.Tag, process))
}

func (d *Data) parseBuildinDataLoader(writer *bufio.Writer) {
	srcAttr := d.Current.AttrOr("src", "mnist")
	writer.WriteString("dataset = tf.keras.datasets.")
	writer.WriteString(srcAttr)
	writer.WriteRune('\n')
	writer.WriteString("(x_train, y_train), (x_test, y_test) = dataset.load_data()\n")
	writer.WriteString("x_train = x_train[..., tf.newaxis]\n")
	writer.WriteString("x_test = x_test[..., tf.newaxis]\n\n")

	d.Ext.HasTestData = true
}

func (d *Data) parseLocalDataLoader(writer *bufio.Writer) {
	trainDir := d.Current.AttrOr("train_dir", "")
	testDir, hasTestData := d.Current.Attr("test_dir")

	writer.WriteString("train_dir = \"")
	writer.WriteString(trainDir)
	writer.WriteString("\"\n")
	
	if hasTestData {
		writer.WriteString("test_dir = \"")
		writer.WriteString(testDir)
		writer.WriteString("\"\n")
	}

	writer.WriteRune('\n')

	d.Ext.HasTestData = hasTestData
}