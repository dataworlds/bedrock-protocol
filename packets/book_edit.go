package packets

import "github.com/dataworlds/bedrock-protocol/types"

const BookEditID uint32 = 97

type BookEdit struct {
	BookSlot  int32
	Operation struct {
		ControlValue              uint32
		BookEditActionReplacePage *types.BookEditActionReplacePage
		BookEditActionAddPage     *types.BookEditActionAddPage
		BookEditActionDeletePage  *types.BookEditActionDeletePage
		BookEditActionSwapPages   *types.BookEditActionSwapPages
		BookEditActionFinalize    *types.BookEditActionFinalize
	}
}
