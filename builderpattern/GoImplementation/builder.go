//Any Desktop Builder must know how to build CPU, RAM, Storage, GPU and finally give me the Desktop.
package main

type DesktopBuilder interface{
	BuildCPU()
	BuildRAM()
	BuildStorage()
	BuildGPU()
	GetResult() Desktop
}