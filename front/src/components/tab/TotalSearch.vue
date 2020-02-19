<template>
  <div class="main_con">
    <p class="result"><em>"<span id="searchword">{{ keyword }}</span>"</em>에 대한 통합검색 결과 입니다. (<span id="total_cnt">{{ totalCnt }}</span>건)</p>
    <!--방송 바로보기-->
    <div class="direct_area" style="display:none;" id="direct_broad"></div>
    <!--bj프로필-->
    <div class="profile_area" v-if="profile">
      <h4 class="stit">BJ 프로필</h4>
      <div class="profile_wrap">
        <div class="detail">
          <div class="pr">
            <div class="img_thum">
              <img :src="profileList.img_file" width="200" height="150" alt="" onerror="//res.afreecatv.com/images/main_new/@thum_s_104x76.gif">
              <p>
                <a href="javascript:void(0);" @click="favorite(profileList.user_id)">
                  <img src="//res.afreecatv.com/images/aftv_search/btn_fa.gif" class="btn_f" alt="즐겨찾기">
                </a>
              </p>
            </div>
            <dl>
              <dt>
                <span>
                  <a href="javascript:void(0);" @click="goStation(profileList.user_id)" target="_blank">{{ profileList.user_nick + '(' + profileList.user_id + ')' }}</a>
                </span>
                <span class="ico_fan">팬</span>
              </dt>
              <dd>
                <em>출생</em>{{ profileList.born_year + '년생 (' + profileList.user_age + '세)' }}
              </dd>
              <dd class="award" v-if="profileList.career_award.length > 0">
                <em>경력/수상</em>{{ profileList.career_award[0].content }}
                <div class="award_layer" v-if="profileList.career_award.length > 1">
                  <img src="//res.afreecatv.com/images/aftv_search/btn_layer_dk.gif" class="btn_layer" alt="레이어" @click="awardListToggle">
                  <div class="award_list" id="divAwardList" style="display:none;">
                    <ul>
                      <li v-for="oData in profileList.career_award" :key="oData">{{ oData.content }}</li>
                    </ul>
                  </div>
                </div>
              </dd>
              <dd>
                <em>방송시간</em>{{ profileList.total_broad_time/3600 | numFormat }}시간
              </dd>
              <dd>
                <em>애청자</em>{{ profileList.fan_count | numFormat }}명
                <span class="bar">|</span>
                <em>누적시청자</em>{{ profileList.view_count | numFormat }}명
              </dd>
              <dd>
                <em>커뮤니티</em>
                <a href="javascript:void(0);" @click="goStation(profileList.user_id)" target="_blank">방송국,</a>
                <a href="javascript:void(0);" @click="goStation(profileList.user_id)" target="_blank">공식까페</a>
              </dd>
              <dd>
                <em>공지사항</em>
                <a href="javascript:void(0);" @click="goNotice(profileList.user_id, profileList.station_no, profileList.bbs_no, profileList.title_no)" target="_blank">{{ profileList.notice }}</a>
              </dd>
            </dl>
          </div>
          <div class="medal" id="divMedal">
            <dl>
              <dt>수상메달</dt>
                <dd v-for="oData in profileList.award_medal.slice(profileAward.start, profileAward.end)" v-bind:key="oData" :class="profileContent[oData.type].class">{{ profileContent[oData.type].content }}</dd>
              </dl><div class="medal_pg">
                <span class="pg_no">
                  <strong>{{ profileAward.curPage }}</strong>/{{ profileAward.totalPage }}
                </span> <!--_n 비활성화-->
                <span class="btn_page">
                  <a href="javascript:;" :class="[profileAward.prevActive ? 'prev' : 'prev_n' , 'btn']" @click="moveMedal('prev')">이전</a>
                  <a href="javascript:;" :class="[profileAward.nextActive ? 'next' : 'next_n', 'btn']" @click="moveMedal('next')">다음</a>
                </span>
              </div>
            </div>
        </div>
        <a href="javascript:;" class="profile_edit" @click="goProfile">내 프로필 수정</a>
      </div>
    </div>
    <div id="broad_order_sec" class="tit_area" style="">
      <h4 id="broad_title_text"><a href="javascript:;">생방송</a><span id="broad_cnt">({{ liveCnt }})</span><img src="//res.afreecatv.com/images/aftv_search/ico_live.gif" alt="live"></h4>
      <div class="select_v">
          <a href="javascript:;" class="v" id="broad_order_text"><strong>정확도순</strong><em></em></a>
          <ul class="sub_list" style="display: none;">
          <li><a href="javascript:;">정확도순</a></li>
          <li><a href="javascript:;">시청인원순</a></li>
          <li><a href="javascript:;">최신순</a></li>
          <li><a href="javascript:;">고화질순</a></li>
          <li><a href="javascript:;" class="last">고해상도순</a></li>
          </ul>
      </div>
    </div>
  </div>
</template>

<script>
import { VueSlideToggle } from 'vue-slide-toggle'

