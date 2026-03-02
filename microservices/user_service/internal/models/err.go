package models

import "errors"

// create user error
var ErrCreateUser = errors.New("create user failed")
var ErrUserIsExist = errors.New("user is exist")
var ErrEmptyValue = errors.New("empty user value")

// get sub error
var ErrGetSub = errors.New("get sub failed")
var ErrInvalidId = errors.New("invalid id")
var ErrUserNotFound = errors.New("user not found")
