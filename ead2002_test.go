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

func TestBasic(t *testing.T) {
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
	testParse10021_MS(path.Join("example-eads", "10021-MS_ead-from-AS.xml"))

	//
	// Test the OAC version of the EAD
	//
	testParse10021_MS(path.Join("example-eads", "10021-MS_ead-from-OAC.xml"))
}
