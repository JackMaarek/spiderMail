package mqservices

import (
	"fmt"
	"github.com/JackMaarek/spiderMail/models"
)

func SendCampaignByID(id uint64) {
	var err error
	var campaign models.Campaign


	campaign, err = models.FindCampaignByID(id)
	if err != nil {
		panic("Cannot get campaign with error:" + err.Error())
	}
	fmt.Println(campaign)

	var recipientList models.RecipientsList
	fmt.Println(recipientList)
	//recipientList, err = models.FindRecipientsListByID(campaign.RecipientsListId)
	//fmt.Println(recipientList.)
	//for recipient := range RecipientsList {



	//CallMailerService()
}
