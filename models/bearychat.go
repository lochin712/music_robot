package models

type BearyChatRequest struct {
	Token       string `json:"token"`
	Ts          int64  `json:"ts"`
	Text        string `json:"text"`
	TriggerWord string `json:"trigger_word"`
	Subdomain   string `json:"subdomain"`
	ChannelName string `json:"channel_name"`
	UserName    string `json:"user_name"`
}

type BearyChatResponse struct {
	Text        string        `json:"text"`
	Attachments []Attachments `json:"attachments"`
}

type Attachments struct {
	Title  string     `json:"title"`
	Text   string     `json:"text"`
	Color  string     `json:"color"`
	Images []ImageUrl `json:"images"`
}
type ImageUrl struct {
	URL string `json:"url"`
}
