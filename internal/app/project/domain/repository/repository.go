package repository

import "context"

type Repository interface {
	Find(ctx context.Context) error
}
