package terminal

import (
	"encoding/json"
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/ansurfen/cushion/go-prompt"
	"github.com/ansurfen/cushion/utils"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
)

type HuloPatiner struct {
	colors map[string]Color
	render map[string]string

	errColor lipgloss.Style
}

func NewHuloPainter() *HuloPatiner {
	painter := &HuloPatiner{
		render: make(map[string]string),
		colors: map[string]Color{
			"transparent": {
				ANSI:    "-1",
				ANSI256: "-1",
				RGB:     "-1",
			},
			"dark-default": {
				ANSI:    "0",
				ANSI256: "16",
				RGB:     "#000000",
			},
			"dark-danger": {
				ANSI:    "1",
				ANSI256: "196",
				RGB:     "#d251a6",
			},
			"dark-success": {
				ANSI:    "2",
				ANSI256: "46",
				RGB:     "#027b77",
			},
			"dark-warning": {
				ANSI:    "3",
				ANSI256: "226",
				RGB:     "#e55426",
			},
			"dark-primary": {
				ANSI:    "4",
				ANSI256: "21",
				RGB:     "#006bbb",
			},
			"dark-important": {
				ANSI:    "5",
				ANSI256: "201",
				RGB:     "#5e38cc",
			},
			"dark-info": {
				ANSI:    "6",
				ANSI256: "51",
				RGB:     "#008080",
			},
			"dark-none": {
				ANSI:    "7",
				ANSI256: "231",
				RGB:     "#4e5458",
			},
			"light-default": {
				ANSI:    "8",
				ANSI256: "238",
				RGB:     "#1c2235",
			},
			"light-danger": {
				ANSI:    "9",
				ANSI256: "203",
				RGB:     "#d251a6",
			},
			"light-success": {
				ANSI:    "10",
				ANSI256: "119",
				RGB:     "#04b9ae",
			},
			"light-warning": {
				ANSI:    "11",
				ANSI256: "227",
				RGB:     "#ff7b52",
			},
			"light-primary": {
				ANSI:    "12",
				ANSI256: "75",
				RGB:     "#00a1e5",
			},
			"light-important": {
				ANSI:    "13",
				ANSI256: "211",
				RGB:     "#8866e9",
			},
			"light-info": {
				ANSI:    "14",
				ANSI256: "123",
				RGB:     "#00FFFF",
			},
			"light-none": {
				ANSI:    "15",
				ANSI256: "255",
				RGB:     "#ffffff",
			},
		},
		errColor: lipgloss.NewStyle().Foreground(lipgloss.Color("#d251a6")),
	}
	return painter
}

type Layout struct {
	Color  map[string]Color  `json:"color"`
	Render map[string]string `json:"render"`
}

func (painter *HuloPatiner) LoadTheme(theme HuloTheme) {
	stream, err := utils.ReadStraemFromFile(path.Join(ThemePath, fmt.Sprintf("%s.json", theme.Name)))
	if err != nil {
		panic(err)
	}
	var layout = Layout{
		Color:  make(map[string]Color),
		Render: make(map[string]string),
	}
	if err = json.Unmarshal(stream, &layout); err != nil {
		panic(err)
	}

	for field, value := range layout.Render {
		c := "0"
		if cc, ok := painter.colors[value]; ok {
			switch theme.Mode {
			case "ansi":
				c = cc.ANSI
			case "ansi256":
				c = cc.ANSI256
			case "rgb":
				c = cc.RGB
			}
		}
		if cc, ok := layout.Color[value]; ok {
			switch theme.Mode {
			case "ansi":
				c = cc.ANSI
			case "ansi256":
				c = cc.ANSI256
			case "rgb":
				c = cc.RGB
			}
		}
		painter.render[field] = c
	}
}

func (painter *HuloPatiner) GenThemeConf(name string) {
	render := utils.NewReflectObject(&prompt.Render{})
	fields := []string{}
	for field := range render.Fields() {
		if strings.HasSuffix(field, "Color") {
			fields = append(fields, field)
		}
	}
	sort.Slice(fields, func(i, j int) bool {
		return fields[i] > fields[j]
	})
	conf := viper.New()
	conf.SetConfigFile(fmt.Sprintf("%s.json", name))
	for _, field := range fields {
		conf.Set(fmt.Sprintf("render.%s", field), "")
	}
	conf.Set("color", map[string]string{})
	conf.WriteConfig()
}

type Color struct {
	RGB     string `json:"rgb"`     // color range 0x000000 - 0xffffff
	ANSI256 string `json:"ansi256"` // color range: 16 - 255
	ANSI    string `json:"ansi"`    // color range: 0 - 15
}

type HuloTheme struct {
	Name string `yaml:"name" json:"name"`
	Mode string `yaml:"mode" json:"mode"`
}
