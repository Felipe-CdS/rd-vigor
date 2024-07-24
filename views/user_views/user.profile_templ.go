// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package user_views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"nugu.dev/rd-vigor/repositories"
	"nugu.dev/rd-vigor/views/layout"
)

func UserProfile(title string, user repositories.User) templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"{open: false}\"><div class=\"flex items-center py-4 px-8 h-14 border-b border-gray-200 border-solid\"><a class=\"flex items-center mr-3 w-28 sm:mr-4\"><img class=\"w-full\" src=\"/static/img/logo-full-2.png\"></a> <button class=\"ml-auto lg:hidden size-8\" x-on:click=\"open = !open; document.body.classList.add(&#39;no-scroll&#39;);\"><img class=\"w-full\" src=\"/static/img/menu.svg\"></button></div><div class=\"flex relative w-screen h-[calc(100vh-3.5rem)]\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = layout.DesktopNavbar().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = layout.MobileNavbar(user).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"overflow-x-hidden flex flex-col py-5 px-5 lg:px-32 2xl:px-64 w-full lg:w-[calc(100vw-18rem)] h-full overflow-y-scroll bg-[#F0F2F5] items-center\"><div class=\"flex flex-col w-full text-left rounded-t-md border border-b-0 border-gray-200 border-solid min-h-32 card-shadow\" style=\"background-image: linear-gradient(to right top, #ffdfa9, #fed18c, #fec371, #feb455, #fea439);\"></div><div class=\"flex relative flex-col p-4 w-full text-left bg-white rounded-b-md border border-gray-200 border-solid card-shadow\"><div class=\"flex absolute left-0 -top-20 flex-col p-4 mr-auto w-fit\"><img class=\"bg-white rounded-3xl size-20\" src=\"/static/img/musk.jpeg\"><h1 class=\"mt-1 text-2xl font-semibold\">Elon Musk</h1><a class=\"font-medium text-sm text-[#ed6c09] hover:underline\" href=\"https://www.spacex.com\">Space X </a><h1 class=\"text-xs font-semibold text-gray-400\">Los Angeles, CA</h1></div><button class=\"px-4 py-2 rounded-full w-fit ml-auto bg-[#FFBD59] text-white font-semibold text-sm\"><span class=\"sm:block\">Mensagem</span></button><div class=\"flex items-center mt-12 space-x-3 w-full\"><div class=\"text-sm font-medium text-gray-400 w-fit\"><span class=\"mt-1 text-base font-semibold text-black\">0</span> Contratos</div><div class=\"text-sm font-medium text-gray-400 w-fit\"><span class=\"mt-1 text-base font-semibold text-black\">0</span> Portifólios</div><div class=\"text-sm font-medium text-gray-400 w-fit\"><span class=\"mt-1 text-base font-semibold text-black\">0</span> Avaliações</div></div></div><div class=\"flex relative flex-col py-4 px-4 mt-2 w-full text-left bg-white rounded-md border border-gray-200 border-solid card-shadow\"><div class=\"flex w-full text-sm font-medium text-gray-400\"><span class=\"font-medium\">Infomações Básicas</span></div></div><div class=\"flex flex-col mt-2 w-full lg:flex-row\"><div class=\"flex relative flex-col py-4 px-4 w-full text-left bg-white rounded-md border border-gray-200 border-solid lg:mr-2 card-shadow\"><div class=\"flex w-full text-sm font-medium text-gray-400\"><span class=\"font-medium\">Sobre Elon</span></div><p class=\"mt-1 text-sm text-justify\">Sou Elon Musk, fundador e CEO da SpaceX, uma empresa que está revolucionando a exploração espacial. Com sede na Califórnia, nossa missão na SpaceX é tornar a humanidade multiplanetária, reduzindo os custos de acesso ao espaço e aumentando a capacidade de transporte espacial. Desde o lançamento bem-sucedido do Falcon 1 em 2008 até as missões internacionais da Crew Dragon para a Estação Espacial Internacional, estamos redefinindo o que é possível no espaço.</p></div><div class=\"flex relative flex-col p-4 text-left bg-white rounded-md border border-gray-200 border-solid w-fit card-shadow\"><div class=\"flex w-full text-sm font-medium text-gray-400\"><span class=\"mb-2 font-medium\">Video de Apresentação</span></div><iframe width=\"560\" height=\"315\" src=\"https://www.youtube.com/embed/zIwLWfaAg-8?si=PwkFascoJp_bYdU9\" title=\"YouTube video player\" frameborder=\"0\" referrerpolicy=\"strict-origin-when-cross-origin\" allowfullscreen></iframe></div></div><div class=\"flex relative flex-col p-4 mt-2 w-full text-left bg-white rounded-md border border-gray-200 border-solid card-shadow\"><div class=\"flex w-full text-sm font-medium text-gray-400\"><span class=\"font-medium\">Avaliações</span> <button class=\"flex items-center ml-auto font-medium\">Ver todas <img class=\"ml-1 size-3\" src=\"/static/img/arrow-right.svg\"></button></div><p class=\"mt-3 text-sm font-medium text-center text-gray-400\">Nenhuma avaliação por enquanto...</p></div><div class=\"flex relative flex-col p-4 mt-2 w-full text-left bg-white rounded-md border border-gray-200 border-solid card-shadow\"><div class=\"flex w-full text-sm font-medium text-gray-400\"><span class=\"font-medium\">Portifólios</span> <button class=\"flex items-center ml-auto font-medium\">Ver todos <img class=\"ml-1 size-3\" src=\"/static/img/arrow-right.svg\"></button></div><p class=\"mt-3 text-sm font-medium text-center text-gray-400\">Nenhum portifólio por enquanto...</p></div></div></div></div>")
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
