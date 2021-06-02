package woocommerce

import (
	"strconv"
	"strings"
	"time"
)

const (
	defaultAPIPrefix = "/wp-json/wc/"
	defaultUserAgent = "WooCommerce API Client-Go"

	defaultVersion = "v3"
	defaultPrefix  = "/wc-api/"
	defaultTimeout = time.Second * 15
)

type BasicOptions struct {
	URL             string          // Store URL
	Key             string          // API Consumer Key
	Secret          string          // API Consumer Secret
	AdvancedOptions AdvancedOptions // Advanced Options (see below)
}

type AdvancedOptions struct {
	WPAPI                bool   // Allow make requests to the new WP REST API integration
	WPAPIPrefix          string // Custom WP REST API URL prefix
	Version              string // API version, default is v3
	Timeout              int    // Request timeout, default is 15
	FollowRedirects      bool   // Allow the API call to follow redirects
	verifySslInitialized bool
	VerifySsl            bool   // Verify SSL when connect
	QueryStringAuth      bool   // Force Basic Authentication
	OAuthTimestamp       string // Custom oAuth timestamp
	UserAgent            string // Custom user-agent
}

// Version of WooCommerce API which will be used
func (o *BasicOptions) Version() string {
	if o.AdvancedOptions.Version == "" {
		return defaultVersion
	}
	return o.AdvancedOptions.Version
}

// VerifySsl return requirement of SSL certificate check
func (o *BasicOptions) VerifySsl() bool {
	if !o.AdvancedOptions.verifySslInitialized {
		return true
	}
	return o.AdvancedOptions.VerifySsl
}

// DisableSslVerification in case of debug
func (o *BasicOptions) DisableSslVerification() {
	o.AdvancedOptions.verifySslInitialized = true
	o.AdvancedOptions.VerifySsl = false
}

// Timeout per request to WooCommerce API
func (o *BasicOptions) Timeout() int {
	if o.AdvancedOptions.Timeout == 0 {
		return int(defaultTimeout)
	}
	return o.AdvancedOptions.Timeout
}

// QueryStringAuth returns true when credentials will pass through query string
func (o *BasicOptions) QueryStringAuth() bool {
	return o.AdvancedOptions.QueryStringAuth
}

// WPAPI returns if used WordPress API or not
func (o *BasicOptions) WPAPI() bool {
	return o.AdvancedOptions.WPAPI
}

// WPAPIPrefix is a prefix for WordPress API requests
func (o *BasicOptions) WPAPIPrefix() string {
	if o.AdvancedOptions.WPAPIPrefix == "" {
		return defaultPrefix
	}
	return o.AdvancedOptions.WPAPIPrefix
}

// OAuthTimestamp return OAuth specific timestamp
func (o *BasicOptions) OAuthTimestamp() string {
	if o.AdvancedOptions.OAuthTimestamp == "" {
		return strconv.FormatInt(time.Now().Unix(), 10)
	}
	return o.AdvancedOptions.OAuthTimestamp
}

// UserAgent which will be use for requests
func (o *BasicOptions) UserAgent() string {
	if o.AdvancedOptions.UserAgent == "" {
		return defaultUserAgent
	}
	return o.AdvancedOptions.UserAgent
}

// FollowRedirects during requests
func (o *BasicOptions) FollowRedirects() bool {
	return o.AdvancedOptions.FollowRedirects
}

func (o *BasicOptions) getAPIPrefix() string {
	if o.WPAPI() {
		return o.WPAPIPrefix()
	}
	return defaultAPIPrefix
}

func (o *BasicOptions) getBaseURL() string {
	return strings.TrimRight(o.URL, "/") + o.getAPIPrefix() + o.Version() + "/"

}

func (o *BasicOptions) buildURL(endpoint string) string {
	return o.getBaseURL() + endpoint

}
