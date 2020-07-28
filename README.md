# tml
train a neural network by using HTML.

## 1. data process

use buildin dataset.
``` html
<data type="buildin" src="fashion_mnist">
    <process x_mean="/ 255.0" y_mean="/ 28.0"></process>
    <batch batch_size="32" shuffle_size="10000"></batch>
</data>
```

## 2. 