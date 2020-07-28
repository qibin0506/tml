package tag

import (
	"bufio"
	"log"
	"os"
	"github.com/PuerkitoBio/goquery"
	"github.com/qibin0506/tml/utils"
)

type Task struct {
	*Tag
}

func NewTask(dom *goquery.Document) TagOp {
	task := &Task{
		Tag: &Tag{},
	}
	
	taskTag := utils.GetTaskTag(dom)
	task.Root = taskTag
	task.Parent = taskTag
	task.Current = taskTag
	task.Ext = NewTagExt()

	return task
}

func CreateTaskFile(t *Task) (*bufio.Writer, string) {
	nameAttr := t.Current.AttrOr("name", "tml")
	// tmlVersionAttr := taskTag.AttrOr("tml_version", utils.VERSION_NAME)

	utils.CreateDir(utils.BUILD_DIR)
	compiledFileName := utils.BUILD_DIR + nameAttr + ".py"

	if utils.FileExists(compiledFileName) {
		if err := os.Remove(compiledFileName); err != nil {
			log.Fatal(err)
		}
	}

	compiledFile, err := os.OpenFile(compiledFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewWriter(compiledFile), compiledFileName
}

func (t *Task) Name() string {
	return "task"
}

func (t *Task) Parse(writer *bufio.Writer) {
	writer.WriteString("import sys\n")
	writer.WriteString("import csv\n")
	writer.WriteString("import tensorflow as tf\n")
	writer.WriteString("import numpy as np\n\n")

	dtype, exists := t.Current.Attr("dtype")
	if exists {
		writer.WriteString("tf.keras.backend.set_floatx(\"")
		writer.WriteString(dtype)
		writer.WriteString("\")\n\n")
	}
}

func (t *Task) Next() TagOp {
	data := utils.GetTagOrFatal(t.Root, "data")
	return tagMap["data"](CreateTagContext(t.Tag, data))
}