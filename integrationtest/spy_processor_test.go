package integrationtest

import (
	"context"

	"github.com/birdayz/kstreams"
)

type SpyProcessor struct {
	out chan [2]string
}

func (p *SpyProcessor) Init(processorContext kstreams.ProcessorContext[string, string]) error {
	return nil
}

func (p *SpyProcessor) Close() error {
	return nil
}

func (p *SpyProcessor) Process(ctx context.Context, k string, v string) error {
	p.out <- [...]string{k, v}
	return nil
}
