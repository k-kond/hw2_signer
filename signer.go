package main

import "fmt"

// ExecutePipeline обеспечивает нам конвейерную обработку функций-воркеров, которые что-то делают.
func ExecutePipeline(jobs ...job) {

}

// SingleHash - считает значение crc32(data)+"~"+crc32(md5(data))
// ( конкатенация двух строк через ~),
// где data - то что пришло на вход (по сути - числа из первой функции)
// crc32 считается через функцию DataSignerCrc32
// md5 считается через DataSignerMd5
//
// 0 SingleHash data 0
// 0 SingleHash md5(data) cfcd208495d565ef66e7dff9f98764da
// 0 SingleHash crc32(md5(data)) 502633748
// 0 SingleHash crc32(data) 4108050209
// 0 SingleHash result 4108050209~502633748
func SingleHash(in, out chan interface{}) {
	inMD5ch := make(chan string)
	crc32md5ch := make(chan string)
	crc32ch := make(chan string, 100)

	go func(inStrCh <-chan string) {
		for {
			select {
			case <-cancelCh:
				return
			case dataCh <- val:
				val++
			}
		}
		for inStr := range inStrCh {
			md5 := DataSignerMd5(inStr)
			fmt.Println("%1 SingleHash md5(data) %2", x, md5)
			crc32md5ch <- md5
		}
	}(inMD5ch)

	for i := range in {
		x := i.(string)

		fmt.Println("%1 SingleHash data %2", x, x)
		// Наполняем очередь МД5
		// Обработка очереди МД5
		go func() {
			j := <-crc32md5ch
			crc32 := DataSignerCrc32(j)
			fmt.Println("%1 SingleHash crc32(md5(data)) %2", x, crc32)
			crc32ch <- crc32
		}()
	}

	crcX := DataSignerCrc32(x)
	fmt.Println("%1 SingleHash crc32(data) %2", x, crcX)
}

// MultiHash - считает значение crc32(th+data))
// (конкатенация цифры, приведённой к строке и строки),
// где th=0..5 ( т.е. 6 хешей на каждое входящее значение ),
// потом берёт конкатенацию результатов в порядке расчета (0..5),
// где data - то что пришло на вход (и ушло на выход из SingleHash)
func MultiHash(in, out chan interface{}) {

}

// CombineResults - получает все результаты,
// сортирует (https://golang.org/pkg/sort/),
//объединяет отсортированный результат через _ (символ подчеркивания) в одну строку
func CombineResults(in, out chan interface{}) {

}
