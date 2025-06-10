package health

type Handler struct{}

func NewHealthHandler() *Handler {
	return &Handler{}
}
