package main

type PeerPicker interface {
	// PickPeer 根据传入的key选择相应节点的PeerGetter
	PickPeer(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	// Get 从对应group查找缓存值
	Get(group string, key string) ([]byte, error)
}
