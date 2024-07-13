package embeddedinterface

type Hitung2d interface {
	Luas() float64
	Keliling() float64
}

type Hitung3d interface {
	Volume() float64
}

type HitungEmbedded interface {
	Hitung2d
	Hitung3d
}

type Kubus struct {
	Sisi float64
}

func (k *Kubus) Luas() float64 {
	return k.Sisi * k.Sisi * 6
}

func (k *Kubus) Volume() float64 {
	return k.Sisi * k.Sisi * k.Sisi
}

func (k *Kubus) Keliling() float64 {
	return k.Sisi * 12
}
