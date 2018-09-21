package factory

const (
	RoleTypeOfPlayer  = 1
	RoleTypeOfMonster = 2
)

/**
1、创建游戏中的多种角色，玩家、NPC、怪物、Boss等
2、这些角色有这大部分类似的属性和行为
3、随时可以增加新的角色
4、根据roleType 读取响应的配置来生产role对象
*/
func Create(roleType int) Behavior {
	switch roleType {
	case RoleTypeOfPlayer:
		player := new(Player)
		player.SetGender()
		player.SetSkin()
		return player
	case RoleTypeOfMonster:
		monster := new(Monster)
		monster.SetGender()
		monster.SetSkin()
		return monster
	default:
		return new(Player)
	}

	return nil
}
