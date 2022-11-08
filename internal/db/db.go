/* SPDX-FileCopyrightText: © Capsule Social, Inc. <nadim@capsule.social>
 * SPDX-License-Identifier: AGPL-3.0-only */

package db

import (
	"sync"

	"github.com/go-fed/activity/streams/vocab"
)

type DB struct {
	// The content of our app, keyed by ActivityPub ID.
	content *sync.Map
	// Enables mutations. A sync.Mutex per ActivityPub ID.
	locks *sync.Map
	// The host domain of our service, for detecting ownership.
	hostname string
}

// Our DBContent map will store this data.
type DBContent struct {
	// The payload of the data: vocab.Type is any type understood by Go-Fed.
	data vocab.Type
	// If true, belongs to our local user and not a federated peer. This is
	// recommended for a solution that just indiscriminately puts everything
	// into a single "table", like this in-memory solution.
	isLocal bool
}

func (db *DB) Construct(content *sync.Map, locks *sync.Map, hostname string) {
	db.content = content
	db.locks = locks
	db.hostname = hostname
}
