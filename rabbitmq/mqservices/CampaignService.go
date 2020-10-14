package mqservices

import (
	"fmt"

	"github.com/JackMaarek/spiderMail/models"
)

func GatherCampaignDataByID(id uint64) error {
	// fmt.Println("test Gather")
	var err error
	var campaign models.Campaign
	campaign, err = models.FindCampaignByID(id)
	if err != nil {
		fmt.Println("Cannot get campaign with error:" + err.Error())
		return err
	}
	// fmt.Println(campaign.ID)
	// fmt.Println(campaign.RecipientsListId)
	var recipientList *[]models.Recipient
	recipientList, err = models.FindRecipientsByListId(uint32(campaign.RecipientsListId))

	if err != nil {
		return err
	}
	fmt.Println("*recipentList",*recipientList)
	var recipient models.Recipient
	for _, recipient = range *recipientList {
		mailData := Mail{
			Recipient: recipient.Email,
			Subject:   campaign.Subject,
			Body:      campaign.Content,
		}
		fmt.Println("Mail data :",mailData)
		err := CallMailerService(&mailData)
		if err != nil {
			fmt.Println("Cannot send email to %d", recipient.ID)
		}
	}

	return nil
}
