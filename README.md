# gpmap

## Using

see `gpmap_test.go`


### with interface

Define a struct that implemented interface
```
	type Data struct {
		Value string
		Lang  string
	}
	
	// map function 
	func (a *Data) Map() gpmap.MapperFunc {
		return func(v gpmap.PMapData) gpmap.PMapData {
			// do something map function
		}
	}
```

using with interface
```
	for i := 0; i < 100; i++ {
		list[i] = &Data{
			Value: fmt.Sprintf("TEST%d", i),
			Lang:  "ja",
		}
	}

	gpm, _ := gpmap.New(list)
	
	result, err := gpm.Map(func(v gpmap.PMapData) gpmap.PMapData {
		return v.Map()(v)
	})
```

### with anonymous function

Define a structure in which the interface is embedded
```
	type Data struct {
		gpmap.PMapData // embedded data interface
		Value string
		Lang  string
	}
```

```
	for i := 0; i < 100; i++ {
		list[i] = &Data{
			Value: fmt.Sprintf("TEST%d", i),
			Lang:  "ja",
		}
	}

	gpm, _ := gpmap.New(list)

	result, err := gpm.Map(func(v gpmap.PMapData) gpmap.PMapData {
		return &Data{
			Value: strings.ReplaceAll(v.(*Data).Value, "TIKARA", "POWER"),
			Lang:  "en",
		}
	})
```

## Benchmark

```
go test ./... -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/tkitsunai/gpmap
BenchmarkParallelMap_AsyncMap-4          1000000              1128 ns/op             206 B/op          7 allocs/op
BenchmarkParallelMap_SyncMap-4           1000000              1196 ns/op             135 B/op          6 allocs/op
PASS
ok      github.com/tkitsunai/gpmap      2.423s
```

_Caution: Parallel Map does not the effect for small size slice. There is a risk that it will slow down._