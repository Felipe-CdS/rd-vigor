// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package landing_page

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "nugu.dev/rd-vigor/views/layout"

func LandingIndex(title string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex items-center py-4 px-8 h-14 border-b border-gray-200 border-solid\"><a class=\"flex items-center mr-3 w-28 sm:mr-4\"><img class=\"w-full\" src=\"static/img/logo-full.png\"></a><div class=\"flex items-center space-x-3 h-8\"><a class=\"p-1 text-[#27272A] cursor-pointer\">Home</a> <a class=\"p-1 text-[#71717A] cursor-pointer\">Eventos</a> <a class=\"p-1 text-[#71717A] cursor-pointer\">Reuniões</a></div><div class=\"flex items-center ml-auto space-x-4\"><button class=\"text-sm px-4 py-2 rounded-full bg-[#27272A] text-white font-medium\"><span class=\"sm:block\">Login</span></button></div></div><div class=\"flex flex-col justify-center items-center grow\"><div class=\"flex relative flex-col w-3/5 h-96 rounded-3xl border border-gray-200 border-solid card-shadow\" style=\"background-image: linear-gradient(to right top, #ffdfa9, #fed18c, #fec371, #feb455, #fea439);\"><div class=\"flex absolute -bottom-7 left-1/2 py-1 px-2 w-4/5 bg-white rounded-3xl border border-gray-200 border-solid -translate-x-1/2 card-shadow\"><div class=\"relative mr-3 w-10 h-10\"><img class=\"w-full\" src=\"static/img/search.svg\"></div><input class=\"flex-grow w-full\" type=\"text\" placeholder=\"Procure por uma ocupação ou empresa especifica\"> <button class=\"ml-3\" type=\"button\">aaaaa</button></div></div></div><div class=\"flex justify-center py-4 px-8 border-t border-gray-200 border-solid\"><span class=\"text-[#71717A]\">© RD Vigor 2024</span></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Base(title).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
