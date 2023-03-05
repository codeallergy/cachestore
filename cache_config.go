/*
 * Copyright (c) 2022-2023 Zander Schwid & Co. LLC.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 */

package cachestore

import (
	"errors"
	"time"
)

var (
	ErrCanceled         = errors.New("operation was canceled")
)

type Config struct {
	DefaultExpiration time.Duration
	CleanupInterval   time.Duration
}

// Option configures memory storage using the functional options paradigm
// popularized by Rob Pike and Dave Cheney. If you're unfamiliar with this style,
// see https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html and
// https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis.
type Option interface {
	apply(*Config)
}

// OptionFunc implements Option interface.
type optionFunc func(*Config)

// apply the configuration to the provided config.
func (fn optionFunc) apply(r *Config) {
	fn(r)
}

// option that do nothing
func WithNope() Option {
	return optionFunc(func(opts *Config) {
	})
}

func WithDefaultExpiration(value time.Duration) Option {
	return optionFunc(func(opts *Config) {
		opts.DefaultExpiration = value
	})
}

func WithCleanupInterval(value time.Duration) Option {
	return optionFunc(func(opts *Config) {
		opts.CleanupInterval = value
	})
}


