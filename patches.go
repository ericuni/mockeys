package mockeys

import "github.com/bytedance/mockey"

type Patches struct {
	mockers    []*mockey.Mocker
	mockerVars []*mockey.MockerVar
}

func NewPatches() *Patches {
	return &Patches{}
}

func (this *Patches) Apply(builders ...*mockey.MockBuilder) *Patches {
	for _, builder := range builders {
		this.mockers = append(this.mockers, builder.Build())
	}
	return this
}

func (this *Patches) ApplyVar(mockerVars ...*mockey.MockerVar) *Patches {
	for _, mockerVar := range mockerVars {
		this.mockerVars = append(this.mockerVars, mockerVar.Patch())
	}
	return this
}

func (this *Patches) Reset() {
	for i := len(this.mockerVars) - 1; i >= 0; i-- {
		this.mockerVars[i].UnPatch()
	}

	for i := len(this.mockers) - 1; i >= 0; i-- {
		this.mockers[i].UnPatch()
	}
}
