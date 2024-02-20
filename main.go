package main

import (
	"emuser/types"
	"fmt"
	"io"
	"os"
	"text/template"
)

func main() {
	resume := types.MyResume{
		Name:     "Dinesh Kumar G",
		Location: "Thanjavur",
		Contact: map[string]string{
			"Github": "https://github.com/bakageddy",
		},
		Email:       "gdineshda@gmail.com",
		Mobile:      "1234567890",
		Summary:     "Dei Naa summa pandren da",
		Designation: "Full Stack Developer",
		Skills: map[string]string{
			"Languages:": "Python, Rust, Go, C",
			"Databases:": "MySQL, PostgreSQL",
		},
		Experience: []types.Experience{
			{
				Designation: "Backend Developer",
				Company:     "Light2Speed",
				Duration:    "May 2023 - August 2023",
				Location: "Pallikaranai",
				Description: []string{
					"Naa ennamo pannen",
					"Vella kudupanga nu sonnaga",
					"Stipend kudupanga num sonnaga",
					"Ana certification eh kudukala",
				},
			},
			{
				Designation: "Student Leader",
				Company:     "IITM PALS - JCE",
				Duration:    "May 2022 - May 2023",
				Location: "Velachery",
				Description: []string{
					"Super ah cut adichen",
					"Ipavum anda pera use panni cut adikuren",
					"Jolly ah irunduchu",
				},
			},
		},
		Education: []types.Education{
			{
				Instituition: "Sree Narayana Mission Senior Secondary School",
				Location:     "West Mambalam, Chennai",
				Degree:       "12th - 91.8",
				Duration:     "2007 - 2021",
			},
			{
				Instituition: "Jerusalem College of Engineering",
				Location:     "Pallikaranai, Chennai",
				Degree:       "B.E. CSE",
				Duration:     "2021 - 2025",
			},
		},

		Projects: []types.Project{
			{
				Name:       "Dyna - Search Engine",
				Href:       "https://github.com/bakageddy/dyna",
				TechStack:  "Rust, Svelte, Python",
				SourceCode: "https://github.com/bakageddy/dyna",
				Description: []string{
					"Ennamo rendu naal la pannen",
					"Work achu",
					"avolo dhan",
					"Kelvi lam kekakudadu",
				},
			},
			{
				Name:       "Hachiman",
				Href:       "https://github.com/hachiman-jce",
				TechStack:  "Rust, Svelte, PostgreSQL",
				SourceCode: "https://github.com/hachiman-jce/",
				Description: []string{
					"Attendance System for my college",
					"Meedhilam solla maten",
					"Neeye poi pathuko!",
				},
			},
		},
	}

	template_file, err := os.Open("./templates/rezume.latex.txt")
	if err != nil {
		fmt.Println("Failed to open file", err)
		return
	}

	templateBody, err := io.ReadAll(template_file)
	if err != nil {
		fmt.Println("Failed to read file", err)
		return
	}
	temp, err := template.New("rezume").Delims("<<", ">>").Parse(string(templateBody))
	if err != nil {
		fmt.Println("Failed to parse template", err)
		return
	}

	latex, err := os.Create("output.latex")
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}

	err = temp.ExecuteTemplate(latex, "rezume", resume)
	if err != nil {
		fmt.Println("Failed to produce to file:", err)
		return
	}
}
