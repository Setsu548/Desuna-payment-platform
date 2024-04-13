package token

import "time"

type IMaker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(toekn string) (*Payload, error)
}
