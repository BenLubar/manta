//go:generate go run gen/game_event.go fixtures/source_1_legacy_game_events_list.pbmsg game_event_lookup.go

package manta

import (
	"github.com/dotabuff/manta/dota"
)

var gameEventNames = map[int32]string{
	0:   "server_spawn",
	1:   "server_pre_shutdown",
	2:   "server_shutdown",
	3:   "server_cvar",
	4:   "server_message",
	5:   "server_addban",
	6:   "server_removeban",
	7:   "player_connect",
	8:   "player_info",
	9:   "player_disconnect",
	10:  "player_activate",
	11:  "player_connect_full",
	12:  "player_say",
	13:  "player_full_update",
	14:  "team_info",
	15:  "team_score",
	16:  "teamplay_broadcast_audio",
	17:  "player_team",
	18:  "player_class",
	19:  "player_death",
	20:  "player_hurt",
	21:  "player_chat",
	22:  "player_score",
	23:  "player_spawn",
	24:  "player_shoot",
	25:  "player_use",
	26:  "player_changename",
	27:  "player_hintmessage",
	28:  "game_init",
	29:  "game_newmap",
	30:  "game_start",
	31:  "game_end",
	32:  "round_start",
	33:  "round_end",
	34:  "round_start_pre_entity",
	35:  "teamplay_round_start",
	36:  "hostname_changed",
	37:  "difficulty_changed",
	38:  "finale_start",
	39:  "game_message",
	40:  "break_breakable",
	41:  "break_prop",
	42:  "npc_spawned",
	43:  "npc_replaced",
	44:  "entity_killed",
	45:  "entity_hurt",
	46:  "bonus_updated",
	47:  "player_stats_updated",
	48:  "achievement_event",
	49:  "achievement_earned",
	50:  "achievement_write_failed",
	51:  "physgun_pickup",
	52:  "flare_ignite_npc",
	53:  "helicopter_grenade_punt_miss",
	54:  "user_data_downloaded",
	55:  "ragdoll_dissolved",
	56:  "gameinstructor_draw",
	57:  "gameinstructor_nodraw",
	58:  "map_transition",
	59:  "instructor_server_hint_create",
	60:  "instructor_server_hint_stop",
	61:  "chat_new_message",
	62:  "chat_members_changed",
	63:  "inventory_updated",
	64:  "cart_updated",
	65:  "store_pricesheet_updated",
	66:  "gc_connected",
	67:  "item_schema_initialized",
	68:  "drop_rate_modified",
	69:  "event_ticket_modified",
	70:  "modifier_event",
	71:  "dota_player_kill",
	72:  "dota_player_deny",
	73:  "dota_barracks_kill",
	74:  "dota_tower_kill",
	75:  "dota_effigy_kill",
	76:  "dota_roshan_kill",
	77:  "dota_courier_lost",
	78:  "dota_courier_respawned",
	79:  "dota_glyph_used",
	80:  "dota_super_creeps",
	81:  "dota_item_purchase",
	82:  "dota_item_gifted",
	83:  "dota_rune_pickup",
	84:  "dota_rune_spotted",
	85:  "dota_item_spotted",
	86:  "dota_no_battle_points",
	87:  "dota_chat_informational",
	88:  "dota_action_item",
	89:  "dota_chat_ban_notification",
	90:  "dota_chat_event",
	91:  "dota_chat_timed_reward",
	92:  "dota_pause_event",
	93:  "dota_chat_kill_streak",
	94:  "dota_chat_first_blood",
	95:  "dota_chat_assassin_announce",
	96:  "dota_chat_assassin_denied",
	97:  "dota_chat_assassin_success",
	98:  "dota_player_update_hero_selection",
	99:  "dota_player_update_selected_unit",
	100: "dota_player_update_query_unit",
	101: "dota_player_update_killcam_unit",
	102: "dota_player_take_tower_damage",
	103: "dota_hud_error_message",
	104: "dota_action_success",
	105: "dota_starting_position_changed",
	106: "dota_money_changed",
	107: "dota_enemy_money_changed",
	108: "dota_portrait_unit_stats_changed",
	109: "dota_portrait_unit_modifiers_changed",
	110: "dota_force_portrait_update",
	111: "dota_inventory_changed",
	112: "dota_item_picked_up",
	113: "dota_inventory_item_changed",
	114: "dota_ability_changed",
	115: "dota_portrait_ability_layout_changed",
	116: "dota_inventory_item_added",
	117: "dota_inventory_changed_query_unit",
	118: "dota_link_clicked",
	119: "dota_set_quick_buy",
	120: "dota_quick_buy_changed",
	121: "dota_player_shop_changed",
	122: "dota_player_show_killcam",
	123: "dota_player_show_minikillcam",
	124: "gc_user_session_created",
	125: "team_data_updated",
	126: "guild_data_updated",
	127: "guild_open_parties_updated",
	128: "fantasy_updated",
	129: "fantasy_league_changed",
	130: "fantasy_score_info_changed",
	131: "player_info_updated",
	132: "player_info_individual_updated",
	133: "game_rules_state_change",
	134: "match_history_updated",
	135: "match_details_updated",
	136: "live_games_updated",
	137: "recent_matches_updated",
	138: "news_updated",
	139: "persona_updated",
	140: "tournament_state_updated",
	141: "party_updated",
	142: "lobby_updated",
	143: "dashboard_caches_cleared",
	144: "last_hit",
	145: "player_completed_game",
	146: "player_reconnected",
	147: "nommed_tree",
	148: "dota_rune_activated_server",
	149: "dota_player_gained_level",
	150: "dota_player_learned_ability",
	151: "dota_player_used_ability",
	152: "dota_non_player_used_ability",
	153: "dota_player_begin_cast",
	154: "dota_non_player_begin_cast",
	155: "dota_ability_channel_finished",
	156: "dota_holdout_revive_complete",
	157: "dota_player_killed",
	158: "bindpanel_open",
	159: "bindpanel_close",
	160: "keybind_changed",
	161: "dota_item_drag_begin",
	162: "dota_item_drag_end",
	163: "dota_shop_item_drag_begin",
	164: "dota_shop_item_drag_end",
	165: "dota_item_purchased",
	166: "dota_item_combined",
	167: "dota_item_used",
	168: "dota_item_auto_purchase",
	169: "dota_unit_event",
	170: "dota_quest_started",
	171: "dota_quest_completed",
	172: "gameui_activated",
	173: "gameui_hidden",
	174: "player_fullyjoined",
	175: "dota_spectate_hero",
	176: "dota_match_done",
	177: "dota_match_done_client",
	178: "set_instructor_group_enabled",
	179: "joined_chat_channel",
	180: "left_chat_channel",
	181: "gc_chat_channel_list_updated",
	182: "today_messages_updated",
	183: "file_downloaded",
	184: "player_report_counts_updated",
	185: "scaleform_file_download_complete",
	186: "item_purchased",
	187: "gc_mismatched_version",
	190: "demo_stop",
	191: "map_shutdown",
	192: "dota_workshop_fileselected",
	193: "dota_workshop_filecanceled",
	194: "rich_presence_updated",
	195: "dota_hero_random",
	196: "dota_rd_chat_turn",
	197: "dota_favorite_heroes_updated",
	198: "profile_opened",
	199: "profile_closed",
	200: "item_preview_closed",
	201: "dashboard_switched_section",
	202: "dota_tournament_item_event",
	203: "dota_hero_swap",
	204: "dota_reset_suggested_items",
	205: "halloween_high_score_received",
	206: "halloween_phase_end",
	207: "halloween_high_score_request_failed",
	208: "dota_hud_skin_changed",
	209: "dota_inventory_player_got_item",
	210: "player_is_experienced",
	211: "player_is_notexperienced",
	212: "dota_tutorial_lesson_start",
	213: "dota_tutorial_task_advance",
	214: "dota_tutorial_shop_toggled",
	215: "map_location_updated",
	216: "richpresence_custom_updated",
	217: "game_end_visible",
	218: "antiaddiction_update",
	219: "highlight_hud_element",
	220: "hide_highlight_hud_element",
	221: "intro_video_finished",
	222: "matchmaking_status_visibility_changed",
	223: "practice_lobby_visibility_changed",
	224: "dota_courier_transfer_item",
	225: "full_ui_unlocked",
	227: "hero_selector_preview_set",
	228: "antiaddiction_toast",
	229: "hero_picker_shown",
	230: "hero_picker_hidden",
	231: "dota_local_quickbuy_changed",
	232: "show_center_message",
	233: "hud_flip_changed",
	234: "frosty_points_updated",
	235: "defeated",
	236: "reset_defeated",
	237: "booster_state_updated",
	238: "event_points_updated",
	239: "local_player_event_points",
	240: "custom_game_difficulty",
	241: "tree_cut",
	242: "ugc_details_arrived",
	243: "ugc_subscribed",
	244: "ugc_unsubscribed",
	245: "ugc_download_requested",
	246: "ugc_installed",
	247: "prizepool_received",
	248: "microtransaction_success",
	249: "dota_rubick_ability_steal",
	250: "compendium_event_actions_loaded",
	251: "compendium_selections_loaded",
	252: "compendium_set_selection_failed",
	253: "compendium_trophies_loaded",
	254: "community_cached_names_updated",
	255: "spec_item_pickup",
	256: "spec_aegis_reclaim_time",
	257: "account_trophies_changed",
	258: "account_all_hero_challenge_changed",
	259: "team_showcase_ui_update",
	260: "ingame_events_changed",
	261: "dota_match_signout",
	262: "dota_illusions_created",
	263: "dota_year_beast_killed",
	264: "dota_hero_undoselection",
	265: "dota_challenge_socache_updated",
	266: "party_invites_updated",
	267: "lobby_invites_updated",
	268: "custom_game_mode_list_updated",
	269: "custom_game_lobby_list_updated",
	270: "friend_lobby_list_updated",
	271: "dota_team_player_list_changed",
	272: "dota_player_details_changed",
	273: "player_profile_stats_updated",
	274: "custom_game_player_count_updated",
	275: "custom_game_friends_played_updated",
	276: "custom_games_friends_play_updated",
	277: "dota_player_update_assigned_hero",
	278: "dota_player_hero_selection_dirty",
	279: "dota_npc_goal_reached",
	280: "dota_player_selected_custom_team",
	281: "hltv_status",
	282: "hltv_cameraman",
	283: "hltv_rank_camera",
	284: "hltv_rank_entity",
	285: "hltv_fixed",
	286: "hltv_chase",
	287: "hltv_message",
	288: "hltv_title",
	289: "hltv_chat",
	290: "hltv_versioninfo",
	291: "dota_chase_hero",
	292: "dota_combatlog",
	293: "dota_game_state_change",
	294: "dota_player_pick_hero",
	295: "dota_team_kill_credit",
}

const (
	EGameEvent_ServerSpawn                        = 0
	EGameEvent_ServerPreShutdown                  = 1
	EGameEvent_ServerShutdown                     = 2
	EGameEvent_ServerCvar                         = 3
	EGameEvent_ServerMessage                      = 4
	EGameEvent_ServerAddban                       = 5
	EGameEvent_ServerRemoveban                    = 6
	EGameEvent_PlayerConnect                      = 7
	EGameEvent_PlayerInfo                         = 8
	EGameEvent_PlayerDisconnect                   = 9
	EGameEvent_PlayerActivate                     = 10
	EGameEvent_PlayerConnectFull                  = 11
	EGameEvent_PlayerSay                          = 12
	EGameEvent_PlayerFullUpdate                   = 13
	EGameEvent_TeamInfo                           = 14
	EGameEvent_TeamScore                          = 15
	EGameEvent_TeamplayBroadcastAudio             = 16
	EGameEvent_PlayerTeam                         = 17
	EGameEvent_PlayerClass                        = 18
	EGameEvent_PlayerDeath                        = 19
	EGameEvent_PlayerHurt                         = 20
	EGameEvent_PlayerChat                         = 21
	EGameEvent_PlayerScore                        = 22
	EGameEvent_PlayerSpawn                        = 23
	EGameEvent_PlayerShoot                        = 24
	EGameEvent_PlayerUse                          = 25
	EGameEvent_PlayerChangename                   = 26
	EGameEvent_PlayerHintmessage                  = 27
	EGameEvent_GameInit                           = 28
	EGameEvent_GameNewmap                         = 29
	EGameEvent_GameStart                          = 30
	EGameEvent_GameEnd                            = 31
	EGameEvent_RoundStart                         = 32
	EGameEvent_RoundEnd                           = 33
	EGameEvent_RoundStartPreEntity                = 34
	EGameEvent_TeamplayRoundStart                 = 35
	EGameEvent_HostnameChanged                    = 36
	EGameEvent_DifficultyChanged                  = 37
	EGameEvent_FinaleStart                        = 38
	EGameEvent_GameMessage                        = 39
	EGameEvent_BreakBreakable                     = 40
	EGameEvent_BreakProp                          = 41
	EGameEvent_NpcSpawned                         = 42
	EGameEvent_NpcReplaced                        = 43
	EGameEvent_EntityKilled                       = 44
	EGameEvent_EntityHurt                         = 45
	EGameEvent_BonusUpdated                       = 46
	EGameEvent_PlayerStatsUpdated                 = 47
	EGameEvent_AchievementEvent                   = 48
	EGameEvent_AchievementEarned                  = 49
	EGameEvent_AchievementWriteFailed             = 50
	EGameEvent_PhysgunPickup                      = 51
	EGameEvent_FlareIgniteNpc                     = 52
	EGameEvent_HelicopterGrenadePuntMiss          = 53
	EGameEvent_UserDataDownloaded                 = 54
	EGameEvent_RagdollDissolved                   = 55
	EGameEvent_GameinstructorDraw                 = 56
	EGameEvent_GameinstructorNodraw               = 57
	EGameEvent_MapTransition                      = 58
	EGameEvent_InstructorServerHintCreate         = 59
	EGameEvent_InstructorServerHintStop           = 60
	EGameEvent_ChatNewMessage                     = 61
	EGameEvent_ChatMembersChanged                 = 62
	EGameEvent_InventoryUpdated                   = 63
	EGameEvent_CartUpdated                        = 64
	EGameEvent_StorePricesheetUpdated             = 65
	EGameEvent_GcConnected                        = 66
	EGameEvent_ItemSchemaInitialized              = 67
	EGameEvent_DropRateModified                   = 68
	EGameEvent_EventTicketModified                = 69
	EGameEvent_ModifierEvent                      = 70
	EGameEvent_DotaPlayerKill                     = 71
	EGameEvent_DotaPlayerDeny                     = 72
	EGameEvent_DotaBarracksKill                   = 73
	EGameEvent_DotaTowerKill                      = 74
	EGameEvent_DotaEffigyKill                     = 75
	EGameEvent_DotaRoshanKill                     = 76
	EGameEvent_DotaCourierLost                    = 77
	EGameEvent_DotaCourierRespawned               = 78
	EGameEvent_DotaGlyphUsed                      = 79
	EGameEvent_DotaSuperCreeps                    = 80
	EGameEvent_DotaItemPurchase                   = 81
	EGameEvent_DotaItemGifted                     = 82
	EGameEvent_DotaRunePickup                     = 83
	EGameEvent_DotaRuneSpotted                    = 84
	EGameEvent_DotaItemSpotted                    = 85
	EGameEvent_DotaNoBattlePoints                 = 86
	EGameEvent_DotaChatInformational              = 87
	EGameEvent_DotaActionItem                     = 88
	EGameEvent_DotaChatBanNotification            = 89
	EGameEvent_DotaChatEvent                      = 90
	EGameEvent_DotaChatTimedReward                = 91
	EGameEvent_DotaPauseEvent                     = 92
	EGameEvent_DotaChatKillStreak                 = 93
	EGameEvent_DotaChatFirstBlood                 = 94
	EGameEvent_DotaChatAssassinAnnounce           = 95
	EGameEvent_DotaChatAssassinDenied             = 96
	EGameEvent_DotaChatAssassinSuccess            = 97
	EGameEvent_DotaPlayerUpdateHeroSelection      = 98
	EGameEvent_DotaPlayerUpdateSelectedUnit       = 99
	EGameEvent_DotaPlayerUpdateQueryUnit          = 100
	EGameEvent_DotaPlayerUpdateKillcamUnit        = 101
	EGameEvent_DotaPlayerTakeTowerDamage          = 102
	EGameEvent_DotaHudErrorMessage                = 103
	EGameEvent_DotaActionSuccess                  = 104
	EGameEvent_DotaStartingPositionChanged        = 105
	EGameEvent_DotaMoneyChanged                   = 106
	EGameEvent_DotaEnemyMoneyChanged              = 107
	EGameEvent_DotaPortraitUnitStatsChanged       = 108
	EGameEvent_DotaPortraitUnitModifiersChanged   = 109
	EGameEvent_DotaForcePortraitUpdate            = 110
	EGameEvent_DotaInventoryChanged               = 111
	EGameEvent_DotaItemPickedUp                   = 112
	EGameEvent_DotaInventoryItemChanged           = 113
	EGameEvent_DotaAbilityChanged                 = 114
	EGameEvent_DotaPortraitAbilityLayoutChanged   = 115
	EGameEvent_DotaInventoryItemAdded             = 116
	EGameEvent_DotaInventoryChangedQueryUnit      = 117
	EGameEvent_DotaLinkClicked                    = 118
	EGameEvent_DotaSetQuickBuy                    = 119
	EGameEvent_DotaQuickBuyChanged                = 120
	EGameEvent_DotaPlayerShopChanged              = 121
	EGameEvent_DotaPlayerShowKillcam              = 122
	EGameEvent_DotaPlayerShowMinikillcam          = 123
	EGameEvent_GcUserSessionCreated               = 124
	EGameEvent_TeamDataUpdated                    = 125
	EGameEvent_GuildDataUpdated                   = 126
	EGameEvent_GuildOpenPartiesUpdated            = 127
	EGameEvent_FantasyUpdated                     = 128
	EGameEvent_FantasyLeagueChanged               = 129
	EGameEvent_FantasyScoreInfoChanged            = 130
	EGameEvent_PlayerInfoUpdated                  = 131
	EGameEvent_PlayerInfoIndividualUpdated        = 132
	EGameEvent_GameRulesStateChange               = 133
	EGameEvent_MatchHistoryUpdated                = 134
	EGameEvent_MatchDetailsUpdated                = 135
	EGameEvent_LiveGamesUpdated                   = 136
	EGameEvent_RecentMatchesUpdated               = 137
	EGameEvent_NewsUpdated                        = 138
	EGameEvent_PersonaUpdated                     = 139
	EGameEvent_TournamentStateUpdated             = 140
	EGameEvent_PartyUpdated                       = 141
	EGameEvent_LobbyUpdated                       = 142
	EGameEvent_DashboardCachesCleared             = 143
	EGameEvent_LastHit                            = 144
	EGameEvent_PlayerCompletedGame                = 145
	EGameEvent_PlayerReconnected                  = 146
	EGameEvent_NommedTree                         = 147
	EGameEvent_DotaRuneActivatedServer            = 148
	EGameEvent_DotaPlayerGainedLevel              = 149
	EGameEvent_DotaPlayerLearnedAbility           = 150
	EGameEvent_DotaPlayerUsedAbility              = 151
	EGameEvent_DotaNonPlayerUsedAbility           = 152
	EGameEvent_DotaPlayerBeginCast                = 153
	EGameEvent_DotaNonPlayerBeginCast             = 154
	EGameEvent_DotaAbilityChannelFinished         = 155
	EGameEvent_DotaHoldoutReviveComplete          = 156
	EGameEvent_DotaPlayerKilled                   = 157
	EGameEvent_BindpanelOpen                      = 158
	EGameEvent_BindpanelClose                     = 159
	EGameEvent_KeybindChanged                     = 160
	EGameEvent_DotaItemDragBegin                  = 161
	EGameEvent_DotaItemDragEnd                    = 162
	EGameEvent_DotaShopItemDragBegin              = 163
	EGameEvent_DotaShopItemDragEnd                = 164
	EGameEvent_DotaItemPurchased                  = 165
	EGameEvent_DotaItemCombined                   = 166
	EGameEvent_DotaItemUsed                       = 167
	EGameEvent_DotaItemAutoPurchase               = 168
	EGameEvent_DotaUnitEvent                      = 169
	EGameEvent_DotaQuestStarted                   = 170
	EGameEvent_DotaQuestCompleted                 = 171
	EGameEvent_GameuiActivated                    = 172
	EGameEvent_GameuiHidden                       = 173
	EGameEvent_PlayerFullyjoined                  = 174
	EGameEvent_DotaSpectateHero                   = 175
	EGameEvent_DotaMatchDone                      = 176
	EGameEvent_DotaMatchDoneClient                = 177
	EGameEvent_SetInstructorGroupEnabled          = 178
	EGameEvent_JoinedChatChannel                  = 179
	EGameEvent_LeftChatChannel                    = 180
	EGameEvent_GcChatChannelListUpdated           = 181
	EGameEvent_TodayMessagesUpdated               = 182
	EGameEvent_FileDownloaded                     = 183
	EGameEvent_PlayerReportCountsUpdated          = 184
	EGameEvent_ScaleformFileDownloadComplete      = 185
	EGameEvent_ItemPurchased                      = 186
	EGameEvent_GcMismatchedVersion                = 187
	EGameEvent_DemoStop                           = 190
	EGameEvent_MapShutdown                        = 191
	EGameEvent_DotaWorkshopFileselected           = 192
	EGameEvent_DotaWorkshopFilecanceled           = 193
	EGameEvent_RichPresenceUpdated                = 194
	EGameEvent_DotaHeroRandom                     = 195
	EGameEvent_DotaRdChatTurn                     = 196
	EGameEvent_DotaFavoriteHeroesUpdated          = 197
	EGameEvent_ProfileOpened                      = 198
	EGameEvent_ProfileClosed                      = 199
	EGameEvent_ItemPreviewClosed                  = 200
	EGameEvent_DashboardSwitchedSection           = 201
	EGameEvent_DotaTournamentItemEvent            = 202
	EGameEvent_DotaHeroSwap                       = 203
	EGameEvent_DotaResetSuggestedItems            = 204
	EGameEvent_HalloweenHighScoreReceived         = 205
	EGameEvent_HalloweenPhaseEnd                  = 206
	EGameEvent_HalloweenHighScoreRequestFailed    = 207
	EGameEvent_DotaHudSkinChanged                 = 208
	EGameEvent_DotaInventoryPlayerGotItem         = 209
	EGameEvent_PlayerIsExperienced                = 210
	EGameEvent_PlayerIsNotexperienced             = 211
	EGameEvent_DotaTutorialLessonStart            = 212
	EGameEvent_DotaTutorialTaskAdvance            = 213
	EGameEvent_DotaTutorialShopToggled            = 214
	EGameEvent_MapLocationUpdated                 = 215
	EGameEvent_RichpresenceCustomUpdated          = 216
	EGameEvent_GameEndVisible                     = 217
	EGameEvent_AntiaddictionUpdate                = 218
	EGameEvent_HighlightHudElement                = 219
	EGameEvent_HideHighlightHudElement            = 220
	EGameEvent_IntroVideoFinished                 = 221
	EGameEvent_MatchmakingStatusVisibilityChanged = 222
	EGameEvent_PracticeLobbyVisibilityChanged     = 223
	EGameEvent_DotaCourierTransferItem            = 224
	EGameEvent_FullUiUnlocked                     = 225
	EGameEvent_HeroSelectorPreviewSet             = 227
	EGameEvent_AntiaddictionToast                 = 228
	EGameEvent_HeroPickerShown                    = 229
	EGameEvent_HeroPickerHidden                   = 230
	EGameEvent_DotaLocalQuickbuyChanged           = 231
	EGameEvent_ShowCenterMessage                  = 232
	EGameEvent_HudFlipChanged                     = 233
	EGameEvent_FrostyPointsUpdated                = 234
	EGameEvent_Defeated                           = 235
	EGameEvent_ResetDefeated                      = 236
	EGameEvent_BoosterStateUpdated                = 237
	EGameEvent_EventPointsUpdated                 = 238
	EGameEvent_LocalPlayerEventPoints             = 239
	EGameEvent_CustomGameDifficulty               = 240
	EGameEvent_TreeCut                            = 241
	EGameEvent_UgcDetailsArrived                  = 242
	EGameEvent_UgcSubscribed                      = 243
	EGameEvent_UgcUnsubscribed                    = 244
	EGameEvent_UgcDownloadRequested               = 245
	EGameEvent_UgcInstalled                       = 246
	EGameEvent_PrizepoolReceived                  = 247
	EGameEvent_MicrotransactionSuccess            = 248
	EGameEvent_DotaRubickAbilitySteal             = 249
	EGameEvent_CompendiumEventActionsLoaded       = 250
	EGameEvent_CompendiumSelectionsLoaded         = 251
	EGameEvent_CompendiumSetSelectionFailed       = 252
	EGameEvent_CompendiumTrophiesLoaded           = 253
	EGameEvent_CommunityCachedNamesUpdated        = 254
	EGameEvent_SpecItemPickup                     = 255
	EGameEvent_SpecAegisReclaimTime               = 256
	EGameEvent_AccountTrophiesChanged             = 257
	EGameEvent_AccountAllHeroChallengeChanged     = 258
	EGameEvent_TeamShowcaseUiUpdate               = 259
	EGameEvent_IngameEventsChanged                = 260
	EGameEvent_DotaMatchSignout                   = 261
	EGameEvent_DotaIllusionsCreated               = 262
	EGameEvent_DotaYearBeastKilled                = 263
	EGameEvent_DotaHeroUndoselection              = 264
	EGameEvent_DotaChallengeSocacheUpdated        = 265
	EGameEvent_PartyInvitesUpdated                = 266
	EGameEvent_LobbyInvitesUpdated                = 267
	EGameEvent_CustomGameModeListUpdated          = 268
	EGameEvent_CustomGameLobbyListUpdated         = 269
	EGameEvent_FriendLobbyListUpdated             = 270
	EGameEvent_DotaTeamPlayerListChanged          = 271
	EGameEvent_DotaPlayerDetailsChanged           = 272
	EGameEvent_PlayerProfileStatsUpdated          = 273
	EGameEvent_CustomGamePlayerCountUpdated       = 274
	EGameEvent_CustomGameFriendsPlayedUpdated     = 275
	EGameEvent_CustomGamesFriendsPlayUpdated      = 276
	EGameEvent_DotaPlayerUpdateAssignedHero       = 277
	EGameEvent_DotaPlayerHeroSelectionDirty       = 278
	EGameEvent_DotaNpcGoalReached                 = 279
	EGameEvent_DotaPlayerSelectedCustomTeam       = 280
	EGameEvent_HltvStatus                         = 281
	EGameEvent_HltvCameraman                      = 282
	EGameEvent_HltvRankCamera                     = 283
	EGameEvent_HltvRankEntity                     = 284
	EGameEvent_HltvFixed                          = 285
	EGameEvent_HltvChase                          = 286
	EGameEvent_HltvMessage                        = 287
	EGameEvent_HltvTitle                          = 288
	EGameEvent_HltvChat                           = 289
	EGameEvent_HltvVersioninfo                    = 290
	EGameEvent_DotaChaseHero                      = 291
	EGameEvent_DotaCombatlog                      = 292
	EGameEvent_DotaGameStateChange                = 293
	EGameEvent_DotaPlayerPickHero                 = 294
	EGameEvent_DotaTeamKillCredit                 = 295
)

type GameEventServerSpawn struct {
	Hostname   string `json:"hostname"`
	Address    string `json:"address"`
	Port       int32  `json:"port"`
	Game       string `json:"game"`
	Mapname    string `json:"mapname"`
	Addonname  string `json:"addonname"`
	Maxplayers int32  `json:"maxplayers"`
	Os         string `json:"os"`
	Dedicated  bool   `json:"dedicated"`
	Password   bool   `json:"password"`
}

type GameEventServerPreShutdown struct {
	Reason string `json:"reason"`
}

type GameEventServerShutdown struct {
	Reason string `json:"reason"`
}

type GameEventServerCvar struct {
	Cvarname  string `json:"cvarname"`
	Cvarvalue string `json:"cvarvalue"`
}

type GameEventServerMessage struct {
	Text string `json:"text"`
}

type GameEventServerAddban struct {
	Name      string `json:"name"`
	Userid    int32  `json:"userid"`
	Networkid string `json:"networkid"`
	Ip        string `json:"ip"`
	Duration  string `json:"duration"`
	By        string `json:"by"`
	Kicked    bool   `json:"kicked"`
}

type GameEventServerRemoveban struct {
	Networkid string `json:"networkid"`
	Ip        string `json:"ip"`
	By        string `json:"by"`
}

type GameEventPlayerConnect struct {
	Name      string `json:"name"`
	Index     int32  `json:"index"`
	Userid    int32  `json:"userid"`
	Networkid string `json:"networkid"`
	Address   string `json:"address"`
}

type GameEventPlayerInfo struct {
	Name      string `json:"name"`
	Index     int32  `json:"index"`
	Userid    int32  `json:"userid"`
	Networkid string `json:"networkid"`
	Bot       bool   `json:"bot"`
}

type GameEventPlayerDisconnect struct {
	Userid    int32  `json:"userid"`
	Reason    int32  `json:"reason"`
	Name      string `json:"name"`
	Networkid string `json:"networkid"`
}

type GameEventPlayerActivate struct {
	Userid int32 `json:"userid"`
}

type GameEventPlayerConnectFull struct {
	Userid int32 `json:"userid"`
	Index  int32 `json:"index"`
}

type GameEventPlayerSay struct {
	Userid int32  `json:"userid"`
	Text   string `json:"text"`
}

type GameEventPlayerFullUpdate struct {
	Userid int32 `json:"userid"`
	Count  int32 `json:"count"`
}

type GameEventTeamInfo struct {
	Teamid   int32  `json:"teamid"`
	Teamname string `json:"teamname"`
}

type GameEventTeamScore struct {
	Teamid int32 `json:"teamid"`
	Score  int32 `json:"score"`
}

type GameEventTeamplayBroadcastAudio struct {
	Team  int32  `json:"team"`
	Sound string `json:"sound"`
}

type GameEventPlayerTeam struct {
	Userid     int32 `json:"userid"`
	Team       int32 `json:"team"`
	Oldteam    int32 `json:"oldteam"`
	Disconnect bool  `json:"disconnect"`
	Autoteam   bool  `json:"autoteam"`
	Silent     bool  `json:"silent"`
}

type GameEventPlayerClass struct {
	Userid int32  `json:"userid"`
	Class  string `json:"class"`
}

type GameEventPlayerDeath struct {
	Userid   int32 `json:"userid"`
	Attacker int32 `json:"attacker"`
}

