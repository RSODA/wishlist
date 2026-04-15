package models

import "errors"

// create user error
var ErrCreateUser = errors.New("create user failed")
var ErrUserIsExist = errors.New("user is exist")
var ErrEmptyValue = errors.New("empty user value")

var ErrGetUser = errors.New("get user failed")

// get sub error
var ErrGetSub = errors.New("get sub failed")
var ErrInvalidId = errors.New("invalid id")
var ErrUserNotFound = errors.New("user not found")
var ErrUserToSubscription = errors.New("user to subscription failed")
var ErrUserAlreadySubscribed = errors.New("user to subscription failed")

// subscribe
var ErrJsonMarshal = errors.New("json marshal failed")
var ErrSubscribe = errors.New("subscribe failed")
var ErrInvalidArgument = errors.New("invalid argument(-s)")

var ErrAcceptedSub = errors.New("accepted sub failed")
