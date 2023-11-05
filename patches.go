package mockeys

import "github.com/bytedance/mockey"

type Patches struct {
	mockers    []*mockey.Mocker
	mockerVars []*mockey.MockerVar
}

func NewPatches() *Patches {
	return &Patches{}
}

func (this *Patches) Apply(builder *mockey.MockBuilder) *Patches {
	this.mockers = append(this.mockers, builder.Build())
	return this
}

func (this *Patches) ApplyVar(mockerVar *mockey.MockerVar) *Patches {
	this.mockerVars = append(this.mockerVars, mockerVar)
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