type GameEventPlayerHurt struct {
	Userid   int32 `json:"userid"`
	Attacker int32 `json:"attacker"`
	Health   int32 `json:"health"`
}

type GameEventPlayerChat struct {
	Teamonly bool   `json:"teamonly"`
	Userid   int32  `json:"userid"`
	Text     string `json:"text"`
}

type GameEventPlayerScore struct {
	Userid int32 `json:"userid"`
	Kills  int32 `json:"kills"`
	Deaths int32 `json:"deaths"`
	Score  int32 `json:"score"`
}

type GameEventPlayerSpawn struct {
	Userid int32 `json:"userid"`
}

type GameEventPlayerShoot struct {
	Userid int32 `json:"userid"`
	Weapon int32 `json:"weapon"`
	Mode   int32 `json:"mode"`
}

type GameEventPlayerUse struct {
	Userid int32 `json:"userid"`
	Entity int32 `json:"entity"`
}

type GameEventPlayerChangename struct {
	Userid  int32  `json:"userid"`
	Oldname string `json:"oldname"`
	Newname string `json:"newname"`
}

type GameEventPlayerHintmessage struct {
	Hintmessage string `json:"hintmessage"`
}

type GameEventGameInit struct {
}

type GameEventGameNewmap struct {
	Mapname string `json:"mapname"`
}

type GameEventGameStart struct {
	Roundslimit int32  `json:"roundslimit"`
	Timelimit   int32  `json:"timelimit"`
	Fraglimit   int32  `json:"fraglimit"`
	Objective   string `json:"objective"`
}

type GameEventGameEnd struct {
	Winner int32 `json:"winner"`
}

type GameEventRoundStart struct {
	Timelimit int32  `json:"timelimit"`
	Fraglimit int32  `json:"fraglimit"`
	Objective string `json:"objective"`
}

type GameEventRoundEnd struct {
	Winner  int32  `json:"winner"`
	Reason  int32  `json:"reason"`
	Message string `json:"message"`
}

type GameEventRoundStartPreEntity struct {
}

type GameEventTeamplayRoundStart struct {
	FullReset bool `json:"full_reset"`
}

type GameEventHostnameChanged struct {
	Hostname string `json:"hostname"`
}

type GameEventDifficultyChanged struct {
	NewDifficulty int32  `json:"newDifficulty"`
	OldDifficulty int32  `json:"oldDifficulty"`
	StrDifficulty string `json:"strDifficulty"`
}

type GameEventFinaleStart struct {
	Rushes int32 `json:"rushes"`
}

type GameEventGameMessage struct {
	Target int32  `json:"target"`
	Text   string `json:"text"`
}

type GameEventBreakBreakable struct {
	Entindex int32 `json:"entindex"`
	Userid   int32 `json:"userid"`
	Material int32 `json:"material"`
}

type GameEventBreakProp struct {
	Entindex int32 `json:"entindex"`
	Userid   int32 `json:"userid"`
}

type GameEventNpcSpawned struct {
	Entindex int32 `json:"entindex"`
}

type GameEventNpcReplaced struct {
	OldEntindex int32 `json:"old_entindex"`
	NewEntindex int32 `json:"new_entindex"`
}

type GameEventEntityKilled struct {
	EntindexKilled    int32 `json:"entindex_killed"`
	EntindexAttacker  int32 `json:"entindex_attacker"`
	EntindexInflictor int32 `json:"entindex_inflictor"`
	Damagebits        int32 `json:"damagebits"`
}

type GameEventEntityHurt struct {
	EntindexKilled    int32 `json:"entindex_killed"`
	EntindexAttacker  int32 `json:"entindex_attacker"`
	EntindexInflictor int32 `json:"entindex_inflictor"`
	Damagebits        int32 `json:"damagebits"`
}

type GameEventBonusUpdated struct {
	Numadvanced int32 `json:"numadvanced"`
	Numbronze   int32 `json:"numbronze"`
	Numsilver   int32 `json:"numsilver"`
	Numgold     int32 `json:"numgold"`
}

type GameEventPlayerStatsUpdated struct {
	Forceupload bool `json:"forceupload"`
}

type GameEventAchievementEvent struct {
	AchievementName string `json:"achievement_name"`
	CurVal          int32  `json:"cur_val"`
	MaxVal          int32  `json:"max_val"`
}

type GameEventAchievementEarned struct {
	Player      int32 `json:"player"`
	Achievement int32 `json:"achievement"`
}

type GameEventAchievementWriteFailed struct {
}

type GameEventPhysgunPickup struct {
	Entindex int32 `json:"entindex"`
}

type GameEventFlareIgniteNpc struct {
	Entindex int32 `json:"entindex"`
}

type GameEventHelicopterGrenadePuntMiss struct {
}

type GameEventUserDataDownloaded struct {
}

type GameEventRagdollDissolved struct {
	Entindex int32 `json:"entindex"`
}

type GameEventGameinstructorDraw struct {
}

type GameEventGameinstructorNodraw struct {
}

type GameEventMapTransition struct {
}

type GameEventInstructorServerHintCreate struct {
	HintName              string  `json:"hint_name"`
	HintReplaceKey        string  `json:"hint_replace_key"`
	HintTarget            int32   `json:"hint_target"`
	HintActivatorUserid   int32   `json:"hint_activator_userid"`
	HintTimeout           int32   `json:"hint_timeout"`
	HintIconOnscreen      string  `json:"hint_icon_onscreen"`
	HintIconOffscreen     string  `json:"hint_icon_offscreen"`
	HintCaption           string  `json:"hint_caption"`
	HintActivatorCaption  string  `json:"hint_activator_caption"`
	HintColor             string  `json:"hint_color"`
	HintIconOffset        float32 `json:"hint_icon_offset"`
	HintRange             float32 `json:"hint_range"`
	HintFlags             int32   `json:"hint_flags"`
	HintBinding           string  `json:"hint_binding"`
	HintAllowNodrawTarget bool    `json:"hint_allow_nodraw_target"`
	HintNooffscreen       bool    `json:"hint_nooffscreen"`
	HintForcecaption      bool    `json:"hint_forcecaption"`
	HintLocalPlayerOnly   bool    `json:"hint_local_player_only"`
}

type GameEventInstructorServerHintStop struct {
	HintName string `json:"hint_name"`
}

type GameEventChatNewMessage struct {
	Channel int32 `json:"channel"`
}

type GameEventChatMembersChanged struct {
	Channel int32 `json:"channel"`
}

type GameEventInventoryUpdated struct {
	Itemdef int32 `json:"itemdef"`
	Itemid  int32 `json:"itemid"`
}

type GameEventCartUpdated struct {
}

type GameEventStorePricesheetUpdated struct {
}

type GameEventGcConnected struct {
}

type GameEventItemSchemaInitialized struct {
}

type GameEventDropRateModified struct {
}

type GameEventEventTicketModified struct {
}

type GameEventModifierEvent struct {
	Eventname string `json:"eventname"`
	Caster    int32  `json:"caster"`
	Ability   int32  `json:"ability"`
}

type GameEventDotaPlayerKill struct {
	VictimUserid  int32 `json:"victim_userid"`
	Killer1Userid int32 `json:"killer1_userid"`
	Killer2Userid int32 `json:"killer2_userid"`
	Killer3Userid int32 `json:"killer3_userid"`
	Killer4Userid int32 `json:"killer4_userid"`
	Killer5Userid int32 `json:"killer5_userid"`
	Bounty        int32 `json:"bounty"`
	Neutral       int32 `json:"neutral"`
	Greevil       int32 `json:"greevil"`
}

type GameEventDotaPlayerDeny struct {
	KillerUserid int32 `json:"killer_userid"`
	VictimUserid int32 `json:"victim_userid"`
}

type GameEventDotaBarracksKill struct {
	BarracksId int32 `json:"barracks_id"`
}

type GameEventDotaTowerKill struct {
	KillerUserid int32 `json:"killer_userid"`
	Teamnumber   int32 `json:"teamnumber"`
	Gold         int32 `json:"gold"`
}

type GameEventDotaEffigyKill struct {
	OwnerUserid int32 `json:"owner_userid"`
}

type GameEventDotaRoshanKill struct {
	Teamnumber int32 `json:"teamnumber"`
	Gold       int32 `json:"gold"`
}

type GameEventDotaCourierLost struct {
	Teamnumber int32 `json:"teamnumber"`
}

type GameEventDotaCourierRespawned struct {
	Teamnumber int32 `json:"teamnumber"`
}

type GameEventDotaGlyphUsed struct {
	Teamnumber int32 `json:"teamnumber"`
}

type GameEventDotaSuperCreeps struct {
	Teamnumber int32 `json:"teamnumber"`
}

type GameEventDotaItemPurchase struct {
	Userid int32 `json:"userid"`
	Itemid int32 `json:"itemid"`
}

type GameEventDotaItemGifted struct {
	Userid   int32 `json:"userid"`
	Itemid   int32 `json:"itemid"`
	Sourceid int32 `json:"sourceid"`
}

type GameEventDotaRunePickup struct {
	Userid int32 `json:"userid"`
	Type   int32 `json:"type"`
	Rune   int32 `json:"rune"`
}

type GameEventDotaRuneSpotted struct {
	Userid int32 `json:"userid"`
	Rune   int32 `json:"rune"`
}

type GameEventDotaItemSpotted struct {
	Userid int32 `json:"userid"`
	Itemid int32 `json:"itemid"`
}

type GameEventDotaNoBattlePoints struct {
	Userid int32 `json:"userid"`
	Reason int32 `json:"reason"`
}

type GameEventDotaChatInformational struct {
	Userid int32 `json:"userid"`
	Type   int32 `json:"type"`
}

type GameEventDotaActionItem struct {
	Reason  int32 `json:"reason"`
	Itemdef int32 `json:"itemdef"`
	Message int32 `json:"message"`
}

type GameEventDotaChatBanNotification struct {
	Userid int32 `json:"userid"`
}

type GameEventDotaChatEvent struct {
	Userid  int32 `json:"userid"`
	Gold    int32 `json:"gold"`
	Message int32 `json:"message"`
}

type GameEventDotaChatTimedReward struct {
	Userid  int32 `json:"userid"`
	Itmedef int32 `json:"itmedef"`
	Message int32 `json:"message"`
}

type GameEventDotaPauseEvent struct {
	Userid  int32 `json:"userid"`
	Value   int32 `json:"value"`
	Message int32 `json:"message"`
}

type GameEventDotaChatKillStreak struct {
	Gold            int32 `json:"gold"`
	KillerId        int32 `json:"killer_id"`
	KillerStreak    int32 `json:"killer_streak"`
	KillerMultikill int32 `json:"killer_multikill"`
	VictimId        int32 `json:"victim_id"`
	VictimStreak    int32 `json:"victim_streak"`
}

type GameEventDotaChatFirstBlood struct {
	Gold     int32 `json:"gold"`
	KillerId int32 `json:"killer_id"`
	VictimId int32 `json:"victim_id"`
}

type GameEventDotaChatAssassinAnnounce struct {
	AssassinId int32 `json:"assassin_id"`
	TargetId   int32 `json:"target_id"`
	Message    int32 `json:"message"`
}

type GameEventDotaChatAssassinDenied struct {
	AssassinId int32 `json:"assassin_id"`
	TargetId   int32 `json:"target_id"`
	Message    int32 `json:"message"`
}

type GameEventDotaChatAssassinSuccess struct {
	AssassinId int32 `json:"assassin_id"`
	TargetId   int32 `json:"target_id"`
	Message    int32 `json:"message"`
}

type GameEventDotaPlayerUpdateHeroSelection struct {
	Tabcycle bool `json:"tabcycle"`
}

type GameEventDotaPlayerUpdateSelectedUnit struct {
}

type GameEventDotaPlayerUpdateQueryUnit struct {
}

type GameEventDotaPlayerUpdateKillcamUnit struct {
}

type GameEventDotaPlayerTakeTowerDamage struct {
	PlayerID int32 `json:"PlayerID"`
	Damage   int32 `json:"damage"`
}

type GameEventDotaHudErrorMessage struct {
	Reason  int32  `json:"reason"`
	Message string `json:"message"`
}

type GameEventDotaActionSuccess struct {
}

type GameEventDotaStartingPositionChanged struct {
}

type GameEventDotaMoneyChanged struct {
}

type GameEventDotaEnemyMoneyChanged struct {
}

type GameEventDotaPortraitUnitStatsChanged struct {
}

type GameEventDotaPortraitUnitModifiersChanged struct {
}

type GameEventDotaForcePortraitUpdate struct {
}

type GameEventDotaInventoryChanged struct {
}

type GameEventDotaItemPickedUp struct {
	Itemname        string `json:"itemname"`
	PlayerID        int32  `json:"PlayerID"`
	ItemEntityIndex int32  `json:"ItemEntityIndex"`
	HeroEntityIndex int32  `json:"HeroEntityIndex"`
}

type GameEventDotaInventoryItemChanged struct {
	EntityIndex int32 `json:"entityIndex"`
}

type GameEventDotaAbilityChanged struct {
}

type GameEventDotaPortraitAbilityLayoutChanged struct {
}

type GameEventDotaInventoryItemAdded struct {
	Itemname string `json:"itemname"`
}

type GameEventDotaInventoryChangedQueryUnit struct {
}

type GameEventDotaLinkClicked struct {
	Link    string `json:"link"`
	Nav     bool   `json:"nav"`
	NavBack bool   `json:"nav_back"`
	Recipe  int32  `json:"recipe"`
	Shop    int32  `json:"shop"`
}

type GameEventDotaSetQuickBuy struct {
	Item   string `json:"item"`
	Recipe int32  `json:"recipe"`
	Toggle bool   `json:"toggle"`
}

type GameEventDotaQuickBuyChanged struct {
	Item   string `json:"item"`
	Recipe int32  `json:"recipe"`
}

type GameEventDotaPlayerShopChanged struct {
	Prevshopmask int32 `json:"prevshopmask"`
	Shopmask     int32 `json:"shopmask"`
}

type GameEventDotaPlayerShowKillcam struct {
	Nodes  int32 `json:"nodes"`
	Player int32 `json:"player"`
}

type GameEventDotaPlayerShowMinikillcam struct {
	Nodes  int32 `json:"nodes"`
	Player int32 `json:"player"`
}

type GameEventGcUserSessionCreated struct {
}

type GameEventTeamDataUpdated struct {
}

type GameEventGuildDataUpdated struct {
}

type GameEventGuildOpenPartiesUpdated struct {
}

type GameEventFantasyUpdated struct {
}

type GameEventFantasyLeagueChanged struct {
}

type GameEventFantasyScoreInfoChanged struct {
}

type GameEventPlayerInfoUpdated struct {
}

type GameEventPlayerInfoIndividualUpdated struct {
	AccountId int32 `json:"account_id"`
}

type GameEventGameRulesStateChange struct {
}

type GameEventMatchHistoryUpdated struct {
	SteamID uint64 `json:"SteamID"`
}

type GameEventMatchDetailsUpdated struct {
	MatchID uint64 `json:"matchID"`
	Result  int32  `json:"result"`
}

type GameEventLiveGamesUpdated struct {
}

type GameEventRecentMatchesUpdated struct {
	Page int32 `json:"Page"`
}

type GameEventNewsUpdated struct {
}

type GameEventPersonaUpdated struct {
	SteamID uint64 `json:"SteamID"`
}

type GameEventTournamentStateUpdated struct {
}

type GameEventPartyUpdated struct {
}

type GameEventLobbyUpdated struct {
}

type GameEventDashboardCachesCleared struct {
}

type GameEventLastHit struct {
	PlayerID   int32 `json:"PlayerID"`
	EntKilled  int32 `json:"EntKilled"`
	FirstBlood bool  `json:"FirstBlood"`
	HeroKill   bool  `json:"HeroKill"`
	TowerKill  bool  `json:"TowerKill"`
}

type GameEventPlayerCompletedGame struct {
	PlayerID int32 `json:"PlayerID"`
	Winner   int32 `json:"Winner"`
}

type GameEventPlayerReconnected struct {
	PlayerID int32 `json:"PlayerID"`
}

type GameEventNommedTree struct {
	PlayerID int32 `json:"PlayerID"`
}

type GameEventDotaRuneActivatedServer struct {
	PlayerID int32 `json:"PlayerID"`
	Rune     int32 `json:"rune"`
}

type GameEventDotaPlayerGainedLevel struct {
	PlayerID int32 `json:"PlayerID"`
	Level    int32 `json:"level"`
}

type GameEventDotaPlayerLearnedAbility struct {
	PlayerID    int32  `json:"PlayerID"`
	Abilityname string `json:"abilityname"`
}

type GameEventDotaPlayerUsedAbility struct {
	PlayerID    int32  `json:"PlayerID"`
	Abilityname string `json:"abilityname"`
}

type GameEventDotaNonPlayerUsedAbility struct {
	Abilityname string `json:"abilityname"`
}

type GameEventDotaPlayerBeginCast struct {
	PlayerID    int32  `json:"PlayerID"`
	Abilityname string `json:"abilityname"`
}

type GameEventDotaNonPlayerBeginCast struct {
	Abilityname string `json:"abilityname"`
}

type GameEventDotaAbilityChannelFinished struct {
	Abilityname string `json:"abilityname"`
	Interrupted bool   `json:"interrupted"`
}

type GameEventDotaHoldoutReviveComplete struct {
	Caster int32 `json:"caster"`
	Target int32 `json:"target"`
}

type GameEventDotaPlayerKilled struct {
	PlayerID  int32 `json:"PlayerID"`
	HeroKill  bool  `json:"HeroKill"`
	TowerKill bool  `json:"TowerKill"`
}

type GameEventBindpanelOpen struct {
}

type GameEventBindpanelClose struct {
}

type GameEventKeybindChanged struct {
}

type GameEventDotaItemDragBegin struct {
}

type GameEventDotaItemDragEnd struct {
}

type GameEventDotaShopItemDragBegin struct {
}

type GameEventDotaShopItemDragEnd struct {
}

type GameEventDotaItemPurchased struct {
	PlayerID int32  `json:"PlayerID"`
	Itemname string `json:"itemname"`
	Itemcost int32  `json:"itemcost"`
}

type GameEventDotaItemCombined struct {
	PlayerID int32  `json:"PlayerID"`
	Itemname string `json:"itemname"`
	Itemcost int32  `json:"itemcost"`
}

type GameEventDotaItemUsed struct {
	PlayerID int32  `json:"PlayerID"`
	Itemname string `json:"itemname"`
}

type GameEventDotaItemAutoPurchase struct {
	ItemId int32 `json:"item_id"`
}

type GameEventDotaUnitEvent struct {
	Victim       int32 `json:"victim"`
	Attacker     int32 `json:"attacker"`
	Basepriority int32 `json:"basepriority"`
	Priority     int32 `json:"priority"`
	Eventtype    int32 `json:"eventtype"`
}

type GameEventDotaQuestStarted struct {
	QuestIndex int32 `json:"questIndex"`
}

type GameEventDotaQuestCompleted struct {
	QuestIndex int32 `json:"questIndex"`
}

type GameEventGameuiActivated struct {
}

type GameEventGameuiHidden struct {
}

type GameEventPlayerFullyjoined struct {
	Userid int32  `json:"userid"`
	Name   string `json:"name"`
}

type GameEventDotaSpectateHero struct {
	Entindex int32 `json:"entindex"`
}

type GameEventDotaMatchDone struct {
	Winningteam int32 `json:"winningteam"`
}

type GameEventDotaMatchDoneClient struct {
}

type GameEventSetInstructorGroupEnabled struct {
	Group   string `json:"group"`
	Enabled int32  `json:"enabled"`
}

type GameEventJoinedChatChannel struct {
	ChannelName string `json:"channelName"`
}

type GameEventLeftChatChannel struct {
	ChannelName string `json:"channelName"`
}

type GameEventGcChatChannelListUpdated struct {
}

type GameEventTodayMessagesUpdated struct {
	NumMessages int32 `json:"num_messages"`
}

type GameEventFileDownloaded struct {
	Success       bool   `json:"success"`
	LocalFilename string `json:"local_filename"`
	RemoteUrl     string `json:"remote_url"`
}

type GameEventPlayerReportCountsUpdated struct {
	PositiveRemaining int32 `json:"positive_remaining"`
	NegativeRemaining int32 `json:"negative_remaining"`
	PositiveTotal     int32 `json:"positive_total"`
	NegativeTotal     int32 `json:"negative_total"`
}

type GameEventScaleformFileDownloadComplete struct {
	Success       bool   `json:"success"`
	LocalFilename string `json:"local_filename"`
	RemoteUrl     string `json:"remote_url"`
}

type GameEventItemPurchased struct {
	Itemid int32 `json:"itemid"`
}

type GameEventGcMismatchedVersion struct {
}

type GameEventDemoStop struct {
}

type GameEventMapShutdown struct {
}

type GameEventDotaWorkshopFileselected struct {
	Filename string `json:"filename"`
}

type GameEventDotaWorkshopFilecanceled struct {
}

type GameEventRichPresenceUpdated struct {
}

type GameEventDotaHeroRandom struct {
	Userid int32 `json:"userid"`
	Heroid int32 `json:"heroid"`
}

type GameEventDotaRdChatTurn struct {
	Userid int32 `json:"userid"`
}

type GameEventDotaFavoriteHeroesUpdated struct {
}

type GameEventProfileOpened struct {
}

type GameEventProfileClosed struct {
}

type GameEventItemPreviewClosed struct {
}

type GameEventDashboardSwitchedSection struct {
	Section int32 `json:"section"`
}

type GameEventDotaTournamentItemEvent struct {
	WinnerCount int32 `json:"winner_count"`
	EventType   int32 `json:"event_type"`
}

type GameEventDotaHeroSwap struct {
	Playerid1 int32 `json:"playerid1"`
	Playerid2 int32 `json:"playerid2"`
}

type GameEventDotaResetSuggestedItems struct {
}

type GameEventHalloweenHighScoreReceived struct {
	Round int32 `json:"round"`
}

type GameEventHalloweenPhaseEnd struct {
	Phase int32 `json:"phase"`
	Team  int32 `json:"team"`
}

type GameEventHalloweenHighScoreRequestFailed struct {
	Round int32 `json:"round"`
}

type GameEventDotaHudSkinChanged struct {
	Skin  string `json:"skin"`
	Style int32  `json:"style"`
}

type GameEventDotaInventoryPlayerGotItem struct {
	Itemname string `json:"itemname"`
}

type GameEventPlayerIsExperienced struct {
}

type GameEventPlayerIsNotexperienced struct {
}

type GameEventDotaTutorialLessonStart struct {
}

type GameEventDotaTutorialTaskAdvance struct {
}

type GameEventDotaTutorialShopToggled struct {
	ShopOpened bool `json:"shop_opened"`
}

type GameEventMapLocationUpdated struct {
}

type GameEventRichpresenceCustomUpdated struct {
}

type GameEventGameEndVisible struct {
}

type GameEventAntiaddictionUpdate struct {
}

type GameEventHighlightHudElement struct {
	Elementname string  `json:"elementname"`
	Duration    float32 `json:"duration"`
}

type GameEventHideHighlightHudElement struct {
}

type GameEventIntroVideoFinished struct {
}

type GameEventMatchmakingStatusVisibilityChanged struct {
}

type GameEventPracticeLobbyVisibilityChanged struct {
}

type GameEventDotaCourierTransferItem struct {
}

type GameEventFullUiUnlocked struct {
}

type GameEventHeroSelectorPreviewSet struct {
	Setindex int32 `json:"setindex"`
}

type GameEventAntiaddictionToast struct {
	Message  string  `json:"message"`
	Duration float32 `json:"duration"`
}

type GameEventHeroPickerShown struct {
}

type GameEventHeroPickerHidden struct {
}

type GameEventDotaLocalQuickbuyChanged struct {
}

type GameEventShowCenterMessage struct {
	Message           string  `json:"message"`
	Duration          float32 `json:"duration"`
	ClearMessageQueue bool    `json:"clear_message_queue"`
}

type GameEventHudFlipChanged struct {
	Flipped bool `json:"flipped"`
}

type GameEventFrostyPointsUpdated struct {
}

type GameEventDefeated struct {
	Entindex int32 `json:"entindex"`
}

type GameEventResetDefeated struct {
}

type GameEventBoosterStateUpdated struct {
}

type GameEventEventPointsUpdated struct {
	EventId       int32 `json:"event_id"`
	Points        int32 `json:"points"`
	PremiumPoints int32 `json:"premium_points"`
	Owned         bool  `json:"owned"`
}

type GameEventLocalPlayerEventPoints struct {
	Points         int32 `json:"points"`
	ConversionRate int32 `json:"conversion_rate"`
}

type GameEventCustomGameDifficulty struct {
	Difficulty int32 `json:"difficulty"`
}

type GameEventTreeCut struct {
	TreeX float32 `json:"tree_x"`
	TreeY float32 `json:"tree_y"`
}

type GameEventUgcDetailsArrived struct {
	PublishedFileId uint64 `json:"published_file_id"`
}

type GameEventUgcSubscribed struct {
	PublishedFileId uint64 `json:"published_file_id"`
}

type GameEventUgcUnsubscribed struct {
	PublishedFileId uint64 `json:"published_file_id"`
}

type GameEventUgcDownloadRequested struct {
	PublishedFileId uint64 `json:"published_file_id"`
}

type GameEventUgcInstalled struct {
	PublishedFileId uint64 `json:"published_file_id"`
}

type GameEventPrizepoolReceived struct {
	Success   bool   `json:"success"`
	Prizepool uint64 `json:"prizepool"`
	Leagueid  uint64 `json:"leagueid"`
}

type GameEventMicrotransactionSuccess struct {
	Txnid uint64 `json:"txnid"`
}

type GameEventDotaRubickAbilitySteal struct {
	AbilityIndex int32 `json:"abilityIndex"`
	AbilityLevel int32 `json:"abilityLevel"`
}

type GameEventCompendiumEventActionsLoaded struct {
	AccountId      uint64 `json:"account_id"`
	LeagueId       uint64 `json:"league_id"`
	LocalTest      bool   `json:"local_test"`
	OriginalPoints uint64 `json:"original_points"`
}

type GameEventCompendiumSelectionsLoaded struct {
	AccountId uint64 `json:"account_id"`
	LeagueId  uint64 `json:"league_id"`
	LocalTest bool   `json:"local_test"`
}

type GameEventCompendiumSetSelectionFailed struct {
	AccountId uint64 `json:"account_id"`
	LeagueId  uint64 `json:"league_id"`
	LocalTest bool   `json:"local_test"`
}

type GameEventCompendiumTrophiesLoaded struct {
	AccountId uint64 `json:"account_id"`
	LeagueId  uint64 `json:"league_id"`
	LocalTest bool   `json:"local_test"`
}

type GameEventCommunityCachedNamesUpdated struct {
}

type GameEventSpecItemPickup struct {
	PlayerId int32  `json:"player_id"`
	ItemName string `json:"item_name"`
	Purchase bool   `json:"purchase"`
}

type GameEventSpecAegisReclaimTime struct {
	ReclaimTime float32 `json:"reclaim_time"`
}

type GameEventAccountTrophiesChanged struct {
	AccountId uint64 `json:"account_id"`
}

type GameEventAccountAllHeroChallengeChanged struct {
	AccountId uint64 `json:"account_id"`
}

type GameEventTeamShowcaseUiUpdate struct {
	Show            bool   `json:"show"`
	AccountId       uint64 `json:"account_id"`
	HeroEntindex    int32  `json:"hero_entindex"`
	DisplayUiOnLeft bool   `json:"display_ui_on_left"`
}

type GameEventIngameEventsChanged struct {
}

type GameEventDotaMatchSignout struct {
}

type GameEventDotaIllusionsCreated struct {
	OriginalEntindex int32 `json:"original_entindex"`
}

type GameEventDotaYearBeastKilled struct {
	KillerPlayerId int32  `json:"killer_player_id"`
	Message        int32  `json:"message"`
	BeastId        uint64 `json:"beast_id"`
}

type GameEventDotaHeroUndoselection struct {
	Playerid1 int32 `json:"playerid1"`
}

type GameEventDotaChallengeSocacheUpdated struct {
}

type GameEventPartyInvitesUpdated struct {
}

type GameEventLobbyInvitesUpdated struct {
}

type GameEventCustomGameModeListUpdated struct {
}

type GameEventCustomGameLobbyListUpdated struct {
}

type GameEventFriendLobbyListUpdated struct {
}

type GameEventDotaTeamPlayerListChanged struct {
}

type GameEventDotaPlayerDetailsChanged struct {
}

type GameEventPlayerProfileStatsUpdated struct {
	AccountId uint64 `json:"account_id"`
}

type GameEventCustomGamePlayerCountUpdated struct {
	CustomGameId uint64 `json:"custom_game_id"`
}

type GameEventCustomGameFriendsPlayedUpdated struct {
	CustomGameId uint64 `json:"custom_game_id"`
}

type GameEventCustomGamesFriendsPlayUpdated struct {
}

type GameEventDotaPlayerUpdateAssignedHero struct {
}

type GameEventDotaPlayerHeroSelectionDirty struct {
}

type GameEventDotaNpcGoalReached struct {
	NpcEntindex      int32 `json:"npc_entindex"`
	GoalEntindex     int32 `json:"goal_entindex"`
	NextGoalEntindex int32 `json:"next_goal_entindex"`
}

type GameEventDotaPlayerSelectedCustomTeam struct {
	PlayerId int32 `json:"player_id"`
	TeamId   int32 `json:"team_id"`
	Success  bool  `json:"success"`
}

type GameEventHltvStatus struct {
	Clients int32  `json:"clients"`
	Slots   int32  `json:"slots"`
	Proxies int32  `json:"proxies"`
	Master  string `json:"master"`
}

type GameEventHltvCameraman struct {
	Index int32 `json:"index"`
}

type GameEventHltvRankCamera struct {
	Index  int32   `json:"index"`
	Rank   float32 `json:"rank"`
	Target int32   `json:"target"`
}

type GameEventHltvRankEntity struct {
	Index  int32   `json:"index"`
	Rank   float32 `json:"rank"`
	Target int32   `json:"target"`
}

type GameEventHltvFixed struct {
	Posx   int32   `json:"posx"`
	Posy   int32   `json:"posy"`
	Posz   int32   `json:"posz"`
	Theta  int32   `json:"theta"`
	Phi    int32   `json:"phi"`
	Offset int32   `json:"offset"`
	Fov    float32 `json:"fov"`
	Target int32   `json:"target"`
}

type GameEventHltvChase struct {
	Target1  int32 `json:"target1"`
	Target2  int32 `json:"target2"`
	Distance int32 `json:"distance"`
	Theta    int32 `json:"theta"`
	Phi      int32 `json:"phi"`
	Inertia  int32 `json:"inertia"`
	Ineye    int32 `json:"ineye"`
}

