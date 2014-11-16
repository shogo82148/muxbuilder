package muxbuilder

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
)

type MUXBuilder struct {
	Package string
	Routes  []Route
}

type Route struct {
	URL  string
	Name string
}

type muxContext struct {
	packageName     string
	importLibraries map[string]string
	contexts        []*contextDefinition
	handlers        []*handlerDefinition
}

var regParamName = regexp.MustCompile("([:*])([^/]+)")

func (b *MUXBuilder) Build() string {
	var buf bytes.Buffer
	bc := newMUXContext(b.Package)

	bc.addImport("http", "net/http")
	bc.addImport("denco", "github.com/naoina/denco")

	for _, route := range b.Routes {
		bc.addRoute(route)
	}

	bc.write(&buf)
	return buf.String()
}

func newMUXContext(packageName string) *muxContext {
	return &muxContext{
		packageName:     packageName,
		importLibraries: map[string]string{},
	}
}

func (bc *muxContext) write(w io.Writer) error {
	fmt.Fprintf(w, "package %s\n", bc.packageName)

	fmt.Fprintln(w, "import (")
	for name, path := range bc.importLibraries {
		fmt.Fprintf(w, "%s \"%s\"\n", name, path)
	}
	fmt.Fprintln(w, ")")

	for _, cd := range bc.contexts {
		cd.write(w)
	}

	fmt.Fprintln(w, `func NewHandler() (http.Handler, error) {
mux := denco.NewMux()
return mux.Build([]denco.Handler{`)
	for _, handler := range bc.handlers {
		handler.write(w)
	}
	fmt.Fprintln(w, "})\n}")

	return nil
}

func (bc *muxContext) addRoute(r Route) {
	cd := newContextDefinition(r.Name + "Context")
	hd := newHandlerDefinition("GET", r.URL, r.Name, cd.name)
	cd.addField("ResponseWriter", "http.ResponseWriter")
	hd.addParam("ResponseWriter", "argWriter")
	cd.addField("Request", "*http.Request")
	hd.addParam("Request", "argRequest")

	for _, match := range regParamName.FindAllStringSubmatch(r.URL, -1) {
		name := match[2]
		cd.addField(name, "string")
		hd.addParam(name, `argParams.Get("`+name+`")`)
	}

	bc.addContext(cd)
	bc.addHandler(hd)
}

func (bc *muxContext) addImport(name, path string) {
	if oldpath, ok := bc.importLibraries[name]; ok && oldpath != path {
		panic("name confliction: " + name)
	}

	bc.importLibraries[name] = path
}

func (bc *muxContext) addContext(cd *contextDefinition) {
	bc.contexts = append(bc.contexts, cd)
}

func (bc *muxContext) addHandler(hd *handlerDefinition) {
	bc.handlers = append(bc.handlers, hd)
}
