package models

type PokemonGameIndices struct {
	Game_index int
	Version    PokemonUrl
}

type PokemonUrl struct {
	Name string
	Url  string
}

type PokemonAbilities struct {
	Ability   PokemonUrl
	Is_hidden *bool
	Slot      int
}

type PokemonForms struct {
	Name string
	Url  string
}

type PokemonVersionGroupDetail struct {
	Level_learned_at  int
	Move_learn_method PokemonUrl
	Version_group     PokemonUrl
}

type PokemonSpritesForms struct {
	Back_default       string
	Back_female        string
	Back_shiny         string
	Back_shiny_female  string
	Front_default      string
	Front_female       string
	Front_shiny        string
	Front_shiny_female string
}

type PokemonSpritesOthers struct {
	Dream_world PokemonSpritesForms
	Home        PokemonSpritesForms
}

type PokemonSprites struct {
	PokemonSpritesForms
	Other PokemonSpritesOthers
}

type PokemonMoves struct {
	Move                  PokemonUrl
	Version_group_details []PokemonVersionGroupDetail
}
type Pokemon struct {
	Abilities                []PokemonAbilities
	Base_experience          int
	Forms                    []PokemonForms
	Game_indices             []PokemonGameIndices
	Height                   int
	Id                       int
	Is_default               *bool
	Location_area_encounters string
	Moves                    []PokemonMoves
	Name                     string
	Order                    int
	Species                  PokemonUrl
	Sprites                  PokemonSprites
}
