package boo

type Boo struct {
	// 大寫開頭, public (不同 pkg 可以 access)
	Foo int
	// 小寫開頭, private (同 pkg 可以 access)
	bar int
}
