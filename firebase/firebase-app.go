package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)

func GetFirebaseApp() (*firebase.App, error) {
	b := []byte(`{
  "type": "service_account",
  "project_id": "cafu-auth-test",
  "private_key_id": "48ad854a53d7abd711bafdb00f0c8adaf0c00342",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDrnSMUJk8orVrb\nMJTdRWKkYj1oQi6wMCF3DAZ16JLYqlvTZ/YQXHzer81TwrY/5KB6V33uiFklT96O\n0GIgSU/o45cwmHjZVX2fxlYUnCF/420N+VsFlFME+Ah11f+P40Z7bRF2RO1e9cur\nhWJHTh/mRaOw85YEAn4lWTy2xnDCDhzJPlMS72zuHCV1x4CwrWEzvdQ3EVv9EJn3\nls0txcbLDOslcD5OGr0SRHiTgTkkZrqlvdr7H9hmbofY64Z6mgLQ8Xx2i3J7MspL\nti1iZOVi4BV8GdPMdVT8QshkXPtl1orf72tmeYPscb8gZ+cFPGnrGEnZes21H3ba\nWtlRAA2fAgMBAAECggEAAo5e3g4YV8GhtbUOOGRTUKM8NVBaf1DWXBtGSqsuRKl8\ne4uqML1AIY5YtOrs3T6Wux0N4bOa+/uk13nNRrpk4ortjfppfTofNceKEYksF94Y\nxckJwsK8L9adCOI8dowprJcfVGAlDx2tivvUeB+R2wawYvARlZONWMxiCz2V5ZUK\nR4fGAVWjEAHGKUtT5UpVxFcf9I21+KyymXzsTiplb3VN6U8alRCkuodxx+Ps8jep\nAosggHFC876HI8Gb2DDteSfHkM9zc12ACvdmLASiJX1tRY4+bkdtEb0Me7AnL4S2\nbnYvID4TRMkWtMK8t4uNcIb9PPuKTIQHSn7PcdLPGQKBgQD//X5Jksqar+XMimKK\nP/8v91dZWVKZZ3M3ZBxbQftFtY9fGxacgPCGxzzkLIj9SpaWhUsRSyyeTxc5NUuw\nFmeb6jEB3g58IAoJO20kw92+gWtyypjp4GN7pc1azgG8Z+NSnlZFiAwL4vaBiX9z\naXz3Q8XJntZCYA3GnElvq7YJowKBgQDrn3G2Sed3/0LjhD4CGXa5GBLXKAl2eUod\nBLU1ZlFPFqoPmaWknv8UYrsi3SyANcTlw4Esf22BBSpXicp3xMYlV1vDeYRt9gBH\n6q4o38SpiWU9Ke/Q49cIzlCfXq2+kfhHUNvoJUlHesU8NBfOmdm8y+TEpmEcxl5z\nNE9dLPBj1QKBgQCCuzWMY0/e0ae83otWaTpsh5dvROOe/irkBddJb5fgUMmIW3Af\nuuWbHxLkwC2xlS70bdwZV9fQfTmx+JrPsZF2GZKvjnS9RSt67YNevDx5NuTlubt/\n7R1BHzqRfhV6319DaSNL3R+xzlXvb4XMLfoo2M6gmeJ4AU/ntxb11a5gmwKBgD2N\nCUd1RuIx7GdCxHAH9tujTkimTNKGHLY35J5ufA8M/J/Wo069uiS5JemVnWpgoW7k\nA3p8Tz7E+qqjdCCVZq2ahH4r4ExZVvlydMx4qAchvTCdI+iNnD7Qpn77XKfLue/n\nT4r1Mo9Zoux3L0DpL9gB62DzlarOpx2xLIpKkwGBAoGBAPsVvXfgkbrG+HILImil\n9FcC4QzNWvA9+T7nU4akjjrJR4wqO3hozffY/PmP6U1o+uHcmvRlDgzo9WUXv0Rk\n9dt0sqYiphER+iVjvPLFbBughrpoOb4jHP29bxedcgNJ+cNSLLFzyW+NmUgN1faG\nGiSiWJO6WCSUFMsM+cBUr3Ng\n-----END PRIVATE KEY-----\n",
  "client_email": "firebase-adminsdk-49bft@cafu-auth-test.iam.gserviceaccount.com",
  "client_id": "114741875966383783692",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-49bft%40cafu-auth-test.iam.gserviceaccount.com"
}`)

	opt := option.WithCredentialsJSON(b)
	config := &firebase.Config{ProjectID: "cafu-auth-test"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
		return nil, err
	}
	return app, err
}
