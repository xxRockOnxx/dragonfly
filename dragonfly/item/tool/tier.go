package tool

// Tier represents the t, or material, that a tool is made of. A higher tier indicates a tool which is able
// to mine more different types of blocks at a higher speed.
type Tier struct {
	// HarvestLevel is the level that this tier of tools is able to harvest. If a block has a harvest level
	// above this one, a tool with this tier won't be able to harvest it.
	HarvestLevel int
	// BaseMiningEfficiency is the base efficiency of the tier, when it comes to mining blocks. This is
	// specifically used for tools such as pickaxes.
	BaseMiningEfficiency float64
	// BaseAttackDamage is the base attack damage of tools with this tiers. All tools have a constant value
	// that is added on top of this.
	BaseAttackDamage float32
	// Durability returns the maximum durability that a tool with this tier has.
	Durability int
}

// TierWood is the tier of wood tools. This is the lowest possible t.
var TierWood = Tier{HarvestLevel: 1, Durability: 59, BaseMiningEfficiency: 2, BaseAttackDamage: 1}

// TierGold is the tier of gold tools.
var TierGold = Tier{HarvestLevel: 1, Durability: 32, BaseMiningEfficiency: 12, BaseAttackDamage: 1}

// TierStone is the tier of stone tools.
var TierStone = Tier{HarvestLevel: 2, Durability: 131, BaseMiningEfficiency: 4, BaseAttackDamage: 2}

// TierIron is the tier of iron tools.
var TierIron = Tier{HarvestLevel: 3, Durability: 250, BaseMiningEfficiency: 6, BaseAttackDamage: 3}

// TierDiamond is the tier of diamond tools.
var TierDiamond = Tier{HarvestLevel: 4, Durability: 1561, BaseMiningEfficiency: 8, BaseAttackDamage: 4}

// TierNetherite is the tier of netherite tools. This is the highest possible tier.
// TODO: Implement netherite tools once 1.16 lands.
var TierNetherite = Tier{HarvestLevel: 4, Durability: 2031, BaseMiningEfficiency: 9, BaseAttackDamage: 5}
