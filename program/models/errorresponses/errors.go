package errorresponses

import "errors"

var UserNotFound = errors.New("User not found!")
var PasswordDontMatch = errors.New("Passwords don't match!")
