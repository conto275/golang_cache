package main

// Эта программа будет являтся простой программой для входа по данным - имя пользователя и пароль, с функцией автоматического кэш хранения правильных значений для быстрого входа.
import (
	"fmt"
	"math/rand"
	cache "pvs/Cache"
	glowne "pvs/Package2"
	"sync"
	"time"
)

type Alrandom_local struct {
	name     string
	password string
}

func main() {
	wg := &sync.WaitGroup{}
	for x := 0; x < 100; x++ {
		wg.Add(1)
		go idd(wg) /// добавить блокировки и вайт и т.д i dodac go roytine

	}
	wg.Wait()
}
func idd(wg *sync.WaitGroup) {
	defer wg.Done()
	names := make([]string, 0)
	names = append(names, "Alesya", "Vlad", "Nika", "Maxim")
	haslos := make([]string, 0)
	haslos = append(haslos, "62291881", "569884", "VDSK56g", "7ft7ghfhr97g") // "6j4jj46i64jih6", "hh6h46hk4", "h6y6yryyr6", "67u6u665u56u", "ujkj7676667", "kj677t7ut7") //количество повторов комманды или комманда, которая будет указывать на действие)
	nms := names[rand.Intn(len(names))]
	//fmt.Println(names)
	hsl := haslos[rand.Intn(len(haslos))]
	all := glowne.NewAlrandom(nms, hsl)
	all_local := Alrandom_local{
		name:     nms,
		password: hsl,
	}
	otv, tof := all.Alrandomstart()
	if tof == true {
		fmt.Println(otv)
		cache.Cac(all_local.name, all_local.password)
		time.Sleep(time.Second * 1)
		cache.Cac(all_local.name, "get")
	}
	if tof == false {
		fmt.Println(otv)
		return
	}

}

//in := []string{"Vlad", "Nika", "Alesya", "Maxim", "Valera"}
