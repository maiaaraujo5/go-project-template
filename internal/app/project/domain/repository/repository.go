package repository

import "context"

type Finder interface {
	Find(ctx context.Context) error
}
