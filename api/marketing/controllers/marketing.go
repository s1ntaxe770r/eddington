package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
)

type MarketingController struct {
	apiKey string
}

func New(apiKey string, routerGroup *gin.RouterGroup) MarketingController {
	fmt.Println("New MarketingController")

	mc := MarketingController{}

	mc.AddAllControllers(routerGroup)

	return MarketingController{apiKey: apiKey}
}

func (m *MarketingController) AddAllControllers(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/email", m.POSTEmailSubscriber())
}

//	@BasePath	/api/v1/

// CreateUser godoc
//
//	@Summary	Create an user
//	@Schemes
//	@Description	create a user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	Helloworld
//	@Router			/marketing/email [post]
func (m *MarketingController) POSTEmailSubscriber() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		fmt.Println("email: " + email)
		err := m.Addrecipients(email)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "New e-mail added to list successfully!"})
	}
}

// GetIDfromEmail : Get ID from email
/*
{
  "result": {
    "jane_doe@example.com": {
      "contact": {
        "address_line_1": "",
        "address_line_2": "",
        "alternate_emails": [
          "janedoe@example1.com"
        ],
        "city": "",
        "country": "",
        "email": "jane_doe@example.com",
        "first_name": "Jane",
        "id": "asdf-Jkl-zxCvBNm",
        "last_name": "Doe",
        "list_ids": [],
        "segment_ids": [],
        "postal_code": "",
        "state_province_region": "",
        "phone_number": "",
        "whatsapp": "",
        "line": "",
        "facebook": "",
        "unique_name": "",
        "custom_fields": {},
        "created_at": "2021-03-02T15:25:47Z",
        "updated_at": "2021-03-30T15:26:16Z",
        "_metadata": {
          "self": "<metadata_url>"
        }
      }
    },
}
*/
func (m *MarketingController) GetIDfromEmail(email string) (string, error) {
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(m.apiKey, "/v3/marketing/contacts/search/emails", host)
	request.Method = "POST"
	v := EmailList{Emails: []string{email}}
	b, e := json.Marshal(v)
	if e != nil {
		fmt.Println(e)
		return "", e
	}
	request.Body = b

	fmt.Println("request: " + string(request.Body))

	response, err := sendgrid.MakeRequest(request)
	//response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
		return "", err
	}
	var result map[string]interface{}

	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)

	err = json.Unmarshal([]byte(response.Body), &result)
	if err != nil {
		log.Println(err)
		return "", err
	}

	result = result["result"].(map[string]interface{})

	for key, result := range result {
		person := result.(map[string]interface{})
		//just the first e-mail is fine. e-mail is all we have to go off of right now anyhow.
		if key == email {
			return person["id"].(string), nil
		}
	}

	return "", fmt.Errorf("no ID found for email: %s", email)

}

func (m *MarketingController) AddRecipientToWaitingList(recipientID string) error {
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(m.apiKey, fmt.Sprintf("/v3/marketing/contacts/lists/bdd5bf34-a5ba-43a5-b24a-e098b2ae3b68/recipients/%s", recipientID), host)
	request.Method = "POST"
	response, err := sendgrid.MakeRequest(request)
	//response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
		return err
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
		return nil
	}

}

// Addrecipients : Add recipients
// POST /contactdb/recipients
func (u *MarketingController) Addrecipients(email string) error {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/marketing/contacts", host)
	request.Method = "PUT"
	v := RecipientList{Contacts: []Contact{Contact{Email: email}}}
	b, e := json.Marshal(v)
	if e != nil {
		fmt.Println(e)
		return e
	}
	request.Body = b
	response, err := sendgrid.MakeRequest(request)
	if err != nil {
		log.Println(err)
		return err
	} else {
		fmt.Println("Addrecipients worked ------")
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	//TODO just return the id once I figure out how to parse the response
	return nil
}

type RecipientList struct {
	Contacts []Contact `json:"contacts"`
}

type EmailList struct {
	Emails []string `json:"emails"`
}

type Contact struct {
	Email string `json:"email"`
}