type GameEventHltvMessage struct {
	Text string `json:"text"`
}

type GameEventHltvTitle struct {
	Text string `json:"text"`
}

type GameEventHltvChat struct {
	Name    string `json:"name"`
	Text    string `json:"text"`
	SteamID uint64 `json:"steamID"`
}

type GameEventHltvVersioninfo struct {
	Version int32 `json:"version"`
}

type GameEventDotaChaseHero struct {
	Target1         int32   `json:"target1"`
	Target2         int32   `json:"target2"`
	Type            int32   `json:"type"`
	Priority        int32   `json:"priority"`
	Gametime        float32 `json:"gametime"`
	Highlight       bool    `json:"highlight"`
	Target1playerid int32   `json:"target1playerid"`
	Target2playerid int32   `json:"target2playerid"`
	Eventtype       int32   `json:"eventtype"`
}

type GameEventDotaCombatlog struct {
	Type             int32   `json:"type"`
	Sourcename       int32   `json:"sourcename"`
	Targetname       int32   `json:"targetname"`
	Attackername     int32   `json:"attackername"`
	Inflictorname    int32   `json:"inflictorname"`
	Attackerillusion bool    `json:"attackerillusion"`
	Targetillusion   bool    `json:"targetillusion"`
	Value            int32   `json:"value"`
	Health           int32   `json:"health"`
	Timestamp        float32 `json:"timestamp"`
	Targetsourcename int32   `json:"targetsourcename"`
	Timestampraw     float32 `json:"timestampraw"`
	Attackerhero     bool    `json:"attackerhero"`
	Targethero       bool    `json:"targethero"`
	AbilityToggleOn  bool    `json:"ability_toggle_on"`
	AbilityToggleOff bool    `json:"ability_toggle_off"`
	AbilityLevel     int32   `json:"ability_level"`
	GoldReason       int32   `json:"gold_reason"`
	XpReason         int32   `json:"xp_reason"`
}

type GameEventDotaGameStateChange struct {
	OldState int32 `json:"old_state"`
	NewState int32 `json:"new_state"`
}

type GameEventDotaPlayerPickHero struct {
	Player    int32  `json:"player"`
	Heroindex int32  `json:"heroindex"`
	Hero      string `json:"hero"`
}

type GameEventDotaTeamKillCredit struct {
	KillerUserid int32 `json:"killer_userid"`
	VictimUserid int32 `json:"victim_userid"`
	Teamnumber   int32 `json:"teamnumber"`
	Herokills    int32 `json:"herokills"`
}

type GameEvents struct {
	onServerSpawn                        []func(*GameEventServerSpawn) error
	onServerPreShutdown                  []func(*GameEventServerPreShutdown) error
	onServerShutdown                     []func(*GameEventServerShutdown) error
	onServerCvar                         []func(*GameEventServerCvar) error
	onServerMessage                      []func(*GameEventServerMessage) error
	onServerAddban                       []func(*GameEventServerAddban) error
	onServerRemoveban                    []func(*GameEventServerRemoveban) error
	onPlayerConnect                      []func(*GameEventPlayerConnect) error
	onPlayerInfo                         []func(*GameEventPlayerInfo) error
	onPlayerDisconnect                   []func(*GameEventPlayerDisconnect) error
	onPlayerActivate                     []func(*GameEventPlayerActivate) error
	onPlayerConnectFull                  []func(*GameEventPlayerConnectFull) error
	onPlayerSay                          []func(*GameEventPlayerSay) error
	onPlayerFullUpdate                   []func(*GameEventPlayerFullUpdate) error
	onTeamInfo                           []func(*GameEventTeamInfo) error
	onTeamScore                          []func(*GameEventTeamScore) error
	onTeamplayBroadcastAudio             []func(*GameEventTeamplayBroadcastAudio) error
	onPlayerTeam                         []func(*GameEventPlayerTeam) error
	onPlayerClass                        []func(*GameEventPlayerClass) error
	onPlayerDeath                        []func(*GameEventPlayerDeath) error
	onPlayerHurt                         []func(*GameEventPlayerHurt) error
	onPlayerChat                         []func(*GameEventPlayerChat) error
	onPlayerScore                        []func(*GameEventPlayerScore) error
	onPlayerSpawn                        []func(*GameEventPlayerSpawn) error
	onPlayerShoot                        []func(*GameEventPlayerShoot) error
	onPlayerUse                          []func(*GameEventPlayerUse) error
	onPlayerChangename                   []func(*GameEventPlayerChangename) error
	onPlayerHintmessage                  []func(*GameEventPlayerHintmessage) error
	onGameInit                           []func(*GameEventGameInit) error
	onGameNewmap                         []func(*GameEventGameNewmap) error
	onGameStart                          []func(*GameEventGameStart) error
	onGameEnd                            []func(*GameEventGameEnd) error
	onRoundStart                         []func(*GameEventRoundStart) error
	onRoundEnd                           []func(*GameEventRoundEnd) error
	onRoundStartPreEntity                []func(*GameEventRoundStartPreEntity) error
	onTeamplayRoundStart                 []func(*GameEventTeamplayRoundStart) error
	onHostnameChanged                    []func(*GameEventHostnameChanged) error
	onDifficultyChanged                  []func(*GameEventDifficultyChanged) error
	onFinaleStart                        []func(*GameEventFinaleStart) error
	onGameMessage                        []func(*GameEventGameMessage) error
	onBreakBreakable                     []func(*GameEventBreakBreakable) error
	onBreakProp                          []func(*GameEventBreakProp) error
	onNpcSpawned                         []func(*GameEventNpcSpawned) error
	onNpcReplaced                        []func(*GameEventNpcReplaced) error
	onEntityKilled                       []func(*GameEventEntityKilled) error
	onEntityHurt                         []func(*GameEventEntityHurt) error
	onBonusUpdated                       []func(*GameEventBonusUpdated) error
	onPlayerStatsUpdated                 []func(*GameEventPlayerStatsUpdated) error
	onAchievementEvent                   []func(*GameEventAchievementEvent) error
	onAchievementEarned                  []func(*GameEventAchievementEarned) error
	onAchievementWriteFailed             []func(*GameEventAchievementWriteFailed) error
	onPhysgunPickup                      []func(*GameEventPhysgunPickup) error
	onFlareIgniteNpc                     []func(*GameEventFlareIgniteNpc) error
	onHelicopterGrenadePuntMiss          []func(*GameEventHelicopterGrenadePuntMiss) error
	onUserDataDownloaded                 []func(*GameEventUserDataDownloaded) error
	onRagdollDissolved                   []func(*GameEventRagdollDissolved) error
	onGameinstructorDraw                 []func(*GameEventGameinstructorDraw) error
	onGameinstructorNodraw               []func(*GameEventGameinstructorNodraw) error
	onMapTransition                      []func(*GameEventMapTransition) error
	onInstructorServerHintCreate         []func(*GameEventInstructorServerHintCreate) error
	onInstructorServerHintStop           []func(*GameEventInstructorServerHintStop) error
	onChatNewMessage                     []func(*GameEventChatNewMessage) error
	onChatMembersChanged                 []func(*GameEventChatMembersChanged) error
	onInventoryUpdated                   []func(*GameEventInventoryUpdated) error
	onCartUpdated                        []func(*GameEventCartUpdated) error
	onStorePricesheetUpdated             []func(*GameEventStorePricesheetUpdated) error
	onGcConnected                        []func(*GameEventGcConnected) error
	onItemSchemaInitialized              []func(*GameEventItemSchemaInitialized) error
	onDropRateModified                   []func(*GameEventDropRateModified) error
	onEventTicketModified                []func(*GameEventEventTicketModified) error
	onModifierEvent                      []func(*GameEventModifierEvent) error
	onDotaPlayerKill                     []func(*GameEventDotaPlayerKill) error
	onDotaPlayerDeny                     []func(*GameEventDotaPlayerDeny) error
	onDotaBarracksKill                   []func(*GameEventDotaBarracksKill) error
	onDotaTowerKill                      []func(*GameEventDotaTowerKill) error
	onDotaEffigyKill                     []func(*GameEventDotaEffigyKill) error
	onDotaRoshanKill                     []func(*GameEventDotaRoshanKill) error
	onDotaCourierLost                    []func(*GameEventDotaCourierLost) error
	onDotaCourierRespawned               []func(*GameEventDotaCourierRespawned) error
	onDotaGlyphUsed                      []func(*GameEventDotaGlyphUsed) error
	onDotaSuperCreeps                    []func(*GameEventDotaSuperCreeps) error
	onDotaItemPurchase                   []func(*GameEventDotaItemPurchase) error
	onDotaItemGifted                     []func(*GameEventDotaItemGifted) error
	onDotaRunePickup                     []func(*GameEventDotaRunePickup) error
	onDotaRuneSpotted                    []func(*GameEventDotaRuneSpotted) error
	onDotaItemSpotted                    []func(*GameEventDotaItemSpotted) error
	onDotaNoBattlePoints                 []func(*GameEventDotaNoBattlePoints) error
	onDotaChatInformational              []func(*GameEventDotaChatInformational) error
	onDotaActionItem                     []func(*GameEventDotaActionItem) error
	onDotaChatBanNotification            []func(*GameEventDotaChatBanNotification) error
	onDotaChatEvent                      []func(*GameEventDotaChatEvent) error
	onDotaChatTimedReward                []func(*GameEventDotaChatTimedReward) error
	onDotaPauseEvent                     []func(*GameEventDotaPauseEvent) error
	onDotaChatKillStreak                 []func(*GameEventDotaChatKillStreak) error
	onDotaChatFirstBlood                 []func(*GameEventDotaChatFirstBlood) error
	onDotaChatAssassinAnnounce           []func(*GameEventDotaChatAssassinAnnounce) error
	onDotaChatAssassinDenied             []func(*GameEventDotaChatAssassinDenied) error
	onDotaChatAssassinSuccess            []func(*GameEventDotaChatAssassinSuccess) error
	onDotaPlayerUpdateHeroSelection      []func(*GameEventDotaPlayerUpdateHeroSelection) error
	onDotaPlayerUpdateSelectedUnit       []func(*GameEventDotaPlayerUpdateSelectedUnit) error
	onDotaPlayerUpdateQueryUnit          []func(*GameEventDotaPlayerUpdateQueryUnit) error
	onDotaPlayerUpdateKillcamUnit        []func(*GameEventDotaPlayerUpdateKillcamUnit) error
	onDotaPlayerTakeTowerDamage          []func(*GameEventDotaPlayerTakeTowerDamage) error
	onDotaHudErrorMessage                []func(*GameEventDotaHudErrorMessage) error
	onDotaActionSuccess                  []func(*GameEventDotaActionSuccess) error
	onDotaStartingPositionChanged        []func(*GameEventDotaStartingPositionChanged) error
	onDotaMoneyChanged                   []func(*GameEventDotaMoneyChanged) error
	onDotaEnemyMoneyChanged              []func(*GameEventDotaEnemyMoneyChanged) error
	onDotaPortraitUnitStatsChanged       []func(*GameEventDotaPortraitUnitStatsChanged) error
	onDotaPortraitUnitModifiersChanged   []func(*GameEventDotaPortraitUnitModifiersChanged) error
	onDotaForcePortraitUpdate            []func(*GameEventDotaForcePortraitUpdate) error
	onDotaInventoryChanged               []func(*GameEventDotaInventoryChanged) error
	onDotaItemPickedUp                   []func(*GameEventDotaItemPickedUp) error
	onDotaInventoryItemChanged           []func(*GameEventDotaInventoryItemChanged) error
	onDotaAbilityChanged                 []func(*GameEventDotaAbilityChanged) error
	onDotaPortraitAbilityLayoutChanged   []func(*GameEventDotaPortraitAbilityLayoutChanged) error
	onDotaInventoryItemAdded             []func(*GameEventDotaInventoryItemAdded) error
	onDotaInventoryChangedQueryUnit      []func(*GameEventDotaInventoryChangedQueryUnit) error
	onDotaLinkClicked                    []func(*GameEventDotaLinkClicked) error
	onDotaSetQuickBuy                    []func(*GameEventDotaSetQuickBuy) error
	onDotaQuickBuyChanged                []func(*GameEventDotaQuickBuyChanged) error
	onDotaPlayerShopChanged              []func(*GameEventDotaPlayerShopChanged) error
	onDotaPlayerShowKillcam              []func(*GameEventDotaPlayerShowKillcam) error
	onDotaPlayerShowMinikillcam          []func(*GameEventDotaPlayerShowMinikillcam) error
	onGcUserSessionCreated               []func(*GameEventGcUserSessionCreated) error
	onTeamDataUpdated                    []func(*GameEventTeamDataUpdated) error
	onGuildDataUpdated                   []func(*GameEventGuildDataUpdated) error
	onGuildOpenPartiesUpdated            []func(*GameEventGuildOpenPartiesUpdated) error
	onFantasyUpdated                     []func(*GameEventFantasyUpdated) error
	onFantasyLeagueChanged               []func(*GameEventFantasyLeagueChanged) error
	onFantasyScoreInfoChanged            []func(*GameEventFantasyScoreInfoChanged) error
	onPlayerInfoUpdated                  []func(*GameEventPlayerInfoUpdated) error
	onPlayerInfoIndividualUpdated        []func(*GameEventPlayerInfoIndividualUpdated) error
	onGameRulesStateChange               []func(*GameEventGameRulesStateChange) error
	onMatchHistoryUpdated                []func(*GameEventMatchHistoryUpdated) error
	onMatchDetailsUpdated                []func(*GameEventMatchDetailsUpdated) error
	onLiveGamesUpdated                   []func(*GameEventLiveGamesUpdated) error
	onRecentMatchesUpdated               []func(*GameEventRecentMatchesUpdated) error
	onNewsUpdated                        []func(*GameEventNewsUpdated) error
	onPersonaUpdated                     []func(*GameEventPersonaUpdated) error
	onTournamentStateUpdated             []func(*GameEventTournamentStateUpdated) error
	onPartyUpdated                       []func(*GameEventPartyUpdated) error
	onLobbyUpdated                       []func(*GameEventLobbyUpdated) error
	onDashboardCachesCleared             []func(*GameEventDashboardCachesCleared) error
	onLastHit                            []func(*GameEventLastHit) error
	onPlayerCompletedGame                []func(*GameEventPlayerCompletedGame) error
	onPlayerReconnected                  []func(*GameEventPlayerReconnected) error
	onNommedTree                         []func(*GameEventNommedTree) error
	onDotaRuneActivatedServer            []func(*GameEventDotaRuneActivatedServer) error
	onDotaPlayerGainedLevel              []func(*GameEventDotaPlayerGainedLevel) error
	onDotaPlayerLearnedAbility           []func(*GameEventDotaPlayerLearnedAbility) error
	onDotaPlayerUsedAbility              []func(*GameEventDotaPlayerUsedAbility) error
	onDotaNonPlayerUsedAbility           []func(*GameEventDotaNonPlayerUsedAbility) error
	onDotaPlayerBeginCast                []func(*GameEventDotaPlayerBeginCast) error
	onDotaNonPlayerBeginCast             []func(*GameEventDotaNonPlayerBeginCast) error
	onDotaAbilityChannelFinished         []func(*GameEventDotaAbilityChannelFinished) error
	onDotaHoldoutReviveComplete          []func(*GameEventDotaHoldoutReviveComplete) error
	onDotaPlayerKilled                   []func(*GameEventDotaPlayerKilled) error
	onBindpanelOpen                      []func(*GameEventBindpanelOpen) error
	onBindpanelClose                     []func(*GameEventBindpanelClose) error
	onKeybindChanged                     []func(*GameEventKeybindChanged) error
	onDotaItemDragBegin                  []func(*GameEventDotaItemDragBegin) error
	onDotaItemDragEnd                    []func(*GameEventDotaItemDragEnd) error
	onDotaShopItemDragBegin              []func(*GameEventDotaShopItemDragBegin) error
	onDotaShopItemDragEnd                []func(*GameEventDotaShopItemDragEnd) error
	onDotaItemPurchased                  []func(*GameEventDotaItemPurchased) error
	onDotaItemCombined                   []func(*GameEventDotaItemCombined) error
	onDotaItemUsed                       []func(*GameEventDotaItemUsed) error
	onDotaItemAutoPurchase               []func(*GameEventDotaItemAutoPurchase) error
	onDotaUnitEvent                      []func(*GameEventDotaUnitEvent) error
	onDotaQuestStarted                   []func(*GameEventDotaQuestStarted) error
	onDotaQuestCompleted                 []func(*GameEventDotaQuestCompleted) error
	onGameuiActivated                    []func(*GameEventGameuiActivated) error
	onGameuiHidden                       []func(*GameEventGameuiHidden) error
	onPlayerFullyjoined                  []func(*GameEventPlayerFullyjoined) error
	onDotaSpectateHero                   []func(*GameEventDotaSpectateHero) error
	onDotaMatchDone                      []func(*GameEventDotaMatchDone) error
	onDotaMatchDoneClient                []func(*GameEventDotaMatchDoneClient) error
	onSetInstructorGroupEnabled          []func(*GameEventSetInstructorGroupEnabled) error
	onJoinedChatChannel                  []func(*GameEventJoinedChatChannel) error
	onLeftChatChannel                    []func(*GameEventLeftChatChannel) error
	onGcChatChannelListUpdated           []func(*GameEventGcChatChannelListUpdated) error
	onTodayMessagesUpdated               []func(*GameEventTodayMessagesUpdated) error
	onFileDownloaded                     []func(*GameEventFileDownloaded) error
	onPlayerReportCountsUpdated          []func(*GameEventPlayerReportCountsUpdated) error
	onScaleformFileDownloadComplete      []func(*GameEventScaleformFileDownloadComplete) error
	onItemPurchased                      []func(*GameEventItemPurchased) error
	onGcMismatchedVersion                []func(*GameEventGcMismatchedVersion) error
	onDemoStop                           []func(*GameEventDemoStop) error
	onMapShutdown                        []func(*GameEventMapShutdown) error
	onDotaWorkshopFileselected           []func(*GameEventDotaWorkshopFileselected) error
	onDotaWorkshopFilecanceled           []func(*GameEventDotaWorkshopFilecanceled) error
	onRichPresenceUpdated                []func(*GameEventRichPresenceUpdated) error
	onDotaHeroRandom                     []func(*GameEventDotaHeroRandom) error
	onDotaRdChatTurn                     []func(*GameEventDotaRdChatTurn) error
	onDotaFavoriteHeroesUpdated          []func(*GameEventDotaFavoriteHeroesUpdated) error
	onProfileOpened                      []func(*GameEventProfileOpened) error
	onProfileClosed                      []func(*GameEventProfileClosed) error
	onItemPreviewClosed                  []func(*GameEventItemPreviewClosed) error
	onDashboardSwitchedSection           []func(*GameEventDashboardSwitchedSection) error
	onDotaTournamentItemEvent            []func(*GameEventDotaTournamentItemEvent) error
	onDotaHeroSwap                       []func(*GameEventDotaHeroSwap) error
	onDotaResetSuggestedItems            []func(*GameEventDotaResetSuggestedItems) error
	onHalloweenHighScoreReceived         []func(*GameEventHalloweenHighScoreReceived) error
	onHalloweenPhaseEnd                  []func(*GameEventHalloweenPhaseEnd) error
	onHalloweenHighScoreRequestFailed    []func(*GameEventHalloweenHighScoreRequestFailed) error
	onDotaHudSkinChanged                 []func(*GameEventDotaHudSkinChanged) error
	onDotaInventoryPlayerGotItem         []func(*GameEventDotaInventoryPlayerGotItem) error
	onPlayerIsExperienced                []func(*GameEventPlayerIsExperienced) error
	onPlayerIsNotexperienced             []func(*GameEventPlayerIsNotexperienced) error
	onDotaTutorialLessonStart            []func(*GameEventDotaTutorialLessonStart) error
	onDotaTutorialTaskAdvance            []func(*GameEventDotaTutorialTaskAdvance) error
	onDotaTutorialShopToggled            []func(*GameEventDotaTutorialShopToggled) error
	onMapLocationUpdated                 []func(*GameEventMapLocationUpdated) error
	onRichpresenceCustomUpdated          []func(*GameEventRichpresenceCustomUpdated) error
	onGameEndVisible                     []func(*GameEventGameEndVisible) error
	onAntiaddictionUpdate                []func(*GameEventAntiaddictionUpdate) error
	onHighlightHudElement                []func(*GameEventHighlightHudElement) error
	onHideHighlightHudElement            []func(*GameEventHideHighlightHudElement) error
	onIntroVideoFinished                 []func(*GameEventIntroVideoFinished) error
	onMatchmakingStatusVisibilityChanged []func(*GameEventMatchmakingStatusVisibilityChanged) error
	onPracticeLobbyVisibilityChanged     []func(*GameEventPracticeLobbyVisibilityChanged) error
	onDotaCourierTransferItem            []func(*GameEventDotaCourierTransferItem) error
	onFullUiUnlocked                     []func(*GameEventFullUiUnlocked) error
	onHeroSelectorPreviewSet             []func(*GameEventHeroSelectorPreviewSet) error
	onAntiaddictionToast                 []func(*GameEventAntiaddictionToast) error
	onHeroPickerShown                    []func(*GameEventHeroPickerShown) error
	onHeroPickerHidden                   []func(*GameEventHeroPickerHidden) error
	onDotaLocalQuickbuyChanged           []func(*GameEventDotaLocalQuickbuyChanged) error
	onShowCenterMessage                  []func(*GameEventShowCenterMessage) error
	onHudFlipChanged                     []func(*GameEventHudFlipChanged) error
	onFrostyPointsUpdated                []func(*GameEventFrostyPointsUpdated) error
	onDefeated                           []func(*GameEventDefeated) error
	onResetDefeated                      []func(*GameEventResetDefeated) error
	onBoosterStateUpdated                []func(*GameEventBoosterStateUpdated) error
	onEventPointsUpdated                 []func(*GameEventEventPointsUpdated) error
	onLocalPlayerEventPoints             []func(*GameEventLocalPlayerEventPoints) error
	onCustomGameDifficulty               []func(*GameEventCustomGameDifficulty) error
	onTreeCut                            []func(*GameEventTreeCut) error
	onUgcDetailsArrived                  []func(*GameEventUgcDetailsArrived) error
	onUgcSubscribed                      []func(*GameEventUgcSubscribed) error
	onUgcUnsubscribed                    []func(*GameEventUgcUnsubscribed) error
	onUgcDownloadRequested               []func(*GameEventUgcDownloadRequested) error
	onUgcInstalled                       []func(*GameEventUgcInstalled) error
	onPrizepoolReceived                  []func(*GameEventPrizepoolReceived) error
	onMicrotransactionSuccess            []func(*GameEventMicrotransactionSuccess) error
	onDotaRubickAbilitySteal             []func(*GameEventDotaRubickAbilitySteal) error
	onCompendiumEventActionsLoaded       []func(*GameEventCompendiumEventActionsLoaded) error
	onCompendiumSelectionsLoaded         []func(*GameEventCompendiumSelectionsLoaded) error
	onCompendiumSetSelectionFailed       []func(*GameEventCompendiumSetSelectionFailed) error
	onCompendiumTrophiesLoaded           []func(*GameEventCompendiumTrophiesLoaded) error
	onCommunityCachedNamesUpdated        []func(*GameEventCommunityCachedNamesUpdated) error
	onSpecItemPickup                     []func(*GameEventSpecItemPickup) error
	onSpecAegisReclaimTime               []func(*GameEventSpecAegisReclaimTime) error
	onAccountTrophiesChanged             []func(*GameEventAccountTrophiesChanged) error
	onAccountAllHeroChallengeChanged     []func(*GameEventAccountAllHeroChallengeChanged) error
	onTeamShowcaseUiUpdate               []func(*GameEventTeamShowcaseUiUpdate) error
	onIngameEventsChanged                []func(*GameEventIngameEventsChanged) error
	onDotaMatchSignout                   []func(*GameEventDotaMatchSignout) error
	onDotaIllusionsCreated               []func(*GameEventDotaIllusionsCreated) error
	onDotaYearBeastKilled                []func(*GameEventDotaYearBeastKilled) error
	onDotaHeroUndoselection              []func(*GameEventDotaHeroUndoselection) error
	onDotaChallengeSocacheUpdated        []func(*GameEventDotaChallengeSocacheUpdated) error
	onPartyInvitesUpdated                []func(*GameEventPartyInvitesUpdated) error
	onLobbyInvitesUpdated                []func(*GameEventLobbyInvitesUpdated) error
	onCustomGameModeListUpdated          []func(*GameEventCustomGameModeListUpdated) error
	onCustomGameLobbyListUpdated         []func(*GameEventCustomGameLobbyListUpdated) error
	onFriendLobbyListUpdated             []func(*GameEventFriendLobbyListUpdated) error
	onDotaTeamPlayerListChanged          []func(*GameEventDotaTeamPlayerListChanged) error
	onDotaPlayerDetailsChanged           []func(*GameEventDotaPlayerDetailsChanged) error
	onPlayerProfileStatsUpdated          []func(*GameEventPlayerProfileStatsUpdated) error
	onCustomGamePlayerCountUpdated       []func(*GameEventCustomGamePlayerCountUpdated) error
	onCustomGameFriendsPlayedUpdated     []func(*GameEventCustomGameFriendsPlayedUpdated) error
	onCustomGamesFriendsPlayUpdated      []func(*GameEventCustomGamesFriendsPlayUpdated) error
	onDotaPlayerUpdateAssignedHero       []func(*GameEventDotaPlayerUpdateAssignedHero) error
	onDotaPlayerHeroSelectionDirty       []func(*GameEventDotaPlayerHeroSelectionDirty) error
	onDotaNpcGoalReached                 []func(*GameEventDotaNpcGoalReached) error
	onDotaPlayerSelectedCustomTeam       []func(*GameEventDotaPlayerSelectedCustomTeam) error
	onHltvStatus                         []func(*GameEventHltvStatus) error
	onHltvCameraman                      []func(*GameEventHltvCameraman) error
	onHltvRankCamera                     []func(*GameEventHltvRankCamera) error
	onHltvRankEntity                     []func(*GameEventHltvRankEntity) error
	onHltvFixed                          []func(*GameEventHltvFixed) error
	onHltvChase                          []func(*GameEventHltvChase) error
	onHltvMessage                        []func(*GameEventHltvMessage) error
	onHltvTitle                          []func(*GameEventHltvTitle) error
	onHltvChat                           []func(*GameEventHltvChat) error
	onHltvVersioninfo                    []func(*GameEventHltvVersioninfo) error
	onDotaChaseHero                      []func(*GameEventDotaChaseHero) error
	onDotaCombatlog                      []func(*GameEventDotaCombatlog) error
	onDotaGameStateChange                []func(*GameEventDotaGameStateChange) error
	onDotaPlayerPickHero                 []func(*GameEventDotaPlayerPickHero) error
	onDotaTeamKillCredit                 []func(*GameEventDotaTeamKillCredit) error
}

func (ge *GameEvents) OnServerSpawn(fn func(*GameEventServerSpawn) error) {
	ge.onServerSpawn = append(ge.onServerSpawn, fn)
}

func (ge *GameEvents) OnServerPreShutdown(fn func(*GameEventServerPreShutdown) error) {
	ge.onServerPreShutdown = append(ge.onServerPreShutdown, fn)
}

func (ge *GameEvents) OnServerShutdown(fn func(*GameEventServerShutdown) error) {
	ge.onServerShutdown = append(ge.onServerShutdown, fn)
}

func (ge *GameEvents) OnServerCvar(fn func(*GameEventServerCvar) error) {
	ge.onServerCvar = append(ge.onServerCvar, fn)
}

func (ge *GameEvents) OnServerMessage(fn func(*GameEventServerMessage) error) {
	ge.onServerMessage = append(ge.onServerMessage, fn)
}

func (ge *GameEvents) OnServerAddban(fn func(*GameEventServerAddban) error) {
	ge.onServerAddban = append(ge.onServerAddban, fn)
}

func (ge *GameEvents) OnServerRemoveban(fn func(*GameEventServerRemoveban) error) {
	ge.onServerRemoveban = append(ge.onServerRemoveban, fn)
}

func (ge *GameEvents) OnPlayerConnect(fn func(*GameEventPlayerConnect) error) {
	ge.onPlayerConnect = append(ge.onPlayerConnect, fn)
}

func (ge *GameEvents) OnPlayerInfo(fn func(*GameEventPlayerInfo) error) {
	ge.onPlayerInfo = append(ge.onPlayerInfo, fn)
}

func (ge *GameEvents) OnPlayerDisconnect(fn func(*GameEventPlayerDisconnect) error) {
	ge.onPlayerDisconnect = append(ge.onPlayerDisconnect, fn)
}

func (ge *GameEvents) OnPlayerActivate(fn func(*GameEventPlayerActivate) error) {
	ge.onPlayerActivate = append(ge.onPlayerActivate, fn)
}

func (ge *GameEvents) OnPlayerConnectFull(fn func(*GameEventPlayerConnectFull) error) {
	ge.onPlayerConnectFull = append(ge.onPlayerConnectFull, fn)
}

func (ge *GameEvents) OnPlayerSay(fn func(*GameEventPlayerSay) error) {
	ge.onPlayerSay = append(ge.onPlayerSay, fn)
}

func (ge *GameEvents) OnPlayerFullUpdate(fn func(*GameEventPlayerFullUpdate) error) {
	ge.onPlayerFullUpdate = append(ge.onPlayerFullUpdate, fn)
}

func (ge *GameEvents) OnTeamInfo(fn func(*GameEventTeamInfo) error) {
	ge.onTeamInfo = append(ge.onTeamInfo, fn)
}

func (ge *GameEvents) OnTeamScore(fn func(*GameEventTeamScore) error) {
	ge.onTeamScore = append(ge.onTeamScore, fn)
}

func (ge *GameEvents) OnTeamplayBroadcastAudio(fn func(*GameEventTeamplayBroadcastAudio) error) {
	ge.onTeamplayBroadcastAudio = append(ge.onTeamplayBroadcastAudio, fn)
}

func (ge *GameEvents) OnPlayerTeam(fn func(*GameEventPlayerTeam) error) {
	ge.onPlayerTeam = append(ge.onPlayerTeam, fn)
}

func (ge *GameEvents) OnPlayerClass(fn func(*GameEventPlayerClass) error) {
	ge.onPlayerClass = append(ge.onPlayerClass, fn)
}

func (ge *GameEvents) OnPlayerDeath(fn func(*GameEventPlayerDeath) error) {
	ge.onPlayerDeath = append(ge.onPlayerDeath, fn)
}

func (ge *GameEvents) OnPlayerHurt(fn func(*GameEventPlayerHurt) error) {
	ge.onPlayerHurt = append(ge.onPlayerHurt, fn)
}

