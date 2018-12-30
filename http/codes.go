package http

// StatusCode is an alias for int using as a http status code
type StatusCode int

// IsInformational check if status code is informational
func (s StatusCode) IsInformational() bool {
	return s >= 100 && s < 200
}

// IsSuccess check if status code is success
func (s StatusCode) IsSuccess() bool {
	return s >= 200 && s < 300
}

// IsRedirection check if status code is redirection
func (s StatusCode) IsRedirection() bool {
	return s >= 300 && s < 400
}

// IsClientError check if status code is client error
func (s StatusCode) IsClientError() bool {
	return s >= 400 && s < 500
}

// IsInternalError check if status code is internal error
func (s StatusCode) IsInternalError() bool {
	return s > 500
}

// Int cast the status code to int value
func (s StatusCode) Int() int {
	return int(s)
}

// Informational HTTP Status Codes
const (
	Continue StatusCode = iota + 100
	SwitchingProtocols
	Processing
)

// Success HTTP Status Codes
const (
	OK StatusCode = iota + 200
	Created
	Accepted
	NonAuthoritativeInformation
	NoContent
	ResetContent
	PartialContent
	MultiStatus
	AlreadyReported
	IMUsed StatusCode = 226
)

// Redirection HTTP Status Code
const (
	MultipleChoices StatusCode = iota + 300
	MovedPermanently
	Found
	SeeOther
	NotModified
	UseProxy
	_
	TemporaryRedirect
	PermanentRedirect
)

// Client Error HTTP Status Code
const (
	BadRequest StatusCode = iota + 400
	Unauthorized
	PaymentRequired
	Forbidden
	NotFound
	MethodNotAllowed
	NotAcceptable
	ProxyAuthenticationRequired
	RequestTimeout
	Conflict
	Gone
	LengthRequired
	PreconditionFailed
	PayloadTooLarge
	RequestURITooLong
	UnsupportedMediaType
	RequestedRangeNotSatisfiable
	ExpectationFailed
	ImATeapot
	_
	_
	MisdirectedRequest
	UnprocessableEntity
	Locked
	FailedDependency
	_
	UpgradeRequired
	_
	PreconditionRequired
	TooManyRequests
	_
	RequestHeaderFieldsTooLarge
	ConnectionClosedWithoutResponse StatusCode = 444
	UnavailableForLegalReasons      StatusCode = 451
	ClientClosedRequest             StatusCode = 499
)

// Server Error HTTP Status Code
const (
	InternalServerError StatusCode = iota + 500
	NotImplemented
	BadGateway
	ServiceUnavailable
	GatewayTimeout
	VersionNotSupported
	VariantAlsoNegotiates
	InsufficientStorage
	LoopDetected
	_
	NotExtended
	NetworkAuthenticationRequired
	NetworkConnectTimeoutError StatusCode = 599
)
