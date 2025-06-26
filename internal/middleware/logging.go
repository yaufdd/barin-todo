package middleware

import (
	"context"
	"log/slog"
	"time"

	"connectrpc.com/connect"
)

func LoggingInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			start := time.Now()

			slog.Info("gRPC request started",
				"procedure", req.Spec().Procedure,
				"peer", req.Peer().Addr,
			)

			resp, err := next(ctx, req)

			duration := time.Since(start)
			if err != nil {
				slog.Error("gRPC request failed",
					"procedure", req.Spec().Procedure,
					"peer", req.Peer().Addr,
					"error", err.Error(),
					"duration", duration,
				)
			} else {
				slog.Info("gRPC request completed",
					"procedure", req.Spec().Procedure,
					"peer", req.Peer().Addr,
					"duration", duration,
				)
			}

			return resp, err
		}
	}
}
