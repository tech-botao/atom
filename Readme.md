# atom

> 只有输入输出的函数写在这里， 不需要依赖其它

## 功能

### Env

> 环境变量的批量取得

```golang
atom.SetPath("/etc/go")
atom.EnvLoad("atmo.env","opts.env")
println(atom.EnvOrPanic("APP_ENV"))
```

### convert 

> 各种变换

```golang
	var v interface{}
	var s = `{"a":1, "b":2}`
	err := DecodeFromString(s, &v)

	if err != nil {
		logger.Error("DecodeFromString()", err)
		return
	}

	logger.Info("decode "+s+" => ", v)
```
