package arc

type Context struct {
	Name       string            `json:"name,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
	Lifespan   int32             `json:"lifespan,omitempty"`
}

type ActionRequest struct {
	UserSays         string            `json:"usersays"`
	Parameters       map[string]string `json:"parameters"`
	Contexts         []Context         `json:"ContextList"`
	ActionName       string            `json:"actionName"`
	SessionID        string            `json:"sessionId"`
	Sender           Sender            `json:"sender"`
	UserResponseType string            `json:userResponseType`
	FeatureToggles   map[string]bool   `json:"featureToggles,omitempty"`
}

type AsyncMessageRequest struct {
	UserEmail     []string `json:"userEmail"`
	Text          string   `json:"text"`
	Data          ChatData `json:"data,omitempty"`
	NotifyEmail   []string `json:"notifyEmail,omitempty"`
	NotifyMessage string   `json:"notifyText,omitempty"`
}

type Sender struct {
	Name           string `json:"name,omitempty"`
	DisplayName    string `json:"displayName"`
	Email          string `json:"email,omitempty"`
	EmployeeId     string `json:"employeeId"`
	LoginName      string `json:"loginName"`
	HomeOffice     string `json:"homeOffice"`
	StaffingOffice string `json:"staffingOffice"`
	WorkingOffice  string `json:"workingOffice"`
	Grade          string `json:"grade"`
	Role           string `json:"role"`
	Assignment        EmpAssignment `json:"assignment"`
}

type HavingName struct {
	Name string `json:"name"`
}

type EmpAssignment struct {
	AccountName HavingName `json:"account"`
	Project EmpProject `json:"project"`
}

type EmpProject struct {
	Id string `json:"id"`
	Name string `json:"name"`

}

type ActionResponse interface {
	ToPlanetResponse() *PlanetResponse
}

type PlanetResponse struct {
	Type                 string            `json:"type"`
	Text                 string            `json:"text,omitempty"`
	Data                 ChatData          `json:"data,omitempty"`
	TemplateKey          string            `json:"templateKey,omitempty"`
	ContextList          []Context         `json:"contextList,omitempty"`
	TemplatePlaceHolders map[string]string `json:"templatePlaceHolders,omitempty"`
	ResetContext         bool              `json:"resetContext,omitempty"`
}

type SimpleResponse struct {
	Text         string
	ContextList  []Context
	TemplateKey  string
	ResetContext bool
}

func (r SimpleResponse) ToPlanetResponse() *PlanetResponse {
	return &PlanetResponse{
		Type:         "simpleResponse",
		Text:         r.Text,
		ContextList:  r.ContextList,
		TemplateKey:  r.TemplateKey,
		ResetContext: r.ResetContext,
	}
}

type ResponseFromTemplate struct {
	TemplateKey          string
	TemplatePlaceHolders map[string]string
	ContextList          []Context
	ResetContext         bool
}

func (r ResponseFromTemplate) ToPlanetResponse() *PlanetResponse {
	return &PlanetResponse{
		Type:                 "templatedResponse",
		ContextList:          r.ContextList,
		TemplatePlaceHolders: r.TemplatePlaceHolders,
		TemplateKey:          r.TemplateKey,
		ResetContext:         r.ResetContext,
	}
}

type FormattedResponse struct {
	Text         string
	Data         ChatData
	TemplateKey  string
	ContextList  []Context
	ResetContext bool
}

func (r FormattedResponse) ToPlanetResponse() *PlanetResponse {
	return &PlanetResponse{
		Type:         "formattedResponse",
		Text:         r.Text,
		Data:         r.Data,
		ContextList:  r.ContextList,
		TemplateKey:  r.TemplateKey,
		ResetContext: r.ResetContext,
	}
}

func (pr *ActionRequest) init() {
	if pr.Parameters == nil {
		pr.Parameters = map[string]string{}
	}
}
