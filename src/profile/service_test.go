package profile

import (
	"context"
	"testing"
)

func TestCreateProfile(t *testing.T) {
	s := NewProfileService()
	p := Profile{}
	saved, err := s.CreateProfile(context.TODO(), p)

	if err != nil {
		t.Error("CreateProfile return error")
	}

	if saved.ID != "TODO" {
		t.Errorf("CreateProfile is expected to return %v, got %v", "TODO", saved.ID)
	}
}
