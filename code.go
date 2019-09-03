package ogs

const (
	StatusOK = 0

	// authorization
	StatusUnauthorized  = 10001
	StatusTokenExpired  = 10002
	StatusInvalidToken  = 10003
	StatusUserNotFound  = 10004
	StatusErrorPassword = 10005

	// system and resources
	StatusSystemError     = 20001
	StatusResourceExpired = 20002

	// crud
	StatusCreateFailed       = 30001
	StatusUpdateFailed       = 30002
	StatusDestroyFailed      = 30003
	StatusBatchCreateFailed  = 30004
	StatusBatchUpdateFailed  = 30005
	StatusBatchDestroyFailed = 30006
)

var codeText = map[int]string{
	StatusOK: "OK",

	StatusUnauthorized:  "Unauthorized",
	StatusTokenExpired:  "Token Expired",
	StatusInvalidToken:  "Invalid Token",
	StatusUserNotFound:  "User Not Found",
	StatusErrorPassword: "Error Password",

	StatusSystemError:     "System Error",
	StatusResourceExpired: "Resource Expired",

	StatusCreateFailed:       "Create Failed",
	StatusUpdateFailed:       "Update Failed",
	StatusDestroyFailed:      "Destroy Failed",
	StatusBatchCreateFailed:  "Batch Create Failed",
	StatusBatchUpdateFailed:  "Batch Update Failed",
	StatusBatchDestroyFailed: "Batch Destroy Failed",
}

// StatusText returns a text for code. It returns the empty
// string if the code is unknown.
func CodeText(code int) string {
	return codeText[code]
}
