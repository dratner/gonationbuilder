package gonb

type Payload struct {
	Person Person `json:"person"`
}

type PersonHook struct {
	Token   string  `json:"token"`
	Payload Payload `json:"payload"`
}
