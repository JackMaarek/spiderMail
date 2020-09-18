package routines

import (
	"github.com/JackMaarek/spiderMail/services/rabbitmq/producer"
	"fmt"
	"github.com/JackMaarek/spiderMail/models"
	"time"
)

// CheckForCampaignsToSend checks every interval(minutes) for campaigns to send to rabbitmq
func CheckForCampaignsToSend(interval int) {
	// Infinite loop
	for true {
		fmt.Println("Checking for campaigns...")
		var campaignIds []uint64
		
		campaignIds = models.GetCampaignsToSend()
		if len(campaignIds) == 0 {
			fmt.Println("No campaign to send!")
		}
		for _, id := range campaignIds {
			var err error
			err = producer.SendToRabbit(id)
			
			if err != nil {
				fmt.Println("Error while sending campaign n°", id ,": ", err)
			} else {
				// If message correctly sent to the rabbitmq, update campaign to done
				campaign, _ := models.FindCampaignByID(id)
				campaign.IsDone = true
				models.EditCampaignByID(&campaign, id)
			}
		}

		fmt.Println("Check finished!")
		// Sleep N minutes before checking again for campaigns
		time.Sleep(time.Duration(interval) * time.Minute)
	}
}