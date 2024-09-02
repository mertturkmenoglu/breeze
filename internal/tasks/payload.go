package tasks

type Payload interface {
	CheckStatusPayload
}

type (
	CheckStatusPayload struct {
		ID  string
		URL string
	}
)
