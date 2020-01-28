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
	"io/ioutil"
	"path"
	"strings"
	"testing"
)

func TestDataEADs(t *testing.T) {
	fname := path.Join("testeads", "2016.0518.0001.test_ead.xml")
	src, err := ioutil.ReadFile(fname)
	if err != nil {
		t.Errorf("Cannot read %s, %s", fname, err)
		t.FailNow()
	}

	ead, err := Parse(src)
	if err != nil {
		t.Errorf("Failed ot parse %s, %s", fname, err)
	}

	if ead.EADHeader == nil {
		t.Errorf("Missing ead.EADHeader")
		t.FailNow()
	}
	if ead.ArchDesc == nil {
		t.Errorf("Missing ead.ArchDesc")
		t.FailNow()
	}
	eadheader := ead.EADHeader
	if eadheader.EADID.Value != "" {
		t.Errorf("Wrong value eadheader.EADID.Value %+v", eadheader.EADID)
	}
	if eadheader.FileDesc == nil {
		t.Errorf("Missing eadheader.FileDesc")
		t.FailNow()
	}
	filedesc := eadheader.FileDesc
	if filedesc.TitleStmt == nil {
		t.Errorf("Missing filedesc.TitleStmt")
		t.FailNow()
	}
	expected := `How did galaxies form?<num>2016.0518.0001.test</num>`
	if len(filedesc.TitleStmt.TitleProper) != 1 ||
		filedesc.TitleStmt.TitleProper[0].Value != expected {
		t.Errorf("Wrong value filedesc.TitleStmt.TitleProper, %+v", filedesc.TitleStmt.TitleProper)
	}
	if filedesc.PublicationStmt == nil {
		t.Errorf("Missing filedesc.PublicationStmt")
		t.FailNow()
	}
	publicationstmt := filedesc.PublicationStmt
	if publicationstmt.Publisher == nil {
		t.Errorf("Missing pulbicationstmt.Publisher")
	}
	if publicationstmt.Publisher.Value != "Caltech Archives" {
		t.Errorf("Wrong value for publications.Publisher, %+v", publicationstmt.Publisher)
	}

	archdesc := ead.ArchDesc
	if len(archdesc.DID) != 1 {
		t.Errorf("Wrong value for archdesc.DID, %+v", archdesc.DID)
	}

	dids := archdesc.DID
	did := dids[0]
	if did.LangMaterial == nil {
		t.Errorf("Missing did.LangMaterial")
	}
	if did.LangMaterial.Language == nil {
		t.Errorf("Missing did.LangMaterial.Language")
	}
	if did.LangMaterial.Language.LangCode != "eng" {
		t.Errorf("Wrong value for did.LangMaterial.Language.LangCode, %+v", did.LangMaterial.Language)
	}
	if did.LangMaterial.Language.Value != "English" {
		t.Errorf("Wrong value for did.LangMaterial.Language.LangCode, %+v", did.LangMaterial.Language)
	}
	if did.Repository == nil {
		t.Errorf("Missing did.Repository")
		t.FailNow()
	}
	if did.Repository.Corpname == nil {
		t.Errorf("Missing did.Repository.Corpname")
	}
	if did.Repository.Corpname.Value != "Caltech Archives" {
		t.Errorf("Missing did.Repository.Corpname.Value, %+v", did.Repository.Corpname)
	}
	if did.UnitID == nil {
		t.Errorf("Missing did.UnitID")
	}
	if did.UnitID.Value != "2016.0518.0001.test" {
		t.Errorf("Wrong value did.UnitID.Value, %+v", did.UnitID.Value)
	}
	if did.PhysDesc == nil {
		t.Errorf("Missing did.PhysDesc")
		t.FailNow()
	}
	physdesc := did.PhysDesc
	if physdesc.AltRender != "whole" {
		t.Errorf("Wrong value for physdesc.AltRender, %+v", physdesc)
	}
	if physdesc.Extent == nil {
		t.Errorf("Missing physdesc.Extent")
	}
	if len(physdesc.Extent) != 1 {
		t.Errorf("Wrong value for physdesc.Extent, %+v", physdesc.Extent)
		t.FailNow()
	}
	if physdesc.Extent[0].AltRender != "materialtype spaceoccupied" {
		t.Errorf("Wrong value for physdesc.Extent, %+v", physdesc.Extent)
		t.FailNow()
	}
	if physdesc.Extent[0].Value != "1 Multimedia" {
		t.Errorf("Wrong value for physdesc.Extent, %+v", physdesc.Extent)
		t.FailNow()
	}
	if physdesc.PhysFacet == nil {
		t.Errorf("Missing physdesc.PhysFacet")
		t.FailNow()
	}
	if physdesc.PhysFacet.Value != "MEDIUM: Videorecording; FORMAT: U-matic master; VHS; LENGTH: 57m 4s; QUANTITY: 1,1" {
		t.Errorf("Wrong value for physdesc.PhysFacet, %+v", physdesc.PhysFacet)
	}
	if did.UnitDate == nil {
		t.Errorf("Missing did.UnitDate")
		t.FailNow()
	}
	unitdate := did.UnitDate
	if unitdate.Value != "1997-05-14" {
		t.Errorf("Wrong value for unitdate.Value, %+v", unitdate)
	}
	if unitdate.Type != "inclusive" {
		t.Errorf("Wrong value for unitdate.Type, %+v", unitdate)
	}
	if did.Container == nil {
		t.Errorf("Missing did.Container")
		t.FailNow()
	}
	if did.Container.ID != "aspace_b456221fc43b9e2bef78728177fef5fa" {
		t.Errorf("missing value for did.Container.ID, %+v", did.Container)
	}
	if did.Container.Label != "Mixed Materials" {
		t.Errorf("Missing value for did.Container.Label, %+v", did.Container)
	}
	if did.Container.Type != "box" {
		t.Errorf("Missing value for did.Container.Type, %+v", did.Container.Type)
	}
	if did.Container.Value != "Test Box Container" {
		t.Errorf("Wrong value for did.Container.Value, %+v", did.Container.Value)
	}

	if archdesc.ScopeContent == nil {
		t.Errorf("Missing ead.ArchDesc.ScopeContent")
		t.FailNow()
	}
	if archdesc.ScopeContent.Audience != "internal" {
		t.Errorf("Wrong value for archdesc.ScopeContent, %+v", archdesc.ScopeContent)
	}
	if archdesc.ScopeContent.ID != "aspace_83fcae08dcbb5f987bcf9db34d048fbe" {
		t.Errorf("Wrong value for archdesc.ScopeContent.ID, %+v", archdesc.ScopeContent)
	}
	if archdesc.ScopeContent.Head != "Content Description" {
		t.Errorf("Wrong value for archdesc.ScopeContent.Head, %+v", archdesc.ScopeContent)
	}
	if archdesc.ScopeContent.P == nil {
		t.Errorf("Missing archdesc.ScopeContent.P")
		t.FailNow()
	}
	if archdesc.ScopeContent.P.Value != `Seventh in series of 10 lectures for non-specialists in this subject area, April -June 1997. George Djorgovski is co-speaker.  See also Seminar 0.1 lectures in other subject areas.` {
		t.Errorf("Wrong value for archdesc.ScopeContent.P, %+v", archdesc.ScopeContent.P)
	}
	if archdesc.ControlAccess == nil {
		t.Errorf("Missing archdesc.ControlAccess")
		t.FailNow()
	}
	if len(archdesc.ControlAccess.Subject) != 2 {
		t.Errorf("Wrong value for archdesc.ControlAccess.Subject, %+v", archdesc.ControlAccess)
	}
	if archdesc.ControlAccess.Function.Source != "local" {
		t.Errorf("Wrong value for archdesc.ControlAccess.Function.Source, %+v", archdesc.ControlAccess.Function)
	}
	if archdesc.ControlAccess.Function.Value != "Seminar 0.1 series" {
		t.Errorf("Wrong value for archdesc.ControlAccess.Function.Value, %+v", archdesc.ControlAccess.Function)
	}
	if archdesc.DSC == nil {
		t.Errorf("Missing archdesc.DSC")
		t.FailNow()
	}
	if archdesc.DSC.Value != "" {
		t.Errorf("Wrong value for archdesc.DSC, %+v", archdesc.DSC)
	}
}

