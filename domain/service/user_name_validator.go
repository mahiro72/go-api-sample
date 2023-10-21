package service

import "context"

type UserNameValidator interface {
	CheckAlreadyUsed(ctx context.Context, name string) (bool, error)
}
