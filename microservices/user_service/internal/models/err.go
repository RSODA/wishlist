package models

import "errors"

var ErrCreateUser = errors.New("create user failed")
var ErrUserIsExist = errors.New("user is exist")
var ErrEmptyValue = errors.New("empty user value")
