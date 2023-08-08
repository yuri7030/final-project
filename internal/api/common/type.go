package common

type AuthJWT struct {
	Email string
	Name  string
	Role  int
	ID    uint
	Exp   int64
}
