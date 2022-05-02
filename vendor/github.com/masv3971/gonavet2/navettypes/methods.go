package navettypes

func (s SimpleAPIError) error() error {
	switch s.Message {
	case "Too many requests":
		return ErrTooManyRequests
	case "Timeout":
		return ErrTimeout
	}
	return nil
}

func (f Felmeddelande) error() error {
	switch f.Orsakskodsbeskrivning {
	case "Felaktigt anrop: indata ej korrekt ifyllt":
		return ErrSkatteverketInvalidCall
	case "Ej inloggad":
		return ErrSkatteverketNotLogedIn
	case "Åtkomst nekad":
		return ErrSkatteverketAuthorizationNotAllowed
	case "Felaktigt anrop: saknar stöd för Content-Type 'application/xml'":
		return ErrSkatteverketNoValidContentType
	case "Tekniskt fel":
		return ErrTechnicalFail
	default:
		return ErrNotSpecified
	}

}
