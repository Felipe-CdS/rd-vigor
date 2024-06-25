// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package login

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func SignupForm() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex justify-between w-full\"><div><h1 class=\"text-2xl font-bold\">Pré-cadastro</h1><h1 class=\"text-sm text-gray-400\">Crie a sua conta agora e tenha a oportunidade de ser um dos primeiros na RD Vigor!</h1></div><button hx-get=\"/login\" hx-target=\"#form-box\" class=\"size-10\"><img class=\"w-full\" src=\"static/img/close.svg\"></button></div><hr class=\"w-full h-px bg-gray-200 border-0\"><form hx-post=\"/signup\" hx-target=\"#form-box\" hx-encoding=\"multipart/form-data\" class=\"flex flex-col space-y-3 w-full\"><div class=\"flex justify-between items-center space-x-3 w-full\"><input type=\"text\" name=\"first_name\" placeholder=\"Primeiro Nome\" class=\"py-3 px-4 w-full font-normal rounded-md border border-gray-200 border-solid\"> <input type=\"text\" name=\"last_name\" placeholder=\"Sobrenome\" class=\"py-3 px-4 w-full font-normal rounded-md border border-gray-200 border-solid\"></div><input type=\"email\" name=\"email\" placeholder=\"E-mail\" class=\"py-3 px-4 w-full font-normal rounded-md border border-gray-200 border-solid\"> <input type=\"password\" name=\"password\" placeholder=\"Nova senha\" class=\"py-3 px-4 w-full font-normal rounded-md border border-gray-200 border-solid\"> <input type=\"password\" name=\"repeat-password\" placeholder=\"Repita a senha\" class=\"py-3 px-4 w-full font-normal rounded-md border border-gray-200 border-solid\"> <input type=\"text\" name=\"occupation_area\" placeholder=\"Area de atuação\" class=\"py-3 px-4 w-full font-normal rounded-md border border-gray-200 border-solid\"> <label class=\"text-sm text-justify text-gray-400\">Para finalizar, precisamos de um documento ou imagem que comprove a sua área de atuação. Pedimos isso para assegurar a qualidade e veracidade dos usuários da RD Vigor, mantendo assim o nosso compromisso de oferecer apenas os melhores profissionais para a sua empresa.</label> <input type=\"file\" class=\"py-3 px-4 w-full font-normal rounded-md border-2 border-gray-200 border-dashed\"> <button type=\"submit\" class=\"p-3 w-full font-semibold text-lg rounded-md text-white bg-[#FFBD59]\">Criar Conta</button></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func SignupFormDone() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col justify-between items-center w-full\"><div><h1 class=\"text-center text-gray-400 text-md\">Obrigado por demonstrar interesse na nossa plataforma! Seu perfil será analisado e em breve enviaremos um email para você com todas as informações que precisa.</h1></div><a hx-get=\"/\" hx-target=\"body\" class=\"text-md mt-3 text-[#FFBD59] cursor-pointer font-normal\">Voltar</a></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func NotReady() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col justify-between items-center w-full\"><div><h1 class=\"text-center text-gray-400 text-md\">Obrigado por se cadastrar na RD Vigor. No momento ainda estamos dando os últimos retoques na plataforma para que você tenha a melhor experiência possivel. Assim que tudo estiver pronto enviaremos um e-mail para você. Nos vemos em breve!</h1></div><a hx-get=\"/\" hx-target=\"body\" class=\"text-md mt-3 text-[#FFBD59] cursor-pointer font-normal\">Voltar</a></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func LoginForm() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form class=\"flex flex-col space-y-3 w-full\"><input class=\"py-3 px-4 w-full font-normal rounded-md border border-gray-200 border-solid\" placeholder=\"E-mail\"> <input class=\"py-3 px-4 w-full font-normal rounded-md border border-gray-200 border-solid\" placeholder=\"Senha\"> <button hx-get=\"/not-ready\" hx-target=\"#form-box\" class=\"p-3 w-full text-lg font-semibold rounded-md text-white bg-[#FFBD59]\">Entrar</button></form><a class=\"text-sm text-[#FFBD59] cursor-pointer font-normal\">Esqueceu a senha?</a><hr class=\"w-full h-px bg-gray-200 border-0\"><button id=\"signup-button\" hx-get=\"/signup\" hx-target=\"#form-box\" class=\"p-3 w-full font-semibold text-lg rounded-md text-white bg-[#441a06]\">Criar nova conta</button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
