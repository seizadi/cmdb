package helm

import (
	"bytes"
	"github.com/Masterminds/sprig/v3"
	"github.com/pkg/errors"
	"text/template"
)

// Values represents a collection of chart values.
// type Values map[string]interface{}
// We wrap this so that we can pass it to a chart using our engine as Values
// Now apply the template to resolve references e.g. {{ Values.this.port }}
type Values struct {
	Values map[interface{}]interface{}
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
//		// Add some extra functionality
//		return template.FuncMap{
//			"tpl": tpl,
//			"required": required,
//		}
}


func tpl(t string, vals Values) (string, error) {
	var out bytes.Buffer
	tt, err := template.New("_").Parse(t)
	if err != nil {
		return "", err
	}
	err = tt.Execute(&out, vals)
	if err != nil {
		return "", err
	}

	return out.String(), nil
	//out := fmt.Sprintf("%T\n", vals)
	//out := reflect.TypeOf(vals).String()
	//return out, nil
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
