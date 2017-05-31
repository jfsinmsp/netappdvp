// Copyright 2017 NetApp, Inc. All Rights Reserved.

package azgo

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
)

// VolumeCloneSplitStartRequest is a structure to represent a volume-clone-split-start ZAPI request object
type VolumeCloneSplitStartRequest struct {
	XMLName xml.Name `xml:"volume-clone-split-start"`

	VolumePtr                *string `xml:"volume"`
}

// ToXML converts this object into an xml string representation
func (o *VolumeCloneSplitStartRequest) ToXML() (string, error) {
	output, err := xml.MarshalIndent(o, " ", "    ")
	if err != nil {
		log.Errorf("error: %v\n", err)
	}
	return string(output), err
}

// NewVolumeCloneSplitStartRequest is a factory method for creating new instances of VolumeCloneSplitStartRequest objects
func NewVolumeCloneSplitStartRequest() *VolumeCloneSplitStartRequest { return &VolumeCloneSplitStartRequest{} }

// ExecuteUsing converts this object to a ZAPI XML representation and uses the supplied ZapiRunner to send to a filer
func (o *VolumeCloneSplitStartRequest) ExecuteUsing(zr *ZapiRunner) (VolumeCloneSplitStartResponse, error) {
	resp, err := zr.SendZapi(o)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Debugf("response Body:\n%s", string(body))

	var n VolumeCloneSplitStartResponse
	xml.Unmarshal(body, &n)
	if err != nil {
		log.Errorf("err: %v", err.Error())
	}
	log.Debugf("volume-clone-split result:\n%s", n.Result)

	return n, err
}

// String returns a string representation of this object's fields and implements the Stringer interface
func (o VolumeCloneSplitStartRequest) String() string {
	var buffer bytes.Buffer
	if o.VolumePtr != nil {
		buffer.WriteString(fmt.Sprintf("%s: %v\n", "volume", *o.VolumePtr))
	} else {
		buffer.WriteString(fmt.Sprintf("volume: nil\n"))
	}
	return buffer.String()
}


// Volume is a fluent style 'getter' method that can be chained
func (o *VolumeCloneSplitStartRequest) Volume() string {
	r := *o.VolumePtr
	return r
}

// SetVolume is a fluent style 'setter' method that can be chained
func (o *VolumeCloneSplitStartRequest) SetVolume(newValue string) *VolumeCloneSplitStartRequest {
	o.VolumePtr = &newValue
	return o
}

// VolumeCloneSplitStartResponse is a structure to represent a volume-clone-split-start ZAPI response object
type VolumeCloneSplitStartResponse struct {
	XMLName xml.Name `xml:"netapp"`

	ResponseVersion string `xml:"version,attr"`
	ResponseXmlns   string `xml:"xmlns,attr"`

	Result VolumeCloneSplitStartResponseResult `xml:"results"`
}

// String returns a string representation of this object's fields and implements the Stringer interface
func (o VolumeCloneSplitStartResponse) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s: %s\n", "version", o.ResponseVersion))
	buffer.WriteString(fmt.Sprintf("%s: %s\n", "xmlns", o.ResponseXmlns))
	buffer.WriteString(fmt.Sprintf("%s: %s\n", "results", o.Result))
	return buffer.String()
}

// VolumeCloneSplitStartResponseResult is a structure to represent a volume-clone-split-start ZAPI object's result
type VolumeCloneSplitStartResponseResult struct {
	XMLName xml.Name `xml:"results"`

	ResultStatusAttr string `xml:"status,attr"`
	ResultReasonAttr string `xml:"reason,attr"`
	ResultErrnoAttr  string `xml:"errno,attr"`
}

// ToXML converts this object into an xml string representation
func (o *VolumeCloneSplitStartResponse) ToXML() (string, error) {
	output, err := xml.MarshalIndent(o, " ", "    ")
	if err != nil {
		log.Debugf("error: %v", err)
	}
	return string(output), err
}

// NewVolumeCloneSplitStartResponse is a factory method for creating new instances of VolumeCloneSplitStartResponse objects
func NewVolumeCloneSplitStartResponse() *VolumeCloneSplitStartResponse { return &VolumeCloneSplitStartResponse{} }

// String returns a string representation of this object's fields and implements the Stringer interface
func (o VolumeCloneSplitStartResponseResult) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s: %s\n", "resultStatusAttr", o.ResultStatusAttr))
	buffer.WriteString(fmt.Sprintf("%s: %s\n", "resultReasonAttr", o.ResultReasonAttr))
	buffer.WriteString(fmt.Sprintf("%s: %s\n", "resultErrnoAttr", o.ResultErrnoAttr))
	return buffer.String()
}
