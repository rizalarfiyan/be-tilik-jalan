package constant

type AuthRole string

const (
	AuthAdmin AuthRole = "admin"
	AuthGuest AuthRole = "guest"
)

func (r AuthRole) String() string {
	return string(r)
}

var MapAuthRole = map[AuthRole]bool{
	AuthAdmin: true,
	AuthGuest: true,
}

func (r AuthRole) IsValid() bool {
	_, ok := MapAuthRole[r]
	return ok
}

func (r AuthRole) IsAdmin() bool {
	return r == AuthAdmin
}

func (r AuthRole) IsValidAuth(roles ...AuthRole) bool {
	if len(roles) == 0 || r.IsAdmin() {
		return true
	}

	for _, role := range roles {
		if role == r {
			return true
		}
	}

	return false
}
