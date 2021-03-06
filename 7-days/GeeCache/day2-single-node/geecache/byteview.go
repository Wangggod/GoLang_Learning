package geecache

//ByteView用来表示缓存值
type ByteView struct {
	b []byte //b将会存储真实的缓存值，且为了防止被修改所以是只读的
}

//使用lru.go中的Value接口的Len方法返回值
func (v ByteView) Len() int {
	return len(v.b)
}

//由于b是只读的，因此要修改我们需要复制一份可以修改的
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

//为了返回String，我们还要将复制的结果转换
func (v ByteView) String() string {
	return string(v.b)
}
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
