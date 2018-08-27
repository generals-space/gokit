package main

func init() {
	userManager = &UserManager{
		Users: []*User{
			&User{
				Name:    "李彦宏",
				Title:   "总经理",
				Company: "北京",
			},
			&User{
				Name:    "马云",
				Title:   "市场总监",
				Company: "杭州",
			},
			&User{
				Name:    "马化腾",
				Title:   "商务专员",
				Company: "广州",
			},
		},
	}
}
