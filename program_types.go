//  Copyright (c) Marty Schoch
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package rds

import (
	"fmt"
)

type RDSInfo struct {
	PI                   uint16
	ProgramType          uint16
	PSName               [8]byte
	RadioText            [64]byte
	TrafficProgram       bool
	TrafficAnnouncement  bool
	IsMusic              bool
	IsStereo             bool
	IsArtificialHead     bool
	IsCompressed         bool
	IsDynamicProgramType bool
}

func NewRDSInfo() *RDSInfo {
	return &RDSInfo{
		IsMusic: true,
	}
}

func (rdsinfo *RDSInfo) Update(a, b, c, d uint16) {
	rdsinfo.PI = a
	rdsinfo.ProgramType = b >> 5 & 0x1F

	// traffic program?
	rdsinfo.TrafficProgram = bool(((b >> 10) & 0x1) == 1)

	// further processing dependent on group code
	group := b >> 12
	isGroupB := bool(((b >> 11) & 0x1) == 1)

	// process groupA
	if !isGroupB {
		switch group {
		case 0:
			rdsinfo.TrafficAnnouncement = bool(((b >> 4) & 0x1) == 1)

			di := bool(((b >> 2) & 0x1) == 1)
			ci := b & 0x3
			switch ci {
			case 0:
				rdsinfo.IsDynamicProgramType = di
				rdsinfo.PSName[0] = byte(d >> 8)
				rdsinfo.PSName[1] = byte(d & 0xff)
			case 1:
				rdsinfo.IsCompressed = di
				rdsinfo.PSName[2] = byte(d >> 8)
				rdsinfo.PSName[3] = byte(d & 0xff)
			case 2:
				rdsinfo.IsArtificialHead = di
				rdsinfo.PSName[4] = byte(d >> 8)
				rdsinfo.PSName[5] = byte(d & 0xff)
			case 3:
				rdsinfo.IsStereo = di
				rdsinfo.PSName[6] = byte(d >> 8)
				rdsinfo.PSName[7] = byte(d & 0xff)
			}
		case 2:
			offset := b & 0xf
			rdsinfo.RadioText[offset*4] = byte(c >> 8)
			rdsinfo.RadioText[(offset*4)+1] = byte(c & 0xff)
			rdsinfo.RadioText[(offset*4)+2] = byte(d >> 8)
			rdsinfo.RadioText[(offset*4)+3] = byte(d & 0xff)
		}
	}
}

func (rdsinfo *RDSInfo) String() string {
	rv := ""
	rv += "PI: " + fmt.Sprintf("%d", rdsinfo.PI) + "\n"
	rv += "Program Type: " + ProgramTypeByCode(int(rdsinfo.ProgramType)).Type + "\n"
	rv += "Program Service Name: " + string(rdsinfo.PSName[:]) + "\n"
	rv += "Radio Text:" + string(rdsinfo.RadioText[:]) + "\n"
	rv += fmt.Sprintf("Traffic Program: %t Traffic Announcement: %t Music: %t\n", rdsinfo.TrafficProgram, rdsinfo.TrafficAnnouncement, rdsinfo.IsMusic)
	rv += fmt.Sprintf("Stereo: %t  Artificial Head: %t Compressed: %t Dynamic Program Type: %t", rdsinfo.IsStereo, rdsinfo.IsArtificialHead, rdsinfo.IsCompressed, rdsinfo.IsDynamicProgramType)
	return rv
}

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
