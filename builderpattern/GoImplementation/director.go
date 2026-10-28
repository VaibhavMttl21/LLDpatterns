package main

type Director struct{
	builder DesktopBuilder
}

func (d * Director) construct(){
	d.builder.BuildCPU()
	d.builder.BuildRAM()
	d.builder.BuildGPU()
	d.builder.BuildStorage()
}