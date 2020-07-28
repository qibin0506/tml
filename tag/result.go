package tag

import (
	"bufio"
	"strings"
	"log"
	"github.com/qibin0506/tml/utils"
)

func NewResult(ctx *TagContext) TagOp {
	return &Result{
		Tag: &Tag{
			Root: ctx.Prev.Root,
			Parent: ctx.Prev.Parent,
			Previous: ctx.Prev.Current,
			Current: ctx.Cur,
			Ext: ctx.Prev.Ext,
		},
	}
}

type Result struct {
	*Tag
}

func (r *Result) Name() string {
	return "result"
}

// <result save_path="./result" ckpt="%2=0"></result>
func (r *Result) Parse(writer *bufio.Writer) {
	model := utils.GetTagOrFatal(r.Root, "model")
	modelName := model.AttrOr("name", "model")

	r.writeTrainStep(modelName, writer)
	if r.Ext.HasTestData {
		r.writeTestStep(modelName, writer)
	}
	r.writeTrainEpoch(modelName, writer)
}

func (r *Result) writeTrainStep(modelName string, writer *bufio.Writer) {
	writer.WriteString("@tf.function\n")
	writer.WriteString("def train_step(x, y):\n")

	writer.WriteString("    ")
	writer.WriteString("with tf.GradientTape() as tape:\n")

	writer.WriteString("        ")
	writer.WriteString("y_pred = ")
	writer.WriteString(modelName)
	writer.WriteString("(x)\n")

	writer.WriteString("        ")
	writer.WriteString("loss_value = loss_object(y_true=y, y_pred=y_pred)\n\n")

	writer.WriteString("    ")
	writer.WriteString("gradients = tape.gradient(loss_value, ")
	writer.WriteString(modelName)
	writer.WriteString(".trainable_variables)\n")

	writer.WriteString("    ")
	writer.WriteString("optimizer.apply_gradients(zip(gradients, ")
	writer.WriteString(modelName)
	writer.WriteString(".trainable_variables))\n\n")

	writer.WriteString("    ")
	writer.WriteString("train_loss(loss_value)\n")

	writer.WriteString("    ")
	writer.WriteString("train_accuracy(y_true=y, y_pred=y_pred)\n\n")
}

func (r *Result) writeTestStep(modelName string, writer *bufio.Writer) {
	writer.WriteString("@tf.function\n")
	writer.WriteString("def test_step(x, y):\n")

	writer.WriteString("    ")
	writer.WriteString("y_pred = ")
	writer.WriteString(modelName)
	writer.WriteString("(x)\n")

	writer.WriteString("    ")
	writer.WriteString("loss_value = loss_object(y_true=y, y_pred=y_pred)\n\n")

	writer.WriteString("    ")
	writer.WriteString("test_loss(loss_value)\n")

	writer.WriteString("    ")
	writer.WriteString("test_accuracy(y_true=y, y_pred=y_pred)\n\n")
}

func (r *Result) writeTrainEpoch(modelName string, writer *bufio.Writer) {
	epoch, exists := r.Parent.Attr("epoch")
	if !exists {
		log.Fatal("tag train must have a epoch attribute.")
	}

	ckptConditions := r.Current.AttrOr("ckpt", "epoch >= 0")
	ckptConditions = strings.TrimPrefix(ckptConditions, "epoch")
	savePath := r.Current.AttrOr("save_result", "./result")
	if !strings.HasSuffix(ckptConditions, "/") {
		savePath += "/"
	}
	
	utils.CreateDir(savePath)

	writer.WriteString("for epoch in range(")
	writer.WriteString(epoch)
	writer.WriteString("):\n")

	writer.WriteString("    ")
	writer.WriteString("for x, y in train_dataset:\n")
	writer.WriteString("        ")
	writer.WriteString("train_step(x, y)\n\n")

	if r.Ext.HasTestData {
		writer.WriteString("    ")
		writer.WriteString("for x, y in test_dataset:\n")
		writer.WriteString("        ")
		writer.WriteString("test_step(x, y)\n\n")
	}

	var printFmt string = "print(\"epoch:{}, train_loss:{}, train_accuracy:{}"
	var printValue string = ".format(epoch + 1, train_loss.result(), train_accuracy.result() * 100"
	
	if r.Ext.HasTestData {
		printFmt += ", test_loss:{}, test_accuracy:{}"
		printValue += ", test_loss.result(), test_accuracy.result() * 100"
	}
	printFmt += ".\""
	printValue += "))\n\n"

	writer.WriteString("    ")
	writer.WriteString(printFmt)
	writer.WriteString(printValue)

	writer.WriteString("    ")
	writer.WriteString("if epoch ")
	writer.WriteString(ckptConditions)
	writer.WriteString(":\n")

	writer.WriteString("        ")
	writer.WriteString("print(\"save result: ")
	writer.WriteString(savePath)
	writer.WriteString(modelName)
	writer.WriteString("_{}.h5\".format(epoch + 1")
	writer.WriteString("))\n")

	writer.WriteString("        ")
	writer.WriteString(modelName)
	writer.WriteString(".save(\"")
	writer.WriteString(savePath)
	writer.WriteString(modelName)
	writer.WriteString("_{}.h5\".format(epoch + 1")
	writer.WriteString("))\n\n")

	writer.WriteString("    ")
	writer.WriteString("sys.stdout.flush()\n")
}

func (r *Result) Next() TagOp {
	// end
	return nil
}