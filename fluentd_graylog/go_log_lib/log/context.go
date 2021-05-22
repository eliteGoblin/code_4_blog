package log

import "context"

type ctxKey int

const logKey ctxKey = 0

// NewContext returns a new context that carries the logger value.
func NewContext(ctx context.Context, l Log) context.Context {
	return context.WithValue(ctx, logKey, l)
}

// FromContext returns the Logger stored in the context, if any.
func FromContext(ctx context.Context) (Log, bool) {
	if ctx == nil {
		return nil, false
	}

	lg, ok := ctx.Value(logKey).(Log)
	return lg, ok
}

// With returns a Log instance for the given context. If not available
// in the context then returns the default log.
func With(ctx context.Context) Log {
	if ctx != nil {
		lg, has := ctx.Value(logKey).(Log)
		if has {
			return lg
		}
	}

	mu.RLock()
	defer mu.RUnlock()

	return defaultLog
}

// deprecated - use With(ctx)
// From returns a Log instance for the given context. If not available
// in the context then returns the default log.
func From(ctx context.Context) Log {
	if ctx != nil {
		lg, has := ctx.Value(logKey).(Log)
		if has {
			return lg
		}
	}

	mu.RLock()
	defer mu.RUnlock()

	return defaultLog
}
