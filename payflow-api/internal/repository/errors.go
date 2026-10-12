package repository

import "errors"

var ErrMerchantNotFound = errors.New("merchant not found")
var ErrMerchantUserNotFound = errors.New("merchant user not found")
var ErrInvalidCredentials = errors.New("invalid email or password")