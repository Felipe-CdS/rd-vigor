package services

type ServiceLayerErr struct {
	Error   error
	Message string
	Code    int
}

type BadRequestErr error
