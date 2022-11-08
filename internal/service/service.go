/* SPDX-FileCopyrightText: © Capsule Social, Inc. <nadim@capsule.social>
 * SPDX-License-Identifier: AGPL-3.0-only */

package service

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/go-fed/activity/pub"
	"github.com/go-fed/activity/streams/vocab"
)

type Service struct{}

func (*Service) AuthenticateGetInbox(c context.Context,
	w http.ResponseWriter,
	r *http.Request) (out context.Context, authenticated bool, err error) {
	// TODO
	return
}

func (*Service) AuthenticateGetOutbox(c context.Context,
	w http.ResponseWriter,
	r *http.Request) (out context.Context, authenticated bool, err error) {
	// TODO
	return
}

func (*Service) GetOutbox(c context.Context,
	r *http.Request) (vocab.ActivityStreamsOrderedCollectionPage, error) {
	// TODO
	return nil, nil
}

func (*Service) NewTransport(c context.Context,
	actorBoxIRI *url.URL,
	gofedAgent string) (t pub.Transport, err error) {
	// TODO
	return
}

func (*Service) PostInboxRequestBodyHook(c context.Context,
	r *http.Request,
	activity Activity) (context.Context, error) {
	// TODO
	return nil, nil
}

func (*Service) AuthenticatePostInbox(c context.Context,
	w http.ResponseWriter,
	r *http.Request) (out context.Context, authenticated bool, err error) {
	// TODO
	return
}

func (*Service) Blocked(c context.Context,
	actorIRIs []*url.URL) (blocked bool, err error) {
	// TODO
	return
}

func (*Service) FederatingCallbacks(c context.Context) (wrapped FederatingWrappedCallbacks, other []interface{}, err error) {
	// TODO
	return
}

func (*Service) DefaultCallback(c context.Context,
	activity Activity) error {
	// TODO
	return nil
}

func (*Service) MaxInboxForwardingRecursionDepth(c context.Context) int {
	// TODO
	return -1
}

func (*Service) MaxDeliveryRecursionDepth(c context.Context) int {
	// TODO
	return -1
}

func (*Service) FilterForwarding(c context.Context,
	potentialRecipients []*url.URL,
	a Activity) (filteredRecipients []*url.URL, err error) {
	// TODO
	return
}

func (*Service) GetInbox(c context.Context,
	r *http.Request) (vocab.ActivityStreamsOrderedCollectionPage, error) {
	// TODO
	return nil, nil
}

func (*Service) Now() time.Time {
	return time.Now()
}
