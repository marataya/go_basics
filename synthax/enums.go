package synthax

import "fmt"

type WeaponType int

func (w WeaponType) String() string {
	switch w {
	case Axe:
		return "AXE"
	case Sword:
		return "SWORD"
	case WoodenStick:
		return "WOODEN STICK"
	}
	return ""
}

const (
	Axe WeaponType = iota
	Sword
	WoodenStick
)

func getDamage(wt WeaponType) int {
	switch wt {
	case Axe:
		return 100
	case Sword:
		return 90
	case WoodenStick:
		return 1
	default:
		panic("weapon doesnt exist")
	}
}

func main() {
	fmt.Printf("Damage of weapon %s is %d\n", Axe, getDamage(Axe))
	fmt.Printf("Damage of weapon %s is %d\n", Sword, getDamage(Sword))
	fmt.Printf("Damage of weapon %s is %d\n", WoodenStick, getDamage(WoodenStick))
}
