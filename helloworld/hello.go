package main //this is the name of the package

import (
	"fmt" //import a package named "fmt"
	"math"
	"strconv"
	//"example.com/user/hello/morestrings"
	//"github.com/google/go-cmp/cmp" //an external package
	"time"
)

func main() {
	//fmt.Println(morestrings.ReverseRunes("!oG ,dlleH"))
	//fmt.Println(cmp.Diff("Hello World", "Hello Go"))

	//in go built in types are passed by val.
	//slices and objects are passed by ref

	//if you want to double an int need to pass by pointer
	n := 2
	doubleit (&n)
	fmt.Println(n)

	//example of using built in error type
	fmt.Println(sqrt(-3))

	//example of defer - which is like a destrutor. It is used to free up resources at the end of a function
	fmt.Println(error_with_defer(-3))

	//exmaple of usting structs
	user := User{
		firstName: "Neil",
		secondName: "McAlpine",
	}

	fmt.Printf("%+v\n", user)

	fmt.Println(user.firstName)

	fmt.Println(user.concat())

	//example of an interface
	PrintUser(user)

	//Channels example
	ch := make(chan int) //make a channel to hold integers

	//this is an anonymous goroutine that sends the number 353 to the channel
	go func() {
		ch <- 353
	}()

	//this sets up a receiver for the channel
	val:= <-ch
	fmt.Printf("got %d\n", val)


	// Send multiple
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	//if you know how many signals we be sent
	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Printf("received %d\n", val)
	}

	//if you don't know how many signals we be sent then the sender ca close the channel when it is done
	// close to signal we're done
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	fmt.Println("signals starting...")

	//this is periodic block/inblock of the main gproutine until the channel closes
	for i := range ch {
		fmt.Printf("received %d\n", i)
	}

	goroutine_worker_example()
}

func goroutine_worker_example() {

	//example of sending work from main to worker goroutines and then stopping the channel when there is no more work.

	//set up two channels
	jobs := make(chan int)
	done := make(chan bool)

	fmt.Println("goroutine_worker_example")

	go func() {
		for {
			//inside we wait for a signal to arrive
			j, more := <-jobs //await on the jobs channel. This is a special two-variable form of channel receive. 
							//more is false unless the jobs channel is closed
			if more  {
				fmt.Printf("received %d\n", j)
			} else {
				fmt.Println("channel is closed so signal back to waiting main")
				done <- true
				return
			}
		}	
		 

	}() //this invokes the anonymous function as a goroutine (as opposed to synchronously)

	//send through signals to the jobs channel
	for i := 0; i < 3; i++ {
		fmt.Printf("sending %d\n", i)
		jobs <- i
		time.Sleep(time.Second)
	}

	close(jobs) //we have finished

	//await for workers to finish before exiting
	<-done
	fmt.Println("all workers finished we can exit")


}

func doubleit(i *int) {
	*i *= 2 //dereference the pointer
}

//error is a built in type
func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("cannot sqrt a negative number")
	} else
	{
		return math.Sqrt(n), nil
	}
}

func error_with_defer(n int) error {

	defer cleanup("A")
	defer cleanup("B")

	if n < 0 {
		return fmt.Errorf("cannot have negative number")
	} else {
		n = n * 2  //do nothing
	}

	return nil

}

func cleanup(res string) {
	fmt.Println("cleaning up" + res)
}


//Go supports simple encapsulation. If a field or method begin with a capital letter then it can be used outsde of the package
//otherwise not.
//This is why all the methods e.. Println start with capital
type User struct {
	firstName string
	secondName string
	referenceId int
}

//to make methods on a struct 
//The (u *User) is known as the receiver and identifies the struct type by which the method is operating
//Note the * means a reference is passed. If this is removed then it passes by val (e.g. a copy is made)
func (u *User) concat () string {
	return u.firstName + "-" + u.secondName
}

//interfaces defines methods
type IAdvancedUser interface {
	concat_all() string
}

//use the interface
func PrintUser(user IAdvancedUser) {
	fmt.Println(user.concat_all())
}

func (u User) concat_all () string {
	return u.firstName + "-" + u.secondName + "-" + strconv.Itoa(u.referenceId)
}