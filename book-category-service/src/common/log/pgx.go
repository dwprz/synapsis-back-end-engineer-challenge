package log

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
)

type QueryTracer struct{}

func (q *QueryTracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	argsStr := ""
	if len(data.Args) > 0 {
		args := make([]string, len(data.Args))
		for i, arg := range data.Args {
			args[i] = fmt.Sprintf("%v", arg)
		}
		argsStr = strings.Join(args, ", ")
	}

	message := fmt.Sprintf("Executing command SQL: %s, with args: [%s]",
		strings.ReplaceAll(strings.ReplaceAll(data.SQL, "\n", ""), "\t", ""), argsStr)

	Logger.Info(message)

	return ctx
}

func (q *QueryTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
}
