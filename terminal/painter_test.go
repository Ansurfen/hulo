package terminal

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ansurfen/cushion/runtime"
	"github.com/ansurfen/cushion/utils"
)

const sandJSON = `
{
    "color": {
        "myprimary": {
            "ansi256": "100",
            "ansi": "5"
        }
    },
    "render": {
        "SuggestionTextColor": "primary",
        "SuggestionBGColor": "myprimary"
    }
}
`

func TestPainter(t *testing.T) {
	painter := NewHuloPainter()
	var layout = Layout{
		Color:  make(map[string]Color),
		Render: make(map[string]string),
	}
	if err := json.Unmarshal([]byte(sandJSON), &layout); err != nil {
		panic(err)
	}
	mode := "ansi"
	p := utils.NewReflectObject(painter)
	for field, value := range layout.Render {
		c := "0"
		if cc, ok := painter.colors[value]; ok {
			switch mode {
			case "ansi":
				c = cc.ANSI
			case "ansi256":
				c = cc.ANSI256
			case "rgb":
				c = cc.RGB
			}
		}
		if cc, ok := layout.Color[value]; ok {
			switch mode {
			case "ansi":
				c = cc.ANSI
			case "ansi256":
				c = cc.ANSI256
			case "rgb":
				c = cc.RGB
			}
		}
		p.Set(field, c)
	}
	fmt.Println(painter)
}

func TestGenThemeConf(t *testing.T) {
	painter := NewHuloPainter()
	painter.GenThemeConf("sand")
}

func TestTmp(t *testing.T) {
	fmt.Println(runtime.LuaAssignLR(runtime.LuaIdent("gcc"), runtime.LuaMap(runtime.Luamap{
		"name": runtime.LuaString("ccc"),
	})))
}
