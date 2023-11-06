package widgets

import (
	"log"

	"github.com/Reloyer/rush/dataservice"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func NewGame(gd dataservice.GameData) *gtk.Box {
	mainBox := gtk.NewBox(gtk.OrientationHorizontal, 6)

	b_gameInfo := gtk.NewBox(gtk.OrientationVertical, 2)

	l_gameMode := gtk.NewLabel(gd.GameType)
	l_gameSince := gtk.NewLabel(gd.TimeSince)
	l_gameResult := gtk.NewLabel(gd.GameResult)
	l_gameDuration := gtk.NewLabel(gd.GameDuration)

	l_gameMode.SetHAlign(gtk.AlignStart)
	l_gameSince.SetHAlign(gtk.AlignStart)
	l_gameResult.SetHAlign(gtk.AlignStart)
	l_gameDuration.SetHAlign(gtk.AlignStart)

	b_gameInfo.Append(l_gameMode)
	b_gameInfo.Append(l_gameResult)
	b_gameInfo.Append(l_gameSince)
	b_gameInfo.Append(l_gameDuration)

	b_userOverview := gtk.NewBox(gtk.OrientationVertical, 2)

	b_userInfo := gtk.NewBox(gtk.OrientationHorizontal, 2)
	i_userChampion := gtk.NewImageFromFile(gd.UserChampionIcon)
	i_userChampion.SetVAlign(gtk.AlignStart)
	i_userChampion.SetSizeRequest(66, 66)
	b_userSummonerSpells := gtk.NewBox(gtk.OrientationVertical, 2)

	i_summonerSpell1 := gtk.NewImageFromFile(gd.SummonerSpell1Icon)
	i_summonerSpell2 := gtk.NewImageFromFile(gd.SummonerSpell2Icon)
	i_summonerSpell1.SetSizeRequest(31, 31)
	i_summonerSpell2.SetSizeRequest(31, 31)
	i_summonerSpell1.AddCSSClass("summmonerspell-image")
	i_summonerSpell2.AddCSSClass("summmonerspell-image")
	b_userSummonerSpells.Append(i_summonerSpell1)
	b_userSummonerSpells.Append(i_summonerSpell2)

	b_userPerks := gtk.NewBox(gtk.OrientationVertical, 2)

	i_perk0 := gtk.NewImageFromFile(gd.Perk0Icon)
	i_perkSubStyle := gtk.NewImageFromFile(gd.PerkSubStyleIcon)
	i_perk0.AddCSSClass("perk-image")
	i_perk0.SetSizeRequest(31, 31)
	i_perkSubStyle.SetSizeRequest(31, 31)
	i_perkSubStyle.AddCSSClass("perk-image")
	b_userPerks.Append(i_perk0)
	b_userPerks.Append(i_perkSubStyle)
	b_userPerksnSums := gtk.NewBox(gtk.OrientationHorizontal, 2)
	b_userPerksnSums.Append(b_userPerks)
	b_userPerksnSums.Append(b_userSummonerSpells)

	b_userKDA := gtk.NewBox(gtk.OrientationVertical, 2)

	l_kda := gtk.NewLabel(gd.UserKDA)
	l_ratio := gtk.NewLabel(gd.UserRatio)

	b_userKDA.Append(l_kda)
	b_userKDA.Append(l_ratio)

	b_userInfo.Append(i_userChampion)
	b_userInfo.Append(b_userPerksnSums)
	b_userInfo.Append(b_userKDA)

	b_userBuild := gtk.NewBox(gtk.OrientationHorizontal, 2)

	log.Println("Adding items icons")
	for _, itemicon := range gd.ItemsIcons {
		log.Println("Adding item icon", itemicon)
		i_item := gtk.NewImageFromFile(itemicon)

		i_item.AddCSSClass("item-image")
		i_item.SetSizeRequest(31, 31)
		b_userBuild.Append(i_item)
	} /*
		i_item0 := gtk.NewImageFromFile(ds.ItemIcon0)
		i_item1 := gtk.NewImageFromFile(ds.ItemIcon1)
		i_item2 := gtk.NewImageFromFile(ds.ItemIcon2)
		i_item3 := gtk.NewImageFromFile(ds.ItemIcon3)
		i_item4 := gtk.NewImageFromFile(ds.ItemIcon4)
		i_item5 := gtk.NewImageFromFile(ds.ItemIcon5)
		i_item6 := gtk.NewImageFromFile(ds.ItemIcon6)
		i_item0.SetSizeRequest(31, 31)
		i_item1.SetSizeRequest(31, 31)
		i_item2.SetSizeRequest(31, 31)
		i_item3.SetSizeRequest(31, 31)
		i_item4.SetSizeRequest(31, 31)
		i_item5.SetSizeRequest(31, 31)
		i_item6.SetSizeRequest(31, 31)

		b_userBuild.Append(i_item0)
		b_userBuild.Append(i_item1)
		b_userBuild.Append(i_item2)
		b_userBuild.Append(i_item3)
		b_userBuild.Append(i_item4)
		b_userBuild.Append(i_item5)
		b_userBuild.Append(i_item6)

	*/
	b_userOverview.Append(b_userInfo)
	b_userOverview.Append(b_userBuild)

	b_team1 := gtk.NewBox(gtk.OrientationVertical, 2)
	for _, player := range gd.Team1 {
		b_player := gtk.NewBox(gtk.OrientationHorizontal, 1)
		i_player := gtk.NewImageFromFile(player.ChampionIcon)
		l_player := gtk.NewLabel(player.Nickname)
		b_player.Append(i_player)
		b_player.Append(l_player)
		b_team1.Append(b_player)
	}

	b_team2 := gtk.NewBox(gtk.OrientationVertical, 2)
	for _, enemy := range gd.Team2 {
		b_enemy := gtk.NewBox(gtk.OrientationHorizontal, 1)
		i_enemy := gtk.NewImageFromFile(enemy.ChampionIcon)
		l_enemy := gtk.NewLabel(enemy.Nickname)
		b_enemy.Append(i_enemy)
		b_enemy.Append(l_enemy)
		b_team2.Append(b_enemy)
	}

	b_teams := gtk.NewBox(gtk.OrientationHorizontal, 10)
	b_teams.Append(b_team1)
	b_teams.Append(b_team2)

	mainBox.SetVAlign(gtk.AlignStart)
	mainBox.AddCSSClass("gamebox")
	if gd.GameResult == "VICTORY" {
		mainBox.AddCSSClass("gamebox-won")
	} else if gd.GameResult == "DEFEAT" {
		mainBox.AddCSSClass("gamebox-lost")
	} else {
		mainBox.AddCSSClass("gamebox-unknown")
	}
	if gd.UserKDASuperior {
		l_kda.AddCSSClass("gamebox-kda-superior")
		l_ratio.AddCSSClass("gamebox-kda-superior")
	}
	if gd.UserKDAGood {
		l_kda.AddCSSClass("gamebox-kda-good")
		l_ratio.AddCSSClass("gamebox-kda-good")
	}
	if gd.UserKDANormal {
		l_kda.AddCSSClass("gamebox-kda-normal")
		l_ratio.AddCSSClass("gamebox-kda-normal")
	}
	if gd.UserKDABad {
		l_kda.AddCSSClass("gamebox-kda-bad")
		l_ratio.AddCSSClass("gamebox-kda-bad")
	}
	mainBox.Append(b_gameInfo)
	mainBox.Append(b_userOverview)
	mainBox.Append(b_teams)
	return mainBox
}
