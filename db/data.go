package db

// Skill アクター、または敵が使用する技データ
type Skill struct {
	AnimationID int `json:"animationId"`
	Damage      struct {
		Critical  bool   `json:"critical"`
		ElementID int    `json:"elementId"`
		Formula   string `json:"formula"`
		Type      int    `json:"type"`
		Variance  int    `json:"variance"`
	} `json:"damage"`
	Description string `json:"description"`
	Effects     []struct {
		Code   int     `json:"code"`
		DataID int     `json:"dataId"`
		Value1 float64 `json:"value1"`
		Value2 float64 `json:"value2"`
	} `json:"effects"`
	HitType          int    `json:"hitType"`
	IconIndex        int    `json:"iconIndex"`
	ID               int    `json:"id"`
	Message1         string `json:"message1"`
	Message2         string `json:"message2"`
	MpCost           int    `json:"mpCost"`
	Name             string `json:"name"`
	Note             string `json:"note"`
	Occasion         int    `json:"occasion"`
	Repeats          int    `json:"repeats"`
	RequiredWtypeID1 int    `json:"requiredWtypeId1"`
	RequiredWtypeID2 int    `json:"requiredWtypeId2"`
	Scope            int    `json:"scope"`
	Speed            int    `json:"speed"`
	StypeID          int    `json:"stypeId"`
	SuccessRate      int    `json:"successRate"`
	TpCost           int    `json:"tpCost"`
	TpGain           int    `json:"tpGain"`
}

