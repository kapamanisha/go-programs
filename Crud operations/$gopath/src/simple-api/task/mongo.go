var database string = "go-development"

func main() {

	session, err := mgo.Dial("mongodb://localhost:27017/" + database)
	if err != nil {
		fmt.Println(err)
	}

	// Cleanup
	defer session.Close()

	c := bootstrap(session)

	create(c)
	read(c)
	update(c)
	delete(c)
}

func bootstrap(s *mgo.Session) *mgo.Collection {

	s.DB(database).DropDatabase()
	c := s.DB(database).C("people")
	index := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		Background: true,
	}

	c.EnsureIndex(index)

	return c
}
