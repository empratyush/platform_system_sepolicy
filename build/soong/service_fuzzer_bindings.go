// Copyright (C) 2022 The Android Open Source Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package selinux

var (
	ServiceFuzzerBindings = map[string][]string{
		"android.hardware.audio.core.IConfig/default":                             []string{},
		"android.hardware.audio.core.IModule/default":                             []string{},
		"android.hardware.audio.effect.IFactory/default":                          []string{},
		"android.hardware.authsecret.IAuthSecret/default":                         []string{},
		"android.hardware.automotive.evs.IEvsEnumerator/hw/0":                     []string{},
		"android.hardware.boot.IBootControl/default":                              []string{},
		"android.hardware.automotive.evs.IEvsEnumerator/hw/1":                     []string{},
		"android.hardware.automotive.remoteaccess.IRemoteAccess/default":          []string{},
		"android.hardware.automotive.vehicle.IVehicle/default":                    []string{},
		"android.hardware.automotive.audiocontrol.IAudioControl/default":          []string{},
		"android.hardware.biometrics.face.IFace/default":                          []string{},
		"android.hardware.biometrics.fingerprint.IFingerprint/default":            []string{},
		"android.hardware.biometrics.fingerprint.IFingerprint/virtual":            []string{},
		"android.hardware.bluetooth.audio.IBluetoothAudioProviderFactory/default": []string{},
		"android.hardware.broadcastradio.IBroadcastRadio/amfm":                    []string{},
		"android.hardware.broadcastradio.IBroadcastRadio/dab":                     []string{},
		"android.hardware.camera.provider.ICameraProvider/internal/0":             []string{},
		"android.hardware.confirmationui.IConfirmationUI/default":                 []string{},
		"android.hardware.contexthub.IContextHub/default":                         []string{},
		"android.hardware.drm.IDrmFactory/clearkey":                               []string{},
		"android.hardware.drm.ICryptoFactory/clearkey":                            []string{},
		"android.hardware.dumpstate.IDumpstateDevice/default":                     []string{},
		"android.hardware.gatekeeper.IGatekeeper/default":                         []string{},
		"android.hardware.gnss.IGnss/default":                                     []string{},
		"android.hardware.graphics.allocator.IAllocator/default":                  []string{},
		"android.hardware.graphics.composer3.IComposer/default":                   []string{},
		"android.hardware.health.storage.IStorage/default":                        []string{},
		"android.hardware.health.IHealth/default":                                 []string{},
		"android.hardware.identity.IIdentityCredentialStore/default":              []string{},
		"android.hardware.input.processor.IInputProcessor/default":                []string{},
		"android.hardware.ir.IConsumerIr/default":                                 []string{},
		"android.hardware.light.ILights/default":                                  []string{},
		"android.hardware.memtrack.IMemtrack/default":                             []string{},
		"android.hardware.net.nlinterceptor.IInterceptor/default":                 []string{},
		"android.hardware.nfc.INfc/default":                                       []string{},
		"android.hardware.oemlock.IOemLock/default":                               []string{},
		"android.hardware.power.IPower/default":                                   []string{},
		"android.hardware.power.stats.IPowerStats/default":                        []string{},
		"android.hardware.radio.config.IRadioConfig/default":                      []string{},
		"android.hardware.radio.data.IRadioData/slot1":                            []string{},
		"android.hardware.radio.data.IRadioData/slot2":                            []string{},
		"android.hardware.radio.data.IRadioData/slot3":                            []string{},
		"android.hardware.radio.ims.IRadioIms/slot1":                              []string{},
		"android.hardware.radio.ims.IRadioIms/slot2":                              []string{},
		"android.hardware.radio.ims.IRadioIms/slot3":                              []string{},
		"android.hardware.radio.ims.media.IImsMedia/default":                      []string{},
		"android.hardware.radio.messaging.IRadioMessaging/slot1":                  []string{},
		"android.hardware.radio.messaging.IRadioMessaging/slot2":                  []string{},
		"android.hardware.radio.messaging.IRadioMessaging/slot3":                  []string{},
		"android.hardware.radio.modem.IRadioModem/slot1":                          []string{},
		"android.hardware.radio.modem.IRadioModem/slot2":                          []string{},
		"android.hardware.radio.modem.IRadioModem/slot3":                          []string{},
		"android.hardware.radio.network.IRadioNetwork/slot1":                      []string{},
		"android.hardware.radio.network.IRadioNetwork/slot2":                      []string{},
		"android.hardware.radio.network.IRadioNetwork/slot3":                      []string{},
		"android.hardware.radio.sim.IRadioSim/slot1":                              []string{},
		"android.hardware.radio.sim.IRadioSim/slot2":                              []string{},
		"android.hardware.radio.sim.IRadioSim/slot3":                              []string{},
		"android.hardware.radio.voice.IRadioVoice/slot1":                          []string{},
		"android.hardware.radio.voice.IRadioVoice/slot2":                          []string{},
		"android.hardware.radio.voice.IRadioVoice/slot3":                          []string{},
		"android.hardware.rebootescrow.IRebootEscrow/default":                     []string{},
		"android.hardware.security.dice.IDiceDevice/default":                      []string{},
		"android.hardware.security.keymint.IKeyMintDevice/default":                []string{},
		"android.hardware.security.keymint.IRemotelyProvisionedComponent/default": []string{},
		"android.hardware.security.secureclock.ISecureClock/default":              []string{},
		"android.hardware.security.sharedsecret.ISharedSecret/default":            []string{},
		"android.hardware.sensors.ISensors/default":                               []string{},
		"android.hardware.soundtrigger3.ISoundTriggerHw/default":                  []string{},
		"android.hardware.thermal.IThermal/default":                               []string{},
		"android.hardware.tv.input.ITvInput/default":                              []string{},
		"android.hardware.tv.tuner.ITuner/default":                                []string{},
		"android.hardware.usb.IUsb/default":                                       []string{},
		"android.hardware.uwb.IUwb/default":                                       []string{},
		"android.hardware.vibrator.IVibrator/default":                             []string{},
		"android.hardware.vibrator.IVibratorManager/default":                      []string{},
		"android.hardware.weaver.IWeaver/default":                                 []string{},
		"android.hardware.wifi.hostapd.IHostapd/default":                          []string{},
		"android.hardware.wifi.supplicant.ISupplicant/default":                    []string{},
		"android.frameworks.stats.IStats/default":                                 []string{},
		"android.se.omapi.ISecureElementService/default":                          []string{},
		"android.system.keystore2.IKeystoreService/default":                       []string{},
		"android.system.net.netd.INetd/default":                                   []string{},
		"android.system.suspend.ISystemSuspend/default":                           []string{},
		"accessibility":      []string{},
		"account":            []string{},
		"activity":           []string{},
		"activity_task":      []string{},
		"adb":                []string{},
		"adservices_manager": []string{},
		"aidl_lazy_test_1":   []string{},
		"aidl_lazy_test_2":   []string{},
		"aidl_lazy_cb_test":  []string{},
		"alarm":              []string{},
		"android.hardware.automotive.evs.IEvsEnumerator/default":          []string{},
		"android.os.UpdateEngineService":                                  []string{},
		"android.os.UpdateEngineStableService":                            []string{},
		"android.frameworks.automotive.display.ICarDisplayProxy/default":  []string{},
		"android.security.apc":                                            []string{},
		"android.security.authorization":                                  []string{},
		"android.security.compat":                                         []string{},
		"android.security.dice.IDiceMaintenance":                          []string{},
		"android.security.dice.IDiceNode":                                 []string{},
		"android.security.identity":                                       []string{},
		"android.security.keystore":                                       []string{},
		"android.security.legacykeystore":                                 []string{},
		"android.security.maintenance":                                    []string{},
		"android.security.metrics":                                        []string{},
		"android.security.remoteprovisioning":                             []string{},
		"android.security.remoteprovisioning.IRemotelyProvisionedKeyPool": []string{},
		"android.service.gatekeeper.IGateKeeperService":                   []string{},
		"android.system.composd":                                          []string{},
		"android.system.virtualizationservice":                            []string{},
		"ambient_context":                                                 []string{},
		"app_binding":                                                     []string{},
		"app_hibernation":                                                 []string{},
		"app_integrity":                                                   []string{},
		"app_prediction":                                                  []string{},
		"app_search":                                                      []string{},
		"apexservice":                                                     []string{},
		"attestation_verification":                                        []string{},
		"blob_store":                                                      []string{},
		"gsiservice":                                                      []string{},
		"appops":                                                          []string{},
		"appwidget":                                                       []string{},
		"artd":                                                            []string{},
		"assetatlas":                                                      []string{},
		"attention":                                                       []string{},
		"audio":                                                           []string{},
		"auth":                                                            []string{},
		"autofill":                                                        []string{},
		"backup":                                                          []string{},
		"batteryproperties":                                               []string{},
		"batterystats":                                                    []string{},
		"battery":                                                         []string{},
		"binder_calls_stats":                                              []string{},
		"biometric":                                                       []string{},
		"bluetooth_manager":                                               []string{},
		"bluetooth":                                                       []string{},
		"broadcastradio":                                                  []string{},
		"bugreport":                                                       []string{},
		"cacheinfo":                                                       []string{},
		"carrier_config":                                                  []string{},
		"clipboard":                                                       []string{},
		"cloudsearch":                                                     []string{},
		"cloudsearch_service":                                             []string{},
		"com.android.net.IProxyService":                                   []string{},
		"companiondevice":                                                 []string{},
		"communal":                                                        []string{},
		"platform_compat":                                                 []string{},
		"platform_compat_native":                                          []string{},
		"connectivity":                                                    []string{},
		"connectivity_native":                                             []string{},
		"connmetrics":                                                     []string{},
		"consumer_ir":                                                     []string{},
		"content":                                                         []string{},
		"content_capture":                                                 []string{},
		"content_suggestions":                                             []string{},
		"contexthub":                                                      []string{},
		"country_detector":                                                []string{},
		"coverage":                                                        []string{},
		"cpuinfo":                                                         []string{},
		"credential":                                                      []string{},
		"crossprofileapps":                                                []string{},
		"dataloader_manager":                                              []string{},
		"dbinfo":                                                          []string{},
		"device_config":                                                   []string{},
		"device_policy":                                                   []string{},
		"device_identifiers":                                              []string{},
		"deviceidle":                                                      []string{},
		"device_lock":                                                     []string{},
		"device_state":                                                    []string{},
		"devicestoragemonitor":                                            []string{},
		"diskstats":                                                       []string{},
		"display":                                                         []string{},
		"dnsresolver":                                                     []string{},
		"domain_verification":                                             []string{},
		"color_display":                                                   []string{},
		"netd_listener":                                                   []string{},
		"network_watchlist":                                               []string{},
		"DockObserver":                                                    []string{},
		"dreams":                                                          []string{},
		"drm.drmManager":                                                  []string{},
		"dropbox":                                                         []string{},
		"dumpstate":                                                       []string{},
		"dynamic_system":                                                  []string{},
		"econtroller":                                                     []string{},
		"emergency_affordance":                                            []string{},
		"euicc_card_controller":                                           []string{},
		"external_vibrator_service":                                       []string{},
		"ethernet":                                                        []string{},
		"face":                                                            []string{},
		"file_integrity":                                                  []string{},
		"fingerprint":                                                     []string{},
		"font":                                                            []string{},
		"android.hardware.fingerprint.IFingerprintDaemon": []string{},
		"game":                         []string{},
		"gfxinfo":                      []string{},
		"gnss_time_update_service":     []string{},
		"graphicsstats":                []string{},
		"gpu":                          []string{},
		"hardware":                     []string{},
		"hardware_properties":          []string{},
		"hdmi_control":                 []string{},
		"healthconnect":                []string{},
		"ions":                         []string{},
		"idmap":                        []string{},
		"incident":                     []string{},
		"incidentcompanion":            []string{},
		"inputflinger":                 []string{},
		"input_method":                 []string{},
		"input":                        []string{},
		"installd":                     []string{},
		"iphonesubinfo_msim":           []string{},
		"iphonesubinfo2":               []string{},
		"iphonesubinfo":                []string{},
		"ims":                          []string{},
		"imms":                         []string{},
		"incremental":                  []string{},
		"ipsec":                        []string{},
		"ircsmessage":                  []string{},
		"iris":                         []string{},
		"isms_msim":                    []string{},
		"isms2":                        []string{},
		"isms":                         []string{},
		"isub":                         []string{},
		"jobscheduler":                 []string{},
		"launcherapps":                 []string{},
		"legacy_permission":            []string{},
		"lights":                       []string{},
		"locale":                       []string{},
		"location":                     []string{},
		"location_time_zone_manager":   []string{},
		"lock_settings":                []string{},
		"logcat":                       []string{},
		"logd":                         []string{},
		"looper_stats":                 []string{},
		"lpdump_service":               []string{},
		"mdns":                         []string{},
		"media.aaudio":                 []string{},
		"media.audio_flinger":          []string{},
		"media.audio_policy":           []string{},
		"media.camera":                 []string{},
		"media.camera.proxy":           []string{},
		"media.log":                    []string{},
		"media.player":                 []string{},
		"media.metrics":                []string{},
		"media.extractor":              []string{},
		"media.transcoding":            []string{},
		"media.resource_manager":       []string{},
		"media.resource_observer":      []string{},
		"media.sound_trigger_hw":       []string{},
		"media.drm":                    []string{},
		"media.tuner":                  []string{},
		"media_communication":          []string{},
		"media_metrics":                []string{},
		"media_projection":             []string{},
		"media_resource_monitor":       []string{},
		"media_router":                 []string{},
		"media_session":                []string{},
		"meminfo":                      []string{},
		"memtrack.proxy":               []string{},
		"midi":                         []string{},
		"mount":                        []string{},
		"music_recognition":            []string{},
		"nearby":                       []string{},
		"netd":                         []string{},
		"netpolicy":                    []string{},
		"netstats":                     []string{},
		"network_stack":                []string{},
		"network_management":           []string{},
		"network_score":                []string{},
		"network_time_update_service":  []string{},
		"nfc":                          []string{},
		"notification":                 []string{},
		"oem_lock":                     []string{},
		"otadexopt":                    []string{},
		"overlay":                      []string{},
		"pac_proxy":                    []string{},
		"package":                      []string{},
		"package_native":               []string{},
		"people":                       []string{},
		"performance_hint":             []string{},
		"permission":                   []string{},
		"permissionmgr":                []string{},
		"permission_checker":           []string{},
		"persistent_data_block":        []string{},
		"phone_msim":                   []string{},
		"phone1":                       []string{},
		"phone2":                       []string{},
		"phone":                        []string{},
		"pinner":                       []string{},
		"powerstats":                   []string{},
		"power":                        []string{},
		"print":                        []string{},
		"processinfo":                  []string{},
		"procstats":                    []string{},
		"profcollectd":                 []string{},
		"radio.phonesubinfo":           []string{},
		"radio.phone":                  []string{},
		"radio.sms":                    []string{},
		"rcs":                          []string{},
		"reboot_readiness":             []string{},
		"recovery":                     []string{},
		"resolver":                     []string{},
		"resources":                    []string{},
		"restrictions":                 []string{},
		"rkpd.registrar":               []string{},
		"rkpd.refresh":                 []string{},
		"role":                         []string{},
		"rollback":                     []string{},
		"rttmanager":                   []string{},
		"runtime":                      []string{},
		"safety_center":                []string{},
		"samplingprofiler":             []string{},
		"scheduling_policy":            []string{},
		"search":                       []string{},
		"search_ui":                    []string{},
		"secure_element":               []string{},
		"sec_key_att_app_id_provider":  []string{},
		"selection_toolbar":            []string{},
		"sensorservice":                []string{},
		"sensor_privacy":               []string{},
		"serial":                       []string{},
		"servicediscovery":             []string{},
		"manager":                      []string{},
		"settings":                     []string{},
		"shortcut":                     []string{},
		"simphonebook_msim":            []string{},
		"simphonebook2":                []string{},
		"simphonebook":                 []string{},
		"sip":                          []string{},
		"slice":                        []string{},
		"smartspace":                   []string{},
		"speech_recognition":           []string{},
		"stats":                        []string{},
		"statsbootstrap":               []string{},
		"statscompanion":               []string{},
		"statsmanager":                 []string{},
		"soundtrigger":                 []string{},
		"soundtrigger_middleware":      []string{},
		"statusbar":                    []string{},
		"storaged":                     []string{},
		"storaged_pri":                 []string{},
		"storagestats":                 []string{},
		"sdk_sandbox":                  []string{},
		"SurfaceFlinger":               []string{},
		"SurfaceFlingerAIDL":           []string{},
		"suspend_control":              []string{},
		"suspend_control_internal":     []string{},
		"system_config":                []string{},
		"system_server_dumper":         []string{},
		"system_update":                []string{},
		"tare":                         []string{},
		"task":                         []string{},
		"telecom":                      []string{},
		"telephony.registry":           []string{},
		"telephony_ims":                []string{},
		"testharness":                  []string{},
		"tethering":                    []string{},
		"textclassification":           []string{},
		"textservices":                 []string{},
		"texttospeech":                 []string{},
		"time_detector":                []string{},
		"time_zone_detector":           []string{},
		"thermalservice":               []string{},
		"tracing.proxy":                []string{},
		"translation":                  []string{},
		"transparency":                 []string{},
		"trust":                        []string{},
		"tv_interactive_app":           []string{},
		"tv_input":                     []string{},
		"tv_tuner_resource_mgr":        []string{},
		"uce":                          []string{},
		"uimode":                       []string{},
		"updatelock":                   []string{},
		"uri_grants":                   []string{},
		"usagestats":                   []string{},
		"usb":                          []string{},
		"user":                         []string{},
		"uwb":                          []string{},
		"vcn_management":               []string{},
		"vibrator":                     []string{},
		"vibrator_manager":             []string{},
		"virtualdevice":                []string{},
		"virtual_touchpad":             []string{},
		"voiceinteraction":             []string{},
		"vold":                         []string{},
		"vpn_management":               []string{},
		"vrmanager":                    []string{},
		"wallpaper":                    []string{},
		"wallpaper_effects_generation": []string{},
		"webviewupdate":                []string{},
		"wifip2p":                      []string{},
		"wifiscanner":                  []string{},
		"wifi":                         []string{},
		"wifinl80211":                  []string{},
		"wifiaware":                    []string{},
		"wifirtt":                      []string{},
		"window":                       []string{},
		"*":                            []string{},
	}
)
