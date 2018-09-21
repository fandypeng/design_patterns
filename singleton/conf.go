package singleton

import "sync"

type gameConf struct {
	GameId   int32
	ServreId int32
	Dsn      string
}

type gameConfOperator struct {
	gameconf *gameConf
}

var operator *gameConfOperator
var onceLoadConf sync.Once

func loadGameConf() *gameConfOperator {
	//load game conf
	return &gameConfOperator{
		gameconf: &gameConf{
			GameId:   1,
			ServreId: 2,
			Dsn:      "username:password@protocol(address)/dbname",
		},
	}
}

func GetGameConf() *gameConfOperator {
	onceLoadConf.Do(func() {
		operator = loadGameConf()
	})

	return operator
}


func (o *gameConfOperator) GetGameId() int32 {
	return o.gameconf.GameId
}

func (o *gameConfOperator) GetServerId() int32 {
	return o.gameconf.ServreId
}

func (o *gameConfOperator) GetDsn() string {
	return o.gameconf.Dsn
}
