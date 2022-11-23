package process

import (
	"context"
	"fmt"
	"github.com/go-kit/log"
	"time"
)

type processReader struct {
	log log.Logger
}

func NewProcessReader(log log.Logger) *processReader {
	return &processReader{log: log}
}

func (p processReader) ProcessReader(ctx context.Context) {
	fmt.Println("Robot en linea", time.Now())
}
