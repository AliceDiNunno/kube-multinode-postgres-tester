package main

func main() {
	LoadEnv()

	gormConfig := LoadGormConfiguration()

	db := StartGormDatabase(gormConfig)

	if db == nil {
		panic("Database connection failed")
	}

	//ginConfig := LoadGinConfiguration()
}