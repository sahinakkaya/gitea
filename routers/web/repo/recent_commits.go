// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package repo

import (
	"net/http"

	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
	contributors_service "code.gitea.io/gitea/services/repository"
)

const (
	tplRecentCommits base.TplName = "repo/activity"
)

// RecentCommits renders the page to show recent commit frequency on repository
func RecentCommits(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("repo.recent_commits")

	ctx.Data["PageIsActivity"] = true
	ctx.Data["PageIsRecentCommits"] = true

	ctx.HTML(http.StatusOK, tplRecentCommits)
}

// RecentCommitsData returns JSON of recent commits data
func RecentCommitsData(ctx *context.Context) {
	if contributorStats, err := contributors_service.GetContributorStats(ctx, ctx.Cache, ctx.Repo.Repository, ctx.Repo.CommitID); err != nil {
		ctx.ServerError("GetCodeFrequencyData", err)
	} else {
		ctx.JSON(http.StatusOK, contributorStats["total"].Weeks)
	}
}