func (ge *GameEvents) OnPlayerChat(fn func(*GameEventPlayerChat) error) {
	ge.onPlayerChat = append(ge.onPlayerChat, fn)
}

func (ge *GameEvents) OnPlayerScore(fn func(*GameEventPlayerScore) error) {
	ge.onPlayerScore = append(ge.onPlayerScore, fn)
}

func (ge *GameEvents) OnPlayerSpawn(fn func(*GameEventPlayerSpawn) error) {
	ge.onPlayerSpawn = append(ge.onPlayerSpawn, fn)
}

func (ge *GameEvents) OnPlayerShoot(fn func(*GameEventPlayerShoot) error) {
	ge.onPlayerShoot = append(ge.onPlayerShoot, fn)
}

func (ge *GameEvents) OnPlayerUse(fn func(*GameEventPlayerUse) error) {
	ge.onPlayerUse = append(ge.onPlayerUse, fn)
}

func (ge *GameEvents) OnPlayerChangename(fn func(*GameEventPlayerChangename) error) {
	ge.onPlayerChangename = append(ge.onPlayerChangename, fn)
}

func (ge *GameEvents) OnPlayerHintmessage(fn func(*GameEventPlayerHintmessage) error) {
	ge.onPlayerHintmessage = append(ge.onPlayerHintmessage, fn)
}

func (ge *GameEvents) OnGameInit(fn func(*GameEventGameInit) error) {
	ge.onGameInit = append(ge.onGameInit, fn)
}

func (ge *GameEvents) OnGameNewmap(fn func(*GameEventGameNewmap) error) {
	ge.onGameNewmap = append(ge.onGameNewmap, fn)
}

func (ge *GameEvents) OnGameStart(fn func(*GameEventGameStart) error) {
	ge.onGameStart = append(ge.onGameStart, fn)
}

func (ge *GameEvents) OnGameEnd(fn func(*GameEventGameEnd) error) {
	ge.onGameEnd = append(ge.onGameEnd, fn)
}

func (ge *GameEvents) OnRoundStart(fn func(*GameEventRoundStart) error) {
	ge.onRoundStart = append(ge.onRoundStart, fn)
}

func (ge *GameEvents) OnRoundEnd(fn func(*GameEventRoundEnd) error) {
	ge.onRoundEnd = append(ge.onRoundEnd, fn)
}

func (ge *GameEvents) OnRoundStartPreEntity(fn func(*GameEventRoundStartPreEntity) error) {
	ge.onRoundStartPreEntity = append(ge.onRoundStartPreEntity, fn)
}

func (ge *GameEvents) OnTeamplayRoundStart(fn func(*GameEventTeamplayRoundStart) error) {
	ge.onTeamplayRoundStart = append(ge.onTeamplayRoundStart, fn)
}

func (ge *GameEvents) OnHostnameChanged(fn func(*GameEventHostnameChanged) error) {
	ge.onHostnameChanged = append(ge.onHostnameChanged, fn)
}

func (ge *GameEvents) OnDifficultyChanged(fn func(*GameEventDifficultyChanged) error) {
	ge.onDifficultyChanged = append(ge.onDifficultyChanged, fn)
}

func (ge *GameEvents) OnFinaleStart(fn func(*GameEventFinaleStart) error) {
	ge.onFinaleStart = append(ge.onFinaleStart, fn)
}

func (ge *GameEvents) OnGameMessage(fn func(*GameEventGameMessage) error) {
	ge.onGameMessage = append(ge.onGameMessage, fn)
}

func (ge *GameEvents) OnBreakBreakable(fn func(*GameEventBreakBreakable) error) {
	ge.onBreakBreakable = append(ge.onBreakBreakable, fn)
}

func (ge *GameEvents) OnBreakProp(fn func(*GameEventBreakProp) error) {
	ge.onBreakProp = append(ge.onBreakProp, fn)
}

func (ge *GameEvents) OnNpcSpawned(fn func(*GameEventNpcSpawned) error) {
	ge.onNpcSpawned = append(ge.onNpcSpawned, fn)
}

func (ge *GameEvents) OnNpcReplaced(fn func(*GameEventNpcReplaced) error) {
	ge.onNpcReplaced = append(ge.onNpcReplaced, fn)
}

func (ge *GameEvents) OnEntityKilled(fn func(*GameEventEntityKilled) error) {
	ge.onEntityKilled = append(ge.onEntityKilled, fn)
}

func (ge *GameEvents) OnEntityHurt(fn func(*GameEventEntityHurt) error) {
	ge.onEntityHurt = append(ge.onEntityHurt, fn)
}

func (ge *GameEvents) OnBonusUpdated(fn func(*GameEventBonusUpdated) error) {
	ge.onBonusUpdated = append(ge.onBonusUpdated, fn)
}

func (ge *GameEvents) OnPlayerStatsUpdated(fn func(*GameEventPlayerStatsUpdated) error) {
	ge.onPlayerStatsUpdated = append(ge.onPlayerStatsUpdated, fn)
}

func (ge *GameEvents) OnAchievementEvent(fn func(*GameEventAchievementEvent) error) {
	ge.onAchievementEvent = append(ge.onAchievementEvent, fn)
}

func (ge *GameEvents) OnAchievementEarned(fn func(*GameEventAchievementEarned) error) {
	ge.onAchievementEarned = append(ge.onAchievementEarned, fn)
}

func (ge *GameEvents) OnAchievementWriteFailed(fn func(*GameEventAchievementWriteFailed) error) {
	ge.onAchievementWriteFailed = append(ge.onAchievementWriteFailed, fn)
}

func (ge *GameEvents) OnPhysgunPickup(fn func(*GameEventPhysgunPickup) error) {
	ge.onPhysgunPickup = append(ge.onPhysgunPickup, fn)
}

func (ge *GameEvents) OnFlareIgniteNpc(fn func(*GameEventFlareIgniteNpc) error) {
	ge.onFlareIgniteNpc = append(ge.onFlareIgniteNpc, fn)
}

func (ge *GameEvents) OnHelicopterGrenadePuntMiss(fn func(*GameEventHelicopterGrenadePuntMiss) error) {
	ge.onHelicopterGrenadePuntMiss = append(ge.onHelicopterGrenadePuntMiss, fn)
}

func (ge *GameEvents) OnUserDataDownloaded(fn func(*GameEventUserDataDownloaded) error) {
	ge.onUserDataDownloaded = append(ge.onUserDataDownloaded, fn)
}

func (ge *GameEvents) OnRagdollDissolved(fn func(*GameEventRagdollDissolved) error) {
	ge.onRagdollDissolved = append(ge.onRagdollDissolved, fn)
}

func (ge *GameEvents) OnGameinstructorDraw(fn func(*GameEventGameinstructorDraw) error) {
	ge.onGameinstructorDraw = append(ge.onGameinstructorDraw, fn)
}

func (ge *GameEvents) OnGameinstructorNodraw(fn func(*GameEventGameinstructorNodraw) error) {
	ge.onGameinstructorNodraw = append(ge.onGameinstructorNodraw, fn)
}

func (ge *GameEvents) OnMapTransition(fn func(*GameEventMapTransition) error) {
	ge.onMapTransition = append(ge.onMapTransition, fn)
}

func (ge *GameEvents) OnInstructorServerHintCreate(fn func(*GameEventInstructorServerHintCreate) error) {
	ge.onInstructorServerHintCreate = append(ge.onInstructorServerHintCreate, fn)
}

func (ge *GameEvents) OnInstructorServerHintStop(fn func(*GameEventInstructorServerHintStop) error) {
	ge.onInstructorServerHintStop = append(ge.onInstructorServerHintStop, fn)
}

func (ge *GameEvents) OnChatNewMessage(fn func(*GameEventChatNewMessage) error) {
	ge.onChatNewMessage = append(ge.onChatNewMessage, fn)
}

func (ge *GameEvents) OnChatMembersChanged(fn func(*GameEventChatMembersChanged) error) {
	ge.onChatMembersChanged = append(ge.onChatMembersChanged, fn)
}

func (ge *GameEvents) OnInventoryUpdated(fn func(*GameEventInventoryUpdated) error) {
	ge.onInventoryUpdated = append(ge.onInventoryUpdated, fn)
}

func (ge *GameEvents) OnCartUpdated(fn func(*GameEventCartUpdated) error) {
	ge.onCartUpdated = append(ge.onCartUpdated, fn)
}

func (ge *GameEvents) OnStorePricesheetUpdated(fn func(*GameEventStorePricesheetUpdated) error) {
	ge.onStorePricesheetUpdated = append(ge.onStorePricesheetUpdated, fn)
}

func (ge *GameEvents) OnGcConnected(fn func(*GameEventGcConnected) error) {
	ge.onGcConnected = append(ge.onGcConnected, fn)
}

func (ge *GameEvents) OnItemSchemaInitialized(fn func(*GameEventItemSchemaInitialized) error) {
	ge.onItemSchemaInitialized = append(ge.onItemSchemaInitialized, fn)
}

func (ge *GameEvents) OnDropRateModified(fn func(*GameEventDropRateModified) error) {
	ge.onDropRateModified = append(ge.onDropRateModified, fn)
}

func (ge *GameEvents) OnEventTicketModified(fn func(*GameEventEventTicketModified) error) {
	ge.onEventTicketModified = append(ge.onEventTicketModified, fn)
}

func (ge *GameEvents) OnModifierEvent(fn func(*GameEventModifierEvent) error) {
	ge.onModifierEvent = append(ge.onModifierEvent, fn)
}

func (ge *GameEvents) OnDotaPlayerKill(fn func(*GameEventDotaPlayerKill) error) {
	ge.onDotaPlayerKill = append(ge.onDotaPlayerKill, fn)
}

func (ge *GameEvents) OnDotaPlayerDeny(fn func(*GameEventDotaPlayerDeny) error) {
	ge.onDotaPlayerDeny = append(ge.onDotaPlayerDeny, fn)
}

func (ge *GameEvents) OnDotaBarracksKill(fn func(*GameEventDotaBarracksKill) error) {
	ge.onDotaBarracksKill = append(ge.onDotaBarracksKill, fn)
}

func (ge *GameEvents) OnDotaTowerKill(fn func(*GameEventDotaTowerKill) error) {
	ge.onDotaTowerKill = append(ge.onDotaTowerKill, fn)
}

func (ge *GameEvents) OnDotaEffigyKill(fn func(*GameEventDotaEffigyKill) error) {
	ge.onDotaEffigyKill = append(ge.onDotaEffigyKill, fn)
}

func (ge *GameEvents) OnDotaRoshanKill(fn func(*GameEventDotaRoshanKill) error) {
	ge.onDotaRoshanKill = append(ge.onDotaRoshanKill, fn)
}

func (ge *GameEvents) OnDotaCourierLost(fn func(*GameEventDotaCourierLost) error) {
	ge.onDotaCourierLost = append(ge.onDotaCourierLost, fn)
}

func (ge *GameEvents) OnDotaCourierRespawned(fn func(*GameEventDotaCourierRespawned) error) {
	ge.onDotaCourierRespawned = append(ge.onDotaCourierRespawned, fn)
}

func (ge *GameEvents) OnDotaGlyphUsed(fn func(*GameEventDotaGlyphUsed) error) {
	ge.onDotaGlyphUsed = append(ge.onDotaGlyphUsed, fn)
}

func (ge *GameEvents) OnDotaSuperCreeps(fn func(*GameEventDotaSuperCreeps) error) {
	ge.onDotaSuperCreeps = append(ge.onDotaSuperCreeps, fn)
}

func (ge *GameEvents) OnDotaItemPurchase(fn func(*GameEventDotaItemPurchase) error) {
	ge.onDotaItemPurchase = append(ge.onDotaItemPurchase, fn)
}

func (ge *GameEvents) OnDotaItemGifted(fn func(*GameEventDotaItemGifted) error) {
	ge.onDotaItemGifted = append(ge.onDotaItemGifted, fn)
}

func (ge *GameEvents) OnDotaRunePickup(fn func(*GameEventDotaRunePickup) error) {
	ge.onDotaRunePickup = append(ge.onDotaRunePickup, fn)
}

func (ge *GameEvents) OnDotaRuneSpotted(fn func(*GameEventDotaRuneSpotted) error) {
	ge.onDotaRuneSpotted = append(ge.onDotaRuneSpotted, fn)
}

func (ge *GameEvents) OnDotaItemSpotted(fn func(*GameEventDotaItemSpotted) error) {
	ge.onDotaItemSpotted = append(ge.onDotaItemSpotted, fn)
}

func (ge *GameEvents) OnDotaNoBattlePoints(fn func(*GameEventDotaNoBattlePoints) error) {
	ge.onDotaNoBattlePoints = append(ge.onDotaNoBattlePoints, fn)
}

func (ge *GameEvents) OnDotaChatInformational(fn func(*GameEventDotaChatInformational) error) {
	ge.onDotaChatInformational = append(ge.onDotaChatInformational, fn)
}

func (ge *GameEvents) OnDotaActionItem(fn func(*GameEventDotaActionItem) error) {
	ge.onDotaActionItem = append(ge.onDotaActionItem, fn)
}

func (ge *GameEvents) OnDotaChatBanNotification(fn func(*GameEventDotaChatBanNotification) error) {
	ge.onDotaChatBanNotification = append(ge.onDotaChatBanNotification, fn)
}

func (ge *GameEvents) OnDotaChatEvent(fn func(*GameEventDotaChatEvent) error) {
	ge.onDotaChatEvent = append(ge.onDotaChatEvent, fn)
}

func (ge *GameEvents) OnDotaChatTimedReward(fn func(*GameEventDotaChatTimedReward) error) {
	ge.onDotaChatTimedReward = append(ge.onDotaChatTimedReward, fn)
}

func (ge *GameEvents) OnDotaPauseEvent(fn func(*GameEventDotaPauseEvent) error) {
	ge.onDotaPauseEvent = append(ge.onDotaPauseEvent, fn)
}

func (ge *GameEvents) OnDotaChatKillStreak(fn func(*GameEventDotaChatKillStreak) error) {
	ge.onDotaChatKillStreak = append(ge.onDotaChatKillStreak, fn)
}

func (ge *GameEvents) OnDotaChatFirstBlood(fn func(*GameEventDotaChatFirstBlood) error) {
	ge.onDotaChatFirstBlood = append(ge.onDotaChatFirstBlood, fn)
}

func (ge *GameEvents) OnDotaChatAssassinAnnounce(fn func(*GameEventDotaChatAssassinAnnounce) error) {
	ge.onDotaChatAssassinAnnounce = append(ge.onDotaChatAssassinAnnounce, fn)
}

func (ge *GameEvents) OnDotaChatAssassinDenied(fn func(*GameEventDotaChatAssassinDenied) error) {
	ge.onDotaChatAssassinDenied = append(ge.onDotaChatAssassinDenied, fn)
}

func (ge *GameEvents) OnDotaChatAssassinSuccess(fn func(*GameEventDotaChatAssassinSuccess) error) {
	ge.onDotaChatAssassinSuccess = append(ge.onDotaChatAssassinSuccess, fn)
}

func (ge *GameEvents) OnDotaPlayerUpdateHeroSelection(fn func(*GameEventDotaPlayerUpdateHeroSelection) error) {
	ge.onDotaPlayerUpdateHeroSelection = append(ge.onDotaPlayerUpdateHeroSelection, fn)
}

func (ge *GameEvents) OnDotaPlayerUpdateSelectedUnit(fn func(*GameEventDotaPlayerUpdateSelectedUnit) error) {
	ge.onDotaPlayerUpdateSelectedUnit = append(ge.onDotaPlayerUpdateSelectedUnit, fn)
}

func (ge *GameEvents) OnDotaPlayerUpdateQueryUnit(fn func(*GameEventDotaPlayerUpdateQueryUnit) error) {
	ge.onDotaPlayerUpdateQueryUnit = append(ge.onDotaPlayerUpdateQueryUnit, fn)
}

func (ge *GameEvents) OnDotaPlayerUpdateKillcamUnit(fn func(*GameEventDotaPlayerUpdateKillcamUnit) error) {
	ge.onDotaPlayerUpdateKillcamUnit = append(ge.onDotaPlayerUpdateKillcamUnit, fn)
}

func (ge *GameEvents) OnDotaPlayerTakeTowerDamage(fn func(*GameEventDotaPlayerTakeTowerDamage) error) {
	ge.onDotaPlayerTakeTowerDamage = append(ge.onDotaPlayerTakeTowerDamage, fn)
}

func (ge *GameEvents) OnDotaHudErrorMessage(fn func(*GameEventDotaHudErrorMessage) error) {
	ge.onDotaHudErrorMessage = append(ge.onDotaHudErrorMessage, fn)
}

func (ge *GameEvents) OnDotaActionSuccess(fn func(*GameEventDotaActionSuccess) error) {
	ge.onDotaActionSuccess = append(ge.onDotaActionSuccess, fn)
}

func (ge *GameEvents) OnDotaStartingPositionChanged(fn func(*GameEventDotaStartingPositionChanged) error) {
	ge.onDotaStartingPositionChanged = append(ge.onDotaStartingPositionChanged, fn)
}

func (ge *GameEvents) OnDotaMoneyChanged(fn func(*GameEventDotaMoneyChanged) error) {
	ge.onDotaMoneyChanged = append(ge.onDotaMoneyChanged, fn)
}

func (ge *GameEvents) OnDotaEnemyMoneyChanged(fn func(*GameEventDotaEnemyMoneyChanged) error) {
	ge.onDotaEnemyMoneyChanged = append(ge.onDotaEnemyMoneyChanged, fn)
}

func (ge *GameEvents) OnDotaPortraitUnitStatsChanged(fn func(*GameEventDotaPortraitUnitStatsChanged) error) {
	ge.onDotaPortraitUnitStatsChanged = append(ge.onDotaPortraitUnitStatsChanged, fn)
}

func (ge *GameEvents) OnDotaPortraitUnitModifiersChanged(fn func(*GameEventDotaPortraitUnitModifiersChanged) error) {
	ge.onDotaPortraitUnitModifiersChanged = append(ge.onDotaPortraitUnitModifiersChanged, fn)
}

func (ge *GameEvents) OnDotaForcePortraitUpdate(fn func(*GameEventDotaForcePortraitUpdate) error) {
	ge.onDotaForcePortraitUpdate = append(ge.onDotaForcePortraitUpdate, fn)
}

func (ge *GameEvents) OnDotaInventoryChanged(fn func(*GameEventDotaInventoryChanged) error) {
	ge.onDotaInventoryChanged = append(ge.onDotaInventoryChanged, fn)
}

func (ge *GameEvents) OnDotaItemPickedUp(fn func(*GameEventDotaItemPickedUp) error) {
	ge.onDotaItemPickedUp = append(ge.onDotaItemPickedUp, fn)
}

func (ge *GameEvents) OnDotaInventoryItemChanged(fn func(*GameEventDotaInventoryItemChanged) error) {
	ge.onDotaInventoryItemChanged = append(ge.onDotaInventoryItemChanged, fn)
}

func (ge *GameEvents) OnDotaAbilityChanged(fn func(*GameEventDotaAbilityChanged) error) {
	ge.onDotaAbilityChanged = append(ge.onDotaAbilityChanged, fn)
}

func (ge *GameEvents) OnDotaPortraitAbilityLayoutChanged(fn func(*GameEventDotaPortraitAbilityLayoutChanged) error) {
	ge.onDotaPortraitAbilityLayoutChanged = append(ge.onDotaPortraitAbilityLayoutChanged, fn)
}

func (ge *GameEvents) OnDotaInventoryItemAdded(fn func(*GameEventDotaInventoryItemAdded) error) {
	ge.onDotaInventoryItemAdded = append(ge.onDotaInventoryItemAdded, fn)
}

func (ge *GameEvents) OnDotaInventoryChangedQueryUnit(fn func(*GameEventDotaInventoryChangedQueryUnit) error) {
	ge.onDotaInventoryChangedQueryUnit = append(ge.onDotaInventoryChangedQueryUnit, fn)
}

func (ge *GameEvents) OnDotaLinkClicked(fn func(*GameEventDotaLinkClicked) error) {
	ge.onDotaLinkClicked = append(ge.onDotaLinkClicked, fn)
}

func (ge *GameEvents) OnDotaSetQuickBuy(fn func(*GameEventDotaSetQuickBuy) error) {
	ge.onDotaSetQuickBuy = append(ge.onDotaSetQuickBuy, fn)
}

func (ge *GameEvents) OnDotaQuickBuyChanged(fn func(*GameEventDotaQuickBuyChanged) error) {
	ge.onDotaQuickBuyChanged = append(ge.onDotaQuickBuyChanged, fn)
}

func (ge *GameEvents) OnDotaPlayerShopChanged(fn func(*GameEventDotaPlayerShopChanged) error) {
	ge.onDotaPlayerShopChanged = append(ge.onDotaPlayerShopChanged, fn)
}

func (ge *GameEvents) OnDotaPlayerShowKillcam(fn func(*GameEventDotaPlayerShowKillcam) error) {
	ge.onDotaPlayerShowKillcam = append(ge.onDotaPlayerShowKillcam, fn)
}

func (ge *GameEvents) OnDotaPlayerShowMinikillcam(fn func(*GameEventDotaPlayerShowMinikillcam) error) {
	ge.onDotaPlayerShowMinikillcam = append(ge.onDotaPlayerShowMinikillcam, fn)
}

func (ge *GameEvents) OnGcUserSessionCreated(fn func(*GameEventGcUserSessionCreated) error) {
	ge.onGcUserSessionCreated = append(ge.onGcUserSessionCreated, fn)
}

func (ge *GameEvents) OnTeamDataUpdated(fn func(*GameEventTeamDataUpdated) error) {
	ge.onTeamDataUpdated = append(ge.onTeamDataUpdated, fn)
}

func (ge *GameEvents) OnGuildDataUpdated(fn func(*GameEventGuildDataUpdated) error) {
	ge.onGuildDataUpdated = append(ge.onGuildDataUpdated, fn)
}

func (ge *GameEvents) OnGuildOpenPartiesUpdated(fn func(*GameEventGuildOpenPartiesUpdated) error) {
	ge.onGuildOpenPartiesUpdated = append(ge.onGuildOpenPartiesUpdated, fn)
}

func (ge *GameEvents) OnFantasyUpdated(fn func(*GameEventFantasyUpdated) error) {
	ge.onFantasyUpdated = append(ge.onFantasyUpdated, fn)
}

func (ge *GameEvents) OnFantasyLeagueChanged(fn func(*GameEventFantasyLeagueChanged) error) {
	ge.onFantasyLeagueChanged = append(ge.onFantasyLeagueChanged, fn)
}

func (ge *GameEvents) OnFantasyScoreInfoChanged(fn func(*GameEventFantasyScoreInfoChanged) error) {
	ge.onFantasyScoreInfoChanged = append(ge.onFantasyScoreInfoChanged, fn)
}

func (ge *GameEvents) OnPlayerInfoUpdated(fn func(*GameEventPlayerInfoUpdated) error) {
	ge.onPlayerInfoUpdated = append(ge.onPlayerInfoUpdated, fn)
}

func (ge *GameEvents) OnPlayerInfoIndividualUpdated(fn func(*GameEventPlayerInfoIndividualUpdated) error) {
	ge.onPlayerInfoIndividualUpdated = append(ge.onPlayerInfoIndividualUpdated, fn)
}

func (ge *GameEvents) OnGameRulesStateChange(fn func(*GameEventGameRulesStateChange) error) {
	ge.onGameRulesStateChange = append(ge.onGameRulesStateChange, fn)
}

func (ge *GameEvents) OnMatchHistoryUpdated(fn func(*GameEventMatchHistoryUpdated) error) {
	ge.onMatchHistoryUpdated = append(ge.onMatchHistoryUpdated, fn)
}

func (ge *GameEvents) OnMatchDetailsUpdated(fn func(*GameEventMatchDetailsUpdated) error) {
	ge.onMatchDetailsUpdated = append(ge.onMatchDetailsUpdated, fn)
}

func (ge *GameEvents) OnLiveGamesUpdated(fn func(*GameEventLiveGamesUpdated) error) {
	ge.onLiveGamesUpdated = append(ge.onLiveGamesUpdated, fn)
}

func (ge *GameEvents) OnRecentMatchesUpdated(fn func(*GameEventRecentMatchesUpdated) error) {
	ge.onRecentMatchesUpdated = append(ge.onRecentMatchesUpdated, fn)
}

func (ge *GameEvents) OnNewsUpdated(fn func(*GameEventNewsUpdated) error) {
	ge.onNewsUpdated = append(ge.onNewsUpdated, fn)
}

func (ge *GameEvents) OnPersonaUpdated(fn func(*GameEventPersonaUpdated) error) {
	ge.onPersonaUpdated = append(ge.onPersonaUpdated, fn)
}

func (ge *GameEvents) OnTournamentStateUpdated(fn func(*GameEventTournamentStateUpdated) error) {
	ge.onTournamentStateUpdated = append(ge.onTournamentStateUpdated, fn)
}

func (ge *GameEvents) OnPartyUpdated(fn func(*GameEventPartyUpdated) error) {
	ge.onPartyUpdated = append(ge.onPartyUpdated, fn)
}

func (ge *GameEvents) OnLobbyUpdated(fn func(*GameEventLobbyUpdated) error) {
	ge.onLobbyUpdated = append(ge.onLobbyUpdated, fn)
}

func (ge *GameEvents) OnDashboardCachesCleared(fn func(*GameEventDashboardCachesCleared) error) {
	ge.onDashboardCachesCleared = append(ge.onDashboardCachesCleared, fn)
}

func (ge *GameEvents) OnLastHit(fn func(*GameEventLastHit) error) {
	ge.onLastHit = append(ge.onLastHit, fn)
}

func (ge *GameEvents) OnPlayerCompletedGame(fn func(*GameEventPlayerCompletedGame) error) {
	ge.onPlayerCompletedGame = append(ge.onPlayerCompletedGame, fn)
}

func (ge *GameEvents) OnPlayerReconnected(fn func(*GameEventPlayerReconnected) error) {
	ge.onPlayerReconnected = append(ge.onPlayerReconnected, fn)
}

func (ge *GameEvents) OnNommedTree(fn func(*GameEventNommedTree) error) {
	ge.onNommedTree = append(ge.onNommedTree, fn)
}

func (ge *GameEvents) OnDotaRuneActivatedServer(fn func(*GameEventDotaRuneActivatedServer) error) {
	ge.onDotaRuneActivatedServer = append(ge.onDotaRuneActivatedServer, fn)
}

func (ge *GameEvents) OnDotaPlayerGainedLevel(fn func(*GameEventDotaPlayerGainedLevel) error) {
	ge.onDotaPlayerGainedLevel = append(ge.onDotaPlayerGainedLevel, fn)
}

func (ge *GameEvents) OnDotaPlayerLearnedAbility(fn func(*GameEventDotaPlayerLearnedAbility) error) {
	ge.onDotaPlayerLearnedAbility = append(ge.onDotaPlayerLearnedAbility, fn)
}

func (ge *GameEvents) OnDotaPlayerUsedAbility(fn func(*GameEventDotaPlayerUsedAbility) error) {
	ge.onDotaPlayerUsedAbility = append(ge.onDotaPlayerUsedAbility, fn)
}

func (ge *GameEvents) OnDotaNonPlayerUsedAbility(fn func(*GameEventDotaNonPlayerUsedAbility) error) {
	ge.onDotaNonPlayerUsedAbility = append(ge.onDotaNonPlayerUsedAbility, fn)
}

func (ge *GameEvents) OnDotaPlayerBeginCast(fn func(*GameEventDotaPlayerBeginCast) error) {
	ge.onDotaPlayerBeginCast = append(ge.onDotaPlayerBeginCast, fn)
}

func (ge *GameEvents) OnDotaNonPlayerBeginCast(fn func(*GameEventDotaNonPlayerBeginCast) error) {
	ge.onDotaNonPlayerBeginCast = append(ge.onDotaNonPlayerBeginCast, fn)
}

func (ge *GameEvents) OnDotaAbilityChannelFinished(fn func(*GameEventDotaAbilityChannelFinished) error) {
	ge.onDotaAbilityChannelFinished = append(ge.onDotaAbilityChannelFinished, fn)
}

func (ge *GameEvents) OnDotaHoldoutReviveComplete(fn func(*GameEventDotaHoldoutReviveComplete) error) {
	ge.onDotaHoldoutReviveComplete = append(ge.onDotaHoldoutReviveComplete, fn)
}

func (ge *GameEvents) OnDotaPlayerKilled(fn func(*GameEventDotaPlayerKilled) error) {
	ge.onDotaPlayerKilled = append(ge.onDotaPlayerKilled, fn)
}

func (ge *GameEvents) OnBindpanelOpen(fn func(*GameEventBindpanelOpen) error) {
	ge.onBindpanelOpen = append(ge.onBindpanelOpen, fn)
}

func (ge *GameEvents) OnBindpanelClose(fn func(*GameEventBindpanelClose) error) {
	ge.onBindpanelClose = append(ge.onBindpanelClose, fn)
}

func (ge *GameEvents) OnKeybindChanged(fn func(*GameEventKeybindChanged) error) {
	ge.onKeybindChanged = append(ge.onKeybindChanged, fn)
}

func (ge *GameEvents) OnDotaItemDragBegin(fn func(*GameEventDotaItemDragBegin) error) {
	ge.onDotaItemDragBegin = append(ge.onDotaItemDragBegin, fn)
}

func (ge *GameEvents) OnDotaItemDragEnd(fn func(*GameEventDotaItemDragEnd) error) {
	ge.onDotaItemDragEnd = append(ge.onDotaItemDragEnd, fn)
}

func (ge *GameEvents) OnDotaShopItemDragBegin(fn func(*GameEventDotaShopItemDragBegin) error) {
	ge.onDotaShopItemDragBegin = append(ge.onDotaShopItemDragBegin, fn)
}

func (ge *GameEvents) OnDotaShopItemDragEnd(fn func(*GameEventDotaShopItemDragEnd) error) {
	ge.onDotaShopItemDragEnd = append(ge.onDotaShopItemDragEnd, fn)
}

func (ge *GameEvents) OnDotaItemPurchased(fn func(*GameEventDotaItemPurchased) error) {
	ge.onDotaItemPurchased = append(ge.onDotaItemPurchased, fn)
}

func (ge *GameEvents) OnDotaItemCombined(fn func(*GameEventDotaItemCombined) error) {
	ge.onDotaItemCombined = append(ge.onDotaItemCombined, fn)
}

func (ge *GameEvents) OnDotaItemUsed(fn func(*GameEventDotaItemUsed) error) {
	ge.onDotaItemUsed = append(ge.onDotaItemUsed, fn)
}

