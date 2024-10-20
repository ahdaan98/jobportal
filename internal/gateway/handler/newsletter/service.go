package newsletter

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"

	"github.com/ahdaan67/jobportal/config"
)

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
</head>
<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333; max-width: 600px; margin: 0 auto; padding: 20px;">
    <header style="background-color: #00ADD8; color: white; padding: 20px; text-align: center;">
        <h1 style="margin: 0;">{{.Title}}</h1>
        <p style="margin: 5px 0 0;">{{.Subtitle}}</p>
    </header>

    <main style="padding: 20px;">
        <section style="margin-bottom: 30px;">
            <h2 style="color: #00ADD8;">Highlights</h2>
            <ul style="padding-left: 20px;">
                {{range .Highlights}}
                <li>{{.}}</li>
                {{end}}
            </ul>
        </section>

        <section style="margin-bottom: 30px;">
            <h2 style="color: #00ADD8;">Featured Article</h2>
            <h3>{{.FeaturedArticle.Title}}</h3>
            <p>{{.FeaturedArticle.Description}}</p>
            <a href="{{.FeaturedArticle.Link}}" style="color: #00ADD8; text-decoration: none;">Read More →</a>
        </section>
    </main>

    <footer style="background-color: #00ADD8; color: white; padding: 20px; text-align: center;">
        <p>© 2024 {{.Title}}. All rights reserved.</p>
        <p>
            <a href="#" style="color: white; margin: 0 10px;">Unsubscribe</a> |
            <a href="#" style="color: white; margin: 0 10px;">Update Preferences</a> |
            <a href="#" style="color: white; margin: 0 10px;">View Online</a>
        </p>
    </footer>
</body>
</html>
`

type Newsletter struct {
	Title          string `json:"title"`
	Subtitle       string `json:"subtitle"`
	Highlights     []string `json:"highlights"`
	FeaturedArticle struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Link        string `json:"link"`
	} `json:"featured_article"`
}

func SendNewsletter(email string, cfg config.Config, newsletter Newsletter) error {
	from := cfg.Email
	password := cfg.Password
	to := email
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"

	tmpl, err := template.New("newsletter").Parse(htmlTemplate)
	if err != nil {
		log.Fatal("Error parsing template:", err)
		return err
	}

	var renderedHTML bytes.Buffer
	err = tmpl.Execute(&renderedHTML, newsletter)
	if err != nil {
		log.Fatal("Error rendering template:", err)
		return err
	}

	subject := "Subject: " + newsletter.Title + " Newsletter\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := renderedHTML.String()
	message := []byte(subject + mime + body)

	auth := smtp.PlainAuth("", from, password, smtpServer)
	err = smtp.SendMail(fmt.Sprintf("%s:%s", smtpServer, smtpPort), auth, from, []string{to}, message)
	if err != nil {
		log.Println("Error sending newsletter:", err)
		return err
	}

	return nil
}
