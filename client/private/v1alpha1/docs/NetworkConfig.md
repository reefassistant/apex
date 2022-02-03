# NetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dhcp** | **bool** |  | 
**Hostname** | **string** |  | 
**Ipaddr** | **string** |  | 
**Netmask** | **string** |  | 
**Gateway** | **string** |  | 
**Dns** | **[]string** |  | 
**HttpPort** | **float32** |  | 
**User** | **string** |  | 
**Password** | **string** |  | 
**DefaultAuth** | **bool** |  | 
**FusionEnable** | **bool** |  | 
**WifiAPLock** | **bool** |  | 
**WifiEnable** | **bool** |  | 
**Ssid** | **string** |  | 
**WifiAP** | **bool** |  | 
**EmailEnable** | **bool** |  | 
**SmtpPort** | **float32** |  | 
**SmtpServer** | **string** |  | 
**EmailFrom** | **string** |  | 
**EmailTo** | **string** |  | 
**ReEmail** | **float32** |  | 
**EmailAuth** | **bool** |  | 
**EmailUser** | **string** |  | 
**EmailPassword** | **string** |  | 
**UpdateFirmware** | **bool** |  | 
**LatestFirmware** | **string** |  | 

## Methods

### NewNetworkConfig

`func NewNetworkConfig(dhcp bool, hostname string, ipaddr string, netmask string, gateway string, dns []string, httpPort float32, user string, password string, defaultAuth bool, fusionEnable bool, wifiAPLock bool, wifiEnable bool, ssid string, wifiAP bool, emailEnable bool, smtpPort float32, smtpServer string, emailFrom string, emailTo string, reEmail float32, emailAuth bool, emailUser string, emailPassword string, updateFirmware bool, latestFirmware string, ) *NetworkConfig`

NewNetworkConfig instantiates a new NetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkConfigWithDefaults

`func NewNetworkConfigWithDefaults() *NetworkConfig`

NewNetworkConfigWithDefaults instantiates a new NetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcp

`func (o *NetworkConfig) GetDhcp() bool`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *NetworkConfig) GetDhcpOk() (*bool, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *NetworkConfig) SetDhcp(v bool)`

SetDhcp sets Dhcp field to given value.


### GetHostname

`func (o *NetworkConfig) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NetworkConfig) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NetworkConfig) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetIpaddr

`func (o *NetworkConfig) GetIpaddr() string`

GetIpaddr returns the Ipaddr field if non-nil, zero value otherwise.

### GetIpaddrOk

`func (o *NetworkConfig) GetIpaddrOk() (*string, bool)`

GetIpaddrOk returns a tuple with the Ipaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddr

`func (o *NetworkConfig) SetIpaddr(v string)`

SetIpaddr sets Ipaddr field to given value.


### GetNetmask

`func (o *NetworkConfig) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *NetworkConfig) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *NetworkConfig) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.


### GetGateway

`func (o *NetworkConfig) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *NetworkConfig) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *NetworkConfig) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetDns

`func (o *NetworkConfig) GetDns() []string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *NetworkConfig) GetDnsOk() (*[]string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *NetworkConfig) SetDns(v []string)`

SetDns sets Dns field to given value.


### GetHttpPort

`func (o *NetworkConfig) GetHttpPort() float32`

GetHttpPort returns the HttpPort field if non-nil, zero value otherwise.

### GetHttpPortOk

`func (o *NetworkConfig) GetHttpPortOk() (*float32, bool)`

GetHttpPortOk returns a tuple with the HttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPort

`func (o *NetworkConfig) SetHttpPort(v float32)`

SetHttpPort sets HttpPort field to given value.


### GetUser

`func (o *NetworkConfig) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *NetworkConfig) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *NetworkConfig) SetUser(v string)`

SetUser sets User field to given value.


### GetPassword

