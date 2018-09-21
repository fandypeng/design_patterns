package main

import (
	"design_patterns/adapter"
	"design_patterns/builder"
	"design_patterns/decorator"
	"design_patterns/factory"
	"design_patterns/observer"
	"design_patterns/singleton"
	"fmt"
)

//a simple game demo
func main() {
	//start game
	singleton.GetDbInstance()
	singleton.GetGameConf()

	player := new(factory.Player)
	player.SetGender()
	player.SetName()

	monsterA := new(factory.Monster)
	monsterA.SetGender()
	monsterA.SetSkin()
	monsterB := new(factory.Monster)
	monsterB.SetGender()
	monsterB.SetSkin()

	//create role with simple factory
	monsterC := factory.Create(factory.RoleTypeOfMonster)
	monsterC.Walk()

	//create monster with builder pattern
	directer := builder.CreateDirecter()
	monster := directer.Create(factory.RoleTypeOfMonster)

	//move and do task
	player.Walk()
	monster.Attack()
	player.Run()

	// add dance to monster with decorator pattern
	danceMonster := decorator.DanceMonster{Behavior: monster}
	danceMonster.Dance()
	danceMonster.Walk()

	//learn and use skill
	freeze := adapter.Frozen{}
	player.ReleaseSkill(freeze)

	talk := adapter.SkillAdapter{NpcSkill: adapter.NpcSkill{}}
	player.ReleaseSkill(talk)

	obsmgr := observer.GetObsMgr()

	//new task watcher
	taskWatcher := observer.NewTask()
	obsmgr.AddObserver(taskWatcher)
	//new achievement watcher
	achieveWatcher := observer.NewAchievement()
	obsmgr.AddObserver(achieveWatcher)

	//kill monster
	monster.BeenAtacked()
	//make friends

	fmt.Println("end")
}
