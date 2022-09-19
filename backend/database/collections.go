package database

import "github.com/fredrikved/buk-gg/common"

func (s *Store) Settings() *Collection[common.Settings] {
	return &Collection[common.Settings]{
		collection: "settings",
		client:     s.client,
	}
}
