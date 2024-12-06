# AvailabilityCheckSettings


## Fields

| Field                                                                                                                                                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                          | Example                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `CheckForString`                                                                                                                                                                                                                                                                                                                                                                                                                                                     | [*components.CheckForString](../../models/components/checkforstring.md)                                                                                                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                   |   Use this field to configure whether availability tests should check for presence or absence of a particular string on a page.<br/>  If the operator is DOES_NOT_CONTAIN and the value is found on the page, the availability test will fail.<br/>  Likewise, if the operator is CONTAINS and the value is not found on the page, the availability test will fail.<br/>  If omitted or set to null, the string checking functionality will be disabled.             |                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| `TestIntervalInSeconds`                                                                                                                                                                                                                                                                                                                                                                                                                                              | *int*                                                                                                                                                                                                                                                                                                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Configure how often availability tests should be performed. Provide a number of seconds that is divisible by 60 and no greater than 14400 (4 hours).                                                                                                                                                                                                                                                                                                                 |                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| `Protocols`                                                                                                                                                                                                                                                                                                                                                                                                                                                          | [][components.WebsiteProtocol](../../models/components/websiteprotocol.md)                                                                                                                                                                                                                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Configure which protocols need availability tests to be performed. At least one protocol must be provided.                                                                                                                                                                                                                                                                                                                                                           |                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| `PlatformOptions`                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [*components.ProbePlatformOptions](../../models/components/probeplatformoptions.md)                                                                                                                                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Configure cloud platforms of the synthetic availability test probes. If omitted or set to null, no particular cloud platform will be enforced.                                                                                                                                                                                                                                                                                                                       |                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| `TestFrom`                                                                                                                                                                                                                                                                                                                                                                                                                                                           | [components.TestFrom](../../models/components/testfrom.md)                                                                                                                                                                                                                                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                   |   Configure locations of the synthetic availability test probes. <br/>  Acceptable values depend on the selected type and actual values of existing probes.                                                                                                                                                                                                                                                                                                          | {<br/>"type": "REGION",<br/>"values": [<br/>"NA"<br/>]<br/>}                                                                                                                                                                                                                                                                                                                                                                                                         |
| `Ssl`                                                                                                                                                                                                                                                                                                                                                                                                                                                                | [*components.SslMonitoring](../../models/components/sslmonitoring.md)                                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                   |   Configure monitoring of SSL/TLS certificates validity. This option is relevant for HTTPS protocol only. <br/>  If omitted or set to null, SSL monitoring will be disabled and its previous configuration discarded.                                                                                                                                                                                                                                                |                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| `CustomHeaders`                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [][components.CustomHeaders](../../models/components/customheaders.md)                                                                                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                   |   Configure custom request headers to be sent with each availability test. It is possible to provide multiple headers with the same name.<br/>  If omitted, set to null or set to an empty array, no custom headers will be sent.                                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| `AllowInsecureRenegotiation`                                                                                                                                                                                                                                                                                                                                                                                                                                         | **bool*                                                                                                                                                                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                   |   Allow insecure SSL renegotiation which introduces a security risk in the communication process.<br/>  Checking this option could lead to exposing credentials to unauthorized entities and the possibility of unauthorized access, interception, or manipulation of sensitive data, compromising the integrity and security of the communication channel.<br/>  Available only with HTTPS check.<br/>  If omitted or set to null, insecure SSL renegotiation won't be allowed. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| `PostData`                                                                                                                                                                                                                                                                                                                                                                                                                                                           | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                   |   Configure data that will be sent as POST request body by the synthetic probe.<br/>  If omitted or set to null/empty string, the probe will send the usual GET requests.                                                                                                                                                                                                                                                                                            |                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |