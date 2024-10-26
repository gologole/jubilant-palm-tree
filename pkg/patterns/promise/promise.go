package main

/*
	func Promise(task func() int) chan int {
		resultCh := make(chan int, 1) // создаем канал для результата

		go func() {
			result := task()   // выполняем задачу
			resultCh <- result // отправляем результат в канал
			close(resultCh)    // закрываем канал после выполнения
		}()

		return resultCh
	}
*/