func (ge *GameEvents) OnDotaItemAutoPurchase(fn func(*GameEventDotaItemAutoPurchase) error) {
	ge.onDotaItemAutoPurchase = append(ge.onDotaItemAutoPurchase, fn)
}

func (ge *GameEvents) OnDotaUnitEvent(fn func(*GameEventDotaUnitEvent) error) {
	ge.onDotaUnitEvent = append(ge.onDotaUnitEvent, fn)
}

func (ge *GameEvents) OnDotaQuestStarted(fn func(*GameEventDotaQuestStarted) error) {
	ge.onDotaQuestStarted = append(ge.onDotaQuestStarted, fn)
}

func (ge *GameEvents) OnDotaQuestCompleted(fn func(*GameEventDotaQuestCompleted) error) {
	ge.onDotaQuestCompleted = append(ge.onDotaQuestCompleted, fn)
}

func (ge *GameEvents) OnGameuiActivated(fn func(*GameEventGameuiActivated) error) {
	ge.onGameuiActivated = append(ge.onGameuiActivated, fn)
}

func (ge *GameEvents) OnGameuiHidden(fn func(*GameEventGameuiHidden) error) {
	ge.onGameuiHidden = append(ge.onGameuiHidden, fn)
}

func (ge *GameEvents) OnPlayerFullyjoined(fn func(*GameEventPlayerFullyjoined) error) {
	ge.onPlayerFullyjoined = append(ge.onPlayerFullyjoined, fn)
}

func (ge *GameEvents) OnDotaSpectateHero(fn func(*GameEventDotaSpectateHero) error) {
	ge.onDotaSpectateHero = append(ge.onDotaSpectateHero, fn)
}

func (ge *GameEvents) OnDotaMatchDone(fn func(*GameEventDotaMatchDone) error) {
	ge.onDotaMatchDone = append(ge.onDotaMatchDone, fn)
}

func (ge *GameEvents) OnDotaMatchDoneClient(fn func(*GameEventDotaMatchDoneClient) error) {
	ge.onDotaMatchDoneClient = append(ge.onDotaMatchDoneClient, fn)
}

func (ge *GameEvents) OnSetInstructorGroupEnabled(fn func(*GameEventSetInstructorGroupEnabled) error) {
	ge.onSetInstructorGroupEnabled = append(ge.onSetInstructorGroupEnabled, fn)
}

func (ge *GameEvents) OnJoinedChatChannel(fn func(*GameEventJoinedChatChannel) error) {
	ge.onJoinedChatChannel = append(ge.onJoinedChatChannel, fn)
}

func (ge *GameEvents) OnLeftChatChannel(fn func(*GameEventLeftChatChannel) error) {
	ge.onLeftChatChannel = append(ge.onLeftChatChannel, fn)
}

func (ge *GameEvents) OnGcChatChannelListUpdated(fn func(*GameEventGcChatChannelListUpdated) error) {
	ge.onGcChatChannelListUpdated = append(ge.onGcChatChannelListUpdated, fn)
}

func (ge *GameEvents) OnTodayMessagesUpdated(fn func(*GameEventTodayMessagesUpdated) error) {
	ge.onTodayMessagesUpdated = append(ge.onTodayMessagesUpdated, fn)
}

func (ge *GameEvents) OnFileDownloaded(fn func(*GameEventFileDownloaded) error) {
	ge.onFileDownloaded = append(ge.onFileDownloaded, fn)
}

func (ge *GameEvents) OnPlayerReportCountsUpdated(fn func(*GameEventPlayerReportCountsUpdated) error) {
	ge.onPlayerReportCountsUpdated = append(ge.onPlayerReportCountsUpdated, fn)
}

func (ge *GameEvents) OnScaleformFileDownloadComplete(fn func(*GameEventScaleformFileDownloadComplete) error) {
	ge.onScaleformFileDownloadComplete = append(ge.onScaleformFileDownloadComplete, fn)
}

func (ge *GameEvents) OnItemPurchased(fn func(*GameEventItemPurchased) error) {
	ge.onItemPurchased = append(ge.onItemPurchased, fn)
}

func (ge *GameEvents) OnGcMismatchedVersion(fn func(*GameEventGcMismatchedVersion) error) {
	ge.onGcMismatchedVersion = append(ge.onGcMismatchedVersion, fn)
}

func (ge *GameEvents) OnDemoStop(fn func(*GameEventDemoStop) error) {
	ge.onDemoStop = append(ge.onDemoStop, fn)
}

func (ge *GameEvents) OnMapShutdown(fn func(*GameEventMapShutdown) error) {
	ge.onMapShutdown = append(ge.onMapShutdown, fn)
}

func (ge *GameEvents) OnDotaWorkshopFileselected(fn func(*GameEventDotaWorkshopFileselected) error) {
	ge.onDotaWorkshopFileselected = append(ge.onDotaWorkshopFileselected, fn)
}

func (ge *GameEvents) OnDotaWorkshopFilecanceled(fn func(*GameEventDotaWorkshopFilecanceled) error) {
	ge.onDotaWorkshopFilecanceled = append(ge.onDotaWorkshopFilecanceled, fn)
}

func (ge *GameEvents) OnRichPresenceUpdated(fn func(*GameEventRichPresenceUpdated) error) {
	ge.onRichPresenceUpdated = append(ge.onRichPresenceUpdated, fn)
}

func (ge *GameEvents) OnDotaHeroRandom(fn func(*GameEventDotaHeroRandom) error) {
	ge.onDotaHeroRandom = append(ge.onDotaHeroRandom, fn)
}

func (ge *GameEvents) OnDotaRdChatTurn(fn func(*GameEventDotaRdChatTurn) error) {
	ge.onDotaRdChatTurn = append(ge.onDotaRdChatTurn, fn)
}

func (ge *GameEvents) OnDotaFavoriteHeroesUpdated(fn func(*GameEventDotaFavoriteHeroesUpdated) error) {
	ge.onDotaFavoriteHeroesUpdated = append(ge.onDotaFavoriteHeroesUpdated, fn)
}

func (ge *GameEvents) OnProfileOpened(fn func(*GameEventProfileOpened) error) {
	ge.onProfileOpened = append(ge.onProfileOpened, fn)
}

func (ge *GameEvents) OnProfileClosed(fn func(*GameEventProfileClosed) error) {
	ge.onProfileClosed = append(ge.onProfileClosed, fn)
}

func (ge *GameEvents) OnItemPreviewClosed(fn func(*GameEventItemPreviewClosed) error) {
	ge.onItemPreviewClosed = append(ge.onItemPreviewClosed, fn)
}

func (ge *GameEvents) OnDashboardSwitchedSection(fn func(*GameEventDashboardSwitchedSection) error) {
	ge.onDashboardSwitchedSection = append(ge.onDashboardSwitchedSection, fn)
}

func (ge *GameEvents) OnDotaTournamentItemEvent(fn func(*GameEventDotaTournamentItemEvent) error) {
	ge.onDotaTournamentItemEvent = append(ge.onDotaTournamentItemEvent, fn)
}

func (ge *GameEvents) OnDotaHeroSwap(fn func(*GameEventDotaHeroSwap) error) {
	ge.onDotaHeroSwap = append(ge.onDotaHeroSwap, fn)
}

func (ge *GameEvents) OnDotaResetSuggestedItems(fn func(*GameEventDotaResetSuggestedItems) error) {
	ge.onDotaResetSuggestedItems = append(ge.onDotaResetSuggestedItems, fn)
}

func (ge *GameEvents) OnHalloweenHighScoreReceived(fn func(*GameEventHalloweenHighScoreReceived) error) {
	ge.onHalloweenHighScoreReceived = append(ge.onHalloweenHighScoreReceived, fn)
}

func (ge *GameEvents) OnHalloweenPhaseEnd(fn func(*GameEventHalloweenPhaseEnd) error) {
	ge.onHalloweenPhaseEnd = append(ge.onHalloweenPhaseEnd, fn)
}

func (ge *GameEvents) OnHalloweenHighScoreRequestFailed(fn func(*GameEventHalloweenHighScoreRequestFailed) error) {
	ge.onHalloweenHighScoreRequestFailed = append(ge.onHalloweenHighScoreRequestFailed, fn)
}

func (ge *GameEvents) OnDotaHudSkinChanged(fn func(*GameEventDotaHudSkinChanged) error) {
	ge.onDotaHudSkinChanged = append(ge.onDotaHudSkinChanged, fn)
}

func (ge *GameEvents) OnDotaInventoryPlayerGotItem(fn func(*GameEventDotaInventoryPlayerGotItem) error) {
	ge.onDotaInventoryPlayerGotItem = append(ge.onDotaInventoryPlayerGotItem, fn)
}

func (ge *GameEvents) OnPlayerIsExperienced(fn func(*GameEventPlayerIsExperienced) error) {
	ge.onPlayerIsExperienced = append(ge.onPlayerIsExperienced, fn)
}

func (ge *GameEvents) OnPlayerIsNotexperienced(fn func(*GameEventPlayerIsNotexperienced) error) {
	ge.onPlayerIsNotexperienced = append(ge.onPlayerIsNotexperienced, fn)
}

func (ge *GameEvents) OnDotaTutorialLessonStart(fn func(*GameEventDotaTutorialLessonStart) error) {
	ge.onDotaTutorialLessonStart = append(ge.onDotaTutorialLessonStart, fn)
}

func (ge *GameEvents) OnDotaTutorialTaskAdvance(fn func(*GameEventDotaTutorialTaskAdvance) error) {
	ge.onDotaTutorialTaskAdvance = append(ge.onDotaTutorialTaskAdvance, fn)
}

func (ge *GameEvents) OnDotaTutorialShopToggled(fn func(*GameEventDotaTutorialShopToggled) error) {
	ge.onDotaTutorialShopToggled = append(ge.onDotaTutorialShopToggled, fn)
}

func (ge *GameEvents) OnMapLocationUpdated(fn func(*GameEventMapLocationUpdated) error) {
	ge.onMapLocationUpdated = append(ge.onMapLocationUpdated, fn)
}

func (ge *GameEvents) OnRichpresenceCustomUpdated(fn func(*GameEventRichpresenceCustomUpdated) error) {
	ge.onRichpresenceCustomUpdated = append(ge.onRichpresenceCustomUpdated, fn)
}

func (ge *GameEvents) OnGameEndVisible(fn func(*GameEventGameEndVisible) error) {
	ge.onGameEndVisible = append(ge.onGameEndVisible, fn)
}

func (ge *GameEvents) OnAntiaddictionUpdate(fn func(*GameEventAntiaddictionUpdate) error) {
	ge.onAntiaddictionUpdate = append(ge.onAntiaddictionUpdate, fn)
}

func (ge *GameEvents) OnHighlightHudElement(fn func(*GameEventHighlightHudElement) error) {
	ge.onHighlightHudElement = append(ge.onHighlightHudElement, fn)
}

func (ge *GameEvents) OnHideHighlightHudElement(fn func(*GameEventHideHighlightHudElement) error) {
	ge.onHideHighlightHudElement = append(ge.onHideHighlightHudElement, fn)
}

func (ge *GameEvents) OnIntroVideoFinished(fn func(*GameEventIntroVideoFinished) error) {
	ge.onIntroVideoFinished = append(ge.onIntroVideoFinished, fn)
}

func (ge *GameEvents) OnMatchmakingStatusVisibilityChanged(fn func(*GameEventMatchmakingStatusVisibilityChanged) error) {
	ge.onMatchmakingStatusVisibilityChanged = append(ge.onMatchmakingStatusVisibilityChanged, fn)
}

func (ge *GameEvents) OnPracticeLobbyVisibilityChanged(fn func(*GameEventPracticeLobbyVisibilityChanged) error) {
	ge.onPracticeLobbyVisibilityChanged = append(ge.onPracticeLobbyVisibilityChanged, fn)
}

func (ge *GameEvents) OnDotaCourierTransferItem(fn func(*GameEventDotaCourierTransferItem) error) {
	ge.onDotaCourierTransferItem = append(ge.onDotaCourierTransferItem, fn)
}

func (ge *GameEvents) OnFullUiUnlocked(fn func(*GameEventFullUiUnlocked) error) {
	ge.onFullUiUnlocked = append(ge.onFullUiUnlocked, fn)
}

func (ge *GameEvents) OnHeroSelectorPreviewSet(fn func(*GameEventHeroSelectorPreviewSet) error) {
	ge.onHeroSelectorPreviewSet = append(ge.onHeroSelectorPreviewSet, fn)
}

func (ge *GameEvents) OnAntiaddictionToast(fn func(*GameEventAntiaddictionToast) error) {
	ge.onAntiaddictionToast = append(ge.onAntiaddictionToast, fn)
}

func (ge *GameEvents) OnHeroPickerShown(fn func(*GameEventHeroPickerShown) error) {
	ge.onHeroPickerShown = append(ge.onHeroPickerShown, fn)
}

func (ge *GameEvents) OnHeroPickerHidden(fn func(*GameEventHeroPickerHidden) error) {
	ge.onHeroPickerHidden = append(ge.onHeroPickerHidden, fn)
}

func (ge *GameEvents) OnDotaLocalQuickbuyChanged(fn func(*GameEventDotaLocalQuickbuyChanged) error) {
	ge.onDotaLocalQuickbuyChanged = append(ge.onDotaLocalQuickbuyChanged, fn)
}

func (ge *GameEvents) OnShowCenterMessage(fn func(*GameEventShowCenterMessage) error) {
	ge.onShowCenterMessage = append(ge.onShowCenterMessage, fn)
}

func (ge *GameEvents) OnHudFlipChanged(fn func(*GameEventHudFlipChanged) error) {
	ge.onHudFlipChanged = append(ge.onHudFlipChanged, fn)
}

func (ge *GameEvents) OnFrostyPointsUpdated(fn func(*GameEventFrostyPointsUpdated) error) {
	ge.onFrostyPointsUpdated = append(ge.onFrostyPointsUpdated, fn)
}

func (ge *GameEvents) OnDefeated(fn func(*GameEventDefeated) error) {
	ge.onDefeated = append(ge.onDefeated, fn)
}

func (ge *GameEvents) OnResetDefeated(fn func(*GameEventResetDefeated) error) {
	ge.onResetDefeated = append(ge.onResetDefeated, fn)
}

func (ge *GameEvents) OnBoosterStateUpdated(fn func(*GameEventBoosterStateUpdated) error) {
	ge.onBoosterStateUpdated = append(ge.onBoosterStateUpdated, fn)
}

func (ge *GameEvents) OnEventPointsUpdated(fn func(*GameEventEventPointsUpdated) error) {
	ge.onEventPointsUpdated = append(ge.onEventPointsUpdated, fn)
}

func (ge *GameEvents) OnLocalPlayerEventPoints(fn func(*GameEventLocalPlayerEventPoints) error) {
	ge.onLocalPlayerEventPoints = append(ge.onLocalPlayerEventPoints, fn)
}

func (ge *GameEvents) OnCustomGameDifficulty(fn func(*GameEventCustomGameDifficulty) error) {
	ge.onCustomGameDifficulty = append(ge.onCustomGameDifficulty, fn)
}

func (ge *GameEvents) OnTreeCut(fn func(*GameEventTreeCut) error) {
	ge.onTreeCut = append(ge.onTreeCut, fn)
}

func (ge *GameEvents) OnUgcDetailsArrived(fn func(*GameEventUgcDetailsArrived) error) {
	ge.onUgcDetailsArrived = append(ge.onUgcDetailsArrived, fn)
}

func (ge *GameEvents) OnUgcSubscribed(fn func(*GameEventUgcSubscribed) error) {
	ge.onUgcSubscribed = append(ge.onUgcSubscribed, fn)
}

func (ge *GameEvents) OnUgcUnsubscribed(fn func(*GameEventUgcUnsubscribed) error) {
	ge.onUgcUnsubscribed = append(ge.onUgcUnsubscribed, fn)
}

func (ge *GameEvents) OnUgcDownloadRequested(fn func(*GameEventUgcDownloadRequested) error) {
	ge.onUgcDownloadRequested = append(ge.onUgcDownloadRequested, fn)
}

func (ge *GameEvents) OnUgcInstalled(fn func(*GameEventUgcInstalled) error) {
	ge.onUgcInstalled = append(ge.onUgcInstalled, fn)
}

func (ge *GameEvents) OnPrizepoolReceived(fn func(*GameEventPrizepoolReceived) error) {
	ge.onPrizepoolReceived = append(ge.onPrizepoolReceived, fn)
}

func (ge *GameEvents) OnMicrotransactionSuccess(fn func(*GameEventMicrotransactionSuccess) error) {
	ge.onMicrotransactionSuccess = append(ge.onMicrotransactionSuccess, fn)
}

func (ge *GameEvents) OnDotaRubickAbilitySteal(fn func(*GameEventDotaRubickAbilitySteal) error) {
	ge.onDotaRubickAbilitySteal = append(ge.onDotaRubickAbilitySteal, fn)
}

func (ge *GameEvents) OnCompendiumEventActionsLoaded(fn func(*GameEventCompendiumEventActionsLoaded) error) {
	ge.onCompendiumEventActionsLoaded = append(ge.onCompendiumEventActionsLoaded, fn)
}

func (ge *GameEvents) OnCompendiumSelectionsLoaded(fn func(*GameEventCompendiumSelectionsLoaded) error) {
	ge.onCompendiumSelectionsLoaded = append(ge.onCompendiumSelectionsLoaded, fn)
}

func (ge *GameEvents) OnCompendiumSetSelectionFailed(fn func(*GameEventCompendiumSetSelectionFailed) error) {
	ge.onCompendiumSetSelectionFailed = append(ge.onCompendiumSetSelectionFailed, fn)
}

func (ge *GameEvents) OnCompendiumTrophiesLoaded(fn func(*GameEventCompendiumTrophiesLoaded) error) {
	ge.onCompendiumTrophiesLoaded = append(ge.onCompendiumTrophiesLoaded, fn)
}

func (ge *GameEvents) OnCommunityCachedNamesUpdated(fn func(*GameEventCommunityCachedNamesUpdated) error) {
	ge.onCommunityCachedNamesUpdated = append(ge.onCommunityCachedNamesUpdated, fn)
}

func (ge *GameEvents) OnSpecItemPickup(fn func(*GameEventSpecItemPickup) error) {
	ge.onSpecItemPickup = append(ge.onSpecItemPickup, fn)
}

func (ge *GameEvents) OnSpecAegisReclaimTime(fn func(*GameEventSpecAegisReclaimTime) error) {
	ge.onSpecAegisReclaimTime = append(ge.onSpecAegisReclaimTime, fn)
}

func (ge *GameEvents) OnAccountTrophiesChanged(fn func(*GameEventAccountTrophiesChanged) error) {
	ge.onAccountTrophiesChanged = append(ge.onAccountTrophiesChanged, fn)
}

func (ge *GameEvents) OnAccountAllHeroChallengeChanged(fn func(*GameEventAccountAllHeroChallengeChanged) error) {
	ge.onAccountAllHeroChallengeChanged = append(ge.onAccountAllHeroChallengeChanged, fn)
}

func (ge *GameEvents) OnTeamShowcaseUiUpdate(fn func(*GameEventTeamShowcaseUiUpdate) error) {
	ge.onTeamShowcaseUiUpdate = append(ge.onTeamShowcaseUiUpdate, fn)
}

func (ge *GameEvents) OnIngameEventsChanged(fn func(*GameEventIngameEventsChanged) error) {
	ge.onIngameEventsChanged = append(ge.onIngameEventsChanged, fn)
}

func (ge *GameEvents) OnDotaMatchSignout(fn func(*GameEventDotaMatchSignout) error) {
	ge.onDotaMatchSignout = append(ge.onDotaMatchSignout, fn)
}

func (ge *GameEvents) OnDotaIllusionsCreated(fn func(*GameEventDotaIllusionsCreated) error) {
	ge.onDotaIllusionsCreated = append(ge.onDotaIllusionsCreated, fn)
}

func (ge *GameEvents) OnDotaYearBeastKilled(fn func(*GameEventDotaYearBeastKilled) error) {
	ge.onDotaYearBeastKilled = append(ge.onDotaYearBeastKilled, fn)
}

func (ge *GameEvents) OnDotaHeroUndoselection(fn func(*GameEventDotaHeroUndoselection) error) {
	ge.onDotaHeroUndoselection = append(ge.onDotaHeroUndoselection, fn)
}

func (ge *GameEvents) OnDotaChallengeSocacheUpdated(fn func(*GameEventDotaChallengeSocacheUpdated) error) {
	ge.onDotaChallengeSocacheUpdated = append(ge.onDotaChallengeSocacheUpdated, fn)
}

func (ge *GameEvents) OnPartyInvitesUpdated(fn func(*GameEventPartyInvitesUpdated) error) {
	ge.onPartyInvitesUpdated = append(ge.onPartyInvitesUpdated, fn)
}

func (ge *GameEvents) OnLobbyInvitesUpdated(fn func(*GameEventLobbyInvitesUpdated) error) {
	ge.onLobbyInvitesUpdated = append(ge.onLobbyInvitesUpdated, fn)
}

func (ge *GameEvents) OnCustomGameModeListUpdated(fn func(*GameEventCustomGameModeListUpdated) error) {
	ge.onCustomGameModeListUpdated = append(ge.onCustomGameModeListUpdated, fn)
}

func (ge *GameEvents) OnCustomGameLobbyListUpdated(fn func(*GameEventCustomGameLobbyListUpdated) error) {
	ge.onCustomGameLobbyListUpdated = append(ge.onCustomGameLobbyListUpdated, fn)
}

func (ge *GameEvents) OnFriendLobbyListUpdated(fn func(*GameEventFriendLobbyListUpdated) error) {
	ge.onFriendLobbyListUpdated = append(ge.onFriendLobbyListUpdated, fn)
}

func (ge *GameEvents) OnDotaTeamPlayerListChanged(fn func(*GameEventDotaTeamPlayerListChanged) error) {
	ge.onDotaTeamPlayerListChanged = append(ge.onDotaTeamPlayerListChanged, fn)
}

func (ge *GameEvents) OnDotaPlayerDetailsChanged(fn func(*GameEventDotaPlayerDetailsChanged) error) {
	ge.onDotaPlayerDetailsChanged = append(ge.onDotaPlayerDetailsChanged, fn)
}

func (ge *GameEvents) OnPlayerProfileStatsUpdated(fn func(*GameEventPlayerProfileStatsUpdated) error) {
	ge.onPlayerProfileStatsUpdated = append(ge.onPlayerProfileStatsUpdated, fn)
}

func (ge *GameEvents) OnCustomGamePlayerCountUpdated(fn func(*GameEventCustomGamePlayerCountUpdated) error) {
	ge.onCustomGamePlayerCountUpdated = append(ge.onCustomGamePlayerCountUpdated, fn)
}

func (ge *GameEvents) OnCustomGameFriendsPlayedUpdated(fn func(*GameEventCustomGameFriendsPlayedUpdated) error) {
	ge.onCustomGameFriendsPlayedUpdated = append(ge.onCustomGameFriendsPlayedUpdated, fn)
}

func (ge *GameEvents) OnCustomGamesFriendsPlayUpdated(fn func(*GameEventCustomGamesFriendsPlayUpdated) error) {
	ge.onCustomGamesFriendsPlayUpdated = append(ge.onCustomGamesFriendsPlayUpdated, fn)
}

func (ge *GameEvents) OnDotaPlayerUpdateAssignedHero(fn func(*GameEventDotaPlayerUpdateAssignedHero) error) {
	ge.onDotaPlayerUpdateAssignedHero = append(ge.onDotaPlayerUpdateAssignedHero, fn)
}

func (ge *GameEvents) OnDotaPlayerHeroSelectionDirty(fn func(*GameEventDotaPlayerHeroSelectionDirty) error) {
	ge.onDotaPlayerHeroSelectionDirty = append(ge.onDotaPlayerHeroSelectionDirty, fn)
}

func (ge *GameEvents) OnDotaNpcGoalReached(fn func(*GameEventDotaNpcGoalReached) error) {
	ge.onDotaNpcGoalReached = append(ge.onDotaNpcGoalReached, fn)
}

func (ge *GameEvents) OnDotaPlayerSelectedCustomTeam(fn func(*GameEventDotaPlayerSelectedCustomTeam) error) {
	ge.onDotaPlayerSelectedCustomTeam = append(ge.onDotaPlayerSelectedCustomTeam, fn)
}

func (ge *GameEvents) OnHltvStatus(fn func(*GameEventHltvStatus) error) {
	ge.onHltvStatus = append(ge.onHltvStatus, fn)
}

func (ge *GameEvents) OnHltvCameraman(fn func(*GameEventHltvCameraman) error) {
	ge.onHltvCameraman = append(ge.onHltvCameraman, fn)
}

func (ge *GameEvents) OnHltvRankCamera(fn func(*GameEventHltvRankCamera) error) {
	ge.onHltvRankCamera = append(ge.onHltvRankCamera, fn)
}

func (ge *GameEvents) OnHltvRankEntity(fn func(*GameEventHltvRankEntity) error) {
	ge.onHltvRankEntity = append(ge.onHltvRankEntity, fn)
}

func (ge *GameEvents) OnHltvFixed(fn func(*GameEventHltvFixed) error) {
	ge.onHltvFixed = append(ge.onHltvFixed, fn)
}

func (ge *GameEvents) OnHltvChase(fn func(*GameEventHltvChase) error) {
	ge.onHltvChase = append(ge.onHltvChase, fn)
}

func (ge *GameEvents) OnHltvMessage(fn func(*GameEventHltvMessage) error) {
	ge.onHltvMessage = append(ge.onHltvMessage, fn)
}

func (ge *GameEvents) OnHltvTitle(fn func(*GameEventHltvTitle) error) {
	ge.onHltvTitle = append(ge.onHltvTitle, fn)
}

func (ge *GameEvents) OnHltvChat(fn func(*GameEventHltvChat) error) {
	ge.onHltvChat = append(ge.onHltvChat, fn)
}

func (ge *GameEvents) OnHltvVersioninfo(fn func(*GameEventHltvVersioninfo) error) {
	ge.onHltvVersioninfo = append(ge.onHltvVersioninfo, fn)
}

func (ge *GameEvents) OnDotaChaseHero(fn func(*GameEventDotaChaseHero) error) {
	ge.onDotaChaseHero = append(ge.onDotaChaseHero, fn)
}

func (ge *GameEvents) OnDotaCombatlog(fn func(*GameEventDotaCombatlog) error) {
	ge.onDotaCombatlog = append(ge.onDotaCombatlog, fn)
}

func (ge *GameEvents) OnDotaGameStateChange(fn func(*GameEventDotaGameStateChange) error) {
	ge.onDotaGameStateChange = append(ge.onDotaGameStateChange, fn)
}

func (ge *GameEvents) OnDotaPlayerPickHero(fn func(*GameEventDotaPlayerPickHero) error) {
	ge.onDotaPlayerPickHero = append(ge.onDotaPlayerPickHero, fn)
}

func (ge *GameEvents) OnDotaTeamKillCredit(fn func(*GameEventDotaTeamKillCredit) error) {
	ge.onDotaTeamKillCredit = append(ge.onDotaTeamKillCredit, fn)
}