`func (o *NetworkConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NetworkConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NetworkConfig) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDefaultAuth

`func (o *NetworkConfig) GetDefaultAuth() bool`

GetDefaultAuth returns the DefaultAuth field if non-nil, zero value otherwise.

### GetDefaultAuthOk

`func (o *NetworkConfig) GetDefaultAuthOk() (*bool, bool)`

GetDefaultAuthOk returns a tuple with the DefaultAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAuth

`func (o *NetworkConfig) SetDefaultAuth(v bool)`

SetDefaultAuth sets DefaultAuth field to given value.


### GetFusionEnable

`func (o *NetworkConfig) GetFusionEnable() bool`

GetFusionEnable returns the FusionEnable field if non-nil, zero value otherwise.

### GetFusionEnableOk

`func (o *NetworkConfig) GetFusionEnableOk() (*bool, bool)`

GetFusionEnableOk returns a tuple with the FusionEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFusionEnable

`func (o *NetworkConfig) SetFusionEnable(v bool)`

SetFusionEnable sets FusionEnable field to given value.


### GetWifiAPLock

`func (o *NetworkConfig) GetWifiAPLock() bool`

GetWifiAPLock returns the WifiAPLock field if non-nil, zero value otherwise.

### GetWifiAPLockOk

`func (o *NetworkConfig) GetWifiAPLockOk() (*bool, bool)`

GetWifiAPLockOk returns a tuple with the WifiAPLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiAPLock

`func (o *NetworkConfig) SetWifiAPLock(v bool)`

SetWifiAPLock sets WifiAPLock field to given value.


### GetWifiEnable

`func (o *NetworkConfig) GetWifiEnable() bool`

GetWifiEnable returns the WifiEnable field if non-nil, zero value otherwise.

### GetWifiEnableOk

`func (o *NetworkConfig) GetWifiEnableOk() (*bool, bool)`

GetWifiEnableOk returns a tuple with the WifiEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiEnable

`func (o *NetworkConfig) SetWifiEnable(v bool)`

SetWifiEnable sets WifiEnable field to given value.


### GetSsid

`func (o *NetworkConfig) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *NetworkConfig) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *NetworkConfig) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetWifiAP

`func (o *NetworkConfig) GetWifiAP() bool`

GetWifiAP returns the WifiAP field if non-nil, zero value otherwise.

### GetWifiAPOk

`func (o *NetworkConfig) GetWifiAPOk() (*bool, bool)`

GetWifiAPOk returns a tuple with the WifiAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiAP

`func (o *NetworkConfig) SetWifiAP(v bool)`

SetWifiAP sets WifiAP field to given value.


### GetEmailEnable

`func (o *NetworkConfig) GetEmailEnable() bool`

GetEmailEnable returns the EmailEnable field if non-nil, zero value otherwise.

### GetEmailEnableOk

`func (o *NetworkConfig) GetEmailEnableOk() (*bool, bool)`

GetEmailEnableOk returns a tuple with the EmailEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailEnable

`func (o *NetworkConfig) SetEmailEnable(v bool)`

SetEmailEnable sets EmailEnable field to given value.


### GetSmtpPort

`func (o *NetworkConfig) GetSmtpPort() float32`

GetSmtpPort returns the SmtpPort field if non-nil, zero value otherwise.

### GetSmtpPortOk

`func (o *NetworkConfig) GetSmtpPortOk() (*float32, bool)`

GetSmtpPortOk returns a tuple with the SmtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPort

`func (o *NetworkConfig) SetSmtpPort(v float32)`

SetSmtpPort sets SmtpPort field to given value.


### GetSmtpServer

`func (o *NetworkConfig) GetSmtpServer() string`

GetSmtpServer returns the SmtpServer field if non-nil, zero value otherwise.

### GetSmtpServerOk

`func (o *NetworkConfig) GetSmtpServerOk() (*string, bool)`

GetSmtpServerOk returns a tuple with the SmtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServer

`func (o *NetworkConfig) SetSmtpServer(v string)`

