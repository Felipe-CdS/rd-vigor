// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package admin_views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"fmt"
	"nugu.dev/rd-vigor/repositories"
	"nugu.dev/rd-vigor/views/layout"
	"strconv"
	"time"
)

func Base(title string, users []repositories.User) templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex items-center py-4 px-8 h-14 border-b border-gray-200 border-solid\"><a class=\"flex items-center mr-3 w-28 sm:mr-4\"><img class=\"w-full\" src=\"/static/img/logo-full-2.png\"></a><div class=\"flex items-center space-x-3 h-8\"><a class=\"p-1 text-[#71717A] cursor-pointer\">Home</a> <a class=\"p-1 text-[#71717A] cursor-pointer\">Eventos</a> <a class=\"p-1 text-[#71717A] cursor-pointer\">Reuniões</a> <a class=\"p-1 text-[#27272A] cursor-pointer\">Dashboard</a></div><div class=\"flex items-center ml-auto space-x-4\"><button hx-get=\"/logout\" hx-target=\"body\" hx-push-url=\"true\" class=\"px-4 py-2 rounded-full bg-[#27272A] text-white font-medium text-sm\"><span class=\"sm:block\">Logout</span></button></div></div><div class=\"flex py-5 px-16 grow justify-center bg-[#F0F2F5] w-screen h-[calc(100vh-7rem)] items-center\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = UsersList(users).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"user-info-div\" class=\"flex flex-col justify-center p-4 w-2/3 h-full text-left bg-white rounded-md border border-gray-200 border-solid card-shadow\"><div class=\"flex flex-col items-center p-4 w-full\"><img class=\"size-16\" src=\"/static/img/user.svg\"><h1 class=\"mt-1 text-2xl font-medium text-gray-400\">Nenhum usuário selecionado</h1><h1 class=\"text-sm font-normal text-gray-400\">Selecione um usuário da lista para editá-lo</h1></div></div></div>")
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

