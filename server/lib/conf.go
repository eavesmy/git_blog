package lib

type Conf struct {
	Root   string `json:"root"`
	Update int    `json:"update"`
}

var RootConf *Conf

func InitConf(root string, update int) {

	RootConf = &Conf{Root: root, Update: update}
	Update()
}
