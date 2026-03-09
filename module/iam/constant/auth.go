package constant

import "time"

const (
	AccessTokenExpiresIn  time.Duration = time.Hour * 1       // 1h
	RefreshTokenExpiresIn               = time.Hour * 25 * 30 // 30d
)