func (ge *GameEvents) onCMsgSource1LegacyGameEvent(m *dota.CMsgSource1LegacyGameEvent) error {
	switch m.GetEventid() {

	case 0: // EGameEvent_ServerSpawn
		if len(m.GetKeys()) < 10 {
			_debugf("short EGameEvent_ServerSpawn: %v", m)
			return nil
		}
		if cbs := ge.onServerSpawn; cbs != nil {
			msg := &GameEventServerSpawn{}
			msg.Hostname = m.GetKeys()[0].GetValString()
			msg.Address = m.GetKeys()[1].GetValString()
			msg.Port = m.GetKeys()[2].GetValShort()
			msg.Game = m.GetKeys()[3].GetValString()
			msg.Mapname = m.GetKeys()[4].GetValString()
			msg.Addonname = m.GetKeys()[5].GetValString()
			msg.Maxplayers = m.GetKeys()[6].GetValLong()
			msg.Os = m.GetKeys()[7].GetValString()
			msg.Dedicated = m.GetKeys()[8].GetValBool()
			msg.Password = m.GetKeys()[9].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 1: // EGameEvent_ServerPreShutdown
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_ServerPreShutdown: %v", m)
			return nil
		}
		if cbs := ge.onServerPreShutdown; cbs != nil {
			msg := &GameEventServerPreShutdown{}
			msg.Reason = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 2: // EGameEvent_ServerShutdown
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_ServerShutdown: %v", m)
			return nil
		}
		if cbs := ge.onServerShutdown; cbs != nil {
			msg := &GameEventServerShutdown{}
			msg.Reason = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 3: // EGameEvent_ServerCvar
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_ServerCvar: %v", m)
			return nil
		}
		if cbs := ge.onServerCvar; cbs != nil {
			msg := &GameEventServerCvar{}
			msg.Cvarname = m.GetKeys()[0].GetValString()
			msg.Cvarvalue = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 4: // EGameEvent_ServerMessage
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_ServerMessage: %v", m)
			return nil
		}
		if cbs := ge.onServerMessage; cbs != nil {
			msg := &GameEventServerMessage{}
			msg.Text = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 5: // EGameEvent_ServerAddban
		if len(m.GetKeys()) < 7 {
			_debugf("short EGameEvent_ServerAddban: %v", m)
			return nil
		}
		if cbs := ge.onServerAddban; cbs != nil {
			msg := &GameEventServerAddban{}
			msg.Name = m.GetKeys()[0].GetValString()
			msg.Userid = m.GetKeys()[1].GetValShort()
			msg.Networkid = m.GetKeys()[2].GetValString()
			msg.Ip = m.GetKeys()[3].GetValString()
			msg.Duration = m.GetKeys()[4].GetValString()
			msg.By = m.GetKeys()[5].GetValString()
			msg.Kicked = m.GetKeys()[6].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 6: // EGameEvent_ServerRemoveban
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_ServerRemoveban: %v", m)
			return nil
		}
		if cbs := ge.onServerRemoveban; cbs != nil {
			msg := &GameEventServerRemoveban{}
			msg.Networkid = m.GetKeys()[0].GetValString()
			msg.Ip = m.GetKeys()[1].GetValString()
			msg.By = m.GetKeys()[2].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 7: // EGameEvent_PlayerConnect
		if len(m.GetKeys()) < 5 {
			_debugf("short EGameEvent_PlayerConnect: %v", m)
			return nil
		}
		if cbs := ge.onPlayerConnect; cbs != nil {
			msg := &GameEventPlayerConnect{}
			msg.Name = m.GetKeys()[0].GetValString()
			msg.Index = m.GetKeys()[1].GetValByte()
			msg.Userid = m.GetKeys()[2].GetValShort()
			msg.Networkid = m.GetKeys()[3].GetValString()
			msg.Address = m.GetKeys()[4].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 8: // EGameEvent_PlayerInfo
		if len(m.GetKeys()) < 5 {
			_debugf("short EGameEvent_PlayerInfo: %v", m)
			return nil
		}
		if cbs := ge.onPlayerInfo; cbs != nil {
			msg := &GameEventPlayerInfo{}
			msg.Name = m.GetKeys()[0].GetValString()
			msg.Index = m.GetKeys()[1].GetValByte()
			msg.Userid = m.GetKeys()[2].GetValShort()
			msg.Networkid = m.GetKeys()[3].GetValString()
			msg.Bot = m.GetKeys()[4].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 9: // EGameEvent_PlayerDisconnect
		if len(m.GetKeys()) < 4 {
			_debugf("short EGameEvent_PlayerDisconnect: %v", m)
			return nil
		}
		if cbs := ge.onPlayerDisconnect; cbs != nil {
			msg := &GameEventPlayerDisconnect{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Reason = m.GetKeys()[1].GetValShort()
			msg.Name = m.GetKeys()[2].GetValString()
			msg.Networkid = m.GetKeys()[3].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 10: // EGameEvent_PlayerActivate
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_PlayerActivate: %v", m)
			return nil
		}
		if cbs := ge.onPlayerActivate; cbs != nil {
			msg := &GameEventPlayerActivate{}
			msg.Userid = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 11: // EGameEvent_PlayerConnectFull
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_PlayerConnectFull: %v", m)
			return nil
		}
		if cbs := ge.onPlayerConnectFull; cbs != nil {
			msg := &GameEventPlayerConnectFull{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Index = m.GetKeys()[1].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 12: // EGameEvent_PlayerSay
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_PlayerSay: %v", m)
			return nil
		}
		if cbs := ge.onPlayerSay; cbs != nil {
			msg := &GameEventPlayerSay{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Text = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 13: // EGameEvent_PlayerFullUpdate
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_PlayerFullUpdate: %v", m)
			return nil
		}
		if cbs := ge.onPlayerFullUpdate; cbs != nil {
			msg := &GameEventPlayerFullUpdate{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Count = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 14: // EGameEvent_TeamInfo
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_TeamInfo: %v", m)
			return nil
		}
		if cbs := ge.onTeamInfo; cbs != nil {
			msg := &GameEventTeamInfo{}
			msg.Teamid = m.GetKeys()[0].GetValByte()
			msg.Teamname = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 15: // EGameEvent_TeamScore
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_TeamScore: %v", m)
			return nil
		}
		if cbs := ge.onTeamScore; cbs != nil {
			msg := &GameEventTeamScore{}
			msg.Teamid = m.GetKeys()[0].GetValByte()
			msg.Score = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 16: // EGameEvent_TeamplayBroadcastAudio
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_TeamplayBroadcastAudio: %v", m)
			return nil
		}
		if cbs := ge.onTeamplayBroadcastAudio; cbs != nil {
			msg := &GameEventTeamplayBroadcastAudio{}
			msg.Team = m.GetKeys()[0].GetValByte()
			msg.Sound = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 17: // EGameEvent_PlayerTeam
		if len(m.GetKeys()) < 6 {
			_debugf("short EGameEvent_PlayerTeam: %v", m)
			return nil
		}
		if cbs := ge.onPlayerTeam; cbs != nil {
			msg := &GameEventPlayerTeam{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Team = m.GetKeys()[1].GetValByte()
			msg.Oldteam = m.GetKeys()[2].GetValByte()
			msg.Disconnect = m.GetKeys()[3].GetValBool()
			msg.Autoteam = m.GetKeys()[4].GetValBool()
			msg.Silent = m.GetKeys()[5].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 18: // EGameEvent_PlayerClass
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_PlayerClass: %v", m)
			return nil
		}
		if cbs := ge.onPlayerClass; cbs != nil {
			msg := &GameEventPlayerClass{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Class = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 19: // EGameEvent_PlayerDeath
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_PlayerDeath: %v", m)
			return nil
		}
		if cbs := ge.onPlayerDeath; cbs != nil {
			msg := &GameEventPlayerDeath{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Attacker = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 20: // EGameEvent_PlayerHurt
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_PlayerHurt: %v", m)
			return nil
		}
		if cbs := ge.onPlayerHurt; cbs != nil {
			msg := &GameEventPlayerHurt{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Attacker = m.GetKeys()[1].GetValShort()
			msg.Health = m.GetKeys()[2].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 21: // EGameEvent_PlayerChat
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_PlayerChat: %v", m)
			return nil
		}
		if cbs := ge.onPlayerChat; cbs != nil {
			msg := &GameEventPlayerChat{}
			msg.Teamonly = m.GetKeys()[0].GetValBool()
			msg.Userid = m.GetKeys()[1].GetValShort()
			msg.Text = m.GetKeys()[2].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 22: // EGameEvent_PlayerScore
		if len(m.GetKeys()) < 4 {
			_debugf("short EGameEvent_PlayerScore: %v", m)
			return nil
		}
		if cbs := ge.onPlayerScore; cbs != nil {
			msg := &GameEventPlayerScore{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Kills = m.GetKeys()[1].GetValShort()
			msg.Deaths = m.GetKeys()[2].GetValShort()
			msg.Score = m.GetKeys()[3].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 23: // EGameEvent_PlayerSpawn
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_PlayerSpawn: %v", m)
			return nil
		}
		if cbs := ge.onPlayerSpawn; cbs != nil {
			msg := &GameEventPlayerSpawn{}
			msg.Userid = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 24: // EGameEvent_PlayerShoot
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_PlayerShoot: %v", m)
			return nil
		}
		if cbs := ge.onPlayerShoot; cbs != nil {
			msg := &GameEventPlayerShoot{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Weapon = m.GetKeys()[1].GetValByte()
			msg.Mode = m.GetKeys()[2].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 25: // EGameEvent_PlayerUse
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_PlayerUse: %v", m)
			return nil
		}
		if cbs := ge.onPlayerUse; cbs != nil {
			msg := &GameEventPlayerUse{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Entity = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 26: // EGameEvent_PlayerChangename
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_PlayerChangename: %v", m)
			return nil
		}
		if cbs := ge.onPlayerChangename; cbs != nil {
			msg := &GameEventPlayerChangename{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Oldname = m.GetKeys()[1].GetValString()
			msg.Newname = m.GetKeys()[2].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 27: // EGameEvent_PlayerHintmessage
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_PlayerHintmessage: %v", m)
			return nil
		}
		if cbs := ge.onPlayerHintmessage; cbs != nil {
			msg := &GameEventPlayerHintmessage{}
			msg.Hintmessage = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 28: // EGameEvent_GameInit
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_GameInit: %v", m)
			return nil
		}
		if cbs := ge.onGameInit; cbs != nil {
			msg := &GameEventGameInit{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 29: // EGameEvent_GameNewmap
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_GameNewmap: %v", m)
			return nil
		}
		if cbs := ge.onGameNewmap; cbs != nil {
			msg := &GameEventGameNewmap{}
			msg.Mapname = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 30: // EGameEvent_GameStart
		if len(m.GetKeys()) < 4 {
			_debugf("short EGameEvent_GameStart: %v", m)
			return nil
		}
		if cbs := ge.onGameStart; cbs != nil {
			msg := &GameEventGameStart{}
			msg.Roundslimit = m.GetKeys()[0].GetValLong()
			msg.Timelimit = m.GetKeys()[1].GetValLong()
			msg.Fraglimit = m.GetKeys()[2].GetValLong()
			msg.Objective = m.GetKeys()[3].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 31: // EGameEvent_GameEnd
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_GameEnd: %v", m)
			return nil
		}
		if cbs := ge.onGameEnd; cbs != nil {
			msg := &GameEventGameEnd{}
			msg.Winner = m.GetKeys()[0].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 32: // EGameEvent_RoundStart
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_RoundStart: %v", m)
			return nil
		}
		if cbs := ge.onRoundStart; cbs != nil {
			msg := &GameEventRoundStart{}
			msg.Timelimit = m.GetKeys()[0].GetValLong()
			msg.Fraglimit = m.GetKeys()[1].GetValLong()
			msg.Objective = m.GetKeys()[2].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 33: // EGameEvent_RoundEnd
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_RoundEnd: %v", m)
			return nil
		}
		if cbs := ge.onRoundEnd; cbs != nil {
			msg := &GameEventRoundEnd{}
			msg.Winner = m.GetKeys()[0].GetValByte()
			msg.Reason = m.GetKeys()[1].GetValByte()
			msg.Message = m.GetKeys()[2].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 34: // EGameEvent_RoundStartPreEntity
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_RoundStartPreEntity: %v", m)
			return nil
		}
		if cbs := ge.onRoundStartPreEntity; cbs != nil {
			msg := &GameEventRoundStartPreEntity{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 35: // EGameEvent_TeamplayRoundStart
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_TeamplayRoundStart: %v", m)
			return nil
		}
		if cbs := ge.onTeamplayRoundStart; cbs != nil {
			msg := &GameEventTeamplayRoundStart{}
			msg.FullReset = m.GetKeys()[0].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 36: // EGameEvent_HostnameChanged
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_HostnameChanged: %v", m)
			return nil
		}
		if cbs := ge.onHostnameChanged; cbs != nil {
			msg := &GameEventHostnameChanged{}
			msg.Hostname = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 37: // EGameEvent_DifficultyChanged
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_DifficultyChanged: %v", m)
			return nil
		}
		if cbs := ge.onDifficultyChanged; cbs != nil {
			msg := &GameEventDifficultyChanged{}
			msg.NewDifficulty = m.GetKeys()[0].GetValShort()
			msg.OldDifficulty = m.GetKeys()[1].GetValShort()
			msg.StrDifficulty = m.GetKeys()[2].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 38: // EGameEvent_FinaleStart
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_FinaleStart: %v", m)
			return nil
		}
		if cbs := ge.onFinaleStart; cbs != nil {
			msg := &GameEventFinaleStart{}
			msg.Rushes = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 39: // EGameEvent_GameMessage
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_GameMessage: %v", m)
			return nil
		}
		if cbs := ge.onGameMessage; cbs != nil {
			msg := &GameEventGameMessage{}
			msg.Target = m.GetKeys()[0].GetValByte()
			msg.Text = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 40: // EGameEvent_BreakBreakable
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_BreakBreakable: %v", m)
			return nil
		}
		if cbs := ge.onBreakBreakable; cbs != nil {
			msg := &GameEventBreakBreakable{}
			msg.Entindex = m.GetKeys()[0].GetValLong()
			msg.Userid = m.GetKeys()[1].GetValShort()
			msg.Material = m.GetKeys()[2].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 41: // EGameEvent_BreakProp
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_BreakProp: %v", m)
			return nil
		}
		if cbs := ge.onBreakProp; cbs != nil {
			msg := &GameEventBreakProp{}
			msg.Entindex = m.GetKeys()[0].GetValLong()
			msg.Userid = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 42: // EGameEvent_NpcSpawned
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_NpcSpawned: %v", m)
			return nil
		}
		if cbs := ge.onNpcSpawned; cbs != nil {
			msg := &GameEventNpcSpawned{}
			msg.Entindex = m.GetKeys()[0].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 43: // EGameEvent_NpcReplaced
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_NpcReplaced: %v", m)
			return nil
		}
		if cbs := ge.onNpcReplaced; cbs != nil {
			msg := &GameEventNpcReplaced{}
			msg.OldEntindex = m.GetKeys()[0].GetValLong()
			msg.NewEntindex = m.GetKeys()[1].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 44: // EGameEvent_EntityKilled
		if len(m.GetKeys()) < 4 {
			_debugf("short EGameEvent_EntityKilled: %v", m)
			return nil
		}
		if cbs := ge.onEntityKilled; cbs != nil {
			msg := &GameEventEntityKilled{}
			msg.EntindexKilled = m.GetKeys()[0].GetValLong()
			msg.EntindexAttacker = m.GetKeys()[1].GetValLong()
			msg.EntindexInflictor = m.GetKeys()[2].GetValLong()
			msg.Damagebits = m.GetKeys()[3].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 45: // EGameEvent_EntityHurt
		if len(m.GetKeys()) < 4 {
			_debugf("short EGameEvent_EntityHurt: %v", m)
			return nil
		}
		if cbs := ge.onEntityHurt; cbs != nil {
			msg := &GameEventEntityHurt{}
			msg.EntindexKilled = m.GetKeys()[0].GetValLong()
			msg.EntindexAttacker = m.GetKeys()[1].GetValLong()
			msg.EntindexInflictor = m.GetKeys()[2].GetValLong()
			msg.Damagebits = m.GetKeys()[3].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 46: // EGameEvent_BonusUpdated
		if len(m.GetKeys()) < 4 {
			_debugf("short EGameEvent_BonusUpdated: %v", m)
			return nil
		}
		if cbs := ge.onBonusUpdated; cbs != nil {
			msg := &GameEventBonusUpdated{}
			msg.Numadvanced = m.GetKeys()[0].GetValShort()
			msg.Numbronze = m.GetKeys()[1].GetValShort()
			msg.Numsilver = m.GetKeys()[2].GetValShort()
			msg.Numgold = m.GetKeys()[3].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 47: // EGameEvent_PlayerStatsUpdated
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_PlayerStatsUpdated: %v", m)
			return nil
		}
		if cbs := ge.onPlayerStatsUpdated; cbs != nil {
			msg := &GameEventPlayerStatsUpdated{}
			msg.Forceupload = m.GetKeys()[0].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 48: // EGameEvent_AchievementEvent
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_AchievementEvent: %v", m)
			return nil
		}
		if cbs := ge.onAchievementEvent; cbs != nil {
			msg := &GameEventAchievementEvent{}
			msg.AchievementName = m.GetKeys()[0].GetValString()
			msg.CurVal = m.GetKeys()[1].GetValShort()
			msg.MaxVal = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 49: // EGameEvent_AchievementEarned
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_AchievementEarned: %v", m)
			return nil
		}
		if cbs := ge.onAchievementEarned; cbs != nil {
			msg := &GameEventAchievementEarned{}
			msg.Player = m.GetKeys()[0].GetValByte()
			msg.Achievement = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 50: // EGameEvent_AchievementWriteFailed
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_AchievementWriteFailed: %v", m)
			return nil
		}
		if cbs := ge.onAchievementWriteFailed; cbs != nil {
			msg := &GameEventAchievementWriteFailed{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 51: // EGameEvent_PhysgunPickup
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_PhysgunPickup: %v", m)
			return nil
		}
		if cbs := ge.onPhysgunPickup; cbs != nil {
			msg := &GameEventPhysgunPickup{}
			msg.Entindex = m.GetKeys()[0].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 52: // EGameEvent_FlareIgniteNpc
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_FlareIgniteNpc: %v", m)
			return nil
		}
		if cbs := ge.onFlareIgniteNpc; cbs != nil {
			msg := &GameEventFlareIgniteNpc{}
			msg.Entindex = m.GetKeys()[0].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 53: // EGameEvent_HelicopterGrenadePuntMiss
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_HelicopterGrenadePuntMiss: %v", m)
			return nil
		}
		if cbs := ge.onHelicopterGrenadePuntMiss; cbs != nil {
			msg := &GameEventHelicopterGrenadePuntMiss{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 54: // EGameEvent_UserDataDownloaded
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_UserDataDownloaded: %v", m)
			return nil
		}
		if cbs := ge.onUserDataDownloaded; cbs != nil {
			msg := &GameEventUserDataDownloaded{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 55: // EGameEvent_RagdollDissolved
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_RagdollDissolved: %v", m)
			return nil
		}
		if cbs := ge.onRagdollDissolved; cbs != nil {
			msg := &GameEventRagdollDissolved{}
			msg.Entindex = m.GetKeys()[0].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 56: // EGameEvent_GameinstructorDraw
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_GameinstructorDraw: %v", m)
			return nil
		}
		if cbs := ge.onGameinstructorDraw; cbs != nil {
			msg := &GameEventGameinstructorDraw{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 57: // EGameEvent_GameinstructorNodraw
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_GameinstructorNodraw: %v", m)
			return nil
		}
		if cbs := ge.onGameinstructorNodraw; cbs != nil {
			msg := &GameEventGameinstructorNodraw{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 58: // EGameEvent_MapTransition
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_MapTransition: %v", m)
			return nil
		}
		if cbs := ge.onMapTransition; cbs != nil {
			msg := &GameEventMapTransition{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 59: // EGameEvent_InstructorServerHintCreate
		if len(m.GetKeys()) < 18 {
			_debugf("short EGameEvent_InstructorServerHintCreate: %v", m)
			return nil
		}
		if cbs := ge.onInstructorServerHintCreate; cbs != nil {
			msg := &GameEventInstructorServerHintCreate{}
			msg.HintName = m.GetKeys()[0].GetValString()
			msg.HintReplaceKey = m.GetKeys()[1].GetValString()
			msg.HintTarget = m.GetKeys()[2].GetValLong()
			msg.HintActivatorUserid = m.GetKeys()[3].GetValShort()
			msg.HintTimeout = m.GetKeys()[4].GetValShort()
			msg.HintIconOnscreen = m.GetKeys()[5].GetValString()
			msg.HintIconOffscreen = m.GetKeys()[6].GetValString()
			msg.HintCaption = m.GetKeys()[7].GetValString()
			msg.HintActivatorCaption = m.GetKeys()[8].GetValString()
			msg.HintColor = m.GetKeys()[9].GetValString()
			msg.HintIconOffset = m.GetKeys()[10].GetValFloat()
			msg.HintRange = m.GetKeys()[11].GetValFloat()
			msg.HintFlags = m.GetKeys()[12].GetValLong()
			msg.HintBinding = m.GetKeys()[13].GetValString()
			msg.HintAllowNodrawTarget = m.GetKeys()[14].GetValBool()
			msg.HintNooffscreen = m.GetKeys()[15].GetValBool()
			msg.HintForcecaption = m.GetKeys()[16].GetValBool()
			msg.HintLocalPlayerOnly = m.GetKeys()[17].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 60: // EGameEvent_InstructorServerHintStop
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_InstructorServerHintStop: %v", m)
			return nil
		}
		if cbs := ge.onInstructorServerHintStop; cbs != nil {
			msg := &GameEventInstructorServerHintStop{}
			msg.HintName = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 61: // EGameEvent_ChatNewMessage
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_ChatNewMessage: %v", m)
			return nil
		}
		if cbs := ge.onChatNewMessage; cbs != nil {
			msg := &GameEventChatNewMessage{}
			msg.Channel = m.GetKeys()[0].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 62: // EGameEvent_ChatMembersChanged
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_ChatMembersChanged: %v", m)
			return nil
		}
		if cbs := ge.onChatMembersChanged; cbs != nil {
			msg := &GameEventChatMembersChanged{}
			msg.Channel = m.GetKeys()[0].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 63: // EGameEvent_InventoryUpdated
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_InventoryUpdated: %v", m)
			return nil
		}
		if cbs := ge.onInventoryUpdated; cbs != nil {
			msg := &GameEventInventoryUpdated{}
			msg.Itemdef = m.GetKeys()[0].GetValShort()
			msg.Itemid = m.GetKeys()[1].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 64: // EGameEvent_CartUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_CartUpdated: %v", m)
			return nil
		}
		if cbs := ge.onCartUpdated; cbs != nil {
			msg := &GameEventCartUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 65: // EGameEvent_StorePricesheetUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_StorePricesheetUpdated: %v", m)
			return nil
		}
		if cbs := ge.onStorePricesheetUpdated; cbs != nil {
			msg := &GameEventStorePricesheetUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 66: // EGameEvent_GcConnected
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_GcConnected: %v", m)
			return nil
		}
		if cbs := ge.onGcConnected; cbs != nil {
			msg := &GameEventGcConnected{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 67: // EGameEvent_ItemSchemaInitialized
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_ItemSchemaInitialized: %v", m)
			return nil
		}
		if cbs := ge.onItemSchemaInitialized; cbs != nil {
			msg := &GameEventItemSchemaInitialized{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 68: // EGameEvent_DropRateModified
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DropRateModified: %v", m)
			return nil
		}
		if cbs := ge.onDropRateModified; cbs != nil {
			msg := &GameEventDropRateModified{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 69: // EGameEvent_EventTicketModified
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_EventTicketModified: %v", m)
			return nil
		}
		if cbs := ge.onEventTicketModified; cbs != nil {
			msg := &GameEventEventTicketModified{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 70: // EGameEvent_ModifierEvent
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_ModifierEvent: %v", m)
			return nil
		}
		if cbs := ge.onModifierEvent; cbs != nil {
			msg := &GameEventModifierEvent{}
			msg.Eventname = m.GetKeys()[0].GetValString()
			msg.Caster = m.GetKeys()[1].GetValShort()
			msg.Ability = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 71: // EGameEvent_DotaPlayerKill
		if len(m.GetKeys()) < 9 {
			_debugf("short EGameEvent_DotaPlayerKill: %v", m)
			return nil
		}
		if cbs := ge.onDotaPlayerKill; cbs != nil {
			msg := &GameEventDotaPlayerKill{}
			msg.VictimUserid = m.GetKeys()[0].GetValShort()
			msg.Killer1Userid = m.GetKeys()[1].GetValShort()
			msg.Killer2Userid = m.GetKeys()[2].GetValShort()
			msg.Killer3Userid = m.GetKeys()[3].GetValShort()
			msg.Killer4Userid = m.GetKeys()[4].GetValShort()
			msg.Killer5Userid = m.GetKeys()[5].GetValShort()
			msg.Bounty = m.GetKeys()[6].GetValShort()
			msg.Neutral = m.GetKeys()[7].GetValShort()
			msg.Greevil = m.GetKeys()[8].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 72: // EGameEvent_DotaPlayerDeny
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaPlayerDeny: %v", m)
			return nil
		}
		if cbs := ge.onDotaPlayerDeny; cbs != nil {
			msg := &GameEventDotaPlayerDeny{}
			msg.KillerUserid = m.GetKeys()[0].GetValShort()
			msg.VictimUserid = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 73: // EGameEvent_DotaBarracksKill
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaBarracksKill: %v", m)
			return nil
		}
		if cbs := ge.onDotaBarracksKill; cbs != nil {
			msg := &GameEventDotaBarracksKill{}
			msg.BarracksId = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 74: // EGameEvent_DotaTowerKill
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_DotaTowerKill: %v", m)
			return nil
		}
		if cbs := ge.onDotaTowerKill; cbs != nil {
			msg := &GameEventDotaTowerKill{}
			msg.KillerUserid = m.GetKeys()[0].GetValShort()
			msg.Teamnumber = m.GetKeys()[1].GetValShort()
			msg.Gold = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 75: // EGameEvent_DotaEffigyKill
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaEffigyKill: %v", m)
			return nil
		}
		if cbs := ge.onDotaEffigyKill; cbs != nil {
			msg := &GameEventDotaEffigyKill{}
			msg.OwnerUserid = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 76: // EGameEvent_DotaRoshanKill
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaRoshanKill: %v", m)
			return nil
		}
		if cbs := ge.onDotaRoshanKill; cbs != nil {
			msg := &GameEventDotaRoshanKill{}
			msg.Teamnumber = m.GetKeys()[0].GetValShort()
			msg.Gold = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 77: // EGameEvent_DotaCourierLost
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaCourierLost: %v", m)
			return nil
		}
		if cbs := ge.onDotaCourierLost; cbs != nil {
			msg := &GameEventDotaCourierLost{}
			msg.Teamnumber = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 78: // EGameEvent_DotaCourierRespawned
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaCourierRespawned: %v", m)
			return nil
		}
		if cbs := ge.onDotaCourierRespawned; cbs != nil {
			msg := &GameEventDotaCourierRespawned{}
			msg.Teamnumber = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 79: // EGameEvent_DotaGlyphUsed
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaGlyphUsed: %v", m)
			return nil
		}
		if cbs := ge.onDotaGlyphUsed; cbs != nil {
			msg := &GameEventDotaGlyphUsed{}
			msg.Teamnumber = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 80: // EGameEvent_DotaSuperCreeps
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaSuperCreeps: %v", m)
			return nil
		}
		if cbs := ge.onDotaSuperCreeps; cbs != nil {
			msg := &GameEventDotaSuperCreeps{}
			msg.Teamnumber = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 81: // EGameEvent_DotaItemPurchase
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaItemPurchase: %v", m)
			return nil
		}
		if cbs := ge.onDotaItemPurchase; cbs != nil {
			msg := &GameEventDotaItemPurchase{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Itemid = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 82: // EGameEvent_DotaItemGifted
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_DotaItemGifted: %v", m)
			return nil
		}
		if cbs := ge.onDotaItemGifted; cbs != nil {
			msg := &GameEventDotaItemGifted{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Itemid = m.GetKeys()[1].GetValShort()
			msg.Sourceid = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 83: // EGameEvent_DotaRunePickup
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_DotaRunePickup: %v", m)
			return nil
		}
		if cbs := ge.onDotaRunePickup; cbs != nil {
			msg := &GameEventDotaRunePickup{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Type = m.GetKeys()[1].GetValShort()
			msg.Rune = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 84: // EGameEvent_DotaRuneSpotted
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaRuneSpotted: %v", m)
			return nil
		}
		if cbs := ge.onDotaRuneSpotted; cbs != nil {
			msg := &GameEventDotaRuneSpotted{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Rune = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 85: // EGameEvent_DotaItemSpotted
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaItemSpotted: %v", m)
			return nil
		}
		if cbs := ge.onDotaItemSpotted; cbs != nil {
			msg := &GameEventDotaItemSpotted{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Itemid = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 86: // EGameEvent_DotaNoBattlePoints
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaNoBattlePoints: %v", m)
			return nil
		}
		if cbs := ge.onDotaNoBattlePoints; cbs != nil {
			msg := &GameEventDotaNoBattlePoints{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Reason = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 87: // EGameEvent_DotaChatInformational
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaChatInformational: %v", m)
			return nil
		}
		if cbs := ge.onDotaChatInformational; cbs != nil {
			msg := &GameEventDotaChatInformational{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Type = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 88: // EGameEvent_DotaActionItem
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_DotaActionItem: %v", m)
			return nil
		}
		if cbs := ge.onDotaActionItem; cbs != nil {
			msg := &GameEventDotaActionItem{}
			msg.Reason = m.GetKeys()[0].GetValShort()
			msg.Itemdef = m.GetKeys()[1].GetValShort()
			msg.Message = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 89: // EGameEvent_DotaChatBanNotification
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaChatBanNotification: %v", m)
			return nil
		}
		if cbs := ge.onDotaChatBanNotification; cbs != nil {
			msg := &GameEventDotaChatBanNotification{}
			msg.Userid = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 90: // EGameEvent_DotaChatEvent
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_DotaChatEvent: %v", m)
			return nil
		}
		if cbs := ge.onDotaChatEvent; cbs != nil {
			msg := &GameEventDotaChatEvent{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Gold = m.GetKeys()[1].GetValShort()
			msg.Message = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 91: // EGameEvent_DotaChatTimedReward
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_DotaChatTimedReward: %v", m)
			return nil
		}
		if cbs := ge.onDotaChatTimedReward; cbs != nil {
			msg := &GameEventDotaChatTimedReward{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Itmedef = m.GetKeys()[1].GetValShort()
			msg.Message = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 92: // EGameEvent_DotaPauseEvent
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_DotaPauseEvent: %v", m)
			return nil
		}
		if cbs := ge.onDotaPauseEvent; cbs != nil {
			msg := &GameEventDotaPauseEvent{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Value = m.GetKeys()[1].GetValShort()
			msg.Message = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 93: // EGameEvent_DotaChatKillStreak
		if len(m.GetKeys()) < 6 {
			_debugf("short EGameEvent_DotaChatKillStreak: %v", m)
			return nil
		}
		if cbs := ge.onDotaChatKillStreak; cbs != nil {
			msg := &GameEventDotaChatKillStreak{}
			msg.Gold = m.GetKeys()[0].GetValShort()
			msg.KillerId = m.GetKeys()[1].GetValShort()
			msg.KillerStreak = m.GetKeys()[2].GetValShort()
			msg.KillerMultikill = m.GetKeys()[3].GetValShort()
			msg.VictimId = m.GetKeys()[4].GetValShort()
			msg.VictimStreak = m.GetKeys()[5].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 94: // EGameEvent_DotaChatFirstBlood
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_DotaChatFirstBlood: %v", m)
			return nil
		}
		if cbs := ge.onDotaChatFirstBlood; cbs != nil {
			msg := &GameEventDotaChatFirstBlood{}
			msg.Gold = m.GetKeys()[0].GetValShort()
			msg.KillerId = m.GetKeys()[1].GetValShort()
			msg.VictimId = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 95: // EGameEvent_DotaChatAssassinAnnounce
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_DotaChatAssassinAnnounce: %v", m)
			return nil
		}
		if cbs := ge.onDotaChatAssassinAnnounce; cbs != nil {
			msg := &GameEventDotaChatAssassinAnnounce{}
			msg.AssassinId = m.GetKeys()[0].GetValShort()
			msg.TargetId = m.GetKeys()[1].GetValShort()
			msg.Message = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 96: // EGameEvent_DotaChatAssassinDenied
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_DotaChatAssassinDenied: %v", m)
			return nil
		}
		if cbs := ge.onDotaChatAssassinDenied; cbs != nil {
			msg := &GameEventDotaChatAssassinDenied{}
			msg.AssassinId = m.GetKeys()[0].GetValShort()
			msg.TargetId = m.GetKeys()[1].GetValShort()
			msg.Message = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 97: // EGameEvent_DotaChatAssassinSuccess
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_DotaChatAssassinSuccess: %v", m)
			return nil
		}
		if cbs := ge.onDotaChatAssassinSuccess; cbs != nil {
			msg := &GameEventDotaChatAssassinSuccess{}
			msg.AssassinId = m.GetKeys()[0].GetValShort()
			msg.TargetId = m.GetKeys()[1].GetValShort()
			msg.Message = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 98: // EGameEvent_DotaPlayerUpdateHeroSelection
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaPlayerUpdateHeroSelection: %v", m)
			return nil
		}
		if cbs := ge.onDotaPlayerUpdateHeroSelection; cbs != nil {
			msg := &GameEventDotaPlayerUpdateHeroSelection{}
			msg.Tabcycle = m.GetKeys()[0].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 99: // EGameEvent_DotaPlayerUpdateSelectedUnit
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaPlayerUpdateSelectedUnit: %v", m)
			return nil
		}
		if cbs := ge.onDotaPlayerUpdateSelectedUnit; cbs != nil {
			msg := &GameEventDotaPlayerUpdateSelectedUnit{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 100: // EGameEvent_DotaPlayerUpdateQueryUnit
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaPlayerUpdateQueryUnit: %v", m)
			return nil
		}
		if cbs := ge.onDotaPlayerUpdateQueryUnit; cbs != nil {
			msg := &GameEventDotaPlayerUpdateQueryUnit{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 101: // EGameEvent_DotaPlayerUpdateKillcamUnit
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaPlayerUpdateKillcamUnit: %v", m)
			return nil
		}
		if cbs := ge.onDotaPlayerUpdateKillcamUnit; cbs != nil {
			msg := &GameEventDotaPlayerUpdateKillcamUnit{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 102: // EGameEvent_DotaPlayerTakeTowerDamage
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaPlayerTakeTowerDamage: %v", m)
			return nil
		}
		if cbs := ge.onDotaPlayerTakeTowerDamage; cbs != nil {
			msg := &GameEventDotaPlayerTakeTowerDamage{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.Damage = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 103: // EGameEvent_DotaHudErrorMessage
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaHudErrorMessage: %v", m)
			return nil
		}
		if cbs := ge.onDotaHudErrorMessage; cbs != nil {
			msg := &GameEventDotaHudErrorMessage{}
			msg.Reason = m.GetKeys()[0].GetValByte()
			msg.Message = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 104: // EGameEvent_DotaActionSuccess
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaActionSuccess: %v", m)
			return nil
		}
		if cbs := ge.onDotaActionSuccess; cbs != nil {
			msg := &GameEventDotaActionSuccess{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 105: // EGameEvent_DotaStartingPositionChanged
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaStartingPositionChanged: %v", m)
			return nil
		}
		if cbs := ge.onDotaStartingPositionChanged; cbs != nil {
			msg := &GameEventDotaStartingPositionChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 106: // EGameEvent_DotaMoneyChanged
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaMoneyChanged: %v", m)
			return nil
		}
		if cbs := ge.onDotaMoneyChanged; cbs != nil {
			msg := &GameEventDotaMoneyChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 107: // EGameEvent_DotaEnemyMoneyChanged
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaEnemyMoneyChanged: %v", m)
			return nil
		}
		if cbs := ge.onDotaEnemyMoneyChanged; cbs != nil {
			msg := &GameEventDotaEnemyMoneyChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 108: // EGameEvent_DotaPortraitUnitStatsChanged
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaPortraitUnitStatsChanged: %v", m)
			return nil
		}
		if cbs := ge.onDotaPortraitUnitStatsChanged; cbs != nil {
			msg := &GameEventDotaPortraitUnitStatsChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 109: // EGameEvent_DotaPortraitUnitModifiersChanged
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaPortraitUnitModifiersChanged: %v", m)
			return nil
		}
		if cbs := ge.onDotaPortraitUnitModifiersChanged; cbs != nil {
			msg := &GameEventDotaPortraitUnitModifiersChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 110: // EGameEvent_DotaForcePortraitUpdate
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaForcePortraitUpdate: %v", m)
			return nil
		}
		if cbs := ge.onDotaForcePortraitUpdate; cbs != nil {
			msg := &GameEventDotaForcePortraitUpdate{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 111: // EGameEvent_DotaInventoryChanged
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaInventoryChanged: %v", m)
			return nil
		}
		if cbs := ge.onDotaInventoryChanged; cbs != nil {
			msg := &GameEventDotaInventoryChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 112: // EGameEvent_DotaItemPickedUp
		if len(m.GetKeys()) < 4 {
			_debugf("short EGameEvent_DotaItemPickedUp: %v", m)
			return nil
		}
		if cbs := ge.onDotaItemPickedUp; cbs != nil {
			msg := &GameEventDotaItemPickedUp{}
			msg.Itemname = m.GetKeys()[0].GetValString()
			msg.PlayerID = m.GetKeys()[1].GetValShort()
			msg.ItemEntityIndex = m.GetKeys()[2].GetValShort()
			msg.HeroEntityIndex = m.GetKeys()[3].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 113: // EGameEvent_DotaInventoryItemChanged
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaInventoryItemChanged: %v", m)
			return nil
		}
		if cbs := ge.onDotaInventoryItemChanged; cbs != nil {
			msg := &GameEventDotaInventoryItemChanged{}
			msg.EntityIndex = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 114: // EGameEvent_DotaAbilityChanged
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaAbilityChanged: %v", m)
			return nil
		}
		if cbs := ge.onDotaAbilityChanged; cbs != nil {
			msg := &GameEventDotaAbilityChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 115: // EGameEvent_DotaPortraitAbilityLayoutChanged
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaPortraitAbilityLayoutChanged: %v", m)
			return nil
		}
		if cbs := ge.onDotaPortraitAbilityLayoutChanged; cbs != nil {
			msg := &GameEventDotaPortraitAbilityLayoutChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 116: // EGameEvent_DotaInventoryItemAdded
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaInventoryItemAdded: %v", m)
			return nil
		}
		if cbs := ge.onDotaInventoryItemAdded; cbs != nil {
			msg := &GameEventDotaInventoryItemAdded{}
			msg.Itemname = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 117: // EGameEvent_DotaInventoryChangedQueryUnit
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaInventoryChangedQueryUnit: %v", m)
			return nil
		}
		if cbs := ge.onDotaInventoryChangedQueryUnit; cbs != nil {
			msg := &GameEventDotaInventoryChangedQueryUnit{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 118: // EGameEvent_DotaLinkClicked
		if len(m.GetKeys()) < 5 {
			_debugf("short EGameEvent_DotaLinkClicked: %v", m)
			return nil
		}
		if cbs := ge.onDotaLinkClicked; cbs != nil {
			msg := &GameEventDotaLinkClicked{}
			msg.Link = m.GetKeys()[0].GetValString()
			msg.Nav = m.GetKeys()[1].GetValBool()
			msg.NavBack = m.GetKeys()[2].GetValBool()
			msg.Recipe = m.GetKeys()[3].GetValShort()
			msg.Shop = m.GetKeys()[4].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 119: // EGameEvent_DotaSetQuickBuy
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_DotaSetQuickBuy: %v", m)
			return nil
		}
		if cbs := ge.onDotaSetQuickBuy; cbs != nil {
			msg := &GameEventDotaSetQuickBuy{}
			msg.Item = m.GetKeys()[0].GetValString()
			msg.Recipe = m.GetKeys()[1].GetValByte()
			msg.Toggle = m.GetKeys()[2].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 120: // EGameEvent_DotaQuickBuyChanged
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaQuickBuyChanged: %v", m)
			return nil
		}
		if cbs := ge.onDotaQuickBuyChanged; cbs != nil {
			msg := &GameEventDotaQuickBuyChanged{}
			msg.Item = m.GetKeys()[0].GetValString()
			msg.Recipe = m.GetKeys()[1].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 121: // EGameEvent_DotaPlayerShopChanged
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaPlayerShopChanged: %v", m)
			return nil
		}
		if cbs := ge.onDotaPlayerShopChanged; cbs != nil {
			msg := &GameEventDotaPlayerShopChanged{}
			msg.Prevshopmask = m.GetKeys()[0].GetValByte()
			msg.Shopmask = m.GetKeys()[1].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 122: // EGameEvent_DotaPlayerShowKillcam
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaPlayerShowKillcam: %v", m)
			return nil
		}
		if cbs := ge.onDotaPlayerShowKillcam; cbs != nil {
			msg := &GameEventDotaPlayerShowKillcam{}
			msg.Nodes = m.GetKeys()[0].GetValByte()
			msg.Player = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 123: // EGameEvent_DotaPlayerShowMinikillcam
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaPlayerShowMinikillcam: %v", m)
			return nil
		}
		if cbs := ge.onDotaPlayerShowMinikillcam; cbs != nil {
			msg := &GameEventDotaPlayerShowMinikillcam{}
			msg.Nodes = m.GetKeys()[0].GetValByte()
			msg.Player = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 124: // EGameEvent_GcUserSessionCreated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_GcUserSessionCreated: %v", m)
			return nil
		}
		if cbs := ge.onGcUserSessionCreated; cbs != nil {
			msg := &GameEventGcUserSessionCreated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 125: // EGameEvent_TeamDataUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_TeamDataUpdated: %v", m)
			return nil
		}
		if cbs := ge.onTeamDataUpdated; cbs != nil {
			msg := &GameEventTeamDataUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 126: // EGameEvent_GuildDataUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_GuildDataUpdated: %v", m)
			return nil
		}
		if cbs := ge.onGuildDataUpdated; cbs != nil {
			msg := &GameEventGuildDataUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 127: // EGameEvent_GuildOpenPartiesUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_GuildOpenPartiesUpdated: %v", m)
			return nil
		}
		if cbs := ge.onGuildOpenPartiesUpdated; cbs != nil {
			msg := &GameEventGuildOpenPartiesUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 128: // EGameEvent_FantasyUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_FantasyUpdated: %v", m)
			return nil
		}
		if cbs := ge.onFantasyUpdated; cbs != nil {
			msg := &GameEventFantasyUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 129: // EGameEvent_FantasyLeagueChanged
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_FantasyLeagueChanged: %v", m)
			return nil
		}
		if cbs := ge.onFantasyLeagueChanged; cbs != nil {
			msg := &GameEventFantasyLeagueChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 130: // EGameEvent_FantasyScoreInfoChanged
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_FantasyScoreInfoChanged: %v", m)
			return nil
		}
		if cbs := ge.onFantasyScoreInfoChanged; cbs != nil {
			msg := &GameEventFantasyScoreInfoChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 131: // EGameEvent_PlayerInfoUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_PlayerInfoUpdated: %v", m)
			return nil
		}
		if cbs := ge.onPlayerInfoUpdated; cbs != nil {
			msg := &GameEventPlayerInfoUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 132: // EGameEvent_PlayerInfoIndividualUpdated
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_PlayerInfoIndividualUpdated: %v", m)
			return nil
		}
		if cbs := ge.onPlayerInfoIndividualUpdated; cbs != nil {
			msg := &GameEventPlayerInfoIndividualUpdated{}
			msg.AccountId = m.GetKeys()[0].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 133: // EGameEvent_GameRulesStateChange
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_GameRulesStateChange: %v", m)
			return nil
		}
		if cbs := ge.onGameRulesStateChange; cbs != nil {
			msg := &GameEventGameRulesStateChange{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 134: // EGameEvent_MatchHistoryUpdated
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_MatchHistoryUpdated: %v", m)
			return nil
		}
		if cbs := ge.onMatchHistoryUpdated; cbs != nil {
			msg := &GameEventMatchHistoryUpdated{}
			msg.SteamID = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 135: // EGameEvent_MatchDetailsUpdated
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_MatchDetailsUpdated: %v", m)
			return nil
		}
		if cbs := ge.onMatchDetailsUpdated; cbs != nil {
			msg := &GameEventMatchDetailsUpdated{}
			msg.MatchID = m.GetKeys()[0].GetValUint64()
			msg.Result = m.GetKeys()[1].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 136: // EGameEvent_LiveGamesUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_LiveGamesUpdated: %v", m)
			return nil
		}
		if cbs := ge.onLiveGamesUpdated; cbs != nil {
			msg := &GameEventLiveGamesUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 137: // EGameEvent_RecentMatchesUpdated
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_RecentMatchesUpdated: %v", m)
			return nil
		}
		if cbs := ge.onRecentMatchesUpdated; cbs != nil {
			msg := &GameEventRecentMatchesUpdated{}
			msg.Page = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 138: // EGameEvent_NewsUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_NewsUpdated: %v", m)
			return nil
		}
		if cbs := ge.onNewsUpdated; cbs != nil {
			msg := &GameEventNewsUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 139: // EGameEvent_PersonaUpdated
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_PersonaUpdated: %v", m)
			return nil
		}
		if cbs := ge.onPersonaUpdated; cbs != nil {
			msg := &GameEventPersonaUpdated{}
			msg.SteamID = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 140: // EGameEvent_TournamentStateUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_TournamentStateUpdated: %v", m)
			return nil
		}
		if cbs := ge.onTournamentStateUpdated; cbs != nil {
			msg := &GameEventTournamentStateUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 141: // EGameEvent_PartyUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_PartyUpdated: %v", m)
			return nil
		}
		if cbs := ge.onPartyUpdated; cbs != nil {
			msg := &GameEventPartyUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 142: // EGameEvent_LobbyUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_LobbyUpdated: %v", m)
			return nil
		}
		if cbs := ge.onLobbyUpdated; cbs != nil {
			msg := &GameEventLobbyUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 143: // EGameEvent_DashboardCachesCleared
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DashboardCachesCleared: %v", m)
			return nil
		}
		if cbs := ge.onDashboardCachesCleared; cbs != nil {
			msg := &GameEventDashboardCachesCleared{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 144: // EGameEvent_LastHit
		if len(m.GetKeys()) < 5 {
			_debugf("short EGameEvent_LastHit: %v", m)
			return nil
		}
		if cbs := ge.onLastHit; cbs != nil {
			msg := &GameEventLastHit{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.EntKilled = m.GetKeys()[1].GetValShort()
			msg.FirstBlood = m.GetKeys()[2].GetValBool()
			msg.HeroKill = m.GetKeys()[3].GetValBool()
			msg.TowerKill = m.GetKeys()[4].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 145: // EGameEvent_PlayerCompletedGame
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_PlayerCompletedGame: %v", m)
			return nil
		}
		if cbs := ge.onPlayerCompletedGame; cbs != nil {
			msg := &GameEventPlayerCompletedGame{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.Winner = m.GetKeys()[1].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 146: // EGameEvent_PlayerReconnected
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_PlayerReconnected: %v", m)
			return nil
		}
		if cbs := ge.onPlayerReconnected; cbs != nil {
			msg := &GameEventPlayerReconnected{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 147: // EGameEvent_NommedTree
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_NommedTree: %v", m)
			return nil
		}
		if cbs := ge.onNommedTree; cbs != nil {
			msg := &GameEventNommedTree{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 148: // EGameEvent_DotaRuneActivatedServer
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaRuneActivatedServer: %v", m)
			return nil
		}
		if cbs := ge.onDotaRuneActivatedServer; cbs != nil {
			msg := &GameEventDotaRuneActivatedServer{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.Rune = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 149: // EGameEvent_DotaPlayerGainedLevel
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaPlayerGainedLevel: %v", m)
			return nil
		}
		if cbs := ge.onDotaPlayerGainedLevel; cbs != nil {
			msg := &GameEventDotaPlayerGainedLevel{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.Level = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 150: // EGameEvent_DotaPlayerLearnedAbility
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaPlayerLearnedAbility: %v", m)
			return nil
		}
		if cbs := ge.onDotaPlayerLearnedAbility; cbs != nil {
			msg := &GameEventDotaPlayerLearnedAbility{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.Abilityname = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 151: // EGameEvent_DotaPlayerUsedAbility
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaPlayerUsedAbility: %v", m)
			return nil
		}
		if cbs := ge.onDotaPlayerUsedAbility; cbs != nil {
			msg := &GameEventDotaPlayerUsedAbility{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.Abilityname = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 152: // EGameEvent_DotaNonPlayerUsedAbility
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaNonPlayerUsedAbility: %v", m)
			return nil
		}
		if cbs := ge.onDotaNonPlayerUsedAbility; cbs != nil {
			msg := &GameEventDotaNonPlayerUsedAbility{}
			msg.Abilityname = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 153: // EGameEvent_DotaPlayerBeginCast
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaPlayerBeginCast: %v", m)
			return nil
		}
		if cbs := ge.onDotaPlayerBeginCast; cbs != nil {
			msg := &GameEventDotaPlayerBeginCast{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.Abilityname = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 154: // EGameEvent_DotaNonPlayerBeginCast
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaNonPlayerBeginCast: %v", m)
			return nil
		}
		if cbs := ge.onDotaNonPlayerBeginCast; cbs != nil {
			msg := &GameEventDotaNonPlayerBeginCast{}
			msg.Abilityname = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 155: // EGameEvent_DotaAbilityChannelFinished
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaAbilityChannelFinished: %v", m)
			return nil
		}
		if cbs := ge.onDotaAbilityChannelFinished; cbs != nil {
			msg := &GameEventDotaAbilityChannelFinished{}
			msg.Abilityname = m.GetKeys()[0].GetValString()
			msg.Interrupted = m.GetKeys()[1].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 156: // EGameEvent_DotaHoldoutReviveComplete
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaHoldoutReviveComplete: %v", m)
			return nil
		}
		if cbs := ge.onDotaHoldoutReviveComplete; cbs != nil {
			msg := &GameEventDotaHoldoutReviveComplete{}
			msg.Caster = m.GetKeys()[0].GetValShort()
			msg.Target = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 157: // EGameEvent_DotaPlayerKilled
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_DotaPlayerKilled: %v", m)
			return nil
		}
		if cbs := ge.onDotaPlayerKilled; cbs != nil {
			msg := &GameEventDotaPlayerKilled{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.HeroKill = m.GetKeys()[1].GetValBool()
			msg.TowerKill = m.GetKeys()[2].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 158: // EGameEvent_BindpanelOpen
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_BindpanelOpen: %v", m)
			return nil
		}
		if cbs := ge.onBindpanelOpen; cbs != nil {
			msg := &GameEventBindpanelOpen{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 159: // EGameEvent_BindpanelClose
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_BindpanelClose: %v", m)
			return nil
		}
		if cbs := ge.onBindpanelClose; cbs != nil {
			msg := &GameEventBindpanelClose{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 160: // EGameEvent_KeybindChanged
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_KeybindChanged: %v", m)
			return nil
		}
		if cbs := ge.onKeybindChanged; cbs != nil {
			msg := &GameEventKeybindChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 161: // EGameEvent_DotaItemDragBegin
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaItemDragBegin: %v", m)
			return nil
		}
		if cbs := ge.onDotaItemDragBegin; cbs != nil {
			msg := &GameEventDotaItemDragBegin{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 162: // EGameEvent_DotaItemDragEnd
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaItemDragEnd: %v", m)
			return nil
		}
		if cbs := ge.onDotaItemDragEnd; cbs != nil {
			msg := &GameEventDotaItemDragEnd{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 163: // EGameEvent_DotaShopItemDragBegin
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaShopItemDragBegin: %v", m)
			return nil
		}
		if cbs := ge.onDotaShopItemDragBegin; cbs != nil {
			msg := &GameEventDotaShopItemDragBegin{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 164: // EGameEvent_DotaShopItemDragEnd
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaShopItemDragEnd: %v", m)
			return nil
		}
		if cbs := ge.onDotaShopItemDragEnd; cbs != nil {
			msg := &GameEventDotaShopItemDragEnd{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 165: // EGameEvent_DotaItemPurchased
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_DotaItemPurchased: %v", m)
			return nil
		}
		if cbs := ge.onDotaItemPurchased; cbs != nil {
			msg := &GameEventDotaItemPurchased{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.Itemname = m.GetKeys()[1].GetValString()
			msg.Itemcost = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 166: // EGameEvent_DotaItemCombined
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_DotaItemCombined: %v", m)
			return nil
		}
		if cbs := ge.onDotaItemCombined; cbs != nil {
			msg := &GameEventDotaItemCombined{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.Itemname = m.GetKeys()[1].GetValString()
			msg.Itemcost = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 167: // EGameEvent_DotaItemUsed
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaItemUsed: %v", m)
			return nil
		}
		if cbs := ge.onDotaItemUsed; cbs != nil {
			msg := &GameEventDotaItemUsed{}
			msg.PlayerID = m.GetKeys()[0].GetValShort()
			msg.Itemname = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 168: // EGameEvent_DotaItemAutoPurchase
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaItemAutoPurchase: %v", m)
			return nil
		}
		if cbs := ge.onDotaItemAutoPurchase; cbs != nil {
			msg := &GameEventDotaItemAutoPurchase{}
			msg.ItemId = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 169: // EGameEvent_DotaUnitEvent
		if len(m.GetKeys()) < 5 {
			_debugf("short EGameEvent_DotaUnitEvent: %v", m)
			return nil
		}
		if cbs := ge.onDotaUnitEvent; cbs != nil {
			msg := &GameEventDotaUnitEvent{}
			msg.Victim = m.GetKeys()[0].GetValShort()
			msg.Attacker = m.GetKeys()[1].GetValShort()
			msg.Basepriority = m.GetKeys()[2].GetValShort()
			msg.Priority = m.GetKeys()[3].GetValShort()
			msg.Eventtype = m.GetKeys()[4].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 170: // EGameEvent_DotaQuestStarted
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaQuestStarted: %v", m)
			return nil
		}
		if cbs := ge.onDotaQuestStarted; cbs != nil {
			msg := &GameEventDotaQuestStarted{}
			msg.QuestIndex = m.GetKeys()[0].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 171: // EGameEvent_DotaQuestCompleted
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaQuestCompleted: %v", m)
			return nil
		}
		if cbs := ge.onDotaQuestCompleted; cbs != nil {
			msg := &GameEventDotaQuestCompleted{}
			msg.QuestIndex = m.GetKeys()[0].GetValLong()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 172: // EGameEvent_GameuiActivated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_GameuiActivated: %v", m)
			return nil
		}
		if cbs := ge.onGameuiActivated; cbs != nil {
			msg := &GameEventGameuiActivated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 173: // EGameEvent_GameuiHidden
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_GameuiHidden: %v", m)
			return nil
		}
		if cbs := ge.onGameuiHidden; cbs != nil {
			msg := &GameEventGameuiHidden{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 174: // EGameEvent_PlayerFullyjoined
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_PlayerFullyjoined: %v", m)
			return nil
		}
		if cbs := ge.onPlayerFullyjoined; cbs != nil {
			msg := &GameEventPlayerFullyjoined{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Name = m.GetKeys()[1].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 175: // EGameEvent_DotaSpectateHero
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaSpectateHero: %v", m)
			return nil
		}
		if cbs := ge.onDotaSpectateHero; cbs != nil {
			msg := &GameEventDotaSpectateHero{}
			msg.Entindex = m.GetKeys()[0].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 176: // EGameEvent_DotaMatchDone
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaMatchDone: %v", m)
			return nil
		}
		if cbs := ge.onDotaMatchDone; cbs != nil {
			msg := &GameEventDotaMatchDone{}
			msg.Winningteam = m.GetKeys()[0].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 177: // EGameEvent_DotaMatchDoneClient
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaMatchDoneClient: %v", m)
			return nil
		}
		if cbs := ge.onDotaMatchDoneClient; cbs != nil {
			msg := &GameEventDotaMatchDoneClient{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 178: // EGameEvent_SetInstructorGroupEnabled
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_SetInstructorGroupEnabled: %v", m)
			return nil
		}
		if cbs := ge.onSetInstructorGroupEnabled; cbs != nil {
			msg := &GameEventSetInstructorGroupEnabled{}
			msg.Group = m.GetKeys()[0].GetValString()
			msg.Enabled = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 179: // EGameEvent_JoinedChatChannel
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_JoinedChatChannel: %v", m)
			return nil
		}
		if cbs := ge.onJoinedChatChannel; cbs != nil {
			msg := &GameEventJoinedChatChannel{}
			msg.ChannelName = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 180: // EGameEvent_LeftChatChannel
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_LeftChatChannel: %v", m)
			return nil
		}
		if cbs := ge.onLeftChatChannel; cbs != nil {
			msg := &GameEventLeftChatChannel{}
			msg.ChannelName = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 181: // EGameEvent_GcChatChannelListUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_GcChatChannelListUpdated: %v", m)
			return nil
		}
		if cbs := ge.onGcChatChannelListUpdated; cbs != nil {
			msg := &GameEventGcChatChannelListUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 182: // EGameEvent_TodayMessagesUpdated
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_TodayMessagesUpdated: %v", m)
			return nil
		}
		if cbs := ge.onTodayMessagesUpdated; cbs != nil {
			msg := &GameEventTodayMessagesUpdated{}
			msg.NumMessages = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 183: // EGameEvent_FileDownloaded
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_FileDownloaded: %v", m)
			return nil
		}
		if cbs := ge.onFileDownloaded; cbs != nil {
			msg := &GameEventFileDownloaded{}
			msg.Success = m.GetKeys()[0].GetValBool()
			msg.LocalFilename = m.GetKeys()[1].GetValString()
			msg.RemoteUrl = m.GetKeys()[2].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 184: // EGameEvent_PlayerReportCountsUpdated
		if len(m.GetKeys()) < 4 {
			_debugf("short EGameEvent_PlayerReportCountsUpdated: %v", m)
			return nil
		}
		if cbs := ge.onPlayerReportCountsUpdated; cbs != nil {
			msg := &GameEventPlayerReportCountsUpdated{}
			msg.PositiveRemaining = m.GetKeys()[0].GetValByte()
			msg.NegativeRemaining = m.GetKeys()[1].GetValByte()
			msg.PositiveTotal = m.GetKeys()[2].GetValShort()
			msg.NegativeTotal = m.GetKeys()[3].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 185: // EGameEvent_ScaleformFileDownloadComplete
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_ScaleformFileDownloadComplete: %v", m)
			return nil
		}
		if cbs := ge.onScaleformFileDownloadComplete; cbs != nil {
			msg := &GameEventScaleformFileDownloadComplete{}
			msg.Success = m.GetKeys()[0].GetValBool()
			msg.LocalFilename = m.GetKeys()[1].GetValString()
			msg.RemoteUrl = m.GetKeys()[2].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 186: // EGameEvent_ItemPurchased
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_ItemPurchased: %v", m)
			return nil
		}
		if cbs := ge.onItemPurchased; cbs != nil {
			msg := &GameEventItemPurchased{}
			msg.Itemid = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 187: // EGameEvent_GcMismatchedVersion
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_GcMismatchedVersion: %v", m)
			return nil
		}
		if cbs := ge.onGcMismatchedVersion; cbs != nil {
			msg := &GameEventGcMismatchedVersion{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 190: // EGameEvent_DemoStop
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DemoStop: %v", m)
			return nil
		}
		if cbs := ge.onDemoStop; cbs != nil {
			msg := &GameEventDemoStop{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 191: // EGameEvent_MapShutdown
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_MapShutdown: %v", m)
			return nil
		}
		if cbs := ge.onMapShutdown; cbs != nil {
			msg := &GameEventMapShutdown{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 192: // EGameEvent_DotaWorkshopFileselected
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaWorkshopFileselected: %v", m)
			return nil
		}
		if cbs := ge.onDotaWorkshopFileselected; cbs != nil {
			msg := &GameEventDotaWorkshopFileselected{}
			msg.Filename = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 193: // EGameEvent_DotaWorkshopFilecanceled
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaWorkshopFilecanceled: %v", m)
			return nil
		}
		if cbs := ge.onDotaWorkshopFilecanceled; cbs != nil {
			msg := &GameEventDotaWorkshopFilecanceled{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 194: // EGameEvent_RichPresenceUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_RichPresenceUpdated: %v", m)
			return nil
		}
		if cbs := ge.onRichPresenceUpdated; cbs != nil {
			msg := &GameEventRichPresenceUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 195: // EGameEvent_DotaHeroRandom
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaHeroRandom: %v", m)
			return nil
		}
		if cbs := ge.onDotaHeroRandom; cbs != nil {
			msg := &GameEventDotaHeroRandom{}
			msg.Userid = m.GetKeys()[0].GetValShort()
			msg.Heroid = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 196: // EGameEvent_DotaRdChatTurn
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaRdChatTurn: %v", m)
			return nil
		}
		if cbs := ge.onDotaRdChatTurn; cbs != nil {
			msg := &GameEventDotaRdChatTurn{}
			msg.Userid = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 197: // EGameEvent_DotaFavoriteHeroesUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaFavoriteHeroesUpdated: %v", m)
			return nil
		}
		if cbs := ge.onDotaFavoriteHeroesUpdated; cbs != nil {
			msg := &GameEventDotaFavoriteHeroesUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 198: // EGameEvent_ProfileOpened
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_ProfileOpened: %v", m)
			return nil
		}
		if cbs := ge.onProfileOpened; cbs != nil {
			msg := &GameEventProfileOpened{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 199: // EGameEvent_ProfileClosed
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_ProfileClosed: %v", m)
			return nil
		}
		if cbs := ge.onProfileClosed; cbs != nil {
			msg := &GameEventProfileClosed{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 200: // EGameEvent_ItemPreviewClosed
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_ItemPreviewClosed: %v", m)
			return nil
		}
		if cbs := ge.onItemPreviewClosed; cbs != nil {
			msg := &GameEventItemPreviewClosed{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 201: // EGameEvent_DashboardSwitchedSection
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DashboardSwitchedSection: %v", m)
			return nil
		}
		if cbs := ge.onDashboardSwitchedSection; cbs != nil {
			msg := &GameEventDashboardSwitchedSection{}
			msg.Section = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 202: // EGameEvent_DotaTournamentItemEvent
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaTournamentItemEvent: %v", m)
			return nil
		}
		if cbs := ge.onDotaTournamentItemEvent; cbs != nil {
			msg := &GameEventDotaTournamentItemEvent{}
			msg.WinnerCount = m.GetKeys()[0].GetValShort()
			msg.EventType = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 203: // EGameEvent_DotaHeroSwap
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaHeroSwap: %v", m)
			return nil
		}
		if cbs := ge.onDotaHeroSwap; cbs != nil {
			msg := &GameEventDotaHeroSwap{}
			msg.Playerid1 = m.GetKeys()[0].GetValByte()
			msg.Playerid2 = m.GetKeys()[1].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 204: // EGameEvent_DotaResetSuggestedItems
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaResetSuggestedItems: %v", m)
			return nil
		}
		if cbs := ge.onDotaResetSuggestedItems; cbs != nil {
			msg := &GameEventDotaResetSuggestedItems{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 205: // EGameEvent_HalloweenHighScoreReceived
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_HalloweenHighScoreReceived: %v", m)
			return nil
		}
		if cbs := ge.onHalloweenHighScoreReceived; cbs != nil {
			msg := &GameEventHalloweenHighScoreReceived{}
			msg.Round = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 206: // EGameEvent_HalloweenPhaseEnd
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_HalloweenPhaseEnd: %v", m)
			return nil
		}
		if cbs := ge.onHalloweenPhaseEnd; cbs != nil {
			msg := &GameEventHalloweenPhaseEnd{}
			msg.Phase = m.GetKeys()[0].GetValByte()
			msg.Team = m.GetKeys()[1].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 207: // EGameEvent_HalloweenHighScoreRequestFailed
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_HalloweenHighScoreRequestFailed: %v", m)
			return nil
		}
		if cbs := ge.onHalloweenHighScoreRequestFailed; cbs != nil {
			msg := &GameEventHalloweenHighScoreRequestFailed{}
			msg.Round = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 208: // EGameEvent_DotaHudSkinChanged
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaHudSkinChanged: %v", m)
			return nil
		}
		if cbs := ge.onDotaHudSkinChanged; cbs != nil {
			msg := &GameEventDotaHudSkinChanged{}
			msg.Skin = m.GetKeys()[0].GetValString()
			msg.Style = m.GetKeys()[1].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 209: // EGameEvent_DotaInventoryPlayerGotItem
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaInventoryPlayerGotItem: %v", m)
			return nil
		}
		if cbs := ge.onDotaInventoryPlayerGotItem; cbs != nil {
			msg := &GameEventDotaInventoryPlayerGotItem{}
			msg.Itemname = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 210: // EGameEvent_PlayerIsExperienced
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_PlayerIsExperienced: %v", m)
			return nil
		}
		if cbs := ge.onPlayerIsExperienced; cbs != nil {
			msg := &GameEventPlayerIsExperienced{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 211: // EGameEvent_PlayerIsNotexperienced
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_PlayerIsNotexperienced: %v", m)
			return nil
		}
		if cbs := ge.onPlayerIsNotexperienced; cbs != nil {
			msg := &GameEventPlayerIsNotexperienced{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 212: // EGameEvent_DotaTutorialLessonStart
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaTutorialLessonStart: %v", m)
			return nil
		}
		if cbs := ge.onDotaTutorialLessonStart; cbs != nil {
			msg := &GameEventDotaTutorialLessonStart{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 213: // EGameEvent_DotaTutorialTaskAdvance
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaTutorialTaskAdvance: %v", m)
			return nil
		}
		if cbs := ge.onDotaTutorialTaskAdvance; cbs != nil {
			msg := &GameEventDotaTutorialTaskAdvance{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 214: // EGameEvent_DotaTutorialShopToggled
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaTutorialShopToggled: %v", m)
			return nil
		}
		if cbs := ge.onDotaTutorialShopToggled; cbs != nil {
			msg := &GameEventDotaTutorialShopToggled{}
			msg.ShopOpened = m.GetKeys()[0].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 215: // EGameEvent_MapLocationUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_MapLocationUpdated: %v", m)
			return nil
		}
		if cbs := ge.onMapLocationUpdated; cbs != nil {
			msg := &GameEventMapLocationUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 216: // EGameEvent_RichpresenceCustomUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_RichpresenceCustomUpdated: %v", m)
			return nil
		}
		if cbs := ge.onRichpresenceCustomUpdated; cbs != nil {
			msg := &GameEventRichpresenceCustomUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 217: // EGameEvent_GameEndVisible
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_GameEndVisible: %v", m)
			return nil
		}
		if cbs := ge.onGameEndVisible; cbs != nil {
			msg := &GameEventGameEndVisible{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 218: // EGameEvent_AntiaddictionUpdate
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_AntiaddictionUpdate: %v", m)
			return nil
		}
		if cbs := ge.onAntiaddictionUpdate; cbs != nil {
			msg := &GameEventAntiaddictionUpdate{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 219: // EGameEvent_HighlightHudElement
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_HighlightHudElement: %v", m)
			return nil
		}
		if cbs := ge.onHighlightHudElement; cbs != nil {
			msg := &GameEventHighlightHudElement{}
			msg.Elementname = m.GetKeys()[0].GetValString()
			msg.Duration = m.GetKeys()[1].GetValFloat()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 220: // EGameEvent_HideHighlightHudElement
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_HideHighlightHudElement: %v", m)
			return nil
		}
		if cbs := ge.onHideHighlightHudElement; cbs != nil {
			msg := &GameEventHideHighlightHudElement{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 221: // EGameEvent_IntroVideoFinished
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_IntroVideoFinished: %v", m)
			return nil
		}
		if cbs := ge.onIntroVideoFinished; cbs != nil {
			msg := &GameEventIntroVideoFinished{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 222: // EGameEvent_MatchmakingStatusVisibilityChanged
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_MatchmakingStatusVisibilityChanged: %v", m)
			return nil
		}
		if cbs := ge.onMatchmakingStatusVisibilityChanged; cbs != nil {
			msg := &GameEventMatchmakingStatusVisibilityChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 223: // EGameEvent_PracticeLobbyVisibilityChanged
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_PracticeLobbyVisibilityChanged: %v", m)
			return nil
		}
		if cbs := ge.onPracticeLobbyVisibilityChanged; cbs != nil {
			msg := &GameEventPracticeLobbyVisibilityChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 224: // EGameEvent_DotaCourierTransferItem
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaCourierTransferItem: %v", m)
			return nil
		}
		if cbs := ge.onDotaCourierTransferItem; cbs != nil {
			msg := &GameEventDotaCourierTransferItem{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 225: // EGameEvent_FullUiUnlocked
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_FullUiUnlocked: %v", m)
			return nil
		}
		if cbs := ge.onFullUiUnlocked; cbs != nil {
			msg := &GameEventFullUiUnlocked{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 227: // EGameEvent_HeroSelectorPreviewSet
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_HeroSelectorPreviewSet: %v", m)
			return nil
		}
		if cbs := ge.onHeroSelectorPreviewSet; cbs != nil {
			msg := &GameEventHeroSelectorPreviewSet{}
			msg.Setindex = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 228: // EGameEvent_AntiaddictionToast
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_AntiaddictionToast: %v", m)
			return nil
		}
		if cbs := ge.onAntiaddictionToast; cbs != nil {
			msg := &GameEventAntiaddictionToast{}
			msg.Message = m.GetKeys()[0].GetValString()
			msg.Duration = m.GetKeys()[1].GetValFloat()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 229: // EGameEvent_HeroPickerShown
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_HeroPickerShown: %v", m)
			return nil
		}
		if cbs := ge.onHeroPickerShown; cbs != nil {
			msg := &GameEventHeroPickerShown{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 230: // EGameEvent_HeroPickerHidden
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_HeroPickerHidden: %v", m)
			return nil
		}
		if cbs := ge.onHeroPickerHidden; cbs != nil {
			msg := &GameEventHeroPickerHidden{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 231: // EGameEvent_DotaLocalQuickbuyChanged
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaLocalQuickbuyChanged: %v", m)
			return nil
		}
		if cbs := ge.onDotaLocalQuickbuyChanged; cbs != nil {
			msg := &GameEventDotaLocalQuickbuyChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 232: // EGameEvent_ShowCenterMessage
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_ShowCenterMessage: %v", m)
			return nil
		}
		if cbs := ge.onShowCenterMessage; cbs != nil {
			msg := &GameEventShowCenterMessage{}
			msg.Message = m.GetKeys()[0].GetValString()
			msg.Duration = m.GetKeys()[1].GetValFloat()
			msg.ClearMessageQueue = m.GetKeys()[2].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 233: // EGameEvent_HudFlipChanged
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_HudFlipChanged: %v", m)
			return nil
		}
		if cbs := ge.onHudFlipChanged; cbs != nil {
			msg := &GameEventHudFlipChanged{}
			msg.Flipped = m.GetKeys()[0].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 234: // EGameEvent_FrostyPointsUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_FrostyPointsUpdated: %v", m)
			return nil
		}
		if cbs := ge.onFrostyPointsUpdated; cbs != nil {
			msg := &GameEventFrostyPointsUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 235: // EGameEvent_Defeated
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_Defeated: %v", m)
			return nil
		}
		if cbs := ge.onDefeated; cbs != nil {
			msg := &GameEventDefeated{}
			msg.Entindex = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 236: // EGameEvent_ResetDefeated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_ResetDefeated: %v", m)
			return nil
		}
		if cbs := ge.onResetDefeated; cbs != nil {
			msg := &GameEventResetDefeated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 237: // EGameEvent_BoosterStateUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_BoosterStateUpdated: %v", m)
			return nil
		}
		if cbs := ge.onBoosterStateUpdated; cbs != nil {
			msg := &GameEventBoosterStateUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 238: // EGameEvent_EventPointsUpdated
		if len(m.GetKeys()) < 4 {
			_debugf("short EGameEvent_EventPointsUpdated: %v", m)
			return nil
		}
		if cbs := ge.onEventPointsUpdated; cbs != nil {
			msg := &GameEventEventPointsUpdated{}
			msg.EventId = m.GetKeys()[0].GetValShort()
			msg.Points = m.GetKeys()[1].GetValShort()
			msg.PremiumPoints = m.GetKeys()[2].GetValShort()
			msg.Owned = m.GetKeys()[3].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 239: // EGameEvent_LocalPlayerEventPoints
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_LocalPlayerEventPoints: %v", m)
			return nil
		}
		if cbs := ge.onLocalPlayerEventPoints; cbs != nil {
			msg := &GameEventLocalPlayerEventPoints{}
			msg.Points = m.GetKeys()[0].GetValShort()
			msg.ConversionRate = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 240: // EGameEvent_CustomGameDifficulty
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_CustomGameDifficulty: %v", m)
			return nil
		}
		if cbs := ge.onCustomGameDifficulty; cbs != nil {
			msg := &GameEventCustomGameDifficulty{}
			msg.Difficulty = m.GetKeys()[0].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 241: // EGameEvent_TreeCut
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_TreeCut: %v", m)
			return nil
		}
		if cbs := ge.onTreeCut; cbs != nil {
			msg := &GameEventTreeCut{}
			msg.TreeX = m.GetKeys()[0].GetValFloat()
			msg.TreeY = m.GetKeys()[1].GetValFloat()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 242: // EGameEvent_UgcDetailsArrived
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_UgcDetailsArrived: %v", m)
			return nil
		}
		if cbs := ge.onUgcDetailsArrived; cbs != nil {
			msg := &GameEventUgcDetailsArrived{}
			msg.PublishedFileId = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 243: // EGameEvent_UgcSubscribed
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_UgcSubscribed: %v", m)
			return nil
		}
		if cbs := ge.onUgcSubscribed; cbs != nil {
			msg := &GameEventUgcSubscribed{}
			msg.PublishedFileId = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 244: // EGameEvent_UgcUnsubscribed
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_UgcUnsubscribed: %v", m)
			return nil
		}
		if cbs := ge.onUgcUnsubscribed; cbs != nil {
			msg := &GameEventUgcUnsubscribed{}
			msg.PublishedFileId = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 245: // EGameEvent_UgcDownloadRequested
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_UgcDownloadRequested: %v", m)
			return nil
		}
		if cbs := ge.onUgcDownloadRequested; cbs != nil {
			msg := &GameEventUgcDownloadRequested{}
			msg.PublishedFileId = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 246: // EGameEvent_UgcInstalled
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_UgcInstalled: %v", m)
			return nil
		}
		if cbs := ge.onUgcInstalled; cbs != nil {
			msg := &GameEventUgcInstalled{}
			msg.PublishedFileId = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 247: // EGameEvent_PrizepoolReceived
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_PrizepoolReceived: %v", m)
			return nil
		}
		if cbs := ge.onPrizepoolReceived; cbs != nil {
			msg := &GameEventPrizepoolReceived{}
			msg.Success = m.GetKeys()[0].GetValBool()
			msg.Prizepool = m.GetKeys()[1].GetValUint64()
			msg.Leagueid = m.GetKeys()[2].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 248: // EGameEvent_MicrotransactionSuccess
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_MicrotransactionSuccess: %v", m)
			return nil
		}
		if cbs := ge.onMicrotransactionSuccess; cbs != nil {
			msg := &GameEventMicrotransactionSuccess{}
			msg.Txnid = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 249: // EGameEvent_DotaRubickAbilitySteal
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaRubickAbilitySteal: %v", m)
			return nil
		}
		if cbs := ge.onDotaRubickAbilitySteal; cbs != nil {
			msg := &GameEventDotaRubickAbilitySteal{}
			msg.AbilityIndex = m.GetKeys()[0].GetValShort()
			msg.AbilityLevel = m.GetKeys()[1].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 250: // EGameEvent_CompendiumEventActionsLoaded
		if len(m.GetKeys()) < 4 {
			_debugf("short EGameEvent_CompendiumEventActionsLoaded: %v", m)
			return nil
		}
		if cbs := ge.onCompendiumEventActionsLoaded; cbs != nil {
			msg := &GameEventCompendiumEventActionsLoaded{}
			msg.AccountId = m.GetKeys()[0].GetValUint64()
			msg.LeagueId = m.GetKeys()[1].GetValUint64()
			msg.LocalTest = m.GetKeys()[2].GetValBool()
			msg.OriginalPoints = m.GetKeys()[3].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 251: // EGameEvent_CompendiumSelectionsLoaded
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_CompendiumSelectionsLoaded: %v", m)
			return nil
		}
		if cbs := ge.onCompendiumSelectionsLoaded; cbs != nil {
			msg := &GameEventCompendiumSelectionsLoaded{}
			msg.AccountId = m.GetKeys()[0].GetValUint64()
			msg.LeagueId = m.GetKeys()[1].GetValUint64()
			msg.LocalTest = m.GetKeys()[2].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 252: // EGameEvent_CompendiumSetSelectionFailed
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_CompendiumSetSelectionFailed: %v", m)
			return nil
		}
		if cbs := ge.onCompendiumSetSelectionFailed; cbs != nil {
			msg := &GameEventCompendiumSetSelectionFailed{}
			msg.AccountId = m.GetKeys()[0].GetValUint64()
			msg.LeagueId = m.GetKeys()[1].GetValUint64()
			msg.LocalTest = m.GetKeys()[2].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 253: // EGameEvent_CompendiumTrophiesLoaded
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_CompendiumTrophiesLoaded: %v", m)
			return nil
		}
		if cbs := ge.onCompendiumTrophiesLoaded; cbs != nil {
			msg := &GameEventCompendiumTrophiesLoaded{}
			msg.AccountId = m.GetKeys()[0].GetValUint64()
			msg.LeagueId = m.GetKeys()[1].GetValUint64()
			msg.LocalTest = m.GetKeys()[2].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 254: // EGameEvent_CommunityCachedNamesUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_CommunityCachedNamesUpdated: %v", m)
			return nil
		}
		if cbs := ge.onCommunityCachedNamesUpdated; cbs != nil {
			msg := &GameEventCommunityCachedNamesUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 255: // EGameEvent_SpecItemPickup
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_SpecItemPickup: %v", m)
			return nil
		}
		if cbs := ge.onSpecItemPickup; cbs != nil {
			msg := &GameEventSpecItemPickup{}
			msg.PlayerId = m.GetKeys()[0].GetValShort()
			msg.ItemName = m.GetKeys()[1].GetValString()
			msg.Purchase = m.GetKeys()[2].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 256: // EGameEvent_SpecAegisReclaimTime
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_SpecAegisReclaimTime: %v", m)
			return nil
		}
		if cbs := ge.onSpecAegisReclaimTime; cbs != nil {
			msg := &GameEventSpecAegisReclaimTime{}
			msg.ReclaimTime = m.GetKeys()[0].GetValFloat()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 257: // EGameEvent_AccountTrophiesChanged
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_AccountTrophiesChanged: %v", m)
			return nil
		}
		if cbs := ge.onAccountTrophiesChanged; cbs != nil {
			msg := &GameEventAccountTrophiesChanged{}
			msg.AccountId = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 258: // EGameEvent_AccountAllHeroChallengeChanged
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_AccountAllHeroChallengeChanged: %v", m)
			return nil
		}
		if cbs := ge.onAccountAllHeroChallengeChanged; cbs != nil {
			msg := &GameEventAccountAllHeroChallengeChanged{}
			msg.AccountId = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 259: // EGameEvent_TeamShowcaseUiUpdate
		if len(m.GetKeys()) < 4 {
			_debugf("short EGameEvent_TeamShowcaseUiUpdate: %v", m)
			return nil
		}
		if cbs := ge.onTeamShowcaseUiUpdate; cbs != nil {
			msg := &GameEventTeamShowcaseUiUpdate{}
			msg.Show = m.GetKeys()[0].GetValBool()
			msg.AccountId = m.GetKeys()[1].GetValUint64()
			msg.HeroEntindex = m.GetKeys()[2].GetValShort()
			msg.DisplayUiOnLeft = m.GetKeys()[3].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 260: // EGameEvent_IngameEventsChanged
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_IngameEventsChanged: %v", m)
			return nil
		}
		if cbs := ge.onIngameEventsChanged; cbs != nil {
			msg := &GameEventIngameEventsChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 261: // EGameEvent_DotaMatchSignout
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaMatchSignout: %v", m)
			return nil
		}
		if cbs := ge.onDotaMatchSignout; cbs != nil {
			msg := &GameEventDotaMatchSignout{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 262: // EGameEvent_DotaIllusionsCreated
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaIllusionsCreated: %v", m)
			return nil
		}
		if cbs := ge.onDotaIllusionsCreated; cbs != nil {
			msg := &GameEventDotaIllusionsCreated{}
			msg.OriginalEntindex = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 263: // EGameEvent_DotaYearBeastKilled
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_DotaYearBeastKilled: %v", m)
			return nil
		}
		if cbs := ge.onDotaYearBeastKilled; cbs != nil {
			msg := &GameEventDotaYearBeastKilled{}
			msg.KillerPlayerId = m.GetKeys()[0].GetValShort()
			msg.Message = m.GetKeys()[1].GetValShort()
			msg.BeastId = m.GetKeys()[2].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 264: // EGameEvent_DotaHeroUndoselection
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_DotaHeroUndoselection: %v", m)
			return nil
		}
		if cbs := ge.onDotaHeroUndoselection; cbs != nil {
			msg := &GameEventDotaHeroUndoselection{}
			msg.Playerid1 = m.GetKeys()[0].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 265: // EGameEvent_DotaChallengeSocacheUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaChallengeSocacheUpdated: %v", m)
			return nil
		}
		if cbs := ge.onDotaChallengeSocacheUpdated; cbs != nil {
			msg := &GameEventDotaChallengeSocacheUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 266: // EGameEvent_PartyInvitesUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_PartyInvitesUpdated: %v", m)
			return nil
		}
		if cbs := ge.onPartyInvitesUpdated; cbs != nil {
			msg := &GameEventPartyInvitesUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 267: // EGameEvent_LobbyInvitesUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_LobbyInvitesUpdated: %v", m)
			return nil
		}
		if cbs := ge.onLobbyInvitesUpdated; cbs != nil {
			msg := &GameEventLobbyInvitesUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 268: // EGameEvent_CustomGameModeListUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_CustomGameModeListUpdated: %v", m)
			return nil
		}
		if cbs := ge.onCustomGameModeListUpdated; cbs != nil {
			msg := &GameEventCustomGameModeListUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 269: // EGameEvent_CustomGameLobbyListUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_CustomGameLobbyListUpdated: %v", m)
			return nil
		}
		if cbs := ge.onCustomGameLobbyListUpdated; cbs != nil {
			msg := &GameEventCustomGameLobbyListUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 270: // EGameEvent_FriendLobbyListUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_FriendLobbyListUpdated: %v", m)
			return nil
		}
		if cbs := ge.onFriendLobbyListUpdated; cbs != nil {
			msg := &GameEventFriendLobbyListUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 271: // EGameEvent_DotaTeamPlayerListChanged
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaTeamPlayerListChanged: %v", m)
			return nil
		}
		if cbs := ge.onDotaTeamPlayerListChanged; cbs != nil {
			msg := &GameEventDotaTeamPlayerListChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 272: // EGameEvent_DotaPlayerDetailsChanged
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaPlayerDetailsChanged: %v", m)
			return nil
		}
		if cbs := ge.onDotaPlayerDetailsChanged; cbs != nil {
			msg := &GameEventDotaPlayerDetailsChanged{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 273: // EGameEvent_PlayerProfileStatsUpdated
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_PlayerProfileStatsUpdated: %v", m)
			return nil
		}
		if cbs := ge.onPlayerProfileStatsUpdated; cbs != nil {
			msg := &GameEventPlayerProfileStatsUpdated{}
			msg.AccountId = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 274: // EGameEvent_CustomGamePlayerCountUpdated
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_CustomGamePlayerCountUpdated: %v", m)
			return nil
		}
		if cbs := ge.onCustomGamePlayerCountUpdated; cbs != nil {
			msg := &GameEventCustomGamePlayerCountUpdated{}
			msg.CustomGameId = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 275: // EGameEvent_CustomGameFriendsPlayedUpdated
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_CustomGameFriendsPlayedUpdated: %v", m)
			return nil
		}
		if cbs := ge.onCustomGameFriendsPlayedUpdated; cbs != nil {
			msg := &GameEventCustomGameFriendsPlayedUpdated{}
			msg.CustomGameId = m.GetKeys()[0].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 276: // EGameEvent_CustomGamesFriendsPlayUpdated
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_CustomGamesFriendsPlayUpdated: %v", m)
			return nil
		}
		if cbs := ge.onCustomGamesFriendsPlayUpdated; cbs != nil {
			msg := &GameEventCustomGamesFriendsPlayUpdated{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 277: // EGameEvent_DotaPlayerUpdateAssignedHero
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaPlayerUpdateAssignedHero: %v", m)
			return nil
		}
		if cbs := ge.onDotaPlayerUpdateAssignedHero; cbs != nil {
			msg := &GameEventDotaPlayerUpdateAssignedHero{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 278: // EGameEvent_DotaPlayerHeroSelectionDirty
		if len(m.GetKeys()) < 0 {
			_debugf("short EGameEvent_DotaPlayerHeroSelectionDirty: %v", m)
			return nil
		}
		if cbs := ge.onDotaPlayerHeroSelectionDirty; cbs != nil {
			msg := &GameEventDotaPlayerHeroSelectionDirty{}

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 279: // EGameEvent_DotaNpcGoalReached
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_DotaNpcGoalReached: %v", m)
			return nil
		}
		if cbs := ge.onDotaNpcGoalReached; cbs != nil {
			msg := &GameEventDotaNpcGoalReached{}
			msg.NpcEntindex = m.GetKeys()[0].GetValShort()
			msg.GoalEntindex = m.GetKeys()[1].GetValShort()
			msg.NextGoalEntindex = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 280: // EGameEvent_DotaPlayerSelectedCustomTeam
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_DotaPlayerSelectedCustomTeam: %v", m)
			return nil
		}
		if cbs := ge.onDotaPlayerSelectedCustomTeam; cbs != nil {
			msg := &GameEventDotaPlayerSelectedCustomTeam{}
			msg.PlayerId = m.GetKeys()[0].GetValShort()
			msg.TeamId = m.GetKeys()[1].GetValShort()
			msg.Success = m.GetKeys()[2].GetValBool()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 281: // EGameEvent_HltvStatus
		if len(m.GetKeys()) < 4 {
			_debugf("short EGameEvent_HltvStatus: %v", m)
			return nil
		}
		if cbs := ge.onHltvStatus; cbs != nil {
			msg := &GameEventHltvStatus{}
			msg.Clients = m.GetKeys()[0].GetValLong()
			msg.Slots = m.GetKeys()[1].GetValLong()
			msg.Proxies = m.GetKeys()[2].GetValShort()
			msg.Master = m.GetKeys()[3].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 282: // EGameEvent_HltvCameraman
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_HltvCameraman: %v", m)
			return nil
		}
		if cbs := ge.onHltvCameraman; cbs != nil {
			msg := &GameEventHltvCameraman{}
			msg.Index = m.GetKeys()[0].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 283: // EGameEvent_HltvRankCamera
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_HltvRankCamera: %v", m)
			return nil
		}
		if cbs := ge.onHltvRankCamera; cbs != nil {
			msg := &GameEventHltvRankCamera{}
			msg.Index = m.GetKeys()[0].GetValByte()
			msg.Rank = m.GetKeys()[1].GetValFloat()
			msg.Target = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 284: // EGameEvent_HltvRankEntity
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_HltvRankEntity: %v", m)
			return nil
		}
		if cbs := ge.onHltvRankEntity; cbs != nil {
			msg := &GameEventHltvRankEntity{}
			msg.Index = m.GetKeys()[0].GetValShort()
			msg.Rank = m.GetKeys()[1].GetValFloat()
			msg.Target = m.GetKeys()[2].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 285: // EGameEvent_HltvFixed
		if len(m.GetKeys()) < 8 {
			_debugf("short EGameEvent_HltvFixed: %v", m)
			return nil
		}
		if cbs := ge.onHltvFixed; cbs != nil {
			msg := &GameEventHltvFixed{}
			msg.Posx = m.GetKeys()[0].GetValLong()
			msg.Posy = m.GetKeys()[1].GetValLong()
			msg.Posz = m.GetKeys()[2].GetValLong()
			msg.Theta = m.GetKeys()[3].GetValShort()
			msg.Phi = m.GetKeys()[4].GetValShort()
			msg.Offset = m.GetKeys()[5].GetValShort()
			msg.Fov = m.GetKeys()[6].GetValFloat()
			msg.Target = m.GetKeys()[7].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 286: // EGameEvent_HltvChase
		if len(m.GetKeys()) < 7 {
			_debugf("short EGameEvent_HltvChase: %v", m)
			return nil
		}
		if cbs := ge.onHltvChase; cbs != nil {
			msg := &GameEventHltvChase{}
			msg.Target1 = m.GetKeys()[0].GetValShort()
			msg.Target2 = m.GetKeys()[1].GetValShort()
			msg.Distance = m.GetKeys()[2].GetValShort()
			msg.Theta = m.GetKeys()[3].GetValShort()
			msg.Phi = m.GetKeys()[4].GetValShort()
			msg.Inertia = m.GetKeys()[5].GetValByte()
			msg.Ineye = m.GetKeys()[6].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 287: // EGameEvent_HltvMessage
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_HltvMessage: %v", m)
			return nil
		}
		if cbs := ge.onHltvMessage; cbs != nil {
			msg := &GameEventHltvMessage{}
			msg.Text = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 288: // EGameEvent_HltvTitle
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_HltvTitle: %v", m)
			return nil
		}
		if cbs := ge.onHltvTitle; cbs != nil {
			msg := &GameEventHltvTitle{}
			msg.Text = m.GetKeys()[0].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 289: // EGameEvent_HltvChat
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_HltvChat: %v", m)
			return nil
		}
		if cbs := ge.onHltvChat; cbs != nil {
			msg := &GameEventHltvChat{}
			msg.Name = m.GetKeys()[0].GetValString()
			msg.Text = m.GetKeys()[1].GetValString()
			msg.SteamID = m.GetKeys()[2].GetValUint64()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 290: // EGameEvent_HltvVersioninfo
		if len(m.GetKeys()) < 1 {
			_debugf("short EGameEvent_HltvVersioninfo: %v", m)
			return nil
		}
		if cbs := ge.onHltvVersioninfo; cbs != nil {
			msg := &GameEventHltvVersioninfo{}
			msg.Version = m.GetKeys()[0].GetValByte()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 291: // EGameEvent_DotaChaseHero
		if len(m.GetKeys()) < 9 {
			_debugf("short EGameEvent_DotaChaseHero: %v", m)
			return nil
		}
		if cbs := ge.onDotaChaseHero; cbs != nil {
			msg := &GameEventDotaChaseHero{}
			msg.Target1 = m.GetKeys()[0].GetValShort()
			msg.Target2 = m.GetKeys()[1].GetValShort()
			msg.Type = m.GetKeys()[2].GetValByte()
			msg.Priority = m.GetKeys()[3].GetValShort()
			msg.Gametime = m.GetKeys()[4].GetValFloat()
			msg.Highlight = m.GetKeys()[5].GetValBool()
			msg.Target1playerid = m.GetKeys()[6].GetValByte()
			msg.Target2playerid = m.GetKeys()[7].GetValByte()
			msg.Eventtype = m.GetKeys()[8].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 292: // EGameEvent_DotaCombatlog
		if len(m.GetKeys()) < 19 {
			_debugf("short EGameEvent_DotaCombatlog: %v", m)
			return nil
		}
		if cbs := ge.onDotaCombatlog; cbs != nil {
			msg := &GameEventDotaCombatlog{}
			msg.Type = m.GetKeys()[0].GetValByte()
			msg.Sourcename = m.GetKeys()[1].GetValShort()
			msg.Targetname = m.GetKeys()[2].GetValShort()
			msg.Attackername = m.GetKeys()[3].GetValShort()
			msg.Inflictorname = m.GetKeys()[4].GetValShort()
			msg.Attackerillusion = m.GetKeys()[5].GetValBool()
			msg.Targetillusion = m.GetKeys()[6].GetValBool()
			msg.Value = m.GetKeys()[7].GetValShort()
			msg.Health = m.GetKeys()[8].GetValShort()
			msg.Timestamp = m.GetKeys()[9].GetValFloat()
			msg.Targetsourcename = m.GetKeys()[10].GetValShort()
			msg.Timestampraw = m.GetKeys()[11].GetValFloat()
			msg.Attackerhero = m.GetKeys()[12].GetValBool()
			msg.Targethero = m.GetKeys()[13].GetValBool()
			msg.AbilityToggleOn = m.GetKeys()[14].GetValBool()
			msg.AbilityToggleOff = m.GetKeys()[15].GetValBool()
			msg.AbilityLevel = m.GetKeys()[16].GetValShort()
			msg.GoldReason = m.GetKeys()[17].GetValShort()
			msg.XpReason = m.GetKeys()[18].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 293: // EGameEvent_DotaGameStateChange
		if len(m.GetKeys()) < 2 {
			_debugf("short EGameEvent_DotaGameStateChange: %v", m)
			return nil
		}
		if cbs := ge.onDotaGameStateChange; cbs != nil {
			msg := &GameEventDotaGameStateChange{}
			msg.OldState = m.GetKeys()[0].GetValShort()
			msg.NewState = m.GetKeys()[1].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 294: // EGameEvent_DotaPlayerPickHero
		if len(m.GetKeys()) < 3 {
			_debugf("short EGameEvent_DotaPlayerPickHero: %v", m)
			return nil
		}
		if cbs := ge.onDotaPlayerPickHero; cbs != nil {
			msg := &GameEventDotaPlayerPickHero{}
			msg.Player = m.GetKeys()[0].GetValShort()
			msg.Heroindex = m.GetKeys()[1].GetValShort()
			msg.Hero = m.GetKeys()[2].GetValString()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	case 295: // EGameEvent_DotaTeamKillCredit
		if len(m.GetKeys()) < 4 {
			_debugf("short EGameEvent_DotaTeamKillCredit: %v", m)
			return nil
		}
		if cbs := ge.onDotaTeamKillCredit; cbs != nil {
			msg := &GameEventDotaTeamKillCredit{}
			msg.KillerUserid = m.GetKeys()[0].GetValShort()
			msg.VictimUserid = m.GetKeys()[1].GetValShort()
			msg.Teamnumber = m.GetKeys()[2].GetValShort()
			msg.Herokills = m.GetKeys()[3].GetValShort()

			for _, fn := range cbs {
				if err := fn(msg); err != nil {
					return err
				}
			}
		}
		return nil

	}

	_panicf("unknown message %d", m.GetEventid())
	return nil
}
