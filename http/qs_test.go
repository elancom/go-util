package http

import (
	"fmt"
	"github.com/elancom/go-util/json"
	"log"
	"testing"
)

func Test1(t *testing.T) {
	//s := "aweme_id=7112283071847599396&type=1&channel_id=0&city=450200&activity=0&item_type=0&is_commerce=0&iid=2168705416692471&device_id=1852046067902142&ac=wifi&channel=tengxun_1128_1025&aid=1128&app_name=aweme&version_code=210400&version_name=21.4.0&device_platform=android&os=android&ssmix=a&device_type=HD1900&device_brand=OnePlus&language=zh&os_api=22&os_version=5.1.1&manifest_version_code=210401&resolution=1080*1920&dpi=480&update_version_code=21409900&_rticket=1657901958289&package=com.ss.android.ugc.aweme&mcc_mnc=46000&cpu_support64=false&host_abi=armeabi-v7a&ts=1657901957&appTheme=light&app_type=normal&need_personal_recommend=1&is_guest_mode=0&minor_status=0&is_android_pad=0&cdid=da52f4b3-2554-4da0-a333-6533464e4c4f"
	//s := "manifest_version_code=210201&app_type=normal&iid=3312193896195406&channel=xiaomi_2329_64_1104&device_type=MI%25208&language=zh&cpu_support64=true&host_abi=arm64-v8a&resolution=1080*2115&update_version_code=21209900&cdid=ba87fed7-1d55-4c12-8159-0469d9c1bc87&os_api=29&dpi=440&ac=wifi&device_id=1341869058957783&os=android&os_version=10&version_code=210200&app_name=douyin_lite&version_name=21.2.0&device_brand=Xiaomi&ssmix=a&device_platform=android&aid=2329&city=441400&device_token=AAA632NRTCZOZHGBIXL3T6AXK3Q5543VGMZ5O5ZXA35HDA5P3PJVHRFTMCSAVPEJ2BSR42K6CKAL23UBNPTAHT7X2BED7MYXX5SGG4B34E6YPBNSILUZUDUG7Z6RK&magic=0d40bb21e12378fb8ecc2fed09bba721662820fd06fdcfa676236a9ac179b6ab&lanusk=cd67a1558e3542c550c260b69adcd6d7b996d19b3266f0c5cba997724a37324984609486196db2930052c33d23c14e2f7ff21463cb983c77898fe100b34a3cd8&x-tt-token=009f08a8f1f0d099183867730d39082fa005acf3721f59f389e1d5b2c0d2f744d0cb37bb9238696a4afdd61ca5efc73a7bb5a25953b598932cca87ef3a750206223b5c94b11bd4599eb76e54f2c32f6098983f84a762f3264eeda5c75e80ec553714a-1.0.1"
	s := "_rticket=1657953363301&ac=wifi&activity=0&aid=1128&appTheme=light&app_name=aweme&app_type=normal&aweme_id=7110022959300873509&cdid=da52f4b3-2554-4da0-a333-6533464e4c4f&channel=tengxun_1128_1025&channel_id=0&city=450200&cpu_support64=false&device_brand=OnePlus&device_id=1852046067902142&device_platform=android&device_type=HD1900&dpi=480&host_abi=armeabi-v7a&iid=2168705416692471&is_android_pad=0&is_commerce=0&is_guest_mode=0&item_type=0&language=zh&manifest_version_code=210401&mcc_mnc=46000&minor_status=0&need_personal_recommend=1&os=android&os_api=22&os_version=5.1.1&package=com.ss.android.ugc.aweme&resolution=1080*1920&ssmix=a&ts=1657953363&type=1&update_version_code=21409900&version_code=210400&version_name=21.4.0"
	log.Println(json.ToJsonStrIndent(ParseQuery(s)))
}

