/***** Select with a periodic timer ***/

package main

import (
	"log"
	"time"
)

const tickRate = 2 * time.Second

func main() {
	log.Println("start")
	ticker := time.NewTicker(tickRate).C // periodic time -- NewTicker is the function thet returning the ticker object -- ticker object.C je novi chanal koji vraca vreme vise od jednom
	stopper := time.After(5 * tickRate)  // one shot - - kada dodjete do odredjene vremenske tacke daje signal da je gotov i zavrsava sa radom
loop:
	for {
		select {
		case <-ticker: // kada ticker posalje vrednost stampaj "tick"
			log.Println("tick")
		case <-stopper: //kada vreme istekne izadji iz for petlje
			break loop
		}
	}
	log.Println("finish")
}
