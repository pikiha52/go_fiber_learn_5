package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"service_auth_go/api/presenter"
)

type Repository interface {
	SigninRepository(request *presenter.AuthRequest) (*presenter.ExternalRequest, error)
}

type repository struct {
	BaseURL string
}

func NewRepo(base_url string) Repository {
	return &repository{
		BaseURL: base_url,
	}
}

func (r *repository) SigninRepository(request *presenter.AuthRequest) (*presenter.ExternalRequest, error) {
	queryParams := make(map[string]string)
	queryParams["username"] = request.Username

	u, err := url.Parse(r.BaseURL + "/user")
	if err != nil {
		return nil, err
	}

	q := u.Query()
	for key, value := range queryParams {
		q.Add(key, value)
	}

	u.RawQuery = q.Encode()
	fullUrl := u.String()

	resp, err := http.Get(fullUrl)
	if err != nil {
		return nil, errors.New("Failed connecting to service Users!")
	}
	defer resp.Body.Close()
	var externalRequest presenter.ExternalRequest

	err = json.NewDecoder(resp.Body).Decode(&externalRequest)
	if err != nil {
		return nil, err
	}

	return &externalRequest, nil
}
