package main

import (
	"flag"
	"log"
	"github.com/qibin0506/tml/utils"
	"github.com/qibin0506/tml/tag"
)

var (
	train string
	check string
)

func init() {
	flag.StringVar(&train, "train", "", "tml -train some.html")
	flag.StringVar(&check, "check", "env", "tml -check env")
}

func main() {
	flag.Parse()

	if train != "" {
		parseTrainFile(train)
	} else if (check != "") {
		utils.CheckEnv()
	}
}

func parseTrainFile(filePath string) {
	dom := utils.ReadFileAsDom(filePath)
	
	task := tag.NewTask(dom)
	taskTag, ok := task.(*tag.Task)
	if !ok {
		log.Fatal("the first tag must be task.")
	}

	writer, compileFilePath := tag.CreateTaskFile(taskTag)
	task.Parse(writer)

	for tag := task.Next(); tag != nil; tag = tag.Next() {
		tag.Parse(writer)
	}

	writer.Flush()

	startTrain(compileFilePath)
}

func startTrain(fileName string) {
	utils.ExecCommand("python3", fileName)
}