func UsersList(users []repositories.User) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col mr-2 w-1/3 h-full text-left bg-white rounded-md border border-gray-200 border-solid card-shadow\"><div class=\"flex items-center py-3 px-4 w-full bg-white rounded-t-md border-b border-gray-200 border-solid\"><h1 class=\"font-medium\">Usuários (")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(len(users)))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/admin_views/dashboard/base.templ`, Line: 55, Col: 64}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(")</h1><button class=\"ml-auto\" title=\"filtrar\"><img class=\"size-6\" src=\"/static/img/filter.svg\"></button></div><div x-data=\"{selected: &#39;&#39;}\" class=\"flex overflow-y-scroll relative flex-col w-full h-full\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(users) == 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col justify-center items-center w-full h-[calc(100%-2.25rem)] text-center\"><img class=\"size-12\" src=\"/static/img/warning.svg\"><h1 class=\"mt-1 text-2xl font-medium text-gray-400\">Nenhum usuário encontrado</h1><h1 class=\"text-sm font-normal text-gray-400\">Consulte o banco de dados se isso for um erro</h1></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			for _, usr := range users {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button hx-get=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("dashboard/details?user=%s", usr.ID))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/admin_views/dashboard/base.templ`, Line: 73, Col: 63}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#user-info-div\" hx-target-3xx=\"body\" hx-swap=\"outerHTML\" class=\"flex flex-col items-start p-4 w-full border-b border-gray-200 border-solid\" :id=\"$id(&#39;user&#39;)\" x-on:click=\"selected = $el.id\" x-bind:style=\"$el.id == selected &amp;&amp; { &#39;border-left&#39;: &#39;4px solid #ED6C09&#39; }\"><h1 class=\"text-sm font-medium\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(usr.FirstName)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/admin_views/dashboard/base.templ`, Line: 82, Col: 53}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(usr.LastName)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/admin_views/dashboard/base.templ`, Line: 82, Col: 70}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1><h1 class=\"text-sm font-normal text-gray-400\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var8 string
				templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("@%s", usr.Username))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/admin_views/dashboard/base.templ`, Line: 83, Col: 86}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1></button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func UserInfoDiv(user repositories.User) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var9 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var9 == nil {
			templ_7745c5c3_Var9 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"user-info-div\" class=\"flex flex-col p-4 w-2/3 h-full text-left bg-white rounded-md border border-gray-200 border-solid card-shadow\"><div class=\"flex flex-col justify-items-start items-start py-2 px-4 w-full\"><p class=\"mr-auto text-2xl font-semibold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(user.FirstName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/admin_views/dashboard/base.templ`, Line: 98, Col: 20}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(user.LastName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/admin_views/dashboard/base.templ`, Line: 99, Col: 19}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><hr class=\"my-2 w-full h-px bg-gray-200 border-0\"><h1 class=\"text-gray-500 text-md\">Informações básicas</h1><div class=\"flex mt-1 space-x-3 w-full\"><label class=\"flex flex-col space-y-1 w-1/2 text-sm text-gray-400\" for=\"email\">Email <input disabled name=\"email\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(user.Email)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/admin_views/dashboard/base.templ`, Line: 112, Col: 24}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"py-1 px-2 mt-1 font-normal text-gray-800 bg-gray-200 rounded-md\"></label> <label class=\"flex flex-col w-1/2 text-sm text-gray-400\" for=\"telephone\">Telefone <input disabled name=\"telephone\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var13 string
		templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(user.Telephone)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/admin_views/dashboard/base.templ`, Line: 124, Col: 28}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"py-1 px-2 mt-1 font-normal text-gray-800 bg-gray-200 rounded-md\"></label></div><div class=\"flex mt-1 space-x-3 w-full\"><label class=\"flex flex-col space-y-1 w-1/2 text-sm text-gray-400\" for=\"occupation_area\">Area de Atuação <input disabled name=\"occupation_area\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var14 string
		templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(user.OccupationArea)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/admin_views/dashboard/base.templ`, Line: 138, Col: 33}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"py-1 px-2 mt-1 font-normal text-gray-800 bg-gray-200 rounded-md\"></label> <label class=\"flex flex-col w-1/2 text-sm text-gray-400\" for=\"refer_friend\">Indicado Por ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if user.ReferFriend != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input disabled name=\"refer_friend\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var15 string
			templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(user.ReferFriend)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/admin_views/dashboard/base.templ`, Line: 151, Col: 31}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"py-1 px-2 mt-1 font-normal text-gray-800 bg-gray-200 rounded-md\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input disabled name=\"refer_friend\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var16 string
			templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs("Sem Indicação")
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/admin_views/dashboard/base.templ`, Line: 158, Col: 32}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"py-1 px-2 mt-1 font-normal text-gray-800 bg-gray-200 rounded-md\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</label></div><div class=\"flex mt-1 space-x-3 w-full\"><label class=\"flex flex-col space-y-1 w-1/2 text-sm text-gray-400\" for=\"created_at\">Data de Criação <input disabled name=\"created_at\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var17 string
		templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(user.CreatedAt.Format(time.RFC822Z))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/admin_views/dashboard/base.templ`, Line: 173, Col: 49}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"py-1 px-2 mt-1 font-normal text-gray-800 bg-gray-200 rounded-md\"></label> <label class=\"flex flex-col w-1/2 text-sm text-gray-400\" for=\"updated_at\">Último Acesso ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !user.UpdatedAt.IsZero() {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input disabled name=\"updated_at\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var18 string
			templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(user.UpdatedAt.Format(time.RFC822Z))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/admin_views/dashboard/base.templ`, Line: 186, Col: 50}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"py-1 px-2 mt-1 font-normal text-gray-800 bg-gray-200 rounded-md\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input disabled name=\"updated_at\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var19 string
			templ_7745c5c3_Var19, templ_7745c5c3_Err = templ.JoinStringErrs("Nenhum acesso registrado")
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/admin_views/dashboard/base.templ`, Line: 193, Col: 41}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var19))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"py-1 px-2 mt-1 font-normal text-gray-800 bg-gray-200 rounded-md\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</label></div><div class=\"flex my-1 space-x-3 w-full\"><label class=\"flex flex-col space-y-1 w-1/2 text-sm text-gray-400\" for=\"document_file\">Documento Anexado <input disabled name=\"document_file\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var20 string
		templ_7745c5c3_Var20, templ_7745c5c3_Err = templ.JoinStringErrs("Sem Documentação")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/admin_views/dashboard/base.templ`, Line: 208, Col: 34}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var20))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"py-1 px-2 mt-1 font-normal text-gray-800 bg-gray-200 rounded-md\"></label></div><hr class=\"my-2 w-full h-px bg-gray-200 border-0\"><h1 class=\"text-gray-500 text-md\">Ações</h1><div class=\"flex mt-1 space-x-3 w-full\"><form hx-post=\"\" class=\"flex flex-col space-y-1 w-1/2 text-sm text-gray-400\"><label>Alterar situação cadastral</label><div class=\"flex mt-1 space-x-2 w-full\"><select x-data=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var21 string
		templ_7745c5c3_Var21, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("{ selected: '%s' }", user.RegistrationStatus))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/admin_views/dashboard/base.templ`, Line: 225, Col: 74}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var21))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"flex py-1 px-2 font-normal text-gray-800 bg-gray-200 rounded-md grow\"><option value=\"rejected\" :selected=\"selected === &#39;rejected&#39;\">Rejeitado</option> <option value=\"pending\" :selected=\"selected === &#39;pending&#39;\">Pendente</option> <option value=\"accepted\" :selected=\"selected === &#39;accepted&#39;\">Aceito</option></select> <button class=\"w-1/5 rounded-md text-white bg-[#FC8713]\">Salvar</button></div></form></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
