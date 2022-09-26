package database

import "github.com/fredrikved/buk-gg/backend/common"

func (s *Store) Settings() *Collection[common.Settings] {
	return &Collection[common.Settings]{
		collection: "settings",
		client:     s.client,
	}
}

func (s *Store) Config() *Collection[common.Config] {
	return &Collection[common.Config]{
		collection: "config",
		client:     s.client,
	}
}
