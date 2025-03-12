## topohub集成pprof

### 本地查看pprof
启动程序
```
.PHONY: local-pprof
local-pprof:
go run cmd/topohub/main.go    --profiling-port=8084
```

在应用启动之后，可以打开浏览器，输入`http://localhost:port/debug/pprof/`， 可以看到UI界面
通过在url后缀可以查看不同的指标，如:"/debug/pprof/heap", "/debug/pprof/goroutine"

### 本地保存pprof文件

将pprof的文件保存下来以便分析：

```bash
go tool pprof http://localhost:port/debug/pprof/profile?seconds=30
```

运行之后会收集30s左右的信息，然后进入交互模式，通过输入web即可生成svg文件。同时对应的pprof文件会被保存在本地。其他类型的指标类似。


### 容器中保存pprof文件
修改deployment文件添加启动pprof参数:

```
     containers:
      - args:
        - --metrics-port=8083
        - --health-probe-port=8081
        - --webhook-port=8082
        - --profiling-port=8084
```

在宿主机上执行curl命令，下载profile文件，传到本地进行分析。
```bash
curl -o profile.pb.gz  http://podip:8084/debug/pprof/profile?seconds=30
```

### 分析pprof文件

想重复分析pprof文件可以直接使用文件打开：

```bash
 go tool pprof profile filename-xxxx
```

打开pprof 文件之后，可以使用`web` 或者 `top`命令查看

### ref
https://github.com/google/pprof/blob/main/doc/README.md
