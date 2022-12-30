package gui

import (
	"github.com/jesseduffield/lazygit/pkg/utils"
)

// list panel functions

func (gui *Gui) onSubCommitFocus() error {
	context := gui.State.Contexts.SubCommits
	if context.GetSelectedLineIdx() > COMMIT_THRESHOLD && context.GetLimitCommits() {
		context.SetLimitCommits(false)
		go utils.Safe(func() {
			if err := gui.refreshSubCommitsWithLimit(); err != nil {
				_ = gui.c.Error(err)
			}
		})
	}

	return nil
}
