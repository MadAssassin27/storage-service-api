package request

type ItemCreateRequest struct {
	Name string `json:"name"`
	Item string `json:"item"`
	// Owner string `json:"owner"`
}
