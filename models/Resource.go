package models

type Resource struct {
	Id           int
	ResourceName string `资源名称`
	ShowName     string `显示名称`
	ParentId     int    `父id`
	Sort         int    `排序`
	Path         string `访问路径`
	Status       int    `状态(0：显示，1，不显示)`
}
