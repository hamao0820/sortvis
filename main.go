package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/hamao0820/sortvis/sort"
	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
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

func main() {
	t, err := tcell.New()
	if err != nil {
		panic(err)
	}
	defer t.Close()

	var values []int
	for i := 0; i < 99; i++ {
		values = append(values, rand.Intn(max+1))
	}

	ctx, cancel := context.WithCancel(context.Background())
	bc, err := barchart.New(
		barchart.ShowValues(),
		barchart.BarColors([]cell.Color{
			cell.ColorRed,
		}),
		barchart.ValueColors([]cell.Color{
			cell.ColorNavy,
		}),
		barchart.BarWidth(1+50/len(values)), // len(values) > 50 ? 1 : 2
	)
	if err != nil {
		panic(err)
	}

	sortChan := make(chan int, 1)
	defer close(sortChan)
	go sort.BubbleSortAsync(values, sortChan)
	// go playBarChartByTick(ctx, bc, values, 300*time.Millisecond, sortChan)
	go playBarChartByKey(ctx, bc, values, 10*time.Millisecond)

	c, err := container.New(
		t,
		container.Border(linestyle.Light),
		container.BorderTitle("PRESS Q/Ctrl+C TO QUIT"),
		container.PlaceWidget(bc),
	)
	if err != nil {
		panic(err)
	}

	keyboardHandler := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' || k.Key == keyboard.KeyCtrlC {
			cancel()
		}
		if k.Key == keyboard.KeySpace || k.Key == 'N' {
			if !sort.ValidSorted(values) {
				sortChan <- 1
			}
		}
	}
	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(keyboardHandler)); err != nil {
		panic(err)
	}
}

func playBarChartByKey(ctx context.Context, bc *barchart.BarChart, values []int, delay time.Duration) {
	ticker := time.NewTicker(delay)
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

func playBarChartByTick(ctx context.Context, bc *barchart.BarChart, values []int, delay time.Duration, channel chan int) {
	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if !sort.ValidSorted(values) {
				channel <- 1
			}
			if err := bc.Values(values, max); err != nil {
				panic(err)
			}

		case <-ctx.Done():
			return
		}
	}
}
