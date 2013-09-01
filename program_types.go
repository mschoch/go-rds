//  Copyright (c) Marty Schoch
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package rds

func ProgramTypeByCode(code int) ProgramType {
	if code < 0 || code > 31 {
		code = 0
	}
	return ProgramTypesNorthAmerica[code]
}

type ProgramType struct {
	Type       string
	Display8   string
	Display16  string
	Definition string
}

var ProgramTypesNorthAmerica = []ProgramType{
	{"No program type or undefined", "None", "None", ""},
	{"News", "News", "News", "News reports, either local or network in origin"},
	{"Information", "Inform", "Information", "Programming that is intended to impart advice"},
	{"Sports", "Sports", "Sports", "Sports reporting, commentary, and/or live event coverage, either local or network in origin"},
	{"Talk", "Talk", "Talk", "Call-in and/or interview talk shows either local or national in origin"},
	{"Rock", "Rock", "Rock", "Album cuts"},
	{"Classic Rock", "Cls_Rock", "Classic_Rock", "Rock oriented oldies, often mixed with hit oldies, from a decade or more ago"},
	{"Adult Hits", "Adlt_Hit", "Adult_Hits", "An up-tempo contemporary hits format with no hard rock and no rap"},
	{"Soft Rock", "Soft_Rck", "Soft_Rock", "Album cuts with a generally soft tempo"},
	{"Top 40", "Top_40", "Top_ 40", "Current hits, often encompassing a variety of rock styles"},
	{"Country", "Country", "Country", "Country music, including contemporary and traditional styles"},
	{"Oldies", "Oldies", "Oldies", "Popular music, usually rock, with 80% or greater non-current music"},
	{"Soft", "Soft", "Soft", "A cross between adult hits and classical, primarily non-current soft-rock originals"},
	{"Nostalgia", "Nostalga", "Nostalgia", "Big-band music"},
	{"Jazz", "Jazz", "Jazz", `Mostly instrumental, includes both traditional jazz and more modern "smooth jazz"`},
	{"Classical", "Classicl", "Classical", "Mostly instrumentals, usually orchestral or symphonic music"},
	{"Rhythm and Blues", "R_&_B", "Rhythm_and_Blues", `A wide range of musical styles, often called "urban contemporary"`},
	{"Soft Rhythm and Blues", "Soft_R&B", "Soft_ R_&_B", "Rhythm and blues with a generally soft tempo"},
	{"Foreign Language", "Language", "Foreign_Language", "Any programming format in a language other than English"},
	{"Religious Music", "Rel_Musc", "Religious_Music", "Music programming with religious lyrics"},
	{"Religious Talk", "Rel_Talk", "Religious_Talk", "Call-in shows, interview programs, etc. with a religious theme"},
	{"Personality", "Persnlty", "Personality", "A radio show where the on-air personality is the main attraction"},
	{"Public", "Public", "Public", "Programming that is supported by listeners and/or corporate sponsors instead of advertising"},
	{"College", "College", "College", "Programming produced by a college or university radio station"},
	{"Spanish Talk", "Habl_Esp", "Hablar_Espanol", "Call-in shows, interview programs, etc. in the Spanish language"},
	{"Spanish Music", "Musc_Esp", "Musica _Espanol", "Music programming in the Spanish language"},
	{"Hip-Hop", "Hip hop", "Hip hop", "Popular music incorporating elements of rap, rhythm-and-blues, funk, and soul"},
	{"Unassigned", "", "", ""},
	{"Unassigned", "", "", ""},
	{"Weather", "Weather", "Weather", "Weather forecasts or bulletins that are nonemergency in nature"},
	{"Emergency Test", "Test", "Emergency_Test", `Broadcast when testing emergency broadcast equipment or receivers. Not intended for searching or dynamic switching for consumer receivers.. Receivers may, if desired, display “TEST” or “Emergency Test”`},
	{"Emergency", "ALERT !", "ALERT!_ALERT!", "Emergency announcement made under exceptional circumstances to give warning of events causing danger of a general nature. Not to be used for searching - only used in a receiver for dynamic switching"},
}
