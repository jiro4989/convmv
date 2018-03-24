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

	buf := &bytes.Buffer{}
	w := csv.NewWriter(buf)
	for _, skill := range skills {
		ef, err := json.Marshal(skill.Effects)
		if err != nil {
			log.Println(err)
			return err
		}

		record := []string{
			strconv.Itoa(skill.ID),
			skill.Name,
			strconv.Itoa(skill.IconIndex),
			skill.Description,
			strconv.Itoa(skill.AnimationID),
			skill.Damage.Formula,
			strconv.FormatBool(skill.Damage.Critical),
			strconv.Itoa(skill.Damage.ElementID),
			strconv.Itoa(skill.Damage.Type),
			strconv.Itoa(skill.Damage.Variance),
			strconv.Itoa(skill.MpCost),
			strconv.Itoa(skill.TpCost),
			strconv.Itoa(skill.TpGain),
			strconv.Itoa(skill.Speed),
			strconv.Itoa(skill.HitType),
			strconv.Itoa(skill.Occasion),
			strconv.Itoa(skill.Repeats),
			strconv.Itoa(skill.Scope),
			strconv.Itoa(skill.SuccessRate),
			strconv.Itoa(skill.StypeID),
			strconv.Itoa(skill.RequiredWtypeID1),
			strconv.Itoa(skill.RequiredWtypeID2),
			skill.Message1,
			skill.Message2,
			skill.Note,
			string(ef),
		}

		if err := w.Write(record); err != nil {
			return err
		}
	}
	fmt.Println(buf.String())

	// read system.json
	// write data/convmv_skills.json

	return nil
}
