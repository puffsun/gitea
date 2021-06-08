// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package setting

import (
	"code.gitea.io/gitea/modules/log"
)

var (
	// Draw.io config that define draw.io configuration
	DrawioConfig = struct {
		Enabled  bool
		Endpoint string
		Suffix   string
	}{
		Enabled:  true,
		Endpoint: "http://shaverif.apac.ads.idt.com:8004/",
		Suffix:   "drawio",
	}
)

func newDrawioService() {
	sec := Cfg.Section("drawio")
	if err := sec.MapTo(&DrawioConfig); err != nil {
		log.Fatal("Failed to map drawio settings: %v", err)
	}

	if DrawioConfig.Enabled {
		log.Info("Draw.io Service Enabled")
	}
}
