package main

import (
	"GoGenius/src/datastructures"
	"fmt"
)

func main() {
	fmt.Print("Welcome")
	// basics.VariableExp()
	// basics.ConditionalExp()
	// basics.SwitchExp()
	// basics.FunctionExp()
	// basics.LoopExp()
	// basics.ArrayExp()
	// basics.SliceExp()
	// basics.MapExp()
	// basics.StructExp()
	// basics.MethodExp()
	// basics.GoroutineExp()
	// basics.ChannelExp()
	// basics.SynchonizationUsingChannelExp()
	// basics.MutexExp()
	// api.HandleRequests()
	// file.ReadExcelFile()
	// ile.ReadJsonFile()
	list := datastructures.LinkedList{}
	list.AddFirst(5)
	list.AddFirst(7)
	list.AddLast(10)
	list.Print()
}
