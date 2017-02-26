package typequery

import "go/types"

func Exec(info *types.Info, q Query) []types.Object {
	objectSet := map[string]types.Object{}

	for _, o := range info.Defs {
		if q.Exec(o) {
			objectSet[o.Id()] = o
		}
	}

	for _, o := range info.Uses {
		if q.Exec(o) {
			objectSet[o.Id()] = o
		}
	}

	objects := make([]types.Object, 0, len(objectSet))
	for _, o := range objectSet {
		objects = append(objects, o)
	}

	return objects
}
