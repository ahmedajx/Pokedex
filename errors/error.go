package errors

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"error_code"`
}
