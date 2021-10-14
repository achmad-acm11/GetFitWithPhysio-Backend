package exception

type failedLogin struct {
	Error string
}

func NewFailedLogin(error string) failedLogin {
	return failedLogin{
		Error: error,
	}
}
