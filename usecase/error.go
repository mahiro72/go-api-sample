package usecase

import "errors"

var ErrFieldsValidation = errors.New("failure fields validation")
var ErrIDValidation = errors.New("failure id validation")
var ErrNotExistsData = errors.New("not exists data")

// user
var ErrUserNameAlreadyUsed = errors.New("user name already used")
