package global

type MccGlobalSettings struct {
	
	// flags 
	loggingCallback,
	warningCallback,
	errorCallback,
}

func LoggingCallbackError(logcallback) {
	MccGlobalSettings.loggingCallback
}