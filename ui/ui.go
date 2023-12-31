package ui

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	g "github.com/hamao0820/sortvis/graph"
	"github.com/hamao0820/sortvis/sort/iter"
	"github.com/hamao0820/sortvis/util"
	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/barchart"
)

const (
	max = 100
)

func Run(num int, duration int, algorithm Algorithm, file, graph string, interact bool) error {
	t, err := tcell.New()
	if err != nil {
		return err
	}
	defer t.Close()

	var values []int
	if file != "" {
		values, err = util.ReadFileAndConvertToIntSlice(file)
		if err != nil {
			return err
		}
	} else if graph != "" {
		flag := false
		for _, v := range g.GraphList {
			if v.Name == graph {
				values = v.Function()
				flag = true
				break
			}
		}
		if !flag {
			return fmt.Errorf("graph type %s is not supported", graph)
		}
	} else {
		for i := 0; i < num; i++ {
			values = append(values, rand.Intn(max+1))
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bc, err := barchart.New(
		barchart.ShowValues(),
		barchart.BarWidth(1+50/len(values)), // len(values) > 50 ? 1 : 2
	)
	if err != nil {
		return err
	}

	sorter, err := algorithm.iterator()
	if err != nil {
		return err
	}
	sorter.Init(values)

	if interact {
		go playBarChartByKey(ctx, bc, values, sorter)
	} else {
		go playBarChartByTick(ctx, bc, values, time.Duration(duration)*time.Millisecond, sorter)
	}

	title := algorithm.Pretty()
	if interact {
		title += " / press space to next step"
	}
	c, err := container.New(
		t,
		container.Border(linestyle.Light),
		container.BorderTitle(title),
		container.PlaceWidget(bc),
	)
	if err != nil {
		return err
	}

	keyboardHandler := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' || k.Key == keyboard.KeyCtrlC {
			cancel()
		}
		if interact {
			if k.Key == keyboard.KeySpace || k.Key == 'N' {
				sorter.Next()
			}
		}
	}
	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(keyboardHandler)); err != nil {
		return err
	}

	return nil
}

func playBarChartByKey(ctx context.Context, bc *barchart.BarChart, values []int, sorter *iter.SortIterator) {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := bc.Values(values, max); err != nil {
				panic(err)
			}

		case <-ctx.Done():
			return
		}
	}
}

func playBarChartByTick(ctx context.Context, bc *barchart.BarChart, values []int, delay time.Duration, sorter *iter.SortIterator) {
	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := bc.Values(values, max); err != nil {
				panic(err)
			}

			sorter.Next()

		case <-ctx.Done():
			return
		}
	}
}
