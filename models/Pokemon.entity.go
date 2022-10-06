package models

type PokemonAbility struct {
  Name string
  Url string

}

type PokemonAbilities struct {
  Ability PokemonAbility
  Is_hidden *bool
  Slot int
}

type PokemonForms struct {
  Name string
  Url string
}

type Pokemon struct {
  Abilities []PokemonAbilities
  Base_experience int
  Forms []PokemonForms
  Height int
  Id int
  Is_default *bool
  Location_area_encounters string
  Name string
  Order int
}
