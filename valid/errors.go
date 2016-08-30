package valid

type Error struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

var DefaultErrors = map[string]Error{
	"required": Error{100, "should be present"},
	"email":    Error{101, "invalid email address"},
	"unique":   Error{102, "should be unique"},
}
