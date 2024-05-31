package structs

type Imgflip struct {
	Success bool `json:"success"`
	Data    Data `json:"data"`
}

type Data struct {
	URL     string `json:"url"`
	PageURL string `json:"page_url"`
}
