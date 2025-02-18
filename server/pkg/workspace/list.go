package workspace

type List []*Workspace

func (l List) Get(id ID) *Workspace {
	for _, w := range l {
		if w.ID() == id {
			return w
		}
	}
	return nil
}