func TestExampleEADs(t *testing.T) {
	// See if we can read an EAD produced by ArchivesSpace
	testParse10021_MS := func(fname string) {
		src, err := ioutil.ReadFile(fname)
		if err != nil {
			t.Errorf("Could not read %s", fname)
		}
		ead, err := Parse(src)
		if err != nil {
			t.Errorf("Could not parse %s, %s\n", fname, err)
			t.FailNow()
		}
		if ead == nil {
			t.Errorf("Expected a populated ead")
			t.FailNow()
		}
		if ead.EADHeader == nil {
			t.Errorf("Expected a populated ead.EADHeader")
			t.FailNow()
		}
		if ead.EADHeader.EADID == nil {
			t.Errorf("Expected a populated ead.EADHeader.EADID")
			t.FailNow()
		}
		if ead.EADHeader.EADID.Value != "10021-MS" {
			t.Errorf("Missing value for EADID")
		}
		if ead.EADHeader.FileDesc == nil {
			t.Errorf("Expected a populated ead.EADHeader.FileDesc")
			t.FailNow()
		}
		if ead.EADHeader.FileDesc.TitleStmt == nil {
			t.Errorf("Expected a populated ead.EADHeader.FileDesc.TitleStmt")
			t.FailNow()
		}
		if len(ead.EADHeader.FileDesc.TitleStmt.TitleProper) != 2 {
			t.Errorf("Expected a populated ead.EADHeader.FileDesc.TitleStmt.TitleProper array")
			t.FailNow()
		}
		// This array is not ordered so adjust sequence check as necessary
		a, b := 0, 1
		tp := ead.EADHeader.FileDesc.TitleStmt.TitleProper
		if tp[0].Type == "" {
			a = 1
			b = 0
		}
		if tp[a].Type != "filing" {
			t.Errorf("Wrong value for TitleProper[0].Type: %+v", tp[a])
		}
		if tp[a].Value != "Millikan (Clark, B.) Papers" {
			t.Errorf("Wrong value for TitleProper[0].Value: %+v", tp[a])
		}
		if tp[b].Value == "" {
			t.Errorf("Wrong value for TitleProper[1].Value %+v", tp[b])
		}
		if ead.EADHeader.FileDesc.TitleStmt.Author == nil {
			t.Errorf("Expected a populated ead.EADHeader.FileDesc.Author")
			t.FailNow()
		}
		expected := "Finding aid created by California Institute of Technology staff using RecordEXPRESS"
		if ead.EADHeader.FileDesc.TitleStmt.Author.Value != expected {
			t.Errorf("Expected a populated Author value: %+v", ead.EADHeader.FileDesc.TitleStmt.Author)
		}
		if ead.EADHeader.FileDesc.PublicationStmt == nil {
			t.Errorf("Expected a populated ead.EADHeader.FileDesc.PublicationStmt")
			t.FailNow()
		}
		if ead.EADHeader.FileDesc.PublicationStmt.Publisher == nil {
			t.Errorf("Expected a populated ead.EADHeader.FileDesc.PublicationStmt.Publisher")
			t.FailNow()
		}
		if strings.HasSuffix(fname, "OAC.xml") {
			if ead.EADHeader.FileDesc.PublicationStmt.Address == nil {
				t.Errorf("Missing FileDesc.PublicationStmt.Address")
				t.FailNow()
			}
			expected = "California Institute of Technology"
			if ead.EADHeader.FileDesc.PublicationStmt.Publisher.Value != expected {
				t.Errorf("Wrong value for ead.EADHeader.FileDesc.PublicationStmt.Publisher.Value: %+v", ead.EADHeader.FileDesc.PublicationStmt.Publisher.Value)
			}
			if len(ead.EADHeader.FileDesc.PublicationStmt.Address.AddressLine) != 5 {
				t.Errorf("Wrong value ead.EADHeader.FileDesc.PublicationStmt.Address, %+v", ead.EADHeader.FileDesc.PublicationStmt)
			}
		} else {
			if ead.EADHeader.FileDesc.PublicationStmt.Publisher.Value != "Caltech Archives" {
				t.Errorf("Wrong value for FileDesc.PublicationStmt.Publisher.Value: %+v", ead.EADHeader.FileDesc.PublicationStmt.Publisher.Value)
			}
			if ead.EADHeader.FileDesc.PublicationStmt.P == nil {
				t.Errorf("Expected a populated ead.EADHeader.FileDesc.PublicationStmt.P")
				t.FailNow()
			}
			if ead.EADHeader.FileDesc.PublicationStmt.P.Value != `<date>2014</date>` {
				t.Errorf("Wrong value for ad.EADHeader.FileDesc.PublicationStmt.P, %+v", ead.EADHeader.FileDesc.PublicationStmt.P)
			}

			if ead.EADHeader.ProfileDesc == nil {
				t.Errorf("Expected a populated ead.EADHeader.ProfileDesc")
				t.FailNow()
			}
			profileDesc := ead.EADHeader.ProfileDesc
			if profileDesc.Creation == nil {
				t.Errorf("Expected a populated profileDesc.Creation")
				t.FailNow()
			}
			expected = "This finding aid was produced using ArchivesSpace on <date>2016-06-01 16:42:46 UTC</date>."
			if strings.Compare(profileDesc.Creation.Value, expected) != 0 {
				t.Errorf("Wrong value for profileDesc.Creation, %+v", profileDesc.Creation)
			}
		}
		if ead.ArchDesc == nil {
			t.Errorf("Expected a populated ead.ArchDesc")
			t.FailNow()
		}
		if len(ead.ArchDesc.DID) != 1 {
			t.Errorf("Wrong value for archDesc.DID")
			t.FailNow()
		}
		archDesc := ead.ArchDesc
		dids := archDesc.DID
		did := dids[0]

		if did.LangMaterial == nil {
			t.Errorf("Expected populated did")
			t.FailNow()
		}
		if did.LangMaterial.Language == nil {
			t.Errorf("Missing did.LangMaterial.Language")
		}
		if strings.HasSuffix(fname, "AS.xml") {
			if did.LangMaterial.Language.Value != "English" {
				t.Errorf("Wrong value for did.LangMaterial, %+v", did.LangMaterial)
			}
			if did.LangMaterial.Language.LangCode != "eng" {
				t.Errorf("Wrong value for did.LangMaterial.LangCode, %+v", did.LangMaterial)
			}
		} else {
			if did.LangMaterial.Language.LangCode != "eng" {
				t.Errorf("Wrong value for did.LangMaterial.LangCode, %+v", did.LangMaterial)
			}
		}
		if did.Repository == nil {
			t.Errorf("Expected a populated did.Repository")
			t.FailNow()
		}
		if did.Repository.Corpname == nil {
			t.Errorf("Missing did.Repository.Corpname")
		}
		if did.Repository.Corpname.Value != "Caltech Archives" &&
			did.Repository.Corpname.Value != "California Institute of Technology" {
			t.Errorf("Wrong value for did.Repository.Corpname.Value, %+v", did.Repository.Corpname)
		}
		if strings.HasSuffix(fname, "OAC.xml") {
			if did.Repository.Address == nil {
				t.Errorf("Missing value for did.Repository.Address")
			}
			if len(did.Repository.Address.AddressLine) != 1 {
				t.Errorf("Wrong value for did.Repository.Address.AddressLine, %+v", did.Repository.Address)
			}
		}
		if did.Origination == nil {
			t.Errorf("Expected populated did.Origination")
			t.FailNow()
		}
		if did.Origination.Persname == nil {
			t.Errorf("Missing did.Origination.Persname, %+v", did.Origination)
		}
		if did.Origination.Persname.Value != "Clark, B." {
			t.Errorf("Wrong value for did.Origination.Persname.Value, %+v", did.Origination.Persname)
		}
		if strings.HasSuffix(fname, "AS.xml") {
			if did.Origination.Persname.Source != "ingest" {
				t.Errorf("Wrong value for did.Origination.Persname.Source, %+v", did.Origination.Persname)
			}
		} else {
			if did.Origination.Famname.Value != "Millikan" {
				t.Errorf("Wrong value for did.Origination.Famname, %+v", did.Origination)
			}
		}
		if did.UnitTitle == nil {
			t.Errorf("Expected a populated did.UnitTitle")
			t.FailNow()
		}
		expected = "Guide to the Papers of Clark B. Millikan, 1922-1965"
		if did.UnitTitle.Value != expected {
			t.Errorf("Wrong value for did.UnitTitle.Value, %+v", did.UnitTitle)
		}

	}

	//
	// Test the ArchivesSpace generated EAD
	//
	testParse10021_MS(path.Join("testeads", "10021-MS_ead-from-AS.xml"))

	//
	// Test the OAC version of the EAD
	//
	testParse10021_MS(path.Join("testeads", "10021-MS_ead-from-OAC.xml"))
}
