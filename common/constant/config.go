package constant

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-07 15:13
**/

type ServerConfig struct {
	ContextPath string
	IpAddr      string
	Port        uint64
}

type ClientConfig struct {
	TimeoutMs         uint64
	ListenInterval    uint64
	BeatInterval      uint64
	SubscribeInterval uint64
}
