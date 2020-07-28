import sys
import csv
import tensorflow as tf
import numpy as np

# start prepare data.

train_dir = "./images/train/"
test_dir = "./images/test/"

train_file = csv.reader(open("./images/train.csv"))
test_file = csv.reader(open("./images/test.csv"))

x_train = []
y_train = []
x_test = []
y_test = []

for train_line in train_file:
    x_train.append(train_dir + train_line[0])
    y_temp = list(map(eval, train_line[1:3]))
    y_train.append(y_temp)

for test_line in test_file:
    x_test.append(test_dir + test_line[0])
    y_temp = list(map(eval, test_line[1:3]))
    y_test.append(y_temp)

x_train = np.array(x_train)
y_train = np.array(y_train)
x_test = np.array(x_test)
y_test = np.array(y_test)

train_dataset = tf.data.Dataset.from_tensor_slices((x_train, y_train))
test_dataset = tf.data.Dataset.from_tensor_slices((x_test, y_test))

def train_process(x, y):
    x = tf.io.read_file(x)
    x = tf.image.decode_image(x)
    x = tf.image.resize_with_pad(x, 3, 3)
    x = tf.keras.backend.cast_to_floatx(x)
    x = x / 255.0
    y = tf.keras.backend.cast_to_floatx(y)
    y = y / 28.0

    return x, y

def test_process(x, y):
    x = tf.io.read_file(x)
    x = tf.image.decode_image(x)
    x = tf.image.resize_with_pad(x, 3, 3)
    x = tf.keras.backend.cast_to_floatx(x)
    x = x / 255.0
    y = tf.keras.backend.cast_to_floatx(y)
    y = y / 28.0

    return x, y

train_dataset = train_dataset.map(train_process)
test_dataset = test_dataset.map(test_process)

train_dataset = train_dataset.shuffle(10000).batch(32)
test_dataset = test_dataset.shuffle(10000).batch(32)

# start build model.

input1 = tf.keras.layers.Input(shape=[28,28,1])
conv1 = tf.keras.layers.Conv2D(filters=16, kernel_size=3, strides=2, activation="relu", padding="same")(input1)
conv2 = tf.keras.layers.Conv2D(filters=32, kernel_size=3, activation="relu")(conv1)
conv3 = tf.keras.layers.Conv2D(filters=32, kernel_size=3, activation="relu")(conv2)
flatten = tf.keras.layers.Flatten()(conv3)
relu = tf.keras.layers.LeakyReLU(0.3)(flatten)
dense1 = tf.keras.layers.Dense(units=10, activation="relu")(relu)
dense2 = tf.keras.layers.Dense(units=10)(dense1)

model = tf.keras.models.Model(inputs=[input1], outputs=[dense2])
model.summary()
sys.stdout.flush()

# start train data.

loss_object = tf.keras.losses.BinaryCrossentropy()
train_loss = tf.keras.metrics.Mean(name="train_loss")
test_loss = tf.keras.metrics.Mean(name="test_loss")

train_accuracy = tf.keras.metrics.BinaryAccuracy(name="train_accuracy")
test_accuracy = tf.keras.metrics.BinaryAccuracy(name="test_accuracy")

optimizer = tf.keras.optimizers.Adam(learning_rate=1e-4)

@tf.function
def train_step(x, y):
    with tf.GradientTape() as tape:
        y_pred = model(x)
        loss_value = loss_object(y_true=y, y_pred=y_pred)

    gradients = tape.gradient(loss_value, model.trainable_variables)
    optimizer.apply_gradients(zip(gradients, model.trainable_variables))

    train_loss(loss_value)
    train_accuracy(y_true=y, y_pred=y_pred)

@tf.function
def test_step(x, y):
    y_pred = model(x)
    loss_value = loss_object(y_true=y, y_pred=y_pred)

    test_loss(loss_value)
    test_accuracy(y_true=y, y_pred=y_pred)

for epoch in range(100):
    for x, y in train_dataset:
        train_step(x, y)

    for x, y in test_dataset:
        test_step(x, y)

    print("epoch:{}, train_loss:{}, train_accuracy:{}, test_loss:{}, test_accuracy:{}.".format(epoch + 1, train_loss.result(), train_accuracy.result() * 100, test_loss.result(), test_accuracy.result() * 100))

    if epoch >=0:
        print("save result: ./result/model_{}.h5".format(epoch + 1))
        model.save("./result/model_{}.h5".format(epoch + 1))

    sys.stdout.flush()
