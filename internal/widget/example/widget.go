package example

type loggerPlugin struct {
}

func New() *loggerPlugin {
	return &loggerPlugin{}
}

func (l *loggerPlugin) Execute() (*string, error) {
	result := "Logger is inactive"
	return &result, nil
}
