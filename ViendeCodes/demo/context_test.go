package demo

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)

	defer cancel() //在退出方法的时候取消context

	time.Sleep(2 * time.Second)

	err := timeoutCtx.Err()
	fmt.Println(err)
}

func TestContext2(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)

	cancel() //先取消context，err就是context canceled

	time.Sleep(500 * time.Millisecond)

	err := timeoutCtx.Err()
	fmt.Println(err) //可以通过err的值看是主动cancel还是timeout等
}

func TestContextValue(t *testing.T) {
	ctx := context.Background()
	ctxValue := context.WithValue(ctx, "key1", "value1")
	val := ctxValue.Value("key1")
	fmt.Println(val) //可以通过err的值看是主动cancel还是timeout等
}

func TestContextParent(t *testing.T) {
	ctx := context.Background()
	dlCtx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Minute))
	ctxValue := context.WithValue(dlCtx, "key1", "value1")
	cancel()              //cancel只是个信号，下层（子）被上层（父）控制
	err := ctxValue.Err() //因为父级dlCtx被cancel了，所以子ctx也会被cancel
	fmt.Println(err)      //可以通过err的值看是主动cancel还是timeout等
}

func TestContextParentValue(t *testing.T) {
	ctx := context.Background()
	dlCtx := context.WithValue(ctx, "map", map[string]string{})
	ctxValue := context.WithValue(dlCtx, "下游参数1", "value1")

	mp := ctxValue.Value("map").(map[string]string) //类型断言，就是这个具体类型
	mp["下游参数1"] = "value1"

	val := dlCtx.Value("map") //即使从父context中拿出，仍然是可以访问到的
	fmt.Println(val)
}
