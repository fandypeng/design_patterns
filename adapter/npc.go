package adapter

type NpcSkill struct {

}

func (n NpcSkill) Talk() {

}

type SkillAdapter struct {
	NpcSkill
}

func (s SkillAdapter) Before() {

}

func (s SkillAdapter) Release() {
	s.NpcSkill.Talk()
}
func (s SkillAdapter) After() {

}
