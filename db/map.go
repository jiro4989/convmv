package db

type (
	Map struct {
		AutoplayBgm     bool   `json:"autoplayBgm"`
		AutoplayBgs     bool   `json:"autoplayBgs"`
		Battleback1Name string `json:"battleback1Name"`
		Battleback2Name string `json:"battleback2Name"`
		Bgm             struct {
			Name   string `json:"name"`
			Pan    int64  `json:"pan"`
			Pitch  int64  `json:"pitch"`
			Volume int64  `json:"volume"`
		} `json:"bgm"`
		Bgs struct {
			Name   string `json:"name"`
			Pan    int64  `json:"pan"`
			Pitch  int64  `json:"pitch"`
			Volume int64  `json:"volume"`
		} `json:"bgs"`
		Data              []int64       `json:"data"`
		DisableDashing    bool          `json:"disableDashing"`
		DisplayName       string        `json:"displayName"`
		EncounterList     []interface{} `json:"encounterList"`
		EncounterStep     int64         `json:"encounterStep"`
		Events            []*MapEvent   `json:"events"`
		Height            int64         `json:"height"`
		Note              string        `json:"note"`
		ParallaxLoopX     bool          `json:"parallaxLoopX"`
		ParallaxLoopY     bool          `json:"parallaxLoopY"`
		ParallaxName      string        `json:"parallaxName"`
		ParallaxShow      bool          `json:"parallaxShow"`
		ParallaxSx        int64         `json:"parallaxSx"`
		ParallaxSy        int64         `json:"parallaxSy"`
		ScrollType        int64         `json:"scrollType"`
		SpecifyBattleback bool          `json:"specifyBattleback"`
		TilesetID         int64         `json:"tilesetId"`
		Width             int64         `json:"width"`
	}
	MapEvent struct {
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Note  string `json:"note"`
		Pages []struct {
			Conditions struct {
				ActorID         int64  `json:"actorId"`
				ActorValid      bool   `json:"actorValid"`
				ItemID          int64  `json:"itemId"`
				ItemValid       bool   `json:"itemValid"`
				SelfSwitchCh    string `json:"selfSwitchCh"`
				SelfSwitchValid bool   `json:"selfSwitchValid"`
				Switch1Id       int64  `json:"switch1Id"`
				Switch1Valid    bool   `json:"switch1Valid"`
				Switch2Id       int64  `json:"switch2Id"`
				Switch2Valid    bool   `json:"switch2Valid"`
				VariableID      int64  `json:"variableId"`
				VariableValid   bool   `json:"variableValid"`
				VariableValue   int64  `json:"variableValue"`
			} `json:"conditions"`
			DirectionFix bool `json:"directionFix"`
			Image        struct {
				CharacterIndex int64  `json:"characterIndex"`
				CharacterName  string `json:"characterName"`
				Direction      int64  `json:"direction"`
				Pattern        int64  `json:"pattern"`
				TileID         int64  `json:"tileId"`
			} `json:"image"`
			List []struct {
				Code       int64         `json:"code"`
				Indent     int64         `json:"indent"`
				Parameters []interface{} `json:"parameters"`
			} `json:"list"`
			MoveFrequency int64 `json:"moveFrequency"`
			MoveRoute     struct {
				List []struct {
					Code       int64         `json:"code"`
					Parameters []interface{} `json:"parameters"`
				} `json:"list"`
				Repeat    bool `json:"repeat"`
				Skippable bool `json:"skippable"`
				Wait      bool `json:"wait"`
			} `json:"moveRoute"`
			MoveSpeed    int64 `json:"moveSpeed"`
			MoveType     int64 `json:"moveType"`
			PriorityType int64 `json:"priorityType"`
			StepAnime    bool  `json:"stepAnime"`
			Through      bool  `json:"through"`
			Trigger      int64 `json:"trigger"`
			WalkAnime    bool  `json:"walkAnime"`
		} `json:"pages"`
		X int64 `json:"x"`
		Y int64 `json:"y"`
	}
)
