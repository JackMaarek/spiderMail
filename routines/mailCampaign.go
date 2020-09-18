package routines

import (
	"github.com/JackMaarek/spiderMail/services/rabbitmq/producer"
	"fmt"
	"github.com/JackMaarek/spiderMail/models"
	"time"
)

func CheckForCampaignsToSend() {
	// Infinite loop
	for true {
		var campaignIds []uint64

		campaignIds = models.GetCampaignsToSend()

		for _, id := range campaignIds {
			var err error
			err = producer.SendToRabbit(id)
			
			if err != nil {
				fmt.Println("Error while sending campaign: ", err)
			} else {
				// If message correctly sent to the rabbitmq, update campaign to done
				campaign, _ := models.FindCampaignByID(id)
				campaign.IsDone = true
				models.EditCampaignByID(&campaign, id)
			}
		}

		// Sleep 2 minutes before checking again for campaigns
		time.Sleep(2 * time.Minute)
	}
}