package main

type Contact struct {
	Phone    string `yaml:"phone"`
	Email    string `yaml:"email"`
	LinkedIn string `yaml:"linkedIn"`
	Github   string `yaml:"github"`
	Location string `yaml:"location"`
}

type Experience struct {
	Company  string   `yaml:"company"`
	Role     string   `yaml:"role"`
	Title    string   `yaml:"title"`
	From     string   `yaml:"from"`
	To       string   `yaml:"to"`
	Location string   `yaml:"location"`
	Details  []string `yaml:"details"`
}

type Education struct {
	School   string `yaml:"school"`
	Diploma  string `yaml:"diploma"`
	Grade    string `yaml:"grade"`
	From     string `yaml:"from"`
	To       string `yaml:"to"`
	Location string `yaml:"location"`
}

type Skills struct {
	Languages            []string `yaml:"languages"`
	ProgrammingLanguages []string `yaml:"programmingLanguages"`
	Tools                []string `yaml:"tools"`
	Devops               []string `yaml:"devops"`
	cloud                []string `yaml:"cloud"`
}

type CV struct {
	Name       string       `yaml:"name"`
	Occupation string       `yaml:"occupation"`
	Contact    Contact      `yaml:"contact"`
	Experience []Experience `yaml:"experience"`
	Education  []Education  `yaml:"eduation"`
	Skills     Skills       `yaml:"skills"`
}
