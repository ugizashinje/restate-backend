package services

import (
	"log"
	"warrant-api/pkg/config"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/enum"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/matcornic/hermes/v2"
	"github.com/wneessen/go-mail"
)

var h *hermes.Hermes
var mailClient *mail.Client

func initialize() {
	if h == nil {

		h = &hermes.Hermes{
			// Optional Theme
			// Theme: new(Default)
			Product: hermes.Product{
				// Appears in header & footer of e-mails
				Name: "ePutni nalog",
				Link: "http://139.162.145.62:9876/",
				// Optional product logo
				Logo: "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
			},
		}

	}

	if mailClient == nil {

		m, err := mail.NewClient(config.Auth.TsMailHost, mail.WithPort(config.Auth.TsMailPort), mail.WithSMTPAuth(mail.SMTPAuthPlain),
			mail.WithUsername("admin"), mail.WithPassword("admin"))
		if err != nil {
			log.Fatalf("failed to create mail client: %s", err)
		}
		m.SetTLSPolicy(mail.NoTLS)
		mailClient = m
	}
}

type TransactionalEmailServiceImpl struct {
	ConfirmationRepo repo.Repo[model.Confirmation]
}

func (svc *TransactionalEmailServiceImpl) VerifyEmail(g *gin.Context, user *model.User) (*model.Confirmation, error) {
	initialize()
	code := uuid.NewString()
	confirmation := &model.Confirmation{
		UserID: user.ID,
		Code:   code,
		Status: enum.Unconfirmed,
		Url:    config.Auth.WarrantServer + "/auth/confirm/" + code,
	}
	dbRes := svc.ConfirmationRepo.Create(g, confirmation)
	utils.Handle(dbRes.Error)

	email := hermes.Email{
		Body: hermes.Body{
			Greeting: "Postovani",
			Name:     user.FirstName + " " + user.LastName,
			Intros: []string{
				"Dobrodosli na ePutne naloge. Molmo vas potvrdite vasu email adresu",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Da nastavite sa registracijom: klinkite dugme",
					Button: hermes.Button{
						Color: "#22BC66", // Optional action button color
						Text:  "Potvrdi nalog",
						Link:  confirmation.Url,
					},
				},
			},
			Outros: []string{
				"Ako vam je potrena pomoc molim vas pozovite nas na +3815223231321323231.",
			},
		},
	}

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	m := mail.NewMsg()
	if err := m.From("register@eputni.com"); err != nil {
		log.Fatalf("failed to set From address: %s", err)
	}
	if err := m.To(user.Email); err != nil {
		log.Fatalf("failed to set To address: %s", err)
	}
	m.Subject("Potvrdite vas email na portalu putni")
	m.SetBodyString(mail.TypeTextHTML, string(emailBody))

	err = mailClient.DialAndSend(m)
	return confirmation, err
}
