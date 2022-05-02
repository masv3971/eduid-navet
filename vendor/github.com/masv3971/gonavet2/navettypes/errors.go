package navettypes

var (
	ErrSkatteverketInvalidCall             = &Errors{Message: "Felaktigt anrop"}
	ErrSkatteverketNotLogedIn              = &Errors{Message: "Ej inloggad"}
	ErrSkatteverketAuthorizationNotAllowed = &Errors{Message: "Åtkomst nekad"}
	ErrSkatteverketNoValidContentType      = &Errors{Message: "Felaktigt anrop"}
	ErrTechnicalFail                       = &Errors{Message: "Tekniskt fel"}
	ErrNotSpecified                        = &Errors{Message: "Not specified"}
	ErrTooManyRequests                     = &Errors{Message: "Too many requests"}
	ErrTimeout                             = &Errors{Message: "Timeout"}
)

// Errors is the bespoke error struct
type Errors struct {
	SkvClientCorrelationID string `json:"skv_client_correlation_id"`
	Message                string `json:"message"`
}

func (e *Errors) addID(id string) {
	e.SkvClientCorrelationID = id
}

func (e *Errors) Error() string {
	//	if e.Ladok != nil && len(e.Internal) > 0 {
	//		return fmt.Sprintf("internal error: %v, ladok error: %v", e.Internal, e.Ladok)
	//	} else if len(e.Internal) > 0 {
	//		return fmt.Sprintf("internal error: %v", e.Internal)
	//	} else if e.Ladok != nil {
	//		return fmt.Sprintf("ladok error: %v", e.Ladok)
	//	}
	return ""
}

// Error interface
type Error interface {
	Error() string
}

func oneError(m, t, f, e string) *Errors {
	return nil
}
