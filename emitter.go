package machina

import (
	"fmt"
)

type EmitFunc func(Event)

type SeqFunc struct {
	seq      int
	callback EmitFunc
}

type SeqFuncMap map[int]*SeqFunc

func (s SeqFuncMap) clear() {
	s = map[int]*SeqFunc{}
}

func (s SeqFuncMap) remove(seq int) {
	delete(s, seq)
}

type EmitOnResult struct {
	eventName string
	seqFunc   *SeqFunc
	emitter   *Emitter
}

func (r *EmitOnResult) Off() {
	if r.emitter != nil && r.eventName != "" {
		r.emitter.offFunc(r.eventName, r.seqFunc)
		r.emitter = nil
	}
}

type Emitter struct {
	funcs map[string]SeqFuncMap
	seq   int
}

func NewEmitter() *Emitter {
	rv := &Emitter{
		funcs: map[string]SeqFuncMap{},
		seq:   0,
	}
	return rv
}

func (e *Emitter) Emit(evtName string, evt Event) {
	funcMap, ok := e.funcs[evtName]
	if !ok {
		return
	}

	if len(funcMap) == 0 {
		return
	}

	// 收集函数seq
	var fnSeqs []int
	for _, v := range funcMap {
		fnSeqs = append(fnSeqs, v.seq)
	}

	// 通过seq来遍历处理
	for _, seq := range fnSeqs {
		seqFn, ok := funcMap[seq]
		if !ok {
			fmt.Println("Emitter::emit func is removed in emit!!!!")
			continue
		}

		// call callback func
		seqFn.callback(evt)
	}

}

func (e *Emitter) On(evtName string, callback EmitFunc) *EmitOnResult {
	e.seq++
	seqFunc := &SeqFunc{
		seq:      e.seq,
		callback: callback,
	}

	if e.funcs[evtName] == nil {
		e.funcs[evtName] = SeqFuncMap{}
	}
	e.funcs[evtName][seqFunc.seq] = seqFunc

	return &EmitOnResult{
		eventName: evtName,
		seqFunc:   seqFunc,
		emitter:   e,
	}
}

func (e *Emitter) offFunc(evtName string, seqFunc *SeqFunc) {
	funmap := e.funcs[evtName]
	funmap.remove(seqFunc.seq)
}

func (e *Emitter) Off(evtName string) {
	delete(e.funcs, evtName)
}

func (e *Emitter) OffAll() {
	e.funcs = map[string]SeqFuncMap{}
}
