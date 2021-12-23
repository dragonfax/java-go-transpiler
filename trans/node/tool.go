package node

func ListOfNodesToNodeList[T Node](list []T) []Node {
	if list == nil {
		return nil
	}
	l := make([]Node, len(list))
	for i, n := range list {
		l[i] = n
	}
	return l
}

func NodeListToStringList[T Node](list []T) []string {
	s := make([]string, 0, len(list))
	for _, n := range list {
		s = append(s, n.String())
	}
	return s
}

func AppendNodeLists[T Node, S Node](listA []T, listB ...S) []Node {
	listC := make([]Node, len(listA)+len(listB))

	for i, l := range listA {
		listC[i] = l
	}

	for i, l := range listB {
		listC[i+len(listA)] = l
	}

	return listC
}
