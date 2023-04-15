package controllers

import (
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/gui/context"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/jesseduffield/lazygit/pkg/utils"
)

type StashController struct {
	baseController
	*controllerCommon
}

var _ types.IController = &StashController{}

func NewStashController(
	common *controllerCommon,
) *StashController {
	return &StashController{
		baseController:   baseController{},
		controllerCommon: common,
	}
}

func (self *StashController) GetKeybindings(opts types.KeybindingsOpts) []*types.Binding {
	bindings := []*types.Binding{
		{
			Key:         opts.GetKey(opts.Config.Universal.Select),
			Handler:     self.checkSelected(self.handleStashApply),
			Description: self.c.Tr.LcApply,
		},
		{
			Key:         opts.GetKey(opts.Config.Stash.PopStash),
			Handler:     self.checkSelected(self.handleStashPop),
			Description: self.c.Tr.LcPop,
		},
		{
			Key:         opts.GetKey(opts.Config.Universal.Remove),
			Handler:     self.checkSelected(self.handleStashDrop),
			Description: self.c.Tr.LcDrop,
		},
		{
			Key:         opts.GetKey(opts.Config.Universal.New),
			Handler:     self.checkSelected(self.handleNewBranchOffStashEntry),
			Description: self.c.Tr.LcNewBranch,
		},
		{
			Key:         opts.GetKey(opts.Config.Stash.RenameStash),
			Handler:     self.checkSelected(self.handleRenameStashEntry),
			Description: self.c.Tr.LcRenameStash,
		},
	}

	return bindings
}

func (self *StashController) GetOnRenderToMain() func() error {
	return func() error {
		return self.helpers.Diff.WithDiffModeCheck(func() error {
			var task types.UpdateTask
			stashEntry := self.context().GetSelected()
			if stashEntry == nil {
				task = types.NewRenderStringTask(self.c.Tr.NoStashEntries)
			} else {
				task = types.NewRunPtyTask(self.git.Stash.ShowStashEntryCmdObj(stashEntry.Index).GetCmd())
			}

			return self.c.RenderToMainViews(types.RefreshMainOpts{
				Pair: self.c.MainViewPairs().Normal,
				Main: &types.ViewUpdateOpts{
					Title: "Stash",
					Task:  task,
				},
			})
		})
	}
}

func (self *StashController) checkSelected(callback func(*models.StashEntry) error) func() error {
	return func() error {
		item := self.context().GetSelected()
		if item == nil {
			return nil
		}

		return callback(item)
	}
}

func (self *StashController) Context() types.Context {
	return self.context()
}

func (self *StashController) context() *context.StashContext {
	return self.contexts.Stash
}

func (self *StashController) handleStashApply(stashEntry *models.StashEntry) error {
	apply := func() error {
		self.c.LogAction(self.c.Tr.Actions.Stash)
		err := self.git.Stash.Apply(stashEntry.Index)
		_ = self.postStashRefresh()
		if err != nil {
			return self.c.Error(err)
		}
		return nil
	}

	if self.c.UserConfig.Gui.SkipStashWarning {
		return apply()
	}

	return self.c.Confirm(types.ConfirmOpts{
		Title:  self.c.Tr.StashApply,
		Prompt: self.c.Tr.SureApplyStashEntry,
		HandleConfirm: func() error {
			return apply()
		},
	})
}

func (self *StashController) handleStashPop(stashEntry *models.StashEntry) error {
	pop := func() error {
		self.c.LogAction(self.c.Tr.Actions.Stash)
		err := self.git.Stash.Pop(stashEntry.Index)
		_ = self.postStashRefresh()
		if err != nil {
			return self.c.Error(err)
		}
		return nil
	}

	if self.c.UserConfig.Gui.SkipStashWarning {
		return pop()
	}

	return self.c.Confirm(types.ConfirmOpts{
		Title:  self.c.Tr.StashPop,
		Prompt: self.c.Tr.SurePopStashEntry,
		HandleConfirm: func() error {
			return pop()
		},
	})
}

func (self *StashController) handleStashDrop(stashEntry *models.StashEntry) error {
	return self.c.Confirm(types.ConfirmOpts{
		Title:  self.c.Tr.StashDrop,
		Prompt: self.c.Tr.SureDropStashEntry,
		HandleConfirm: func() error {
			self.c.LogAction(self.c.Tr.Actions.Stash)
			err := self.git.Stash.Drop(stashEntry.Index)
			_ = self.c.Refresh(types.RefreshOptions{Scope: []types.RefreshableView{types.STASH}})
			if err != nil {
				return self.c.Error(err)
			}
			return nil
		},
	})
}

func (self *StashController) postStashRefresh() error {
	return self.c.Refresh(types.RefreshOptions{Scope: []types.RefreshableView{types.STASH, types.FILES}})
}

func (self *StashController) handleNewBranchOffStashEntry(stashEntry *models.StashEntry) error {
	return self.helpers.Refs.NewBranch(stashEntry.RefName(), stashEntry.Description(), "")
}

func (self *StashController) handleRenameStashEntry(stashEntry *models.StashEntry) error {
	message := utils.ResolvePlaceholderString(
		self.c.Tr.RenameStashPrompt,
		map[string]string{
			"stashName": stashEntry.RefName(),
		},
	)

	return self.c.Prompt(types.PromptOpts{
		Title:          message,
		InitialContent: stashEntry.Name,
		HandleConfirm: func(response string) error {
			self.c.LogAction(self.c.Tr.Actions.RenameStash)
			err := self.git.Stash.Rename(stashEntry.Index, response)
			_ = self.c.Refresh(types.RefreshOptions{Scope: []types.RefreshableView{types.STASH}})
			if err != nil {
				return err
			}
			self.context().SetSelectedLineIdx(0) // Select the renamed stash
			return nil
		},
	})
}
