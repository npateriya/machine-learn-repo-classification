1. PowerShell-WebClient
=======================

A WebClient tool used to fetch updates for Cisco networking appliances.

Use the following command in the Windows Powershell Command Prompt while in the same directory as the script files to get script help:

>> get-help .\WebClient.ps1 -Detailed


2. Equipment.config File Requirements
=====================================

This script requires the creation of an equipment.config XML file.  The format of the XML file is as described by the following XML Schema definition 
and should be placed in the same directory as the scripts:

<?xml version="1.0" encoding="utf-8"?>
<xs:schema xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xs="http://www.w3.org/2001/XMLSchema" attributeFormDefault="unqualified" elementFormDefault="qualified">
  <xsd:element name="Devices">
    <xsd:complexType>
      <xsd:sequence>
        <xsd:element maxOccurs="unbounded" name="Device">
          <xsd:complexType>
            <xsd:sequence>
              <xsd:element name="Type">
                <xs:simpleType>
                  <xs:restriction base="xs:string">
                    <xs:enumeration value="serials"/>
                    <xs:enumeration value="vln"/>
                  </xs:restriction>
                </xs:simpleType>
              </xsd:element>
              <xsd:element name="SerialNumber" type="xsd:string" />
              <xsd:element name="Model" type="xsd:string" />
              <xsd:element name="Applications">
                <xsd:complexType>
                  <xsd:sequence>
                    <xsd:element maxOccurs="unbounded" name="Application" type="xsd:string" />
                  </xsd:sequence>
                </xsd:complexType>
              </xsd:element>
              <xsd:element name="Version" type="xsd:string" />
            </xsd:sequence>
          </xsd:complexType>
        </xsd:element>
      </xsd:sequence>
    </xsd:complexType>
  </xsd:element>
</xs:schema>

Here is an example of what an equipment.config file would look like with one entry:

<?xml version="1.0" encoding="utf-8"?>
<Devices xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
  <Device>
    <Type>serials</Type>
    <SerialNumber>0026B98A7A55-13895M1</SerialNumber>
    <Model>C300</Model>
    <Applications>
      <Application>wbrs</Application>
      <Application>sophos</Application>
    </Applications>
    <Version>phoebe-8-5-6-092</Version>
  </Device>
</Devices>

3. Client Certificate Requirements
==================================
You must acquire and register a client certificate in the Windows Certificate Store in order to run this script successfully.
You may add the certificate together with it's private key to the Personal directory in either the Local User or Local Machine store.

4. Windows Requirements
==================================
Powershell version 4.0 or higher.