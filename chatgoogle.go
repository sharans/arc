package arc

type ChatData struct {
	Tables []Table `json:"cards,omitempty"`
}

type Table struct {
	Rows   []TableRow `json:"rows,omitempty"`
	Header string     `json:"header,omitempty"`
}

type TableRow struct {
	Id          int           `json:"id,omitempty"`
	Title       string        `json:"title,omitempty"`
	Text        string        `json:"text,omitempty"`
	Description string        `json:"description,omitempty"`
	Url         string        `json:"url,omitempty"`
	Image       string        `json:"image,omitempty"`
	Action      OnClickAction `json:"action,omitempty"`
}

type ActionMethod string

const (
	NO_ACTION_METHOD string = "NO_ACTION_METHOD"
)

type OnClickAction struct {
	ActionMethodName string         `json:"actionMethodName"`
	Parameters       []KeyValuePair `json:"parameters"`
	Url              string         `json:"url"`
}

type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"string"`
}
