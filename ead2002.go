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
	"encoding/json"
	"encoding/xml"
)

const (
	Version = "0.0.1"
)

type EAD struct {
	XMLName   xml.Name   `json:"-"`
	EADHeader *EADHeader `xml:"eadheader,omitempty" json:"eadheader,omitempty"`
	ArchDesc  *ArchDesc  `xml:"archdesc,omitempty" json:"archdesc,omitempty"`
}

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
	PublicID       string   `xml:"publicid,attr,omitempty" json:"publicid,omitempty"`
	PathParent     string   `xml:"path:parent,attr,omitempty" json:"path_parent,omitempty"`
	Value          string   `xml:",innerxml" json:"value"`
}

type FileDesc struct {
	XMLName         xml.Name         `json:"-"`
	TitleStmt       *TitleStmt       `xml:"titlestmt,omitempty" json:"titlestmt,omitempty"`
	PublicationStmt *PublicationStmt `xml:"publicationstmt,omitempty", json:"publicationstmt,omitempty"`
}

type TitleStmt struct {
	XMLName     xml.Name      `json:"-"`
	TitleProper []TitleProper `xml:"titleproper,omitempty" json:"titleproper,omitempty"`
	//Subtitle    string        `xml:"subtitle,omitempty" json:"subtitle,omitempty"`
	Author *Author `xml:"author,omitempty" json:"author,omitempty"`
}

type TitleProper struct {
	XMLName        xml.Name `json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Type           string   `xml:"type,attr,omitempty" json:"type,omitempty"`
	Value          string   `xml:",innerxml" json:"value"`
}

type Author struct {
	XMLName        xml.Name `json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Value          string   `xml:",chardata" json:"value"`
}

type PublicationStmt struct {
	XMLName   xml.Name      `json:"-"`
	Publisher *Publisher    `xml:"publisher,omitempty" json:"publisher,omitempty"`
	P         *P            `xml:"p,omitempty" json:"p,omitempty"`
	Address   *AddressLines `xml:"address,omitempty" json:"address,omitempty"`
	Date      string        `xml:"date,omitempty" json:"date,omitempty"`
}

type P struct {
	XMLName xml.Name `json:"-"`
	Value   string   `xml:",innerxml" json:"value,omitempty"`
}

type AddressLines struct {
	AddressLine []string `xml:"addressline,omitempty" json:"addressline,omitempty"`
}

type Data struct {
	XMLName        xml.Name `json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Normal         string   `xml:"normal,attr,omitempty" json:"normal,omitempty"`
	Value          string   `xml:",innerxml" json:"value"`
}

type Publisher struct {
	XMLName        xml.Name `json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Value          string   `xml:",innerxml" json:"value,omitempty"`
}

type ProfileDesc struct {
	XMLName   xml.Name   `json:"-"`
	Creation  *Creation  `xml:"creation,omitempty" json:"creation,omitempty"`
	LangUsage *LangUsage `xml:"langusage,omitempty" json:"langusage,omitempty"`
}

type Creation struct {
	XMLName xml.Name `json:"-"`
	Value   string   `xml:",innerxml" json:"value,omitempty"`
}

type LangUsage struct {
	XMLName  xml.Name     `json:"-"`
	Language LanguageList `xml:"language,omitempty" json:"language,omitempty"`
}

