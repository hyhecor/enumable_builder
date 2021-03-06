package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"strings"
)

const content = `
{{- $package := .Package  -}}
// This code was auto-generated by github.com/hyhecor/enumable_builder; DO NOT EDIT.
// $ enumable_builder
{{- printf " -P %s" $package -}}
{{- range .Types }}
{{- printf " %s" . }}
{{- end }}

package {{ $package }}

`

const content2 = `
{{- $TARGET_TYPE := .Target.TypeName -}}
{{- $TARGET_ENUM := printf "Slice%s"  .Target.Title -}}

//{{ $TARGET_ENUM }} is actual []{{ $TARGET_TYPE }}
type {{ $TARGET_ENUM }} []{{ $TARGET_TYPE }}

//Append elements for {{ $TARGET_ENUM }}
func (s {{ $TARGET_ENUM }}) Append(e ...{{ $TARGET_TYPE }}) {{ $TARGET_ENUM }} {
	out := make({{ $TARGET_ENUM }}, 0, len(s)+len(e))
	out = append(out, s...)
	out = append(out, e...)
	return out
}

//Len {{ $TARGET_ENUM }}
func (s {{ $TARGET_ENUM }}) Len() int {
	return len(s)
}

//Fold {{ $TARGET_ENUM }}
func (s {{ $TARGET_ENUM }}) Fold(init {{ $TARGET_TYPE }}, folder func(a, b {{ $TARGET_TYPE }}) {{ $TARGET_TYPE }}) {{ $TARGET_TYPE }} {
	out := init
	for _, item := range s {
		out = folder(out, item)
	}
	return out
}

//IndexFold {{ $TARGET_ENUM }}
func (s {{ $TARGET_ENUM }}) IndexFold(init {{ $TARGET_TYPE }}, folder func(index int, a, b {{ $TARGET_TYPE }}) {{ $TARGET_TYPE }}) {{ $TARGET_TYPE }} {
	out := init
	for index, item := range s {
		out = folder(index, out, item)
	}
	return out
}

{{ range .Dest }}

{{- $DESTINATION_TYPE := .TypeName  -}}
{{- $DESTINATION_ENUM := printf "Slice%s"  .Title  -}}
{{- $FUNC_MAP := printf "Map%s"  .Title  -}}
{{- $FUNC_INDEX_MAP := printf "IndexMap%s"  .Title  -}}

//{{ $FUNC_MAP }} []{{ $TARGET_TYPE }}->[]{{ $DESTINATION_TYPE }}
func (s {{ $TARGET_ENUM }}) {{ $FUNC_MAP }}(mapper func(item {{ $TARGET_TYPE }}) {{ $DESTINATION_TYPE }}) {{ $DESTINATION_ENUM }} {
	out := make({{ $DESTINATION_ENUM }}, len(s))
	for index, item := range s {
		out[index] = mapper(item)
	}
	return out
}

//{{ $FUNC_INDEX_MAP }} []{{ $TARGET_TYPE }}->[]{{ $DESTINATION_TYPE }}
func (s {{ $TARGET_ENUM }}) {{ $FUNC_INDEX_MAP }}(mapper func(index int, item {{ $TARGET_TYPE }}) {{ $DESTINATION_TYPE }}) {{ $DESTINATION_ENUM }} {
	out := make({{ $DESTINATION_ENUM }}, len(s))
	for index, item := range s {
		out[index] = mapper(index, item)
	}
	return out
}

{{ end }}
`

//Args Commnads
type Args struct {
	Help    bool
	Version bool
	// WorkDir string
	Package string
	Types   []string
}

var (
	version = "0.0.0"
	args    Args
	stdout  io.Writer
)

func init() {
	stdout = os.Stdout

	log.SetOutput(os.Stderr)
	flag.CommandLine.SetOutput(os.Stderr)
}

func init() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s %s:\n", os.Args[0], version)
		fmt.Fprintf(flag.CommandLine.Output(), "[OPTION ...] Type ...\n")
		flag.PrintDefaults()
	}

	// flag.StringVar(&args.WorkDir, "I", "enumable", "Working Dirctory")
	flag.StringVar(&args.Package, "P", "enumable", "Package")

	flag.BoolVar(&args.Help, "h", false, "help")
	flag.BoolVar(&args.Version, "version", false, "version")

	flag.Parse()

	if args.Help {
		flag.Usage()
		os.Exit(1)
	}
	if args.Version {
		fmt.Fprintf(flag.CommandLine.Output(), "%s", version)
		os.Exit(1)
	}
	if 1 > len(flag.Args()) {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {

	title := func(s string) tmplType {

		pre := ""
		title := s

		dim := []string{"[]", "*"}
		r := true
		for r {
			indexies := []int{len(s), len(s)}
			for idx, d := range dim {
				i := strings.Index(title, d)
				if -1 != i {
					indexies[idx] = i
				}
			}
			if -1 != indexies[0] && indexies[0] < indexies[1] {
				title = title[indexies[0]+len(dim[0]):]
				pre += "Arr"
			} else if -1 != indexies[1] && indexies[0] > indexies[1] {
				title = title[indexies[1]+len(dim[1]):]
				pre += "Ptr"
			} else {
				r = false
			}
		}

		return tmplType{
			TypeName: s,
			Title:    fmt.Sprintf("%s%s", pre, strings.Title(title)),
		}
	}

	//workingDir := flagWorkingDir
	packageName := args.Package
	// taget := "string"
	dest := strings.Split("int string byte []byte [][]byte *byte **byte []*byte *[]byte", " ")
	if 0 < len(flag.Args()) {
		dest = []string{}
	}
	for _, arg := range flag.Args() {
		dest = append(dest, arg)
	}

	datahead := tmplHead{
		Package: packageName,
	}

	databodys := []tmplData{}
	for _, taget := range dest {

		datahead.Types = append(datahead.Types, taget)

		data := tmplData{}
		data.Target = title(taget)
		for _, v := range dest {
			data.Dest = append(data.Dest, title(v))
		}
		databodys = append(databodys, data)
	}

	tmpl, err := template.New("content").Parse(content)
	catch(err)

	err = tmpl.Execute(stdout, datahead)
	catch(err)

	tmpl, err = template.New("content2").Parse(content2)
	catch(err)

	for _, data := range databodys {
		err = tmpl.Execute(stdout, data)
		catch(err)
	}

}

type tmplHead struct {
	Package string
	Types   []string
}

type tmplType struct {
	TypeName string
	Title    string
}

type tmplData struct {
	Target tmplType
	Dest   []tmplType
}

func catch(err error) {
	if err != nil {
		log.Panic(err)
	}
}
