package site_blocker_config

import (
	"os"
	"slices"
	"time"

	"gopkg.in/yaml.v3"
)

type SiteBlockerConfig struct {
	Everyday struct {
		BlockStart int      `yaml:"start"`
		BlockEnd   int      `yaml:"end"`
		Sites      []string `yaml:"sites"`
	}
}

func NewSiteBlockerConfig(path string) *SiteBlockerConfig {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	newConfig := &SiteBlockerConfig{}
	err = yaml.Unmarshal(data, &newConfig)
	if err != nil {
		panic(err)
	}

	return newConfig
}

func (sbc *SiteBlockerConfig) IsBlocked(url string) bool {
	if !slices.Contains(sbc.Everyday.Sites, url) {
		return false
	}

	now := time.Now()

	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	startAt := startOfDay.Add(time.Minute * time.Duration(sbc.Everyday.BlockStart))

	startOfDay = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endAt := startOfDay.Add(time.Minute * time.Duration(sbc.Everyday.BlockEnd))

	isBlocked := now.After(startAt) && now.Before(endAt)

	return isBlocked
}
