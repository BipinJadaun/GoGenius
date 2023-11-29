package cache


type Node struct {
	data string
	prev *Node
	next *Node
}

type Cache struct {
	cpacity int
	size int
	ivictionPolicy string
	cache map[int]*Node
	head Node
	tail Node
}

func (c Cache) init(maxSize int, policy string) {
	c.cpacity = maxSize
	c.size = 0
	c.ivictionPolicy = policy
}

func(c Cache) laod(key int) string {
	if _, ok := c.cache[key]; ok {
		
	}

	return ""
}