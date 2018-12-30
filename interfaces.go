package pantopoda

import "github.com/Kamva/nautilus"

// RequestData is an interface for incoming request payload
type RequestData interface {
	nautilus.Taggable
}
