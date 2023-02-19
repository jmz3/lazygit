// THIS FILE IS AUTO-GENERATED. You can regenerate it by running `go generate ./...` at the root of the lazygit repo.

package tests

import (
	"github.com/jesseduffield/lazygit/pkg/integration/components"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/bisect"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/branch"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/cherry_pick"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/commit"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/config"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/custom_commands"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/diff"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/file"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/filter_by_path"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/interactive_rebase"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/misc"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/patch_building"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/stash"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/submodule"
	"github.com/jesseduffield/lazygit/pkg/integration/tests/sync"
)

var tests = []*components.IntegrationTest{
	bisect.Basic,
	bisect.FromOtherBranch,
	branch.CheckoutByName,
	branch.Delete,
	branch.DetachedHead,
	branch.Rebase,
	branch.RebaseAndDrop,
	branch.RebaseDoesNotAutosquash,
	branch.Reset,
	branch.Suggestions,
	cherry_pick.CherryPick,
	cherry_pick.CherryPickConflicts,
	commit.Commit,
	commit.CommitMultiline,
	commit.DiscardOldFileChange,
	commit.NewBranch,
	commit.Revert,
	commit.StageRangeOfLines,
	commit.Staged,
	commit.StagedWithoutHooks,
	commit.Unstaged,
	config.RemoteNamedStar,
	custom_commands.Basic,
	custom_commands.FormPrompts,
	custom_commands.MenuFromCommand,
	custom_commands.MenuFromCommandsOutput,
	custom_commands.MultiplePrompts,
	diff.Diff,
	diff.DiffAndApplyPatch,
	diff.DiffCommits,
	diff.IgnoreWhitespace,
	file.DirWithUntrackedFile,
	file.DiscardChanges,
	file.DiscardStagedChanges,
	file.Gitignore,
	filter_by_path.CliArg,
	filter_by_path.SelectFile,
	filter_by_path.TypeFile,
	interactive_rebase.AmendMerge,
	interactive_rebase.One,
	misc.ConfirmOnQuit,
	patch_building.CopyPatchToClipboard,
	stash.Rename,
	stash.Stash,
	stash.StashIncludingUntrackedFiles,
	submodule.Add,
	submodule.Enter,
	submodule.Remove,
	submodule.Reset,
	sync.FetchPrune,
	sync.ForcePush,
	sync.Pull,
	sync.PullAndSetUpstream,
	sync.RenameBranchAndPull,
}
