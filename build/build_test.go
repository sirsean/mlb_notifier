package build

import (
    "testing"
    "strings"
)

func TestOneGame(t *testing.T) {
    games, _ := BuildGames(strings.NewReader(onlyOneGame))
    if (len(games) != 1) {
        t.Errorf("Should have 1 game")
    }
}

func TestPitchers(t *testing.T) {
    boxscore, _ := BuildBoxscore(strings.NewReader(finalBoxscoreWithShutoutThru7))
    if len(boxscore.AwayPitchers) == 0 {
        t.Errorf("Shouldn't be 0 AwayPitchers")
    }
    if len(boxscore.HomePitchers) != 3 {
        t.Errorf("Should be 3 HomePitchers")
    }
    pitcher := boxscore.HomePitchers[0]
    if pitcher.Name != "Tanner Roark" {
        t.Errorf("Pitcher's name is Tanner Roark")
    }
    if pitcher.Outs != 21 {
        t.Errorf("Should have 21 outs")
    }
    if pitcher.R != 0 {
        t.Errorf("Should have 0 runs")
    }
}

var finalBoxscoreWithShutoutThru7 = `{ "subject": "boxscore", "copyright" : "Copyright 2013 MLB Advanced Media, L.P.  Use of any content on this page acknowledges agreement to the terms posted here http://gdx.mlb.com/components/copyright.txt", "data" : {"boxscore":{"home_sname":"Washington","away_fname":"Atlanta Braves","game_info":"\n      <wild_pitches>Garcia, F.<\/wild_pitches><br/>\n      <pitches_to_strikes>Garcia, F 84-56, Walden 20-14, Roark 101-67, Stammen 18-10, Soriano, R 17-12.<\/pitches_to_strikes><br/>\n      <groundouts_to_flyouts>Garcia, F 6-2, Walden 2-1, Roark 7-2, Stammen 0-0, Soriano, R 2-0.<\/groundouts_to_flyouts><br/>\n      <batters_faced>Garcia, F 27, Walden 7, Roark 24, Stammen 4, Soriano, R 5.<\/batters_faced><br/>\n      <umpires>\n         <umpire id=\"490319\" position=\"HP\" name=\"Jordan Baker\"><\/umpire>\n         <umpire id=\"427044\" position=\"1B\" name=\"CB Bucknor\"><\/umpire>\n         <umpire id=\"427457\" position=\"2B\" name=\"Dale Scott\"><\/umpire>\n         <umpire id=\"427108\" position=\"3B\" name=\"Dana DeMuth\"><\/umpire>\n      <\/umpires>\n      <weather>63 degrees, clear<\/weather>\n      <wind>3 mph, In from RF<\/wind>\n      <time>2:25<\/time>\n      <attendance>28,369<\/attendance>\n      <wind>Nationals Park<\/wind>\n      <wind>September 17, 2013<\/wind>","home_loss":"70","venue_name":"Nationals Park","linescore":{"home_team_runs":"4","home_team_errors":"1","home_team_hits":"11","inning_line_score":[{"home":"0","away":"0","inning":"1"},{"home":"1","away":"0","inning":"2"},{"home":"0","away":"0","inning":"3"},{"home":"0","away":"0","inning":"4"},{"home":"0","away":"0","inning":"5"},{"home":"0","away":"0","inning":"6"},{"home":"0","away":"0","inning":"7"},{"home":"3","away":"0","inning":"8"},{"home":"x","away":"0","inning":"9"}],"away_team_errors":"1","away_team_runs":"0","away_team_hits":"5"},"home_sport_code":"mlb","away_wins":"89","game_pk":"348993","status_ind":"F","date":"September 17, 2013","pitching":[{"hr":"1","so":"6","r":"4","pitcher":[{"hr":"0","s_er":"37","loss":"true","np":"84","name_display_first_last":"Freddy Garcia","s_h":"77","era":"4.52","bs":"0","pos":"P","id":"150119","name":"Garcia, F","bf":"27","sv":"0","bb":"2","note":"(L, 4-7)","hld":"0","so":"6","l":"7","s_so":"39","h":"7","s_ip":"73.2","w":"4","s_bb":"16","s_r":"38","s":"56","r":"1","er":"1","out":"21"},{"hr":"1","s_er":"17","np":"20","name_display_first_last":"Jordan Walden","s_h":"37","era":"3.30","bs":"2","pos":"P","id":"477229","name":"Walden","bf":"7","sv":"1","bb":"0","hld":"14","so":"0","l":"3","s_so":"52","h":"4","s_ip":"46.1","w":"4","s_bb":"14","s_r":"17","s":"14","r":"3","er":"3","out":"3"}],"bf":"34","era":"3.21","team_flag":"away","bb":"2","er":"4","h":"11","out":"24"},{"hr":"0","so":"9","r":"0","pitcher":[{"hr":"0","s_er":"5","np":"101","name_display_first_last":"Tanner Roark","s_h":"26","era":"1.08","bs":"0","pos":"P","id":"543699","name":"Roark","bf":"24","sv":"0","bb":"1","note":"(W, 7-0)","hld":"1","so":"6","l":"0","s_so":"32","win":"true","h":"2","s_ip":"41.2","w":"7","s_bb":"9","s_r":"6","s":"67","r":"0","er":"0","out":"21"},{"hr":"0","s_er":"24","np":"18","name_display_first_last":"Craig Stammen","s_h":"76","era":"2.75","bs":"1","pos":"P","id":"489334","name":"Stammen","bf":"4","sv":"0","bb":"0","note":"(H, 6)","hld":"6","so":"3","l":"6","s_so":"77","h":"1","s_ip":"78.2","w":"7","s_bb":"26","s_r":"29","s":"10","r":"0","er":"0","out":"3"},{"hr":"0","s_er":"23","np":"17","name_display_first_last":"Rafael Soriano","s_h":"64","era":"3.25","bs":"6","pos":"P","id":"400089","name":"Soriano, R","bf":"5","sv":"41","bb":"0","hld":"0","so":"0","l":"3","s_so":"48","h":"2","s_ip":"63.2","w":"2","s_bb":"14","s_r":"24","s":"12","r":"0","er":"0","out":"3"}],"bf":"33","era":"3.63","team_flag":"home","bb":"1","er":"0","h":"5","out":"27"}],"home_id":"120","away_loss":"62","batting":[{"hr":"1","d":"1","da":"10","so":"6","batter":[{"hr":"0","sac":"0","name_display_first_last":"Denard Span","s_h":"162","s_hr":"4","s_rbi":"42","pos":"CF","id":"452655","rbi":"0","bo":"100","lob":"1","name":"Span","avg":".283","bb":"0","fldg":"1.000","hbp":"0","d":"0","e":"0","so":"0","a":"0","s_so":"71","sf":"0","h":"1","cs":"0","s_bb":"40","s_r":"68","t":"0","ao":"0","r":"0","sb":"1","po":"2","ab":"4","go":"3"},{"hr":"1","sac":"0","name_display_first_last":"Ryan Zimmerman","s_h":"148","s_hr":"25","s_rbi":"74","pos":"3B","id":"475582","rbi":"1","bo":"200","lob":"1","name":"Zimmerman","avg":".281","bb":"0","fldg":"1.000","hbp":"0","d":"0","e":"0","so":"2","a":"1","s_so":"123","sf":"0","h":"2","cs":"0","s_bb":"58","s_r":"81","t":"0","ao":"0","r":"1","sb":"0","po":"1","ab":"4"},{"hr":"0","sac":"0","name_display_first_last":"Jayson Werth","s_h":"137","s_hr":"23","s_rbi":"74","pos":"RF","id":"150029","rbi":"0","bo":"300","lob":"1","name":"Werth","avg":".320","bb":"1","fldg":"1.000","hbp":"0","d":"0","e":"0","so":"0","a":"0","s_so":"95","sf":"0","h":"0","cs":"0","s_bb":"53","s_r":"78","t":"0","ao":"3","r":"0","sb":"0","po":"1","ab":"3"},{"hr":"0","sac":"0","name_display_first_last":"Bryce Harper","s_h":"108","s_hr":"19","s_rbi":"51","pos":"LF","id":"547180","rbi":"0","bo":"400","lob":"2","name":"Harper","avg":".280","bb":"0","fldg":"1.000","hbp":"0","d":"0","e":"0","so":"1","a":"0","s_so":"88","sf":"0","h":"2","cs":"0","s_bb":"60","s_r":"68","t":"0","ao":"0","r":"2","sb":"1","po":"2","ab":"4","go":"1"},{"hr":"0","sac":"0","name_display_first_last":"Ian Desmond","s_h":"161","s_hr":"20","s_rbi":"79","pos":"SS","id":"435622","rbi":"1","bo":"500","lob":"1","name":"Desmond","avg":".285","bb":"0","fldg":".833","hbp":"0","d":"1","e":"1","so":"1","a":"5","s_so":"138","sf":"0","h":"2","cs":"0","s_bb":"41","s_r":"76","t":"0","ao":"1","r":"1","sb":"0","po":"0","ab":"4"},{"hr":"0","sac":"0","name_display_first_last":"Adam LaRoche","s_h":"117","s_hr":"20","s_rbi":"62","pos":"1B","id":"425560","rbi":"1","bo":"600","lob":"1","name":"LaRoche","avg":".239","bb":"1","fldg":"1.000","hbp":"0","d":"0","e":"0","so":"1","a":"1","s_so":"127","sf":"0","h":"1","cs":"0","s_bb":"66","s_r":"69","t":"0","ao":"1","r":"0","sb":"0","po":"7","ab":"3"},{"hr":"0","sac":"0","name_display_first_last":"Steve Lombardozzi","s_h":"69","s_hr":"2","s_rbi":"21","pos":"2B","id":"543459","rbi":"1","bo":"700","lob":"1","name":"Lombardozzi","avg":".255","bb":"0","fldg":"1.000","hbp":"0","d":"0","e":"0","so":"0","a":"1","s_so":"33","sf":"0","h":"2","cs":"0","s_bb":"8","s_r":"24","t":"0","ao":"0","r":"0","sb":"0","po":"5","ab":"4","go":"2"},{"hr":"0","gidp":"1","sac":"0","name_display_first_last":"Jhonatan Solano","s_h":"7","s_hr":"0","s_rbi":"2","pos":"C","id":"500207","rbi":"0","bo":"800","lob":"4","name":"Solano, J","avg":".167","bb":"0","fldg":"1.000","hbp":"0","d":"0","e":"0","so":"0","a":"1","s_so":"6","sf":"0","h":"0","cs":"0","s_bb":"2","s_r":"2","t":"0","ao":"2","r":"0","sb":"0","po":"8","ab":"3","go":"2"},{"hr":"0","sac":"0","name_display_first_last":"Tanner Roark","s_h":"3","s_hr":"0","s_rbi":"1","pos":"P","id":"543699","rbi":"0","bo":"900","lob":"0","name":"Roark","avg":".300","bb":"0","fldg":"1.000","hbp":"0","d":"0","e":"0","so":"1","a":"0","s_so":"4","sf":"0","h":"1","cs":"0","s_bb":"0","s_r":"0","t":"0","ao":"0","r":"0","sb":"0","po":"1","ab":"2"},{"hr":"0","sac":"0","name_display_first_last":"Chad Tracy","s_h":"23","s_hr":"3","s_rbi":"10","pos":"PH","id":"429710","rbi":"0","bo":"901","lob":"0","name":"Tracy","avg":".187","bb":"0","fldg":".000","note":"a-","hbp":"0","d":"0","e":"0","so":"0","a":"0","s_so":"25","sf":"0","h":"0","cs":"0","s_bb":"6","s_r":"5","t":"0","ao":"0","r":"0","sb":"0","po":"0","ab":"1","go":"1"},{"hr":"0","sac":"0","name_display_first_last":"Craig Stammen","s_h":"0","s_hr":"0","s_rbi":"0","pos":"P","id":"489334","rbi":"0","bo":"902","lob":"0","name":"Stammen","avg":".000","bb":"0","fldg":".000","hbp":"0","d":"0","e":"0","so":"0","a":"0","s_so":"2","sf":"0","h":"0","cs":"0","s_bb":"1","s_r":"0","t":"0","ao":"0","r":"0","sb":"0","po":"0","ab":"0"},{"hr":"0","sac":"0","name_display_first_last":"Rafael Soriano","s_h":"0","s_hr":"0","s_rbi":"0","pos":"P","id":"400089","rbi":"0","bo":"903","lob":"0","name":"Soriano, R","avg":".000","bb":"0","fldg":"1.000","hbp":"0","d":"0","e":"0","so":"0","a":"1","s_so":"0","sf":"0","h":"0","cs":"0","s_bb":"0","s_r":"0","t":"0","ao":"0","r":"0","sb":"0","po":"0","ab":"0"}],"h":"11","text_data":"\n         <doubles>Desmond (37, Walden).<\/doubles><br/>\n         <home_runs>Zimmerman (25, 8th inning off Walden, 0 on, 1 out).<\/home_runs><br/>\n         <total_bases>Zimmerman 5; Lombardozzi 2; Span; Desmond 3; Roark; Harper 2; LaRoche.<\/total_bases><br/>\n         <rbi>Lombardozzi (21), Zimmerman (74), Desmond (79), LaRoche (62).<\/rbi><br/>\n         <two_out_rbi>Desmond; LaRoche.<\/two_out_rbi><br/>\n         <risp_lob_two_out>Harper; Werth.<\/risp_lob_two_out><br/>\n         <grounded_into_dp>Solano, J.<\/grounded_into_dp><br/>\n         <team_risp>2-for-6.<\/team_risp>\n         <team_lob>6.<\/team_lob>\n         <stolen_bases>Harper (11, 2nd base off Garcia, F/Laird), Span (18, 2nd base off Garcia, F/Laird).<\/stolen_bases><br/>\n         <errors>Desmond (19, throw).<\/errors><br/>\n         <double_plays>(Zimmerman-Lombardozzi).<\/double_plays><br/><br/>","rbi":"4","lob":"12","t":"0","r":"4","team_flag":"home","ab":"32","po":"27","avg":".254","bb":"2","note":"\n         <pinch_hitters>a-Grounded out for Roark in the 7th. <\/pinch_hitters>"},{"hr":"0","d":"0","da":"12","so":"9","batter":[{"hr":"0","sac":"0","name_display_first_last":"Jordan Schafer","s_h":"54","s_hr":"3","s_rbi":"19","pos":"CF","id":"457788","rbi":"0","bo":"100","lob":"1","name":"Schafer, J","avg":".249","bb":"1","fldg":"1.000","hbp":"0","d":"0","e":"0","so":"1","a":"0","s_so":"68","sf":"0","h":"0","cs":"0","s_bb":"27","s_r":"30","t":"0","ao":"0","r":"0","sb":"0","po":"1","ab":"3","go":"2"},{"hr":"0","sac":"0","name_display_first_last":"Justin Upton","s_h":"134","s_hr":"25","s_rbi":"65","pos":"RF","id":"457708","rbi":"0","bo":"200","lob":"1","name":"Upton, J","avg":".259","bb":"0","fldg":"1.000","hbp":"0","d":"0","e":"0","so":"1","a":"0","s_so":"152","sf":"0","h":"1","cs":"0","s_bb":"72","s_r":"90","t":"0","ao":"1","r":"0","sb":"0","po":"1","ab":"4","go":"1"},{"hr":"0","sac":"0","name_display_first_last":"Freddie Freeman","s_h":"162","s_hr":"21","s_rbi":"100","pos":"1B","id":"518692","rbi":"0","bo":"300","lob":"1","name":"Freeman, F","avg":".313","bb":"0","fldg":"1.000","hbp":"0","d":"0","e":"0","so":"1","a":"3","s_so":"112","sf":"0","h":"1","cs":"0","s_bb":"61","s_r":"81","t":"0","ao":"1","r":"0","sb":"0","po":"6","ab":"4","go":"1"},{"hr":"0","sac":"0","name_display_first_last":"Evan Gattis","s_h":"73","s_hr":"20","s_rbi":"59","pos":"LF","id":"594828","rbi":"0","bo":"400","lob":"2","name":"Gattis","avg":".235","bb":"0","fldg":"1.000","hbp":"0","d":"0","e":"0","so":"1","a":"0","s_so":"75","sf":"0","h":"0","cs":"0","s_bb":"20","s_r":"41","t":"0","ao":"3","r":"0","sb":"0","po":"2","ab":"4"},{"hr":"0","sac":"0","name_display_first_last":"Chris Johnson","s_h":"157","s_hr":"10","s_rbi":"64","pos":"3B","id":"453400","rbi":"0","bo":"500","lob":"1","name":"Johnson, C","avg":".327","bb":"0","fldg":"1.000","hbp":"0","d":"0","e":"0","so":"0","a":"1","s_so":"105","sf":"0","h":"0","cs":"0","s_bb":"27","s_r":"50","t":"0","ao":"2","r":"0","sb":"0","po":"1","ab":"4","go":"2"},{"hr":"0","sac":"0","name_display_first_last":"Gerald Laird","s_h":"28","s_hr":"1","s_rbi":"13","pos":"C","id":"408042","rbi":"0","bo":"600","lob":"2","name":"Laird","avg":".262","bb":"0","fldg":"1.000","hbp":"0","d":"0","e":"0","so":"1","a":"0","s_so":"20","sf":"0","h":"1","cs":"0","s_bb":"11","s_r":"8","t":"0","ao":"1","r":"0","sb":"0","po":"7","ab":"4","go":"1"},{"hr":"0","sac":"0","name_display_first_last":"Dan Uggla","s_h":"78","s_hr":"21","s_rbi":"54","pos":"2B","id":"462564","rbi":"0","bo":"700","lob":"0","name":"Uggla","avg":".183","bb":"0","fldg":"1.000","hbp":"0","d":"0","e":"0","so":"2","a":"5","s_so":"160","sf":"0","h":"1","cs":"0","s_bb":"71","s_r":"59","t":"0","ao":"0","r":"0","sb":"0","po":"2","ab":"3"},{"hr":"0","sac":"0","name_display_first_last":"Andrelton Simmons","s_h":"140","s_hr":"15","s_rbi":"52","pos":"SS","id":"592743","rbi":"0","bo":"800","lob":"2","name":"Simmons","avg":".247","bb":"0","fldg":".857","hbp":"0","d":"0","e":"1","so":"1","a":"3","s_so":"51","sf":"0","h":"0","cs":"0","s_bb":"38","s_r":"70","t":"0","ao":"0","r":"0","sb":"0","po":"3","ab":"3","go":"2"},{"hr":"0","sac":"0","name_display_first_last":"Freddy Garcia","s_h":"0","s_hr":"0","s_rbi":"0","pos":"P","id":"150119","rbi":"0","bo":"900","lob":"0","name":"Garcia, F","avg":".000","bb":"0","fldg":"1.000","hbp":"0","d":"0","e":"0","so":"1","a":"0","s_so":"3","sf":"0","h":"0","cs":"0","s_bb":"0","s_r":"0","t":"0","ao":"1","r":"0","sb":"0","po":"1","ab":"2"},{"hr":"0","sac":"0","name_display_first_last":"Elliot Johnson","s_h":"47","s_hr":"2","s_rbi":"14","pos":"PH","id":"471107","rbi":"0","bo":"901","lob":"0","name":"Johnson, E","avg":".205","bb":"0","fldg":".000","note":"a-","hbp":"0","d":"0","e":"0","so":"0","a":"0","s_so":"61","sf":"0","h":"1","cs":"0","s_bb":"11","s_r":"24","t":"0","ao":"0","r":"0","sb":"1","po":"0","ab":"1"},{"hr":"0","sac":"0","name_display_first_last":"Jordan Walden","s_h":"0","s_hr":"0","s_rbi":"0","pos":"P","id":"477229","rbi":"0","bo":"902","lob":"0","name":"Walden","avg":".000","bb":"0","fldg":".000","hbp":"0","d":"0","e":"0","so":"0","a":"0","s_so":"0","sf":"0","h":"0","cs":"0","s_bb":"0","s_r":"0","t":"0","ao":"0","r":"0","sb":"0","po":"0","ab":"0"}],"h":"5","text_data":"\n         <total_bases>Johnson, E; Freeman, F; Laird; Uggla; Upton, J.<\/total_bases><br/>\n         <risp_lob_two_out>Simmons; Schafer, J; Laird.<\/risp_lob_two_out><br/>\n         <team_risp>0-for-4.<\/team_risp>\n         <team_lob>6.<\/team_lob>\n         <stolen_bases>Johnson, E (21, 2nd base off Stammen/Solano, J).<\/stolen_bases><br/>\n         <errors>Simmons (13, throw).<\/errors><br/>\n         <double_plays>2 (Simmons-Freeman, F, Simmons-Uggla-Freeman, F).<\/double_plays><br/><br/>","rbi":"0","lob":"10","t":"0","r":"0","team_flag":"away","ab":"32","po":"24","avg":".248","bb":"1","note":"\n         <pinch_hitters>a-Singled for Garcia, F in the 8th. <\/pinch_hitters>"}],"away_sname":"Atlanta","game_id":"2013/09/17/atlmlb-wasmlb-2","home_team_code":"was","venue_id":"3309","home_wins":"81","away_team_code":"atl","away_id":"144","home_fname":"Washington Nationals"}}}`
var onlyOneGame = `{ "subject": "master_scoreboard_mlb_2013_09_30", "copyright" : "Copyright 2013 MLB Advanced Media, L.P.  Use of any content on this page acknowledges agreement to the terms posted here http://gdx.mlb.com/components/copyright.txt", "data" : {"games":{"next_day_date":"2013-10-01","modified_date":"2013-10-01T16:02:49Z","month":"09","year":"2013","game":{"game_type":"R","double_header_sw":"N","away_time":"8:07","broadcast":{"away":{"tv":"TBS","radio":{}},"home":{"tv":"TBS","radio":{}}},"time":"8:07","home_time":"7:07","home_team_name":"Rangers","description":"AL Tiebreaker Game","original_date":"2013/09/30","home_team_city":"Texas","venue_id":"13","gameday_sw":"P","away_win":"92","home_games_back_wildcard":"1.0","save_pitcher":{"id":"","last":"","saves":"0","losses":"0","era":"0","name_display_roster":"","number":"","svo":"0","first":"","wins":"0"},"away_team_id":"139","tz_hm_lg_gen":"ET","status":{"top_inning":"N","s":"0","b":"0","reason":"","ind":"F","status":"Final","o":"3","inning":"9","inning_state":""},"home_loss":"72","home_games_back":"5.5","home_code":"tex","away_sport_code":"mlb","home_win":"91","time_hm_lg":"8:07","away_name_abbrev":"TB","league":"AA","time_zone_aw_lg":"-4","away_games_back":"5.5","home_file_code":"tex","game_data_directory":"/components/game/mlb/year_2013/month_09/day_30/gid_2013_09_30_tbamlb_texmlb_1","time_zone":"ET","away_league_id":"103","home_team_id":"140","day":"MON","time_aw_lg":"8:07","away_team_city":"Tampa Bay","tbd_flag":"N","tz_aw_lg_gen":"ET","away_code":"tba","winning_pitcher":{"id":"456034","last":"Price","losses":"8","era":"3.33","number":"14","name_display_roster":"Price","first":"David","wins":"10"},"game_media":{"media":[{"free":"NO","title":"TB @ TEX","thumbnail":"http://mediadownloads.mlb.com/mlbam/preview/tbatex_385292_th_7_preview.jpg","media_state":"media_archive","start":"2013-09-30T20:07:00-0400","has_mlbtv":"true","calendar_event_id":"14-385292-2013-09-30","type":"game"},{"headline":"Longoria's three-hit game","thumbnail":"http://mediadownloads.mlb.com/mlbam/2013/10/01/images/mlbf_31065127_th_7.jpg","content_id":"31065127","topic_id":"","type":"vpp"}]},"game_nbr":"1","time_date_aw_lg":"2013/09/30 8:07","away_games_back_wildcard":"-","scheduled_innings":"9","linescore":{"hr":{"home":"0","away":"1"},"e":{"home":"1","away":"0"},"so":{"home":"8","away":"4"},"r":{"home":"2","away":"5","diff":"3"},"sb":{"home":"1","away":"1"},"inning":[{"home":"0","away":"1"},{"home":"0","away":"0"},{"home":"1","away":"2"},{"home":"0","away":"0"},{"home":"0","away":"0"},{"home":"1","away":"1"},{"home":"0","away":"0"},{"home":"0","away":"0"},{"home":"0","away":"1"}],"h":{"home":"7","away":"7"}},"venue_w_chan_loc":"USTX0045","first_pitch_et":"","away_team_name":"Rays","home_runs":{"player":{"std_hr":"32","hr":"1","id":"446334","last":"Longoria","team_code":"tba","inning":"3","runners":"1","number":"3","name_display_roster":"Longoria","first":"Evan"}},"time_date_hm_lg":"2013/09/30 8:07","id":"2013/09/30/tbamlb-texmlb-1","home_name_abbrev":"TEX","tiebreaker_sw":"Y","ampm":"PM","home_division":"W","home_time_zone":"CT","away_time_zone":"ET","hm_lg_ampm":"PM","home_sport_code":"mlb","time_date":"2013/09/30 8:07","links":{"away_audio":"bam.media.launchPlayer({calendar_event_id:'14-385292-2013-09-30',media_type:'audio'})","wrapup":"/mlb/gameday/index.jsp?gid=2013_09_30_tbamlb_texmlb_1&mode=wrap&c_id=mlb","preview":"/mlb/gameday/index.jsp?gid=2013_09_30_tbamlb_texmlb_1&mode=preview&c_id=mlb","home_preview":"/mlb/gameday/index.jsp?gid=2013_09_30_tbamlb_texmlb_1&mode=preview&c_id=mlb","away_preview":"/mlb/gameday/index.jsp?gid=2013_09_30_tbamlb_texmlb_1&mode=preview&c_id=mlb","tv_station":"TBS","home_audio":"bam.media.launchPlayer({calendar_event_id:'14-385292-2013-09-30',media_type:'audio'})","mlbtv":"bam.media.launchPlayer({calendar_event_id:'14-385292-2013-09-30',media_type:'video'})"},"home_ampm":"PM","game_pk":"385292","venue":"Rangers Ballpark in Arlington","home_league_id":"103","video_thumbnail":"http://mediadownloads.mlb.com/mlbam/preview/tbatex_385292_th_7_preview.jpg","away_loss":"71","resume_date":"","away_file_code":"tb","losing_pitcher":{"id":"527048","last":"Perez","losses":"6","era":"3.62","number":"33","name_display_roster":"Perez, M","first":"Martin","wins":"10"},"aw_lg_ampm":"PM","video_thumbnails":{"thumbnail":[{"content":"http://mediadownloads.mlb.com/mlbam/preview/tbatex_385292_th_7_preview.jpg","height":"70","scenario":"7","width":"124"},{"content":"http://mediadownloads.mlb.com/mlbam/preview/tbatex_385292_th_37_preview.jpg","height":"90","scenario":"37","width":"160"}]},"time_zone_hm_lg":"-4","away_ampm":"PM","gameday":"2013_09_30_tbamlb_texmlb_1","away_division":"E"},"day":"30"}}}`