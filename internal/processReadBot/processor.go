package processReadBot

import "context"

type IProcessReaderBot interface {
	ProcessReaderInitializer()
}

type ReaderProccesor interface {
	ProcessReader(ctx context.Context)
}
