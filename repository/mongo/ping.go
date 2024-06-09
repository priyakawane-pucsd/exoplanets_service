package mongo

import (
	"context"
	"exoplanetservice/utils"
)

func (r *Repository) Ping(ctx context.Context) error {
	err := r.conn.Ping(ctx, nil)
	if err != nil {
		return utils.NewInternalServerError("Failed to ping MongoDB")
	}
	return nil
}
