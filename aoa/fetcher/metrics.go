// Copyright 2021 The go-aoa Authors
// This file is part of the go-aoa library.
//
// The the go-aoa library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The the go-aoa library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-aoa library. If not, see <http://www.gnu.org/licenses/>.

// Contains the metrics collected by the fetcher.

package fetcher

import (
	"github.com/Aurorachain-io/go-aoa/metrics"
)

var (
	propAnnounceInMeter   = metrics.NewMeter("em/fetcher/prop/announces/in")
	propAnnounceOutTimer  = metrics.NewTimer("em/fetcher/prop/announces/out")
	propAnnounceDropMeter = metrics.NewMeter("em/fetcher/prop/announces/drop")
	propAnnounceDOSMeter  = metrics.NewMeter("em/fetcher/prop/announces/dos")

	propBroadcastInMeter   = metrics.NewMeter("em/fetcher/prop/broadcasts/in")
	propBroadcastOutTimer  = metrics.NewTimer("em/fetcher/prop/broadcasts/out")
	propBroadcastDropMeter = metrics.NewMeter("em/fetcher/prop/broadcasts/drop")
	propBroadcastDOSMeter  = metrics.NewMeter("em/fetcher/prop/broadcasts/dos")

	headerFetchMeter = metrics.NewMeter("em/fetcher/fetch/headers")
	bodyFetchMeter   = metrics.NewMeter("em/fetcher/fetch/bodies")

	headerFilterInMeter  = metrics.NewMeter("em/fetcher/filter/headers/in")
	headerFilterOutMeter = metrics.NewMeter("em/fetcher/filter/headers/out")
	bodyFilterInMeter    = metrics.NewMeter("em/fetcher/filter/bodies/in")
	bodyFilterOutMeter   = metrics.NewMeter("em/fetcher/filter/bodies/out")
)
