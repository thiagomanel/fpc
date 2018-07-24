//based on Katherine Cox-Buday, Concurrency in Go
package main

import (
	"fmt"
	"time"
)

//esta função variádica recebe um número arbitrário de canais de leitura
//e retorna um canal de leitura. O canal desbloqueia (por um close)
//quando pelo menos um dos canais recebidos desbloqueia.
func or(channels ...<-chan interface{}) <-chan interface{} {

	//a função é recursiva, aqui temos o critério de parada
	switch len(channels) {
	case 0:
		//importante. caso contrário, bloqueríamos indefinidamente
		return nil
	case 1:
		return channels[0]
	}

	//canal de resposta
	orDone := make(chan interface{})

	//uhm! criamos uma nova goroutine para não bloquear nos cases
	//do switch. Será uma decisão do cliente desta função or bloquear
	//ou não a depender do que ele queira fazer com o canal recebido
	go func() {
		//ao fim do contexto, o canal de resposta é fechado. isso só
		//vai acontecer, ao receber algo dos canais no "case 2" ou no
		//"case default"
		defer close(orDone)
		switch len(channels) {
		//esse caso é uma otimização. poderíamos evitá-lo, todavida.
		//como seria o código sem ele?
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}

func msg(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		fmt.Println("I'm going to sleep for ", after)
		time.Sleep(after)
	}()
	return c
}

func main() {

	start := time.Now()
	<-or(
		msg(2*time.Hour),
		msg(5*time.Minute),
		msg(1*time.Second),
		msg(1*time.Hour),
		msg(1*time.Minute),
	)
	fmt.Printf("done after %v\n", time.Since(start))
}
