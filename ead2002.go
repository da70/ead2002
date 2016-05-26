//
// Package ead2002 implements a light wrapper for working with ead2002 XML content.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2016, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// * Neither the name of epgo nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package ead2002

import (
	"encoding/xml"
	"time"
)

type EADHeader struct {
	XMLName            xml.Name     `json:"-"`
	Audience           string       `xml:"audience,attr,omitempty" json:"audience,omitempty"`
	CountryEncoding    string       `xml:"countryencoding,attr,omitempty" json:"countryencoding,omitempty"`
	DateEncoding       string       `xml:"dateencoding,attr,omitempty" json:"dateencoding,omitempty"`
	LangEncoding       string       `xml:"langencoding,attr,omitempty" json:"langencoding,omitempty"`
	RelatedEncoding    string       `xml:"relatedencoding,attr,omitempty" json:"relatedencoding,omitempty"`
	RepositoryEncoding string       `xml:"repositoryencoding,attr,omitempty" json:"repositoryencoding,omitempty"`
	ScriptEncoding     string       `xml:"scriptencoding,attr,omitempty" json:"scriptencoding,omitempty"`
	EADID              *EADID       `xml:"eadid,omitempty" json:"eadid,omitempty"`
	FileDesc           *FileDesc    `xml:"filedesc,omitempty" json:"filedesc,omitempty"`
	ProfileDesc        *ProfileDesc `xml:"profiledesc,omitempty" json:"profiledesc,omitempty"`
}

type EADID struct {
	XMLName        xml.Name `json:"-"`
	CountryCode    string   `xml:"countrycode,attr,omitempty" json:"countrycode,omitempty"`
	Identifier     string   `xml:"identifier,attr,omitempty" json:"identifier,omitempty"`
	MainAgencyCode string   `xml:"mainagencycode,attr,omitempty" json:"mainagencycode,omitempty"`
	Value          string   `xml:",innerxml,omitempty" json:"value,omitempty"`
}

type FileDesc struct {
	XMLName         xml.Name         `json:"-"`
	TitleStmt       *TitleStmt       `xml:"titlestmt,omitempty" json:"titlestmt,omitempty"`
	PublicationStmt *PublicationStmt `xml:"publicationstmt,omitempty", json:"publicationstmt,omitempty"`
}

type TitleStmt struct {
	XMLName     xml.Name     `json:"-"`
	TitleProper *TitleProper `xml:"titleproper,omitempty" json:"titleproper,omitempty"`
	Subtitle    string       `xml:"subtitle,omitempty" json:"subtitle,omitempty"`
	Author      *Author      `xml:"author,omitempty" json:"author,omitempty"`
}

type TitleProper struct {
	XMLName        xml.Name `json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Value          string   `xml:",innerxml,omitempty" json:"value,omitempty"`
}

type Author struct {
	XMLName        xml.Name `json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Value          string   `xml:",innerxml,omitempty" json:"value,omitempty"`
}

type PublicationStmt struct {
	XMLName   xml.Name   `json:"-"`
	Publisher *Publisher `xml:"publisher,omitempty" json:"publisher,omitempty"`
	Date      *time.Time `xml:"date,omitempty" json:"date,omitempty"`
}

type Data struct {
	XMLName        xml.Name `json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Normal         string   `xml:"normal,attr,omitempty" json:"normal,omitempty"`
	Value          string   `xml:",innerxml,omitempty" json:"value,omitempty"`
}

type Publisher struct {
	XMLName        xml.Name `json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Value          string   `xml:",innerxml,omitempty" json:"value,omitempty"`
}

type ProfileDesc struct {
	XMLName   xml.Name   `json:"-"`
	Creation  *Creation  `xml:"creation,omitempty" json:"creation,omitempty"`
	LangUsage *LangUsage `xml:"langusage,omitempty" json:"langusage,omitempty"`
}

type Creation struct {
	XMLName xml.Name `json:"-"`
	Value   string   `xml:",innerxml,omitempty" json:"value,omitempty"`
}

type LangUsage struct {
	XMLName  xml.Name     `json:"-"`
	Language LanguageList `xml:"language,omitempty" json:"language,omitempty"`
}

type Language struct {
	XMLName        xml.Name `json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	LangCode       string   `xml:"langcode,attr,omitempty" json:"langcode,omitempty"`
	Value          string   `xml:",innerxml,omitempty" json:"value,omitempty"`
}

type LanguageList []*Language

type ArchDesc struct {
	XMLName         xml.Name `json:"-"`
	Level           string   `xml:"level,attr,omitempty" json:"level,omitempty"`
	Type            string   `xml:"type,attr,omitempty" json:"type,omitempty"`
	RelatedEncoding `xml:"relatedencoding,attr,omitempty" json:"relatedencoding,omitempty"`
	DID             *DID `xml:"did,omitempty" json:"did,omitempty"`
}

type DID struct {
	XMLName      xml.Name      `json:"-"`
	Head         string        `xml:"head,omitempty" json:"omitempty"`
	Repository   *Repository   `xml:"repository,omitempty" json:"repository,omitempty"`
	Organization *Organization `xml:"organization,omitempty" json:"organization,omitempty"`
	UnitTitle    *UnitTitle    `xml:"unittitle,omitempty" json:"unittitle,omitempty"`
	UnitDate     *UnitDate     `xml:"unitdate,omitempty" json:"unitdate,omitempty"`
	PhysDesc     *PhysDesc     `xml:"physdesc,omitempty" json:"physdesc,omitempty"`
	Abstract     *Abstract     `xml:"abstract,omitempty" json:"abstract,omitempty"`
	UnitID       *UnitID       `xml:"unitid,omitempty" json:"unitid,omitempty"`
	LangMaterial *LangMaterial `xml:"langmaterial,omitempty" json:"langmaterial,omitempty"`
}

type Organization struct {
	XMLName  xml.Name  `json:"-"`
	Label    string    `xml:"label,attr,omitempty" json:"label,omitempty"`
	PersName *PersName `xml:"persname,omitempty" json:"persname,omitempty"`
}

type PersName struct {
	XMLName        xml.Name `json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Value          string   `xml:",innerxml,omitempty" json:"value,omitempty"`
}

type UnitTitle struct {
	XMLName        xml.Name `json:"-"`
	Label          string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Value          string   `xml:",innerxml,omitempty" json:"value,omitempty"`
}

type UnitDate struct {
	XMLName        xml.Name `json:"-"`
	Normal         string   `xml:"normal,attr,omitempty" json:"normal,omitempty"`
	Type           string   `xml:"type,attr,omitempty" json:"type,omitempty"`
	Label          string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Value          string   `xml:",innerxml,omitempty" json:"value,omitempty"`
}

type PhysDesc struct {
	XMLName        xml.Name   `json:"-"`
	EncodingAnalog string     `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Label          string     `xml:"label,attr,omitempty" json:"label,omitempty"`
	Extent         ExtentList `xml:"extent,omitempty" json:"extent,omitempty"`
}

type Extent struct {
}

type ExtentList []Extent
