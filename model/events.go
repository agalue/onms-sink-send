package model

import (
	"encoding/xml"
	"time"
)

type Time struct {
	time.Time
}

func (t Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.IsZero() {
		return e.EncodeElement("", start)
	}
	return e.EncodeElement(t.Format("2006-01-02T15:04:05.999999-07:00"), start)
}

func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	var err error
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	t.Time, err = time.Parse(s, s)
	if err != nil {
		return err
	}
	return nil
}

type MaskElement struct {
	XMLName xml.Name `xml:"maskelement"`
	MEname  string   `xml:"mename"`
	MEvalue []string `xml:"mevalue"`
}

type Mask struct {
	XMLName  xml.Name      `xml:"mask"`
	Elements []MaskElement `xml:"maskelement"`
}

type SNMP struct {
	XMLName   xml.Name `xml:"snmp"`
	ID        string   `xml:"id"`
	IDText    string   `xml:"idtext,omitempty"`
	Version   string   `xml:"version"`
	Specific  int      `xml:"specific"`
	Generic   int      `xml:"generic"`
	Community string   `xml:"community"`
	Timestamp int      `xml:"time-stamp"`
}

type Param struct {
	XMLName xml.Name `xml:"parm"`
	Name    string   `xml:"parmName"`
	Value   string   `xml:"value"`
}

type Parms struct {
	XMLName xml.Name `xml:"parms"`
	Params  []Param  `xml:"parm"`
}

type LogMsg struct {
	XMLName     xml.Name `xml:"logmsg"`
	Destination string   `xml:"dest,attr"` // logndisplay|displayonly|logonly|suppress|donotpersist
	Notify      bool     `xml:"notify,attr,omitempty"`
	Content     string   `xml:",chardata"`
}

type Correlation struct {
	XMLName xml.Name `xml:"correlation"`
	State   string   `xml:"state,attr"` // on|off
	Path    string   `xml:"path,attr"`
	CUEI    []string `xml:"cuei"`
	CMin    string   `xml:"cmin"`
	CMax    string   `xml:"cmax"`
	CTime   *Time    `xml:"ctime"`
}

type AutoAction struct {
	XMLName xml.Name `xml:"autoaction"`
	State   string   `xml:"state,attr"` // on|off
	Content string   `xml:",chardata"`
}

type OperAction struct {
	XMLName  xml.Name `xml:"operaction"`
	State    string   `xml:"state,attr"` // on|off
	MenuText string   `xml:"menutext,attr"`
	Content  string   `xml:",chardata"`
}

type AutoAcknowledge struct {
	XMLName xml.Name `xml:"autoacknowledge"`
	State   string   `xml:"state,attr"` // on|off
	Content string   `xml:",chardata"`
}

type TTicket struct {
	XMLName xml.Name `xml:"tticket"`
	State   string   `xml:"state,attr"` // on|off
	Content string   `xml:",chardata"`
}

type Forward struct {
	XMLName   xml.Name `xml:"forward"`
	State     string   `xml:"state,attr"`     // on|off
	Mechanism string   `xml:"mechanism,attr"` // snmpudp|snmptcp|xmltcp|xmludp
	Content   string   `xml:",chardata"`
}

type Script struct {
	XMLName  xml.Name `xml:"script"`
	Language string   `xml:"language,attr"`
	Content  string   `xml:",chardata"`
}

type UpdateField struct {
	XMLName           xml.Name `xml:"update-field"`
	FieldName         string   `xml:"reduction-key,attr"`
	UpdateOnReduction bool     `xml:"update-on-reduction,attr"`
	ValueExpression   string   `xml:"value-expression,attr"`
}

type ManagedObject struct {
	XMLName xml.Name `xml:"managed-object"`
	Type    string   `xml:"type,attr"`
}

type AlarmData struct {
	XMLName           xml.Name       `xml:"alarm-data"`
	ReductionKey      string         `xml:"reduction-key,attr"`
	AlarmType         int            `xml:"alarm-type,attr"`
	ClearKey          string         `xml:"clear-key,attr,omitempty"`
	AutoClean         bool           `xml:"auto-clean,attr,omitempty"`
	X733AlarmType     string         `xml:"x733-alarm-type,attr,omitempty"`
	X733ProbableCause string         `xml:"x733-probable-cause,attr,omitempty"`
	UpdateField       []UpdateField  `xml:"update-field,omitempty"`
	ManagedObject     *ManagedObject `xml:"managed-object,omitempty"`
}

type Event struct {
	XMLName         xml.Name         `xml:"event"`
	UUID            string           `xml:"uuid,attr,omitempty"`
	DBID            int              `xml:"dbid,omitempty"`
	DistPoller      string           `xml:"dist-poller,omitempty"`
	CreationTime    *Time            `xml:"creation-time,omitempty"`
	MasterStation   string           `xml:"master-station,omitempty"`
	Mask            *Mask            `xml:"mask,omitempty"`
	UEI             string           `xml:"uei,omitempty"`
	Source          string           `xml:"source,omitempty"`
	NodeID          int              `xml:"nodeid,omitempty"`
	EventTime       *Time            `xml:"time,omitempty"`
	Host            string           `xml:"host,omitempty"`
	Interface       string           `xml:"interface,omitempty"`
	SNMPHost        string           `xml:"snmphost,omitempty"`
	Service         string           `xml:"service,omitempty"`
	SNMP            *SNMP            `xml:"snmp,omitempty"`
	Parameters      *Parms           `xml:"parms,omitempty"`
	Descr           string           `xml:"descr,omitempty"`
	LogMsg          *LogMsg          `xml:"logmsg,omitempty"`
	Severity        string           `xml:"severity,omitempty"`
	PathOutage      string           `xml:"pathoutage,omitempty"`
	Correlation     *Correlation     `xml:"correlation,omitempty"`
	OperInstruct    string           `xml:"operinstruct,omitempty"`
	AutoAction      []AutoAction     `xml:"autoaction,omitempty"`
	OperAction      []OperAction     `xml:"operaction,omitempty"`
	AutoAcknowledge *AutoAcknowledge `xml:"autoacknowledge,omitempty"`
	LogGroup        []string         `xml:"loggroup,omitempty"`
	TTicket         *TTicket         `xml:"tticket,omitempty"`
	Forward         *Forward         `xml:"forward,omitempty"`
	Script          *Script          `xml:"script,omitempty"`
	IfIndex         int              `xml:"ifIndex,omitempty"`
	IfAlias         string           `xml:"ifAlias,omitempty"`
	MouseOverText   string           `xml:"mouseovertext,omitempty"`
	AlarmData       *AlarmData       `xml:"alarm-data,omitempty"`
}
