package lotto

const (
	/* ssc */
	SscNumber1          = "10001" // 数字盘 第1球 (范围0～9)
	SscNumber2          = "10002" // 数字盘 第2球 (范围0～9)
	SscNumber3          = "10003" // 数字盘 第3球 (范围0～9)
	SscNumber4          = "10004" // 数字盘 第4球 (范围0～9)
	SscNumber5          = "10005" // 数字盘 第5球 (范围0～9)
	SscTwoSide1         = "10006" // 双面盘 第1球 (大,小,单,双)
	SscTwoSide2         = "10007" // 双面盘 第2球 (大,小,单,双)
	SscTwoSide3         = "10008" // 双面盘 第3球 (大,小,单,双)
	SscTwoSide4         = "10009" // 双面盘 第4球 (大,小,单,双)
	SscTwoSide5         = "10010" // 双面盘 第5球 (大,小,单,双)
	SscTwoSideSum       = "10011" // 双面盘 总和大小单双
	SscExtraFirstThree  = "10012" // 趣味 前三
	SscExtraMiddleThree = "10013" // 趣味 中三
	SscExtraAfterThree  = "10014" // 趣味 后三
	SscNiuNiu           = "10015" // 牛牛
	SscNiuNiuTwoSide    = "10016" // 牛牛 双面

	/* 11选5*/
	SyxwNumber1     = "11001" // 数字盘 第1球
	SyxwNumber2     = "11002" // 数字盘 第2球
	SyxwNumber3     = "11003" // 数字盘 第3球
	SyxwNumber4     = "11004" // 数字盘 第4球
	SyxwNumber5     = "11005" // 数字盘 第5球
	SyxwTwoSide1    = "11006" // 双面盘 第5球
	SyxwTwoSide2    = "11007" // 双面盘 第5球
	SyxwTwoSide3    = "11008" // 双面盘 第5球
	SyxwTwoSide4    = "11009" // 双面盘 第5球
	SyxwTwoSide5    = "11010" // 双面盘 第5球
	SyxwTwoSideSum  = "11011" // 双面盘 总和大小单双龙虎
	SyxwOption1Win1 = "11012" // 任选1中1
	SyxwOption2Win2 = "11013" // 任选2中2
	SyxwOption3Win3 = "11014" // 任选3中3
	SyxwOption4Win4 = "11015" // 任选4中4
	SyxwOption5Win5 = "11016" // 任选5中5
	SyxwOption6Win5 = "11017" // 任选6中5
	SyxwOption7Win5 = "11018" // 任选7中5
	SyxwOption8Win5 = "11019" // 任选8中5
	SyxwGroup2First = "11020" // 前二组选
	SyxwGroup3First = "11021" // 前三组选
	Syxw2First      = "11022" // 前二直选
	Syxw3First      = "11023" // 前三直选

	/* 快3 */
	K3Sum     = "12001" //  点数 由4点至17点，三粒骰子平面点数之总和；
	K33Jun    = "12002" //  三军 任何一粒骰子出现选定之平面点数
	K3TwoSide = "12003" //  大小 三粒骰子平面点数相同，通吃「大」、「小」各注
	K3Weitou  = "12004" //  围骰：三粒骰子平面与选定点数相同；
	K3Quantou = "12005" //  全骰：在一点至六点内，三粒骰子平面点数相同；
	K3Chang   = "12006" //  长牌：任两粒骰子之平面点数
	K3Duan    = "12007" //  短牌：选定两粒骰子之平面点数；

	/* pk10 */
	PK10Number1   = "13001" // 数字盘 冠军
	PK10Number2   = "13002" // 数字盘 亚军
	PK10Number3   = "13003" // 数字盘 第三名
	PK10Number4   = "13004" // 数字盘 第四名
	PK10Number5   = "13005" // 数字盘 第五名
	PK10Number6   = "13006" // 数字盘 第六名
	PK10Number7   = "13007" // 数字盘 第七名
	PK10Number8   = "13008" // 数字盘 第八名
	PK10Number9   = "13009" // 数字盘 第九名
	PK10Number10  = "13010" // 数字盘 第十名
	PK10TwoSide1  = "13011" // 双面盘 冠军 龙虎
	PK10TwoSide2  = "13012" // 双面盘 亚军 龙虎
	PK10TwoSide3  = "13013" // 双面盘 第三名 龙虎
	PK10TwoSide4  = "13014" // 双面盘 第四名 龙虎
	PK10TwoSide5  = "13015" // 双面盘 第五名 龙虎
	PK10TwoSide6  = "13016" // 双面盘 第六名
	PK10TwoSide7  = "13017" // 双面盘 第七名
	PK10TwoSide8  = "13018" // 双面盘 第八名
	PK10TwoSide9  = "13019" // 双面盘 第九名
	PK10TwoSide10 = "13020" // 双面盘 第十名
	PK10TwoSideGY = "13021" // 双面盘 冠亚和
	PK10NumberGY  = "13022" // 数字盘 冠亚和

	/* 3D */
	F3DMaster        = "14001" // 主势盘
	F3DCombo1        = "14002" // 一字组合
	F3DCombo2        = "14003" // 二字组合
	F3DCombo3        = "14004" // 三字组合
	F3DNumberB       = "14005" // 一字定位 百
	F3DNumberS       = "14006" // 一字定位 十
	F3DNumberG       = "14007" // 一字定位 个
	F3DNumberBS      = "14008" // 二字定位 百十
	F3DNumberBG      = "14009" // 二字定位 百个
	F3DNumberSG      = "14010" // 二字定位 十个
	F3DNumberBSG     = "14011" // 三字定位 百十个
	F3DSumBS         = "14012" // 二字合数 百十
	F3DSumBG         = "14013" // 二字合数 百个
	F3DSumSG         = "14014" // 二字合数 十个
	F3DSumBSG        = "14015" // 三字合数 百十个
	F3DSumWeiBS      = "14016" // 二字合数尾 百十
	F3DSumWeiBG      = "14017" // 二字合数尾 百个
	F3DSumWeiSG      = "14018" // 二字合数尾 十个
	F3DSumWeiBSG     = "14019" // 三字合数尾 百十个
	F3DInterval      = "14020" // 跨度
	F3DTwoSideB      = "14021" // 双面盘 百十个 百 (大小单双质合)
	F3DTwoSideS      = "14022" // 双面盘 百十个 十 (大小单双质合)
	F3DTwoSideG      = "14023" // 双面盘 百十个 个 (大小单双质合)
	F3DTwoSideSumBS  = "14024" // 双面盘 合数 百十 (单双)
	F3DTwoSideSumBG  = "14025" // 双面盘 合数 百个 (单双)
	F3DTwoSideSumSG  = "14026" // 双面盘 合数 十个 (单双)
	F3DTwoSideSumBSG = "14027" // 双面盘 合数 百十个 (大小单双)

	/* 幸运农场 */
	HappyFarmNumber1    = "15001" // 数字盘 第一球
	HappyFarmNumber2    = "15002" // 数字盘 第二球
	HappyFarmNumber3    = "15003" // 数字盘 第三球
	HappyFarmNumber4    = "15004" // 数字盘 第四球
	HappyFarmNumber5    = "15005" // 数字盘 第五球
	HappyFarmNumber6    = "15006" // 数字盘 第六球
	HappyFarmNumber7    = "15007" // 数字盘 第七球
	HappyFarmNumber8    = "15008" // 数字盘 第八球
	HappyFarmTwoSide1   = "15009" // 双面盘 第一球
	HappyFarmTwoSide2   = "15010" // 双面盘 第二球
	HappyFarmTwoSide3   = "15011" // 双面盘 第三球
	HappyFarmTwoSide4   = "15012" // 双面盘 第四球
	HappyFarmTwoSide5   = "15013" // 双面盘 第五球
	HappyFarmTwoSide6   = "15014" // 双面盘 第六球
	HappyFarmTwoSide7   = "15015" // 双面盘 第七球
	HappyFarmTwoSide8   = "15016" // 双面盘 第八球
	HappyFarmTwoSideSum = "15017" // 双面盘 总和
	HappyFarmZhenma     = "15018" // 正码

	/* keno */
	KenoSum = "16001" // 双面盘 总和

	/* luck28 */
	Lucky28Number  = "17001" // 数字
	Lucky28Color   = "17002" // 色波
	Lucky28TwoSide = "17003" // 四项 大小单双
	Lucky28Combo   = "17004" // 组合 大单,大双,小单,小双
	Lucky28Extra   = "17005" // 混合 极大,极小,豹子

	/* MarkSix */
	MarkSixNumberA        = "18001" // 特码A
	MarkSixNumberB        = "18002" // 特码B
	MarkSixSpecialTwoSide = "18003" // 特码 双面盘
	MarkSixColor          = "18004" // 色波
	MarkSixHalfColor      = "18005" // 半波
	MarkSixHalfHalfColor  = "18006" // 半半波
	MarkSixSpecialZodiac  = "18007" // 特码生肖
	MarkSixComeZodiac     = "18008" // 合肖
	MarkSixSpecialTou     = "18009" // 特码 头数
	MarkSixSpecialWei     = "18010" // 特码 尾数
	MarkSixZhenMa         = "18011" // 正码
	MarkSixZhen1te        = "18012" // 正一特
	MarkSixZhen2te        = "18013" // 正二特
	MarkSixZhen3te        = "18014" // 正三特
	MarkSixZhen4te        = "18015" // 正四特
	MarkSixZhen5te        = "18016" // 正五特
	MarkSixZhen6te        = "18017" // 正六特
	MarkSixTwoSideZhen1te = "18018" // 正一特 双面盘
	MarkSixTwoSideZhen2te = "18019" // 正二特 双面盘
	MarkSixTwoSideZhen3te = "18020" // 正三特 双面盘
	MarkSixTwoSideZhen4te = "18021" // 正四特 双面盘
	MarkSixTwoSideZhen5te = "18022" // 正五特 双面盘
	MarkSixTwoSideZhen6te = "18023" // 正六特 双面盘
	MarkSixWuxing         = "18024" // 五行
	MarkSixPingTeYiXiao   = "18025" // 平特一肖
	MarkSixPingTeWeiShu   = "18026" // 平特尾数
	MarkSixZhenXiao       = "18027" // 正肖
	MarkSix7SeBo          = "18028" // 7色波
	MarkSixZongXiao       = "18029" // 总肖
	MarkSixZiXuanBuZhong  = "18030" // 自选不中
)
