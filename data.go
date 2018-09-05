package main

var inputJSON = []byte(`{
  "nodes": [
    {
      "id": "0"
    },
    {
      "id": "1"
    },
    {
      "id": "2"
    },
    {
      "id": "3"
    },
    {
      "id": "4"
    },
    {
      "id": "5"
    },
    {
      "id": "6"
    },
    {
      "id": "7"
    },
    {
      "id": "8"
    },
    {
      "id": "9"
    },
    {
      "id": "10"
    },
    {
      "id": "11"
    },
    {
      "id": "12"
    },
    {
      "id": "13"
    },
    {
      "id": "14"
    },
    {
      "id": "15"
    },
    {
      "id": "16"
    },
    {
      "id": "17"
    },
    {
      "id": "18"
    },
    {
      "id": "19"
    },
    {
      "id": "20"
    },
    {
      "id": "21"
    },
    {
      "id": "22"
    },
    {
      "id": "23"
    },
    {
      "id": "24"
    },
    {
      "id": "25"
    },
    {
      "id": "26"
    },
    {
      "id": "27"
    },
    {
      "id": "28"
    },
    {
      "id": "29"
    },
    {
      "id": "30"
    },
    {
      "id": "31"
    },
    {
      "id": "32"
    },
    {
      "id": "33"
    },
    {
      "id": "34"
    },
    {
      "id": "35"
    },
    {
      "id": "36"
    },
    {
      "id": "37"
    },
    {
      "id": "38"
    },
    {
      "id": "39"
    },
    {
      "id": "40"
    },
    {
      "id": "41"
    },
    {
      "id": "42"
    },
    {
      "id": "43"
    },
    {
      "id": "44"
    },
    {
      "id": "45"
    },
    {
      "id": "46"
    },
    {
      "id": "47"
    },
    {
      "id": "48"
    },
    {
      "id": "49"
    },
    {
      "id": "50"
    },
    {
      "id": "51"
    },
    {
      "id": "52"
    },
    {
      "id": "53"
    },
    {
      "id": "54"
    },
    {
      "id": "55"
    },
    {
      "id": "56"
    },
    {
      "id": "57"
    },
    {
      "id": "58"
    },
    {
      "id": "59"
    },
    {
      "id": "60"
    },
    {
      "id": "61"
    },
    {
      "id": "62"
    },
    {
      "id": "63"
    },
    {
      "id": "64"
    },
    {
      "id": "65"
    },
    {
      "id": "66"
    },
    {
      "id": "67"
    },
    {
      "id": "68"
    },
    {
      "id": "69"
    },
    {
      "id": "70"
    },
    {
      "id": "71"
    },
    {
      "id": "72"
    },
    {
      "id": "73"
    },
    {
      "id": "74"
    },
    {
      "id": "75"
    },
    {
      "id": "76"
    },
    {
      "id": "77"
    },
    {
      "id": "78"
    },
    {
      "id": "79"
    },
    {
      "id": "80"
    },
    {
      "id": "81"
    },
    {
      "id": "82"
    },
    {
      "id": "83"
    },
    {
      "id": "84"
    },
    {
      "id": "85"
    },
    {
      "id": "86"
    },
    {
      "id": "87"
    },
    {
      "id": "88"
    },
    {
      "id": "89"
    },
    {
      "id": "90"
    },
    {
      "id": "91"
    },
    {
      "id": "92"
    },
    {
      "id": "93"
    },
    {
      "id": "94"
    },
    {
      "id": "95"
    },
    {
      "id": "96"
    },
    {
      "id": "97"
    },
    {
      "id": "98"
    },
    {
      "id": "99"
    }
  ],
  "links": [
    {
      "source": "1",
      "target": "0"
    },
    {
      "source": "2",
      "target": "1"
    },
    {
      "source": "3",
      "target": "2"
    },
    {
      "source": "4",
      "target": "3"
    },
    {
      "source": "5",
      "target": "4"
    },
    {
      "source": "6",
      "target": "5"
    },
    {
      "source": "7",
      "target": "6"
    },
    {
      "source": "8",
      "target": "7"
    },
    {
      "source": "9",
      "target": "8"
    },
    {
      "source": "10",
      "target": "0"
    },
    {
      "source": "11",
      "target": "10"
    },
    {
      "source": "11",
      "target": "1"
    },
    {
      "source": "12",
      "target": "11"
    },
    {
      "source": "12",
      "target": "2"
    },
    {
      "source": "13",
      "target": "12"
    },
    {
      "source": "13",
      "target": "3"
    },
    {
      "source": "14",
      "target": "13"
    },
    {
      "source": "14",
      "target": "4"
    },
    {
      "source": "15",
      "target": "14"
    },
    {
      "source": "15",
      "target": "5"
    },
    {
      "source": "16",
      "target": "15"
    },
    {
      "source": "16",
      "target": "6"
    },
    {
      "source": "17",
      "target": "16"
    },
    {
      "source": "17",
      "target": "7"
    },
    {
      "source": "18",
      "target": "17"
    },
    {
      "source": "18",
      "target": "8"
    },
    {
      "source": "19",
      "target": "18"
    },
    {
      "source": "19",
      "target": "9"
    },
    {
      "source": "20",
      "target": "10"
    },
    {
      "source": "21",
      "target": "20"
    },
    {
      "source": "21",
      "target": "11"
    },
    {
      "source": "22",
      "target": "21"
    },
    {
      "source": "22",
      "target": "12"
    },
    {
      "source": "23",
      "target": "22"
    },
    {
      "source": "23",
      "target": "13"
    },
    {
      "source": "24",
      "target": "23"
    },
    {
      "source": "24",
      "target": "14"
    },
    {
      "source": "25",
      "target": "24"
    },
    {
      "source": "25",
      "target": "15"
    },
    {
      "source": "26",
      "target": "25"
    },
    {
      "source": "26",
      "target": "16"
    },
    {
      "source": "27",
      "target": "26"
    },
    {
      "source": "27",
      "target": "17"
    },
    {
      "source": "28",
      "target": "27"
    },
    {
      "source": "28",
      "target": "18"
    },
    {
      "source": "29",
      "target": "28"
    },
    {
      "source": "29",
      "target": "19"
    },
    {
      "source": "30",
      "target": "20"
    },
    {
      "source": "31",
      "target": "30"
    },
    {
      "source": "31",
      "target": "21"
    },
    {
      "source": "32",
      "target": "31"
    },
    {
      "source": "32",
      "target": "22"
    },
    {
      "source": "33",
      "target": "32"
    },
    {
      "source": "33",
      "target": "23"
    },
    {
      "source": "34",
      "target": "33"
    },
    {
      "source": "34",
      "target": "24"
    },
    {
      "source": "35",
      "target": "34"
    },
    {
      "source": "35",
      "target": "25"
    },
    {
      "source": "36",
      "target": "35"
    },
    {
      "source": "36",
      "target": "26"
    },
    {
      "source": "37",
      "target": "36"
    },
    {
      "source": "37",
      "target": "27"
    },
    {
      "source": "38",
      "target": "37"
    },
    {
      "source": "38",
      "target": "28"
    },
    {
      "source": "39",
      "target": "38"
    },
    {
      "source": "39",
      "target": "29"
    },
    {
      "source": "40",
      "target": "30"
    },
    {
      "source": "41",
      "target": "40"
    },
    {
      "source": "41",
      "target": "31"
    },
    {
      "source": "42",
      "target": "41"
    },
    {
      "source": "42",
      "target": "32"
    },
    {
      "source": "43",
      "target": "42"
    },
    {
      "source": "43",
      "target": "33"
    },
    {
      "source": "44",
      "target": "43"
    },
    {
      "source": "44",
      "target": "34"
    },
    {
      "source": "45",
      "target": "44"
    },
    {
      "source": "45",
      "target": "35"
    },
    {
      "source": "46",
      "target": "45"
    },
    {
      "source": "46",
      "target": "36"
    },
    {
      "source": "47",
      "target": "46"
    },
    {
      "source": "47",
      "target": "37"
    },
    {
      "source": "48",
      "target": "47"
    },
    {
      "source": "48",
      "target": "38"
    },
    {
      "source": "49",
      "target": "48"
    },
    {
      "source": "49",
      "target": "39"
    },
    {
      "source": "50",
      "target": "40"
    },
    {
      "source": "51",
      "target": "50"
    },
    {
      "source": "51",
      "target": "41"
    },
    {
      "source": "52",
      "target": "51"
    },
    {
      "source": "52",
      "target": "42"
    },
    {
      "source": "53",
      "target": "52"
    },
    {
      "source": "53",
      "target": "43"
    },
    {
      "source": "54",
      "target": "53"
    },
    {
      "source": "54",
      "target": "44"
    },
    {
      "source": "55",
      "target": "54"
    },
    {
      "source": "55",
      "target": "45"
    },
    {
      "source": "56",
      "target": "55"
    },
    {
      "source": "56",
      "target": "46"
    },
    {
      "source": "57",
      "target": "56"
    },
    {
      "source": "57",
      "target": "47"
    },
    {
      "source": "58",
      "target": "57"
    },
    {
      "source": "58",
      "target": "48"
    },
    {
      "source": "59",
      "target": "58"
    },
    {
      "source": "59",
      "target": "49"
    },
    {
      "source": "60",
      "target": "50"
    },
    {
      "source": "61",
      "target": "60"
    },
    {
      "source": "61",
      "target": "51"
    },
    {
      "source": "62",
      "target": "61"
    },
    {
      "source": "62",
      "target": "52"
    },
    {
      "source": "63",
      "target": "62"
    },
    {
      "source": "63",
      "target": "53"
    },
    {
      "source": "64",
      "target": "63"
    },
    {
      "source": "64",
      "target": "54"
    },
    {
      "source": "65",
      "target": "64"
    },
    {
      "source": "65",
      "target": "55"
    },
    {
      "source": "66",
      "target": "65"
    },
    {
      "source": "66",
      "target": "56"
    },
    {
      "source": "67",
      "target": "66"
    },
    {
      "source": "67",
      "target": "57"
    },
    {
      "source": "68",
      "target": "67"
    },
    {
      "source": "68",
      "target": "58"
    },
    {
      "source": "69",
      "target": "68"
    },
    {
      "source": "69",
      "target": "59"
    },
    {
      "source": "70",
      "target": "60"
    },
    {
      "source": "71",
      "target": "70"
    },
    {
      "source": "71",
      "target": "61"
    },
    {
      "source": "72",
      "target": "71"
    },
    {
      "source": "72",
      "target": "62"
    },
    {
      "source": "73",
      "target": "72"
    },
    {
      "source": "73",
      "target": "63"
    },
    {
      "source": "74",
      "target": "73"
    },
    {
      "source": "74",
      "target": "64"
    },
    {
      "source": "75",
      "target": "74"
    },
    {
      "source": "75",
      "target": "65"
    },
    {
      "source": "76",
      "target": "75"
    },
    {
      "source": "76",
      "target": "66"
    },
    {
      "source": "77",
      "target": "76"
    },
    {
      "source": "77",
      "target": "67"
    },
    {
      "source": "78",
      "target": "77"
    },
    {
      "source": "78",
      "target": "68"
    },
    {
      "source": "79",
      "target": "78"
    },
    {
      "source": "79",
      "target": "69"
    },
    {
      "source": "80",
      "target": "70"
    },
    {
      "source": "81",
      "target": "80"
    },
    {
      "source": "81",
      "target": "71"
    },
    {
      "source": "82",
      "target": "81"
    },
    {
      "source": "82",
      "target": "72"
    },
    {
      "source": "83",
      "target": "82"
    },
    {
      "source": "83",
      "target": "73"
    },
    {
      "source": "84",
      "target": "83"
    },
    {
      "source": "84",
      "target": "74"
    },
    {
      "source": "85",
      "target": "84"
    },
    {
      "source": "85",
      "target": "75"
    },
    {
      "source": "86",
      "target": "85"
    },
    {
      "source": "86",
      "target": "76"
    },
    {
      "source": "87",
      "target": "86"
    },
    {
      "source": "87",
      "target": "77"
    },
    {
      "source": "88",
      "target": "87"
    },
    {
      "source": "88",
      "target": "78"
    },
    {
      "source": "89",
      "target": "88"
    },
    {
      "source": "89",
      "target": "79"
    },
    {
      "source": "90",
      "target": "80"
    },
    {
      "source": "91",
      "target": "90"
    },
    {
      "source": "91",
      "target": "81"
    },
    {
      "source": "92",
      "target": "91"
    },
    {
      "source": "92",
      "target": "82"
    },
    {
      "source": "93",
      "target": "92"
    },
    {
      "source": "93",
      "target": "83"
    },
    {
      "source": "94",
      "target": "93"
    },
    {
      "source": "94",
      "target": "84"
    },
    {
      "source": "95",
      "target": "94"
    },
    {
      "source": "95",
      "target": "85"
    },
    {
      "source": "96",
      "target": "95"
    },
    {
      "source": "96",
      "target": "86"
    },
    {
      "source": "97",
      "target": "96"
    },
    {
      "source": "97",
      "target": "87"
    },
    {
      "source": "98",
      "target": "97"
    },
    {
      "source": "98",
      "target": "88"
    },
    {
      "source": "99",
      "target": "98"
    },
    {
      "source": "99",
      "target": "89"
    }
  ]
}`)
