package poller

import (
	"context"

	"github.com/eycia/rudita/basic"
)

type Poller func(ctx context.Context) <-chan basic.RawLog
