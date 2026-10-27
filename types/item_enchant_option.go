package types

type ItemEnchantOption struct {
	Cost         uint8
	Enchants     ItemEnchants
	EnchantName  string
	EnchantNetId RecipeNetId
}
