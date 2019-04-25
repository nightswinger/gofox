package sigfox

type ApiUserService service

type ApiUser struct {
	Name     string `json:"name,omitempty"`
	Timezone string `json:"timezone,omitempty"`
	//Group    MinimalGroup
	CreationTime int64  `json:"creationTime,omitempty"`
	ID           string `json:"id,omitempty"`
	AccessToken  string `json:"accessToken,omitempty"`
	//Profiles []Profile
}
