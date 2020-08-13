# ttools
my tools collection

### 1. 单词格式转换工具
```
➜  ttools git:(master) go run main.go help word
该子命令支持各种单词格式转换，模式如下：
1: 全部单词转换为大写
2: 全部单词转换为小写
3: 下划线单词转换为大写驼峰单词
4: 下划线单词转换为小写驼峰单词
5: 驼峰单词转换为下划线单词

Usage:
   word [flags]

Flags:
  -h, --help         help for word
  -m, --mode int8    请输入单词转换模式
  -s, --str string   请输入单词内容
```

- 使用说明
```
➜  ttools git:(master) go run main.go word -s=eafwf -m=1
2020/08/13 11:04:09 输出结果: EAFWF
```

### 2. 便捷时间工具

```
➜  ttools git:(master) ✗ go run main.go help time
时间格式处理

Usage:
   time [flags]
   time [command]

Available Commands:
  calc        计算所需时间
  now         获取当前时间

Flags:
  -h, --help   help for time

Use " time [command] --help" for more information about a command.
```

- 使用说明
```
➜  ttools git:(master) ✗ go run main.go time now
2020/08/13 14:19:28 输出结果: 2020-08-13 14:19:28, 1597299568
```
```
➜  ttools git:(master) ✗ go run main.go time calc -c="2029-09-04 12:02:03" -d=5m
2020/08/13 14:21:44 输出结果: 2029-09-04 12:07:03, 1883218023
```
