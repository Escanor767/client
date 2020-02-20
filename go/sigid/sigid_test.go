package sigid

import (
	"testing"

	"github.com/keybase/client/go/jsonparserw"
	"github.com/keybase/client/go/kbcrypto"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	"github.com/stretchr/testify/require"
)

func testOne(t *testing.T, sig string, sigID string) {
	si, err := kbcrypto.DecodeArmoredNaclSigInfoPacket(sig)
	require.NoError(t, err)
	payload := si.Payload
	name, err := jsonparserw.GetString(payload, "client", "name")
	require.NoError(t, err)
	version, err := jsonparserw.GetString(payload, "client", "version")
	require.NoError(t, err)
	expectedSigID, err := keybase1.SigIDFromString(sigID, true)
	require.NoError(t, err)
	_, computedSigID, err := ComputeSigBodyAndID(&si, name, version)
	require.NoError(t, err)
	require.True(t, expectedSigID.Eq(computedSigID))
}

func TestSignatures(t *testing.T) {

	var tests = []struct {
		username string
		seqno    int
		sig      string
		sigID    string
	}{
		{
			"max", 689,
			"hKRib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEg7AkiR0GhYL2jZ6J16YMdkCv2ncoVa9M599aTk4/ndBYKp3BheWxvYWTFBDd7ImJvZHkiOnsiZGV2aWNlIjp7ImlkIjoiMjkxOWZjOGQ5YThlODU1ZmVkNTljOGU5YmM1MjNkMTgiLCJraWQiOiIwMTIxOTA0MDU4Njg2YjEwNThkZGNhMzJjMGIwOWI2NDc2NjA1MzZmMzY0M2JlMmIwYzcwM2MxM2M3NjM2MGNkYjg2YTBhIiwic3RhdHVzIjoxfSwia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTAxM2VmOTBiNGM0ZTYyMTIxZDEyYTUxZDE4NTY5YjU3OTk2MDAyYzhiZGNjYzliMjc0MDkzNWM5ZTRhMDdkMjBiNDBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwZWMwOTIyNDc0MWExNjBiZGEzNjdhMjc1ZTk4MzFkOTAyYmY2OWRjYTE1NmJkMzM5ZjdkNjkzOTM4ZmU3NzQxNjBhIiwidWlkIjoiZGJiMTY1Yjc4NzlmZTdiMTE3NGRmNzNiZWQwYjk1MDAiLCJ1c2VybmFtZSI6Im1heCJ9LCJtZXJrbGVfcm9vdCI6eyJjdGltZSI6MTU3OTczODQxMSwiaGFzaCI6IjE2NmFjNjA1OTkwNGFjOTQ1NjJhZWY5NjA0MmVlNDI4MmJjMzI5MzI4NmNlOGI4NDUyNTliMjQ0M2UzZTJkOTM4M2JjZjAyYmZhYTgwYzliNTVlOTZmMzIwNzgwNjM5MWViMjAzOGVlNzQwOTFjMjMwZGFmMThmZTYyNzgxZWFjIiwiaGFzaF9tZXRhIjoiYWU3YTNkMWFkMDZhZmM2ODU3MjIzM2I4ZmI0MjhhMDIyM2VkOGFlYTcwNDQyZGI4NTViOWNhOGFkZDQ2MTc3NCIsInNlcW5vIjoxNDMzMDA3Mn0sInN1YmtleSI6eyJraWQiOiIwMTIxOTA0MDU4Njg2YjEwNThkZGNhMzJjMGIwOWI2NDc2NjA1MzZmMzY0M2JlMmIwYzcwM2MxM2M3NjM2MGNkYjg2YTBhIiwicGFyZW50X2tpZCI6IjAxMjBlYzA5MjI0NzQxYTE2MGJkYTM2N2EyNzVlOTgzMWQ5MDJiZjY5ZGNhMTU2YmQzMzlmN2Q2OTM5MzhmZTc3NDE2MGEifSwidHlwZSI6InN1YmtleSIsInZlcnNpb24iOjF9LCJjbGllbnQiOnsibmFtZSI6ImtleWJhc2UuaW8gZ28gY2xpZW50IiwidmVyc2lvbiI6IjUuMS4wIn0sImN0aW1lIjoxNTc5NzM4NDEyLCJleHBpcmVfaW4iOjUwNDU3NjAwMCwicHJldiI6ImZmODM1MjFiZTBjZjIyNWVmZGQ3MmY0OTg2NjdhMWM1ZTEwZTFkMTkzMTY4MGI5MjI2ZGZjN2EwM2VkZTYzNWIiLCJzZXFubyI6Njg5LCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQLIPxQXsmzF2ay/hF708QwfRO9nDqMcriCm9fP36EzchNYNCc4HKVBJ/eWYiIldeMMo4iftwWtUxnc1mwcA6iwGoc2lnX3R5cGUgpGhhc2iCpHR5cGUIpXZhbHVlxCAKMI5a2CV84+RLhJBxoKT0sTrDJ5gwe2Kxeb3Tx8vYQ6N0YWfNAgKndmVyc2lvbgE=",
			"b1d9c9cd6f1dce5c572630124410221ecdab57a6e78a2e1b5fa5a9aa3277344c0f",
		},
		{
			"max", 218,
			"hKRib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEg00WLvs38DQrjn+wFcixuPol8FpIjg1l3qKogjfzZAtMKp3BheWxvYWTFDIF7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTAxM2VmOTBiNGM0ZTYyMTIxZDEyYTUxZDE4NTY5YjU3OTk2MDAyYzhiZGNjYzliMjc0MDkzNWM5ZTRhMDdkMjBiNDBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwZDM0NThiYmVjZGZjMGQwYWUzOWZlYzA1NzIyYzZlM2U4OTdjMTY5MjIzODM1OTc3YThhYTIwOGRmY2Q5MDJkMzBhIiwidWlkIjoiZGJiMTY1Yjc4NzlmZTdiMTE3NGRmNzNiZWQwYjk1MDAiLCJ1c2VybmFtZSI6Im1heCJ9LCJzaWJrZXkiOnsiZmluZ2VycHJpbnQiOiI0NDc1MjkzMzA2MjQzNDA4ZmE1OTU4ZGM2Mzg0N2I0YjgzOTMwZjBjIiwiZnVsbF9oYXNoIjoiNzkwOTczOWQyMTY0N2E1ODNhZjk3NTdlYmQ0YTY4MTdjZjllMWM2MGUwZDIxZjM4MDFhYzM3ODZiZTNlZTQ3MyIsImtleV9pZCI6IjYzODQ3QjRCODM5MzBGMEMiLCJraWQiOiIwMTAxN2E5Y2E0NDYxNzE4ZDU4OWRlNGQzNGRlOGRjNzAzNDcxNmNkNjk0MDk4OGY5OWI5ZmVkNmM1MjBjMzgyMDYxYjBhIiwicmV2ZXJzZV9zaWciOiItLS0tLUJFR0lOIFBHUCBNRVNTQUdFLS0tLS1cbkNvbW1lbnQ6IGh0dHBzOi8va2V5YmFzZS5pby9kb3dubG9hZFxuVmVyc2lvbjogS2V5YmFzZSBHbyAxLjAuMTYgKGRhcndpbilcblxueEEwREFBb0JMK0FjUlVOSTJqa0J5K0YwQU9JQUFBQUE2WHNpWW05a2VTSTZleUpyWlhraU9uc2laV3hrWlhOMFxuWDJ0cFpDSTZJakF4TURFelpXWTVNR0kwWXpSbE5qSXhNakZrTVRKaE5URmtNVGcxTmpsaU5UYzVPVFl3TURKalxuT0dKa1kyTmpPV0l5TnpRd09UTTFZemxsTkdFd04yUXlNR0kwTUdFaUxDSm9iM04wSWpvaWEyVjVZbUZ6WlM1cFxuYnlJc0ltdHBaQ0k2SWpBeE1qQmtNelExT0dKaVpXTmtabU13WkRCaFpUTTVabVZqTURVM01qSmpObVV6WlRnNVxuTjJNeE5qa3lNak00TXpVNU56ZGhPR0ZoTWpBNFpHWmpaRGt3TW1Rek1HRWlMQ0oxYVdRaU9pSmtZbUl4TmpWaVxuTnpnM09XWmxOMkl4TVRjMFpHWTNNMkpsWkRCaU9UVXdNQ0lzSW5WelpYSnVZVzFsSWpvaWJXRjRJbjBzSW5OcFxuWW10bGVTSTZleUptYVc1blpYSndjbWx1ZENJNklqUTBOelV5T1RNek1EWXlORE0wTURobVlUVTVOVGhrWXpZelxuT0RRM1lqUmlPRE01TXpCbU1HTWlMQ0ptZFd4c1gyaGhjMmdpT2lJM09UQTVOek01WkRJeE5qUTNZVFU0TTJGbVxuT1RjMU4yVmlaRFJoTmpneE4yTm1PV1V4WXpZd1pUQmtNakZtTXpnd01XRmpNemM0Tm1KbE0yVmxORGN6SWl3aVxuYTJWNVgybGtJam9pTmpNNE5EZENORUk0TXprek1FWXdReUlzSW10cFpDSTZJakF4TURFM1lUbGpZVFEwTmpFM1xuTVRoa05UZzVaR1UwWkRNMFpHVTRaR00zTURNME56RTJZMlEyT1RRd09UZzRaams1WWpsbVpXUTJZelV5TUdNelxuT0RJd05qRmlNR0VpNkN3aWNtVjJaWEp6WlY5emFXY2lPbTUxYkd4OUxDSjBlWEJsSWpvaWMybGlhMlY1SWl3aVxuZG1WeWMybHZiaUk2TVgwc0ltTnNhV1Z1ZENJNmV5SnVZVzFsSWpvaWEyVjVZbUZ6WlM1cGJ5Qm5ieUJqYkdsbFxuYm5RaUxDSjJaWEp6YVc5dUlqb2lNUzR3TGpFMkluMHNJbU4wYVcxbElqb3hORFk1TXpnNU5ESTJMQ0psZUhCcFxuY21WZmFXNGlPalV3TkRVM05qQXdNQ3dpYldWeWEyeGxYM0p2YjNRaU9uc2lZM1JwYldVaU9qRTBOamt6T0RreFxuTlRVc0ltaGhjMmdpT2lKa1pXRTBaRFpoTTJRMk1XVTNaV000Wm1NMVpUY3pNemd4TkdNd016YzBNelkwT0dVeVxuWWpsaU56TTNZamxsTWpSaE9EazVZalF4Wm1abVl6VmhabUxuWVRFNU5ETTROelZsWkRVME4yUTVaV1V4TVdRM1xuTjJKaE5UYzBZMlZtWVRkak1qRTBZMkkxTnpKbFptUTBZVFV4TkRRMU5qQmpNVEZpTXpFMFpEY3paR01pTENKelxuWlhGdWJ5STZOVFF3TkRRM2ZTd2ljSEpsZGlJNkltRXlaakExTlRGbFpEa3pPVGN5TVRrM1ptVTRaRGM0T1RsaVxuTmpjeE0yVmhZekpoTmpYbE1EYzNNbUV5WTJFNE5UVXlaV000TVRVNE1tTTFabVE1TTJabUlpd2ljMlhrY1c1dlxuSWpveU1UZ3NJblJoWnlJNkl1TnphV2R1WVhSMWN1RmxJdUI5QU1MQlhBUUFBUW9BRUFVQ1Y1VWFjZ2tRTCtBY1xuUlVOSTJqa0FBSlBTRUFDTzZRU3I0YkRYK3VkclNnVHY4MHk2UzhMaHdoTFhYOW9BWEZTeHR6YVZGRjdYN0ZMb1xuWjUrZUhJc2ZTNHFOKzcwczhGcGxISkplMGxidHRicnlrdG5mdmRKVkdmMFBRaFM5Nnk3QjUrMWpxQXdlTC93cFxudEdNZG9XRDZneHZnZWw1bjF0U1NPK25ySzBRMlExR1FtcVBGUVllVXJwbk1LcEdGZTRYMTZncXlXK2NMT0k0Y1xuY3dMaElNbWRUMng0V3l4ZjdlZFA4WndncVF5b0pJM0w2V2tocjR0Wi9FVEZWaGdCQm94Uk41ZEtUQ1RnelFYZFxueC9zQktxY3l1K0s1Mjg0ZmJvV3Z1L3IzNzZYdk1LV3NBMjNlS0tZYjNxVTVLU1I3SzhXT1pJZDZQNTM3bVVybFxuejg5MXRRMGc3eEQ4UFdKK3B2QnhwOXQ2UHNkSndneW0wM20wWVZFOE9YWTA2UDJSQ010VGFnVncvVVNGU3lPTlxuQko3ajN4WFRnQ3dncG9WeTRDSTdqUEJ1aUpNNVljeXpSS0l4Qk1XNmZQbW5xdUg5MW9HNlNqSGNFekpMTlFCa1xuL1F0UXVmdGNzd0w2YjloaTNUd3lvRmt2SzVUWndldjYxMk9HRFhXdERhYlBTZWZtaGtrd05BZFowa3JoZzVJd1xuSlYzYllkMnpkL2lhZ1dZdWJjSXNuNTNWeDkzUlh4eWFSdUNST1JFUE9ZOGcyYmhHc3A0THhFUnpDUEcwT3pIR1xueHE5YnliZzdKREdWdHRxWUhzb2FFUzN6RDh4djVBQWxXUEJPZUhRUlhjRDgxZFNIOFFFWmxJcVZMamNWeVl6c1xuOVRjRFRPY08xUGYydzBBZVhVdS9XZStteUo3Wi9WUi9ZOFNZZ3h0dnBEQS8yZ2pLR3NlZXplNlZYQT09XG49Qnd1MVxuLS0tLS1FTkQgUEdQIE1FU1NBR0UtLS0tLVxuIn0sInR5cGUiOiJzaWJrZXkiLCJ2ZXJzaW9uIjoxfSwiY2xpZW50Ijp7Im5hbWUiOiJrZXliYXNlLmlvIGdvIGNsaWVudCIsInZlcnNpb24iOiIxLjAuMTYifSwiY3RpbWUiOjE0NjkzODk0MjYsImV4cGlyZV9pbiI6NTA0NTc2MDAwLCJtZXJrbGVfcm9vdCI6eyJjdGltZSI6MTQ2OTM4OTE1NSwiaGFzaCI6ImRlYTRkNmEzZDYxZTdlYzhmYzVlNzMzODE0YzAzNzQzNjQ4ZTJiOWI3MzdiOWUyNGE4OTliNDFmZmZjNWFmYmExOTQzODc1ZWQ1NDdkOWVlMTFkNzdiYTU3NGNlZmE3YzIxNGNiNTcyZWZkNGE1MTQ0NTYwYzExYjMxNGQ3M2RjIiwic2Vxbm8iOjU0MDQ0N30sInByZXYiOiJhMmYwNTUxZWQ5Mzk3MjE5N2ZlOGQ3ODk5YjY3MTNlYWMyYTY1MDc3MmEyY2E4NTUyZWM4MTU4MmM1ZmQ5M2ZmIiwic2Vxbm8iOjIxOCwidGFnIjoic2lnbmF0dXJlIn2jc2lnxEDJYzrSNKm/vmkIv7SRH72oQkyQVgHbFPC4ZeYCeeBq2xUC+OwueYrW+vSMD8Nz5cCeZhJ3nNhuX17CrvAYbTgCqHNpZ190eXBlIKRoYXNogqR0eXBlCKV2YWx1ZcQgTDTCul2YiKRGX/Hr5BJ7Mip9Vvx4PblCwvnYUSBtJ0WjdGFnzQICp3ZlcnNpb24B",
			"d7d09db9572687ca6e90c5f1a9003b310e04e435f5b8d270c2692db75122b5fb0f",
		},
		{
			"max", 199,
			"hKRib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgXS40pxPgKl7FDD9B3NzaC5od50u1VETPa0379UFIfSYKp3BheWxvYWTFBEd7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTAxM2VmOTBiNGM0ZTYyMTIxZDEyYTUxZDE4NTY5YjU3OTk2MDAyYzhiZGNjYzliMjc0MDkzNWM5ZTRhMDdkMjBiNDBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwNWQyZTM0YTcxM2UwMmE1ZWM1MGMzZjQxZGNkY2RhMGI5YTFkZTc0YmI1NTQ0NGNmNmI0ZGZiZjU0MTQ4N2QyNjBhIiwidWlkIjoiZGJiMTY1Yjc4NzlmZTdiMTE3NGRmNzNiZWQwYjk1MDAiLCJ1c2VybmFtZSI6Im1heCJ9LCJ0cmFjayI6eyJiYXNpY3MiOnsiaWRfdmVyc2lvbiI6MiwibGFzdF9pZF9jaGFuZ2UiOjE0NjQ4OTI4MzQsInVzZXJuYW1lIjoiY3J1ZGRlciJ9LCJpZCI6IjEyMzRmNDc4MmUwOWFiOWNhOTFlZDZiZWViNjM3ZDE5Iiwia2V5Ijp7ImtpZCI6IjAxMjAyZWVmMjg2YTJiZTU4Nzk3YTc3MTk1MDE5NGVhNDQ4OTU3ZTUwYzA4YTJjM2FiNTNkZDM5NzljNDkxMGQ0M2FkMGEifSwicmVtb3RlX3Byb29mcyI6W10sInNlcV90YWlsIjp7InBheWxvYWRfaGFzaCI6IjFlODcyNDVlMmMzMTU3NjM5NWNiNWUzYWFkMjg4MDg0YzZlODllNDQ0ZjBmNDRhYzNmMmJhNmNiODk0OWVmNTMiLCJzZXFubyI6Nywic2lnX2lkIjoiODEyNGRlZWVjOTAyY2E5OWUxZjQ3YzZmZDQ1OTAzNWFhYjVhMWUwZjgxMmE3MWMwNDdjMDRmZWQyOGJiMzIzMDBmIn19LCJ0eXBlIjoidHJhY2siLCJ2ZXJzaW9uIjoxfSwiY2xpZW50Ijp7Im5hbWUiOiJrZXliYXNlLmlvIGdvIGNsaWVudCIsInZlcnNpb24iOiIxLjAuMTYifSwiY3RpbWUiOjE0NjU5MjE5NjMsImV4cGlyZV9pbiI6NTA0NTc2MDAwLCJtZXJrbGVfcm9vdCI6eyJjdGltZSI6MTQ2NTkyMTg2MywiaGFzaCI6Ijg3MmNlZWI0YmE1ODUzZGNiNTE4OGZhYTc1ZGZiZjVkZGNlNjliMTAxOTgwOGNjNjM4ZTg4YWZkYjA3OTJiZDEwZmVkMjk0MDg3MDQ0ZTQ4MGYzOTg4MDQ4ZDViYmM2ZGEzMzlkM2QxYWE1MzgwMzk1OTVkMWQxYzc5NjRhNjJmIiwic2Vxbm8iOjQ4Mzc5Nn0sInByZXYiOiJmNzIxMjgyNjBhNzVlMjE5YjQwOWEyNzgyNzIzMzYzMmIwMjQyMGE2YjM5NGJlMDIyZjAxMzQ2MTRkMmRjNDM3Iiwic2Vxbm8iOjE5OSwidGFnIjoic2lnbmF0dXJlIn2jc2lnxEAu/Fbt4R6mPlNxdjutjXsTemQ+p9Nur1ZSZxToQs2DCMlza9CfeBzgw1gkpxOSYWhzv+VZUyct0fPKsu5CY+4IqHNpZ190eXBlIKRoYXNogqR0eXBlCKV2YWx1ZcQgXr7GSfR0hNb2WHt5Ic7htUKoFngEwHPbP7P4Px8/8u+jdGFnzQICp3ZlcnNpb24B",
			"a183de7e98d179d929e3704ee592e84e1db0cdabda1ec448fbef51cbc0b3f9370f",
		},
		{
			"max", 120,
			"g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEg00WLvs38DQrjn+wFcixuPol8FpIjg1l3qKogjfzZAtMKp3BheWxvYWTFBpl7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTAxM2VmOTBiNGM0ZTYyMTIxZDEyYTUxZDE4NTY5YjU3OTk2MDAyYzhiZGNjYzliMjc0MDkzNWM5ZTRhMDdkMjBiNDBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwZDM0NThiYmVjZGZjMGQwYWUzOWZlYzA1NzIyYzZlM2U4OTdjMTY5MjIzODM1OTc3YThhYTIwOGRmY2Q5MDJkMzBhIiwidWlkIjoiZGJiMTY1Yjc4NzlmZTdiMTE3NGRmNzNiZWQwYjk1MDAiLCJ1c2VybmFtZSI6Im1heCJ9LCJ0cmFjayI6eyJiYXNpY3MiOnsiaWRfdmVyc2lvbiI6MTEsImxhc3RfaWRfY2hhbmdlIjoxNDI4NTkwOTAyLCJ1c2VybmFtZSI6InJqIn0sImlkIjoiMjYzMDhhZmIzZDhjZWU0MjYwMzFlMzAwYmIwOWEzMDAiLCJrZXkiOnsia2V5X2ZpbmdlcnByaW50IjoiM2ZkODFjY2UwNWQ5YWUwYmM4ZmM1YWQyN2QxYzFmNTVhODU0OGRlNiIsImtpZCI6IjAxMDFlN2Y4NGFhNjc3YWIwNDFmOGNiYjJhYzFmNDI4ZWI5N2RmMjY3YTQxN2Q4NmQ5MmE4MDgwNWY1OTEwNWY3MWZjMGEifSwicGdwX2tleXMiOlt7ImtleV9maW5nZXJwcmludCI6IjNmZDgxY2NlMDVkOWFlMGJjOGZjNWFkMjdkMWMxZjU1YTg1NDhkZTYiLCJraWQiOiIwMTAxZTdmODRhYTY3N2FiMDQxZjhjYmIyYWMxZjQyOGViOTdkZjI2N2E0MTdkODZkOTJhODA4MDVmNTkxMDVmNzFmYzBhIn1dLCJyZW1vdGVfcHJvb2ZzIjpbeyJjdGltZSI6MTQyODU5MDg4NiwiY3VyciI6ImUwOGVmNTYxYmJiNTZjMmZhMmM3NGQ0NDMzMjgwNzcxMDVjMWZmNzMwNGZlNThmMWJmZTA5NDgxMDczOTc4NTMiLCJldGltZSI6MTU4NjI3MDg4NiwicHJldiI6ImFkOGU2NjJiMDIzYjVkOGUyMDVjNzA1MDdiNGFhN2IzNGRmNGFjMGM2ZGE4MTk4Mzg4YzQyM2ZhZmU2MTA0MjQiLCJyZW1vdGVfa2V5X3Byb29mIjp7ImNoZWNrX2RhdGFfanNvbiI6eyJuYW1lIjoidHdpdHRlciIsInVzZXJuYW1lIjoibWV0YWJyZXcifSwicHJvb2ZfdHlwZSI6Miwic3RhdGUiOjF9LCJzaWdfaWQiOiIzZTljZmZmNTg3ODY4YTc2MmFkMjUwNTgyODE0NGIxZGIzZGFjZmE1MTJmNDQxOWMyNDViZDNhZDllODU5MTVhMGYiLCJzaWdfdHlwZSI6Mn1dLCJzZXFfdGFpbCI6eyJwYXlsb2FkX2hhc2giOiJlMDhlZjU2MWJiYjU2YzJmYTJjNzRkNDQzMzI4MDc3MTA1YzFmZjczMDRmZTU4ZjFiZmUwOTQ4MTA3Mzk3ODUzIiwic2Vxbm8iOjcsInNpZ19pZCI6IjNlOWNmZmY1ODc4NjhhNzYyYWQyNTA1ODI4MTQ0YjFkYjNkYWNmYTUxMmY0NDE5YzI0NWJkM2FkOWU4NTkxNWEwZiJ9fSwidHlwZSI6InRyYWNrIiwidmVyc2lvbiI6MX0sImNsaWVudCI6eyJuYW1lIjoia2V5YmFzZS5pbyBnbyBjbGllbnQiLCJ2ZXJzaW9uIjoiMS4wLjAifSwiY3RpbWUiOjE0NDUwNDQ2NDEsImV4cGlyZV9pbiI6NTA0NTc2MDAwLCJtZXJrbGVfcm9vdCI6eyJjdGltZSI6MTQ0NTA0NDM3MCwiaGFzaCI6IjgxZTMwZjhkNDQ4MzUyZGQzODc5MmViM2M1M2U2ODcwY2Y1NmFhNTZiYzllNDU0ZDhjOTE1MzFmZmFkNDUxYzNlNDczYTJiNWRjODcyYzM1MzgxZmVkNWJkM2ViODUzNTkxYWQ1ZDk2NGU1Yzc4ZDhkYzAxNTk5NGVkMjZlYjJhIiwic2Vxbm8iOjMwNzcxNX0sInByZXYiOiJlNDYzYjc3ODdkMjc2N2VhNTQ5NjMyNmI4MmJiNjkyNDJhOWU2MjE0NzQyMzBmZTE0Y2VlN2JlMDk1YmViZWZlIiwic2Vxbm8iOjEyMCwidGFnIjoic2lnbmF0dXJlIn2jc2lnxEBRF1xCPvQ7Lq3tOeRMIAIiI8lHAnEMRFMsLl8/7o0dgpGLOUvWNLac7tnawmse4MBfhfIn/BHfsM2SYfUUBYULqHNpZ190eXBlIKN0YWfNAgKndmVyc2lvbgE=",
			"4badf05d87bd3b0041ad2aaba1ad866268834811680568a191a31882257df2300f",
		},
		{
			"00dani", 123,
			"hKRib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgr/lcvDb3c1+eD199Q8eEG2ZG1gZ6whVsvjlfQ09EzfkKp3BheWxvYWTFBAJ7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwYmNlZDExODZkMDkwM2RhMWYyOGQwN2MwMDkzMDg0ZmI4MWYxMDcwMjhiMjViNDY5MTVjM2Q4NDVhNTBjZDc5MzBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwYWZmOTVjYmMzNmY3NzM1ZjllMGY1ZjdkNDNjNzg0MWI2NjQ2ZDYwNjdhYzIxNTZjYmUzOTVmNDM0ZjQ0Y2RmOTBhIiwidWlkIjoiZmUxZTBjOTY2NTQ2MDNjNTg4MGIzZjU3MWM4NjI5MTkiLCJ1c2VybmFtZSI6IjAwZGFuaSJ9LCJtZXJrbGVfcm9vdCI6eyJjdGltZSI6MTU1NDY5MjQ3NiwiaGFzaCI6IjZmNjgzODA4NGI0MzE2Njc5NjM5ZDcxNTNjNGY3YmIwM2VhNzU0MjU1YTRhNjIwMzQxYzA2MTcwMTM0MWUyMDdmZDE4ZDk3NjhjNjliMGVlMzkxYWU0N2QwMzVlZmE5NzYxODQ1M2YwZjg1NjdlYmFlNzUzMmZjZDA4NWFiNzlkIiwiaGFzaF9tZXRhIjoiZWYzNjUyNDdkYTAyM2I3ZGFlZDYzODdmN2Y4Yjk3NzRhNDg3ODA4Y2Q3YzA1ZmNhOTFmMmRmMWY4YzAzZDkwZSIsInNlcW5vIjo1MTE1MDg3fSwicGdwX3VwZGF0ZSI6eyJmaW5nZXJwcmludCI6IjgzZjNkY2VjOThkNTIyYjZhMzhhZjVkOTI3ZDA3NmQyYWNhN2JhYmUiLCJmdWxsX2hhc2giOiJkMmFlZDU4MDc4NzNlODc5NDg0ODMxYTJlMDUyZDJmMGNmYzM1ZmFkZjgyNjNhZDc2ODg1N2RlNzc3MTQ4MzVmIiwia2V5X2lkIjoiMjdEMDc2RDJBQ0E3QkFCRSIsImtpZCI6IjAxMDFjMjc4N2JlNTUxNjYzYzg2MDI3YmE3ODY5MzdmNTZkN2UzNjUwYzE3ZTUzNTkxNmRiYWMzM2M5ZWRiNGQ4NTUxMGEifSwidHlwZSI6InBncF91cGRhdGUiLCJ2ZXJzaW9uIjoxfSwiY2xpZW50Ijp7Im5hbWUiOiJrZXliYXNlLmlvIGdvIGNsaWVudCIsInZlcnNpb24iOiIzLjEuMiJ9LCJjdGltZSI6MTU1NDY5MjUxOCwiZXhwaXJlX2luIjo1MDQ1NzYwMDAsInByZXYiOiI2ZTVjOGUwMTk3NjA3NzQ1YTFlYTNhMDlhYjgxMjdiZmMxNTI0N2ZlZTI3NDUyNDFiNTE2ZTU2MDM0NmMyZjEzIiwic2Vxbm8iOjEyMywidGFnIjoic2lnbmF0dXJlIn2jc2lnxEBw/Lo8QUqgvlGXY4961coKy3UuOhxAhvrpUjjABvxtJmZqdy3WCrz9KkZvmrwfItkN1nAVV+f7ykjRa3VxLAgCqHNpZ190eXBlIKRoYXNogqR0eXBlCKV2YWx1ZcQg2O1YYOHk8zdytvKhK1jKnnSVB1JaaQ3UU28joqj3BASjdGFnzQICp3ZlcnNpb24B",
			"7dace7012cc3825a8ae8be294cab5fb1f24929db93197b795aece5b016cd9cdb0f",
		},
		{
			"00dani", 2,
			"g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgvO0RhtCQPaHyjQfACTCE+4HxBwKLJbRpFcPYRaUM15MKp3BheWxvYWTFA+h7ImJvZHkiOnsiZGV2aWNlIjp7ImlkIjoiNWNjNjc4YTcyMjI1NzFjZDk4NTBiNTQ3MTkyOTcyMTgiLCJraWQiOiIwMTIxNDc4NDMwNzFhMGE1OGU3MzhjMjYzZjc2ZmY4ZDVkNzhkNjVmYzdiNzE4YTg2ZTdiN2FkZmVmNmU1Njc4Y2QyZjBhIiwic3RhdHVzIjoxfSwia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwYmNlZDExODZkMDkwM2RhMWYyOGQwN2MwMDkzMDg0ZmI4MWYxMDcwMjhiMjViNDY5MTVjM2Q4NDVhNTBjZDc5MzBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwYmNlZDExODZkMDkwM2RhMWYyOGQwN2MwMDkzMDg0ZmI4MWYxMDcwMjhiMjViNDY5MTVjM2Q4NDVhNTBjZDc5MzBhIiwidWlkIjoiZmUxZTBjOTY2NTQ2MDNjNTg4MGIzZjU3MWM4NjI5MTkiLCJ1c2VybmFtZSI6IjAwZGFuaSJ9LCJzdWJrZXkiOnsia2lkIjoiMDEyMTQ3ODQzMDcxYTBhNThlNzM4YzI2M2Y3NmZmOGQ1ZDc4ZDY1ZmM3YjcxOGE4NmU3YjdhZGZlZjZlNTY3OGNkMmYwYSIsInBhcmVudF9raWQiOiIwMTIwYmNlZDExODZkMDkwM2RhMWYyOGQwN2MwMDkzMDg0ZmI4MWYxMDcwMjhiMjViNDY5MTVjM2Q4NDVhNTBjZDc5MzBhIn0sInR5cGUiOiJzdWJrZXkiLCJ2ZXJzaW9uIjoxfSwiY2xpZW50Ijp7Im5hbWUiOiJrZXliYXNlLmlvIGdvIGNsaWVudCIsInZlcnNpb24iOiIxLjAuMTYifSwiY3RpbWUiOjE0NjcwNzg2ODUsImV4cGlyZV9pbiI6NTA0NTc2MDAwLCJtZXJrbGVfcm9vdCI6eyJjdGltZSI6MTQ2NzA3ODEzOSwiaGFzaCI6IjFkNTdjNzc1MWZhYjA1NjVmNjcyZjA1OTg4OGE4NWQ4NGUzZTJjZGNmNmUyNjk3MGM3ZjVjZDM0Y2NlMjE1NzAwYzhhZTEzMDQ1MDk0OWUzNTk4ZThlODc5OTY5ZjdmOTRhOTUzNGY1MWNjMGQ5NDlmZjZjMWZmMTFhZWJjYTljIiwic2Vxbm8iOjUwMTU5MX0sInByZXYiOiJkZjU3YTJkZjA5NzVjNTkyZjgzZmJkZGFmZjdmNWM4MTllOTFkYzYxZGUzYzY5ODQ1NzE2NzdiMWUwYTAxNjg5Iiwic2Vxbm8iOjIsInRhZyI6InNpZ25hdHVyZSJ9o3NpZ8RAahEKAwz9Ls+ahooouPFMHNjCgkAYD3j9cieBizsM9u378f5rxl52yWKrGb78VUWOE3/kHf/Tt1AO9TOUyNH+D6hzaWdfdHlwZSCjdGFnzQICp3ZlcnNpb24B",
			"bd2ba14da159c9bdb1c9eece386ec1f8aba37fc9f2848c79052ee52055c6ddcd0f",
		},
		{
			"00dani", 4,
			"g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgL48JmNbp2t3wvVwSlcD6AF8IawVLE757igZov2jW9s8Kp3BheWxvYWTFA+h7ImJvZHkiOnsiZGV2aWNlIjp7ImlkIjoiM2RiMzllYzkxMWUyOTJlZTQ3MDE1MzVkY2RiNjcyMTgiLCJraWQiOiIwMTIxMzgxNTFkNDViYmZkODA3ZDgxYTk0ZjkxOGI0OTgwZDlhZDJkMWU1YmRkNDRlZmY5ODAzZjJlMjZlMzc4ODM0MjBhIiwic3RhdHVzIjoxfSwia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwYmNlZDExODZkMDkwM2RhMWYyOGQwN2MwMDkzMDg0ZmI4MWYxMDcwMjhiMjViNDY5MTVjM2Q4NDVhNTBjZDc5MzBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwMmY4ZjA5OThkNmU5ZGFkZGYwYmQ1YzEyOTVjMGZhMDA1ZjA4NmIwNTRiMTNiZTdiOGEwNjY4YmY2OGQ2ZjZjZjBhIiwidWlkIjoiZmUxZTBjOTY2NTQ2MDNjNTg4MGIzZjU3MWM4NjI5MTkiLCJ1c2VybmFtZSI6IjAwZGFuaSJ9LCJzdWJrZXkiOnsia2lkIjoiMDEyMTM4MTUxZDQ1YmJmZDgwN2Q4MWE5NGY5MThiNDk4MGQ5YWQyZDFlNWJkZDQ0ZWZmOTgwM2YyZTI2ZTM3ODgzNDIwYSIsInBhcmVudF9raWQiOiIwMTIwMmY4ZjA5OThkNmU5ZGFkZGYwYmQ1YzEyOTVjMGZhMDA1ZjA4NmIwNTRiMTNiZTdiOGEwNjY4YmY2OGQ2ZjZjZjBhIn0sInR5cGUiOiJzdWJrZXkiLCJ2ZXJzaW9uIjoxfSwiY2xpZW50Ijp7Im5hbWUiOiJrZXliYXNlLmlvIGdvIGNsaWVudCIsInZlcnNpb24iOiIxLjAuMTYifSwiY3RpbWUiOjE0NjcwNzg2ODYsImV4cGlyZV9pbiI6NTA0NTc2MDAwLCJtZXJrbGVfcm9vdCI6eyJjdGltZSI6MTQ2NzA3ODEzOSwiaGFzaCI6IjFkNTdjNzc1MWZhYjA1NjVmNjcyZjA1OTg4OGE4NWQ4NGUzZTJjZGNmNmUyNjk3MGM3ZjVjZDM0Y2NlMjE1NzAwYzhhZTEzMDQ1MDk0OWUzNTk4ZThlODc5OTY5ZjdmOTRhOTUzNGY1MWNjMGQ5NDlmZjZjMWZmMTFhZWJjYTljIiwic2Vxbm8iOjUwMTU5MX0sInByZXYiOiIzZDMzZTUzYTRlOGFlNTI0ODU4ZmY1N2Q2MGQxOWU1MTNhZDM3ZjhkMmFiMDc1ZDczMWNlNzFmNzJlNzkyM2Q0Iiwic2Vxbm8iOjQsInRhZyI6InNpZ25hdHVyZSJ9o3NpZ8RAwsjrKy5PsaV+QzoMVIzIZo83dqhO4V3mbFKaMbXAfSzDbhYo5AxRq7FFG1QujSHWiwwKjS8JeE8eDnsSJjwkDahzaWdfdHlwZSCjdGFnzQICp3ZlcnNpb24B",
			"1009bf4aac80d81e7cd734392a158475fd97b75fac78bf11c840cc876ab051b10f",
		},
	}

	for _, test := range tests {
		t.Logf("testing at link %s %d\n", test.username, test.seqno)
		testOne(t, test.sig, test.sigID)
	}

}
