package models

/*JsonWebToken type struct for response*/
type JsonWebToken struct {
	AccessToken string `json: "accessToken, omitempty"`
	ExpiresIn   string `json: "expiresIn, omitempty"`
}
