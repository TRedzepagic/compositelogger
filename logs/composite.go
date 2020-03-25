package logs

// CompositeLog structure was needed to implement the "true for debug, false for no debug" functionality.
type CompositeLog struct {
	slicelog []SuperLogger
	flag     bool
}

// NewCustomLogger adds all passed loggers into a slice of SuperLoggers (Variadic args...SuperLogger)
// We see here that we can add more loggers (we just need to pass them through)
func NewCustomLogger(flag bool, args ...SuperLogger) CompositeLog {
	var compositelogger CompositeLog
	compositelogger.slicelog = args
	compositelogger.flag = flag
	return compositelogger
}

// The following functions change the prefix, then call each of the loggers' Print/f functions respectively.

// Info ..
func (composite CompositeLog) Info(v ...interface{}) {
	for _, logger := range composite.slicelog {
		logger.SetPrefix("Info:")
		logger.Println(v)
	}

}

// Infof ..
func (composite CompositeLog) Infof(format string, v ...interface{}) {

	for _, logger := range composite.slicelog {
		logger.SetPrefix("Infof:")
		logger.Printf(format, v)
	}

}

// Warn ..
func (composite CompositeLog) Warn(v ...interface{}) {

	for _, logger := range composite.slicelog {
		logger.SetPrefix("WARNING:")
		logger.Println(v)
	}

}

// Warnf ..
func (composite CompositeLog) Warnf(format string, v ...interface{}) {

	for _, logger := range composite.slicelog {
		logger.SetPrefix("WARNINGf:")
		logger.Printf(format, v)
	}

}

// Error ..
func (composite CompositeLog) Error(v ...interface{}) {

	for _, logger := range composite.slicelog {
		logger.SetPrefix("ERROR:")
		logger.Println(v)
	}

}

// Errorf ..
func (composite CompositeLog) Errorf(format string, v ...interface{}) {

	for _, logger := range composite.slicelog {
		logger.SetPrefix("ERRORf:")
		logger.Printf(format, v)
	}

}

// Debug (skipped if flag == false)
func (composite CompositeLog) Debug(v ...interface{}) {
	if composite.flag {
		for _, logger := range composite.slicelog {
			logger.SetPrefix("DEBUG:")
			logger.Println(v)
		}
	}

}

// Debugf (skipped if flag == false)
func (composite CompositeLog) Debugf(format string, v ...interface{}) {
	if composite.flag {
		for _, logger := range composite.slicelog {
			logger.SetPrefix("DEBUGf:")
			logger.Printf(format, v)
		}
	}

}
