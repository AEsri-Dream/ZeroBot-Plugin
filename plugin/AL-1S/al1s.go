/*
Package atri 本文件基于 https://github.com/Kyomotoi/ATRI
本项目遵守 AGPL v3 协议进行开源
*/
package al1s

import (
	"encoding/base64"
	"math/rand"
	"time"

	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
)

type datagetter func(string, bool) ([]byte, error)

func (dgtr datagetter) randImage(file ...string) message.MessageSegment {
	data, err := dgtr(file[rand.Intn(len(file))], true)
	if err != nil {
		return message.Text("ERROR: ", err)
	}
	return message.ImageBytes(data)
}

func (dgtr datagetter) randRecord(file ...string) message.MessageSegment {
	data, err := dgtr(file[rand.Intn(len(file))], true)
	if err != nil {
		return message.Text("ERROR: ", err)
	}
	return message.Record("base64://" + base64.StdEncoding.EncodeToString(data))
}

func randText(text ...string) message.MessageSegment {
	return message.Text(text[rand.Intn(len(text))])
}


func init() { // 插件主体
	engine := control.AutoRegister(&ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "AL-1S人格文本回复",
		Help: "本插件基于 ATRI ，梦的魔改版\n" +
			"- \n- 喜欢 | 爱你 | 爱 | suki | daisuki | すき | 好き | 贴贴 | 老婆 | 亲一个 | mua\n" +
			"- 草你妈 | 操你妈 | 脑瘫 | 废柴 | fw | 废物 | 战斗 | 爬 | 爪巴 | sb | SB | 傻B\n- 早安 | 早哇 | 早上好 | ohayo | 哦哈哟 | お早う | 早好 | 早 | 早早早\n" +
			"- 中午好 | 午安 | 午好\n- 晚安 | oyasuminasai | おやすみなさい | 晚好 | 晚上好\n- 高性能 | 太棒了 | すごい | sugoi | 斯国一 | よかった\n" +
			"- 没事 | 没关系 | 大丈夫 | 还好 | 不要紧 | 没出大问题 | 没伤到哪\n- 好吗 | 是吗 | 行不行 | 能不能 | 可不可以\n- 啊这\n- 我好了\n- ？ | ? | ¿\n" +
			"- 离谱\n- 答应我",
		PublicDataFolder: "AL-1S",
		OnEnable: func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("嗯呜呜……sensei……？"))
		},
		OnDisable: func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("Zzz……Zzz……"))
		},
	})
	var dgtr datagetter = engine.GetLazyData
	engine.OnFullMatch("萝卜子").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			switch rand.Intn(2) {
			case 0:
				ctx.SendChain(randText("萝卜子是对机器人的蔑称！", "是爱丽丝......萝卜子可是对机器人的蔑称"))
			case 1:
				ctx.SendChain(dgtr.randRecord("RocketPunch.amr"))
			}
		})
	engine.OnFullMatchGroup([]string{"喜欢", "爱你", "爱", "suki", "daisuki", "すき", "好き"}, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(dgtr.randImage("X1.png", "X2.png", "X3.png"))
		})
	engine.OnFullMatchGroup([]string{"讨厌", "讨厌你", "不爱了", "没爱了", "不喜欢你了", "不喜欢", "一边去", "谁问你了", "这机器人好烦", "你好烦"}, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(dgtr.randImage("K1.png", "K2.png", "K3.png", "K4.png"))
		})
	engine.OnFullMatchGroup([]string{"透透", "透透老婆", "开银趴", "炼铜", "萝莉控", "变态", "Hentai", "hentai"}, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(dgtr.randImage("BY1.png", "BY2.png", "HY1.png", "HY2.png", "HY3.png"))
		})
	engine.OnFullMatchGroup([]string{"贴贴", "老婆", "我爱你", "摸摸", "亲一个", "mua"}, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(dgtr.randImage("X3.png"))
		})
	engine.OnFullMatchGroup([]string{"骂我", "狠狠骂我", "tister!", "tister！", "Tister!", "Tister！"}, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(dgtr.randImage("HY2.png", "HY3.png"))
		})
	engine.OnKeywordGroup([]string{"啊米诺斯", "草你妈", "曹尼玛", "神经病", "操你妈", "脑瘫", "废柴", "fw", "kkp", "five", "废物", "战斗", "爬", "爪巴", "sb", "SB", "滚", "傻B"}, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(dgtr.randImage("KB1.png", "KB2.png", "KB3.png", "N1.png", "N2.png", "N3.png"))
		})
	engine.OnFullMatchGroup([]string{"早安", "早哇", "早上好", "ohayo", "哦哈哟", "お早う", "早好", "早", "早早早"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			now := time.Now().Hour()
			switch {
			case now < 6: // 凌晨
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"zzzz......",
					"zzzzzzzz......",
					"zzz...sensei..zzz....",
					"sensei...不要..zzz..那..zzz..",
					"sensei..zzz..呐~..zzzz..",
					"...zzz....哧溜哧溜....",
				))
			case now >= 6 && now < 9:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"啊......早上好...(哈欠)",
					"唔......吧唧...早上...哈啊啊~~~\n早上好......",
					"早上好......",
					"早上好呜......呼啊啊~~~~",
					"啊......早上好。\n昨晚也很激情呢！\n？爱丽丝是说游戏哦",
					"senaei......怎么了...已经早上了么...",
					"早上好！",
					"早上好......欸~~~脸好近呢",
					"邦邦卡邦！午安，sensei！爱丽丝今天精神饱满，准备好迎接新的冒险了！",
				))
			case now >= 9 && now < 18:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"哼！这个点还早啥，昨晚干啥去了！？",
					"熬夜了对吧熬夜了对吧熬夜了对吧？？？！",
					"是不是熬夜是不是熬夜是不是熬夜？！",
				))
			case now >= 18 && now < 24:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"早个啥？哼唧！我都准备洗洗睡了！",
					"不是...你看看几点了，哼！",
					"晚上好哇",
				))
			}
		})
	engine.OnFullMatchGroup([]string{"中午好", "午安", "午好"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			now := time.Now().Hour()
			if now > 11 && now < 15 { // 中午
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"午安w",
					"午觉要好好睡哦，AL-1S会陪伴在你身旁的w",
					"嗯哼哼~睡吧，就像平常一样安眠吧~o(≧▽≦)o",
					"睡你午觉去！哼唧！！",
				))
			}
		})
	engine.OnFullMatchGroup([]string{"晚安", "oyasuminasai", "おやすみなさい", "晚好", "晚上好"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			now := time.Now().Hour()
			switch {
			case now < 6: // 凌晨
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"zzzz......",
					"zzzzzzzz......",
					"zzz...好涩哦..zzz....",
					"别...不要..zzz..那..zzz..",
					"嘻嘻..zzz..呐~..zzzz..",
					"...zzz....哧溜哧溜....",
				))
			case now >= 6 && now < 11:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"sensei可不要猝死！",
					"？都这个时间了！",
					"sensei，快点睡觉！",
				))
			case now >= 11 && now < 15:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"午安w",
					"午觉要好好睡哦，爱丽丝会陪伴在你身旁的w",
					"嗯哼哼~睡吧，就像平常一样安眠吧~o(≧▽≦)o",
				))
			case now >= 15 && now < 19:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"难不成？？晚上不想睡觉？？现在休息",
					"就......挺离谱的...现在睡觉",
					"现在还是白天哦，睡觉还太早了",
				))
			case now >= 19 && now < 24:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"嗯哼哼~睡吧，就像平常一样安眠吧，爱丽丝会一直陪在sensei身边的~o(≧▽≦)o",
					"......(打瞌睡)",
					"呼...呼...爱丽丝已经睡着了哦~...呼......",
					"......爱丽丝会在这守着你的",
				))
			}
		})
	
	engine.OnKeywordGroup([]string{"好吗", "是吗", "行不行", "能不能", "可不可以"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			if rand.Intn(2) == 0 {
				ctx.SendChain(dgtr.randImage("YES1.png", "YES2.png", "NO1.png", "NO2.png"))
			}
		})
	engine.OnKeywordGroup([]string{"啊这", "AZ", "az"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			if rand.Intn(2) == 0 {
				ctx.SendChain(dgtr.randImage("AZ.png", "AZ1.png"))
			}
		})
	engine.OnKeywordGroup([]string{"我好了", "准备好了", "准备完成", "准备出发", "好了", "出发"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(dgtr.randRecord("Aris_Formation_In_2.ogg", "Aris_Formation_Select.ogg", "Aris_Formation_In_1.ogg", "Formation_In_1.ogg", "Formation_In_2.ogg", "Formation_Select.ogg"))
		})
	engine.OnFullMatchGroup([]string{"？", "?", "¿", "离谱"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			switch rand.Intn(5) {
			case 0:
				ctx.SendChain(randText("?", "？", "嗯？", "(。´・ω・)ん?", "ん？"))
			case 1, 2:
				ctx.SendChain(dgtr.randImage("WH1.png", "WH2.png"))
			}
		})
	engine.OnKeyword("答应我", zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(randText("爱丽丝会努力的！"))
		})
}
