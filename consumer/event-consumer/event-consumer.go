package event_consumer

import (
	"log/slog"
	"time"

	"github.com/kekehahaa/gogram/events"
	"github.com/kekehahaa/gogram/internal/lib/logger/sl"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
	log       slog.Logger
}

func New(log slog.Logger, fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
		log:       log,
	}
}

func (c Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			c.log.Error("can't fetch events", sl.Err(err))

			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		if err := c.handleEvents(gotEvents); err != nil {
			c.log.Error("can't handle events", sl.Err(err))

			continue
		}
	}
}

// 1. потеря событий ретраи, фоллбэк в опертивке, подтверждение для фетчера
// 2. обработка всей пачки: счетчик ошибок и остановка
// 3. параллельная обработка: понадобится WaitGroup
func (c *Consumer) handleEvents(events []events.Event) error {
	for _, event := range events {
		c.log.Info("got new event", slog.String("text", event.Text))

		if err := c.processor.Process(event); err != nil {
			c.log.Error("can't handle event", sl.Err(err))

			continue
		}
	}

	return nil
}
