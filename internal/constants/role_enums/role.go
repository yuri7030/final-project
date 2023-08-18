package role_enums

type UserRole int

const (
    Admin UserRole = iota + 1
    User
)