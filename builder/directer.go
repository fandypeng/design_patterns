package builder

import "design_patterns/factory"

type Directer struct {
	builder builder
}

/**
1、创建游戏中的多种角色，玩家、NPC、怪物、Boss等
2、这些角色有这大部分类似的属性和行为
3、随时可以增加新的角色
4、不同角色创建过程有不一样的处理
 */
func(d Directer) Create(roleType int) factory.Behavior {
	d.builder = builder{}
	switch roleType {
	case factory.RoleTypeOfPlayer:
		d.builder.SetRole(factory.Player{})
		// build by sequence 人物要先设置性别再设置皮肤
		d.builder.SetGender().SetSkin().SetName().SetHeight().Build()
	case factory.RoleTypeOfMonster:
		d.builder.SetRole(factory.Monster{})
		// build by sequence 怪物要先设置性别和身高再设置皮肤
		d.builder.SetGender().SetHeight().SetSkin().SetName().Build()
	default:
		d.builder.SetRole(factory.Player{})
		d.builder.SetGender().SetSkin().SetName().SetHeight().Build()
	}

	return d.builder.role
}


func CreateDirecter() Directer {
	return Directer{}
}