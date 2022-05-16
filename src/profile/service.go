package profile

import (
	"context"
	"errors"
)

type Service interface {
	CreateProfile(ctx context.Context, p Profile) (Profile, error)
	GetProfile(ctx context.Context, id string) (Profile, error)
	EditProfile(ctx context.Context, id string, p Profile) (Profile, error)
	SearchProfile(ctx context.Context, name string) ([]Profile, error)
	ListProfiles(ctx context.Context) ([]Profile, error)
	DeleteProfile(ctx context.Context, id string) error
	ActivateProfile(ctx context.Context, id string) error
	DeactivateProfile(ctx context.Context, id string) error
	IsActive(ctx context.Context, id string) (bool, error)
}

type Profile struct {
	ID       string
	Name     string
	Avatar   string
	IsActive bool
}

var (
	ErrProfileNotFound = errors.New("profile not found")
	ErrNotImplemented  = errors.New("not implemented")
)

type profileService struct {
	// Add DB for persistent later
}

func NewProfileService() Service {
	return &profileService{}
}

// Service implementation
func (s *profileService) CreateProfile(ctx context.Context, p Profile) (Profile, error) {
	p.ID = "TODO"
	return p, nil
}

func (s *profileService) GetProfile(ctx context.Context, id string) (Profile, error) {
	// TODO
	return Profile{}, ErrNotImplemented
}

func (s *profileService) EditProfile(ctx context.Context, id string, p Profile) (Profile, error) {
	// TODO
	return p, ErrNotImplemented
}

func (s *profileService) SearchProfile(ctx context.Context, name string) ([]Profile, error) {
	// TODO
	return []Profile{}, ErrNotImplemented
}

func (s *profileService) ListProfiles(ctx context.Context) ([]Profile, error) {
	// TODO
	return []Profile{}, ErrNotImplemented
}

func (s *profileService) DeleteProfile(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *profileService) DeleteProfile(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *profileService) ActivateProfile(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *profileService) DeactivateProfile(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *profileService) IsActive(ctx context.Context, id string) (bool, error) {
	return false, ErrNotImplemented
}
