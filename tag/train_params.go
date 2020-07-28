package tag

var losses map[string]string = make(map[string]string)
var optimizers map[string]string = make(map[string]string)
var metrics map[string]string = make(map[string]string)

func init() {
	// losses
	losses["mae"] = "MeanAbsoluteError"
	losses["mse"] = "MeanSquaredError"
	losses["binary_cross_entropy"] = "BinaryCrossentropy"
	losses["categorical_cross_entropy"] = "CategoricalCrossentropy"
	losses["sparse_categorical_cross_entropy"] = "SparseCategoricalCrossentropy"
	losses["bce"] = "BinaryCrossentropy"
	losses["cce"] = "CategoricalCrossentropy"
	losses["scce"] = "SparseCategoricalCrossentropy"

	// optimizers
	optimizers["adam"] = "Adam"
	optimizers["sgd"] = "SGD"
	optimizers["rmsprop"] = "RMSprop"

	// metrics
	metrics["binary_metrics"] = "BinaryAccuracy"
	metrics["categorical_metrics"] = "CategoricalAccuracy"
	metrics["sparse_categorical_metrics"] = "SparseCategoricalAccuracy"
	metrics["bm"] = "BinaryAccuracy"
	metrics["cm"] = "CategoricalAccuracy"
	metrics["scm"] = "SparseCategoricalAccuracy"
}