SetSmtpServer sets SmtpServer field to given value.


### GetEmailFrom

`func (o *NetworkConfig) GetEmailFrom() string`

GetEmailFrom returns the EmailFrom field if non-nil, zero value otherwise.

### GetEmailFromOk

`func (o *NetworkConfig) GetEmailFromOk() (*string, bool)`

GetEmailFromOk returns a tuple with the EmailFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFrom

`func (o *NetworkConfig) SetEmailFrom(v string)`

SetEmailFrom sets EmailFrom field to given value.


### GetEmailTo

`func (o *NetworkConfig) GetEmailTo() string`

GetEmailTo returns the EmailTo field if non-nil, zero value otherwise.

### GetEmailToOk

`func (o *NetworkConfig) GetEmailToOk() (*string, bool)`

GetEmailToOk returns a tuple with the EmailTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTo

`func (o *NetworkConfig) SetEmailTo(v string)`

SetEmailTo sets EmailTo field to given value.


### GetReEmail

`func (o *NetworkConfig) GetReEmail() float32`

GetReEmail returns the ReEmail field if non-nil, zero value otherwise.

### GetReEmailOk

`func (o *NetworkConfig) GetReEmailOk() (*float32, bool)`

GetReEmailOk returns a tuple with the ReEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReEmail

`func (o *NetworkConfig) SetReEmail(v float32)`

SetReEmail sets ReEmail field to given value.


### GetEmailAuth

`func (o *NetworkConfig) GetEmailAuth() bool`

GetEmailAuth returns the EmailAuth field if non-nil, zero value otherwise.

### GetEmailAuthOk

`func (o *NetworkConfig) GetEmailAuthOk() (*bool, bool)`

GetEmailAuthOk returns a tuple with the EmailAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAuth

`func (o *NetworkConfig) SetEmailAuth(v bool)`

SetEmailAuth sets EmailAuth field to given value.


### GetEmailUser

`func (o *NetworkConfig) GetEmailUser() string`

GetEmailUser returns the EmailUser field if non-nil, zero value otherwise.

### GetEmailUserOk

`func (o *NetworkConfig) GetEmailUserOk() (*string, bool)`

GetEmailUserOk returns a tuple with the EmailUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailUser

`func (o *NetworkConfig) SetEmailUser(v string)`

SetEmailUser sets EmailUser field to given value.


### GetEmailPassword

`func (o *NetworkConfig) GetEmailPassword() string`

GetEmailPassword returns the EmailPassword field if non-nil, zero value otherwise.

### GetEmailPasswordOk

`func (o *NetworkConfig) GetEmailPasswordOk() (*string, bool)`

GetEmailPasswordOk returns a tuple with the EmailPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPassword

`func (o *NetworkConfig) SetEmailPassword(v string)`

SetEmailPassword sets EmailPassword field to given value.


### GetUpdateFirmware

`func (o *NetworkConfig) GetUpdateFirmware() bool`

GetUpdateFirmware returns the UpdateFirmware field if non-nil, zero value otherwise.

### GetUpdateFirmwareOk

`func (o *NetworkConfig) GetUpdateFirmwareOk() (*bool, bool)`

GetUpdateFirmwareOk returns a tuple with the UpdateFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateFirmware

`func (o *NetworkConfig) SetUpdateFirmware(v bool)`

SetUpdateFirmware sets UpdateFirmware field to given value.


### GetLatestFirmware

`func (o *NetworkConfig) GetLatestFirmware() string`

GetLatestFirmware returns the LatestFirmware field if non-nil, zero value otherwise.

### GetLatestFirmwareOk

`func (o *NetworkConfig) GetLatestFirmwareOk() (*string, bool)`

GetLatestFirmwareOk returns a tuple with the LatestFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestFirmware

`func (o *NetworkConfig) SetLatestFirmware(v string)`

SetLatestFirmware sets LatestFirmware field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


