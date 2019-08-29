// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

type MarketplacePluginState int

const (
	MarketPlacePluginStateNotInstalled MarketplacePluginState = iota
	MarketPlacePluginStateInstalled
)

// MarketplacePlugin provides a state-aware view of marketplace plugin.
type BaseMarketplacePlugin struct {
	HomepageURL  string
	DownloadURL  string
	SignatureURL string
	Manifest     *Manifest
}

type MarketplacePlugin struct {
	*BaseMarketplacePlugin
	State MarketplacePluginState
}

// type BaseMarketplacePlugins []*BaseMarketplacePlugin

// PluginFromReader decodes a json-encoded cluster from the given io.Reader.
func PluginFromReader(reader io.Reader) (*MarketplacePlugin, error) {
	plugin := MarketplacePlugin{}
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&plugin)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return &plugin, nil
}

// PluginsFromReader decodes a json-encoded list of plugins from the given io.Reader.
func BasePluginsFromReader(reader io.Reader) ([]*BaseMarketplacePlugin, error) {
	plugins := []*BaseMarketplacePlugin{}
	decoder := json.NewDecoder(reader)

	err := decoder.Decode(&plugins)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return plugins, nil
}

func PluginsFromReader(reader io.Reader) ([]*MarketplacePlugin, error) {
	plugins := []*MarketplacePlugin{}
	decoder := json.NewDecoder(reader)

	err := decoder.Decode(&plugins)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return plugins, nil
}