export default {
  name: 'TotalSearch',
  components: {
    VueSlideToggle
  }
  data () {
    return {
      totalCnt: 0,
      keyword: '',
      profile: false,
      profileList: '',
      profileAward: {
        expose: false,
        start: 0,
        end: 5,
        curPage: 1,
        totalPage: 1,
        prevActive: false,
        nextActive: false
      },
      profileContent: ''
    }
  },
  filters: {
    numFormat (value) {
      let num = Number(value)
      return num.toFixed(0).replace(/(\d)(?=(\d{3})+(?:\.\d+)?$)/g, '$1,')
    }
  },
  created: function () {
    this.setting()
    this.profileTheme()
    this.awardContentSetting()
  },
  methods: {
    setting () {
      this.keyword = this.$route.query.keyword
      this.$store.dispatch('callChangeTab', { tab: 'totalSearch' })
    },
    profileTheme () {
      let params = {}
      params['m'] = 'profileTheme'
      params['v'] = '1.0'
      params['t'] = 'json'
      params['c'] = 'EUC-KR'
      params['d'] = this.keyword
      params['w'] = 'webk'

      this.$http.get('//scketc.afreecatv.com/api.php', {
        params: params
      }).then((res) => {
        if (res.data.PROFILE.length !== 0) {
          this.profile = true
          this.profileList = res.data.PROFILE[0]
          // 메달 세팅
          if (this.profileList.award_medal.length !== 0) {
            this.profileAward.expose = 0
            this.profileAward.totalPage = Math.floor(this.profileList.award_medal.length / 5) + 1
            if (this.profileAward.totalPage > 1) {
              this.profileAward.nextActive = true
            }
          }
        }
      })
    },
    goStation (szUserId) {
      let oNewWindow = window.open('about:blank')
      oNewWindow.location.href = '//bj.afreecatv.com/' + szUserId
    },
    goNotice (szUserId, nStationNo, nBbsNo, nTitleNo) {
      let oNewWindow = window.open('about:blank')
      oNewWindow.location.href = '//live.afreecatv.com:8079/app/index.cgi?szBoard=read_bbs&szBjId=' + szUserId + '&nStationNo=' + nStationNo + '&nBbsNo=' + nBbsNo + '&nTitleNo=' + nTitleNo
    },
    goProfile () {
      let oNewWindow = window.open('about:blank')
      oNewWindow.location.href = '//bj.afreecatv.com/' + 'pjh08190819' + '/setting/profile'
    },
    awardListToggle () {
      this.profileAward.expose = !this.profileAward.expose

      let oAwardList = document.getElementById('divAwardList')
      if (this.profileAward.expose) {
        oAwardList.style.display = 'block'
      } else {
        oAwardList.style.display = 'none'
      }
    },
    awardContentSetting () {
      this.profileContent = {
        1: {
          class: 'md_partner',
          content: '파트너BJ'
        },
        2: {
          class: 'md_best',
          content: '베스트BJ'
        },
        3: {
          class: 'md_a2013',
          content: '방송대상2013'
        },
        4: {
          class: 'md_a2012',
          content: '방송대상2012'
        },
        5: {
          class: 'md_a2011',
          content: '방송대상2011'
        },
        6: {
          class: 'md_progamer',
          content: '프로게이머BJ'
        },
        7: {
          class: 'md_sports',
          content: '스포츠BJ'
        },
        8: {
          class: 'md_a2011',
          content: '방송대상2011'
        },
        9: {
          class: 'md_mgame',
          content: '모바일게임BJ'
        },
        10: {
          class: 'md_clan10',
          content: '클랜TOP10'
        },
        11: {
          class: 'md_game_god',
          content: '게임신BJ'
        },
        12: {
          class: 'md_angel',
          content: '엔젤BJ'
        },
        13: {
          class: 'md_a2014',
          content: '방송대상2014'
        },
        14: {
          class: 'md_a2015',
          content: '2015 BJ대상'
        },
        15: {
          class: 'md_a2016',
          content: '2016 BJ대상'
        },
        16: {
          class: 'md_tech',
          content: 'Tech BJ'
        },
        17: {
          class: 'md_a2017',
          content: '2017 BJ대상'
        },
        18: {
          class: 'md_a2018',
          content: '2018 BJ대상'
        },
        19: {
          class: 'md_a2019',
          content: '2019 BJ대상'
        },
        20: {
          class: 'md_muse',
          content: '2019 뮤즈BJ'
        }
      }
    },
    moveMedal (szAction) {
      if (szAction === 'prev') {
        if (this.profileAward.curPage > 1) {
          this.profileAward.start -= 5
          this.profileAward.end -= 5
          this.profileAward.curPage -= 1
          this.profileAward.nextActive = true
          if (this.profileAward.curPage === 1) {
            this.profileAward.prevActive = false
          }
        }
      } else if (szAction === 'next') {
        if (this.profileAward.curPage < this.profileAward.totalPage) {
          this.profileAward.start += 5
          this.profileAward.end += 5
          this.profileAward.curPage += 1
          this.profileAward.prevActive = true
          if (this.profileAward.curPage === this.profileAward.totalPage) {
            this.profileAward.nextActive = false
          }
        }
      }
    },
    favorite (szBjId) {
      console.log('추가예정')
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
.fade-enter-active, .fade-leave-active {
  transition: opacity .5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}
</style>
