package resp

const RES = "res"

const (
	NeedToLogin  = 444
	LoginExpired = 445
)

type R struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

func Success(data any) R {
	return R{
		Code: 200,
		Data: data,
		Msg:  "success",
	}
}

func Fail(msg string) R {
	return R{
		Code: 500,
		Msg:  msg,
		Data: nil,
	}
}

func FailWithCode(code int, msg string) R {
	return R{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
