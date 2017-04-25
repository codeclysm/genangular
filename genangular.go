package genangular

import (
	"goa.design/goa.v2/codegen"
	"goa.design/goa.v2/design"
	"goa.design/goa.v2/eval"
)

func Generate(roots ...eval.Root) ([]codegen.File, error) {
	var types []codegen.File

	for _, root := range roots {
		switch r := root.(type) {
		case *design.RootExpr:
			for _, t := range r.Types {
				types = append(types, Type(t))
			}
		}
	}
	return types, nil
}

func Type(usertype design.UserType) codegen.File {
	name := JavaScriptify(usertype.Name(), false, false)

	return codegen.NewSource(name+".ts", func(genPkg string) []*codegen.Section {
		data := typeData{
			Name:  name,
			Props: []string{},
		}
		body := codegen.Section{
			typeTmpl,
			data,
		}

		return []*codegen.Section{&body}
	})
}
