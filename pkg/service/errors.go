package service

import "errors"

var ErrNoPermission = errors.New("no permission")
var ErrNotFound = errors.New("not found")
var ErrUnauthenticated = errors.New("unauthenticated")
var ErrRateLimiting = errors.New("rate limiting")
var ErrNamespaceQuotaExceed = errors.New("namespace quota exceed")
