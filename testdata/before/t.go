package testdata

func Unchanged1(n int)
func Unchanged2(n int) error
func Unchanged3(n int) error

func Compatible1(n int)
func Compatible2(n int)

func Breaking1(n int)
func Breaking2(n int) []byte
func Breaking3(n int, s string)
func Breaking4(n int) string
func Removed1()

type RemovedT1 bool
type UnchangedT1 int
type UnchangedT2 struct {
	Foo string
}
type UnchangedT3 struct {
	// Foo is a foo
	Foo string
}
type CompatibleT1 struct {
	Foo string
}
type CompatibleT2 struct {
	Foo string
}
type BreakingT1 struct {
	XXX string
}
