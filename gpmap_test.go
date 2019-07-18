package gpmap_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/tkitsunai/gpmap"
)

func BenchmarkParallelMap_AsyncMap(b *testing.B) {
	b.ResetTimer()

	limit := b.N

	list := make([]gpmap.PMapData, limit)

	for i := 0; i < limit; i++ {
		list[i] = &Data{
			Value: fmt.Sprintf("TIKARA%d", i),
			Lang:  "ja",
		}
	}

	gpm, _ := gpmap.New(list)

	_, _ = gpm.Map(func(v gpmap.PMapData) gpmap.PMapData {
		return v.Map()(v)
	})
}

func BenchmarkParallelMap_SyncMap(b *testing.B) {
	b.ResetTimer()
	limit := b.N
	list := make([]gpmap.PMapData, limit)

	for i := 0; i < limit; i++ {
		list[i] = &Data{
			Value: fmt.Sprintf("TIKARA%d", i),
			Lang:  "ja",
		}
	}

	gpm, _ := gpmap.New(list)

	_, _ = gpm.SyncMap(func(v gpmap.PMapData) gpmap.PMapData {
		return v.Map()(v)
	})
}

func TestGpMapHowToUse_MapperFunc(t *testing.T) {
	list := make([]gpmap.PMapData, 100)

	for i := 0; i < 100; i++ {
		list[i] = &Data{
			Value: fmt.Sprintf("TIKARA%d", i),
			Lang:  "ja",
		}
	}

	gpm, err := gpmap.New(list)

	if err != nil {
		t.Fatal(err)
	}

	actual, err := gpm.Map(func(v gpmap.PMapData) gpmap.PMapData {
		return &Data{
			Value: strings.ReplaceAll(v.(*Data).Value, "TIKARA", "POWER"),
			Lang:  "en",
		}
	})

	if err != nil {
		t.Fatal(err)
	}

	expected := make([]gpmap.PMapData, 100)
	for i := 0; i < 100; i++ {
		expected[i] = &Data{
			Value: fmt.Sprintf("POWER%d", i),
			Lang:  "en",
		}
	}

	assert.True(t, assert.ObjectsAreEqualValues(expected, actual))
}

func TestGpMapHowToUse_MapperFuncInterface(t *testing.T) {

	list := make([]gpmap.PMapData, 100)

	for i := 0; i < 100; i++ {
		list[i] = &Data{
			Value: fmt.Sprintf("TIKARA%d", i),
			Lang:  "ja",
		}
	}

	gpm, err := gpmap.New(list)

	if err != nil {
		t.Fatal(err)
	}

	actual, err := gpm.Map(func(v gpmap.PMapData) gpmap.PMapData {
		return v.Map()(v)
	})

	if err != nil {
		t.Fatal(err)
	}

	expected := make([]gpmap.PMapData, 100)
	for i := 0; i < 100; i++ {
		expected[i] = &Data{
			Value: fmt.Sprintf("POWER%d", i),
			Lang:  "en",
		}
	}

	assert.True(t, assert.ObjectsAreEqualValues(expected, actual))
}

type Data struct {
	Value string
	Lang  string
}

func (a *Data) Map() gpmap.MapperFunc {
	return func(v gpmap.PMapData) gpmap.PMapData {
		time.Sleep(5)
		result := a
		result.Value = strings.ReplaceAll(a.Value, "TIKARA", "POWER")
		result.Lang = "en"
		return result
	}
}
