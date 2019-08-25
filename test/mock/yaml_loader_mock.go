package mock

import (
    "github.com/HiroyukiNakatsuma/canon-go/internal/action"
    "github.com/HiroyukiNakatsuma/canon-go/internal/config"
)

type YamlLoaderMock struct {
    actions []action.Action
    config  *config.Config
}

func NewYamlLoaderMock(config *config.Config, actions ...action.Action) *YamlLoaderMock {
    return &YamlLoaderMock{config: config, actions: actions}
}

func (yamlLoader *YamlLoaderMock) LoadConfig() *config.Config {
    return &config.Config{Threads: 1, Loop: 1}
}

func (yamlLoader *YamlLoaderMock) LoadActions() []action.Action {
    return yamlLoader.actions
}
