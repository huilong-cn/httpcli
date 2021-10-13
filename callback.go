package httpcli

var timeoutCallback func(string, string)

func SetTimeoutNotify(fn func(string, string)) {
	timeoutCallback = fn
}
