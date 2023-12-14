package user

import (
	"context"
)

func (r *repository) Delete(ctx context.Context, ID int64) error {
	_, err := r.db.ExecContext(
		ctx,
		`DELETE FROM user WHERE id = ?`,
		ID,
	)
	return err
}
