# hajimeteno-self-reference-golang

Result

```shell
go run main.go
```

```shell
{"level":"info","ts":1746001094.6603029,"caller":"hajimeteno-self-reference-golang/main.go:36","msg":"Is some.parent == st?","result":true}
{"level":"info","ts":1746001094.660334,"caller":"hajimeteno-self-reference-golang/main.go:40","msg":"Is some.parent.some == st.some?","result":true}
{"level":"info","ts":1746001094.6603417,"caller":"hajimeteno-self-reference-golang/main.go:44","msg":"Is some.parent.some.parent == st?","result":true}
{"level":"info","ts":1746001094.6603491,"caller":"hajimeteno-self-reference-golang/main.go:50","msg":"Is st == st2?","result":false}
```

