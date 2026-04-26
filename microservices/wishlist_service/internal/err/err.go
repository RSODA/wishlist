package errors_entity

import "errors"

var ErrCreateWish = errors.New("create new wish error")
var ErrGetWishs = errors.New("get wishs error")

var ErrValidateWish = errors.New("invalid param")

var ErrUserNotFound = errors.New("user not found")

var ErrGetWishById = errors.New("get wish by id error")

var ErrInvalidId = errors.New("invalid id")

var ErrSubNotConfirmed = errors.New("sub not confirmed")
