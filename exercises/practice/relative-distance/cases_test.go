package relativedistance

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: ad7e332 fix: relative distance (#2545)

var testCases = []struct {
	description    string
	tree           map[string][]string
	personA        string
	personB        string
	expectedDegree int
	expectedOk     bool
}{
	{
		description:    "Direct parent-child relation",
		tree:           map[string][]string{"Tomoko": []string{"Aditi"}, "Vera": []string{"Tomoko"}},
		personA:        "Vera",
		personB:        "Tomoko",
		expectedDegree: 1,
		expectedOk:     true,
	},
	{
		description:    "Sibling relationship",
		tree:           map[string][]string{"Dalia": []string{"Olga", "Yassin"}},
		personA:        "Olga",
		personB:        "Yassin",
		expectedDegree: 1,
		expectedOk:     true,
	},
	{
		description:    "Two degrees of separation, grandchild",
		tree:           map[string][]string{"Khadija": []string{"Mateo"}, "Mateo": []string{"Rami"}},
		personA:        "Khadija",
		personB:        "Rami",
		expectedDegree: 2,
		expectedOk:     true,
	},
	{
		description:    "Unrelated individuals",
		tree:           map[string][]string{"Kaito": []string{"Elif"}, "Priya": []string{"Rami"}},
		personA:        "Priya",
		personB:        "Kaito",
		expectedDegree: 0,
		expectedOk:     false,
	},
	{
		description:    "Complex graph, cousins",
		tree:           map[string][]string{"Aditi": []string{"Nia"}, "Aiko": []string{"Bao", "Carlos"}, "Bao": []string{"Dalia", "Elias"}, "Boris": []string{"Oscar"}, "Carlos": []string{"Fatima", "Gustavo"}, "Celine": []string{"Priya"}, "Dalia": []string{"Hassan", "Isla"}, "Diego": []string{"Qi"}, "Elias": []string{"Javier"}, "Elif": []string{"Rami"}, "Farah": []string{"Sven"}, "Fatima": []string{"Khadija", "Liam"}, "Giorgio": []string{"Tomoko"}, "Gustavo": []string{"Mina"}, "Hana": []string{"Umar"}, "Hassan": []string{"Noah", "Olga"}, "Ian": []string{"Vera"}, "Isla": []string{"Pedro"}, "Javier": []string{"Quynh", "Ravi"}, "Jing": []string{"Wyatt"}, "Kaito": []string{"Xia"}, "Khadija": []string{"Sofia"}, "Leila": []string{"Yassin"}, "Liam": []string{"Tariq", "Uma"}, "Mateo": []string{"Zara"}, "Mina": []string{"Viktor", "Wang"}, "Nia": []string{"Antonio"}, "Noah": []string{"Xiomara"}, "Olga": []string{"Yuki"}, "Oscar": []string{"Bianca"}, "Pedro": []string{"Zane", "Aditi"}, "Priya": []string{"Cai"}, "Qi": []string{"Dimitri"}, "Quynh": []string{"Boris"}, "Rami": []string{"Ewa"}, "Ravi": []string{"Celine"}, "Sofia": []string{"Diego", "Elif"}, "Sven": []string{"Fabio"}, "Tariq": []string{"Farah"}, "Tomoko": []string{"Gabriela"}, "Uma": []string{"Giorgio"}, "Umar": []string{"Helena"}, "Vera": []string{"Igor"}, "Viktor": []string{"Hana", "Ian"}, "Wang": []string{"Jing"}, "Wyatt": []string{"Jun"}, "Xia": []string{"Kim"}, "Xiomara": []string{"Kaito"}, "Yassin": []string{"Lucia"}, "Yuki": []string{"Leila"}, "Zane": []string{"Mateo"}, "Zara": []string{"Mohammed"}},
		personA:        "Dimitri",
		personB:        "Fabio",
		expectedDegree: 9,
		expectedOk:     true,
	},
	{
		description:    "Complex graph, no shortcut, far removed nephew",
		tree:           map[string][]string{"Aditi": []string{"Nia"}, "Aiko": []string{"Bao", "Carlos"}, "Bao": []string{"Dalia", "Elias"}, "Boris": []string{"Oscar"}, "Carlos": []string{"Fatima", "Gustavo"}, "Celine": []string{"Priya"}, "Dalia": []string{"Hassan", "Isla"}, "Diego": []string{"Qi"}, "Elias": []string{"Javier"}, "Elif": []string{"Rami"}, "Farah": []string{"Sven"}, "Fatima": []string{"Khadija", "Liam"}, "Giorgio": []string{"Tomoko"}, "Gustavo": []string{"Mina"}, "Hana": []string{"Umar"}, "Hassan": []string{"Noah", "Olga"}, "Ian": []string{"Vera"}, "Isla": []string{"Pedro"}, "Javier": []string{"Quynh", "Ravi"}, "Jing": []string{"Wyatt"}, "Kaito": []string{"Xia"}, "Khadija": []string{"Sofia"}, "Leila": []string{"Yassin"}, "Liam": []string{"Tariq", "Uma"}, "Mateo": []string{"Zara"}, "Mina": []string{"Viktor", "Wang"}, "Nia": []string{"Antonio"}, "Noah": []string{"Xiomara"}, "Olga": []string{"Yuki"}, "Oscar": []string{"Bianca"}, "Pedro": []string{"Zane", "Aditi"}, "Priya": []string{"Cai"}, "Qi": []string{"Dimitri"}, "Quynh": []string{"Boris"}, "Rami": []string{"Ewa"}, "Ravi": []string{"Celine"}, "Sofia": []string{"Diego", "Elif"}, "Sven": []string{"Fabio"}, "Tariq": []string{"Farah"}, "Tomoko": []string{"Gabriela"}, "Uma": []string{"Giorgio"}, "Umar": []string{"Helena"}, "Vera": []string{"Igor"}, "Viktor": []string{"Hana", "Ian"}, "Wang": []string{"Jing"}, "Wyatt": []string{"Jun"}, "Xia": []string{"Kim"}, "Xiomara": []string{"Kaito"}, "Yassin": []string{"Lucia"}, "Yuki": []string{"Leila"}, "Zane": []string{"Mateo"}, "Zara": []string{"Mohammed"}},
		personA:        "Lucia",
		personB:        "Jun",
		expectedDegree: 14,
		expectedOk:     true,
	},
	{
		description:    "Complex graph, some shortcuts, cross-down and cross-up, cousins several times removed, with unrelated family tree",
		tree:           map[string][]string{"Aditi": []string{"Nia"}, "Aiko": []string{"Bao", "Carlos"}, "Bao": []string{"Dalia"}, "Boris": []string{"Oscar"}, "Carlos": []string{"Fatima", "Gustavo"}, "Celine": []string{"Priya"}, "Dalia": []string{"Hassan", "Isla"}, "Diego": []string{"Qi"}, "Elif": []string{"Rami"}, "Farah": []string{"Sven"}, "Fatima": []string{"Khadija", "Liam"}, "Giorgio": []string{"Tomoko"}, "Gustavo": []string{"Mina"}, "Hana": []string{"Umar"}, "Hassan": []string{"Noah", "Olga"}, "Ian": []string{"Vera"}, "Isla": []string{"Pedro"}, "Javier": []string{"Quynh", "Ravi"}, "Jing": []string{"Wyatt"}, "Kaito": []string{"Xia"}, "Khadija": []string{"Sofia"}, "Leila": []string{"Yassin"}, "Liam": []string{"Tariq", "Uma"}, "Mateo": []string{"Zara"}, "Mina": []string{"Viktor", "Wang"}, "Nia": []string{"Antonio"}, "Noah": []string{"Xiomara"}, "Olga": []string{"Yuki"}, "Oscar": []string{"Bianca"}, "Pedro": []string{"Zane", "Aditi"}, "Priya": []string{"Cai"}, "Qi": []string{"Dimitri"}, "Quynh": []string{"Boris"}, "Rami": []string{"Ewa"}, "Ravi": []string{"Celine"}, "Sofia": []string{"Diego", "Elif"}, "Sven": []string{"Fabio"}, "Tariq": []string{"Farah"}, "Tomoko": []string{"Gabriela"}, "Uma": []string{"Giorgio"}, "Umar": []string{"Helena"}, "Vera": []string{"Igor"}, "Viktor": []string{"Hana", "Ian"}, "Wang": []string{"Jing"}, "Wyatt": []string{"Jun"}, "Xia": []string{"Kim"}, "Xiomara": []string{"Kaito"}, "Yassin": []string{"Lucia"}, "Yuki": []string{"Leila"}, "Zane": []string{"Mateo"}, "Zara": []string{"Mohammed"}},
		personA:        "Wyatt",
		personB:        "Xia",
		expectedDegree: 12,
		expectedOk:     true,
	},
}
