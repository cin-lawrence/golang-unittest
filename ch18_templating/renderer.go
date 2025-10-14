package blogrenderer

import (
	"embed"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	tmpl     *template.Template
	mdParser *parser.Parser
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostVM(p Post, r *PostRenderer) postViewModel {
	return postViewModel{
		Post:     p,
		HTMLBody: template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil)),
	}
}

func NewPostRenderer() (*PostRenderer, error) {
	tmpl, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &PostRenderer{tmpl: tmpl, mdParser: parser}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.tmpl.ExecuteTemplate(w, "blog.gohtml", newPostVM(p, r))
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.tmpl.ExecuteTemplate(w, "index.gohtml", posts)
}