type Language struct {
	XMLName        xml.Name `json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	LangCode       string   `xml:"langcode,attr,omitempty" json:"langcode,omitempty"`
	Value          string   `xml:",innerxml" json:"value,omitempty"`
}

type LanguageList []*Language

type ArchDesc struct {
	XMLName         xml.Name        `json:"-"`
	Level           string          `xml:"level,attr,omitempty" json:"level,omitempty"`
	Type            string          `xml:"type,attr,omitempty" json:"type,omitempty"`
	RelatedEncoding string          `xml:"relatedencoding,attr,omitempty" json:"relatedencoding,omitempty"`
	DID             []DID           `xml:"did,omitempty" json:"did,omitempty"`
	AccessRestrict  *AccessRestrict `xml:"accessrestrict,omitempty" json:"accessrestrict,omitempty"`
	UserRestrict    *UserRestrict   `xml:"userrestrict,omitempty" json:"userrestrict,omitempty"`
	Prefercite      *Prefercite     `xml:"prefercite,omitempty" json:"prefercite,omitempty"`
	AcqInfo         *AcqInfo        `xml:"acqinfo,omitempty" json:"acqinfo,omitempty"`
	BiogHist        *BiogHist       `xml:"bioghist,omitempty" json:"bioghist,omitempty"`
	ScopeContent    *ScopeContent   `xml:"scopecontent,omitempty" json:"scopecontent,omitempty"`
	ControlAccess   *ControlAccess  `xml:"controlaccess,omitempty" json:"controlaccess,omitempty"`
	OtherFindAid    *OtherFindAid   `xml:"otherfindaid,omitempty" json:"otherfindaid,omitempty"`
	DSC             *DSC            `xml:"dsc,omitempty" json:"dsc,omitempty"`
}

type DID struct {
	XMLName      xml.Name      `json:"-"`
	Head         string        `xml:"head,omitempty" json:"omitempty"`
	Repository   *Repository   `xml:"repository,omitempty" json:"repository,omitempty"`
	Origination  *Origination  `xml:"origination,omitempty" json:"origination,omitempty"`
	UnitTitle    *UnitTitle    `xml:"unittitle,omitempty" json:"unittitle,omitempty"`
	UnitDate     *UnitDate     `xml:"unitdate,omitempty" json:"unitdate,omitempty"`
	PhysDesc     *PhysDesc     `xml:"physdesc,omitempty" json:"physdesc,omitempty"`
	Abstract     *Abstract     `xml:"abstract,omitempty" json:"abstract,omitempty"`
	UnitID       *UnitID       `xml:"unitid,omitempty" json:"unitid,omitempty"`
	LangMaterial *LangMaterial `xml:"langmaterial,omitempty" json:"langmaterial,omitempty"`
	Container    *Container    `xml:"container,omitempty" json:"container,omitempty"`
}

type Container struct {
	XMLName xml.Name `json:"-"`
	ID      string   `xml:"id,attr,omitempty" json:"id,omitempty"`
	Label   string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	Type    string   `xml:"type,attr,omitempty" json:"type,omitempty"`
	Value   string   `xml:",innerxml" json:"value,omitempty"`
}

type LangMaterial struct {
	XMLName  xml.Name  `json:"-"`
	Language *Language `xml:"language,omitempty" json:"language,omitempty"`
}

type UnitID struct {
	XMLName xml.Name `json:"-"`
	Value   string   `xml:",innerxml" json:"value,omitempty"`
}

type Abstract struct {
	XMLName xml.Name `json:"-"`
	Label   string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	Value   string   `xml:",innerxml" json:"value,omitempty"`
}

type Origination struct {
	XMLName  xml.Name  `json:"-"`
	Label    string    `xml:"label,attr,omitempty" json:"label,omitempty"`
	Persname *Persname `xml:"persname,omitempty" json:"persname,omitempty"`
	Famname  *Famname  `xml:"famname,omitempty" json:"famname,omitempty"`
}

type Persname struct {
	XMLName        xml.Name `json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Source         string   `xml:"source,attr,omitempty" json:"source,omitempty"`
	Value          string   `xml:",innerxml" json:"value,omitempty"`
}

type Famname struct {
	XMLName        xml.Name `json:"-"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Source         string   `xml:"source,attr,omitempty" json:"source,omitempty"`
	Value          string   `xml:",innerxml" json:"value,omitempty"`
}

type UnitTitle struct {
	XMLName        xml.Name `json:"-"`
	Label          string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Value          string   `xml:",innerxml" json:"value,omitempty"`
}

type UnitDate struct {
	XMLName        xml.Name `json:"-"`
	Normal         string   `xml:"normal,attr,omitempty" json:"normal,omitempty"`
	Type           string   `xml:"type,attr,omitempty" json:"type,omitempty"`
	Label          string   `xml:"label,attr,omitempty" json:"label,omitempty"`
	EncodingAnalog string   `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Value          string   `xml:",innerxml" json:"value,omitempty"`
}

type PhysDesc struct {
	XMLName        xml.Name   `json:"-"`
	EncodingAnalog string     `xml:"encodinganalog,attr,omitempty" json:"encodinganalog,omitempty"`
	Label          string     `xml:"label,attr,omitempty" json:"label,omitempty"`
	Extent         ExtentList `xml:"extent,omitempty" json:"extent,omitempty"`
	AltRender      string     `xml:"altrender,attr,omitempty" json:"altrender,omitempty"`
	PhysFacet      *PhysFacet `xml:"physfacet,omitempty" json:"physfacet,omitempty"`
}

type PhysFacet struct {
	XMLName xml.Name `json:"-"`
	Value   string   `xml:",innerxml" json:"value,omitempty"`
}

type Extent struct {
	XMLName   xml.Name `json:"-"`
	AltRender string   `xml:"altrender,attr,omitempty" json:"altrender,omitempty"`
	Value     string   `xml:",innerxml" json:"value,omitempty"`
}

type ExtentList []Extent

