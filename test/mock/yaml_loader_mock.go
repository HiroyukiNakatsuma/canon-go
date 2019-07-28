package mock

import "github.com/HiroyukiNakatsuma/canon-go/internal"

type YamlLoaderMock struct {
    actions []internal.Action
}

func NewYamlLoaderMock(actions ...internal.Action) *YamlLoaderMock {
    return &YamlLoaderMock{actions: actions}
}

func (yamlLoader *YamlLoaderMock) LoadConfig() *internal.ActionConfig {
    return &internal.ActionConfig{Threads: 1, Loop: 1}
}

func (yamlLoader *YamlLoaderMock) LoadActions() []internal.Action {
    return yamlLoader.actions
}
