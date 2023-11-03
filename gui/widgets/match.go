package widgets

import (
	"github.com/Reloyer/rush/dataservice"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func NewGame(ds dataservice.HomeData) *gtk.Box {
	mainBox := gtk.NewBox(gtk.OrientationHorizontal, 6)

	b_gameInfo := gtk.NewBox(gtk.OrientationVertical, 2)

	l_gameMode := gtk.NewLabel("GameMode")
	l_gameDate := gtk.NewLabel("GameDate")
	l_gameResult := gtk.NewLabel("GameResult")
	l_gameDuration := gtk.NewLabel("GameDuration")

	b_gameInfo.Append(l_gameMode)
	b_gameInfo.Append(l_gameDate)
	b_gameInfo.Append(l_gameResult)
	b_gameInfo.Append(l_gameDuration)

	b_userOverview := gtk.NewBox(gtk.OrientationHorizontal, 2)

	b_userInfo := gtk.NewBox(gtk.OrientationVertical, 2)
	i_userChampion := gtk.NewImage()

	b_userSummonerSpells := gtk.NewBox(gtk.OrientationVertical, 2)

	i_stSpell := gtk.NewImage()
	i_ndSpell := gtk.NewImage()

	b_userSummonerSpells.Append(i_stSpell)
	b_userSummonerSpells.Append(i_ndSpell)

	b_userRunes := gtk.NewBox(gtk.OrientationVertical, 2)

	i_stRune := gtk.NewImage()
	i_ndRune := gtk.NewImage()

	b_userRunes.Append(i_stRune)
	b_userRunes.Append(i_ndRune)

	b_userKDA := gtk.NewBox(gtk.OrientationVertical, 2)

	l_kda := gtk.NewLabel("KDA")
	l_ratio := gtk.NewLabel("ratio")

	b_userKDA.Append(l_kda)
	b_userKDA.Append(l_ratio)

	b_userInfo.Append(i_userChampion)
	b_userInfo.Append(b_userSummonerSpells)
	b_userInfo.Append(b_userRunes)
	b_userInfo.Append(b_userKDA)

	b_userBuild := gtk.NewBox(gtk.OrientationHorizontal, 1)

	i_item1 := gtk.NewImage()
	i_item2 := gtk.NewImage()
	i_item3 := gtk.NewImage()
	i_item4 := gtk.NewImage()
	i_item5 := gtk.NewImage()
	i_item6 := gtk.NewImage()
	i_trinket := gtk.NewImage()

	b_userBuild.Append(i_item1)
	b_userBuild.Append(i_item2)
	b_userBuild.Append(i_item3)
	b_userBuild.Append(i_item4)
	b_userBuild.Append(i_item5)
	b_userBuild.Append(i_item6)
	b_userBuild.Append(i_trinket)

	b_userOverview.Append(b_userInfo)
	b_userOverview.Append(b_userBuild)

	mainBox.Append(b_gameInfo)
	mainBox.Append(b_userOverview)

	return mainBox
}
