package serviceapi

type UserError struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}
