package presenter

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Success struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
