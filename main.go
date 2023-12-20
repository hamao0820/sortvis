package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/barchart"
)

func main() {
	t, err := tcell.New()
	if err != nil {
		panic(err)
	}
	defer t.Close()

	ctx, cancel := context.WithCancel(context.Background())
	bc, err := barchart.New(
		barchart.BarColors([]cell.Color{
			cell.ColorBlue,
			cell.ColorRed,
			cell.ColorYellow,
			cell.ColorBlue,
			cell.ColorGreen,
			cell.ColorRed,
		}),
		barchart.ValueColors([]cell.Color{
			cell.ColorRed,
			cell.ColorYellow,
			cell.ColorNumber(33),
			cell.ColorGreen,
			cell.ColorRed,
			cell.ColorNumber(33),
		}),
		barchart.ShowValues(),
		barchart.BarWidth(8),
		barchart.Labels([]string{
			"CPU1",
			"",
			"CPU3",
		}),
	)
	if err != nil {
		panic(err)
	}
	go playBarChart(ctx, bc, 1*time.Second)

	c, err := container.New(
		t,
		container.Border(linestyle.Light),
		container.BorderTitle("PRESS Q TO QUIT"),
		container.PlaceWidget(bc),
	)
	if err != nil {
		panic(err)
	}

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}

	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter)); err != nil {
		panic(err)
	}
}

func playBarChart(ctx context.Context, bc *barchart.BarChart, delay time.Duration) {
	const max = 100

	ticker := time.NewTicker(delay)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			var values []int
			for i := 0; i < bc.ValueCapacity(); i++ {
				values = append(values, int(rand.Int31n(max+1)))
			}

			if err := bc.Values(values, max); err != nil {
				panic(err)
			}

		case <-ctx.Done():
			return
		}
	}
}
