package helm

import (
	"strings"
	"text/template"
	
	"github.com/Masterminds/sprig/v3"
	"github.com/pkg/errors"
)

// Values represents a collection of chart values.
// type Values map[string]interface{}
// We wrap this so that we can pass it to a chart using our engine as Values
// Now apply the template to resolve references e.g. {{ Values.this.port }}
type Values struct {
	Values map[string]interface{}
}

// renderable is an object that can be rendered.
type Renderable struct {
	// tpl is the current template.
	Tpl string
	// vals are the values to be supplied to the template.
	Vals Values
}

// We duplicate helm tempalte functions here so that we can
// use helm templates, if this pattern does not work we will
// resort to call helm instead of using go template package

func FuncMap() template.FuncMap {
	// TODO - Helm includes sprig function as listed below exclude for now
	f := sprig.TxtFuncMap()
	delete(f, "env")
	delete(f, "expandenv")

	// Add some extra functionality
	extra := template.FuncMap{
		"tpl": tpl,
		"required": required,
	}

	for k, v := range extra {
		f[k] = v
	}

	return f
}

func tpl(tpl string, vals Values) (string, error) {
	r := Renderable{
		Tpl:      tpl,
		Vals:     vals,
	}

	result, err := RenderWithReferences(r)
	if err != nil {
		return "", err
	}
	return result, nil
}

var ErrRequiredArg0Nil = errors.New("required first arg is nil")
var ErrRequiredArg0Empty = errors.New("required first arg is empty")

// Add the `required` function
func required (warn string, val interface{}) (interface{}, error) {
	if val == nil {
		return val, ErrRequiredArg0Nil
	} else if _, ok := val.(string); ok {
		if val == "" {
			return val, ErrRequiredArg0Empty
		}
	}
	return val, nil
}

func RenderWithReferences(referenceTpl Renderable) (rendered string, err error) {
	// Basically, what we do here is start with an empty parent template and then
	// build up a list of templates -- one for each file. Once all of the templates
	// have been parsed, we loop through again and execute every template.
	//
	// The idea with this process is to make it possible for more complex templates
	// to share common blocks, but to make the entire thing feel like a file-based
	// template engine.
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("rendering template failed: %v", r)
		}
	}()
	t := template.New("gotpl")
	t.Option("missingkey=zero")
	// e.initFunMap(t, referenceTpls)

	if _, err := t.New("test").Funcs(FuncMap()).Parse(referenceTpl.Tpl); err != nil {
		return "", err
	}

	vals := referenceTpl.Vals
	var buf strings.Builder
	if err := t.ExecuteTemplate(&buf, "test", vals); err != nil {
		return "", err
	}

	// Work around the issue where Go will emit "<no value>" even if Options(missing=zero)
	// is set. Since missing=error will never get here, we do not need to handle
	// the Strict case.
	rendered = strings.ReplaceAll(buf.String(), "<no value>", "")

	return rendered, nil
}
