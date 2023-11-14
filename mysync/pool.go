package mysync

import (
	"fmt"
	"sync"
)

type Operation interface {
	Hello()
}

type Person struct {
	Name string
}

func (p *Person) Hello() {
	fmt.Println(p.Name + "：你好")
}

type USAPerson struct {
	Name string
}

func (p *USAPerson) Hello() {
	fmt.Println(p.Name + "：hello")
}

func Pool() {
	p := sync.Pool{
		New: func() any {
			return new(Operation)
		},
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		person := p.Get()
		fmt.Println(person)
		//person, ok := p.Get().(*Person)
		//if !ok {
		//	panic("is not person")
		//}
		//
		//person.Hello()
	}()

	go func() {
		defer wg.Done()
		p.Put(USAPerson{
			"Q",
		})

		person := p.Get().(USAPerson)
		person.Hello()
	}()

	wg.Wait()

}
