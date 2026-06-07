package generator

type CodeGenerator interface {
	Generate(id int64) string
}
