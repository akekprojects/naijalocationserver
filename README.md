# Naija Location Server
![code coverage badge](https://github.com/muhammadolammi/naijalocationserver/actions/workflows/ci.yml/badge.svg)

This is the Repository for http://naijalocationserver.com/, an api endpoint that provides the necesary infomation about Nigeria locations, ranging from States, Capitals, Local Government Areas and more to come.

## Motivation Behind NaijaLocationSever
The development of NaijaLocationSever is guided by four key principles:

* Correctness: This project is personally managed by me, and information will be consistently updated to reflect the latest data.

* Simplicity: The process is straightforward; all that's required is to send an HTTP GET request to one of our endpoints. (Details on how to do this will be explained in the next section.)

* Accessibility: The server's primary goal is to make information easily accessible for use. We handle the heavy lifting in terms of data sourcing, and all you need to do is use HttpClient.Get() on your desired endpoint.

* Performance: Being a Go application, I chose this language for its exceptional performance and speed. If you grasp concepts like memory management and concurrency, you'll appreciate the efficiency of this server.

## 🚀 Quick Start

Navigate to http://naijalocationserver.com/api/states to access a list of all Nigerian states and their capitals in "application/json" format.


## Examples
### Note: All my examples will be in Golang, You can use any http Client From any Language
* Get All States 

```
func getFromEndpoint() interface{} {
	httpClient := http.Client{
		Timeout: time.Second * 10, // Maximum of 10 secs

	}
	resp, err := httpClient.Get(" http://naijalocationserver.com/api/states")
	if err != nil {
		log.Printf("error getting resp from endpoint: %v", err)
		return "error getting resp from endpoint"
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	var bodyData interface{}
	//this is a bad practice, you should create a struct for body data to be marshalled into that correlate with the json data from endpoint. To manage type better.
	err = decoder.Decode(&bodyData)
	if err != nil {
		log.Printf("error creating decoder from resp.Body(): %v", err)
		return "error creating decoder from resp.Body()"
	}

	return bodyData

}

```
### The code here will return the following json
```
[
{
"id": "9e29eb1c-e724-453e-99fe-f58ef3b523e6",
"name": "Abia",
"capital": "Umuahia"
},
{
"id": "9cbb384d-5a18-40f4-88f5-4abef4391c18",
"name": "Adamawa",
"capital": "Yola"
},
{
"id": "c976340a-9000-4f92-a0a4-79acbb9c4688",
"name": "Akwa Ibom",
"capital": "Uyo"
},
{
"id": "3e111a57-e2c9-4c39-9dea-85649061ecdb",
"name": "Anambra",
"capital": "Awka"
},
{
"id": "b91009ad-9a64-4013-b393-4bbdb61a1ee7",
"name": "Bauchi",
"capital": "Bauchi"
},
{
"id": "8a6d6e71-8537-4371-b4a4-26a5fcb47cee",
"name": "Bayelsa",
"capital": "Yenagoa"
},
{
"id": "13b491d2-9513-4108-9e83-f68a0d7f9024",
"name": "Benue",
"capital": "Makurdi"
},
{
"id": "170aaed7-bbb9-4789-8bc9-f57297cdd3c8",
"name": "Borno",
"capital": "Maiduguri"
},
{
"id": "15a2a316-1a29-470c-9e1a-0b9f14ac4fd0",
"name": "Cross River",
"capital": "Calabar"
},
{
"id": "ab3aaebf-3e9b-492d-bc37-bd548baa9723",
"name": "Delta",
"capital": "Asaba"
},
{
"id": "7223111d-4c2b-40f5-99b0-73a504e981ea",
"name": "Ebonyi",
"capital": "Abakaliki"
},
{
"id": "0c6e1c11-9546-4ea6-b71e-4f8140a2326b",
"name": "Edo",
"capital": "Benin City"
},
{
"id": "6dc47d8d-154f-4879-bd73-a0e1307878db",
"name": "Ekiti",
"capital": "Ado Ekiti"
},
{
"id": "b9465e1f-ae53-4cbb-a3e9-dba8b4fb3789",
"name": "Enugu",
"capital": "Enugu"
},
{
"id": "c1a49e84-e090-4d2a-8be3-03fb3dfd8d72",
"name": "Gombe",
"capital": "Gombe"
},
{
"id": "c6c24049-8415-451d-875f-4c9970b7b4b0",
"name": "Imo",
"capital": "Owerri"
},
{
"id": "ce32a7ec-e391-4664-a5d0-cb113a54115d",
"name": "Jigawa",
"capital": "Dutse"
},
{
"id": "8e8f207d-3b6d-4aed-bd1c-4a1a803b502c",
"name": "Kaduna",
"capital": "Kaduna"
},
{
"id": "01c8554b-1ece-4876-a4f5-84a49d665c07",
"name": "Kano",
"capital": "Kano"
},
{
"id": "a7dc2546-54e5-4088-9123-8d623e1dbe95",
"name": "Katsina",
"capital": "Kastina"
},
{
"id": "00109792-575c-423a-a840-9d66ef0075a4",
"name": "Kebbi",
"capital": "Birnin Kebbi"
},
{
"id": "a6837d6d-57e6-4570-84b5-cef2370bf6e9",
"name": "Kogi",
"capital": "Lokoja"
},
{
"id": "d0e04cc0-10ec-408e-906e-ee3d4e3e2d22",
"name": "Kwara",
"capital": "Ilorin"
},
{
"id": "7bf63e43-7d1e-4608-ac36-4342ceeb654a",
"name": "Lagos",
"capital": "Ikeja"
},
{
"id": "499af18e-66cd-4696-b40e-cbc0e2afea63",
"name": "Nasarawa",
"capital": "Lafia"
},
{
"id": "4efa6883-0546-48fa-809d-22f10863219a",
"name": "Niger",
"capital": "Minna"
},
{
"id": "227336ee-b742-41e5-9e92-f1c50037a53c",
"name": "Ogun",
"capital": "Abeokuta"
},
{
"id": "d06258db-261a-4f65-8fa4-81426d37d233",
"name": "Ondo",
"capital": "Akure"
},
{
"id": "bb28069a-b1f2-4c32-92f6-f1d737f00911",
"name": "Osun",
"capital": "Oshogbo"
},
{
"id": "3595705f-d4ef-4083-ad56-84c31729a0dc",
"name": "Oyo",
"capital": "Ibadan"
},
{
"id": "045e1f79-7714-48b6-82e5-9feada7e8689",
"name": "Plateau",
"capital": "Jos"
},
{
"id": "e6362bf4-fbf0-4bb1-a0cd-d86889bb0864",
"name": "Rivers",
"capital": "Port Harcourt"
},
{
"id": "70f42e1b-8745-4279-b1a3-cdc2b1349156",
"name": "Sokoto",
"capital": "Sokoto"
},
{
"id": "878ff196-607e-4ef2-ab56-23528bcb13f4",
"name": "Taraba",
"capital": "Jalingo"
},
{
"id": "0ecfc484-19cc-478c-aa71-9ad235f000e9",
"name": "Yobe",
"capital": "Damaturu"
},
{
"id": "94bc574e-410d-4754-b4eb-ef96a3677daf",
"name": "Zamfara",
"capital": "Gusau"
},
{
"id": "a927afbe-fbc5-43d2-9791-0cd21c721665",
"name": "FCT",
"capital": "Abuja"
}
]
```

* Get All States with Leading Prefixes (eg, k, kw, ka) .

```
func getFromEndpoint() interface{} {
	httpClient := http.Client{
		Timeout: time.Second * 10, // Maximum of 10 secs

	}
	resp, err := httpClient.Get(" http://naijalocationserver.com/api/states/k")
	if err != nil {
		log.Printf("error getting resp from endpoint: %v", err)
		return "error getting resp from endpoint"
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	var bodyData interface{}
	//this is a bad practice, you should create a struct for body data to be marshalled into that correlate with the json data from endpoint. To manage type better.
	err = decoder.Decode(&bodyData)
	if err != nil {
		log.Printf("error creating decoder from resp.Body(): %v", err)
		return "error creating decoder from resp.Body()"
	}

	return bodyData

}

```
### The code here will return the following json

```
[{"id":"8e8f207d-3b6d-4aed-bd1c-4a1a803b502c","name":"Kaduna","capital":"Kaduna"},{"id":"01c8554b-1ece-4876-a4f5-84a49d665c07","name":"Kano","capital":"Kano"},{"id":"a7dc2546-54e5-4088-9123-8d623e1dbe95","name":"Katsina","capital":"Kastina"},{"id":"00109792-575c-423a-a840-9d66ef0075a4","name":"Kebbi","capital":"Birnin Kebbi"},{"id":"a6837d6d-57e6-4570-84b5-cef2370bf6e9","name":"Kogi","capital":"Lokoja"},{"id":"d0e04cc0-10ec-408e-906e-ee3d4e3e2d22","name":"Kwara","capital":"Ilorin"}]
```
* Get A State (by specifying the state name) endpoint : "http://naijalocationserver.com/api/state/{stateNAME}"

```
func getFromEndpoint() interface{} {
	httpClient := http.Client{
		Timeout: time.Second * 10, // Maximum of 10 secs

	}
	resp, err := httpClient.Get("http://naijalocationserver.com/api/state/kwara")
	if err != nil {
		log.Printf("error getting resp from endpoint: %v", err)
		return "error getting resp from endpoint"
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	var bodyData interface{}
	//this is a bad practice, you should create a struct for body data to be marshalled into that correlate with the json data from endpoint. To manage type better.
	err = decoder.Decode(&bodyData)
	if err != nil {
		log.Printf("error creating decoder from resp.Body(): %v", err)
		return "error creating decoder from resp.Body()"
	}

	return bodyData

}

```
### The code here will return the following json
```
{"id":"d0e04cc0-10ec-408e-906e-ee3d4e3e2d22","name":"Kwara","capital":"Ilorin"}
```

* Get All Lgas  endpoint : "http://naijalocationserver.com/api/local-government-areas"

```
func getFromEndpoint() interface{} {
	httpClient := http.Client{
		Timeout: time.Second * 10, // Maximum of 10 secs

	}
	resp, err := httpClient.Get("http://naijalocationserver.com/api/local-government-areas")
	if err != nil {
		log.Printf("error getting resp from endpoint: %v", err)
		return "error getting resp from endpoint"
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	var bodyData interface{}
	//this is a bad practice, you should create a struct for body data to be marshalled into that correlate with the json data from endpoint. To manage type better.
	err = decoder.Decode(&bodyData)
	if err != nil {
		log.Printf("error creating decoder from resp.Body(): %v", err)
		return "error creating decoder from resp.Body()"
	}

	return bodyData

}

```
### The code here will return the following json
```
{
"Abia": [
"Aba North",
"Aba South",
"Arochukwu",
"Bende",
"Ikwuano",
"Isiala-Ngwa North",
"Isiala-Ngwa South",
"Isuikwato",
"Obi Nwa",
"Ohafia",
"Osisioma",
"Ngwa",
"Ugwunagbo",
"Ukwa East",
"Ukwa West",
"Umuahia North",
"Umuahia South",
"Umu-Neochi"
],
"Adamawa": [
"Demsa",
"Fufore",
"Ganaye",
"Gireri",
"Gombi",
"Guyuk",
"Hong",
"Jada",
"Lamurde",
"Madagali",
"Maiha",
"Mayo-Belwa",
"Michika",
"Mubi North",
"Mubi South",
"Numan",
"Shelleng",
"Song",
"Toungo",
"Yola North",
"Yola South"
],
"Akwa Ibom": [
"Abak",
"Eastern Obolo",
"Eket",
"Esit Eket",
"Essien Udim",
"Etim Ekpo",
"Etinan",
"Ibeno",
"Ibesikpo Asutan",
"Ibiono Ibom",
"Ika",
"Ikono",
"Ikot Abasi",
"Ikot Ekpene",
"Ini",
"Itu",
"Mbo",
"Mkpat Enin",
"Nsit Atai",
"Nsit Ibom",
"Nsit Ubium",
"Obot Akara",
"Okobo",
"Onna",
"Oron",
"Oruk Anam",
"Udung Uko",
"Ukanafun",
"Uruan",
"Urue-Offong/Oruko ",
"Uyo"
],
"Anambra": [
"Aguata",
"Anambra East",
"Anambra West",
"Anaocha",
"Awka North",
"Awka South",
"Ayamelum",
"Dunukofia",
"Ekwusigo",
"Idemili North",
"Idemili south",
"Ihiala",
"Njikoka",
"Nnewi North",
"Nnewi South",
"Ogbaru",
"Onitsha North",
"Onitsha South",
"Orumba North",
"Orumba South",
"Oyi"
],
"Bauchi": [
"Alkaleri",
"Bauchi",
"Bogoro",
"Damban",
"Darazo",
"Dass",
"Ganjuwa",
"Giade",
"Itas/Gadau",
"Jama'are",
"Katagum",
"Kirfi",
"Misau",
"Ningi",
"Shira",
"Tafawa-Balewa",
"Toro",
"Warji",
"Zaki"
],
"Bayelsa": [
"Brass",
"Ekeremor",
"Kolokuma/Opokuma",
"Nembe",
"Ogbia",
"Sagbama",
"Southern Jaw",
"Yenegoa"
],
"Benue": [
"Ado",
"Agatu",
"Apa",
"Buruku",
"Gboko",
"Guma",
"Gwer East",
"Gwer West",
"Katsina-Ala",
"Konshisha",
"Kwande",
"Logo",
"Makurdi",
"Obi",
"Ogbadibo",
"Oju",
"Okpokwu",
"Ohimini",
"Oturkpo",
"Tarka",
"Ukum",
"Ushongo",
"Vandeikya"
],
"Borno": [
"Abadam",
"Askira/Uba",
"Bama",
"Bayo",
"Biu",
"Chibok",
"Damboa",
"Dikwa",
"Gubio",
"Guzamala",
"Gwoza",
"Hawul",
"Jere",
"Kaga",
"Kala/Balge",
"Konduga",
"Kukawa",
"Kwaya Kusar",
"Mafa",
"Magumeri",
"Maiduguri",
"Marte",
"Mobbar",
"Monguno",
"Ngala",
"Nganzai",
"Shani"
],
"Cross River": [
"Akpabuyo",
"Odukpani",
"Akamkpa",
"Biase",
"Abi",
"Ikom",
"Yarkur",
"Odubra",
"Boki",
"Ogoja",
"Yala",
"Obanliku",
"Obudu",
"Calabar South",
"Etung",
"Bekwara",
"Bakassi",
"Calabar Municipality"
],
"Delta": [
"Oshimili",
"Aniocha",
"Aniocha South",
"Ika South",
"Ika North-East",
"Ndokwa West",
"Ndokwa East",
"Isoko south",
"Isoko North",
"Bomadi",
"Burutu",
"Ughelli South",
"Ughelli North",
"Ethiope West",
"Ethiope East",
"Sapele",
"Okpe",
"Warri North",
"Warri South",
"Uvwie",
"Udu",
"Warri Central",
"Ukwani",
"Oshimili North",
"Patani"
],
"Ebonyi": [
"Edda",
"Afikpo",
"Onicha",
"Ohaozara",
"Abakaliki",
"Ishielu",
"lkwo",
"Ezza",
"Ezza South",
"Ohaukwu",
"Ebonyi",
"Ivo"
],
"Edo": [
"Esan North-East",
"Esan Central",
"Esan West",
"Egor",
"Ukpoba",
"Central",
"Etsako Central",
"Igueben",
"Oredo",
"Ovia SouthWest",
"Ovia South-East",
"Orhionwon",
"Uhunmwonde",
"Etsako East",
"Esan South-East"
],
"Ekiti": [
"Ado",
"Ekiti-East",
"Ekiti-West",
"Emure/Ise/Orun",
"Ekiti South-West",
"Ikere",
"Irepodun",
"Ijero,",
"Ido/Osi",
"Oye",
"Ikole",
"Moba",
"Gbonyin",
"Efon",
"Ise/Orun",
"Ilejemeje."
],
"Enugu": [
"Enugu South,",
"Igbo-Eze South",
"Enugu North",
"Nkanu",
"Udi Agwu",
"Oji-River",
"Ezeagu",
"IgboEze North",
"Isi-Uzo",
"Nsukka",
"Igbo-Ekiti",
"Uzo-Uwani",
"Enugu Eas",
"Aninri",
"Nkanu East",
"Udenu."
],
"FCT": [
"Abaji",
"Abuja Municipal",
"Bwari",
"Gwagwalada",
"Kuje",
"Kwali"
],
"Gombe": [
"Akko",
"Balanga",
"Billiri",
"Dukku",
"Kaltungo",
"Kwami",
"Shomgom",
"Funakaye",
"Gombe",
"Nafada/Bajoga",
"Yamaltu/Delta."
],
"Imo": [
"Aboh-Mbaise",
"Ahiazu-Mbaise",
"Ehime-Mbano",
"Ezinihitte",
"Ideato North",
"Ideato South",
"Ihitte/Uboma",
"Ikeduru",
"Isiala Mbano",
"Isu",
"Mbaitoli",
"Mbaitoli",
"Ngor-Okpala",
"Njaba",
"Nwangele",
"Nkwerre",
"Obowo",
"Oguta",
"Ohaji/Egbema",
"Okigwe",
"Orlu",
"Orsu",
"Oru East",
"Oru West",
"Owerri-Municipal",
"Owerri North",
"Owerri West"
],
"Jigawa": [
"Auyo",
"Babura",
"Birni Kudu",
"Biriniwa",
"Buji",
"Dutse",
"Gagarawa",
"Garki",
"Gumel",
"Guri",
"Gwaram",
"Gwiwa",
"Hadejia",
"Jahun",
"Kafin Hausa",
"Kaugama Kazaure",
"Kiri Kasamma",
"Kiyawa",
"Maigatari",
"Malam Madori",
"Miga",
"Ringim",
"Roni",
"Sule-Tankarkar",
"Taura",
"Yankwashi"
],
"Kaduna": [
"Birni-Gwari",
"Chikun",
"Giwa",
"Igabi",
"Ikara",
"jaba",
"Jema'a",
"Kachia",
"Kaduna North",
"Kaduna South",
"Kagarko",
"Kajuru",
"Kaura",
"Kauru",
"Kubau",
"Kudan",
"Lere",
"Makarfi",
"Sabon-Gari",
"Sanga",
"Soba",
"Zango-Kataf",
"Zaria"
],
"Kano": [
"Ajingi",
"Albasu",
"Bagwai",
"Bebeji",
"Bichi",
"Bunkure",
"Dala",
"Dambatta",
"Dawakin Kudu",
"Dawakin Tofa",
"Doguwa",
"Fagge",
"Gabasawa",
"Garko",
"Garum",
"Mallam",
"Gaya",
"Gezawa",
"Gwale",
"Gwarzo",
"Kabo",
"Kano Municipal",
"Karaye",
"Kibiya",
"Kiru",
"kumbotso",
"Ghari",
"Kura",
"Madobi",
"Makoda",
"Minjibir",
"Nasarawa",
"Rano",
"Rimin Gado",
"Rogo",
"Shanono",
"Sumaila",
"Takali",
"Tarauni",
"Tofa",
"Tsanyawa",
"Tudun Wada",
"Ungogo",
"Warawa",
"Wudil"
],
"Katsina": [
"Bakori",
"Batagarawa",
"Batsari",
"Baure",
"Bindawa",
"Charanchi",
"Dandume",
"Danja",
"Dan Musa",
"Daura",
"Dutsi",
"Dutsin-Ma",
"Faskari",
"Funtua",
"Ingawa",
"Jibia",
"Kafur",
"Kaita",
"Kankara",
"Kankia",
"Katsina",
"Kurfi",
"Kusada",
"Mai'Adua",
"Malumfashi",
"Mani",
"Mashi",
"Matazuu",
"Musawa",
"Rimi",
"Sabuwa",
"Safana",
"Sandamu",
"Zango"
],
"Kebbi": [
"Aleiro",
"Arewa-Dandi",
"Argungu",
"Augie",
"Bagudo",
"Birnin Kebbi",
"Bunza",
"Dandi",
"Fakai",
"Gwandu",
"Jega",
"Kalgo",
"Koko/Besse",
"Maiyama",
"Ngaski",
"Sakaba",
"Shanga",
"Suru",
"Wasagu/Danko",
"Yauri",
"Zuru"
],
"Kogi": [
"Adavi",
"Ajaokuta",
"Ankpa",
"Bassa",
"Dekina",
"Ibaji",
"Idah",
"Igalamela-Odolu",
"Ijumu",
"Kabba/Bunu",
"Kogi",
"Lokoja",
"Mopa-Muro",
"Ofu",
"Ogori/Mangongo",
"Okehi",
"Okene",
"Olamabolo",
"Omala",
"Yagba East",
"Yagba West"
],
"Kwara": [
"Asa",
"Baruten",
"Edu",
"Ekiti",
"Ifelodun",
"Ilorin East",
"Ilorin West",
"Irepodun",
"Isin",
"Kaiama",
"Moro",
"Offa",
"Oke-Ero",
"Oyun",
"Pategi"
],
"Lagos": [
"Agege",
"Ajeromi-Ifelodun",
"Alimosho",
"Amuwo-Odofin",
"Apapa",
"Badagry",
"Epe",
"Eti-Osa",
"Ibeju/Lekki",
"Ifako-Ijaye",
"Ikeja",
"Ikorodu",
"Kosofe",
"Lagos Island",
"Lagos Mainland",
"Mushin",
"Ojo",
"Oshodi-Isolo",
"Shomolu",
"Surulere"
],
"Nasarawa": [
"Akwanga",
"Awe",
"Doma",
"Karu",
"Keana",
"Keffi",
"Kokona",
"Lafia",
"Nasarawa",
"Nasarawa-Eggon",
"Obi",
"Toto",
"Wamba"
],
"Niger": [
"Agaie",
"Agwara",
"Bida",
"Borgu",
"Bosso",
"Chanchaga",
"Edati",
"Gbako",
"Gurara",
"Katcha",
"Kontagora",
"Lapai",
"Lavun",
"Magama",
"Mariga",
"Mashegu",
"Mokwa",
"Muya",
"Pailoro",
"Rafi",
"Rijau",
"Shiroro",
"Suleja",
"Tafa",
"Wushishi"
],
"Ogun": [
"Abeokuta North",
"Abeokuta South",
"Ado-Odo/Ota",
"Yewa North",
"Yewa South",
"Ewekoro",
"Ifo",
"Ijebu East",
"Ijebu North",
"Ijebu North East",
"Ijebu Ode",
"Ikenne",
"Imeko-Afon",
"Ipokia",
"Obafemi-Owode",
"Ogun Waterside",
"Odeda",
"Odogbolu",
"Remo North",
"Shagamu"
],
"Ondo": [
"Akoko North East",
"Akoko North West",
"Akoko South Akure East",
"Akoko South West",
"Akure North",
"Akure South",
"Ese-Odo",
"Idanre",
"Ifedore",
"Ilaje",
"Ile-Oluji",
"Okeigbo",
"Irele",
"Odigbo",
"Okitipupa",
"Ondo East",
"Ondo West",
"Ose",
"Owo"
],
"Osun": [
"Aiyedade",
"Aiyedire",
"Atakumosa East",
"Atakumosa West",
"Boluwaduro",
"Boripe",
"Ede North",
"Ede South",
"Egbedore",
"Ejigbo",
"Ife Central",
"Ife East",
"Ife North",
"Ife South",
"Ifedayo",
"Ifelodun",
"Ila",
"Ilesha East",
"Ilesha West",
"Irepodun",
"Irewole",
"Isokan",
"Iwo",
"Obokun",
"Odo-Otin",
"Ola-Oluwa",
"Olorunda",
"Oriade",
"Orolu",
"Osogbo"
],
"Oyo": [
"Afijio",
"Akinyele",
"Atiba",
"Atisbo",
"Egbeda",
"Ibadan Central",
"Ibadan North",
"Ibadan North West",
"Ibadan South East",
"Ibadan South West",
"Ibarapa Central",
"Ibarapa East",
"Ibarapa North",
"Ido",
"Irepo",
"Iseyin",
"Itesiwaju",
"Iwajowa",
"Kajola",
"Lagelu Ogbomosho North",
"Ogbomosho South",
"Ogo Oluwa",
"Olorunsogo",
"Oluyole",
"Ona-Ara",
"Orelope",
"Ori Ire",
"Oyo East",
"Oyo West",
"Saki East",
"Saki West",
"Surulere"
],
"Plateau": [
"Barikin Ladi",
"Bassa",
"Bokkos",
"Jos East",
"Jos North",
"Jos South",
"Kanam",
"Kanke",
"Langtang North",
"Langtang South",
"Mangu",
"Mikang",
"Pankshin",
"Qua'an Pan",
"Riyom",
"Shendam",
"Wase"
],
"Rivers": [
"Abua/Odual",
"Ahoada East",
"Ahoada West",
"Akuku Toru",
"Andoni",
"Asari-Toru",
"Bonny",
"Degema",
"Emohua",
"Eleme",
"Etche",
"Gokana",
"Ikwerre",
"Khana",
"Obio/Akpor",
"Ogba/Egbema/Ndoni",
"Ogu/Bolo",
"Okrika",
"Omumma",
"Opobo/Nkoro",
"Oyigbo",
"Port-Harcourt",
"Tai"
],
"Sokoto": [
"Binji",
"Bodinga",
"Dange-shnsi",
"Gada",
"Goronyo",
"Gudu",
"Gawabawa",
"Illela",
"Isa",
"Kware",
"kebbe",
"Rabah",
"Sabon birni",
"Shagari",
"Silame",
"Sokoto North",
"Sokoto South",
"Tambuwal",
"Tqngaza",
"Tureta",
"Wamako",
"Wurno",
"Yabo"
],
"Taraba": [
"Ardo-kola",
"Bali",
"Donga",
"Gashaka",
"Cassol",
"Ibi",
"Jalingo",
"Karin-Lamido",
"Kurmi",
"Lau",
"Sardauna",
"Takum",
"Ussa",
"Wukari",
"Yorro",
"Zing"
],
"Yobe": [
"Bade",
"Bursari",
"Damaturu",
"Fika",
"Fune",
"Geidam",
"Gujba",
"Gulani",
"Jakusko",
"Karasuwa",
"Karawa",
"Machina",
"Nangere",
"Nguru Potiskum",
"Tarmua",
"Yunusari",
"Yusufari"
],
"Zamfara": [
"Anka",
"Bakura",
"Birnin Magaji",
"Bukkuyum",
"Bungudu",
"Gummi",
"Gusau",
"Kaura",
"Namoda",
"Maradun",
"Maru",
"Shinkafi",
"Talata Mafara",
"Tsafe",
"Zurmi"
]
}
```


* Get All Lgas For States Prefixe(eg k, ka , kwa) endpoint : "http://naijalocationserver.com/api/local-government-areas/{statePrefix}"

```
func getFromEndpoint() interface{} {
	httpClient := http.Client{
		Timeout: time.Second * 10, // Maximum of 10 secs

	}
	resp, err := httpClient.Get("http://naijalocationserver.com/api/local-government-areas/k")
	if err != nil {
		log.Printf("error getting resp from endpoint: %v", err)
		return "error getting resp from endpoint"
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	var bodyData interface{}
	//this is a bad practice, you should create a struct for body data to be marshalled into that correlate with the json data from endpoint. To manage type better.
	err = decoder.Decode(&bodyData)
	if err != nil {
		log.Printf("error creating decoder from resp.Body(): %v", err)
		return "error creating decoder from resp.Body()"
	}

	return bodyData

}

```
### The code here will return the following json

```
{

"Kaduna": [
"Birni-Gwari",
"Chikun",
"Giwa",
"Igabi",
"Ikara",
"jaba",
"Jema'a",
"Kachia",
"Kaduna North",
"Kaduna South",
"Kagarko",
"Kajuru",
"Kaura",
"Kauru",
"Kubau",
"Kudan",
"Lere",
"Makarfi",
"Sabon-Gari",
"Sanga",
"Soba",
"Zango-Kataf",
"Zaria"
],
"Kano": [
"Ajingi",
"Albasu",
"Bagwai",
"Bebeji",
"Bichi",
"Bunkure",
"Dala",
"Dambatta",
"Dawakin Kudu",
"Dawakin Tofa",
"Doguwa",
"Fagge",
"Gabasawa",
"Garko",
"Garum",
"Mallam",
"Gaya",
"Gezawa",
"Gwale",
"Gwarzo",
"Kabo",
"Kano Municipal",
"Karaye",
"Kibiya",
"Kiru",
"kumbotso",
"Ghari",
"Kura",
"Madobi",
"Makoda",
"Minjibir",
"Nasarawa",
"Rano",
"Rimin Gado",
"Rogo",
"Shanono",
"Sumaila",
"Takali",
"Tarauni",
"Tofa",
"Tsanyawa",
"Tudun Wada",
"Ungogo",
"Warawa",
"Wudil"
],
"Katsina": [
"Bakori",
"Batagarawa",
"Batsari",
"Baure",
"Bindawa",
"Charanchi",
"Dandume",
"Danja",
"Dan Musa",
"Daura",
"Dutsi",
"Dutsin-Ma",
"Faskari",
"Funtua",
"Ingawa",
"Jibia",
"Kafur",
"Kaita",
"Kankara",
"Kankia",
"Katsina",
"Kurfi",
"Kusada",
"Mai'Adua",
"Malumfashi",
"Mani",
"Mashi",
"Matazuu",
"Musawa",
"Rimi",
"Sabuwa",
"Safana",
"Sandamu",
"Zango"
],
"Kebbi": [
"Aleiro",
"Arewa-Dandi",
"Argungu",
"Augie",
"Bagudo",
"Birnin Kebbi",
"Bunza",
"Dandi",
"Fakai",
"Gwandu",
"Jega",
"Kalgo",
"Koko/Besse",
"Maiyama",
"Ngaski",
"Sakaba",
"Shanga",
"Suru",
"Wasagu/Danko",
"Yauri",
"Zuru"
],
"Kogi": [
"Adavi",
"Ajaokuta",
"Ankpa",
"Bassa",
"Dekina",
"Ibaji",
"Idah",
"Igalamela-Odolu",
"Ijumu",
"Kabba/Bunu",
"Kogi",
"Lokoja",
"Mopa-Muro",
"Ofu",
"Ogori/Mangongo",
"Okehi",
"Okene",
"Olamabolo",
"Omala",
"Yagba East",
"Yagba West"
],
"Kwara": [
"Asa",
"Baruten",
"Edu",
"Ekiti",
"Ifelodun",
"Ilorin East",
"Ilorin West",
"Irepodun",
"Isin",
"Kaiama",
"Moro",
"Offa",
"Oke-Ero",
"Oyun",
"Pategi"
]
}

```



* Get All Lgas For State (eg k, ka , kwa) endpoint : "http://naijalocationserver.com/api/local-government-areas/{state}"

```
func getFromEndpoint() interface{} {
	httpClient := http.Client{
		Timeout: time.Second * 10, // Maximum of 10 secs

	}
	resp, err := httpClient.Get("http://naijalocationserver.com/api/local-government-areas/kwara")
	if err != nil {
		log.Printf("error getting resp from endpoint: %v", err)
		return "error getting resp from endpoint"
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	var bodyData interface{}
	//this is a bad practice, you should create a struct for body data to be marshalled into that correlate with the json data from endpoint. To manage type better.
	err = decoder.Decode(&bodyData)
	if err != nil {
		log.Printf("error creating decoder from resp.Body(): %v", err)
		return "error creating decoder from resp.Body()"
	}

	return bodyData

}

```
### The code here will return the following json

```
{
"Kwara": [
"Asa",
"Baruten",
"Edu",
"Ekiti",
"Ifelodun",
"Ilorin East",
"Ilorin West",
"Irepodun",
"Isin",
"Kaiama",
"Moro",
"Offa",
"Oke-Ero",
"Oyun",
"Pategi"
]
}
```