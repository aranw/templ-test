// Code generated by templ@v0.2.364 DO NOT EDIT.

package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Main(current_page string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<!doctype html>")
		if err != nil {
			return err
		}
		var var_2 = []any{
			"h-full",
			templ.KV("bg-gray-50", current_page == "login"),
			templ.KV("bg-gray-50", current_page == "register"),
		}
		err = templ.RenderCSSItems(ctx, templBuffer, var_2...)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<html lang=\"en\" class=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(var_2).String()))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><head><meta charset=\"utf-8\"><link rel=\"icon\" href=\"/favicon.png\"><meta name=\"viewport\" content=\"width=device-width\"><style type=\"text/css\">")
		if err != nil {
			return err
		}
		var_3 := `
                .h-full {
                    height: 100%;
                }


            `
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</style></head><body class=\"h-full\">")
		if err != nil {
			return err
		}
		var_4 := `Hello World!`
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</body></html>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}