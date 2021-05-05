package request

type RequestType interface {
	Validate() error
}
