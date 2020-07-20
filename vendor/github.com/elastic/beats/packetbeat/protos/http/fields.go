// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package http

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("packetbeat", "http", asset.ModuleFieldsPri, AssetHttp); err != nil {
		panic(err)
	}
}

// AssetHttp returns asset data.
// This is the base64 encoded gzipped contents of protos/http.
func AssetHttp() string {
	return "eJzUVMuOlUAQ3fMVJ7N2+AAWJm6MszEubuJyUkAB7dCP6a72yt+b5nEdoKMxNy5kBfU451T1aR7xwlOFQcQVgCgZucLDp8vly0MBtBwar5woayqk4GNw3KhONeDvbASd4rENZYH1rSoA4BGGNN9Q0yOT4wq9t3GL7LCfTGe9pvQBqm0UyMAzIzy/Rg4CMi08B2dN4HLFeEv6lnjtucUzk2RqzhpzHLsBmVr2YZfbcGz9jRs5pJbg81LxwtPV+vZQslP6/pAEPkCTQ2ONkDLK9POiGnISPberoFUzOm/1nF9nLU9oXwfVDNsYELshQYXE0ak+eqpHLvHU3cquSoYZNpDmE+QqIS0I5BnOc0hWUWbu0RwC9fwufUy4qnFEzQjsyJNwi3o6ITZWawplkT2C1KfzJ0CjomNGq97Tslvx8ajekQwVoh/L18h+KjLGWgz4B2ediv7eWkFIYnh2g6dwlLmjvGwXZenA0nE8av5B2qW7/dkKPtpo2vw6/x9H7/8Ev567LH1CSxa/09InzHr6vaUb2x6F3Gfo9B8ubxtbnTWz5K9UznP/QMHK8zMAAP//lgPSsQ=="
}
