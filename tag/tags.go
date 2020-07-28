package tag

var tagMap map[string](func(*TagContext) TagOp) = make(map[string](func(*TagContext) TagOp))

func init() {
	tagMap["data"] = NewData
	tagMap["csv"] = NewCSV
	tagMap["process"] = NewProcess
	tagMap["batch"] = NewBatch
	
	tagMap["model"] = NewModel
	tagMap["input"] = NewInput
	tagMap["conv2d"] = NewConv2D
	tagMap["add"] = NewAdd
	tagMap["concat"] = NewConcat
	tagMap["flatten"] = NewFlatten
	tagMap["reshape"] = NewReshape
	tagMap["conv2d-transpose"] = NewConv2DTranspose
	tagMap["softmax"] = NewSoftmax
	tagMap["dropout"] = NewDropout
	tagMap["relu"] = NewRelu
	tagMap["lrelu"] = NewLRelu
	tagMap["maxpool"] = NewMaxPool
	tagMap["dense"] = NewDense
	tagMap["activation"] = NewActivation
	tagMap["batch-normalization"] = NewBatchNormalization
	tagMap["padding"] = NewPadding

	tagMap["train"] = NewTrain
	tagMap["loss"] = NewLoss
	tagMap["metrics"] = NewMetrics
	tagMap["optimizer"] = NewOptimizer
	tagMap["result"] = NewResult
}