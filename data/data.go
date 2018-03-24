package data

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
		Code   int64   `json:"code"`
		DataID int64   `json:"dataId"`
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

type System struct {
	Airship struct {
		Bgm struct {
			Name   string `json:"name"`
			Pan    int64  `json:"pan"`
			Pitch  int64  `json:"pitch"`
			Volume int64  `json:"volume"`
		} `json:"bgm"`
		CharacterIndex int64  `json:"characterIndex"`
		CharacterName  string `json:"characterName"`
		StartMapID     int64  `json:"startMapId"`
		StartX         int64  `json:"startX"`
		StartY         int64  `json:"startY"`
	} `json:"airship"`
	ArmorTypes    []string `json:"armorTypes"`
	AttackMotions []struct {
		Type          int64 `json:"type"`
		WeaponImageID int64 `json:"weaponImageId"`
	} `json:"attackMotions"`
	BattleBgm struct {
		Name   string `json:"name"`
		Pan    int64  `json:"pan"`
		Pitch  int64  `json:"pitch"`
		Volume int64  `json:"volume"`
	} `json:"battleBgm"`
	Battleback1Name string `json:"battleback1Name"`
	Battleback2Name string `json:"battleback2Name"`
	BattlerHue      int64  `json:"battlerHue"`
	BattlerName     string `json:"battlerName"`
	Boat            struct {
		Bgm struct {
			Name   string `json:"name"`
			Pan    int64  `json:"pan"`
			Pitch  int64  `json:"pitch"`
			Volume int64  `json:"volume"`
		} `json:"bgm"`
		CharacterIndex int64  `json:"characterIndex"`
		CharacterName  string `json:"characterName"`
		StartMapID     int64  `json:"startMapId"`
		StartX         int64  `json:"startX"`
		StartY         int64  `json:"startY"`
	} `json:"boat"`
	CurrencyUnit string `json:"currencyUnit"`
	DefeatMe     struct {
		Name   string `json:"name"`
		Pan    int64  `json:"pan"`
		Pitch  int64  `json:"pitch"`
		Volume int64  `json:"volume"`
	} `json:"defeatMe"`
	EditMapID  int64    `json:"editMapId"`
	Elements   []string `json:"elements"`
	EquipTypes []string `json:"equipTypes"`
	GameTitle  string   `json:"gameTitle"`
	GameoverMe struct {
		Name   string `json:"name"`
		Pan    int64  `json:"pan"`
		Pitch  int64  `json:"pitch"`
		Volume int64  `json:"volume"`
	} `json:"gameoverMe"`
	Locale         string  `json:"locale"`
	MagicSkills    []int64 `json:"magicSkills"`
	MenuCommands   []bool  `json:"menuCommands"`
	OptDisplayTp   bool    `json:"optDisplayTp"`
	OptDrawTitle   bool    `json:"optDrawTitle"`
	OptExtraExp    bool    `json:"optExtraExp"`
	OptFloorDeath  bool    `json:"optFloorDeath"`
	OptFollowers   bool    `json:"optFollowers"`
	OptSideView    bool    `json:"optSideView"`
	OptSlipDeath   bool    `json:"optSlipDeath"`
	OptTransparent bool    `json:"optTransparent"`
	PartyMembers   []int64 `json:"partyMembers"`
	Ship           struct {
		Bgm struct {
			Name   string `json:"name"`
			Pan    int64  `json:"pan"`
			Pitch  int64  `json:"pitch"`
			Volume int64  `json:"volume"`
		} `json:"bgm"`
		CharacterIndex int64  `json:"characterIndex"`
		CharacterName  string `json:"characterName"`
		StartMapID     int64  `json:"startMapId"`
		StartX         int64  `json:"startX"`
		StartY         int64  `json:"startY"`
	} `json:"ship"`
	SkillTypes []string `json:"skillTypes"`
	Sounds     []struct {
		Name   string `json:"name"`
		Pan    int64  `json:"pan"`
		Pitch  int64  `json:"pitch"`
		Volume int64  `json:"volume"`
	} `json:"sounds"`
	StartMapID int64    `json:"startMapId"`
	StartX     int64    `json:"startX"`
	StartY     int64    `json:"startY"`
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
		ActorID int64   `json:"actorId"`
		Equips  []int64 `json:"equips"`
		Level   int64   `json:"level"`
	} `json:"testBattlers"`
	TestTroopID int64  `json:"testTroopId"`
	Title1Name  string `json:"title1Name"`
	Title2Name  string `json:"title2Name"`
	TitleBgm    struct {
		Name   string `json:"name"`
		Pan    int64  `json:"pan"`
		Pitch  int64  `json:"pitch"`
		Volume int64  `json:"volume"`
	} `json:"titleBgm"`
	Variables []string `json:"variables"`
	VersionID int64    `json:"versionId"`
	VictoryMe struct {
		Name   string `json:"name"`
		Pan    int64  `json:"pan"`
		Pitch  int64  `json:"pitch"`
		Volume int64  `json:"volume"`
	} `json:"victoryMe"`
	WeaponTypes []string `json:"weaponTypes"`
	WindowTone  []int64  `json:"windowTone"`
}
