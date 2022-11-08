/* SPDX-FileCopyrightText: © Capsule Social, Inc. <nadim@capsule.social>
 * SPDX-License-Identifier: AGPL-3.0-only */

package cmd

import (
	"log"
	"mastogon/internal/db"
	"mastogon/internal/service"
	"sync"

	"github.com/go-fed/activity/pub"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mastogon",
	Short: "Mastodon but in Go, basically. ActivityPub! Fediverse!",
	Long:  `Long description`,
	Run: func(cmd *cobra.Command, args []string) {
		s := &service.Service{}
		db := db.DB{}
		db.Construct(&sync.Map{}, &sync.Map{}, "localhost")
		actor := pub.NewFederatingActor(s, s, db, s)
	},
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
