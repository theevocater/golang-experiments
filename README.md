Use [go-torch](https://github.com/uber/go-torch) to create flame graph:

```
go test -bench . -cpuprofile=cpu.prof
go-torch golang-experiments.test cpu.prof
```