// System システムデータ
type System struct {
	Airship struct {
		Bgm struct {
			Name   string `json:"name"`
			Pan    int    `json:"pan"`
			Pitch  int    `json:"pitch"`
			Volume int    `json:"volume"`
		} `json:"bgm"`
		CharacterIndex int    `json:"characterIndex"`
		CharacterName  string `json:"characterName"`
		StartMapID     int    `json:"startMapId"`
		StartX         int    `json:"startX"`
		StartY         int    `json:"startY"`
	} `json:"airship"`
	ArmorTypes    []string `json:"armorTypes"`
	AttackMotions []struct {
		Type          int `json:"type"`
		WeaponImageID int `json:"weaponImageId"`
	} `json:"attackMotions"`
	BattleBgm struct {
		Name   string `json:"name"`
		Pan    int    `json:"pan"`
		Pitch  int    `json:"pitch"`
		Volume int    `json:"volume"`
	} `json:"battleBgm"`
	Battleback1Name string `json:"battleback1Name"`
	Battleback2Name string `json:"battleback2Name"`
	BattlerHue      int    `json:"battlerHue"`
	BattlerName     string `json:"battlerName"`
	Boat            struct {
		Bgm struct {
			Name   string `json:"name"`
			Pan    int    `json:"pan"`
			Pitch  int    `json:"pitch"`
			Volume int    `json:"volume"`
		} `json:"bgm"`
		CharacterIndex int    `json:"characterIndex"`
		CharacterName  string `json:"characterName"`
		StartMapID     int    `json:"startMapId"`
		StartX         int    `json:"startX"`
		StartY         int    `json:"startY"`
	} `json:"boat"`
	CurrencyUnit string `json:"currencyUnit"`
	DefeatMe     struct {
		Name   string `json:"name"`
		Pan    int    `json:"pan"`
		Pitch  int    `json:"pitch"`
		Volume int    `json:"volume"`
	} `json:"defeatMe"`
	EditMapID  int      `json:"editMapId"`
	Elements   []string `json:"elements"`
	EquipTypes []string `json:"equipTypes"`
	GameTitle  string   `json:"gameTitle"`
	GameoverMe struct {
		Name   string `json:"name"`
		Pan    int    `json:"pan"`
		Pitch  int    `json:"pitch"`
		Volume int    `json:"volume"`
	} `json:"gameoverMe"`
	Locale         string `json:"locale"`
	MagicSkills    []int  `json:"magicSkills"`
	MenuCommands   []bool `json:"menuCommands"`
	OptDisplayTp   bool   `json:"optDisplayTp"`
	OptDrawTitle   bool   `json:"optDrawTitle"`
	OptExtraExp    bool   `json:"optExtraExp"`
	OptFloorDeath  bool   `json:"optFloorDeath"`
	OptFollowers   bool   `json:"optFollowers"`
	OptSideView    bool   `json:"optSideView"`
	OptSlipDeath   bool   `json:"optSlipDeath"`
	OptTransparent bool   `json:"optTransparent"`
	PartyMembers   []int  `json:"partyMembers"`
	Ship           struct {
		Bgm struct {
			Name   string `json:"name"`
			Pan    int    `json:"pan"`
			Pitch  int    `json:"pitch"`
			Volume int    `json:"volume"`
		} `json:"bgm"`
		CharacterIndex int    `json:"characterIndex"`
		CharacterName  string `json:"characterName"`
		StartMapID     int    `json:"startMapId"`
		StartX         int    `json:"startX"`
		StartY         int    `json:"startY"`
	} `json:"ship"`
	SkillTypes []string `json:"skillTypes"`
	Sounds     []struct {
		Name   string `json:"name"`
		Pan    int    `json:"pan"`
		Pitch  int    `json:"pitch"`
		Volume int    `json:"volume"`
	} `json:"sounds"`
	StartMapID int      `json:"startMapId"`
	StartX     int      `json:"startX"`
	StartY     int      `json:"startY"`
	Switches   []string `json:"switches"`
	Terms      struct {
		Basic    []string      `json:"basic"`
		Commands []interface{} `json:"commands"`
		Messages struct {
			ActionFailure   string `json:"actionFailure"`
			ActorDamage     string `json:"actorDamage"`
			ActorDrain      string `json:"actorDrain"`
			ActorGain       string `json:"actorGain"`
			ActorLoss       string `json:"actorLoss"`
			ActorNoDamage   string `json:"actorNoDamage"`
			ActorNoHit      string `json:"actorNoHit"`
			ActorRecovery   string `json:"actorRecovery"`
			AlwaysDash      string `json:"alwaysDash"`
			BgmVolume       string `json:"bgmVolume"`
			BgsVolume       string `json:"bgsVolume"`
			BuffAdd         string `json:"buffAdd"`
			BuffRemove      string `json:"buffRemove"`
			CommandRemember string `json:"commandRemember"`
			CounterAttack   string `json:"counterAttack"`
			CriticalToActor string `json:"criticalToActor"`
			CriticalToEnemy string `json:"criticalToEnemy"`
			DebuffAdd       string `json:"debuffAdd"`
			Defeat          string `json:"defeat"`
			Emerge          string `json:"emerge"`
			EnemyDamage     string `json:"enemyDamage"`
			EnemyDrain      string `json:"enemyDrain"`
			EnemyGain       string `json:"enemyGain"`
			EnemyLoss       string `json:"enemyLoss"`
			EnemyNoDamage   string `json:"enemyNoDamage"`
			EnemyNoHit      string `json:"enemyNoHit"`
			EnemyRecovery   string `json:"enemyRecovery"`
			EscapeFailure   string `json:"escapeFailure"`
			EscapeStart     string `json:"escapeStart"`
			Evasion         string `json:"evasion"`
			ExpNext         string `json:"expNext"`
			ExpTotal        string `json:"expTotal"`
			File            string `json:"file"`
			LevelUp         string `json:"levelUp"`
			LoadMessage     string `json:"loadMessage"`
			MagicEvasion    string `json:"magicEvasion"`
			MagicReflection string `json:"magicReflection"`
			MeVolume        string `json:"meVolume"`
			ObtainExp       string `json:"obtainExp"`
			ObtainGold      string `json:"obtainGold"`
			ObtainItem      string `json:"obtainItem"`
			ObtainSkill     string `json:"obtainSkill"`
			PartyName       string `json:"partyName"`
			Possession      string `json:"possession"`
			Preemptive      string `json:"preemptive"`
			SaveMessage     string `json:"saveMessage"`
			SeVolume        string `json:"seVolume"`
			Substitute      string `json:"substitute"`
			Surprise        string `json:"surprise"`
			UseItem         string `json:"useItem"`
			Victory         string `json:"victory"`
		} `json:"messages"`
		Params []string `json:"params"`
	} `json:"terms"`
	TestBattlers []struct {
		ActorID int   `json:"actorId"`
		Equips  []int `json:"equips"`
		Level   int   `json:"level"`
	} `json:"testBattlers"`
	TestTroopID int    `json:"testTroopId"`
	Title1Name  string `json:"title1Name"`
	Title2Name  string `json:"title2Name"`
	TitleBgm    struct {
		Name   string `json:"name"`
		Pan    int    `json:"pan"`
		Pitch  int    `json:"pitch"`
		Volume int    `json:"volume"`
	} `json:"titleBgm"`
	Variables []string `json:"variables"`
	VersionID int      `json:"versionId"`
	VictoryMe struct {
		Name   string `json:"name"`
		Pan    int    `json:"pan"`
		Pitch  int    `json:"pitch"`
		Volume int    `json:"volume"`
	} `json:"victoryMe"`
	WeaponTypes []string `json:"weaponTypes"`
	WindowTone  []int    `json:"windowTone"`
}

// システムに固定で決まっている文言
type SystemText struct {
	HitTypes    []string
	DamageTypes []string
	Repeats     []string
	Scope       []string
}

// OS言語ごとのシステム文言
var SystemTexts = map[string]SystemText{
	"ja": SystemText{
		HitTypes:    []string{"必中", "物理攻撃", "魔法攻撃"},
		DamageTypes: []string{"なし", "HPダメージ", "MPダメージ", "HP回復", "MP回復", "HP吸収", "MP吸収"},
	},
}
