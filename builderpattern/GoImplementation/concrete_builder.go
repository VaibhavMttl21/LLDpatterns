package main

type GamingDesktopBuilder struct{
	desktop Desktop
}

func(d *GamingDesktopBuilder) BuildCPU(){
	d.desktop.CPU= "Intel i9"
}

func (d *GamingDesktopBuilder) BuildRAM(){
	d.desktop.RAM=32
}

func (d *GamingDesktopBuilder) BuildStorage(){
	d.desktop.RAM = 512
}

func (d *GamingDesktopBuilder) BuildGPU(){
	d.desktop.GPU = "RTX 5070"
}

func (b *GamingDesktopBuilder) GetResult() Desktop {
	return b.desktop
}

// DesktopBuilder
//        ▲
//        │ implements
//        │
// GamingDesktopBuilder