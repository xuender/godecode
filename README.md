GoDecode
========
[![Build Status](https://travis-ci.org/xuender/godecode.png)](https://travis-ci.org/xuender/godecode)
ASCII transliterations of Unicode text for go lang.

## How to Use ##


```shell
github.com/xuender/godecode
```

### decode
```go
var gd godecode.GoDecode
gd.Init("/path/data")
fmt.Println(gd.Decode("南无阿弥陀佛"))
Nan Wu A Mi Tuo Fo

fmt.Println(gd.Decode("一条会走路的鱼"));
Yi Tiao Hui Zou Lu De Yu

fmt.Println(gd.Decode("あみだにょらい"));
amidaniyorai
```
### initials
```go
fmt.Println(gd.Initials("南无阿弥陀佛"));
NWAMTF

fmt.Println(gd.Initials("不怨人就是成佛的大道根"));
BYRJSCFDDDG

fmt.Println(gd.Initials("Κνωσός"));
K
```

## Reference ##

[Unidecode for Java](https://github.com/xuender/unidecode)
[Unidecode for Python](https://pypi.python.org/pypi/Unidecode)
