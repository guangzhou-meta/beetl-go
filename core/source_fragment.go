package core

type SourceFragment struct {
	List           []FragmentI
	startLineIndex int
	startLine      int
}

func NewSourceFragment() *SourceFragment {
	return &SourceFragment{
		startLine:      0,
		startLineIndex: 0,
		List:           make([]FragmentI, 0),
	}
}

func (f *SourceFragment) Add(fr FragmentI) {
	f.List = append(f.List, fr)
	if cr, ok := fr.(*CRFragment); ok {
		f.TrimSpace()
		f.MoveNext(cr)
		return
	}

	if fr.GetEndLine() != f.startLine {
		f.TrimSpace()
		f.MoveNext(fr)
	}
}

func (f *SourceFragment) TrimSpace() {
	hasScript := false
	for i := f.startLineIndex; i < len(f.List); i++ {
		fr := f.List[i]
		switch t := fr.(type) {
		case *CRFragment:
			continue
		case *PlaceholderFragment:
			return
		case *TextFragment:
			if !t.OnlySpace() {
				return
			} else {
				continue
			}
		default:
			hasScript = true
		}
	}

	if !hasScript {
		return
	}

	var lastScript FragmentI
	for i := f.startLineIndex; i < len(f.List); i++ {
		fr := f.List[i]
		switch fr.(type) {
		case *TextFragment:
			fr.SetStatus(del)
		case *CRFragment:
			if lastScript != nil {
				fr.SetStatus(del)
				if ls, ok := lastScript.(ScriptFragmentI); ok {
					ls.AppendCr()
				}
			}
		default:
			lastScript = fr
		}
	}
}

func (f *SourceFragment) MoveNext(fr FragmentI) {
	f.startLineIndex = len(f.List)
	if _, ok := fr.(*CRFragment); ok {
		f.startLine = fr.GetEndLine() + 1
	} else {
		f.startLine = fr.GetEndLine()
	}
}

func (f *SourceFragment) Merge() error {
	f.Check()
	for i := range f.List {
		fr := f.List[i]
		if fr.GetStatus() == del {
			continue
		}
		if text, ok := fr.(*TextFragment); ok {
			for z := i + 1; z < len(f.List); z++ {
				nextFr := f.List[z]
				if _, ok := nextFr.(BeetlFragmentI); ok {
					i = z
					break
				}
				if nextFr.GetStatus() == del {
					continue
				}
				nextFr.SetStatus(del)
				text.AppendTextFragment(nextFr)
			}
		}
	}
	return nil
}

func (f *SourceFragment) Check() {
	if f.startLineIndex != len(f.List) {
		f.TrimSpace()
	}
}
