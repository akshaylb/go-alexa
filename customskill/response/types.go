package response

type ability string

const (
	Yes   ability = "YES"
	Maybe ability = "MAYBE"
	No    ability = "NO"
)

type envelope struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Response          *Response              `json:"response"`
}

type Response struct {
	OutputSpeech     *OutputSpeech     `json:"outputSpeech,omitempty"`
	Reprompt         *Reprompt         `json:"reprompt,omitempty"`
	Card             *Card             `json:"card,omitempty"`
	Directives       []interface{}     `json:"directives,omitempty"`
	ShouldEndSession *bool             `json:"shouldEndSession,omitempty"`
	CanFulfillIntent *CanFulfillIntent `json:"canFulfillIntent,omitempty"`
}

type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
	SSML string `json:"ssml,omitempty"`
}

type Reprompt struct {
	OutputSpeech *OutputSpeech `json:"outputSpeech,omitempty"`
}

type Card struct {
	Type        string   `json:"type"`
	Title       string   `json:"title,omitempty"`
	Content     string   `json:"content,omitempty"`
	Text        string   `json:"text,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	Image       *Image   `json:"image,omitempty"`
}

type Image struct {
	SmallImageURL string `json:"smallImageUrl,omitempty"`
	LargeImageURL string `json:"largeImageUrl,omitempty"`
}

type CanFulfillIntent struct {
	CanFulfill string          `json:"canFulfill"`
	Slots      map[string]Slot `json:"slots,omitempty"`
}

type Slot struct {
	CanUnderstand string `json:"canUnderstand"`
	CanFulfill    string `json:"canFulfill"`
}
