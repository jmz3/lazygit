package controllers

import (
	"github.com/jesseduffield/lazygit/pkg/gui/context"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

type SubCommitsController struct {
	baseController
	*controllerCommon
	context *context.SubCommitsContext
}

var _ types.IController = &SubCommitsController{}

func NewSubCommitsController(
	common *controllerCommon,
	context *context.SubCommitsContext,
) *SubCommitsController {
	return &SubCommitsController{
		baseController:   baseController{},
		controllerCommon: common,
		context:          context,
	}
}

func (self *SubCommitsController) Context() types.Context {
	return self.context
}

func (self *SubCommitsController) GetOnRenderToMain() func() error {
	return func() error {
		return self.helpers.Diff.WithDiffModeCheck(func() error {
			commit := self.context.GetSelected()
			var task types.UpdateTask
			if commit == nil {
				task = types.NewRenderStringTask("No commits")
			} else {
				cmdObj := self.git.Commit.ShowCmdObj(commit.Sha, self.modes.Filtering.GetPath(), self.c.State().GetIgnoreWhitespaceInDiffView())

				task = types.NewRunPtyTask(cmdObj.GetCmd())
			}

			return self.c.RenderToMainViews(types.RefreshMainOpts{
				Pair: self.c.MainViewPairs().Normal,
				Main: &types.ViewUpdateOpts{
					Title: "Commit",
					Task:  task,
				},
			})
		})
	}
}
