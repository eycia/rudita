package poller

import (
	"bufio"
	"context"
	"io"
	"strings"

	"github.com/eycia/rudita/basic"
	"github.com/sirupsen/logrus"
)

func ReaderStreamPoller(reader io.ReadCloser) func(ctx context.Context) <-chan basic.RawLog {
	return func(ctx context.Context) <-chan basic.RawLog {
		c := make(chan basic.RawLog, 1)

		go func() {
			defer close(c)
			defer reader.Close()

			buf := bufio.NewReaderSize(reader, 2*1024*1024) // 2M, TODO: move to config
			var err error
			var line string
			var linea string

			for err != io.EOF {
				line, err = buf.ReadString('\n')
				if err != nil && err != io.EOF {
					logrus.Errorln(err.Error()) // XXX: return a error?
					return
				}

				linea = strings.TrimSpace(line)
				if linea == "" {
					continue
				}
				c <- basic.RawLog(linea) // wait send

				select {
				case <-ctx.Done():
					break
				default:
				}
			}
		}()
		return c
	}
}
