package routines

import (
	"github.com/JackMaarek/spiderMail/services/rabbitmq/producer"
	"github.com/JackMaarek/spiderMail/models"
	"time"
	log "github.com/sirupsen/logrus"
)

// CheckForCampaignsToSend checks every interval(minutes) for campaigns to send to rabbitmq
func CheckForCampaignsToSend(interval int) {
	// Infinite loop
	for true {
		log.Info("Checking for campaigns to send...")
		var campaignIds []uint64
		campaignIds = models.GetCampaignsToSend()
		if len(campaignIds) != 0 {
			log.WithField("number", len(campaignIds)).Info("Found campaigns that need to be sent")
		} else {
			log.Info("Nothing to send")
		}

		for _, id := range campaignIds {
			var err error
			err = producer.SendToRabbit(id)
			
			if err != nil {
				log.WithField("campaign id", id).Warn("Error while sending campaign: ", err)
			} else {
				// If message correctly sent to the rabbitmq, update campaign to done
				campaign, _ := models.FindCampaignByID(id)
				campaign.IsDone = true
				models.EditCampaignByID(&campaign, id)
			}
		}

		// Sleep N minutes before checking again for campaigns
		time.Sleep(time.Duration(interval) * time.Minute)
	}
}