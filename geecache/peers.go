package geecache

// PeerPicker is the interface to locate the peers that owns a specific
type PeerPicker interface {
	// PickPeer 根据传入传入的key选择相应节点的PeerGetter
	PickPeer(key string) (peer PeerGetter, ok bool)
}
type PeerGetter interface {
	// Get 从对应的group中查找缓存值
	Get(group string, key string) ([]byte, error)
}
