package mqservices

import (
//"github.com/JackMaarek/spiderMail/models"
)

func SendCampaignByID(id uint64) {
	var err error
	//var campaign models.Campaign

	//campaign, err = models.FindCampaignByID(id)
	if err != nil {
		panic("Cannot get campaign:  " + err.Error())
	}
	//fmt.Println(campaign)
	//CallMailerService()
}
