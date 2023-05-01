package terminal

import (
	"path"

	"github.com/ansurfen/cushion/utils"
)

type HuloEnv struct {
	Lang     string `yaml:"lang" json:"lang"`
	Daemon   int    `yaml:"daemon" json:"daemon"`
	Terminal struct {
		Title  string `yaml:"title" json:"title"`
		Prefix string `yaml:"prefix" json:"prefix"`
	} `yaml:"terminal" json:"terminal"`
	Boot   []string  `yaml:"boot" json:"boot"`
	Theme  HuloTheme `yaml:"theme" json:"theme"`
	Global bool      `yaml:"global" json:"global"`
}

func NewHuloEnv() *HuloEnv {
	env := utils.NewEnv(utils.EnvOpt[HuloEnv]{
		Workdir: ".hulo",
		Subdirs: []string{"bin", "lib", "loader", "theme"},
		Payload: HuloEnv{},
		BlankConf: `workdir: 
lang: 
daemon: 
boot: 
global: true
theme: 
 name: 
 mode: asni`,
	})

	Workdir = utils.GetEnv().Workdir()
	BinPath = path.Join(Workdir, "bin")
	LoaderPath = path.Join(Workdir, "loader")
	ThemePath = path.Join(Workdir, "theme")
	LibPath = path.Join(Workdir, "lib")
	CompletePath = path.Join(Workdir, "complete")
	return env
}

func (env *HuloEnv) SetLang(lang string) *HuloEnv {
	env.Lang = lang
	utils.GetEnv().Commit("lang", env.Lang)
	return env
}

func (env *HuloEnv) SetDaemon() *HuloEnv {
	utils.GetEnv().Commit("daemon", env.Daemon)
	return env
}

func (env *HuloEnv) CommitAll() *HuloEnv {
	utils.GetEnv().Commit("", env)
	return env
}

func (env *HuloEnv) Write() {
	utils.GetEnv().Write()
}

var (
	Workdir      string
	BinPath      string
	LoaderPath   string
	ThemePath    string
	LibPath      string
	CompletePath string
)
