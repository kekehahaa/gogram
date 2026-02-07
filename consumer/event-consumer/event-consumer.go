package event_consumer

import (
	"time"

	"github.com/kekehahaa/gogram/events"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			// log

			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		if err := c.handleEvents(gotEvents); err != nil {
			// log

			continue
		}
	}
}

// 1. потеря событий ретраи, фоллбэк в опертивке, подтверждение для фетчера
// 2. обработка всей пачки: счетчик ошибок и остановка
// 3. параллельная обработка: понадобится WaitGroup
func (c *Consumer) handleEvents(events []events.Event) error {
	for _, event := range events {
		// log

		if err := c.processor.Process(event); err != nil {
			// log

			continue
		}
	}

	return nil
}
