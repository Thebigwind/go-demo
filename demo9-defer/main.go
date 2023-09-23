package demo9_defer

import "fmt"

func test() (err error) {
	defer func() {
		p := recover()
		if p != nil {
			switch e := p.(type) {
			case error:
				err = fmt.Errorf("%+v", e)
			default:
				err = fmt.Errorf("%+v", p)
			}
		}
	}()
	return nil
}

func test2() (err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recover_panic")
		}
	}()
	return nil
}
