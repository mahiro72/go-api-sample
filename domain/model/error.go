package model

import "errors"

// user
var ErrUserFieldsValidation = errors.New("failure user fields validation")
var ErrUserIDValidation = errors.New("failure userID validation")

// post
var ErrPostFieldsValidation = errors.New("failure post fields validation")
var ErrPostIDValidation = errors.New("failure postID validation")
