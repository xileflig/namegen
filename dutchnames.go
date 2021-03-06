package namegen

var (
	// https://en.wikipedia.org/wiki/Category:Dutch_masculine_given_names
	// https://en.wikipedia.org/wiki/Category:Dutch_feminine_given_names
	dutchMaleFirstNames = []string{
		"Aad",
		"Aart",
		"Ab",
		"Abraham",
		"Ad",
		"Adam",
		"Adem",
		"Adolf",
		"Adriaan",
		"Adriaen",
		"Adrian",
		"Adrianus",
		"Aernout",
		"Albert",
		"Albertus",
		"Alexander",
		"Alfred",
		"Alvin",
		"Alwin",
		"Andreas",
		"Andries",
		"Anne",
		"Anthonie",
		"Anton",
		"Antonie",
		"Antonius",
		"Antoon",
		"Arend",
		"Arie",
		"Arjan",
		"Arjen",
		"Arno",
		"Arnoldus",
		"Arnoud",
		"Arnout",
		"Arthur",
		"Axel",
		"Barend",
		"Bart",
		"Bartholomeus",
		"Bas",
		"Bastiaan",
		"Bauke",
		"Berend",
		"Bernardus",
		"Bernhard",
		"Bert",
		"Bertus",
		"Bjorn",
		"Boudewijn",
		"Bram",
		"Bruno",
		"Calvin",
		"Carel",
		"Carl",
		"Cees",
		"Christiaan",
		"Christoffel",
		"Claes",
		"Clemens",
		"Coen",
		"Coenraad",
		"Conradus",
		"Cor",
		"Cornelis",
		"Daan",
		"Daniel",
		"David",
		"Derek",
		"Derk",
		"Diederik",
		"Dirk",
		"Dixie",
		"Dolf",
		"Donald",
		"Douwe",
		"Dries",
		"Eduard",
		"Edwin",
		"Eelco",
		"Egbert",
		"Egon",
		"Elbert",
		"Emiel",
		"Eric",
		"Erik",
		"Evert",
		"Ewout",
		"Felix",
		"Florian",
		"Floris",
		"Franciscus",
		"Frank",
		"Frans",
		"Frederik",
		"Frits",
		"Gabriel",
		"Geert",
		"Gerard",
		"Gerardus",
		"Gerbrand",
		"Gerhard",
		"Gerhardus",
		"Gerrit",
		"Gert",
		"Gerwin",
		"Giel",
		"Gijs",
		"Gijsbert",
		"Gilbert",
		"Godfried",
		"Goos",
		"Govert",
		"Gregoor",
		"Guido",
		"Gustav",
		"Guus",
		"Han",
		"Hannes",
		"Hans",
		"Harm",
		"Harmen",
		"Harrie",
		"Hein",
		"Heinke",
		"Helge",
		"Hendrick",
		"Hendrik",
		"Hendrikus",
		"Henk",
		"Hennie",
		"Henricus",
		"Herbert",
		"Herman",
		"Hermanus",
		"Hessel",
		"Hilbert",
		"Hubert",
		"Hubertus",
		"Huib",
		"Huub",
		"IJsbrand",
		"Ivo",
		"Jaap",
		"Jacco",
		"Jack",
		"Jaco",
		"Jacob",
		"Jacobus",
		"Jan",
		"Jannes",
		"Jasper",
		"Jef",
		"Jelle",
		"Jeroen",
		"Jesse",
		"Jo",
		"Joan",
		"Jochem",
		"Joep",
		"Joeri",
		"Johan",
		"Johannes",
		"Jonas",
		"Jonathan",
		"Joop",
		"Joost",
		"Jorien",
		"Joris",
		"Jorrit",
		"Jos",
		"Josephus",
		"Jozef",
		"Jurjen",
		"Karel",
		"Karl",
		"Kees",
		"Klaas",
		"Klaus",
		"Kobus",
		"Koen",
		"Koenraad",
		"Koos",
		"Lambert",
		"Lambertus",
		"Lambrecht",
		"Lammert",
		"Lars",
		"Laurens",
		"Leendert",
		"Leonardus",
		"Leopold",
		"Lex",
		"Lieve",
		"Lieven",
		"Linus",
		"Lodewijk",
		"Ludger",
		"Ludo",
		"Ludovicus",
		"Lukas",
		"Luuk",
		"Maarten",
		"Machiel",
		"Maikel",
		"Manuel",
		"Marcel",
		"Marco",
		"Marijn",
		"Marinus",
		"Marius",
		"Mark",
		"Marko",
		"Marnix",
		"Mart",
		"Marten",
		"Martijn",
		"Martin",
		"Martinus",
		"Mattheus",
		"Matthias",
		"Matthijs",
		"Maurits",
		"Mauritz",
		"Melis",
		"Melvin",
		"Menno",
		"Michiel",
		"Milan",
		"Nanne",
		"Nelis",
		"Nico",
		"Nicolaas",
		"Niek",
		"Nikolaas",
		"Norbert",
		"Norman",
		"Olaf",
		"Olivier",
		"Onno",
		"Otto",
		"Pascal",
		"Paul",
		"Paulus",
		"Pepijn",
		"Peter",
		"Petrus",
		"Philip",
		"Pier",
		"Piet",
		"Pieter",
		"Pim",
		"Quirijn",
		"Ralph",
		"Rein",
		"Reinder",
		"Reindert",
		"Reinhardt",
		"Reinhart",
		"Reinier",
		"Remco",
		"Rens",
		"Reynier",
		"Rien",
		"Rik",
		"Rik",
		"Rinus",
		"Robbert",
		"Robert",
		"Roel",
		"Roeland",
		"Roelof",
		"Rogier",
		"Roland",
		"Rombout",
		"Romeyn",
		"Ronald",
		"Rudolf",
		"Rudolph",
		"Rutger",
		"Ruud",
		"Samuel",
		"Sander",
		"Sebastiaan",
		"Sem",
		"Simon",
		"Sjaak",
		"Sjoerd",
		"Stef",
		"Stefaan",
		"Stefan",
		"Stijn",
		"Sven",
		"Sylvester",
		"Teun",
		"Teunis",
		"Thaddeus",
		"Theo",
		"Theodoor",
		"Theodor",
		"Theodoros",
		"Theunis",
		"Thijmen",
		"Thijs",
		"Thomas",
		"Tielman",
		"Tijs",
		"Tim",
		"Tinus",
		"Tjalling",
		"Tjapko",
		"Tjeerd",
		"Tobias",
		"Tom",
		"Ton",
		"Toon",
		"Twan",
		"Valentijn",
		"Valentin",
		"Vincent",
		"Vincenz",
		"Walraven",
		"Walter",
		"Werner",
		"Wiebe",
		"Wijnand",
		"Wilco",
		"Wilhelmus",
		"Willem",
		"Wim",
		"Wisse",
		"Wolter",
		"Wout",
		"Wouter",
		"Zeger",
	}

	dutchFemaleFirstNames = []string{
		"Ada",
		"Adriana",
		"Agatha",
		"Alberta",
		"Aleida",
		"Aletta",
		"Alida",
		"Amalia",
		"Anna",
		"Anne",
		"Anneke",
		"Anneliese",
		"Annemarie",
		"Annemieke",
		"Annie",
		"Annika",
		"Anouk",
		"Ans",
		"Antje",
		"Antonia",
		"Astrid",
		"Barbara",
		"Carlijn",
		"Carolina",
		"Caroline",
		"Catharina",
		"Charlotte",
		"Claudia",
		"Cobie",
		"Cor",
		"Cornelia",
		"Corrie",
		"Crystal",
		"Cynthia",
		"Danielle",
		"Eefje",
		"Elke",
		"Ellen",
		"Elly",
		"Els",
		"Emily",
		"Emma",
		"Estelle",
		"Esther",
		"Ethel",
		"Eva",
		"Femke",
		"Fleur",
		"Floortje",
		"Frida",
		"Geertruida",
		"Gerarda",
		"Gerda",
		"Gesina",
		"Greet",
		"Grietje",
		"Hanne",
		"Hanneke",
		"Heidi",
		"Heleen",
		"Helga",
		"Hendrika",
		"Hennie",
		"Henriëtte",
		"Hilda",
		"Ilse",
		"Ineke",
		"Inge",
		"Ingrid",
		"Irene",
		"Isabella",
		"Jacoba",
		"Jacqueline",
		"Janneke",
		"Jet",
		"Jo",
		"Johanna",
		"Joke",
		"Jolanda",
		"Jolijn",
		"Jorien",
		"Josephina",
		"Judith",
		"Karen",
		"Karin",
		"Katja",
		"Kristi",
		"Kristin",
		"Lieke",
		"Liesbeth",
		"Lieve",
		"Loes",
		"Lonneke",
		"Lotte",
		"Lowiena",
		"Maaike",
		"Maartje",
		"Manon",
		"Margaretha",
		"Margriet",
		"Maria",
		"Marianne",
		"Marieke",
		"Marijke",
		"Marijn",
		"Mariska",
		"Marja",
		"Marjan",
		"Marjolein",
		"Marleen",
		"Marlies",
		"Marloes",
		"Marta",
		"Meike",
		"Melanie",
		"Merel",
		"Michelle",
		"Mieke",
		"Miranda",
		"Mirjam",
		"Myrthe",
		"Nadine",
		"Nanne",
		"Natalie",
		"Nathalie",
		"Nel",
		"Nienke",
		"Nina",
		"Olga",
		"Olivia",
		"Paula",
		"Paulien",
		"Pearl",
		"Petronella",
		"Petunia",
		"Rachel",
		"Ria",
		"Rie",
		"Sanne",
		"Saskia",
		"Sophie",
		"Suzanne",
		"Tineke",
		"Toos",
		"Truus",
		"Valerie",
		"Victoria",
		"Viola",
		"Wilhelmina",
		"Willeke",
		"Willemina",
		"Wilma",
		"Yvette",
		"Yvonne"}

	// https://en.wikipedia.org/wiki/List_of_Dutch_family_names
	dutchLastNames = []string{
		"baas",
		"bakker",
		"van beek",
		"beenhouwer",
		"van der bijl",
		"bos",
		"van der berg",
		"berkenbosch",
		"de boer",
		"boogaard",
		"boogman",
		"van der boor",
		"boswel",
		"bouwman",
		"braam",
		"brouwer",
		"de bruin",
		"van buskirk",
		"van der byl",
		"van coevorden",
		"citroen",
		"cornelissen",
		"dekker",
		"van dijk",
		"dijkstra",
		"de graaf",
		"de groot",
		"de haan",
		"de haas",
		"van der heide",
		"hendriks",
		"van den heuvel",
		"hoebee",
		"van het hoff",
		"van der kleij",
		"kuiper",
		"van leeuwen",
		"de jaager",
		"jansen",
		"de jong",
		"de koning",
		"de lange",
		"van der linden",
		"meijer",
		"van der meer",
		"mesman",
		"meulenbelt",
		"van der molen",
		"muis",
		"mulder",
		"maarschalkerweerd",
		"peters",
		"prins",
		"ruis",
		"rynsburger",
		"van sittart",
		"smit",
		"spaans",
		"stegenga",
		"teuling",
		"timmerman",
		"tuinstra",
		"vinke",
		"visser",
		"van vliet",
		"de vries",
		"vos",
		"vroom",
		"de wees",
		"van der westhuizen",
		"willems",
		"de wit",
		"van zijl"}
)
