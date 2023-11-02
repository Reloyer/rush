package dataservice

import (
	"fmt"
	"strings"

	"github.com/Reloyer/rush/lcu/datatypes"
	"github.com/Reloyer/rush/utility"
)

type HomeData struct {
	Nickname     string
	ProfileIcon  string
	SoloqIcon    string
	SoloqTier    string
	SoloqDivison string
	FlexqIcon    string
	FlexqTier    string
	FlexqDivison string
}
type DataService struct {
	Homedata HomeData
}

func NewDataService() *DataService {
	return &DataService{}
}
func (ds *DataService) GetHomePageData(s datatypes.Summoner, r datatypes.Ranked) {
	assetsDir := "./assets/"
	ds.Homedata.Nickname = s.DisplayName
	ds.Homedata.ProfileIcon = fmt.Sprintf("%sprofileIcons/%d.jpg", assetsDir, s.ProfileIconId)
	ds.Homedata.SoloqIcon = fmt.Sprintf("%srankedIcons/%s%d.png", assetsDir, strings.ToLower(r.Queues[0].Tier), utility.DivisonToDec(r.Queues[0].Division))
	ds.Homedata.SoloqTier = r.Queues[0].Tier
	ds.Homedata.SoloqTier = r.Queues[0].Division
	ds.Homedata.FlexqIcon = fmt.Sprintf("%srankedIcons/%s%d.png", assetsDir, strings.ToLower(r.Queues[1].Tier), utility.DivisonToDec(r.Queues[1].Division))
	ds.Homedata.FlexqTier = r.Queues[1].Tier
	ds.Homedata.FlexqTier = r.Queues[1].Division
}