func Test2(t *testing.T) {
	a := "[{\"Cookie\":\"tt_webid=640da8fc47f8f9d73458930bd852fdf9; d_ticket=db449fc73ef367fb9f0686d159af7a520ebf3; multi_sids=3835528075285980%3A3adf0f57a1c8cb7219b4b99e0858061d; odin_tt=f0832e411ee37a92238765ae1a43117b8f2c74d104aeb4b84a8e1d6b4e6f68ddfa66b2079faf132fe13b51b53a9462f4cd1200134d58314d89e55ec63f41ffd7109bf34354dfc350fd3f622aa3e293b0; n_mh=7bZQH-iTWqDCSUvi6Rb7djaTLmtqRGTkjzHVsImzoSg; sid_guard=3adf0f57a1c8cb7219b4b99e0858061d%7C1657900918%7C5184000%7CTue%2C+13-Sep-2022+16%3A01%3A58+GMT; uid_tt=570708be911049a396bf4176eb5978a2; sid_tt=3adf0f57a1c8cb7219b4b99e0858061d; sessionid=3adf0f57a1c8cb7219b4b99e0858061d; passport_csrf_token_default=2cc6edcbe17370d1139ed3790974a879; install_id=2168705416692471; ttreq=1$c5353c2b8a4931bab13bbd9c9f1c2471653f3817\"},{\"x-tt-dt\":\"AAA2G563DXCRQQBTH7IH65TFPPZOJSMC7FBJNT2RK7YP3QWFFFBSGSCBKC72V2OOGZ3TTA4OEHOI34QISOCOZOEJZC4ORLP7AMOVLACTSCPJ3SIRH6NSB7JNMPTDTUNEWAM64FHMKEGP7CHF2YZGVZQ\"},{\"activity_now_client\":\"1657953019620\"},{\"X-SS-REQ-TICKET\":\"1657953019717\"},{\"x-bd-client-key\":\"bd9779d620cbd92868fb41b7cc820324ccaabe16bef849e56c0efc64f6ed628384609486196db2930052c33d23c14e2f7ff21463cb983c77898fe100b34a3cd8\"},{\"x-bd-kmsv\":\"1\"},{\"passport-sdk-version\":\"20369\"},{\"X-Tt-Token\":\"003adf0f57a1c8cb7219b4b99e0858061d03a6f2875037249fc3ae91fc2fd7933b1ef3d4e67154f09fd0db00b2bee713270d67bbad0db9b0008ecc8c90fd73115fe1c90756e3cea59ec4231f99ebd38bf567ede90a8229a389c00a6490de30671ff94-1.0.1\"},{\"sdk-version\":\"2\"},{\"x-vc-bdturing-sdk-version\":\"2.2.1.cn\"}]"
	ls := make([]any, 0)
	json.ToObj(a, &ls)
	//m := make(map[string]string, 0)
	for _, v := range ls {
		ms := v.(map[string]any)
		for k, v2 := range ms {
			fmt.Println(k, ":", v2)
		}
	}
}

func Test3(t *testing.T) {
	a := "{\"x-tt-dt\":\"AAA2G563DXCRQQBTH7IH65TFPPZOJSMC7FBJNT2RK7YP3QWFFFBSGSCBKC72V2OOGZ3TTA4OEHOI34QISOCOZOEJZC4ORLP7AMOVLACTSCPJ3SIRH6NSB7JNMPTDTUNEWAM64FHMKEGP7CHF2YZGVZQ\",\"activity_now_client\":\"1657953364066\",\"X-SS-REQ-TICKET\":\"1657953363317\",\"x-bd-client-key\":\"bd9779d620cbd92868fb41b7cc820324ccaabe16bef849e56c0efc64f6ed628384609486196db2930052c33d23c14e2f7ff21463cb983c77898fe100b34a3cd8\",\"x-bd-kmsv\":\"1\",\"passport-sdk-version\":\"20369\",\"X-Tt-Token\":\"003adf0f57a1c8cb7219b4b99e0858061d03a6f2875037249fc3ae91fc2fd7933b1ef3d4e67154f09fd0db00b2bee713270d67bbad0db9b0008ecc8c90fd73115fe1c90756e3cea59ec4231f99ebd38bf567ede90a8229a389c00a6490de30671ff94-1.0.1\",\"sdk-version\":\"2\",\"x-vc-bdturing-sdk-version\":\"2.2.1.cn\",\"X-Khronos\":\"1657953365\",\"X-Gorgon\":\"0404a00700802717f38c584eaa288172bef90a3514f6f15aa1d7\",\"X-Ladon\":\"AenVHW4a63tSMtYXXlX/IHzFXKAI69G7ro13WEpber1+e3OH\",\"X-Argus\":\"lOqWdrlCJmswkhP5sC4LRCUzaI8HuqGQZdwVUbivrXX8OGAOSTsnuZNfRVl7nToI/ZQHT3VgcsKvKWzztJRw3bu2MTncWNQXngqm9LKfZXqhrva/jsnEfCM80Kk3qI+ujKDathHrz4svuyb4bUrckf6m/HNy2krn3QWaQV9qH2zQACdqkNSN6n+yRjYQxCYXs6iZJxitTA7vYadMyNlEmZ+gvetE3t7TtEzn4hzVv5wu6l1B1+yXbS18CQBSInk66cNjsLW4tUtmysMnGS4ji8qb\",\"Host\":\"aweme.snssdk.com\",\"Connection\":\"Keep-Alive\",\"User-Agent\":\"okhttp/3.10.0.1\"}"
	m := make(map[string]any)
	obj, e := json.ToObj(a, &m)
	log.Println(e)
	//m := make(map[string]string, 0)
	for k, v := range *obj {
		fmt.Println(k, ":", v)
	}
}