type AccessRestrict struct {
	XMLName xml.Name `json:"-"`
	ID      string   `xml:"id,attr,omitempty" json:"id,omitempty"`
	Head    string   `xml:"head,omitemtpy" json:"head,omitempty"`
	P       *P       `xml:"p,omitempty" json:"p,omitempty"`
}

type UserRestrict struct {
	XMLName xml.Name `json:"-"`
	ID      string   `xml:"id,attr,omitempty" json:"id,omitempty"`
	Head    string   `xml:"head,omitemtpy" json:"head,omitempty"`
	P       *P       `xml:"p,omitempty" json:"p,omitempty"`
}

type Prefercite struct {
	XMLName xml.Name `json:"-"`
	ID      string   `xml:"id,attr,omitempty" json:"id,omitempty"`
	Head    string   `xml:"head,omitemtpy" json:"head,omitempty"`
	P       *P       `xml:"p,omitempty" json:"p,omitempty"`
}

type AcqInfo struct {
	XMLName xml.Name `json:"-"`
	ID      string   `xml:"id,attr,omitempty" json:"id,omitempty"`
	Head    string   `xml:"head,omitemtpy" json:"head,omitempty"`
	P       *P       `xml:"p,omitempty" json:"p,omitempty"`
}

type BiogHist struct {
	XMLName xml.Name `json:"-"`
	ID      string   `xml:"id,attr,omitempty" json:"id,omitempty"`
	Head    string   `xml:"head,omitemtpy" json:"head,omitempty"`
	P       *P       `xml:"p,omitempty" json:"p,omitempty"`
}

type ScopeContent struct {
	XMLName  xml.Name `json:"-"`
	Audience string   `xml:"audience,attr,omitempty" json:"audience,omitempty"`
	ID       string   `xml:"id,attr,omitempty" json:"id,omitempty"`
	Head     string   `xml:"head,omitemtpy" json:"head,omitempty"`
	P        *P       `xml:"p,omitempty" json:"p,omitempty"`
}

type OtherFindAid struct {
	XMLName xml.Name `json:"-"`
	ID      string   `xml:"id,attr,omitempty" json:"id,omitempty"`
	Head    string   `xml:"head,omitemtpy" json:"head,omitempty"`
	List    []List   `xml:"list,omitempty" json:"list,omitempty"`
}

type List struct {
	XMLName xml.Name `json:"-"`
	Item    []Item   `xml:"item,omitempty" json:"item,omitempty"`
}

type Item struct {
	XMLName xml.Name `json:"-"`
	ExtRef  *ExtRef  `xml:"extref,omitempty" json:"extref,omitempty"`
}

type ExtRef struct {
	XMLName xml.Name `json:"-"`
	HRef    string   `xml:"href,attr,omitempty" json:"href,omitempty"`
	Value   string   `xml:",innerxml" json:"value,omitempty"`
}

type ControlAccess struct {
	XMLName  xml.Name  `json:"-"`
	Subject  []Subject `xml:"subject,omitempty" json:"subject,omitempty"`
	Corpname *Corpname `xml:"corpname,omitempty" json:"corpname,omitempty"`
	Function *Function `xml:"function,omitempty" json:"function,omitempty"`
}

type Subject struct {
	XMLName xml.Name `json:"-"`
	Source  string   `xml:"source,attr,omitempty" json:"source,omitempty"`
	Value   string   `xml:",innerxml" json:"value,omitempty"`
}

type Corpname struct {
	XMLName xml.Name `json:"-"`
	Source  string   `xml:"source,attr,omitempty" json:"source,omitempty"`
	Value   string   `xml:",innerxml" json:"value,omitempty"`
}

type Function struct {
	XMLName xml.Name `json:"-"`
	Source  string   `xml:"source,attr,omitempty" json:"source,omitempty"`
	Value   string   `xml:",innerxml" json:"value,omitempty"`
}

type DSC struct {
	XMLName xml.Name `json:"-"`
	Source  string   `xml:"source,attr,omitempty" json:"source,omitempty"`
	Value   string   `xml:",innerxml" json:"value,omitempty"`
}

type Repository struct {
	XMLName  xml.Name      `json:"-"`
	Corpname *Corpname     `xml:"corpname,omitempty" json:"corpname,omitempty"`
	Address  *AddressLines `xml:"address,omitempty" json:"address,omitempty"`
}

// Parse parses a ead2002 XML file into a set of Go structures
func Parse(src []byte) (*EAD, error) {
	ead := new(EAD)
	err := xml.Unmarshal(src, ead)
	return ead, err
}

// String write the EAD structure as JSON
func (ead *EAD) String() string {
	src, _ := json.Marshal(ead)
	return string(src)
}

// ToXML write the EAD structure as XML
func (ead *EAD) ToXML() []byte {
	src, _ := xml.Marshal(ead)
	return src
}
