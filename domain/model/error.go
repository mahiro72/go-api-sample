package model

import "errors"

// user
var ErrUserFieldsValidation = errors.New("failue user fields validation")
var ErrUserIDValidation = errors.New("failue userID validation")

// post
var ErrPostFieldsValidation = errors.New("failue post fields validation")
var ErrPostIDValidation = errors.New("failue postID validation")
