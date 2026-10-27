package types

type RecipeIngredient struct {
	ItemDescriptor struct {
		ControlValue         uint32
		EmptyItemDescriptor  *EmptyItemDescriptor
		ItemNameDescriptor   *ItemNameDescriptor
		MolangItemDescriptor *MolangItemDescriptor
		ItemTagDescriptor    *ItemTagDescriptor
	}
	StackSize uint16
}
