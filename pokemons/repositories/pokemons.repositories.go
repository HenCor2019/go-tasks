package PokemonsRepositories

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"github.com/HenCor2019/task-go/models"
)

func Fetch(ids []int, wg *sync.WaitGroup) []models.Pokemon {
  pokemons := []models.Pokemon{}
  for _,pokemonId := range ids {
    wg.Add(1)
    go func(id int) {
      resp,err := http.Get("https://pokeapi.co/api/v2/pokemon/" + strconv.Itoa(id))
      if err != nil {
        log.Fatal("Cannot fetch the request with id: ", id)
        return
      }

      if resp.Status != "200 OK" {
        log.Fatal("Cannot fetch the request with id: ", id)
        return
      }

      defer resp.Body.Close()
      bodyBytes, _ := io.ReadAll(resp.Body)
      var pokemonStruct models.Pokemon
      err = json.Unmarshal(bodyBytes, &pokemonStruct)
      if err != nil {
        log.Fatal("Cannot fetch the request with id: ", id)
        return
      }

      pokemons = append(pokemons,pokemonStruct)
      wg.Done()
    }(pokemonId)
  }
  wg.Wait()

  return pokemons
}

func FindManyById(pokemonsIds []int) ([]models.Pokemon,error) {
  wg := &sync.WaitGroup {}
  pokemons := Fetch(pokemonsIds,wg)

  return pokemons,nil
}
