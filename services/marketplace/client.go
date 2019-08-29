// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package marketplace

import (
	"net/http"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/services/httpservice"
	"github.com/pkg/errors"
)

type MarketplaceClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string, httpService httpservice.HTTPService) *MarketplaceClient {
	return &MarketplaceClient{
		baseURL:    baseURL,
		httpClient: httpService.MakeClient(true),
	}
}

func (m *MarketplaceClient) GetPlugins() ([]*model.BaseMarketplacePlugin, error) {
	res, err := m.httpClient.Get(m.baseURL + "/api/v1/plugins")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get plugins from marketplace")
	}

	plugins, err := model.BasePluginsFromReader(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal plugins marketplace")
	}

	return plugins, nil
}
