package generator

type IDGenerator interface {
	NextID() int64
}
