package tag

import (
	"bufio"
	"log"
	"github.com/qibin0506/tml/utils"
)

func NewCSV(ctx *TagContext) TagOp {
	return &CSV{
		Tag: &Tag{
			Root: ctx.Prev.Root,
			Parent: ctx.Prev.Parent,
			Previous: ctx.Prev.Current,
			Current: ctx.Cur,
			Ext: ctx.Prev.Ext,
		},
	}
}

type CSV struct {
	*Tag
}

func (c *CSV) Name() string {
	return "csv"
}

func (c *CSV) Parse(writer *bufio.Writer) {
	hasTestData := c.Ext.HasTestData

	trainFile, exists := c.Current.Attr("train_file")
	if !exists {
		log.Fatal("tag csv must have a train_file attribute.")
	}

	testFile, exists := c.Current.Attr("test_file")
	if hasTestData && !exists {
		log.Fatal("tag csv must have a test_file attribute.")
	}
	
	imageIdx, exists := c.Current.Attr("image")
	if !exists {
		log.Fatal("tag csv must have a image attribute.")
	}

	labelIdx, exists := c.Current.Attr("label")
	if !exists {
		log.Fatal("tag csv must have a label attribute.")
	}

	writer.WriteString("train_file = csv.reader(open(\"")
	writer.WriteString(trainFile)
	writer.WriteString("\"))\n")

	if hasTestData {
		writer.WriteString("test_file = csv.reader(open(\"")
		writer.WriteString(testFile)
		writer.WriteString("\"))\n")
	}
	writer.WriteRune('\n')

	writer.WriteString("x_train = []\n")
	writer.WriteString("y_train = []\n")

	if hasTestData {
		writer.WriteString("x_test = []\n")
		writer.WriteString("y_test = []\n")
	}
	writer.WriteRune('\n')
	
	writer.WriteString("for train_line in train_file:\n")
	
	writer.WriteString("    ")
	writer.WriteString("x_train.append(train_dir + train_line")
	writer.WriteString(imageIdx)
	writer.WriteString(")\n")

	writer.WriteString("    ")
	writer.WriteString("y_temp = list(map(eval, train_line")
	writer.WriteString(labelIdx)
	writer.WriteString("))\n")
	writer.WriteString("    ")
	writer.WriteString("y_train.append(y_temp)\n\n")

	if hasTestData {
		writer.WriteString("for test_line in test_file:\n")
		
		writer.WriteString("    ")
		writer.WriteString("x_test.append(test_dir + test_line")
		writer.WriteString(imageIdx)
		writer.WriteString(")\n")

		writer.WriteString("    ")
		writer.WriteString("y_temp = list(map(eval, test_line")
		writer.WriteString(labelIdx)
		writer.WriteString("))\n")
		writer.WriteString("    ")
		writer.WriteString("y_test.append(y_temp)\n\n")
	}

	writer.WriteString("x_train = np.array(x_train)\n")
	writer.WriteString("y_train = np.array(y_train)\n")

	if hasTestData {
		writer.WriteString("x_test = np.array(x_test)\n")
		writer.WriteString("y_test = np.array(y_test)\n")
	}
	writer.WriteRune('\n')
}

func (c *CSV) Next() TagOp {
	process := utils.GetTagOrFatal(c.Parent, "process")
	return tagMap["process"](CreateTagContext(c.Tag, process))
}
