package cache

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	INFINITY = -1
	DEFAULT  = 0
)

var cac *Cache

func init() {
	cac = New(10*time.Hour, 20*time.Minute)
}

func Cac(imia string, otpv string) {

	switch otpv {
	case "get":
		parol, z := (cac).Get(imia)
		if z == true {
			fmt.Printf("Password for imie %s is %s\n", imia, parol)
		} else {
			fmt.Printf("Password for imie %s not found\n", imia)
		}

	default:
		//fmt.Println(cache.defaultExpiration)
		(cac).set(imia, otpv, 2*time.Minute)
		fmt.Println("Cach is writed\n", cac.items)
	}
}

type Item struct {
	Object     interface{}
	Expiration int64 ///its a field for save unix time varable, после которого ты не видишь переменную "ключ-значение", тюе время истечения.
}
type janitor struct {
	Interval time.Duration // время, после которого тип janitor придет чистить кэш от файлов
	stop     chan bool     // переменнаяю которая информирует janitor что очитска кэша больше не требуется
}
type cache struct {
	defaultExpiration time.Duration   //Время жизни кеша по-умолчанию (этот параметр можно будет переопределить для каждого элемента)
	items             map[string]Item // мара с параметрами ключ - значение
	mu                sync.RWMutex    //mu - это блокировка. Так как мар не потокообразный Блокировка помогает гарантировать, что к карте не обращаются два разных потока для записи одновременно. sync.RWMutex имеет два разных вида блокировки – Lock() и RLock(). Lock() позволяет одновременно считывать и записывать только одной подпрограмме. RLock() позволяет нескольким подпрограммам одновременно читать, но не записывать.
	onEvicted         func(string)    //вызывается пользователем и действует как обратный вызов при удалении данных, что бы заменить старые данные актуальными
	janitor           *janitor        //очищает данные из кэша по истечении заданного времени
}
type Cache struct { // эта струткура необходима что бы сослаться на cahe из другого файла или сервера, так как переменные написанные с маленькой буквы- локальные
	*cache
}

func New(defaultExpiration time.Duration, cleanUpInterval time.Duration) *Cache {
	if defaultExpiration == 0 {
		defaultExpiration = INFINITY
	}
	c := &cache{
		defaultExpiration: defaultExpiration,
		items:             make(map[string]Item), // инициализация мапы
	}
	C := &Cache{c}           //структура кэш с новым кэшем
	if cleanUpInterval > 0 { // если время интервала времени после которого янитор все удаляет больне 0
		runJanitor(c, cleanUpInterval)       // новый кэш и время ci
		runtime.SetFinalizer(C, stopJanitor) //
	}
	return C
}

func (j *janitor) Run(c *cache) { //янитор.кэш передается сюда, т.е янитор и подпись к какоум кэщу он относится
	ticker := time.NewTicker(j.Interval)
	for {
		select {
		case <-ticker.C:
			c.DeleteExpired()
		case <-j.stop:
			ticker.Stop()
			return
		}
	}
}
func stopJanitor(cache *Cache) {
	cache.janitor.stop <- true
}
func runJanitor(c *cache, ci time.Duration) {
	j := &janitor{
		Interval: ci,
		stop:     make(chan bool),
	}
	c.janitor = j // заппись янитора для файла cashe
	go j.Run(c)
}

func (c *cache) Add(k string, v interface{}, d time.Duration) error {
	c.mu.Lock()
	_, found := c.Get(k)
	if found {
		c.mu.Unlock()
		return fmt.Errorf("Item %s already exists", k)
	}
	c.set(k, v, d)
	c.mu.Unlock()
	return nil
}
func (c *cache) set(k string, v interface{}, d time.Duration) {
	var e int64
	if d == 0 {
		d = c.defaultExpiration
	}
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.items[k] = Item{
		Object:     v,
		Expiration: e,
	}
}
func (c *cache) Get(k string) (interface{}, bool) {
	c.mu.RLock()
	item, found := c.items[k]
	if !found {
		c.mu.RUnlock()
		return "NO1", false
	}
	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			c.mu.RUnlock()
			return "NO2EXP", false
		}
	}
	c.mu.RUnlock()
	return item.Object, true
}

func (c *cache) DeleteExpired() {
	var evictedItems []string
	now := time.Now().UnixNano()
	c.mu.Lock()
	for k, v := range c.items {
		// "Inlining" of expired
		if v.Expiration > 0 && now > v.Expiration {
			delete(c.items, k) // удаляем ключ из кэша и значение
			evictedItems = append(evictedItems, k)

		}
	}
	c.mu.Unlock()
	for _, v := range evictedItems {
		c.onEvicted(v)
	}
}
