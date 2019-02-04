package response

import "github.com/mikeflynn/go-alexa/customskill/request"

func New() *Response {
	return &Response{
		ShouldEndSession: Bool(true),
	}
}

func (r *Response) SetOutputSpeech(text string) *Response {
	r.OutputSpeech = &OutputSpeech{
		Type: "PlainText",
		Text: text,
	}

	return r
}

func (r *Response) SetOutputSpeechSSML(text string) *Response {
	r.OutputSpeech = &OutputSpeech{
		Type: "SSML",
		SSML: text,
	}

	return r
}

func (r *Response) SetReprompt(text string) *Response {
	r.Reprompt = &Reprompt{
		OutputSpeech: &OutputSpeech{
			Type: "PlainText",
			Text: text,
		},
	}

	return r
}

func (r *Response) SetRepromptSSML(text string) *Response {
	r.Reprompt = &Reprompt{
		OutputSpeech: &OutputSpeech{
			Type: "SSML",
			SSML: text,
		},
	}

	return r
}

func (r *Response) SetCard(title, content string) *Response {
	return r.SetSimpleCard(title, content)
}

func (r *Response) SetSimpleCard(title, content string) *Response {
	r.Card = &Card{
		Type:    "Simple",
		Title:   title,
		Content: content,
	}

	return r
}

func (r *Response) SetStandardCard(title, content, smallImg, largeImg string) *Response {
	r.Card = &Card{
		Type:    "Standard",
		Title:   title,
		Content: content,
	}

	if (smallImg != "") || (largeImg != "") {
		r.Card.Image = &Image{
			SmallImageURL: smallImg,
			LargeImageURL: largeImg,
		}
	}

	return r
}

func (r *Response) SetLinkAccountCard() *Response {
	r.Card = &Card{
		Type: "LinkAccount",
	}

	return r
}

func (r *Response) SetAskForPermissionsConsentCard(permissions []string) *Response {
	r.Card = &Card{
		Type:        "AskForPermissionsConsent",
		Permissions: permissions,
	}

	return r
}

func (r *Response) SetEndSession(flag *bool) *Response {
	r.ShouldEndSession = flag

	return r
}

func (r *Response) SetCanFulfillIntent(flag ability) *Response {
	r.CanFulfillIntent = &CanFulfillIntent{
		CanFulfill: string(flag),
	}
	return r
}

func (r *Response) SetCanFulfillSlot(slot request.Slot, canUnderstand ability, canFulfill ability) {
	if r.CanFulfillIntent.Slots == nil {
		r.CanFulfillIntent.Slots = make(map[string]Slot)
	}
	r.CanFulfillIntent.Slots[slot.Name] = Slot{
		CanFulfill:    string(canFulfill),
		CanUnderstand: string(canUnderstand),
	}
}

func (r *Response) String() string {
	b, err := jsonMarshal(r)
	if err != nil {
		return "failed to marshal response to JSON: " + err.Error()
	}

	return string(b)
}
