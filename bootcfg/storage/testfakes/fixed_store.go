package testfakes

import (
	"fmt"

	"github.com/coreos/coreos-baremetal/bootcfg/storage/storagepb"
)

// FixedStore is used for testing purposes.
type FixedStore struct {
	Groups          map[string]*storagepb.Group
	Profiles        map[string]*storagepb.Profile
	IgnitionConfigs map[string]string
	CloudConfigs    map[string]string
}

// GroupGet returns the Group from the Groups map with the given id.
func (s *FixedStore) GroupGet(id string) (*storagepb.Group, error) {
	if group, present := s.Groups[id]; present {
		return group, nil
	}
	return nil, fmt.Errorf("Group not found")
}

// GroupList returns the groups in the Groups map.
func (s *FixedStore) GroupList() ([]*storagepb.Group, error) {
	groups := make([]*storagepb.Group, len(s.Groups))
	i := 0
	for _, g := range s.Groups {
		groups[i] = g
		i++
	}
	return groups, nil
}

// ProfilePut writes the given Profile to the Profiles map.
func (s *FixedStore) ProfilePut(profile *storagepb.Profile) error {
	s.Profiles[profile.Id] = profile
	return nil
}

// ProfileGet returns the Profile from the Profiles map with the given id.
func (s *FixedStore) ProfileGet(id string) (*storagepb.Profile, error) {
	if profile, present := s.Profiles[id]; present {
		return profile, nil
	}
	return nil, fmt.Errorf("Profile not found")
}

// ProfileList returns the profiles in the Profiles map.
func (s *FixedStore) ProfileList() ([]*storagepb.Profile, error) {
	profiles := make([]*storagepb.Profile, len(s.Profiles))
	i := 0
	for _, p := range s.Profiles {
		profiles[i] = p
		i++
	}
	return profiles, nil
}

// IgnitionGet returns the Ignition config with the given id.
func (s *FixedStore) IgnitionGet(id string) (string, error) {
	if config, present := s.IgnitionConfigs[id]; present {
		return config, nil
	}
	return "", fmt.Errorf("no Ignition Config %s", id)
}

// CloudGet returns the Cloud config with the given id.
func (s *FixedStore) CloudGet(id string) (string, error) {
	if config, present := s.CloudConfigs[id]; present {
		return config, nil
	}
	return "", fmt.Errorf("no Cloud Config %s", id)
}
