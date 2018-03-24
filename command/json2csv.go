package command

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/jiro4989/convmv/data"
)

func CmdJson2csv(c *cli.Context) error {
	// read skills.json
	b, err := ioutil.ReadFile("./data/Skills.json")
	if err != nil {
		log.Println(err)
		return err
	}

	var skills []data.Skill
	if err := json.Unmarshal(b, &skills); err != nil {
		log.Println(err)
		return err
	}

	// read system.json
	b, err = ioutil.ReadFile("./data/System.json")
	if err != nil {
		log.Println(err)
		return err
	}

	var system data.System
	if err := json.Unmarshal(b, &system); err != nil {
		log.Println(err)
		return err
	}

	buf := &bytes.Buffer{}
	w := csv.NewWriter(buf)
	for _, skill := range skills {
		record, err := convSkillStrings(skill, system)
		if err != nil {
			log.Println(err)
			return err
		}

		if err := w.Write(record); err != nil {
			return err
		}
	}
	fmt.Println(buf.String())

	// fmt.Println(system)

	// write data/convmv_skills.json

	return nil
}

func convSkillStrings(skill data.Skill, system data.System) ([]string, error) {
	ef, err := json.Marshal(skill.Effects)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	st := data.SystemTexts["ja"]
	system.SkillTypes[0] = "なし"
	system.Elements[0] = "なし"
	// 先頭appendはパフォーマンス悪いけれど一回だけなんで気にしない
	system.Elements = append([]string{"通常攻撃"}, system.Elements...)

	skillStrs := []string{
		strconv.Itoa(skill.ID),
		skill.Name,
		strconv.Itoa(skill.IconIndex),
		skill.Description,
		strconv.Itoa(skill.AnimationID),
		skill.Damage.Formula,
		strconv.FormatBool(skill.Damage.Critical),
		system.Elements[skill.Damage.ElementID+1],
		st.DamageTypes[skill.Damage.Type],
		strconv.Itoa(skill.Damage.Variance),
		strconv.Itoa(skill.MpCost),
		strconv.Itoa(skill.TpCost),
		strconv.Itoa(skill.TpGain),
		strconv.Itoa(skill.Speed),
		st.HitTypes[skill.HitType],
		strconv.Itoa(skill.Occasion),
		strconv.Itoa(skill.Repeats), //
		strconv.Itoa(skill.Scope),   //
		strconv.Itoa(skill.SuccessRate),
		system.SkillTypes[skill.StypeID],
		strconv.Itoa(skill.RequiredWtypeID1),
		strconv.Itoa(skill.RequiredWtypeID2),
		skill.Message1,
		skill.Message2,
		skill.Note,
		string(ef),
	}

	return skillStrs, nil
}
