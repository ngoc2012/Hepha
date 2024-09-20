package Test

type Test struct {
	Path string
	Html string
}

func (t Test) TestInit() {
	t.Path = "Test.html"
	t.Html = "<h1>Test</h1>"
}
