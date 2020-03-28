// Copyright 2015 Light Code Labs, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package caddyhttp

import (
	// plug in the server
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/httpserver"

	// plug in the standard directives
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/basicauth"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/bind"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/browse"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/errors"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/expvar"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/extensions"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/fastcgi"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/gzip"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/header"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/index"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/internalsrv"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/limits"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/log"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/markdown"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/mime"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/pprof"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/proxy"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/push"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/redirect"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/requestid"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/rewrite"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/root"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/status"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/templates"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/timeouts"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/caddyhttp/websocket"
	_ "github.com/haochudao/caddyserver-caddy_v1.0.3/onevent"
)
