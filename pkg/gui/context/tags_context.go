package context

import (
	"github.com/jesseduffield/gocui"
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

type TagsContext struct {
	*BasicViewModel[*models.Tag]
	*ListContextTrait
}

var (
	_ types.IListContext    = (*TagsContext)(nil)
	_ types.DiffableContext = (*TagsContext)(nil)
)

func NewTagsContext(
	getModel func() []*models.Tag,
	view *gocui.View,
	getDisplayStrings func(startIdx int, length int) [][]string,

	c *types.HelperCommon,
) *TagsContext {
	viewModel := NewBasicViewModel(getModel)

	return &TagsContext{
		BasicViewModel: viewModel,
		ListContextTrait: &ListContextTrait{
			Context: NewSimpleContext(NewBaseContext(NewBaseContextOpts{
				View:       view,
				WindowName: "branches",
				Key:        TAGS_CONTEXT_KEY,
				Kind:       types.SIDE_CONTEXT,
				Focusable:  true,
			}), ContextCallbackOpts{}),
			list:              viewModel,
			getDisplayStrings: getDisplayStrings,
			c:                 c,
		},
	}
}

func (self *TagsContext) GetSelectedItemId() string {
	item := self.GetSelected()
	if item == nil {
		return ""
	}

	return item.ID()
}

func (self *TagsContext) GetSelectedRef() types.Ref {
	tag := self.GetSelected()
	if tag == nil {
		return nil
	}
	return tag
}

func (self *TagsContext) GetDiffTerminals() []string {
	itemId := self.GetSelectedItemId()

	return []string{itemId}
}
