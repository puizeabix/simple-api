package profile

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateProfileEndpoint     endpoint.Endpoint
	GetProfileEndpoint        endpoint.Endpoint
	EditProfileEndpoint       endpoint.Endpoint
	SearchProfileEndpoint     endpoint.Endpoint
	ListProfilesEndpoint      endpoint.Endpoint
	DeleteProfileEndpoint     endpoint.Endpoint
	ActivateProfileEndpoint   endpoint.Endpoint
	DeactivateProfileEndpoint endpoint.Endpoint
	IsActiveEndpoint          endpoint.Endpoint
}

func (e Endpoints) CreateProfile(ctx context.Context, p Profile) (Profile, error) {
	request := createProfileRequest{Profile: p}
	response, err := e.CreateProfileEndpoint(ctx, request)
	if err != nil {
		return p, err
	}

	resp := response.(createProfileResponse)

	return resp.Profile, nil

}

type createProfileRequest struct {
	Profile Profile
}

type createProfileResponse struct {
	Profile Profile `json:"profile,omitempty"`
	Err     error   `json:"err,omitempty"`
}
