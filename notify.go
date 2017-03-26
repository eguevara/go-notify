package notify

import "fmt"

// Notifier defines the behavior required by types that want to implement Notify.
type Notifier interface {
	Notify(interface{}) error
}

// Message stores notify data to be sent as a message.
type Message struct {

	// Title is the title (or name) of the thing that was checked.
	Title string `json:"title,omitempty"`

	// Text is the text body that will carry notify information
	Text string `json:"message,omitempty"`

	// Endpoint is the URL/address/path/identifier/locator
	Endpoint string `json:"endpoint,omitempty"`
}

// String returns a string representation of Result type.
func (r Message) String() string {
	msg := ""
	msg += fmt.Sprintf("Title: %s\n", r.Title)
	msg += fmt.Sprintf("Endpoint: %s\n", r.Endpoint)
	msg += fmt.Sprintf("Text: %s", r.Text)
	return msg
}
